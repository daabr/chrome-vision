package websocket

import (
	"bufio"
	"context"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

const (
	dialTimeout   = 5 * time.Second
	websocketGUID = "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"
)

// Conn defines a WebSocket connection with a buffered I/O interface.
type Conn struct {
	nc net.Conn
	rw *bufio.ReadWriter
}

// NewConn initializes a WebSocket connection based on an open TCP connection.
func NewConn(c net.Conn) *Conn {
	rw := bufio.NewReadWriter(bufio.NewReader(c), bufio.NewWriter(c))
	return &Conn{nc: c, rw: rw}
}

// Handshake initiates a WebSocket connection with "ws://addr/path",
// based on https://datatracker.ietf.org/doc/html/rfc6455#section-4.1.
func Handshake(ctx context.Context, addr, path string) (*Conn, error) {
	// Network connection.
	d := net.Dialer{Timeout: dialTimeout}
	netConn, err := d.DialContext(ctx, "tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to dial to browser: %v", err)
	}
	wsConn := NewConn(netConn)

	// Upgrade request.
	nonce, err := nonce()
	if err != nil {
		return nil, fmt.Errorf("failed to generate a nonce: %v", err)
	}
	sendUpgradeRequest(wsConn, addr, path, nonce)

	// Upgrade response.
	err = receiveUpgradeResponse(wsConn, nonce)
	if err != nil {
		return nil, fmt.Errorf("failed to receive WebSocket upgrade response: %v", err)
	}
	return wsConn, nil
}

