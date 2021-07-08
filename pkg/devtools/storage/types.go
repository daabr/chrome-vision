package storage

// StorageType data type. Enum of possible storage types.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#type-StorageType
type Type string

// Type valid values.
const (
	TypeAppcache       Type = "appcache"
	TypeCookies        Type = "cookies"
	TypeFileSystems    Type = "file_systems"
	TypeIndexeddb      Type = "indexeddb"
	TypeLocalStorage   Type = "local_storage"
	TypeShaderCache    Type = "shader_cache"
	TypeWebsql         Type = "websql"
	TypeServiceWorkers Type = "service_workers"
	TypeCacheStorage   Type = "cache_storage"
	TypeAll            Type = "all"
	TypeOther          Type = "other"
)

// String returns the Type value as a built-in string.
func (t Type) String() string {
	return string(t)
}

// UsageForType data type. Usage for a storage type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#type-UsageForType
type UsageForType struct {
	// Name of storage type.
	StorageType Type `json:"storageType"`
	// Storage usage (bytes).
	Usage float64 `json:"usage"`
}

// TrustTokens data type. Pair of issuer origin and number of available (signed, but not used) Trust
// Tokens from that issuer.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#type-TrustTokens
//
// This CDP type is experimental.
type TrustTokens struct {
	IssuerOrigin string  `json:"issuerOrigin"`
	Count        float64 `json:"count"`
}
