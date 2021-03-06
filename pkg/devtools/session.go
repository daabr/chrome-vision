// The types and functions in this file follow the example in comment of the
// context.Context.Value field in https://golang.org/pkg/context/#Context.
// See also https://golang.org/pkg/context/#example_WithValue.

package devtools

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"github.com/daabr/chrome-vision/pkg/websocket"
)

// OutputRootEnv is the name of an optional environment variable to override
// the default path for CDP output directories (Go's `os.TempDir()` -
// see https://golang.org/pkg/os/#TempDir)
const OutputRootEnv = "CDP_OUTPUT_ROOT"

// TargetTimeout is the maximum amount of time to wait for the browser to start
// using a new tab, when constructing a new session by calling `NewContext`.
const TargetTimeout = 1 * time.Minute

// SafeString is a "thread-safe" string: multiple goroutines may
// read and write it simultaneously.
type SafeString struct {
	mu sync.RWMutex
	string
}

func newSafeString() *SafeString {
	return &SafeString{mu: sync.RWMutex{}}
}

func (t *SafeString) Read() string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.string
}

func (t *SafeString) Write(s string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.string = s
}

// Session contains CDP runtime details, stored in a `context.Context`,
// and retrievable with the `devtools.FromContext` function.
type Session struct {
	// For immediate canceling of the context returned by `devtools.NewContext`,
	// and any descendant contexts, when the browser process ends.
	cancel context.CancelFunc

	// Data directory per browser process. Created under Go's `os.TempDir()`,
	// or the path specified in the environment variable "CDP_OUTPUT_ROOT",
	// if set (see `devtools.OutputRootEnv`). Contains STDOUT and STDERR dumps,
	// logs, user profile data, screenshots, etc. Not deleted automatically!
	OutputDir string
	// User data directory per browser process, instead of the system's
	// default location. Contains data such as history, bookmarks and cookies
	// (https://chromium.googlesource.com/chromium/src/+/master/docs/user_data_dir.md).
	// Created by default under this session's `OutputDir`, unless the caller
	// of the `devtools.NewContext` function overrides this behavior by
	// specifying the `devtools.UserDataDir` session option.
	UserDataDir string

	// Browser execution details. Not shared with descendant contexts because
	// the browser was already started by the first call to `devtools.NewContext`.
	browserPath  *string
	browserFlags map[string]interface{}
	// TODO: environment variables.

	browserDone chan struct{}

	// Communication with the browser...
	// ...On POSIX-compliant operating systems - via pipes.
	browserInputWriter, browserOutputReader *os.File
	// ...On Windows - via WebSocket.
	wsAddress, wsPath *SafeString
	webSocket         *websocket.Conn

	msgLog *log.Logger
	msgID  int64
	msgQ   chan asyncMessage // https://blog.golang.org/codelab-share

	// Exactly one subscriber per response (created and used in devtools.Send).
	responseSubscribers map[int64]chan *Message
	// Zero or more subscribers per event type.
	eventSubscribers map[string][]chan *Message

	// IDs for the attached browser tab. Not shared with descendant contexts
	// because they create their own tabs, targets and sessions IDs. See also:
	// https://github.com/aslushnikov/getting-started-with-cdp#targets--sessions.
	TargetID, SessionID *SafeString
}

type sessionKey struct{}

// Copy of `target.CreateTargetResult`, for parsing the created target ID.
type createTargetResult struct {
	TargetID string `json:"targetId"`
}

// Partial copy of `target.TargetInfo`, for parsing target states.
type targetInfo struct {
	TargetID string `json:"targetId"`
	Type     string `json:"type"`
	Attached bool   `json:"attached"`
}

// Copy of `target.GetTargetsResult`, for parsing target states.
type getTargetsResult struct {
	TargetInfos []targetInfo `json:"targetInfos"`
}

// FromContext returns the `devtools.Session` stored in the given `context.Context`,
// if that context was initialized with the `devtools.NewContext` function.
func FromContext(ctx context.Context) (*Session, bool) {
	s, ok := ctx.Value(sessionKey{}).(*Session)
	if ok && s == nil {
		ok = false
	}
	return s, ok
}

// SessionOption is used for customization in the `devtools.NewContext` function.
//
// See https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html.
type SessionOption = func(*Session)

