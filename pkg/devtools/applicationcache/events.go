package applicationcache

// StatusUpdated asynchronous event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#event-applicationCacheStatusUpdated
type StatusUpdated struct {
	// Identifier of the frame containing document whose application cache updated status.
	FrameID string `json:"frameId"`
	// Manifest URL.
	ManifestURL string `json:"manifestURL"`
	// Updated application cache status.
	Status int64 `json:"status"`
}

// NetworkStateUpdated asynchronous event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#event-networkStateUpdated
type NetworkStateUpdated struct {
	IsNowOnline bool `json:"isNowOnline"`
}
