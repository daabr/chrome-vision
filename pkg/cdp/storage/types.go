package storage

// Enum of possible storage types.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#type-StorageType
type StorageType string

// StorageType valid values.
const (
	StorageTypeAppcache       StorageType = "appcache"
	StorageTypeCookies        StorageType = "cookies"
	StorageTypeFileSystems    StorageType = "file_systems"
	StorageTypeIndexeddb      StorageType = "indexeddb"
	StorageTypeLocalStorage   StorageType = "local_storage"
	StorageTypeShaderCache    StorageType = "shader_cache"
	StorageTypeWebsql         StorageType = "websql"
	StorageTypeServiceWorkers StorageType = "service_workers"
	StorageTypeCacheStorage   StorageType = "cache_storage"
	StorageTypeAll            StorageType = "all"
	StorageTypeOther          StorageType = "other"
)

// Usage for a storage type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#type-UsageForType
type UsageForType struct {
	// Name of storage type.
	StorageType StorageType `json:"storageType"`
	// Storage usage (bytes).
	Usage float64 `json:"usage"`
}

// Pair of issuer origin and number of available (signed, but not used) Trust
// Tokens from that issuer.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#type-TrustTokens
//
// This CDP type is experimental.
type TrustTokens struct {
	IssuerOrigin string  `json:"issuerOrigin"`
	Count        float64 `json:"count"`
}
