# Chrome Vision

[![Go Reference](https://pkg.go.dev/badge/github.com/daabr/chrome-vision.svg)](https://pkg.go.dev/github.com/daabr/chrome-vision)
[![Go Report Card](https://goreportcard.com/badge/github.com/daabr/chrome-vision)](https://goreportcard.com/report/github.com/daabr/chrome-vision)
[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/daabr/chrome-vision/Go/main)](https://github.com/daabr/chrome-vision/actions/workflows/go.yml?query=branch:main)

Advanced cross-platform web automation with a convenient Go API.

This project is based on the
[Chrome DevTools Protocol](https://chromedevtools.github.io/devtools-protocol)
(CDP) and
[WebDriver specification](https://w3c.github.io/webdriver/webdriver-spec.html),
to control web sessions in Chrome and other Blink-based browsers.

It improves, consolidates and extends the approaches of
[Selenium WebDriver](https://www.selenium.dev/documentation),
[Puppeteer](https://pptr.dev) and [ChromeDP](https://github.com/chromedp/chromedp).
The first two are very popular, but only the last one offers integration with
[Go](https://golang.org) without heavy non-Go dependencies, albeit with a very
limited high-level API.

## Goals

* Make the Chrome DevTools Protocol more accessible
* Simplify the low-level API even more than ChromeDP
* Support the WebDriver specification as a higher-level API layer,
  **together** with the lower-level CDP API
* Add Computer Vision (CV) and Optical Character Recognition (OCR)
  capabilities, which are missing in all of the above

## Differences From ChromeDP

* Simpler session initialization, and simpler customization of browser flags
  (see [example](./examples/init/main.go))

* Simpler CDP command execution API:

  * More idiomatic command execution, instead of limited JavaScript-await-like
    wrapper functions
    ([`chromedp.Run`](https://pkg.go.dev/github.com/chromedp/chromedp#Run),
    [`chromedp.RunResponse`](https://pkg.go.dev/github.com/chromedp/chromedp#RunResponse),
    [`chromedp.ActionFunc`](https://pkg.go.dev/github.com/chromedp/chromedp#ActionFunc)):
    * **Synchronous/blocking execution (simplest):** the command's `Do`
      function waits for the browser's response and returns it as a parsed
      struct (see [example](./examples/navigation/sync/main.go))
    * **Asynchronous/non-blocking execution (more advanced):** the command's
      `Start` function returns a Go channel to receive from the browser a raw
      [`devtools.Message`](https://pkg.go.dev/github.com/daabr/chrome-vision/pkg/devtools#Message),
      and the optional `ParseResponse` function parses it (see
      [example](./examples/navigation/async/main.go))
    * **Return values:** for commands with output, `Do` and `ParseResponse`
      return a single struct, to avoid the clutter and brittleness of many
      return value assignments
    * **Errors:** both Go and CDP errors are returned as Go errors by `Do` and
      `ParseResponse`, no need for multiple error checks per command

  * No need to initialize executors, either with
    [`chromedp.Run`](https://pkg.go.dev/github.com/chromedp/chromedp#Run) or
    with
    [`cdproto.cdp.WithExecutor`](https://pkg.go.dev/github.com/chromedp/cdproto/cdp#WithExecutor)

* Communication with the browser:
  * **Non-Windows operating systems:** using POSIX pipes instead of WebSockets
    (faster, more reliable and more secure)
  * **Windows:** using an internal implementation of the WebSocket protocol
    (optimized for Chrome DevTools, faster and more efficient - see
    [documentation](https://pkg.go.dev/github.com/daabr/chrome-vision/pkg/websocket))

* Stronger adherence to idiomatic Go coding style and
  <https://github.com/golang-standards/project-layout>
