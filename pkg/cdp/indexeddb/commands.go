package indexeddb

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/cdp"
)

// ClearObjectStore contains the parameters, and acts as
// a Go receiver, for the CDP command `clearObjectStore`.
//
// Clears all entries from an object store.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-clearObjectStore
type ClearObjectStore struct {
	// Security origin.
	SecurityOrigin string `json:"securityOrigin"`
	// Database name.
	DatabaseName string `json:"databaseName"`
	// Object store name.
	ObjectStoreName string `json:"objectStoreName"`
}

// NewClearObjectStore constructs a new ClearObjectStore struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-clearObjectStore
func NewClearObjectStore(securityOrigin string, databaseName string, objectStoreName string) *ClearObjectStore {
	return &ClearObjectStore{
		SecurityOrigin:  securityOrigin,
		DatabaseName:    databaseName,
		ObjectStoreName: objectStoreName,
	}
}

// Do sends the ClearObjectStore CDP command to a browser,
// and returns the browser's response.
func (t *ClearObjectStore) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "ClearObjectStore", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// DeleteDatabase contains the parameters, and acts as
// a Go receiver, for the CDP command `deleteDatabase`.
//
// Deletes a database.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-deleteDatabase
type DeleteDatabase struct {
	// Security origin.
	SecurityOrigin string `json:"securityOrigin"`
	// Database name.
	DatabaseName string `json:"databaseName"`
}

// NewDeleteDatabase constructs a new DeleteDatabase struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-deleteDatabase
func NewDeleteDatabase(securityOrigin string, databaseName string) *DeleteDatabase {
	return &DeleteDatabase{
		SecurityOrigin: securityOrigin,
		DatabaseName:   databaseName,
	}
}

// Do sends the DeleteDatabase CDP command to a browser,
// and returns the browser's response.
func (t *DeleteDatabase) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "DeleteDatabase", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// DeleteObjectStoreEntries contains the parameters, and acts as
// a Go receiver, for the CDP command `deleteObjectStoreEntries`.
//
// Delete a range of entries from an object store
//
// https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-deleteObjectStoreEntries
type DeleteObjectStoreEntries struct {
	SecurityOrigin  string `json:"securityOrigin"`
	DatabaseName    string `json:"databaseName"`
	ObjectStoreName string `json:"objectStoreName"`
	// Range of entry keys to delete
	KeyRange KeyRange `json:"keyRange"`
}

// NewDeleteObjectStoreEntries constructs a new DeleteObjectStoreEntries struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-deleteObjectStoreEntries
func NewDeleteObjectStoreEntries(securityOrigin string, databaseName string, objectStoreName string, keyRange KeyRange) *DeleteObjectStoreEntries {
	return &DeleteObjectStoreEntries{
		SecurityOrigin:  securityOrigin,
		DatabaseName:    databaseName,
		ObjectStoreName: objectStoreName,
		KeyRange:        keyRange,
	}
}

// Do sends the DeleteObjectStoreEntries CDP command to a browser,
// and returns the browser's response.
func (t *DeleteObjectStoreEntries) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "DeleteObjectStoreEntries", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// Disable contains the parameters, and acts as
// a Go receiver, for the CDP command `disable`.
//
// Disables events from backend.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-disable
type Disable struct{}

// NewDisable constructs a new Disable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-disable
func NewDisable() *Disable {
	return &Disable{}
}

