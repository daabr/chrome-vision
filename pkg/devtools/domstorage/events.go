package domstorage

// ItemAdded asynchronous event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#event-domStorageItemAdded
type ItemAdded struct {
	StorageID StorageID `json:"storageId"`
	Key       string    `json:"key"`
	NewValue  string    `json:"newValue"`
}

// ItemRemoved asynchronous event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#event-domStorageItemRemoved
type ItemRemoved struct {
	StorageID StorageID `json:"storageId"`
	Key       string    `json:"key"`
}

// ItemUpdated asynchronous event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#event-domStorageItemUpdated
type ItemUpdated struct {
	StorageID StorageID `json:"storageId"`
	Key       string    `json:"key"`
	OldValue  string    `json:"oldValue"`
	NewValue  string    `json:"newValue"`
}

// ItemsCleared asynchronous event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#event-domStorageItemsCleared
type ItemsCleared struct {
	StorageID StorageID `json:"storageId"`
}