// NewContext constructs a new `devtools.Session`, and returns a copy of the
// parent context which carries this new session.
//
// If the parent context already carries a CDP session, this function reuses
// that session's browser and opens a new tab in it. Otherwise, this function
// creates a new output directory and starts a new browser.
//
// The output directory contains the browser's STDOUT and STDERR, user data
// directory (unless the caller overrides the browser flag "user-data-dir"),
// proxy logs, screenshots, etc.
//
// The default path for output directories is Go's `os.TempDir()`, but it can be
// overridden with an optional environment variable (see `devtools.OutputRootEnv`).
func NewContext(parent context.Context, opts ...SessionOption) (context.Context, error) {
	// Store the new session in a cancelable copy of the parent context.
	ctx, cancel := context.WithCancel(parent)
	session := &Session{cancel: cancel}
	ctx = context.WithValue(ctx, sessionKey{}, session)

	// Report when this context will be canceled. No need to clean-up
	// anything: the cancelation of the context automatically kills the
	// browser, which triggers its own cleanup: see the goroutine at the
	// bottom of the start function in browser.go.
	go func() {
		<-ctx.Done()
		log.Printf("CDP context ending reason: %v", ctx.Err())
	}()

	if ps, ok := FromContext(parent); ok {
		// Reuse the existing session stored in the parent context.
		session.cancel = ps.cancel

		session.OutputDir = ps.OutputDir
		session.UserDataDir = ps.UserDataDir

		session.browserDone = ps.browserDone
		session.browserInputWriter = ps.browserInputWriter
		session.browserOutputReader = ps.browserOutputReader
		session.webSocket = ps.webSocket

		session.msgLog = ps.msgLog
		session.msgID = ps.msgID
		session.msgQ = ps.msgQ

		session.responseSubscribers = ps.responseSubscribers
		session.eventSubscribers = ps.eventSubscribers

		// Open a new tab.
		session.TargetID, session.SessionID = newSafeString(), newSafeString()
		targetID, err := createTarget(ctx)
		if err != nil {
			session.cancel()
			return parent, fmt.Errorf(`"Target.createTarget" command error: %v`, err)
		}
		session.TargetID.Write(targetID)
	} else {
		// Construct a new session.

		// Set optional session options specified by the caller, if any.
		for _, o := range opts {
			o(session)
		}
		// Initialize the session's output directory.
		path, err := mkdirOutput()
		if err != nil {
			err = fmt.Errorf("failed to create CDP output directory (%s): %v", path, err)
			return parent, err
		}
		session.OutputDir = path
		// Initialize a new log file for incoming/outgoing JSON messages.
		f, err := os.Create(filepath.Join(path, "cdp_json.log"))
		if err != nil {
			return parent, fmt.Errorf("failed to initialize CDP log file: %v", err)
		}
		session.msgLog = log.New(f, "", log.Ldate|log.Ltime|log.Lmicroseconds)
		// Start a new browser.
		if err := start(ctx, session); err != nil {
			return parent, err
		}
		// Initialize channels to send JSON messages to and from the browser.
		session.msgID = 1
		session.msgQ = make(chan asyncMessage)
		session.responseSubscribers = make(map[int64]chan *Message)
		session.eventSubscribers = make(map[string][]chan *Message)
		go func(s *Session) {
			for {
				asyncMsg, ok := <-s.msgQ
				if !ok {
					// s.msgQ is closed when the browser process ends (see the
					// goroutine at the bottom of the start function in browser.go).
					return
				}
				if runtime.GOOS != "windows" {
					sendToPipe(s, asyncMsg)
				} else {
					sendToWebSocket(s, asyncMsg)
				}
				s.msgID++
			}
		}(session)

		// Attach this session to the first tab.
		session.TargetID, session.SessionID = newSafeString(), newSafeString()
		targetID, err := fetchTargetID(ctx)
		if err != nil {
			session.cancel()
			return parent, fmt.Errorf(`"Target.getTargets" command error: %v`, err)
		}
		session.TargetID.Write(targetID)
	}

	// Finish attaching this session to a browser tab.
	sessionID, err := attach(ctx, session.TargetID.Read())
	if err != nil {
		session.cancel()
		return parent, fmt.Errorf(`"Target.attachToTarget" command error: %v`, err)
	}
	log.Printf("Target ID: %s", session.TargetID.Read())
	log.Printf("Session ID: %s", sessionID)
	session.SessionID.Write(sessionID)

	// Enable receiving various asynchronous events from the browser.
	if _, err := SendAndWait(ctx, "Page.enable", nil); err != nil {
		session.cancel()
		return parent, err
	}
	params := []byte(`{"enabled":true}`)
	if _, err := SendAndWait(ctx, "Page.setLifecycleEventsEnabled", params); err != nil {
		session.cancel()
		return parent, err
	}

	return ctx, nil
}

