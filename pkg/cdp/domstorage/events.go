package domstorage

// https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#event-domStorageItemAdded
type DomStorageItemAdded struct {
	StorageID StorageID
	Key       string
	NewValue  string
}

// https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#event-domStorageItemRemoved
type DomStorageItemRemoved struct {
	StorageID StorageID
	Key       string
}

// https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#event-domStorageItemUpdated
type DomStorageItemUpdated struct {
	StorageID StorageID
	Key       string
	OldValue  string
	NewValue  string
}

// https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#event-domStorageItemsCleared
type DomStorageItemsCleared struct {
	StorageID StorageID
}
