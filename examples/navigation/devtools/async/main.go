// The async program is an example for running Chrome DevTools Protocol (CDP)
// commands in an asynchronous (non-blocking) mode, which enables more
// advanced logic than the synchronous (blocking) mode.
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/daabr/chrome-vision/pkg/devtools"
	"github.com/daabr/chrome-vision/pkg/devtools/page"
)

func main() {
	// Initialize the browsing session.
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	ctx, err := devtools.NewContext(ctx)
	if err != nil {
		log.Fatalf("initialization error: %v", err)
	}
	defer func() {
		devtools.Close(ctx)
		if session, ok := devtools.FromContext(ctx); ok {
			os.RemoveAll(session.OutputDir)
		}
	}()

	// Prepare to receive a couple of events from the browser.
	fnEventChan, err := devtools.SubscribeEvent(ctx, "Page.frameNavigated")
	if err != nil {
		log.Fatalf("event subscription error: %v", err)
	}
	defer close(fnEventChan)

	leEventChan, err := devtools.SubscribeEvent(ctx, "Page.lifecycleEvent")
	if err != nil {
		log.Fatalf("event subscription error: %v", err)
	}
	defer close(leEventChan)

	// Navigate to Amazon's homepage.
	n := page.NewNavigate("https://amazon.com/")
	navCmdChan, err := n.Start(ctx)
	if err != nil {
		log.Fatalf("navigation command error: %v", err)
	}
	defer close(navCmdChan)

	// Wait for multiple conditions to determine that page loading as really done
	// (per https://tour.golang.org/concurrency/5, the order of the conditions
	// doesn't matter).
	gotNavResponse, gotNavEvent, isStable := false, false, false
	d := 3 * time.Second
	t := time.NewTimer(d)
	for !gotNavResponse || !gotNavEvent || !isStable {
		select {
		// Condition 1: the `page.Navigate` command returns a response.
		case msg := <-navCmdChan:
			response, err := n.ParseResponse(msg)
			if err != nil {
				log.Fatalf("navigation response error: %v", err)
			}
			log.Printf("navigation command is done: %#v", response)
			gotNavResponse = true

		// Condition 2: we received a `Page.frameNavigated` event.
		case msg := <-fnEventChan:
			fn := &page.FrameNavigated{}
			if err := json.Unmarshal(msg.Params, fn); err != nil {
				log.Fatalf("JSON event parsing error: %v", err)
			}
			log.Printf("Page.frameNavigated event: %#v", fn)
			gotNavEvent = true

		// Condition 3: no page lifecycle events for 3 seconds.
		case msg := <-leEventChan:
			e := &page.LifecycleEvent{}
			if err := json.Unmarshal(msg.Params, e); err != nil {
				log.Fatalf("JSON event parsing error: %v", err)
			}
			log.Printf("frame ID %q: lifecycle event %q", e.FrameID, e.Name)
			if !t.Stop() {
				<-t.C
			}
			t.Reset(d)
			isStable = false

		case <-t.C:
			log.Print("a few seconds passed without any lifecycle events")
			isStable = true
		}

		// TODO: wait for title? wait for a visible element?
	}
}
