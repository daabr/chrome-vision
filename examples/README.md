# Examples

Each directory focuses on a specific functionality or website, and demonstrates
multiple implementations of the same workflow, using different API layers:

* Chrome DevTools Protocol (CDP) - low level API
* WebDriver - high level API
* CV and OCR API extensions

## Initialization

Various initializations and customizations of browser sessions:

* [CDP](./init/main.go)

## Navigation

Navigation is a fundamental functionality in web automation, but these samples
also demonstrate the different modes of command execution in Chrome Vision:

* CDP
  * [Synchronous (blocking) execution](./navigation/devtools/sync/main.go)
    * Simple, but not always reliable or efficient timing-wise
    * Case in point:
      [page.Navigate](https://pkg.go.dev/github.com/daabr/chrome-vision/pkg/devtools/page#Navigate)
      may return before the URL is fully loaded
    * This sample also uses a navigation-specific event as an extra safety
      measure
  * [Asynchronous (non-blocking) execution](./navigation/devtools/async/main.go)
    * Pro: enables advanced logic to synchronize multiple results and events
    * Con: requires the caller to use Go channels directly

## Google Search

An "end-to-end" example of a full, non-trivial workflow:

1. Navigate to Google Search
   * Web search homepage (`/webhp`)
   * English UI language (`hl=eng`)
   * No personalization (`pws=0`)
2. Type the search query "kittens", and press the Enter key
3. Click the "Images" tab to show image search results
4. Scroll down to the bottom of the page
5. Print the title and URL of the last result

Implementations:

* [CDP](./googlesearch/devtools/main.go)
