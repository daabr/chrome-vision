package devtools

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
	ID        int64           `json:"id,omitempty"`
	SessionID string          `json:"sessionId,omitempty"`
	Method    string          `json:"method,omitempty"`
	Params    json.RawMessage `json:"params,omitempty"`
	Result    json.RawMessage `json:"result,omitempty"`
	Error     *Error          `json:"error,omitempty"`
}

type asyncMessage struct {
	requestMsg   Message
	responseChan chan<- *Message
}

// Parse and relay incoming CDP messages.
func parseAndRelay(s *Session, b []byte) {
	s.msgLog.Printf("<- %s\n", b)

	// Parse the raw JSON content.
	m := &Message{}
	if err := json.Unmarshal(b, m); err != nil {
		log.Printf("JSON error: %v", err)
		return
	}

	if len(m.Method) == 0 {
		// Solicited response: relay to the request caller.
		log.Printf("Received response: ID %d (%d bytes)", m.ID, len(b))
		if ch, ok := s.responseSubscribers[m.ID]; ok {
			ch <- m
		}
	} else {
		// Unsolicited event: relay to any subscribers.
		log.Printf("Received event: %q (%d bytes)", m.Method, len(b))
		if subscribers, ok := s.eventSubscribers[m.Method]; ok {
			for _, ch := range subscribers {
				ch <- m
			}
			switch len(subscribers) {
			case 1:
				log.Printf("Relayed to 1 subscriber")
			default:
				log.Printf("Relayed to %d subscribers", len(subscribers))
			}
		}
	}
}

// Asynchronously receive incoming CDP messages from the browser through a
// POSIX pipe on non-Windows operating systems, as long as the pipe is open.
// Called as a goroutine in the `start` function in `browser.go`.
func receiveFromPipe(s *Session) {
	// This scanner wraps the browser's POSIX pipe, which is closed when the
	// browser process ends (see the goroutine at the bottom of the `start`
	// function in `browser.go`).
	scanner := bufio.NewScanner(s.browserOutputReader)
	scanner.Split(scanMessages)
	for scanner.Scan() {
		b := scanner.Bytes()
		parseAndRelay(s, b)
	}
}

// Helper function based on `bufio.ScanLines`, using \0 instead of \n as a
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

// Asynchronously receive incoming CDP messages from the browser through a
// WebSocket on Windows operating systems, as long as the connection is open.
// Called as a goroutine in the `start` function in `browser.go`.
func receiveFromWebSocket(s *Session) {
	for {
		b, err := s.webSocket.Read()
		if err != nil {
			if err == io.EOF {
				continue
			}
			log.Printf("WARNING: failed to read incoming CDP message: %v", err)
			return
		}
		parseAndRelay(s, b)
	}
}

func preSend(s *Session, async *asyncMessage) ([]byte, error) {
	// Discard malformed data.
	if len(async.requestMsg.Method) == 0 {
		log.Printf("Discarding malformed message: %#v", async.requestMsg)
		if async.responseChan != nil {
			m := &Message{ID: s.msgID, Error: &Error{}}
			m.Error.Message = fmt.Sprintf("malformed message: %#v", async.requestMsg)
			async.responseChan <- m
		}
		return nil, errors.New("malformed message")
	}
	// Construct the JSON message, and prepare to receive the response.
	async.requestMsg.ID = s.msgID
	b, err := json.Marshal(async.requestMsg)
	if err != nil {
		m := &Message{ID: s.msgID, Error: &Error{Message: err.Error()}}
		async.responseChan <- m
		return nil, errors.New(m.Error.Message)
	}

	s.responseSubscribers[s.msgID] = make(chan *Message)
	log.Printf("Sending: %s", b)
	return b, nil
}

func postSend(s *Session, async asyncMessage, b []byte) {
	// Wait for the response, clean-up, and relay back to the caller of devtools.Send.
	s.msgLog.Printf("-> %s\n", b)
	m := <-s.responseSubscribers[s.msgID]

	close(s.responseSubscribers[s.msgID])
	delete(s.responseSubscribers, m.ID)

	async.responseChan <- m
}

// Construct and send CDP messages to the browser through a POSIX pipe on non-Windows
// operating systems, in a thread-safe manner (https://blog.golang.org/codelab-share).
// Called in a goroutine in `session.go` as long as the browser is running.
func sendToPipe(s *Session, async asyncMessage) {
	b, err := preSend(s, &async)
	if err != nil {
		return // Already reported to the caller by marshalJSON().
	}

	// Send the JSON message.
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

	postSend(s, async, b)
}

// Construct and send CDP messages to the browser through a WebSocket on Windows
// operating systems, in a thread-safe manner (https://blog.golang.org/codelab-share).
// Called in a goroutine in `session.go` as long as the browser is running.
func sendToWebSocket(s *Session, async asyncMessage) {
	b, err := preSend(s, &async)
	if err != nil {
		return // Already reported to the caller by preSend.
	}

	// Send the JSON message.
	err = s.webSocket.WriteText(b)
	if err != nil {
		m := &Message{ID: s.msgID, Error: &Error{Message: err.Error()}}
		async.responseChan <- m
		return
	}

	postSend(s, async, b)
}

// Send constructs and sends a CDP message to the browser associated
// with the given context, and returns the browser's response message.
// Multiple goroutines may call this function simultaneously.
func Send(ctx context.Context, method string, params json.RawMessage) (*Message, error) {
	s, ok := FromContext(ctx)
	if !ok {
		return nil, errors.New("context not initialized with devtools.NewContext")
	}
	// https://github.com/aslushnikov/getting-started-with-cdp#targets--sessions
	m := &Message{Method: method, SessionID: s.sessionID.read(), Params: params}
	ch := make(chan *Message)
	// https://blog.golang.org/codelab-share
	s.msgQ <- asyncMessage{requestMsg: *m, responseChan: ch}
	m = <-ch
	close(ch)
	return m, nil
}

// SubscribeEvent returns a channel to receive event messages of
// the given type from the browser associated with the given context.
func SubscribeEvent(ctx context.Context, name string) (chan *Message, error) {
	s, ok := FromContext(ctx)
	if !ok {
		return nil, errors.New("context not initialized with devtools.NewContext")
	}
	ch := make(chan *Message)
	s.eventSubscribers[name] = append(s.eventSubscribers[name], ch)
	return ch, nil
}

// TODO: unsubscribe.
