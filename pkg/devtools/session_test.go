package devtools_test

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/daabr/chrome-vision/pkg/devtools"
)

func ExampleNewContext() {
	ctx1 := context.Background()
	ctx1, cancel := context.WithTimeout(ctx1, 2*time.Second)
	defer cancel()

	// Start a new browser.
	ctx2, err := devtools.NewContext(ctx1)
	if err != nil {
		log.Fatal(err)
	}

	// Open a new tab in the existing browser.
	ctx3, err := devtools.NewContext(ctx2)
	if err != nil {
		log.Fatal(err)
	}

	// Close the browser.
	devtools.Close(ctx3)
}

func ExampleBrowserFlags() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	// Customize the browser flags, before starting the browser.
	flags := devtools.DefaultBrowserFlags()
	flags["disable-gpu"] = true // https://crbug.com/765284
	flags["window-size"] = "1920,1080"
	delete(flags, "headless")

	ctx, err := devtools.NewContext(ctx, devtools.BrowserFlags(flags))
	if err != nil {
		log.Fatal(err)
	}

	devtools.Close(ctx)
}

func TestFromContext(t *testing.T) {
	// Set up.
	parentCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	ctx, err := devtools.NewContext(parentCtx)
	if err != nil {
		t.Fatalf("devtools.NewContext(ctx); got error: %s", err.Error())
	}
	defer func() {
		// Tear down.
		devtools.Cancel(ctx)
		devtools.Wait(ctx)
		if session, ok := devtools.FromContext(ctx); ok {
			os.RemoveAll(session.OutputDir)
		}
	}()

	// Test.
	session, ok := devtools.FromContext(ctx)
	if !ok {
		t.Fatalf("devtools.FromContext(ctx); ok = %v, want %v", ok, !ok)
	}
	if session.OutputDir == "" {
		t.Error(`session.OutputDir = "", want != ""`)
	}
	if _, err := os.Stat(session.OutputDir); err != nil {
		t.Errorf("os.Stat(session.OutputDir); got error: %s", err.Error())
	}
	if session.UserDataDir == "" {
		t.Error(`session.UserDataDir = "", want != ""`)
	}
	if _, err := os.Stat(session.UserDataDir); err != nil {
		t.Errorf("os.Stat(session.UserDataDir); got error: %s", err.Error())
	}
}

func TestUserDataDir(t *testing.T) {
	// Set up.
	parentCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	dir, err := os.MkdirTemp("", "")
	if err != nil {
		t.Fatalf("os.MkdirTemp(%q, %q); got error: %s", "", "", err.Error())
	}
	defer func() {
		os.RemoveAll(dir)
	}()

	ctx, err := devtools.NewContext(parentCtx, devtools.UserDataDir(dir))
	if err != nil {
		t.Fatalf("devtools.NewContext(ctx); got error: %s", err.Error())
	}
	session, ok := devtools.FromContext(ctx)
	if !ok {
		t.Fatalf("devtools.FromContext(ctx); ok = %v, want %v", ok, !ok)
	}
	defer func() {
		// Tear down.
		devtools.Cancel(ctx)
		devtools.Wait(ctx)
		os.RemoveAll(session.OutputDir)
	}()

	// Test.
	if session.UserDataDir == "" {
		t.Fatalf(`session.UserDataDir = "", want != ""`)
	}
	if session.UserDataDir != dir {
		t.Errorf("session.UserDataDir = %v, want %v", session.UserDataDir, dir)
	}
}
