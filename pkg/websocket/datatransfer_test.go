package websocket_test

import (
	"encoding/binary"
	"io"
	"net"
	"testing"

	"github.com/daabr/chrome-vision/pkg/websocket"
	"github.com/google/go-cmp/cmp"
)

func TestReadErrors(t *testing.T) {
	tests := []struct {
		desc string
		b    []byte
	}{
		{
			"reserved bits",
			[]byte{0x70},
		},
		{
			"invalid opcode",
			[]byte{0x0f},
		},
		{
			"mask bit",
			[]byte{0x80, 0x80},
		},
	}
	for i, tc := range tests {
		server, client := net.Pipe()
		conn := websocket.NewConn(client)
		defer server.Close()
		defer client.Close()

		go func() {
			server.Write(tc.b)
			server.Read(make([]byte, 8))
		}()

		got, err := conn.Read()
		if err == nil {
			t.Errorf("TC %d: Conn.Read() = %#v, want %s error", i, got, tc.desc)
		}
	}
}

func TestReadSingleEmptyFrame(t *testing.T) {
	server, client := net.Pipe()
	conn := websocket.NewConn(client)
	defer server.Close()
	defer client.Close()

	go func() {
		b := []byte{0x81, 0x00}
		server.Write(b)
	}()

	got, err := conn.Read()
	if err != nil {
		t.Fatalf("Conn.Read(); got unexpected error: %v", err)
	}
	if len(got) > 0 {
		t.Errorf("Conn.Read() = %#v, want %#v", got, []byte{})
	}
}

func TestReadThreeFrames(t *testing.T) {
	server, client := net.Pipe()
	conn := websocket.NewConn(client)
	defer server.Close()
	defer client.Close()

	go func() {
		b := []byte{0x01, 0x01, 0xaa, 0x00, 0x02, 0xbb, 0xcc, 0x80, 0x03, 0xdd, 0xee, 0xff}
		server.Write(b)
	}()

	got, err := conn.Read()
	want := []byte{0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff}
	if err != nil {
		t.Fatalf("Conn.Read(); got unexpected error: %v", err)
	}
	if !cmp.Equal(got, want) {
		t.Errorf("Conn.Read() = %#v, want %#v", got, want)
	}
}

func TestReadWithControlFrames(t *testing.T) {
	server, client := net.Pipe()
	conn := websocket.NewConn(client)
	defer server.Close()
	defer client.Close()

	go func() {
		b := []byte{0x01, 0x01, 0xaa, 0x89, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x8a, 0x00, 0x80, 0x03, 0xdd, 0xee, 0xff}
		server.Write(b)
		server.Read(make([]byte, 10))
	}()

	got, err := conn.Read()
	want := []byte{0xaa, 0xdd, 0xee, 0xff}
	if err != nil {
		t.Fatalf("Conn.Read(); got unexpected error: %v", err)
	}
	if !cmp.Equal(got, want) {
		t.Errorf("Conn.Read() = %#v, want %#v", got, want)
	}
}

func TestCloseDuringRead(t *testing.T) {
	server, client := net.Pipe()
	conn := websocket.NewConn(client)
	defer server.Close()
	defer client.Close()

	go func() {
		b := []byte{0x01, 0x01, 0xaa, 0x88, 0x08, 0x03, 0xe9, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e}
		server.Write(b)
		server.Read(make([]byte, 14))
	}()

	got, err := conn.Read()
	if err == nil {
		t.Errorf("Conn.Read() = %#v, want connection closing error", got)
	}
}

func TestRead1KB(t *testing.T) {
	server, client := net.Pipe()
	conn := websocket.NewConn(client)
	defer server.Close()
	defer client.Close()

	want := make([]byte, 1024)
	want[0] = 0x01
	want[1023] = 0x02

	go func() {
		extendedLength := make([]byte, 2)
		binary.BigEndian.PutUint16(extendedLength, uint16(len(want)))
		b := append([]byte{0x81, 0x7e}, extendedLength...)
		b = append(b, want...)
		server.Write(b)
	}()

	got, err := conn.Read()
	if err != nil {
		t.Fatalf("Conn.Read(); got unexpected error: %v", err)
	}
	if !cmp.Equal(got, want) {
		t.Errorf("Conn.Read() = %#v, want %#v", got, want)
	}
}

