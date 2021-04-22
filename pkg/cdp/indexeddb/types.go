package indexeddb

import "github.com/daabr/chrome-vision/pkg/cdp/runtime"

// Database with an array of object stores.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#type-DatabaseWithObjectStores
type DatabaseWithObjectStores struct {
	// Database name.
	Name string
	// Database version (type is not 'integer', as the standard
	// requires the version number to be 'unsigned long long')
	Version float64
	// Object stores in this database.
	ObjectStores []ObjectStore
}

// Object store.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#type-ObjectStore
type ObjectStore struct {
	// Object store name.
	Name string
	// Object store key path.
	KeyPath KeyPath
	// If true, object store has auto increment flag set.
	AutoIncrement bool
	// Indexes in this object store.
	Indexes []ObjectStoreIndex
}

// Object store index.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#type-ObjectStoreIndex
type ObjectStoreIndex struct {
	// Index name.
	Name string
	// Index key path.
	KeyPath KeyPath
	// If true, index is unique.
	Unique bool
	// If true, index allows multiple entries for a key.
	MultiEntry bool
}

// Key.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#type-Key
type Key struct {
	// Key type.
	Type string
	// Number value.
	Number float64 `json:"number,omitempty"`
	// String value.
	String string `json:"string,omitempty"`
	// Date value.
	Date float64 `json:"date,omitempty"`
	// Array value.
	Array []Key `json:"array,omitempty"`
}

// Key range.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#type-KeyRange
type KeyRange struct {
	// Lower bound.
	Lower *Key `json:"lower,omitempty"`
	// Upper bound.
	Upper *Key `json:"upper,omitempty"`
	// If true lower bound is open.
	LowerOpen bool
	// If true upper bound is open.
	UpperOpen bool
}

// Data entry.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#type-DataEntry
type DataEntry struct {
	// Key object.
	Key runtime.RemoteObject
	// Primary key object.
	PrimaryKey runtime.RemoteObject
	// Value object.
	Value runtime.RemoteObject
}

// Key path.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#type-KeyPath
type KeyPath struct {
	// Key path type.
	Type string
	// String value.
	String string `json:"string,omitempty"`
	// Array value.
	Array []string `json:"array,omitempty"`
}
