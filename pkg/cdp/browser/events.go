package browser

// DownloadWillBegin asynchronous event. Fired when page is about to start a download.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#event-downloadWillBegin
//
// This CDP event is experimental.
type DownloadWillBegin struct {
	// Id of the frame that caused the download to begin.
	FrameID string `json:"frameId"`
	// Global unique identifier of the download.
	Guid string `json:"guid"`
	// URL of the resource being downloaded.
	URL string `json:"url"`
	// Suggested file name of the resource (the actual name of the file saved on disk may differ).
	SuggestedFilename string `json:"suggestedFilename"`
}

// DownloadProgress asynchronous event. Fired when download makes progress. Last call has |done| == true.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#event-downloadProgress
//
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
