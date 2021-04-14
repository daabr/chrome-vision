// The types and functions in this file follow the example in comment of the
// context.Context.Value field in https://golang.org/pkg/context/#Context.
// See also https://golang.org/pkg/context/#example_WithValue.

package cdp

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// OutputRootEnv is the name of an optional environment variable to override
// the default path for CDP output directories instead of Go's os.TempDir().
const OutputRootEnv = "CDP_OUTPUT_ROOT"

// Session contains CDP runtime details, stored in a context.Context,
// and retrievable with the cdp.FromContext function.
type Session struct {
	cancel context.CancelFunc

	OutputDir, UserDataDir *string

	browserPath  *string
	browserFlags map[string]interface{}
	// TODO: environment variables.
	browserDone chan struct{}

	browserInputWriter, browserOutputReader *os.File

	msgLog *log.Logger
	msgID  int64
	msgQ   chan asyncMessage
}

type sessionKey struct{}

// FromContext returns the cdp.Session stored in the given context.Context,
// if that context was initialized with the cdp.NewContext function.
func FromContext(ctx context.Context) (*Session, bool) {
	s, ok := ctx.Value(sessionKey{}).(*Session)
	if ok && s == nil {
		ok = false
	}
	return s, ok
}

// SessionOption is used for customization in the cdp.NewContext function.
// See https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html.
type SessionOption = func(*Session)

// NewContext constructs a new cdp.Session, and returns a copy of the parent
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
// The default path for output directories is Go's os.TempDir(), but it can be
// overriden with an optional environment variable (see cdp.OutputRootEnv).
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
		// Initialize channel to send JSON messages to the browser.
		session.msgID = 1
		session.msgQ = make(chan asyncMessage)
		go func(s *Session) {
			for {
				asyncMsg, ok := <-s.msgQ
				if !ok {
					return
				}
				sendToPipe(s, asyncMsg)
				s.msgID++
			}
		}(session)
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

// mkdirOutput creates a uniquely-named session output directory.
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

// BrowserPath allows the caller of the cdp.NewContext function to force this
// Go package to run the browser from a specific path, instead of looking
// for the first available OS-specific default path.
//
// This is useful if you want to run a non-default browser or build.
func BrowserPath(path string) SessionOption {
	return func(s *Session) {
		s.browserPath = &path
	}
}

// UserDataDir allows the caller of the cdp.NewContext function to customize
// the browser's default user data directory.
//
// This is useful if you want to customize or monitor the content of the user
// profile directory.
func UserDataDir(path string) SessionOption {
	return func(s *Session) {
		s.UserDataDir = &path
	}
}

// BrowserFlags allows the caller of the cdp.NewContext function to override
// this Go package's default browser flags, retrieved with the function
// cdp.DefaultBrowserFlags.
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
