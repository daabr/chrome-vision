[![Go Reference](https://pkg.go.dev/badge/github.com/daabr/chrome-vision.svg)](https://pkg.go.dev/github.com/daabr/chrome-vision)
[![Go Report Card](https://goreportcard.com/badge/github.com/daabr/chrome-vision)](https://goreportcard.com/report/github.com/daabr/chrome-vision)
[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/daabr/chrome-vision/Go/main)](https://github.com/daabr/chrome-vision/actions/workflows/go.yml?query=branch:main)

# Chrome Vision

Advanced cross-platform web automation with a convenient Go API.

This is based on the
[Chrome DevTools Protocol](https://chromedevtools.github.io/devtools-protocol)
(CDP) to control web sessions in Chrome and other Blink-based browsers.

This project improves, consolidates and extends the approaches of other
products such as [Selenium WebDriver](https://www.selenium.dev/documentation),
[Puppeteer](https://pptr.dev) and [ChromeDP](https://github.com/chromedp/chromedp).
The first two are very popular, but only the last one offers integration with
[Go](https://golang.org) without heavy non-Go dependencies.

This is an experimental project, aiming to:

* Learn more about the Chrome DevTools Protocol
* Simplify the low-level API even more than ChromeDP
* Support the [WebDriver API specification](https://www.w3.org/TR/webdriver)
  as a higher-level API layer
* Add Computer Vision (CV) and Optical Character Recognition (OCR) capabilities,
  which are missing in all of the above
