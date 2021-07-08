// Package websocket is a lightweight client implementation of the WebSocket
// protocol (RFC 6455). It is designed specifically for fast, efficient and
// idiomatic client-side communication with Chrome DevTools in Blink-based
// browsers.
//
// Examples for unsupported features: server-side functionality, proxies, TLS
// for "wss://" addresses, and handshake customizations such as sub-protocols,
// extensions, and other HTTP headers (e.g. authentication, cookies).
//
// Specifically, this implementation does not support the "permessage-deflate"
// extension (RFC 7692), for the same reasons it doesn't support TLS: almost
// all communication with Chrome DevTools should happen on the same localhost,
// and most transactions involve very small amounts of data, so the security
// benefits of TLS and performance benefits of compression become irrelevant
// in this use-case, and are in fact wasteful and unnecessarily slow.
//
// More readaing materials about other WebSocket implementations:
//
// • https://centrifugal.github.io/centrifugo/blog/scaling_websocket/
//
// • https://www.freecodecamp.org/news/million-websockets-and-go-cc58418460bb/
package websocket

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
)

// Based on https://datatracker.ietf.org/doc/html/rfc6455#section-11.8.
type opcode byte

const (
	continuationFrame opcode = iota
	textFrame
	binaryFrame
	// Opcodes 3-7 are reserved for further non-control frames.
	_
	_
	_
	_
	_
	connectionCloseFrame
	pingFrame
	pongFrame
	// Opcodes 11-15 are reserved for further control frames.
)

// Based on https://datatracker.ietf.org/doc/html/rfc6455#section-5.2.
type frame struct {
	// Bit 0: Indicates that this is the final fragment in a message.
	fin bool
	// Bits 1-3: Reserved. MUST be 0 unless an extension is negotiated that
	// defines meanings for non-zero values.
	rsv [3]bool
	// Bits 4-7: Defines the interpretation of the "Payload data".
	opcode opcode
	// Bit 8: Defines whether the "Payload data" is masked. If set to 1, a masking key
	// is present in masking-key, and this is used to unmask the "Payload data" as per
	// Section 5.3. All frames sent from client to server have this bit set to 1.
	mask bool
	// Bits 9-15 +0 or +2 or +8 bytes: The length of the "Payload data", in bytes: if
	// 0-125, that is the payload length. If 126, the following 2 bytes interpreted as
	// a 16-bit unsigned integer are the payload length. If 127, the following 8 bytes
	// interpreted as a 64-bit unsigned integer (the most significant bit MUST be 0) are
	// the payload length. Multibyte length quantities are expressed in network byte
	// order. Note that in all cases, the minimal number of bytes MUST be used to encode
	// the length, for example, the length of a 124-byte-long string can't be encoded as
	// the sequence 126, 0, 124. The payload length is the length of the "Extension data"
	// + the length of the "Application data". The length of the "Extension data" may be
	// zero, in which case the payload length is the length of the "Application data".
	payloadLength uint64
	// 0 or 4 bytes: All frames sent from the client to the server are masked by a 32-bit
	// value that is contained within the frame. This field is present if the mask bit is
	// set to 1 and is absent if the mask bit is set to 0. See Section 5.3 for further
	// information on client-to-server masking.
	maskingKey []byte
	// "Extension data" concatenated with "Application data". The "Extension data" is 0
	// bytes unless an extension has been negotiated. Any extension MUST specify the
	// length of the "Extension data", or how that length may be calculated, and how the
	// extension use MUST be negotiated during the opening handshake. If present, the
	// "Extension data" is included in the total payload length.
	payloadData []byte
}