// Do sends the Disable CDP command to a browser,
// and returns the browser's response.
func (t *Disable) Do(ctx context.Context) error {
	response, err := cdp.Send(ctx, "Disable", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// Enable contains the parameters, and acts as
// a Go receiver, for the CDP command `enable`.
//
// Enables events from backend.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-enable
type Enable struct{}

// NewEnable constructs a new Enable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-enable
func NewEnable() *Enable {
	return &Enable{}
}

// Do sends the Enable CDP command to a browser,
// and returns the browser's response.
func (t *Enable) Do(ctx context.Context) error {
	response, err := cdp.Send(ctx, "Enable", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// RequestData contains the parameters, and acts as
// a Go receiver, for the CDP command `requestData`.
//
// Requests data from object store or index.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-requestData
type RequestData struct {
	// Security origin.
	SecurityOrigin string `json:"securityOrigin"`
	// Database name.
	DatabaseName string `json:"databaseName"`
	// Object store name.
	ObjectStoreName string `json:"objectStoreName"`
	// Index name, empty string for object store data requests.
	IndexName string `json:"indexName"`
	// Number of records to skip.
	SkipCount int64 `json:"skipCount"`
	// Number of records to fetch.
	PageSize int64 `json:"pageSize"`
	// Key range.
	KeyRange *KeyRange `json:"keyRange,omitempty"`
}

// NewRequestData constructs a new RequestData struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-requestData
func NewRequestData(securityOrigin string, databaseName string, objectStoreName string, indexName string, skipCount int64, pageSize int64) *RequestData {
	return &RequestData{
		SecurityOrigin:  securityOrigin,
		DatabaseName:    databaseName,
		ObjectStoreName: objectStoreName,
		IndexName:       indexName,
		SkipCount:       skipCount,
		PageSize:        pageSize,
	}
}

// SetKeyRange adds or modifies the value of the optional
// parameter `keyRange` in the RequestData CDP command.
//
// Key range.
func (t *RequestData) SetKeyRange(v KeyRange) *RequestData {
	t.KeyRange = &v
	return t
}

// RequestDataResponse contains the browser's response
// to calling the RequestData CDP command with Do().
type RequestDataResponse struct {
	// Array of object store data entries.
	ObjectStoreDataEntries []DataEntry `json:"objectStoreDataEntries"`
	// If true, there are more entries to fetch in the given range.
	HasMore bool `json:"hasMore"`
}

// Do sends the RequestData CDP command to a browser,
// and returns the browser's response.
func (t *RequestData) Do(ctx context.Context) (*RequestDataResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "RequestData", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &RequestDataResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetMetadata contains the parameters, and acts as
// a Go receiver, for the CDP command `getMetadata`.
//
// Gets metadata of an object store
//
// https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-getMetadata
type GetMetadata struct {
	// Security origin.
	SecurityOrigin string `json:"securityOrigin"`
	// Database name.
	DatabaseName string `json:"databaseName"`
	// Object store name.
	ObjectStoreName string `json:"objectStoreName"`
}

// NewGetMetadata constructs a new GetMetadata struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-getMetadata
func NewGetMetadata(securityOrigin string, databaseName string, objectStoreName string) *GetMetadata {
	return &GetMetadata{
		SecurityOrigin:  securityOrigin,
		DatabaseName:    databaseName,
		ObjectStoreName: objectStoreName,
	}
}

// GetMetadataResponse contains the browser's response
// to calling the GetMetadata CDP command with Do().
type GetMetadataResponse struct {
	// the entries count
	EntriesCount float64 `json:"entriesCount"`
	// the current value of key generator, to become the next inserted
	// key into the object store. Valid if objectStore.autoIncrement
	// is true.
	KeyGeneratorValue float64 `json:"keyGeneratorValue"`
}

// Do sends the GetMetadata CDP command to a browser,
// and returns the browser's response.
func (t *GetMetadata) Do(ctx context.Context) (*GetMetadataResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "GetMetadata", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetMetadataResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// RequestDatabase contains the parameters, and acts as
// a Go receiver, for the CDP command `requestDatabase`.
//
// Requests database with given name in given frame.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-requestDatabase
type RequestDatabase struct {
	// Security origin.
	SecurityOrigin string `json:"securityOrigin"`
	// Database name.
	DatabaseName string `json:"databaseName"`
}

// NewRequestDatabase constructs a new RequestDatabase struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-requestDatabase
func NewRequestDatabase(securityOrigin string, databaseName string) *RequestDatabase {
	return &RequestDatabase{
		SecurityOrigin: securityOrigin,
		DatabaseName:   databaseName,
	}
}

// RequestDatabaseResponse contains the browser's response
// to calling the RequestDatabase CDP command with Do().
type RequestDatabaseResponse struct {
	// Database with an array of object stores.
	DatabaseWithObjectStores DatabaseWithObjectStores `json:"databaseWithObjectStores"`
}

// Do sends the RequestDatabase CDP command to a browser,
// and returns the browser's response.
func (t *RequestDatabase) Do(ctx context.Context) (*RequestDatabaseResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "RequestDatabase", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &RequestDatabaseResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// RequestDatabaseNames contains the parameters, and acts as
// a Go receiver, for the CDP command `requestDatabaseNames`.
//
// Requests database names for given security origin.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-requestDatabaseNames
type RequestDatabaseNames struct {
	// Security origin.
	SecurityOrigin string `json:"securityOrigin"`
}

// NewRequestDatabaseNames constructs a new RequestDatabaseNames struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB/#method-requestDatabaseNames
func NewRequestDatabaseNames(securityOrigin string) *RequestDatabaseNames {
	return &RequestDatabaseNames{
		SecurityOrigin: securityOrigin,
	}
}

// RequestDatabaseNamesResponse contains the browser's response
// to calling the RequestDatabaseNames CDP command with Do().
type RequestDatabaseNamesResponse struct {
	// Database names for origin.
	DatabaseNames []string `json:"databaseNames"`
}

// Do sends the RequestDatabaseNames CDP command to a browser,
// and returns the browser's response.
func (t *RequestDatabaseNames) Do(ctx context.Context) (*RequestDatabaseNamesResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "RequestDatabaseNames", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &RequestDatabaseNamesResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}
