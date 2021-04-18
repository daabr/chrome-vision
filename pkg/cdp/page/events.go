package page

import (
	"github.com/daabr/chrome-vision/pkg/cdp/dom"
	"github.com/daabr/chrome-vision/pkg/cdp/network"
	"github.com/daabr/chrome-vision/pkg/cdp/runtime"
)

// https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-domContentEventFired
type DomContentEventFired struct {
	Timestamp network.MonotonicTime `json:"timestamp"`
}

// Emitted only when `page.interceptFileChooser` is enabled.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-fileChooserOpened
type FileChooserOpened struct {
	// Id of the frame containing input node.
	//
	// This CDP parameter is experimental.
	FrameID FrameID `json:"frameId"`
	// Input node id.
	//
	// This CDP parameter is experimental.
	BackendNodeID dom.BackendNodeID `json:"backendNodeId"`
	// Input mode.
	Mode string `json:"mode"`
}

// Fired when frame has been attached to its parent.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameAttached
type FrameAttached struct {
	// Id of the frame that has been attached.
	FrameID FrameID `json:"frameId"`
	// Parent frame identifier.
	ParentFrameID FrameID `json:"parentFrameId"`
	// JavaScript stack trace of when frame was attached, only set if frame initiated from script.
	Stack *runtime.StackTrace `json:"stack,omitempty"`
}

// Fired when frame no longer has a scheduled navigation.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameClearedScheduledNavigation
//
// This CDP event is deprecated.
type FrameClearedScheduledNavigation struct {
	// Id of the frame that has cleared its scheduled navigation.
	FrameID FrameID `json:"frameId"`
}

// Fired when frame has been detached from its parent.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameDetached
type FrameDetached struct {
	// Id of the frame that has been detached.
	FrameID FrameID `json:"frameId"`
	// This CDP parameter is experimental.
	Reason string `json:"reason"`
}

// Fired once navigation of the frame has completed. Frame is now associated with the new loader.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameNavigated
type FrameNavigated struct {
	// Frame object.
	Frame Frame `json:"frame"`
	// This CDP parameter is experimental.
	Type NavigationType `json:"type"`
}

// Fired when opening document to write to.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-documentOpened
//
// This CDP event is experimental.
type DocumentOpened struct {
	// Frame object.
	Frame Frame `json:"frame"`
}

// https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameResized
//
// This CDP event is experimental.
type FrameResized struct{}

// Fired when a renderer-initiated navigation is requested.
// Navigation may still be cancelled after the event is issued.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameRequestedNavigation
//
// This CDP event is experimental.
type FrameRequestedNavigation struct {
	// Id of the frame that is being navigated.
	FrameID FrameID `json:"frameId"`
	// The reason for the navigation.
	Reason ClientNavigationReason `json:"reason"`
	// The destination URL for the requested navigation.
	URL string `json:"url"`
	// The disposition for the navigation.
	Disposition ClientNavigationDisposition `json:"disposition"`
}

// Fired when frame schedules a potential navigation.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameScheduledNavigation
//
// This CDP event is deprecated.
type FrameScheduledNavigation struct {
	// Id of the frame that has scheduled a navigation.
	FrameID FrameID `json:"frameId"`
	// Delay (in seconds) until the navigation is scheduled to begin. The navigation is not
	// guaranteed to start.
	Delay float64 `json:"delay"`
	// The reason for the navigation.
	Reason ClientNavigationReason `json:"reason"`
	// The destination URL for the scheduled navigation.
	URL string `json:"url"`
}

// Fired when frame has started loading.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameStartedLoading
//
// This CDP event is experimental.
type FrameStartedLoading struct {
	// Id of the frame that has started loading.
	FrameID FrameID `json:"frameId"`
}

// Fired when frame has stopped loading.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-frameStoppedLoading
//
// This CDP event is experimental.
type FrameStoppedLoading struct {
	// Id of the frame that has stopped loading.
	FrameID FrameID `json:"frameId"`
}

// Fired when page is about to start a download.
// Deprecated. Use Browser.downloadWillBegin instead.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-downloadWillBegin
//
// This CDP event is deprecated.
// This CDP event is experimental.
type DownloadWillBegin struct {
	// Id of the frame that caused download to begin.
	FrameID FrameID `json:"frameId"`
	// Global unique identifier of the download.
	Guid string `json:"guid"`
	// URL of the resource being downloaded.
	URL string `json:"url"`
	// Suggested file name of the resource (the actual name of the file saved on disk may differ).
	SuggestedFilename string `json:"suggestedFilename"`
}

