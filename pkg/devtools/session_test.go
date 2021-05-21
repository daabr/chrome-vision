package devtools_test

import (
	"context"
	"log"
	"os"
	"path"
	"testing"
	"time"

	"github.com/daabr/chrome-vision/pkg/devtools"
)

func ExampleNewContext() {
	// Set an idiomatic timeout (or deadline) for the entire session,
	// i.e. the browser will be killed immediately when the time comes.
	ctx1, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Customize the browser command-line flags, before starting it.
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

func TestBrowserFlags(t *testing.T) {
	// Set up.
	dir, err := os.MkdirTemp("", "")
	if err != nil {
		t.Fatalf(`os.MkdirTemp("", ""); got error: %v`, err)
	}
	defer func() {
		os.RemoveAll(dir)
	}()
	os.Setenv(devtools.OutputRootEnv, dir)
	defer os.Unsetenv(devtools.OutputRootEnv)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	flags := devtools.DefaultBrowserFlags()
	flags["disable-gpu"] = true

	// Test.
	_, err = devtools.NewContext(ctx, devtools.BrowserFlags(flags))
	if err != nil {
		t.Errorf("devtools.NewContext(ctx, flags); got error: %v", err)
	}
}

func TestFromContext(t *testing.T) {
	// Set up.
	dir, err := os.MkdirTemp("", "")
	if err != nil {
		t.Fatalf(`os.MkdirTemp("", ""); got error: %v`, err)
	}
	defer func() {
		os.RemoveAll(dir)
	}()
	os.Setenv(devtools.OutputRootEnv, dir)
	defer os.Unsetenv(devtools.OutputRootEnv)

	ctx, err := devtools.NewContext(context.Background())
	if err != nil {
		t.Fatalf("devtools.NewContext(ctx); got error: %v", err)
	}
	defer devtools.Cancel(ctx)

	// Test.
	session, ok := devtools.FromContext(ctx)
	if !ok {
		t.Fatalf("devtools.FromContext(ctx); ok = %v, want %v", ok, !ok)
	}
	if session.OutputDir == "" {
		t.Error(`session.OutputDir = "", want != ""`)
	}
	if _, err := os.Stat(session.OutputDir); err != nil {
		t.Errorf("os.Stat(session.OutputDir); got error: %v", err)
	}
	if session.UserDataDir == "" {
		t.Error(`session.UserDataDir = "", want != ""`)
	}
	if _, err := os.Stat(session.UserDataDir); err != nil {
		t.Errorf("os.Stat(session.UserDataDir); got error: %v", err)
	}
	if session.TargetID.Read() == "" {
		t.Error(`session.TargetID.Read() = "", want != ""`)
	}
	if session.SessionID.Read() == "" {
		t.Error(`session.SessionID.Read() = "", want != ""`)
	}
}

func TestUserDataDir(t *testing.T) {
	// Set up.
	dir, err := os.MkdirTemp("", "")
	if err != nil {
		t.Fatalf(`os.MkdirTemp("", ""); got error: %v`, err)
	}
	defer func() {
		os.RemoveAll(dir)
	}()
	os.Setenv(devtools.OutputRootEnv, dir)
	defer os.Unsetenv(devtools.OutputRootEnv)

	userDir := path.Join(dir, "userDir")

	ctx, err := devtools.NewContext(context.Background(), devtools.UserDataDir(userDir))
	if err != nil {
		t.Fatalf("devtools.NewContext(ctx, userDir); got error: %v", err)
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
	if session.UserDataDir != userDir {
		t.Errorf("session.UserDataDir = %q, want %q", session.UserDataDir, userDir)
	}
	if _, err := os.Stat(session.UserDataDir); err != nil {
		t.Errorf("os.Stat(session.UserDataDir); got error: %v", err)
	}
}
