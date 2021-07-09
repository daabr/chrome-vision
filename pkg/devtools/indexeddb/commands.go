package indexeddb

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
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
	m, err := devtools.SendAndWait(ctx, "IndexedDB.clearObjectStore", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the ClearObjectStore CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *ClearObjectStore) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "IndexedDB.clearObjectStore", b)
}

// ParseResponse parses the browser's response
// to the ClearObjectStore CDP command.
func (t *ClearObjectStore) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
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
	m, err := devtools.SendAndWait(ctx, "IndexedDB.deleteDatabase", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the DeleteDatabase CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *DeleteDatabase) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "IndexedDB.deleteDatabase", b)
}

// ParseResponse parses the browser's response
// to the DeleteDatabase CDP command.
func (t *DeleteDatabase) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
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
	m, err := devtools.SendAndWait(ctx, "IndexedDB.deleteObjectStoreEntries", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the DeleteObjectStoreEntries CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *DeleteObjectStoreEntries) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "IndexedDB.deleteObjectStoreEntries", b)
}

// ParseResponse parses the browser's response
// to the DeleteObjectStoreEntries CDP command.
func (t *DeleteObjectStoreEntries) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
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
	m, err := devtools.SendAndWait(ctx, "IndexedDB.disable", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the Disable CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *Disable) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "IndexedDB.disable", nil)
}

// ParseResponse parses the browser's response
// to the Disable CDP command.
func (t *Disable) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
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
	m, err := devtools.SendAndWait(ctx, "IndexedDB.enable", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the Enable CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *Enable) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "IndexedDB.enable", nil)
}

// ParseResponse parses the browser's response
// to the Enable CDP command.
func (t *Enable) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
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

// RequestDataResult contains the browser's response
// to calling the RequestData CDP command with Do().
type RequestDataResult struct {
	// Array of object store data entries.
	ObjectStoreDataEntries []DataEntry `json:"objectStoreDataEntries"`
	// If true, there are more entries to fetch in the given range.
	HasMore bool `json:"hasMore"`
}

// Do sends the RequestData CDP command to a browser,
// and returns the browser's response.
func (t *RequestData) Do(ctx context.Context) (*RequestDataResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "IndexedDB.requestData", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the RequestData CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *RequestData) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "IndexedDB.requestData", b)
}

// ParseResponse parses the browser's response
// to the RequestData CDP command.
func (t *RequestData) ParseResponse(m *devtools.Message) (*RequestDataResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &RequestDataResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
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

// GetMetadataResult contains the browser's response
// to calling the GetMetadata CDP command with Do().
type GetMetadataResult struct {
	// the entries count
	EntriesCount float64 `json:"entriesCount"`
	// the current value of key generator, to become the next inserted
	// key into the object store. Valid if objectStore.autoIncrement
	// is true.
	KeyGeneratorValue float64 `json:"keyGeneratorValue"`
}

// Do sends the GetMetadata CDP command to a browser,
// and returns the browser's response.
func (t *GetMetadata) Do(ctx context.Context) (*GetMetadataResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "IndexedDB.getMetadata", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetMetadata CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetMetadata) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "IndexedDB.getMetadata", b)
}

// ParseResponse parses the browser's response
// to the GetMetadata CDP command.
func (t *GetMetadata) ParseResponse(m *devtools.Message) (*GetMetadataResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetMetadataResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
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

// RequestDatabaseResult contains the browser's response
// to calling the RequestDatabase CDP command with Do().
type RequestDatabaseResult struct {
	// Database with an array of object stores.
	DatabaseWithObjectStores DatabaseWithObjectStores `json:"databaseWithObjectStores"`
}

// Do sends the RequestDatabase CDP command to a browser,
// and returns the browser's response.
func (t *RequestDatabase) Do(ctx context.Context) (*RequestDatabaseResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "IndexedDB.requestDatabase", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the RequestDatabase CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *RequestDatabase) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "IndexedDB.requestDatabase", b)
}

// ParseResponse parses the browser's response
// to the RequestDatabase CDP command.
func (t *RequestDatabase) ParseResponse(m *devtools.Message) (*RequestDatabaseResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &RequestDatabaseResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
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

// RequestDatabaseNamesResult contains the browser's response
// to calling the RequestDatabaseNames CDP command with Do().
type RequestDatabaseNamesResult struct {
	// Database names for origin.
	DatabaseNames []string `json:"databaseNames"`
}

// Do sends the RequestDatabaseNames CDP command to a browser,
// and returns the browser's response.
func (t *RequestDatabaseNames) Do(ctx context.Context) (*RequestDatabaseNamesResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "IndexedDB.requestDatabaseNames", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the RequestDatabaseNames CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *RequestDatabaseNames) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "IndexedDB.requestDatabaseNames", b)
}

// ParseResponse parses the browser's response
// to the RequestDatabaseNames CDP command.
func (t *RequestDatabaseNames) ParseResponse(m *devtools.Message) (*RequestDatabaseNamesResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &RequestDatabaseNamesResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}
