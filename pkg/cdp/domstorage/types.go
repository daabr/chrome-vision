package domstorage

// DOM Storage identifier.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#type-StorageId
type StorageID struct {
	// Security origin for the storage.
	SecurityOrigin string
	// Whether the storage is local storage (not session storage).
	IsLocalStorage bool
}

// DOM Storage item.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#type-Item
type Item []string
