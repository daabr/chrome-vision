package cdp

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
)

// defaultBrowserFlags is largely based on the results of other related projects:
//
// • https://source.chromium.org/chromium/chromium/src/+/master:chrome/test/chromedriver/chrome_launcher.cc?q=k.*Switches
//
// • https://github.com/puppeteer/puppeteer/blob/main/src/node/Launcher.ts
//
// • https://github.com/chromedp/chromedp/blob/master/allocate.go
//
// • https://github.com/GoogleChrome/chrome-launcher/blob/master/src/flags.ts
var defaultBrowserFlags = map[string]interface{}{
	// https://peter.sh/experiments/chromium-command-line-switches/#disable-background-networking
	// (includes extension updating, safe browsing service, upgrade detector, translate, UMA)
	"disable-background-networking":          true,
	"disable-background-timer-throttling":    true,
	"disable-backgrounding-occluded-windows": true,
	"disable-breakpad":                       true,
	// https://github.com/mjzffr/puppeteer/commit/4af0911b90773e59c3ff552f6c4cc6e42af715a
	"disable-client-side-phishing-detection": true,
	// https://github.com/puppeteer/puppeteer/pull/4704
	"disable-component-extensions-with-background-pages": true,
	"disable-default-apps":                               true,
	// https://peter.sh/experiments/chromium-command-line-switches/#disable-dev-shm-usage
	// https://github.com/puppeteer/puppeteer/issues/1834
	"disable-dev-shm-usage": true,
	// https://github.com/puppeteer/puppeteer/blob/main/docs/troubleshooting.md#chrome-headless-doesnt-launch-on-windows
	"disable-extensions":   true,
	"disable-features":     "Translate",
	"disable-hang-monitor": true,
	// https://github.com/puppeteer/puppeteer/pull/3474
	"disable-ipc-flooding-protection": true,
	// https://crrev.com/725724
	"disable-popup-blocking":         true,
	"disable-prompt-on-repost":       true,
	"disable-renderer-backgrounding": true,
	"disable-sync":                   true,
	// https://github.com/GoogleChrome/chrome-launcher/blob/master/docs/chrome-flags-for-tools.md#test--debugging-flags
	"enable-automation": true,
	// https://github.com/puppeteer/puppeteer/pull/6410
	"enable-blink-features": "IdleDetection",
	// https://github.com/puppeteer/puppeteer/pull/3738
	"enable-features": "NetworkService,NetworkServiceInProcess",
	// https://github.com/puppeteer/puppeteer/pull/3732
	"force-color-profile": "srgb",
	// https://developers.google.com/web/updates/2017/04/headless-chrome
	// https://chromium.googlesource.com/chromium/src/+/lkgr/headless/README.md
	"headless":                 true,
	"metrics-recording-only":   true,
	"mute-audio":               true,
	"no-default-browser-check": true,
	"no-first-run":             true,
	// https://chromium.googlesource.com/chromium/src/+/master/docs/linux/password_storage.md
	"password-store":    "basic",
	"use-mock-keychain": true,
}

// DefaultBrowserFlags returns a map of default command-line flags for starting the browser.
//
// The caller may modify this map for their particular use-cases. For more information, see
// https://github.com/GoogleChrome/chrome-launcher/blob/master/docs/chrome-flags-for-tools.md
// and https://peter.sh/experiments/chromium-command-line-switches.
//
// Examples of a few other common flags to consider:
// 	flags := cdp.DefaultBrowserFlags()
// 	flags["allow-running-insecure-content"] = true
// 	flags["auto-open-devtools-for-tabs"] = true
// 	flags["disable-device-discovery-notifications"] = true
// 	flags["disable-gpu"] = true // https://crbug.com/765284
// 	flags["disable-web-security"] = true
// 	flags["enable-logging"] = "stderr"
// 	flags["ignore-certificate-errors"] = true
// 	flags["js-flags"] = "--random-seed=1157259157" // Initialize V8's RNG with a fixed seed.
// 	flags["log-level"] = "0", // INFO = 0, WARNING = 1, LOG_ERROR = 2, LOG_FATAL = 3.
// 	flags["proxy-server"] = "..."
// 	flags["start-maximized"] = true
// 	flags["use-fake-device-for-media-stream"] = true // Required for the next flag.
// 	flags["use-file-for-fake-video-capture"] = "<path-to-file>" // .y4m or .mjpeg.
// 	flags["user-agent"] = "..."
// 	flags["windows-position"] = "x,y" // "x" and "y" are non-negative integers.
// 	flags["window-size"] = "x,y" // "x" and "y" are positive integers.
func DefaultBrowserFlags() map[string]interface{} {
	// Initialize a copy of the default map.
	copy := make(map[string]interface{})
	for k, v := range defaultBrowserFlags {
		copy[k] = v
	}
	// Runtime customizations.
	if os.Getuid() == 0 {
		// https://chromium.googlesource.com/chromium/src.git/+/master/docs/linux/sandboxing.md
		// https://github.com/puppeteer/puppeteer/blob/main/docs/troubleshooting.md#setting-up-chrome-linux-sandbox
		copy["no-sandbox"] = true
	}
	return copy
}

func adjustFlags(s *Session) []string {
	// Runtime adjustments to set-up communication with the browser process.
	if runtime.GOOS != "windows" {
		// Prefer pipes over websockets in POSIX-compliant operating systems.
		s.browserFlags["remote-debugging-pipe"] = true
		delete(s.browserFlags, "remote-debugging-port")
	} else {
		// https://golang.org/pkg/os/exec/#Cmd.ExtraFiles
		// is not supported by Go in Windows.
		delete(s.browserFlags, "remote-debugging-pipe")
		if _, ok := s.browserFlags["remote-debugging-port"]; !ok {
			s.browserFlags["remote-debugging-port"] = "0"
		}
	}

	// Allow the caller to specify their own directory, but if they don't
	// then create a new one - for isolation. The existence of the directory,
	// (in the session's output directory, or a custom path provided by the user)
	// is ensured right after calling this function.
	if s.UserDataDir == nil {
		path := filepath.Join(*s.OutputDir, "user_data")
		s.UserDataDir = &path
	}
	s.browserFlags["user-data-dir"] = *s.UserDataDir

	// Convert the map to a sorted slice.
	var args, keys []string
	for k := range s.browserFlags {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		flag := "--" + k
		switch v := s.browserFlags[k].(type) {
		case bool:
			if v {
				args = append(args, flag)
			}
		default:
			args = append(args, fmt.Sprintf("%s=%v", flag, v))
		}
	}
	return args
}
