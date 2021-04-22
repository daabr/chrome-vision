package applicationcache

// Detailed application cache resource information.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#type-ApplicationCacheResource
type ApplicationCacheResource struct {
	// Resource url.
	URL string
	// Resource size.
	Size int64
	// Resource type.
	Type string
}

// Detailed application cache information.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#type-ApplicationCache
type ApplicationCache struct {
	// Manifest URL.
	ManifestURL string
	// Application cache size.
	Size float64
	// Application cache creation time.
	CreationTime float64
	// Application cache update time.
	UpdateTime float64
	// Application cache resources.
	Resources []ApplicationCacheResource
}

// Frame identifier - manifest URL pair.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#type-FrameWithManifest
type FrameWithManifest struct {
	// Frame identifier.
	FrameID string
	// Manifest URL.
	ManifestURL string
	// Application cache status.
	Status int64
}
