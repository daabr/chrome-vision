package websocket_test

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/daabr/chrome-vision/pkg/websocket"
)

func expectedKey(r *http.Request) string {
	h := sha1.New()
	h.Write([]byte(r.Header.Get("Sec-WebSocket-Key")))
	h.Write([]byte("258EAFA5-E914-47DA-95CA-C5AB0DC85B11"))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func subTestHandshake(f func(http.ResponseWriter, *http.Request)) func(t *testing.T) {
	return func(t *testing.T) {
		// Start a test server.
		ts := httptest.NewServer(http.HandlerFunc(f))
		defer ts.Close()

		// Derive WebSocket handshake parameters from the test server's URL.
		address := strings.TrimPrefix(ts.URL, "http://")
		path := "/devtools/browser/01234567-89ab-cdef-0123-456789abcdef"

		// Test.
		_, err := websocket.Handshake(context.Background(), address, path)
		if err == nil {
			t.Error("Handshake() = Conn, want error")
		}
	}
}

func TestHandshakeExpectedErrors(t *testing.T) {
	t.Run("incorrect status code", subTestHandshake(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Upgrade", "websocket")
		w.Header().Add("Connection", "Upgrade")
		w.Header().Add("Sec-WebSocket-Accept", expectedKey(r))
		w.WriteHeader(http.StatusOK)
	}))
	t.Run("incorrect upgrade header", subTestHandshake(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Upgrade", "FOO")
		w.Header().Add("Connection", "Upgrade")
		w.Header().Add("Sec-WebSocket-Accept", expectedKey(r))
		w.WriteHeader(http.StatusSwitchingProtocols)
	}))
	t.Run("incorrect connection header", subTestHandshake(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Upgrade", "websocket")
		w.Header().Add("Connection", "BAR")
		w.Header().Add("Sec-WebSocket-Accept", expectedKey(r))
		w.WriteHeader(http.StatusSwitchingProtocols)
	}))
	t.Run("incorrect accept header", subTestHandshake(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Upgrade", "websocket")
		w.Header().Add("Connection", "Upgrade")
		w.Header().Add("Sec-WebSocket-Accept", "BAZ")
		w.WriteHeader(http.StatusSwitchingProtocols)
	}))
	t.Run("missing upgrade header", subTestHandshake(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Connection", "Upgrade")
		w.Header().Add("Sec-WebSocket-Accept", expectedKey(r))
		w.WriteHeader(http.StatusSwitchingProtocols)
	}))
	t.Run("missing connection header", subTestHandshake(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Upgrade", "websocket")
		w.Header().Add("Sec-WebSocket-Accept", expectedKey(r))
		w.WriteHeader(http.StatusSwitchingProtocols)
	}))
	t.Run("missing accept header", subTestHandshake(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Upgrade", "websocket")
		w.Header().Add("Connection", "Upgrade")
		w.WriteHeader(http.StatusSwitchingProtocols)
	}))
}

func TestHandshakeUnexpectedHeader(t *testing.T) {
	// Start a test server.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Upgrade", "websocket")
		w.Header().Add("Connection", "Upgrade")
		w.Header().Add("Sec-WebSocket-Accept", expectedKey(r))
		w.Header().Add("Foo", "Bar")
		w.WriteHeader(http.StatusSwitchingProtocols)
	}))
	defer ts.Close()

	// Derive WebSocket handshake parameters from the test server's URL.
	address := strings.TrimPrefix(ts.URL, "http://")
	path := "/devtools/browser/01234567-89ab-cdef-0123-456789abcdef"

	// Test.
	if _, err := websocket.Handshake(context.Background(), address, path); err != nil {
		t.Errorf("Handshake(); got unexpected error: %v", err)
	}
}
