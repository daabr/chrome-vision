// The sync program is an example for running Chrome DevTools Protocol (CDP)
// commands in a synchronous (blocking) mode, which is more straightforward
// than the asynchronous (non-blocking) mode.
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

	// Prepare to receive navigation events from the browser (before calling the
	// navigation command, so we won't lose any events due to a race condition).
	ch, err := devtools.SubscribeEvent(ctx, "Page.frameNavigated")
	if err != nil {
		log.Fatalf("event subscription error: %v", err)
	}
	defer close(ch)

	// Navigate to Amazon's homepage.
	n := page.NewNavigate("https://amazon.com/")
	if _, err := n.Do(ctx); err != nil {
		log.Fatalf("navigation error: %v", err)
	}

	// Wait until the page is really loaded (though not necessarily stable).
	msg := <-ch
	e := &page.FrameNavigated{}
	if err := json.Unmarshal(msg.Params, e); err != nil {
		log.Fatalf("JSON event parsing error: %v", err)
	}
	log.Printf("Page.frameNavigated event: %#v", e)
}
