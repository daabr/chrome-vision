# Chrome Vision

Advanced web automation with a convenient Go API.

This is based on the [Chrome DevTools Protocol](https://chromedevtools.github.io/devtools-protocol)
(CDP) to control web sessions in Chrome and other Blink-based browsers.

This approach is similar to other products such as:

* [Selenium WebDriver](https://www.selenium.dev/documentation)
* [Puppeteer](https://pptr.dev)
* [ChromeDP](https://github.com/chromedp/chromedp)

The first two are more popular, but only the last one offers integration with
[Go](https://golang.org) without heavy non-Go dependencies.

This is an experimental project, aiming to minimize dependencies and simplify
the API even more than ChromeDP, but even more importantly: add support for the
[WebDriver API specification](https://www.w3.org/TR/webdriver), and integration
with Computer Vision (CV) and Optical Character Recognition (OCR) capabilities.
