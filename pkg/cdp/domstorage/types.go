package domstorage

// DOM Storage identifier.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#type-StorageId
type StorageID struct {
	// Security origin for the storage.
	SecurityOrigin string `json:"securityOrigin"`
	// Whether the storage is local storage (not session storage).
	IsLocalStorage bool `json:"isLocalStorage"`
}

// DOM Storage item.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#type-Item
type Item []string