// Generate a nonce consisting of a randomly selected 16-byte value
// that has been Base64-encoded.
func nonce() (string, error) {
	b := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func sendUpgradeRequest(c *Conn, addr, path, nonce string) {
	// The handshake MUST be a valid HTTP request as specified by [RFC2616].
	// The method of the request MUST be GET, and the HTTP version MUST	be at least 1.1.
	// The "Request-URI" part of the request MUST match the /resource name/ defined in
	// Section 3 (a relative URI) or be an absolute	HTTP/HTTPS URI that, when parsed,
	// has a /resource name/, /host/, and /port/ that match the corresponding ws/wss URI.
	fmt.Fprintf(c.nc, "GET %s HTTP/1.1\r\n", path)
	// The request MUST contain a |Host| header field whose value contains /host/ plus
	// optionally ":" followed by /port/ (when not using the default port).
	fmt.Fprintf(c.nc, "Host: %s\r\n", addr)
	// The request MUST contain an |Upgrade| header field whose value MUST include the
	// "websocket" keyword.
	fmt.Fprint(c.nc, "Upgrade: websocket\r\n")
	// The request MUST contain a |Connection| header field whose value MUST include the
	// "Upgrade" token.
	fmt.Fprint(c.nc, "Connection: Upgrade\r\n")
	// The request MUST include a header field with the name |Sec-WebSocket-Key|.
	// The value of this header field MUST be a nonce consisting of a randomly selected
	// 16-byte value that has been Base64-encoded (see Section 4 of [RFC4648]).
	// The nonce MUST be selected randomly for each connection.
	fmt.Fprintf(c.nc, "Sec-WebSocket-Key: %s\r\n", nonce)

	// The request MUST include a header field with the name |Origin| [RFC6454] if the
	// request is coming from a browser client. If the connection is from a non-browser
	// client, the request MAY include this header field if the semantics of that client
	// match the use-case described here for browser clients. The value of this header
	// field is the ASCII serialization of origin of the context in which the code
	// establishing the connection is running. See [RFC6454] for the details of how this
	// header field value is constructed.

	// The request MUST include a header field with the name |Sec-WebSocket-Version|.
	// The value of this header field MUST be 13.
	fmt.Fprint(c.nc, "Sec-WebSocket-Version: 13\r\n")

	// The request MAY include a header field with the name |Sec-WebSocket-Protocol|. If
	// present, this value indicates one or more comma-separated subprotocol the client
	// wishes to speak, ordered by preference. The elements that comprise this value MUST
	// be non-empty strings with characters in the range U+0021 to U+007E not including
	// separator characters as defined in [RFC2616] and MUST all be unique strings. The
	// ABNF for the value of this header field is 1#token, where the definitions of
	// constructs and rules are as given in [RFC2616].

	// The request MAY include a header field with the name |Sec-WebSocket-Extensions|.
	// If present, this value indicates the protocol-level extension(s) the client wishes
	// to speak. The interpretation and format of this header field is described in
	// Section 9.1.

	// The request MAY include any other header fields, for example, cookies [RFC6265]
	// and/or authentication-related header fields such as the |Authorization| header
	// field [RFC2616], which are processed according to documents that define them.

	fmt.Fprint(c.nc, "\r\n")
}

func receiveUpgradeResponse(c *Conn, nonce string) error {
	line, err := c.rw.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read the response status line: %v", err)
	}
	// If the status code received from the server is not 101, the client handles the
	// response per HTTP [RFC2616] procedures. In particular, the client might perform
	// authentication if it receives a 401 status code; the server might redirect the
	// client using a 3xx status code (but clients are not required to follow them),
	// etc. Otherwise, proceed as follows.
	if !strings.HasPrefix(line, "HTTP/1.1 101") {
		return fmt.Errorf("expected status code 101, got %s", strings.TrimSpace(line))
	}
	// Parse the headers after the status line.
	gotUpgrade, gotConnection, gotNonce := false, false, false
	for line != "" {
		line, err = c.rw.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read a response header line: %v", err)
		}
		line = strings.TrimSpace(line)
		tokens := strings.Split(line, " ")
		// According to [RFC2616], all header field names in both HTTP requests and
		// HTTP responses are case-insensitive.
		switch strings.ToLower(tokens[0]) {
		// If the response lacks an |Upgrade| header field or the |Upgrade| header field
		// contains a value that is not an ASCII case-insensitive match for the value
		// "websocket", the client MUST _Fail the WebSocket Connection_.
		case "upgrade:":
			gotUpgrade = true
			if strings.ToLower(tokens[1]) != "websocket" {
				return fmt.Errorf("unexpected value in the Upgrade header: %q", tokens[1])
			}
		// If the response lacks a |Connection| header field or the |Connection| header
		// field doesn't contain a token that is an ASCII case-insensitive match for the
		// value "Upgrade", the client MUST _Fail the WebSocket Connection_.
		case "connection:":
			gotConnection = true
			if strings.ToLower(tokens[1]) != "upgrade" {
				return fmt.Errorf("unexpected value in the Connection header: %q", tokens[1])
			}
		// If the response lacks a |Sec-WebSocket-Accept| header field or the
		// |Sec-WebSocket-Accept| contains a value other than the Base64-encoded SHA-1 of
		// the concatenation of the |Sec-WebSocket-Key| (as a string, not base64-decoded)
		// with the string "258EAFA5-E914-47DA-95CA-C5AB0DC85B11" but ignoring any
		// leading and trailing whitespace, the client MUST _Fail the WebSocket
		// Connection_.
		case "sec-websocket-accept:":
			gotNonce = true
			expected, err := expectedKey(nonce)
			if err != nil {
				return err
			}
			if tokens[1] != expected {
				e := "unexpected value in the Sec-WebSocket-Accept header"
				return fmt.Errorf("%s: got %q, expected %q", e, tokens[1], expected)
			}
		case "":
			continue // End of response.
		default:
			log.Printf("unexpected WebSocket upgrade response header: %q", line)
			continue
		}
	}
	// Check we got all the necessary headers.
	if !gotUpgrade {
		return errors.New("WebSocket upgrade response lacks an Upgrade header")
	}
	if !gotConnection {
		return errors.New("WebSocket upgrade response lacks a Connection header")
	}
	if !gotNonce {
		return errors.New("WebSocket upgrade response lacks a Sec-WebSocket-Accept header")
	}
	return nil
}

// Generate the "Sec-WebSocket-Accept" header value expected from the server:
// the Base64-encoded SHA-1 hash of the client-provided Base64 key and the
// GUID "258EAFA5-E914-47DA-95CA-C5AB0DC85B11".
func expectedKey(nonce string) (string, error) {
	h := sha1.New()
	_, err := h.Write([]byte(nonce))
	if err != nil {
		return "", errors.New("failed to hash the server's Sec-WebSocket-Accept header")
	}
	_, err = h.Write([]byte(websocketGUID))
	if err != nil {
		return "", errors.New("failed to hash the server's Sec-WebSocket-Accept header")
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}
