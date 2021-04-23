package cdp_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/daabr/chrome-vision/pkg/cdp"
)

func TestCancel(t *testing.T) {
	// Set up.
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	ctx, err := cdp.NewContext(ctx)
	if err != nil {
		t.Fatalf("cdp.NewContext(ctx); got error: %s", err.Error())
	}
	defer func() {
		// Tear down.
		if session, ok := cdp.FromContext(ctx); ok {
			os.RemoveAll(session.OutputDir)
		}
	}()

	// Test.
	cdp.Cancel(ctx)
	cdp.Wait(ctx)
	if ctx.Err() == nil {
		t.Error("cdp.Cancel(ctx); ctx.Err() = nil, want !nil")
	}
}

func TestClose(t *testing.T) {
	// Set up.
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	ctx, err := cdp.NewContext(ctx)
	if err != nil {
		t.Fatalf("cdp.NewContext(ctx); got error: %s", err.Error())
	}
	defer func() {
		// Tear down.
		if session, ok := cdp.FromContext(ctx); ok {
			os.RemoveAll(session.OutputDir)
		}
	}()

	// Test.
	cdp.Close(ctx)
	if ctx.Err() == nil {
		t.Error("cdp.Close(ctx); ctx.Err() = nil, want !nil")
	}
}