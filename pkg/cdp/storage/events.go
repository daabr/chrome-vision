package storage

// A cache's contents have been modified.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#event-cacheStorageContentUpdated
type CacheStorageContentUpdated struct {
	// Origin to update.
	Origin string `json:"origin"`
	// Name of cache in origin.
	CacheName string `json:"cacheName"`
}

// A cache has been added/deleted.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#event-cacheStorageListUpdated
type CacheStorageListUpdated struct {
	// Origin to update.
	Origin string `json:"origin"`
}

// The origin's IndexedDB object store has been modified.
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

// The origin's IndexedDB database list has been modified.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#event-indexedDBListUpdated
type IndexedDBListUpdated struct {
	// Origin to update.
	Origin string `json:"origin"`
}
