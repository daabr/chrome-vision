package cdp_test

import (
	"context"
	"log"
	"os"
	"testing"
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

func TestFromContext(t *testing.T) {
	// Set up.
	parentCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	ctx, err := cdp.NewContext(parentCtx)
	if err != nil {
		t.Fatalf("cdp.NewContext(ctx); got error: %s", err.Error())
	}
	defer func() {
		// Tear down.
		cdp.Cancel(ctx)
		cdp.Wait(ctx)
		if session, ok := cdp.FromContext(ctx); ok {
			if session.OutputDir != nil {
				os.RemoveAll(*session.OutputDir)
			}
		}
	}()

	// Test.
	session, ok := cdp.FromContext(ctx)
	if !ok {
		t.Fatalf("cdp.FromContext(ctx); ok = %v, want %v", ok, !ok)
	}
	if session.OutputDir == nil {
		t.Error("session.OutputDir = nil, want !nil")
	}
	if _, err := os.Stat(*session.OutputDir); err != nil {
		t.Errorf("os.Stat(*session.OutputDir); got error: %s", err.Error())
	}
	if session.UserDataDir == nil {
		t.Error("session.UserDataDir = nil, want !nil")
	}
	if _, err := os.Stat(*session.UserDataDir); err != nil {
		t.Errorf("os.Stat(*session.UserDataDir); got error: %s", err.Error())
	}
	t.Logf("%#v", session)
}

func TestUserDataDir(t *testing.T) {
	// Set up.
	parentCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	dir := os.TempDir()
	defer func() {
		os.RemoveAll(dir)
	}()

	ctx, err := cdp.NewContext(parentCtx, cdp.UserDataDir(dir))
	if err != nil {
		t.Fatalf("cdp.NewContext(ctx); got error: %s", err.Error())
	}
	session, ok := cdp.FromContext(ctx)
	if !ok {
		t.Fatalf("cdp.FromContext(ctx); ok = %v, want %v", ok, !ok)
	}
	defer func() {
		// Tear down.
		cdp.Cancel(ctx)
		cdp.Wait(ctx)
		if session.OutputDir != nil {
			os.RemoveAll(*session.OutputDir)
		}
	}()

	// Test.
	if session.UserDataDir == nil {
		t.Fatalf("session.UserDataDir = nil, want !nil")
	}
	if *session.UserDataDir != dir {
		t.Errorf("*session.UserDataDir = %v, want %v", *session.UserDataDir, dir)
	}
	t.Logf("%#v", session)
}
