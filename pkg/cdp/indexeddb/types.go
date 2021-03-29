package indexeddb

import (
	"github.com/daabr/chrome-vision/pkg/cdp/runtime"
)

// Database with an array of object stores.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#type-DatabaseWithObjectStores
type DatabaseWithObjectStores struct {
	// Database name.
	Name string `json:"name"`
	// Database version (type is not 'integer', as the standard
	// requires the version number to be 'unsigned long long')
	Version float64 `json:"version"`
	// Object stores in this database.
	ObjectStores []ObjectStore `json:"objectStores"`
}

// Object store.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#type-ObjectStore
type ObjectStore struct {
	// Object store name.
	Name string `json:"name"`
	// Object store key path.
	KeyPath KeyPath `json:"keyPath"`
	// If true, object store has auto increment flag set.
	AutoIncrement bool `json:"autoIncrement"`
	// Indexes in this object store.
	Indexes []ObjectStoreIndex `json:"indexes"`
}

// Object store index.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#type-ObjectStoreIndex
type ObjectStoreIndex struct {
	// Index name.
	Name string `json:"name"`
	// Index key path.
	KeyPath KeyPath `json:"keyPath"`
	// If true, index is unique.
	Unique bool `json:"unique"`
	// If true, index allows multiple entries for a key.
	MultiEntry bool `json:"multiEntry"`
}

// Key.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#type-Key
type Key struct {
	// Key type.
	Type string `json:"type"`
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
	LowerOpen bool `json:"lowerOpen"`
	// If true upper bound is open.
	UpperOpen bool `json:"upperOpen"`
}

// Data entry.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#type-DataEntry
type DataEntry struct {
	// Key object.
	Key runtime.RemoteObject `json:"key"`
	// Primary key object.
	PrimaryKey runtime.RemoteObject `json:"primaryKey"`
	// Value object.
	Value runtime.RemoteObject `json:"value"`
}

// Key path.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#type-KeyPath
type KeyPath struct {
	// Key path type.
	Type string `json:"type"`
	// String value.
	String string `json:"string,omitempty"`
	// Array value.
	Array []string `json:"array,omitempty"`
}
