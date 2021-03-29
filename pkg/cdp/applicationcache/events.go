package applicationcache

import "github.com/daabr/chrome-vision/pkg/cdp"

// https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#event-applicationCacheStatusUpdated
type ApplicationCacheStatusUpdated struct {
	// Identifier of the frame containing document whose application cache updated status.
	FrameID cdp.FrameID `json:"frameId"`
	// Manifest URL.
	ManifestURL string `json:"manifestURL"`
	// Updated application cache status.
	Status int64 `json:"status"`
}

// https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#event-networkStateUpdated
type NetworkStateUpdated struct {
	IsNowOnline bool `json:"isNowOnline"`
}
