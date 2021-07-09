# Chrome Vision

[![Go Reference](https://pkg.go.dev/badge/github.com/daabr/chrome-vision.svg)](https://pkg.go.dev/github.com/daabr/chrome-vision)
[![Go Report Card](https://goreportcard.com/badge/github.com/daabr/chrome-vision)](https://goreportcard.com/report/github.com/daabr/chrome-vision)
[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/daabr/chrome-vision/Go/main)](https://github.com/daabr/chrome-vision/actions/workflows/go.yml?query=branch:main)

Advanced cross-platform web automation with a convenient Go API.

This is based on the
[Chrome DevTools Protocol](https://chromedevtools.github.io/devtools-protocol)
(CDP) to control web sessions in Chrome and other Blink-based browsers.

This project improves, consolidates and extends the approaches of other
products such as [Selenium WebDriver](https://www.selenium.dev/documentation),
[Puppeteer](https://pptr.dev) and [ChromeDP](https://github.com/chromedp/chromedp).
The first two are very popular, but only the last one offers integration with
[Go](https://golang.org) without heavy non-Go dependencies.

## Goals

* Make the Chrome DevTools Protocol more accessible
* Simplify the low-level API even more than ChromeDP
* Support the [WebDriver API specification](https://www.w3.org/TR/webdriver)
  as a higher-level API layer
* Add Computer Vision (CV) and Optical Character Recognition (OCR) capabilities,\
  which are missing in all of the above

## Differences From ChromeDP

* Simpler session initialization, and simpler customization of browser flags
  (see [example](https://github.com/daabr/chrome-vision/blob/main/examples/init/main.go>))
* Simpler CDP command execution API:
  * More idiomatic command execution: non-blocking `Start` (returns a channel)
    and blocking `Do` (returns the command's result), instead of non-idiomaic
    and sometimes confusing JavaScript-await-like wrapper functions
    ([`chromedp.Run`](https://pkg.go.dev/github.com/chromedp/chromedp#Run),
    [`chromedp.RunResponse`](https://pkg.go.dev/github.com/chromedp/chromedp#RunResponse),
    [`chromedp.ActionFunc`](https://pkg.go.dev/github.com/chromedp/chromedp#ActionFunc))
  * No need to initialize executors, either with
    [`chromedp.Run`](https://pkg.go.dev/github.com/chromedp/chromedp#Run), or with
    [`cdproto.cdp.WithExecutor`](https://pkg.go.dev/github.com/chromedp/cdproto/cdp#WithExecutor)
* Communication with the browser:
  * Non-Windows OSs: using POSIX pipes instead of WebSockets (faster, more
    reliable and more secure)
  * Windows: using an internal implementation of the WebSocket protocol
    (optimized for Chrome DevTools, faster and more efficient - see
    [documentation](https://pkg.go.dev/github.com/daabr/chrome-vision/pkg/websocket))
* Stronger adherence to idiomatic Go coding style and
  <https://github.com/golang-standards/project-layout>
