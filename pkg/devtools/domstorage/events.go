package domstorage

// DomStorageItemAdded asynchronous event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#event-domStorageItemAdded
type DomStorageItemAdded struct {
	StorageID StorageID `json:"storageId"`
	Key       string    `json:"key"`
	NewValue  string    `json:"newValue"`
}

// DomStorageItemRemoved asynchronous event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#event-domStorageItemRemoved
type DomStorageItemRemoved struct {
	StorageID StorageID `json:"storageId"`
	Key       string    `json:"key"`
}

// DomStorageItemUpdated asynchronous event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#event-domStorageItemUpdated
type DomStorageItemUpdated struct {
	StorageID StorageID `json:"storageId"`
	Key       string    `json:"key"`
	OldValue  string    `json:"oldValue"`
	NewValue  string    `json:"newValue"`
}

// DomStorageItemsCleared asynchronous event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#event-domStorageItemsCleared
type DomStorageItemsCleared struct {
	StorageID StorageID `json:"storageId"`
}
