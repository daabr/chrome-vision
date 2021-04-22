// Package cdp provides Go bindings for all the commands, events and
// types in the Chrome DevTools Protocol (CDP) - from the latest Chromium
// "tip-of-tree" (tot) definitions (see the details and API documentation
// in https://chromedevtools.github.io/devtools-protocol/tot, mirrored in
// https://github.com/ChromeDevTools/devtools-protocol). Each sub-package
// here corresponds to a single "domain" in the CDP.
package cdp

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

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
	log.Printf("Browser executable path: %s\n", *p)

	// Initialize the command-line.
	if len(s.browserFlags) == 0 {
		s.browserFlags = DefaultBrowserFlags()
	}
	args := append(adjustFlags(s), "about:blank")
	log.Printf("Browser command-line args: %q\n", args)
	cmd := exec.CommandContext(ctx, *p, args...)

	// Ensure that the user data directory exists, whether it's the
	// default (in the output directory) or a custom one specified by the caller.
	// If starting the browser fails immediately, delete the unused output directory.
	os.MkdirAll(*s.UserDataDir, 0755)
	defer func() {
		if cmd.Process == nil {
			os.RemoveAll(*s.OutputDir)
		}
	}()

	// Redirect the browser process's output to files.
	stdout, err := os.Create(filepath.Join(*s.OutputDir, "stdout.txt"))
	if err != nil {
		return fmt.Errorf("failed to initialize browser process's stdout file: %v", err)
	}
	stderr, err := os.Create(filepath.Join(*s.OutputDir, "stderr.txt"))
	if err != nil {
		return fmt.Errorf("failed to initialize browser process's stderr file: %v", err)
	}
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	// On POSIX-compliant operating systems, prepare input and output pipes to
	// communicate with the browser process. On Windows, we have to use websockets.
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
	// TODO: initialize websockets otherwise.

	// Start a new broswer process.
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start browser process: %v", err)
	}
	log.Printf("Browser process started: PID %d\n", cmd.Process.Pid)
	s.browserDone = make(chan struct{})

	// On POSIX-compliant operating systems, start using the
	// previously-prepared pipes to communicate with the browser.
	if runtime.GOOS != "windows" {
		go receiveFromPipe(s)
	}
	// TODO: use websockets otherwise.

	// Wait in the background for the browser process to end, and clean-up
	// any resources associated with it. See also the goroutine in the
	// NewContext function in session.go.
	go func(s *Session, c *exec.Cmd) {
		log.Println("Waiting for the browser process to end")
		if err := c.Wait(); err != nil {
			log.Printf("Browser process has ended with an error: %v\n", err)
		} else {
			log.Println("Browser process has ended without an error")
		}
		s.cancel()

		close(s.msgQ)
		s.browserInputWriter.Close()
		s.browserOutputReader.Close()
		s.msgLog.Writer().(*os.File).Sync()
		s.msgLog.Writer().(*os.File).Close()
		// TODO: unsubscribe (close channels) for all existing subscribers.
		close(s.browserDone)
	}(s, cmd)

	return nil
}

// Cancel cancels the given context, if it was initialized with the
// `cdp.NewContext` function, and returns immediately.
//
// This will kill the browser forcefully. If you want to let it close
// gracefully, call the `cdp.Close` function instead. Either way, any resources
// associated with the CDP session will be released when the process ends.
//
// Remember that this does not impact ancestor or parent contexts specified
// in the `cdp.NewContext` function call, only contexts returned by it.
// Context cancelation propogates only from ancestors to descendants.
func Cancel(ctx context.Context) {
	if session, ok := FromContext(ctx); ok {
		session.cancel()
	}
}

// Wait waits for the browser associated with the given context (which was
// started by the `cdp.NewContext` function) to end - no matter how - and for
// any resources associated with the CDP session to be released.
func Wait(ctx context.Context) {
	if session, ok := FromContext(ctx); ok {
		if session.browserDone != nil {
			<-session.browserDone
		}
	}
}

// Close sends to the browser a command to close itself gracefully,
// and waits until this is done.
//
// If you want to kill the browser forcefully, call `cdp.Cancel` instead.
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
