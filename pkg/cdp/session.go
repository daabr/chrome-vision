// The types and functions in this file follow the example in comment of the
// context.Context.Value field in https://golang.org/pkg/context/#Context.
// See also https://golang.org/pkg/context/#example_WithValue.

package cdp

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// OutputRootEnv is the name of an optional environment variable to override
// the default path for CDP output directories instead of Go's `os.TempDir()`.
const OutputRootEnv = "CDP_OUTPUT_ROOT"

// Session contains CDP runtime details, stored in a `context.Context`,
// and retrievable with the `cdp.FromContext` function.
type Session struct {
	// For immediate canceling of the context returned by cdp.NewContext,
	// and any descendant contexts, when the browser process ends.
	cancel context.CancelFunc

	OutputDir, UserDataDir *string

	// Browser execution details. Not shared with descendant contexts because
	// the browser was already started by the first call to cdp.NewContext.
	browserPath  *string
	browserFlags map[string]interface{}
	// TODO: environment variables.

	browserDone chan struct{}

	browserInputWriter, browserOutputReader *os.File

	msgLog *log.Logger
	msgID  int64
	msgQ   chan asyncMessage // https://blog.golang.org/codelab-share

	// Exactly one subscriber per response (created and used in cdp.Send).
	responseSubscribers map[int64]chan *Message
	// Zero or more subscribers per event type.
	eventSubscribers map[string][]chan *Message

	// Metadata about the browser, its tabs (pages) and their iframes.
	targets map[string]targetInfo
	// The attached target. Not shared with descendant contexts
	// because they create their own tabs, targets and sessions IDs. See also:
	// https://github.com/aslushnikov/getting-started-with-cdp#targets--sessions.
	targetID, sessionID string
}

type sessionKey struct{}

// FromContext returns the `cdp.Session` stored in the given `context.Context`,
// if that context was initialized with the `cdp.NewContext` function.
func FromContext(ctx context.Context) (*Session, bool) {
	s, ok := ctx.Value(sessionKey{}).(*Session)
	if ok && s == nil {
		ok = false
	}
	return s, ok
}

// SessionOption is used for customization in the `cdp.NewContext` function.
//
// See https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html.
type SessionOption = func(*Session)

// NewContext constructs a new `cdp.Session`, and returns a copy of the parent
// context which carries this new session.
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
// overriden with an optional environment variable (see `cdp.OutputRootEnv`).
func NewContext(parent context.Context, opts ...SessionOption) (context.Context, error) {
	// Store the new session in a cancelable copy of the parent context.
	ctx, cancel := context.WithCancel(parent)
	session := &Session{cancel: cancel}
	ctx = context.WithValue(ctx, sessionKey{}, session)

	if ps, ok := FromContext(parent); ok {
		// Reuse the existing session stored in the parent context.
		session.cancel = ps.cancel

		session.OutputDir = ps.OutputDir
		session.UserDataDir = ps.UserDataDir

		session.browserDone = ps.browserDone
		session.browserInputWriter = ps.browserInputWriter
		session.browserOutputReader = ps.browserOutputReader

		session.msgLog = ps.msgLog
		session.msgID = ps.msgID
		session.msgQ = ps.msgQ

		session.responseSubscribers = ps.responseSubscribers
		session.eventSubscribers = ps.eventSubscribers

		session.targets = ps.targets

		// TODO: open a new tab, attach to the target, and set the session ID.
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
		session.OutputDir = &path
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
				sendToPipe(s, asyncMsg)
				s.msgID++
			}
		}(session)

		// Enable and receive target update events.
		session.targets = make(map[string]targetInfo)
		events := [...]string{
			"Target.targetCreated",
			"Target.targetInfoChanged",
			"Target.targetDestroyed",
		}
		channels := []chan *Message{}
		for _, e := range events {
			ch, err := SubscribeEvent(ctx, e)
			if err != nil {
				log.Printf("WARNING: %q subscription error: %v", e, err)
			}
			channels = append(channels, ch)
		}
		go receiveTargetUpdates(ctx, channels)

		// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-setDiscoverTargets
		// (we don't use the target sub-package to avoid circular dependencies).
		method, params := "Target.setDiscoverTargets", `{"discover":true}`
		if _, err := Send(ctx, method, []byte(params)); err != nil {
			log.Printf("WARNING: %q command error: %v", method, err)
		}
		for _, t := range session.targets {
			if t.Type == "page" && !t.Attached {
				// Attach to the new page target to get a session ID
				// (https://github.com/aslushnikov/getting-started-with-cdp#targets--sessions).
				session.targetID = t.TargetID
				log.Printf("Target ID: %s", t.TargetID)
				session.sessionID = attachToTarget(ctx, t.TargetID)
				log.Printf("Session ID: %s", session.sessionID)
			}
		}

		// TODO: Runtime.enable, Log.enable, Network.enable, Inspector.enable,
		// Page.enable, DOM.enable, CSS.enable, Page.setLifecycleEventsEnabled?

		// Report when the context will be canceled. No need to clean-up
		// anything: the cancelation of the context automatically kills the
		// browser, which triggers its own cleanup: see the goroutine at the
		// bottom of the start function in browser.go.
		go func() {
			<-ctx.Done()
			log.Printf("CDP context ending reason: %v\n", ctx.Err())
		}()
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
	log.Printf("Session output directory: %s\n", path)
	return path, nil
}

// Attach to the given target, and return the resulting session ID. Based on
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-attachToTarget.
func attachToTarget(ctx context.Context, targetID string) string {
	// We don't use the target sub-package to avoid circular dependencies.
	params := fmt.Sprintf(`{"targetId":%q,"flatten":true}`, targetID)
	response, err := Send(ctx, "Target.attachToTarget", []byte(params))
	if err != nil {
		log.Printf(`WARNING: "attachToTarget" command error: %v`, err)
	}
	result := &Message{}
	json.Unmarshal(response.Result, result)
	return result.SessionID
}

// BrowserPath allows the caller of the `cdp.NewContext` function to force
// this Go package to run the browser from a specific path, instead of looking
// for the first available OS-specific default path.
//
// This is useful if you want to run a non-default browser or build.
func BrowserPath(path string) SessionOption {
	return func(s *Session) {
		s.browserPath = &path
	}
}

// UserDataDir allows the caller of the `cdp.NewContext` function to customize
// the browser's default user data directory.
//
// This is useful if you want to customize or monitor the content of the user
// profile directory.
func UserDataDir(path string) SessionOption {
	return func(s *Session) {
		s.UserDataDir = &path
	}
}

// BrowserFlags allows the caller of the `cdp.NewContext` function to override
// this Go package's default browser flags, retrieved with the function
// `cdp.DefaultBrowserFlags`.
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
