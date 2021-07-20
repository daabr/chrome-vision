// The devtools program is an example for interacting with Google Search,
// using Chrome Vision's Chrome DevTools Protocol (CDP) API.
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/daabr/chrome-vision/pkg/devtools"
	"github.com/daabr/chrome-vision/pkg/devtools/dom"
	"github.com/daabr/chrome-vision/pkg/devtools/input"
	"github.com/daabr/chrome-vision/pkg/devtools/page"
)

func main() {
	// Initialize the browsing session.
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	flags := devtools.DefaultBrowserFlags()
	flags["window-size"] = "1280,720"
	delete(flags, "headless")

	ctx, err := devtools.NewContext(ctx, devtools.BrowserFlags(flags))
	if err != nil {
		log.Fatalf("initialization error: %v", err)
	}
	defer func() {
		devtools.Close(ctx)
		if session, ok := devtools.FromContext(ctx); ok {
			os.RemoveAll(session.OutputDir)
		}
	}()

	// Prepare to receive page lifecycle events from the browser.
	ch, err := devtools.SubscribeEvent(ctx, "Page.lifecycleEvent")
	if err != nil {
		log.Fatalf("event subscription error: %v", err)
	}
	defer close(ch)

	// Interact with the browsing session.
	if err := navigate(ctx, ch); err != nil {
		log.Fatalf("navigation error: %v", err)
	}
	if err := search(ctx, ch); err != nil {
		log.Fatalf("failed to type search query: %v", err)
	}
	if err := imageResults(ctx, ch); err != nil {
		log.Fatalf("failed to switch to image results: %v", err)
	}
	if err := scrollDown(ctx, ch); err != nil {
		log.Fatalf("failed to scroll down: %v", err)
	}
	if err := lastResult(ctx); err != nil {
		log.Fatalf("failed to get the last result: %v", err)
	}
}

// Navigate to Google Search.
func navigate(ctx context.Context, ch <-chan *devtools.Message) error {
	nav := page.NewNavigate("https://google.com/webhp?hl=en&pws=0")
	if _, err := nav.Do(ctx); err != nil {
		return err
	}
	return waitUntilStable(ch)
}

// Wait until there are no page lifecycle events for 2 seconds.
func waitUntilStable(ch <-chan *devtools.Message) error {
	d := 2 * time.Second
	t := time.NewTimer(d)
	for {
		select {
		case msg := <-ch:
			e := &page.LifecycleEvent{}
			if err := json.Unmarshal(msg.Params, e); err != nil {
				return fmt.Errorf("JSON event parsing error: %v", err)
			}
			if !t.Stop() {
				<-t.C
			}
			t.Reset(d)

		case <-t.C:
			return nil
		}
	}
}

// Type the search query "kittens", and press the Enter key.
func search(ctx context.Context, ch <-chan *devtools.Message) error {
	down := input.NewDispatchKeyEvent("keyDown")
	up := input.NewDispatchKeyEvent("keyUp")
	for _, key := range "kittens\r" {
		if err := down.SetText(string(key)).Do(ctx); err != nil {
			return err
		}
		// These `time.Sleep` function calls are not "load-bearing", i.e. they
		// are not required for asynchronous event handling or safety, they
		// simply simulate human latency - feel free to remove them.
		time.Sleep(50 * time.Millisecond)
		if err := up.Do(ctx); err != nil {
			return err
		}
		time.Sleep(50 * time.Millisecond)
	}
	return waitUntilStable(ch)
}

// Click the "Images" tab to show image search results.
func imageResults(ctx context.Context, ch <-chan *devtools.Message) error {
	doc, err := dom.NewGetDocument().Do(ctx)
	if err != nil {
		return err
	}
	links, err := dom.NewQuerySelectorAll(doc.Root.NodeID, "a").Do(ctx)
	if err != nil {
		return err
	}
	clicked := false
	for _, nodeID := range links.NodeIds {
		if s, err := dom.NewGetOuterHTML().SetNodeID(nodeID).Do(ctx); err == nil {
			if strings.Contains(s.OuterHTML, ">Images<") {
				if box, err := dom.NewGetBoxModel().SetNodeID(nodeID).Do(ctx); err == nil {
					// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-Quad
					x := (box.Model.Content[0] + box.Model.Content[2]) / 2
					y := (box.Model.Content[3] + box.Model.Content[5]) / 2

					mouse := input.NewDispatchMouseEvent("mousePressed", x, y)
					mouse = mouse.SetButton(input.MouseButtonLeft)
					if err := mouse.SetClickCount(1).Do(ctx); err != nil {
						return err
					}
					mouse.Type = "mouseReleased"
					if err := mouse.Do(ctx); err != nil {
						return err
					}
					clicked = true
					break
				}
			}
		}
	}
	if !clicked {
		return errors.New("image search results tab not found")
	}
	return waitUntilStable(ch)
}

// Scroll down to the bottom of the page.
func scrollDown(ctx context.Context, ch <-chan *devtools.Message) error {
	doc, err := dom.NewGetDocument().Do(ctx)
	if err != nil {
		return err
	}
	footer, err := dom.NewQuerySelector(doc.Root.NodeID, "footer").Do(ctx)
	if err != nil {
		return err
	}
	if footer.NodeID == 0 {
		return errors.New("image search results footer not found")
	}
	box := dom.NewGetBoxModel().SetNodeID(footer.NodeID)

	down := input.NewDispatchKeyEvent("keyDown")
	down = down.SetNativeVirtualKeyCode(34).SetWindowsVirtualKeyCode(34)
	up := input.NewDispatchKeyEvent("keyUp")

	// Continue to press page-down three more times after the footer
	// becomes "visible", to ensure it's really visible to the user.
	bottomIsVisible := 0
	for bottomIsVisible < 3 {
		if err := down.Do(ctx); err != nil {
			return err
		}
		if err := up.Do(ctx); err != nil {
			return err
		}
		// This `time.Sleep` function call is not "load-bearing", i.e. it's
		// not required for asynchronous event handling or safety, it simply
		// simulates human latency - feel free to remove it.
		time.Sleep(500 * time.Millisecond)
		if _, err := box.Do(ctx); err == nil {
			bottomIsVisible++
		}
	}
	return waitUntilStable(ch)
}

// Print the title and URL of the last result.
func lastResult(ctx context.Context) error {
	doc, err := dom.NewGetDocument().Do(ctx)
	if err != nil {
		return err
	}
	s := "div#islrg > div.islrc > div:last-child > a:last-child"
	last, err := dom.NewQuerySelector(doc.Root.NodeID, s).Do(ctx)
	if err != nil {
		return err
	}
	a, err := dom.NewGetAttributes(last.NodeID).Do(ctx)
	if err != nil {
		return err
	}
	attr := make(map[string]string)
	for i := 0; i < len(a.Attributes); i += 2 {
		attr[a.Attributes[i]] = a.Attributes[i+1]
	}
	log.Print(attr["title"])
	log.Print(attr["href"])
	return nil
}
