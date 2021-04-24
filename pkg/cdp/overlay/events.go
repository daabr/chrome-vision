package overlay

import "github.com/daabr/chrome-vision/pkg/cdp/page"

// InspectNodeRequested asynchronous event.
//
// Fired when the node should be inspected. This happens after call to `setInspectMode` or when
// user manually inspects an element.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#event-inspectNodeRequested
type InspectNodeRequested struct {
	// Id of the node to inspect.
	BackendNodeID int64 `json:"backendNodeId"`
}

// NodeHighlightRequested asynchronous event.
//
// Fired when the node should be highlighted. This happens after call to `setInspectMode`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#event-nodeHighlightRequested
type NodeHighlightRequested struct {
	NodeID int64 `json:"nodeId"`
}

// ScreenshotRequested asynchronous event.
//
// Fired when user asks to capture screenshot of some area on the page.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#event-screenshotRequested
type ScreenshotRequested struct {
	// Viewport to capture, in device independent pixels (dip).
	Viewport page.Viewport `json:"viewport"`
}

// InspectModeCanceled asynchronous event.
//
// Fired when user cancels the inspect mode.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#event-inspectModeCanceled
type InspectModeCanceled struct{}