// Based on https://datatracker.ietf.org/doc/html/rfc6455#section-5.2.
func (c *Conn) readFrame() (f frame, closeConnection bool, err error) {
	// First header byte.
	b, err := c.rw.ReadByte()
	if err != nil {
		if err == io.EOF {
			return f, false, err // TODO: any way to block until available?
		}
		return f, false, fmt.Errorf("failed to read the first header byte: %v", err)
	}
	f.fin = (b & 0x80) != 0    // Bit 0.
	f.rsv[0] = (b & 0x40) != 0 // Bit 1.
	f.rsv[1] = (b & 0x20) != 0 // Bit 2.
	f.rsv[2] = (b & 0x10) != 0 // Bit 3.
	if (b & 0x70) != 0 {
		// Reserved bits MUST be 0 unless an extension is negotiated that defines
		// meanings for non-zero values. If a nonzero value is received and none of
		// the negotiated extensions defines the meaning of such a nonzero value,
		// the receiving endpoint MUST _Fail the WebSocket Connection_.
		return f, true, errors.New("server unexpectedly sent non-0 reserved bits")
	}
	f.opcode = opcode(b & 0x0f) // Bits 4-7.
	if (f.opcode > 2 && f.opcode < 8) || f.opcode > 10 {
		// If an unknown opcode is received, the receiving endpoint MUST
		// _Fail the WebSocket Connection_.
		return f, true, fmt.Errorf("server sent an unknown opcode %d", f.opcode)
	}

	// Second header byte.
	b, err = c.rw.ReadByte()
	if err != nil {
		return f, false, fmt.Errorf("failed to read the second header byte: %v", err)
	}
	f.mask = (b & 0x80) != 0 // Bit 0.
	if f.mask {
		// https://datatracker.ietf.org/doc/html/rfc6455#section-5.1: A server
		// MUST NOT mask any frames that it sends to the client. A client MUST
		// close a connection if it detects a masked frame.
		return f, true, errors.New("server unexpectedly masked the payload data")
	}
	b &= 0x7f // Bits 1-7.

	// Extended payload length (0 or 2 or 8 bytes).
	switch {
	case b <= 125: // Up to 125 bytes.
		f.payloadLength = uint64(b)
	case b == 126: // Up to 64 KiB.
		extendedLength := make([]byte, 2)
		_, err = io.ReadFull(c.rw, extendedLength)
		f.payloadLength = uint64(binary.BigEndian.Uint16(extendedLength))
	default: // Up to 16 EiB.
		extendedLength := make([]byte, 8)
		_, err = io.ReadFull(c.rw, extendedLength)
		f.payloadLength = binary.BigEndian.Uint64(extendedLength)
	}
	if err != nil {
		return f, false, fmt.Errorf("failed to read extended payload length: %v", err)
	}

	// Unmasked payload data (variable length).
	f.payloadData = make([]byte, f.payloadLength)
	_, err = io.ReadFull(c.rw, f.payloadData)
	if err != nil {
		return f, false, fmt.Errorf("failed to read the payload: %v", err)
	}
	return f, false, nil
}

