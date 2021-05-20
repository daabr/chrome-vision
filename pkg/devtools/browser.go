// Package devtools provides Go bindings for all the commands, events and
// types in the Chrome DevTools Protocol (CDP) - from the latest Chromium
// "tip-of-tree" (tot) definitions (see the details and API documentation
// in https://chromedevtools.github.io/devtools-protocol/tot, mirrored in
// https://github.com/ChromeDevTools/devtools-protocol). Each sub-package
// here corresponds to a single "domain" in the CDP.
package devtools

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"time"

	"github.com/daabr/chrome-vision/pkg/websocket"
)

// WebSocketTimeout is the maximum amount of time to wait for the browser to
// report its WebSocket address via STDERR, in order to perform a handshake.
const WebSocketTimeout = 1 * time.Minute

type windowsStderrWriter struct {
	session *Session
	stderr  *os.File
}

// Write passes lines from the browser's STDERR to our log file, but also extract
// and remember the browser's WebSocket address for initialization.
func (w *windowsStderrWriter) Write(b []byte) (n int, err error) {
	if w.session.wsAddress.read() == "" {
		re := regexp.MustCompile(`ws://(.+:\d+)(/devtools/browser/[\w-]{36})`)
		m := re.FindSubmatch(b)
		if m != nil {
			log.Printf("WebSocket address: %s", m[0])
			w.session.wsAddress.write(string(m[1]))
			w.session.wsPath.write(string(m[2]))
		}
	}
	return w.stderr.Write(b)
}

func start(ctx context.Context, s *Session) error {
	// Find the executable if it's not specified. Compare with:
	// https://github.com/karma-runner/karma-chrome-launcher/blob/master/index.js
	// https://github.com/SeleniumHQ/selenium/wiki/ChromeDriver#requirements
	p := s.browserPath
	if p == nil {
		for _, e := range executables {
			if path, err := exec.LookPath(e); err == nil {
				p = &path
				break
			}
		}
		if p == nil {
			return exec.ErrNotFound
		}
	}
	log.Printf("Browser executable path: %s", *p)

	// Initialize the command-line.
	if len(s.browserFlags) == 0 {
		s.browserFlags = DefaultBrowserFlags()
	}
	args := append(adjustFlags(s), "about:blank")
	log.Printf("Browser command-line args: %q", args)
	cmd := exec.CommandContext(ctx, *p, args...)

	// Ensure that the user data directory exists, whether it's the
	// default (in the output directory) or a custom one specified by the caller.
	// If starting the browser fails immediately, delete the unused output directory.
	os.MkdirAll(s.UserDataDir, 0755)
	defer func() {
		if cmd.Process == nil {
			os.RemoveAll(s.OutputDir)
		}
	}()

	// Redirect the browser process's output to files.
	stdout, err := os.Create(filepath.Join(s.OutputDir, "stdout.txt"))
	if err != nil {
		return fmt.Errorf("failed to initialize browser process's STDOUT file: %v", err)
	}
	stderr, err := os.Create(filepath.Join(s.OutputDir, "stderr.txt"))
	if err != nil {
		return fmt.Errorf("failed to initialize browser process's STDERR file: %v", err)
	}
	cmd.Stdout = stdout
	if runtime.GOOS != "windows" {
		cmd.Stderr = stderr
	} else {
		// On Windows, we also need to read STDERR during runtime,
		// in order to know how to communicate with the browser.
		s.wsAddress, s.wsPath = newSafeString(), newSafeString()
		cmd.Stderr = &windowsStderrWriter{session: s, stderr: stderr}
	}

	// On POSIX-compliant operating systems, prepare input and output pipes to
	// communicate with the browser process. On Windows, we use a WebSocket.
	if runtime.GOOS != "windows" {
		inputReader, inputWriter, err := os.Pipe()
		if err != nil {
			return fmt.Errorf("failed to initialize browser input pipe: %v", err)
		}
		outputReader, outputWriter, err := os.Pipe()
		if err != nil {
			return fmt.Errorf("failed to initialize browser output pipe: %v", err)
		}
		s.browserInputWriter, s.browserOutputReader = inputWriter, outputReader
		cmd.ExtraFiles = []*os.File{inputReader, outputWriter}
	}

	// Start a new broswer process.
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start browser process: %v", err)
	}
	log.Printf("Browser process started: PID %d", cmd.Process.Pid)
	s.browserDone = make(chan struct{})

	if runtime.GOOS != "windows" {
		// On POSIX-compliant operating systems, start using the
		// previously-prepared pipes to communicate with the browser.
		go receiveFromPipe(s)
	} else {
		// On Windows, initialize a WebSocket to communicate with the browser.
		ticker := time.NewTicker(10 * time.Millisecond)
		timer := time.NewTimer(WebSocketTimeout)
		defer ticker.Stop()
		defer timer.Stop()
		for s.wsPath.read() == "" {
			select {
			case <-ticker.C:
				continue // Recheck the loop condition.
			case <-timer.C:
				s.cancel()
				return errors.New("failed to detect WebSocket address in STDERR")
			}
		}
		conn, err := websocket.Handshake(ctx, s.wsAddress.read(), s.wsPath.read())
		if err != nil {
			s.cancel()
			return err
		}
		s.webSocket = conn
		go receiveFromWebSocket(s)
	}

	// Wait in the background for the browser process to end, and clean-up
	// any resources associated with it. See also the goroutine in the
	// NewContext function in session.go.
	go func(s *Session, c *exec.Cmd) {
		log.Print("Waiting in the background for the browser process to end")
		if err := c.Wait(); err != nil {
			log.Printf("Browser process has ended with an error: %v", err)
		} else {
			log.Print("Browser process has ended without an error")
		}
		s.cancel()

		close(s.msgQ)
		if runtime.GOOS != "windows" {
			s.browserInputWriter.Close()
			s.browserOutputReader.Close()
		}
		s.msgLog.Writer().(*os.File).Sync()
		s.msgLog.Writer().(*os.File).Close()
		// TODO: unsubscribe (close channels) for all existing subscribers.
		close(s.browserDone)
	}(s, cmd)

	return nil
}

// Cancel cancels the given context, if it was initialized with the
// `devtools.NewContext` function, and returns immediately.
//
// This will kill the browser forcefully. If you want to close it gracefully,
// call the `devtools.Close` function instead. Either way, any resources
// associated with the CDP session will be released when the process ends.
//
// Remember that this does not impact ancestor or parent contexts specified
// in the `devtools.NewContext` function call, only contexts returned by it.
// Context cancelation propagates only from ancestors to descendants.
func Cancel(ctx context.Context) {
	if session, ok := FromContext(ctx); ok {
		session.cancel()
	}
}

// Wait waits for the browser associated with the given context (which was
// started by the `devtools.NewContext` function) to end - no matter how - and
// for any resources associated with the CDP session to be released.
func Wait(ctx context.Context) {
	if session, ok := FromContext(ctx); ok {
		if session.browserDone != nil {
			<-session.browserDone
		}
	}
}

// Close sends to the browser a command to close itself gracefully, and waits
// until this is done.
//
// If you want to kill the browser forcefully, call `devtools.Cancel` instead.
// Either way, any resources associated with the CDP session will be
// released when the process ends.
func Close(ctx context.Context) {
	if _, ok := FromContext(ctx); ok {
		// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-close
		// (we don't use the browser sub-package to avoid circular dependencies).
		Send(ctx, "Browser.close", nil)
		Wait(ctx)
	}
}
