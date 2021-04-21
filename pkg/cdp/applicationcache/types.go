package applicationcache

// Detailed application cache resource information.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#type-ApplicationCacheResource
type ApplicationCacheResource struct {
	// Resource url.
	URL string `json:"url"`
	// Resource size.
	Size int64 `json:"size"`
	// Resource type.
	Type string `json:"type"`
}

// Detailed application cache information.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#type-ApplicationCache
type ApplicationCache struct {
	// Manifest URL.
	ManifestURL string `json:"manifestURL"`
	// Application cache size.
	Size float64 `json:"size"`
	// Application cache creation time.
	CreationTime float64 `json:"creationTime"`
	// Application cache update time.
	UpdateTime float64 `json:"updateTime"`
	// Application cache resources.
	Resources []ApplicationCacheResource `json:"resources"`
}

// Frame identifier - manifest URL pair.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#type-FrameWithManifest
type FrameWithManifest struct {
	// Frame identifier.
	FrameID string `json:"frameId"`
	// Manifest URL.
	ManifestURL string `json:"manifestURL"`
	// Application cache status.
	Status int64 `json:"status"`
}