func TestRead1MB(t *testing.T) {
	server, client := net.Pipe()
	conn := websocket.NewConn(client)
	defer server.Close()
	defer client.Close()

	want := make([]byte, 1024*1024)
	want[0] = 0x01
	want[1024*1024-1] = 0x02

	go func() {
		extendedLength := make([]byte, 8)
		binary.BigEndian.PutUint64(extendedLength, uint64(len(want)))
		b := append([]byte{0x81, 0x7f}, extendedLength...)
		b = append(b, want...)
		server.Write(b)
	}()

	got, err := conn.Read()
	if err != nil {
		t.Fatalf("Conn.Read(); got unexpected error: %v", err)
	}
	if !cmp.Equal(got, want) {
		t.Errorf("Conn.Read() = %#v, want %#v", got, want)
	}
}

func TestWriteSingleEmptyFrame(t *testing.T) {
	server, client := net.Pipe()
	conn := websocket.NewConn(client)
	defer server.Close()
	defer client.Close()

	go func() {
		conn.WriteBinary([]byte{})
	}()

	b := make([]byte, 7)
	n, err := server.Read(b)
	if err != nil {
		t.Fatalf("server.Read(b); got unexpected error: %v", err)
	}
	if n != 6 {
		t.Errorf("server.Read(b) = %d, want 6", n)
	}
	if b[0] != 0x82 {
		t.Errorf("server.Read(b); b[0] = %b, want %b", b, 0x82)
	}
	if b[1] != 0x80 {
		t.Errorf("server.Read(b); b[1] = %b, want %b", b, 0x80)
	}
}

func TestWrite1Byte(t *testing.T) {
	server, client := net.Pipe()
	conn := websocket.NewConn(client)
	defer server.Close()
	defer client.Close()

	go func() {
		conn.WriteText([]byte{0x00})
	}()

	// 2 (minimum header) + 4 (masking key) + 1 (payload)
	// + 1 (to detect unexpected bytes).
	want := 2 + 4 + 1
	b := make([]byte, want+1)
	n, err := server.Read(b)
	if err != nil {
		t.Fatalf("server.Read(b); got unexpected error: %v", err)
	}
	if n != want {
		t.Errorf("server.Read(b) = %d, want %d", n, want)
	}
	if b[0] != 0x81 { // Fin, text frame.
		t.Errorf("server.Read(b); b[0] = %b, want %b", b[0], 0x81)
	}
	if b[1] != 0x81 { // Mask, payload length 1.
		t.Errorf("server.Read(b); b[1] = %b, want %b", b[1], 0x81)
	}
}

func TestWrite1KB(t *testing.T) {
	server, client := net.Pipe()
	conn := websocket.NewConn(client)
	defer server.Close()
	defer client.Close()

	go func() {
		conn.WriteBinary(make([]byte, 1024))
	}()

	// 2 (minimum header) + 2 (extended length) + 4 (masking key)
	// + 1024 (payload) + 1 (to detect unexpected bytes).
	want := 2 + 2 + 4 + 1024
	b := make([]byte, want+1)
	n, err := server.Read(b)
	if err != nil {
		t.Fatalf("server.Read(b); got unexpected error: %v", err)
	}
	if n != want {
		t.Errorf("server.Read(b) = %d, want %d", n, want)
	}
	if b[0] != 0x82 { // Fin, binary frame.
		t.Errorf("server.Read(b); b[0] = %b, want %b", b[0], 0x82)
	}
	if b[1] != 0xfe { // Mask, payload length 126.
		t.Errorf("server.Read(b); b[1] = %b, want %b", b[1], 0xfe)
	}
}

