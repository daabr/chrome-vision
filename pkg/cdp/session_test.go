package cdp_test

import (
	"context"
	"log"
	"time"

	"github.com/daabr/chrome-vision/pkg/cdp"
)

func ExampleNewContext() {
	ctx1 := context.Background()
	ctx1, cancel := context.WithTimeout(ctx1, 2*time.Second)
	defer cancel()

	// Start a new browser.
	ctx2, err := cdp.NewContext(ctx1)
	if err != nil {
		log.Fatal(err)
	}

	// Open a new tab in the existing browser.
	ctx3, err := cdp.NewContext(ctx2)
	if err != nil {
		log.Fatal(err)
	}

	// Close the browser.
	cdp.Close(ctx3)
}

func ExampleBrowserFlags() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	// Customize the browser flags, before starting the browser.
	flags := cdp.DefaultBrowserFlags()
	flags["disable-gpu"] = true // https://crbug.com/765284
	flags["window-size"] = "1920,1080"
	delete(flags, "headless")

	ctx, err := cdp.NewContext(ctx, cdp.BrowserFlags(flags))
	if err != nil {
		log.Fatal(err)
	}

	cdp.Close(ctx)
}
