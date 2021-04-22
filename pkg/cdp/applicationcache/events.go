package applicationcache

// https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#event-applicationCacheStatusUpdated
type ApplicationCacheStatusUpdated struct {
	// Identifier of the frame containing document whose application cache updated status.
	FrameID string
	// Manifest URL.
	ManifestURL string
	// Updated application cache status.
	Status int64
}

// https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#event-networkStateUpdated
type NetworkStateUpdated struct {
	IsNowOnline bool
}
