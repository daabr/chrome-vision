package devtools_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/daabr/chrome-vision/pkg/devtools"
)

func TestCancel(t *testing.T) {
	// Set up.
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	ctx, err := devtools.NewContext(ctx)
	if err != nil {
		t.Fatalf("devtools.NewContext(ctx); got error: %s", err.Error())
	}
	defer func() {
		// Tear down.
		if session, ok := devtools.FromContext(ctx); ok {
			os.RemoveAll(session.OutputDir)
		}
	}()

	// Test.
	devtools.Cancel(ctx)
	devtools.Wait(ctx)
	if ctx.Err() == nil {
		t.Error("devtools.Cancel(ctx); ctx.Err() = nil, want !nil")
	}
}

func TestClose(t *testing.T) {
	// Set up.
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	ctx, err := devtools.NewContext(ctx)
	if err != nil {
		t.Fatalf("devtools.NewContext(ctx); got error: %s", err.Error())
	}
	defer func() {
		// Tear down.
		if session, ok := devtools.FromContext(ctx); ok {
			os.RemoveAll(session.OutputDir)
		}
	}()

	// Test.
	devtools.Close(ctx)
	if ctx.Err() == nil {
		t.Error("devtools.Close(ctx); ctx.Err() = nil, want !nil")
	}
}
