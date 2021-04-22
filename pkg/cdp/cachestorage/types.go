package cachestorage

// Unique identifier of the Cache object.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#type-CacheId
type CacheID string

// type of HTTP response cached
//
// https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#type-CachedResponseType
type CachedResponseType string

// CachedResponseType valid values.
const (
	CachedResponseTypeBasic          CachedResponseType = "basic"
	CachedResponseTypeCors           CachedResponseType = "cors"
	CachedResponseTypeDefault        CachedResponseType = "default"
	CachedResponseTypeError          CachedResponseType = "error"
	CachedResponseTypeOpaqueResponse CachedResponseType = "opaqueResponse"
	CachedResponseTypeOpaqueRedirect CachedResponseType = "opaqueRedirect"
)

// Data entry.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#type-DataEntry
type DataEntry struct {
	// Request URL.
	RequestURL string
	// Request method.
	RequestMethod string
	// Request headers
	RequestHeaders []Header
	// Number of seconds since epoch.
	ResponseTime float64
	// HTTP response status code.
	ResponseStatus int64
	// HTTP response status text.
	ResponseStatusText string
	// HTTP response type
	ResponseType string
	// Response headers
	ResponseHeaders []Header
}

// Cache identifier.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#type-Cache
type Cache struct {
	// An opaque unique id of the cache.
	CacheID string
	// Security origin of the cache.
	SecurityOrigin string
	// The name of the cache.
	CacheName string
}

// https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#type-Header
type Header struct {
	Name  string
	Value string
}

// Cached response
//
// https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#type-CachedResponse
type CachedResponse struct {
	// Entry content, base64-encoded. (Encoded as a base64 string when passed over JSON)
	Body string
}
