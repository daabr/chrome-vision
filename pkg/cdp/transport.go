package cdp

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

// Error details passed within a CDP response message.
type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// Error satisfies the Go error interface (https://golang.org/pkg/builtin/#error).
func (e *Error) Error() string {
	if e.Code == 0 {
		return e.Message
	}
	return fmt.Sprintf("%s (%d)", e.Message, e.Code)
}

// Message is a generic CDP message sent to or received from a browser.
type Message struct {
	ID     int64           `json:"id,omitempty"`
	Method string          `json:"method,omitempty"`
	Params json.RawMessage `json:"params,omitempty"`
	Result json.RawMessage `json:"result,omitempty"`
	Error  *Error          `json:"error,omitempty"`
}

type asyncMessage struct {
	requestMsg   Message
	responseChan chan<- *Message
}

// Parse and relay incoming CDP messages (solicited responses or unsolicited events),
// received from the browser through a POSIX pipe on non-Windows operating systems,
// as long as the pipe is open. Called as a goroutine in browser.go.
func receiveFromPipe(s *Session) {
	// This scanner wraps the browser's POSIX pipe, which is closed when the browser
	// process ends (see the goroutine at the bottom of the start function in browser.go).
	scanner := bufio.NewScanner(s.browserOutputReader)
	scanner.Split(scanMessages)
	for scanner.Scan() {
		// Receive a new messsage.
		b := scanner.Bytes()
		s.msgLog.Printf("<- %s\n", b)
		// Parse the raw JSON content.
		m := &Message{}
		if err := json.Unmarshal(b, m); err != nil {
			log.Printf("JSON error: %v\n", err)
			continue
		}
		if len(m.Method) == 0 {
			// Solicited response.
			log.Printf("Received response: ID %d (%d bytes)\n", m.ID, len(b))
			if ch, ok := s.responseSubscribers[m.ID]; ok {
				ch <- m
			}
		} else {
			// Unsolicited event.
			log.Printf("Received event: %q (%d bytes)\n", m.Method, len(b))
			if subscribers, ok := s.eventSubscribers[m.Method]; ok {
				for _, ch := range subscribers {
					ch <- m
				}
				switch len(subscribers) {
				case 1:
					log.Printf("Relayed to %d subscribers", len(subscribers))
				default:
					log.Printf("Relayed to 1 subscriber")
				}
			}
		}
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

// Construct and send CDP messages to the browser through a POSIX pipe on non-Windows
// operating systems, in a thread-safe manner (https://blog.golang.org/codelab-share).
// Called in a goroutine in session.go as long as the browser is running.
func sendToPipe(s *Session, async asyncMessage) {
	// Discard malformed data.
	if len(async.requestMsg.Method) == 0 {
		log.Printf("Discarding malformed message: %#v\n", async.requestMsg)
		if async.responseChan != nil {
			m := &Message{ID: s.msgID, Error: &Error{}}
			m.Error.Message = fmt.Sprintf("malformed message: %#v", async.requestMsg)
			async.responseChan <- m
		}
		return
	}
	// Construct the JSON message, and prepare to receive the response.
	async.requestMsg.ID = s.msgID
	b, err := json.Marshal(async.requestMsg)
	if err != nil {
		m := &Message{ID: s.msgID, Error: &Error{Message: err.Error()}}
		async.responseChan <- m
		return
	}
	s.responseSubscribers[s.msgID] = make(chan *Message)
	// Send the JSON message.
	log.Printf("Sending: %s\n", b)
	n, err := s.browserInputWriter.Write(b)
	if err != nil {
		m := &Message{ID: s.msgID, Error: &Error{Message: err.Error()}}
		async.responseChan <- m
		return
	}
	if n < len(b) {
		m := &Message{ID: s.msgID, Error: &Error{}}
		m.Error.Message = fmt.Sprintf("sent %d bytes instead of %d", n, len(b))
		// Don't return like other errors - send \0 and expect an error result.
	}
	// Send \0 to mark the end of the message.
	n, err = s.browserInputWriter.Write([]byte("\000"))
	if err != nil {
		m := &Message{ID: s.msgID, Error: &Error{Message: err.Error()}}
		async.responseChan <- m
		return
	}
	if n != 1 {
		m := &Message{ID: s.msgID, Error: &Error{}}
		m.Error.Message = fmt.Sprintf(`sent %d bytes instead of one \0`, n)
		async.responseChan <- m
		return
	}
	// Wait for the response, clean-up, and relay back to the caller of cdp.Send.
	s.msgLog.Printf("-> %s\n", b)
	m := <-s.responseSubscribers[s.msgID]

	close(s.responseSubscribers[s.msgID])
	delete(s.responseSubscribers, m.ID)

	async.responseChan <- m
}

// Send constructs and sends a JSON message to the browser associated
// with the given context, and returns the browser's response message.
// This is guaranteed to be thread-safe.
func Send(ctx context.Context, method string, params json.RawMessage) (*Message, error) {
	session, ok := FromContext(ctx)
	if !ok {
		return nil, errors.New("context not initialized with cdp.NewContext")
	}
	m := &Message{Method: method, Params: params}
	ch := make(chan *Message)
	// https://blog.golang.org/codelab-share
	session.msgQ <- asyncMessage{requestMsg: *m, responseChan: ch}
	m = <-ch
	close(ch)
	return m, nil
}

// Subscribe returns a channel to receive event messages of the
// given type from the browser associated with the given context.
func SubscribeEvent(ctx context.Context, name string) (chan *Message, error) {
	session, ok := FromContext(ctx)
	if !ok {
		return nil, errors.New("context not initialized with cdp.NewContext")
	}
	ch := make(chan *Message)
	session.eventSubscribers[name] = append(session.eventSubscribers[name], ch)
	return ch, nil
}

// TODO: ubsubscribe.