// Based on https://datatracker.ietf.org/doc/html/rfc6455#section-6.2,
// https://datatracker.ietf.org/doc/html/rfc6455#section-5.4,
// https://datatracker.ietf.org/doc/html/rfc6455#section-5.5,
// and https://datatracker.ietf.org/doc/html/rfc6455#section-7.
func (c *Conn) readMessage() ([]byte, error) {
	msg := bytes.NewBuffer([]byte{})
	for {
		f, close, err := c.readFrame()
		if close {
			log.Printf("Closing WebSocket connection due to protocol error: %v", err)
			c.Close(1002, []byte{})
		}
		if err != nil {
			if err == io.EOF {
				continue // TODO: remove this if/when readFrame() becomes blocking.
			}
			return nil, err
		}
		// Control frames (see Section 5.5) MAY be injected in the middle of a
		// fragmented message. Control frames themselves MUST NOT be fragmented.
		// (https://datatracker.ietf.org/doc/html/rfc6455#section-5.5)
		if f.opcode == connectionCloseFrame {
			statusCode := uint16(1005) // Status not specified (RFC 6455 section 7.4.1).
			if f.payloadLength >= 2 {
				statusCode = binary.BigEndian.Uint16(f.payloadData[0:2])
			}
			c.Close(statusCode, f.payloadData[2:])
			s := "WebSocket server closing connection: status code"
			return nil, fmt.Errorf("%s %d, reason %q", s, statusCode, f.payloadData[2:])
		}
		if f.opcode == pingFrame {
			log.Printf("Received ping control frame from WebSocket server: %q", f.payloadData)
			c.WritePong(f.payloadData)
			continue
		}
		if f.opcode == pongFrame {
			log.Printf("Received pong control frame from WebSocket server: %q", f.payloadData)
			continue
		}
		// Handle data frames. The fragments of one message MUST NOT be interleaved
		// between the fragments of another message unless an extension has been
		// negotiated that can interpret the interleaving.
		if f.fin {
			// An unfragmented message consists of a single frame with the FIN
			// bit set (Section 5.2) and an opcode other than 0.
			if f.opcode != continuationFrame {
				return f.payloadData, nil
			}
			// A fragmented message [...] terminated by a single frame with
			// the FIN bit set and an opcode of 0.
			msg.Write(f.payloadData)
			return msg.Bytes(), nil
		}
		if f.opcode != continuationFrame {
			// A fragmented message consists of a single frame with the
			// FIN bit clear and an opcode other than 0...
			msg = bytes.NewBuffer(f.payloadData)
		} else {
			// ...Followed by zero or more frames with the FIN bit clear
			// and the opcode set to 0.
			msg.Write(f.payloadData)
		}
	}
}

// Based on https://datatracker.ietf.org/doc/html/rfc6455#section-5.2
// and https://datatracker.ietf.org/doc/html/rfc6455#section-6.1.
func (c *Conn) writeFrame(f frame) error {
	// First header byte.
	var b byte
	if f.fin {
		b |= 0x80 // Bit 0.
	}
	for i := 0; i < 3; i++ { // Bits 1-3.
		if f.rsv[i] {
			b |= 1 << (6 - i)
		}
	}
	b |= byte(f.opcode) // Bits 4-7.
	if err := c.rw.WriteByte(b); err != nil {
		return fmt.Errorf("failed to write the first header byte: %v", err)
	}

	// Second header byte.
	b = 0x80 // Bit 0 (client-sent frames MUST be masked).
	extendedLength := 0
	switch {
	case f.payloadLength <= 125: // Up to 125 bytes.
		b |= byte(f.payloadLength)
	case f.payloadLength <= 65535: // Up to 64 KiB.
		b |= 126
		extendedLength = 2
	default: // Up to 16 EiB.
		b |= 127
		extendedLength = 8
	}
	if err := c.rw.WriteByte(b); err != nil {
		return fmt.Errorf("failed to write the second header byte: %v", err)
	}

	// Extended payload length (0 or 2 or 8 bytes).
	for i := 0; i < extendedLength; i++ {
		j := uint((extendedLength - i - 1) * 8)
		b = byte((f.payloadLength >> j) & 0xff)
		if err := c.rw.WriteByte(b); err != nil {
			return fmt.Errorf("failed to write the extended payload length: %v", err)
		}
	}

	// Masking key (0 or 4 bytes).
	_, err := c.rw.Write(f.maskingKey)
	if err != nil {
		return fmt.Errorf("failed to write the masking key: %v", err)
	}

	// Masked payload data (variable length).
	_, err = c.rw.Write(f.payloadData)
	if err != nil {
		return fmt.Errorf("failed to write the masked payload: %v", err)
	}

	// Flush.
	if err := c.rw.Flush(); err != nil {
		return fmt.Errorf("failed to flush frame after writing: %v", err)
	}
	return nil
}

