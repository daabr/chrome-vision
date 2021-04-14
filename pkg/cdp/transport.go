package cdp

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

// Error details passed within a CDP message.
type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// Error satisfies the Go error interface
// (https://golang.org/pkg/builtin/#error).
func (e *Error) Error() string {
	if e.Code == 0 {
		return e.Message
	}
	return fmt.Sprintf("%s (%d)", e.Message, e.Code)
}

// CDP message sent to or received from the browser.
type Message struct {
	ID     int64           `json:"id,omitempty"`
	Method string          `json:"method,omitempty"`
	Params json.RawMessage `json:"params,omitempty"`
	Result json.RawMessage `json:"result,omitempty"`
	Error  *Error          `json:"error,omitempty"`
}

type asyncMessage struct {
	m Message
	c chan<- Message
}

// Called as a goroutine in browser.go.
func receiveFromPipe(reader *os.File, msgLog *log.Logger) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(scanMessages)
	for scanner.Scan() {
		b := scanner.Bytes()
		msgLog.Printf("<- %s\n", b)
		// TODO: parse and pass to a channel to actually use this.
		// TODO: the following log line should contain only method, ID and byte length.
		log.Printf("Received: %s\n", b)
	}
}

// Helper function based on bufio.ScanLines, using \0 instead of \n as a
// separator - see https://golang.org/pkg/bufio/#example_Scanner_custom.
func scanMessages(data []byte, atEOF bool) (int, []byte, error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '\000'); i >= 0 {
		// We have a full \0-terminated message.
		return i + 1, data[0:i], nil
	}
	// If we're at EOF, we have a final, non-terminated message. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}

// Called in a goroutine in session.go.
func sendToPipe(s *Session, a asyncMessage) {
	// Discard malformed data.
	if len(a.m.Method) == 0 {
		log.Printf("Discarding malformed message: %v\n", a.m)
		if a.c != nil {
			m := Message{ID: s.msgID, Error: &Error{}}
			m.Error.Message = fmt.Sprintf("malformed message: %v", a.m)
			a.c <- m
		}
		return
	}
	// Construct the JSON message.
	a.m.ID = s.msgID
	b, err := json.Marshal(a.m)
	if err != nil {
		m := Message{ID: s.msgID, Error: &Error{Message: err.Error()}}
		a.c <- m
		return
	}
	// Send the JSON message.
	log.Printf("Sending: %s\n", b)
	n, err := s.browserInputWriter.Write(b)
	if err != nil {
		m := Message{ID: s.msgID, Error: &Error{Message: err.Error()}}
		a.c <- m
		return
	}
	if n < len(b) {
		m := Message{ID: s.msgID, Error: &Error{}}
		m.Error.Message = fmt.Sprintf("sent %d bytes instead of %d", n, len(b))
	}
	// Send \0 to mark the end of the message.
	n, err = s.browserInputWriter.Write([]byte("\000"))
	if err != nil {
		m := Message{ID: s.msgID, Error: &Error{Message: err.Error()}}
		a.c <- m
		return
	}
	if n != 1 {
		m := Message{ID: s.msgID, Error: &Error{}}
		m.Error.Message = fmt.Sprintf("sent %d bytes instead of 1", n)
	}
	// Success
	s.msgLog.Printf("-> %s\n", b)
	a.c <- Message{ID: s.msgID}
}

// Send constructs and sends a JSON message to the browser associated with the
// given context, and returns the auto-assigned ID used for that message.
//
// TODO: block and return result message?
func Send(ctx context.Context, method string, params json.RawMessage) (int64, error) {
	session, ok := FromContext(ctx)
	if !ok {
		return 0, errors.New("context not initialized with cdp.NewContext")
	}
	m := Message{Method: method}
	if len(params) > 0 {
		m.Params = params
	}
	c := make(chan Message)
	session.msgQ <- asyncMessage{m: m, c: c}
	m = <-c
	return m.ID, nil
}
