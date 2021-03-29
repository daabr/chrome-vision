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

This is an experimental project, aiming to:

* Learn more about the Chrome DevTools Protocol
* Simplify the low-level API even more than ChromeDP?
* Support the [WebDriver API specification](https://www.w3.org/TR/webdriver)
  as a higher-level API layer
* Add Computer Vision (CV) and Optical Character Recognition (OCR) capabilities