func (c *Conn) writeMessage(o opcode, msg []byte) error {
	// https://datatracker.ietf.org/doc/html/rfc6455#section-5.4: An unfragmented
	// message consists of a single frame with the FIN bit set (Section 5.2) and
	// an opcode other than 0.
	// https://datatracker.ietf.org/doc/html/rfc6455#section-5.2: Reserved bits
	// MUST be 0 unless an extension is negotiated that defines meanings for
	// non-zero values.
	// https://datatracker.ietf.org/doc/html/rfc6455#section-5.1: To avoid confusing
	// network intermediaries (such as intercepting proxies) and for security reasons
	// that are further discussed in Section 10.3, a client MUST mask all frames that
	// it sends to the server (see Section 5.3 for further details).
	f := frame{fin: true, opcode: o, mask: true, payloadLength: uint64(len(msg))}

	// https://datatracker.ietf.org/doc/html/rfc6455#section-5.3: masking.
	f.maskingKey = make([]byte, 4)
	if _, err := io.ReadFull(rand.Reader, f.maskingKey); err != nil {
		return fmt.Errorf("failed to generate frame masking key: %v", err)
	}
	f.payloadData = make([]byte, f.payloadLength)
	for i := range msg {
		f.payloadData[i] = msg[i] ^ f.maskingKey[i%4]
	}

	return c.writeFrame(f)
}

// Read receives a full message from a WebSocket server. It handles all
// the implementation details internally, such as frame de/fragmentation,
// masking, and handling control frames.
func (c *Conn) Read() ([]byte, error) {
	b, err := c.readMessage()
	if err != nil {
		err = fmt.Errorf("failed to read message from WebSocket: %v", err)
	}
	return b, err
}

// WriteText sends a full UTF-8 text message to a WebSocket server. It handles
// all the implementation details internally, such as frame de/fragmentation,
// masking, and handling control frames.
func (c *Conn) WriteText(msg []byte) error {
	err := c.writeMessage(textFrame, msg)
	if err != nil {
		err = fmt.Errorf("failed to write text message to WebSocket: %v", err)
	}
	return err
}

// WriteBinary sends a full binary text message to a WebSocket server.
// It handles all the implementation details internally, such as frame
// de/fragmentation, masking, and handling control frames.
func (c *Conn) WriteBinary(msg []byte) error {
	err := c.writeMessage(binaryFrame, msg)
	if err != nil {
		err = fmt.Errorf("failed to write binary message to WebSocket: %v", err)
	}
	return err
}

// WritePing sends a "ping" control frame to a WebSocket server, as a
// keep-alive or as a means to verify that the server is still responsive.
// Based on https://datatracker.ietf.org/doc/html/rfc6455#section-5.5.2.
func (c *Conn) WritePing(appData []byte) error {
	if len(appData) > 125 {
		return errors.New("control frames must have a payload of 0-125 bytes")
	}
	err := c.writeMessage(pingFrame, appData)
	if err != nil {
		err = fmt.Errorf("failed to write ping control frame to WebSocket: %v", err)
	}
	return err
}

// WritePong sends a "pong" control frame to a WebSocket server, as a
// response to a "ping" (with the ping's "application data" payload),
// or as an unsolicited unidirectional heartbeat. Based on
// https://datatracker.ietf.org/doc/html/rfc6455#section-5.5.3.
func (c *Conn) WritePong(appData []byte) error {
	if len(appData) > 125 {
		return errors.New("control frames must have a payload of 0-125 bytes")
	}
	err := c.writeMessage(pongFrame, appData)
	if err != nil {
		err = fmt.Errorf("failed to write pong control frame to WebSocket: %v", err)
	}
	return err
}

// Close sends a "connection close" control frame to a WebSocket server, and
// then closes the underlying TCP network connection. Valid status codes are
// described in https://datatracker.ietf.org/doc/html/rfc6455#section-7.4.1,
// and the reason is a UTF-8-encoded string which may be empty. Based on
// https://datatracker.ietf.org/doc/html/rfc6455#section-5.5.1 and
// https://datatracker.ietf.org/doc/html/rfc6455#section-7.
func (c *Conn) Close(statusCode uint16, reason []byte) error {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, statusCode)
	c.writeMessage(connectionCloseFrame, append(b, reason...))
	// The client SHOULD wait for the server to close the connection but MAY close
	// the connection at any time after sending and receiving a Close message.
	return c.nc.Close()
}
