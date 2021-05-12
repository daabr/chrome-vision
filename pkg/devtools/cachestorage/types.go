package cachestorage

// CacheID data type. Unique identifier of the Cache object.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#type-CacheId
type CacheID string

// CachedResponseType data type. type of HTTP response cached
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

// String returns the CachedResponseType value as a built-in string.
func (t CachedResponseType) String() string {
	return string(t)
}

// DataEntry data type. Data entry.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#type-DataEntry
type DataEntry struct {
	// Request URL.
	RequestURL string `json:"requestURL"`
	// Request method.
	RequestMethod string `json:"requestMethod"`
	// Request headers
	RequestHeaders []Header `json:"requestHeaders"`
	// Number of seconds since epoch.
	ResponseTime float64 `json:"responseTime"`
	// HTTP response status code.
	ResponseStatus int64 `json:"responseStatus"`
	// HTTP response status text.
	ResponseStatusText string `json:"responseStatusText"`
	// HTTP response type
	ResponseType CachedResponseType `json:"responseType"`
	// Response headers
	ResponseHeaders []Header `json:"responseHeaders"`
}

// Cache identifier.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#type-Cache
type Cache struct {
	// An opaque unique id of the cache.
	CacheID string `json:"cacheId"`
	// Security origin of the cache.
	SecurityOrigin string `json:"securityOrigin"`
	// The name of the cache.
	CacheName string `json:"cacheName"`
}

// Header data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#type-Header
type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// CachedResponse data type. Cached response
//
// https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#type-CachedResponse
type CachedResponse struct {
	// Entry content, base64-encoded. (Encoded as a base64 string when passed over JSON)
	Body string `json:"body"`
}
