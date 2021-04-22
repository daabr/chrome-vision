package browser

// Fired when page is about to start a download.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#event-downloadWillBegin
//
// This CDP event is experimental.
type DownloadWillBegin struct {
	// Id of the frame that caused the download to begin.
	FrameID string
	// Global unique identifier of the download.
	Guid string
	// URL of the resource being downloaded.
	URL string
	// Suggested file name of the resource (the actual name of the file saved on disk may differ).
	SuggestedFilename string
}

// Fired when download makes progress. Last call has |done| == true.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#event-downloadProgress
//
// This CDP event is experimental.
type DownloadProgress struct {
	// Global unique identifier of the download.
	Guid string
	// Total expected bytes to download.
	TotalBytes float64
	// Total bytes received.
	ReceivedBytes float64
	// Download status.
	State string
}