// Fired when download makes progress. Last call has |done| == true.
// Deprecated. Use Browser.downloadProgress instead.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-downloadProgress
//
// This CDP event is deprecated.
// This CDP event is experimental.
type DownloadProgress struct {
	// Global unique identifier of the download.
	Guid string `json:"guid"`
	// Total expected bytes to download.
	TotalBytes float64 `json:"totalBytes"`
	// Total bytes received.
	ReceivedBytes float64 `json:"receivedBytes"`
	// Download status.
	State string `json:"state"`
}

// Fired when interstitial page was hidden
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-interstitialHidden
type InterstitialHidden struct{}

// Fired when interstitial page was shown
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-interstitialShown
type InterstitialShown struct{}

// Fired when a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload) has been
// closed.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-javascriptDialogClosed
type JavascriptDialogClosed struct {
	// Whether dialog was confirmed.
	Result bool `json:"result"`
	// User input in case of prompt.
	UserInput string `json:"userInput"`
}

// Fired when a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload) is about to
// open.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-javascriptDialogOpening
type JavascriptDialogOpening struct {
	// Frame url.
	URL string `json:"url"`
	// Message that will be displayed by the dialog.
	Message string `json:"message"`
	// Dialog type.
	Type DialogType `json:"type"`
	// True iff browser is capable showing or acting on the given dialog. When browser has no
	// dialog handler for given target, calling alert while Page domain is engaged will stall
	// the page execution. Execution can be resumed via calling Page.handleJavaScriptDialog.
	HasBrowserHandler bool `json:"hasBrowserHandler"`
	// Default dialog prompt.
	DefaultPrompt string `json:"defaultPrompt,omitempty"`
}

// Fired for top level page lifecycle events such as navigation, load, paint, etc.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-lifecycleEvent
type LifecycleEvent struct {
	// Id of the frame.
	FrameID FrameID `json:"frameId"`
	// Loader identifier. Empty string if the request is fetched from worker.
	LoaderID  network.LoaderID      `json:"loaderId"`
	Name      string                `json:"name"`
	Timestamp network.MonotonicTime `json:"timestamp"`
}

// Fired for failed bfcache history navigations if BackForwardCache feature is enabled. Do
// not assume any ordering with the Page.frameNavigated event. This event is fired only for
// main-frame history navigation where the document changes (non-same-document navigations),
// when bfcache navigation fails.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-backForwardCacheNotUsed
//
// This CDP event is experimental.
type BackForwardCacheNotUsed struct {
	// The loader id for the associated navgation.
	LoaderID network.LoaderID `json:"loaderId"`
	// The frame id of the associated frame.
	FrameID FrameID `json:"frameId"`
}

// https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-loadEventFired
type LoadEventFired struct {
	Timestamp network.MonotonicTime `json:"timestamp"`
}

// Fired when same-document navigation happens, e.g. due to history API usage or anchor navigation.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-navigatedWithinDocument
//
// This CDP event is experimental.
type NavigatedWithinDocument struct {
	// Id of the frame.
	FrameID FrameID `json:"frameId"`
	// Frame's new url.
	URL string `json:"url"`
}

// Compressed image data requested by the `startScreencast`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-screencastFrame
//
// This CDP event is experimental.
type ScreencastFrame struct {
	// Base64-encoded compressed image. (Encoded as a base64 string when passed over JSON)
	Data string `json:"data"`
	// Screencast frame metadata.
	Metadata ScreencastFrameMetadata `json:"metadata"`
	// Frame number.
	SessionID int64 `json:"sessionId"`
}

// Fired when the page with currently enabled screencast was shown or hidden `.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-screencastVisibilityChanged
//
// This CDP event is experimental.
type ScreencastVisibilityChanged struct {
	// True if the page is visible.
	Visible bool `json:"visible"`
}

// Fired when a new window is going to be opened, via window.open(), link click, form submission,
// etc.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-windowOpen
type WindowOpen struct {
	// The URL for the new window.
	URL string `json:"url"`
	// Window name.
	WindowName string `json:"windowName"`
	// An array of enabled window features.
	WindowFeatures []string `json:"windowFeatures"`
	// Whether or not it was triggered by user gesture.
	UserGesture bool `json:"userGesture"`
}

// Issued for every compilation cache generated. Is only available
// if Page.setGenerateCompilationCache is enabled.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#event-compilationCacheProduced
//
// This CDP event is experimental.
type CompilationCacheProduced struct {
	URL string `json:"url"`
	// Base64-encoded data (Encoded as a base64 string when passed over JSON)
	Data string `json:"data"`
}
