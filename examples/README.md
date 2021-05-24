# Examples

Each directory focuses on a specific functionality or website, and demonstrates
multiple implementations of the same workflow, using different API layers:

* Chrome DevTools Protocol (CDP) - low level API
* WebDriver - high level API
* CV and OCR API extensions

## Initialization

Initialization and customization of browser sessions ([code](./init/main.go),
[godoc](https://pkg.go.dev/github.com/daabr/chrome-vision/examples/init)).

## Google Search

1. Navigate to Google Search
   * Web search homepage (`/webhp`)
   * English UI language (`hl=eng`)
   * No personalization (`pws=0`)
2. Type the search query "kittens", and press the Enter key
3. Click the "Images" tab to show image search results
4. Scroll down to the bottom of the page
5. Print the title and URL of the last result