func TestWrite1MB(t *testing.T) {
	server, client := net.Pipe()
	conn := websocket.NewConn(client)
	defer server.Close()
	defer client.Close()

	go func() {
		conn.WriteText(make([]byte, 1024*1024))
	}()

	// 2 (minimum header) + 8 (extended length) + 4 (masking key)
	// + 1024^2 (payload) + 1 (to detect unexpected bytes).
	want := 2 + 8 + 4 + (1024 * 1024)
	b := make([]byte, want+1)
	n, err := io.ReadAtLeast(server, b, 1000*1000)
	if err != nil {
		t.Fatalf("server.Read(b); got unexpected error: %v", err)
	}
	if n != want {
		t.Errorf("server.Read(b) = %d, want %d", n, want)
	}
	if b[0] != 0x81 { // Fin, text frame.
		t.Errorf("server.Read(b); b[0] = %b, want %b", b[0], 0x81)
	}
	if b[1] != 0xff { // Mask, payload length 127.
		t.Errorf("server.Read(b); b[1] = %b, want %b", b[1], 0xff)
	}
}

func TestWritePing(t *testing.T) {
	server, client := net.Pipe()
	conn := websocket.NewConn(client)
	defer server.Close()
	defer client.Close()

	go func() {
		conn.WritePing([]byte("ping"))
	}()

	// 2 (minimum header) + 4 (masking key) + 4 (payload)
	// + 1 (to detect unexpected bytes).
	want := 2 + 4 + 4
	b := make([]byte, want+1)
	n, err := server.Read(b)
	if err != nil {
		t.Fatalf("server.Read(b); got unexpected error: %v", err)
	}
	if n != want {
		t.Errorf("server.Read(b) = %d, want %d", n, want)
	}
	if b[0] != 0x89 { // Fin, ping control frame.
		t.Errorf("server.Read(b); b[0] = %b, want %b", b[0], 0x89)
	}
	if b[1] != 0x84 { // Mask, payload length 4.
		t.Errorf("server.Read(b); b[1] = %b, want %b", b[1], 0x84)
	}
}

func TestWritePong(t *testing.T) {
	server, client := net.Pipe()
	conn := websocket.NewConn(client)
	defer server.Close()
	defer client.Close()

	go func() {
		conn.WritePong([]byte{})
	}()

	// 2 (minimum header) + 4 (masking key) + 0 (payload)
	// + 1 (to detect unexpected bytes).
	want := 2 + 4 + 0
	b := make([]byte, want+1)
	n, err := server.Read(b)
	if err != nil {
		t.Fatalf("server.Read(b); got unexpected error: %v", err)
	}
	if n != want {
		t.Errorf("server.Read(b) = %d, want %d", n, want)
	}
	if b[0] != 0x8a { // Fin, pong control frame.
		t.Errorf("server.Read(b); b[0] = %b, want %b", b[0], 0x8a)
	}
	if b[1] != 0x80 { // Mask, payload length 0.
		t.Errorf("server.Read(b); b[1] = %b, want %b", b[1], 0x80)
	}
}

func TestClose(t *testing.T) {
	server, client := net.Pipe()
	conn := websocket.NewConn(client)
	defer server.Close()
	defer client.Close()

	go func() {
		conn.Close(1001, []byte("reason"))
	}()

	// 2 (minimum header) + 4 (masking key) + 8 (payload)
	// + 1 (to detect unexpected bytes).
	want := 2 + 4 + 8
	b := make([]byte, want+1)
	n, err := server.Read(b)
	if err != nil {
		t.Fatalf("server.Read(b); got unexpected error: %v", err)
	}
	if n != want {
		t.Errorf("server.Read(b) = %d, want %d", n, want)
	}
	if b[0] != 0x88 { // Fin, pong control frame.
		t.Errorf("server.Read(b); b[0] = %b, want %b", b[0], 0x88)
	}
	if b[1] != 0x88 { // Mask, payload length 8.
		t.Errorf("server.Read(b); b[1] = %b, want %b", b[1], 0x88)
	}
}