// Create a uniquely-named session output directory.
func mkdirOutput() (string, error) {
	path, ok := os.LookupEnv(OutputRootEnv)
	if !ok {
		path = os.TempDir()
	}
	dir := time.Now().UTC().Format("chrome_vision_20060102_150405.000000000")
	path = filepath.Join(path, dir)
	if err := os.MkdirAll(path, 0755); err != nil {
		return "", nil
	}
	log.Printf("Session output directory: %s", path)
	return path, nil
}

// Create a new browser tab and return its target ID.
func createTarget(ctx context.Context) (string, error) {
	// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-createTarget
	// (we don't use the target sub-package to avoid circular dependencies).
	response, err := SendAndWait(ctx, "Target.createTarget", []byte(`{"url":""}`))
	if err != nil {
		return "", err
	}
	if response.Error != nil {
		return "", errors.New(response.Error.Error())
	}
	result := &createTargetResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return "", err
	}
	return result.TargetID, nil
}

// Return the ID of the browser's first unattached page target.
func fetchTargetID(ctx context.Context) (string, error) {
	ticker := time.NewTicker(10 * time.Millisecond)
	timer := time.NewTimer(TargetTimeout)
	defer ticker.Stop()
	defer timer.Stop()
	var lastErr error
	for {
		select {
		case <-ticker.C:
			// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-getTargets
			// (we don't use the target sub-package to avoid circular dependencies).
			response, err := SendAndWait(ctx, "Target.getTargets", nil)
			if err != nil {
				lastErr = err
				break
			}
			if response.Error != nil {
				lastErr = errors.New(response.Error.Error())
				break
			}
			result := &getTargetsResult{}
			if err := json.Unmarshal(response.Result, result); err != nil {
				lastErr = err
				break
			}
			for _, t := range result.TargetInfos {
				if t.Type == "page" && !t.Attached {
					return t.TargetID, nil // Success.
				}
			}
			lastErr = errors.New("unattached page target not found")
		case <-timer.C:
			return "", lastErr // Timeout - return the last recorded error.
		}
	}
}

// Attach to the given target and return the resulting session ID.
func attach(ctx context.Context, targetID string) (string, error) {
	// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-attachToTarget
	// (we don't use the target sub-package to avoid circular dependencies).
	params := fmt.Sprintf(`{"targetId":%q,"flatten":true}`, targetID)
	response, err := SendAndWait(ctx, "Target.attachToTarget", []byte(params))
	if err != nil {
		return "", err
	}
	if response.Error != nil {
		return "", errors.New(response.Error.Error())
	}
	result := &Message{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return "", err
	}
	return result.SessionID, nil
}

// BrowserPath allows the caller of the `devtools.NewContext` function to
// force this Go package to run the browser from a specific path, instead of
// looking for the first available OS-specific default path.
//
// This is useful if you want to run a non-default browser or build.
func BrowserPath(path string) SessionOption {
	return func(s *Session) {
		s.browserPath = &path
	}
}

// UserDataDir allows the caller of the `devtools.NewContext` function to
// customize the browser's default user data directory.
//
// This is useful if you want to customize or monitor the content of the user
// profile directory.
func UserDataDir(path string) SessionOption {
	return func(s *Session) {
		s.UserDataDir = path
	}
}

// BrowserFlags allows the caller of the `devtools.NewContext` function to
// override this Go package's default browser flags, retrieved with the
// function `devtools.DefaultBrowserFlags`.
//
// This is useful if you want to customize the browser's behavior.
func BrowserFlags(flags map[string]interface{}) SessionOption {
	return func(s *Session) {
		s.browserFlags = make(map[string]interface{})
		for k, v := range flags {
			s.browserFlags[k] = v
		}
	}
}
