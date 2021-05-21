// The init program contains several examples of initialization and
// customization of browser sessions in Chrome Vision.
package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/daabr/chrome-vision/pkg/devtools"
	"github.com/daabr/chrome-vision/pkg/devtools/page"
	"github.com/daabr/chrome-vision/pkg/devtools/target"
)

func main() {
	ExampleMinimalSession()
	ExampleSessionTimeout()
	ExampleBrowserCustomizations()
	ExamplehMultipleBrowserTabs()
	ExampleCleanupAfterSession()
	ExampleCustomOutputDir()
}

func ExampleMinimalSession() {
	// Start a new browser.
	ctx, err := devtools.NewContext(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// Kill the browser forcefully. If you want to close it gracefully,
	// call the `devtools.Close` function instead.
	defer devtools.Cancel(ctx)

	// Do stuff...

	// Optional (because of the `defer` above): send the browser a command
	// to close itself gracefully, and wait until this is done.
	devtools.Close(ctx)
}

func ExampleSessionTimeout() {
	// Set an idiomatic timeout (or deadline) for the entire session,
	// i.e. the browser will be killed immediately when the time comes.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Start a new browser.
	ctx, err := devtools.NewContext(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// No need to call `defer devtools.Cancel(ctx)` because `defer cancel()`
	// above will propagate the cancellation to its descendant - the browsing
	// session context.

	// Do stuff...

	// Optional (because of the `defer` above): send the browser a command
	// to close itself gracefully, and wait until this is done.
	devtools.Close(ctx)
}

func ExampleBrowserCustomizations() {
	// Customize the browser command-line flags, before starting it.
	flags := devtools.DefaultBrowserFlags()
	flags["disable-gpu"] = true // https://crbug.com/765284
	flags["window-size"] = "1920,1080"
	delete(flags, "headless")

	// The `devtools.NewContext` function supports 0 or more customizations.
	// Other options to consider: `devtools.BrowserPath` (to run a custom
	// binary from a specific location) and `devtools.UserDataDir` (to use
	// an existing non-temporary user data directory - useful if you want
	// non-default user settings or user data in the browsing session).
	ctx, err := devtools.NewContext(context.Background(), devtools.BrowserFlags(flags))
	if err != nil {
		log.Fatal(err)
	}
	defer devtools.Cancel(ctx)

	// Do stuff...

	devtools.Close(ctx)
}

func ExamplehMultipleBrowserTabs() {
	// Chrome doesn't allow using tabs in headless mode.
	flags := devtools.DefaultBrowserFlags()
	flags["window-size"] = "1920,1080"
	delete(flags, "headless")

	// Start a new browser (first tab).
	ctx1, err := devtools.NewContext(context.Background(), devtools.BrowserFlags(flags))
	if err != nil {
		log.Fatal(err)
	}
	defer devtools.Cancel(ctx1)

	// Do stuff in the first tab...
	nav := page.NewNavigate("https://en.wikipedia.org/wiki/1")
	if _, err := nav.Do(ctx1); err != nil {
		log.Fatal(err)
	}

	// Open a second tab in the existing browser.
	ctx2, err := devtools.NewContext(ctx1)
	if err != nil {
		log.Fatal(err)
	}
	// No need to call `defer devtools.Cancel(ctx2)` because the `defer` above
	// will propagate the cancellation of the main browsing session context
	// (ctx1) to its descendant - the second tab's context (ctx2).

	// Do stuff in the second tab...
	nav = page.NewNavigate("https://en.wikipedia.org/wiki/2")
	if _, err := nav.Do(ctx2); err != nil {
		log.Fatal(err)
	}

	// Do stuff in the first tab again...
	if session, ok := devtools.FromContext(ctx1); ok {
		focus := target.NewActivateTarget(session.TargetID.Read())
		if err := focus.Do(ctx1); err != nil {
			log.Fatal(err)
		}
	}
	nav = page.NewNavigate("https://en.wikipedia.org/wiki/3")
	if _, err := nav.Do(ctx1); err != nil {
		log.Fatal(err)
	}
}

func ExampleCleanupAfterSession() {
	// Start a new browser.
	ctx, err := devtools.NewContext(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if session, ok := devtools.FromContext(ctx); ok {
			log.Printf("Session output directory: %s", session.OutputDir)
			os.RemoveAll(session.OutputDir)
			log.Print("Delorted!")
		}
	}()
	// Defer the `Cancel` call AFTER the cleanup function
	// (https://tour.golang.org/flowcontrol/13).
	defer devtools.Cancel(ctx)

	// Do stuff...
}

func ExampleCustomOutputDir() {
	customDir, err := os.MkdirTemp("", "foo_*")
	if err != nil {
		log.Fatalf(`os.MkdirTemp("", "foo_*"); got unexpected error: %v`, err)
	}
	defer func() {
		os.RemoveAll(customDir)
	}()

	// We create an output directory per browser process, which contains logs,
	// user data, screenshots, etc. By default, this directory is created under
	// Go's `os.TempDir()` - see https://golang.org/pkg/os/#TempDir). The
	// optional environment variable "CDP_OUTPUT_ROOT" overrides this path.
	os.Setenv(devtools.OutputRootEnv, customDir)
	defer os.Unsetenv(devtools.OutputRootEnv)

	// Start a new browser.
	ctx, err := devtools.NewContext(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer devtools.Cancel(ctx)

	// Do stuff...

	devtools.Close(ctx)
}
