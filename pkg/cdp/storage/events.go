package storage

// CacheStorageContentUpdated asynchronous event. A cache's contents have been modified.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#event-cacheStorageContentUpdated
type CacheStorageContentUpdated struct {
	// Origin to update.
	Origin string `json:"origin"`
	// Name of cache in origin.
	CacheName string `json:"cacheName"`
}

// CacheStorageListUpdated asynchronous event. A cache has been added/deleted.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#event-cacheStorageListUpdated
type CacheStorageListUpdated struct {
	// Origin to update.
	Origin string `json:"origin"`
}

// IndexedDBContentUpdated asynchronous event. The origin's IndexedDB object store has been modified.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#event-indexedDBContentUpdated
type IndexedDBContentUpdated struct {
	// Origin to update.
	Origin string `json:"origin"`
	// Database to update.
	DatabaseName string `json:"databaseName"`
	// ObjectStore to update.
	ObjectStoreName string `json:"objectStoreName"`
}

// IndexedDBListUpdated asynchronous event. The origin's IndexedDB database list has been modified.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#event-indexedDBListUpdated
type IndexedDBListUpdated struct {
	// Origin to update.
	Origin string `json:"origin"`
}
