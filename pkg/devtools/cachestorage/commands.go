package cachestorage

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
)

// DeleteCache contains the parameters, and acts as
// a Go receiver, for the CDP command `deleteCache`.
//
// Deletes a cache.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-deleteCache
type DeleteCache struct {
	// Id of cache for deletion.
	CacheID string `json:"cacheId"`
}

// NewDeleteCache constructs a new DeleteCache struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-deleteCache
func NewDeleteCache(cacheID string) *DeleteCache {
	return &DeleteCache{
		CacheID: cacheID,
	}
}

// Do sends the DeleteCache CDP command to a browser,
// and returns the browser's response.
func (t *DeleteCache) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "CacheStorage.deleteCache", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the DeleteCache CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *DeleteCache) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "CacheStorage.deleteCache", b)
}

// ParseResponse parses the browser's response
// to the DeleteCache CDP command.
func (t *DeleteCache) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// DeleteEntry contains the parameters, and acts as
// a Go receiver, for the CDP command `deleteEntry`.
//
// Deletes a cache entry.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-deleteEntry
type DeleteEntry struct {
	// Id of cache where the entry will be deleted.
	CacheID string `json:"cacheId"`
	// URL spec of the request.
	Request string `json:"request"`
}

// NewDeleteEntry constructs a new DeleteEntry struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-deleteEntry
func NewDeleteEntry(cacheID string, request string) *DeleteEntry {
	return &DeleteEntry{
		CacheID: cacheID,
		Request: request,
	}
}

// Do sends the DeleteEntry CDP command to a browser,
// and returns the browser's response.
func (t *DeleteEntry) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "CacheStorage.deleteEntry", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the DeleteEntry CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *DeleteEntry) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "CacheStorage.deleteEntry", b)
}

// ParseResponse parses the browser's response
// to the DeleteEntry CDP command.
func (t *DeleteEntry) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// RequestCacheNames contains the parameters, and acts as
// a Go receiver, for the CDP command `requestCacheNames`.
//
// Requests cache names.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-requestCacheNames
type RequestCacheNames struct {
	// Security origin.
	SecurityOrigin string `json:"securityOrigin"`
}

// NewRequestCacheNames constructs a new RequestCacheNames struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-requestCacheNames
func NewRequestCacheNames(securityOrigin string) *RequestCacheNames {
	return &RequestCacheNames{
		SecurityOrigin: securityOrigin,
	}
}

// RequestCacheNamesResult contains the browser's response
// to calling the RequestCacheNames CDP command with Do().
type RequestCacheNamesResult struct {
	// Caches for the security origin.
	Caches []Cache `json:"caches"`
}

// Do sends the RequestCacheNames CDP command to a browser,
// and returns the browser's response.
func (t *RequestCacheNames) Do(ctx context.Context) (*RequestCacheNamesResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "CacheStorage.requestCacheNames", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the RequestCacheNames CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *RequestCacheNames) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "CacheStorage.requestCacheNames", b)
}

// ParseResponse parses the browser's response
// to the RequestCacheNames CDP command.
func (t *RequestCacheNames) ParseResponse(m *devtools.Message) (*RequestCacheNamesResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &RequestCacheNamesResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// RequestCachedResponse contains the parameters, and acts as
// a Go receiver, for the CDP command `requestCachedResponse`.
//
// Fetches cache entry.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-requestCachedResponse
type RequestCachedResponse struct {
	// Id of cache that contains the entry.
	CacheID string `json:"cacheId"`
	// URL spec of the request.
	RequestURL string `json:"requestURL"`
	// headers of the request.
	RequestHeaders []Header `json:"requestHeaders"`
}

// NewRequestCachedResponse constructs a new RequestCachedResponse struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-requestCachedResponse
func NewRequestCachedResponse(cacheID string, requestURL string, requestHeaders []Header) *RequestCachedResponse {
	return &RequestCachedResponse{
		CacheID:        cacheID,
		RequestURL:     requestURL,
		RequestHeaders: requestHeaders,
	}
}

// RequestCachedResponseResult contains the browser's response
// to calling the RequestCachedResponse CDP command with Do().
type RequestCachedResponseResult struct {
	// Response read from the cache.
	Response CachedResponse `json:"response"`
}

// Do sends the RequestCachedResponse CDP command to a browser,
// and returns the browser's response.
func (t *RequestCachedResponse) Do(ctx context.Context) (*RequestCachedResponseResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "CacheStorage.requestCachedResponse", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the RequestCachedResponse CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *RequestCachedResponse) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "CacheStorage.requestCachedResponse", b)
}

// ParseResponse parses the browser's response
// to the RequestCachedResponse CDP command.
func (t *RequestCachedResponse) ParseResponse(m *devtools.Message) (*RequestCachedResponseResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &RequestCachedResponseResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// RequestEntries contains the parameters, and acts as
// a Go receiver, for the CDP command `requestEntries`.
//
// Requests data from cache.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-requestEntries
type RequestEntries struct {
	// ID of cache to get entries from.
	CacheID string `json:"cacheId"`
	// Number of records to skip.
	SkipCount int64 `json:"skipCount,omitempty"`
	// Number of records to fetch.
	PageSize int64 `json:"pageSize,omitempty"`
	// If present, only return the entries containing this substring in the path
	PathFilter string `json:"pathFilter,omitempty"`
}

// NewRequestEntries constructs a new RequestEntries struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-requestEntries
func NewRequestEntries(cacheID string) *RequestEntries {
	return &RequestEntries{
		CacheID: cacheID,
	}
}

// SetSkipCount adds or modifies the value of the optional
// parameter `skipCount` in the RequestEntries CDP command.
//
// Number of records to skip.
func (t *RequestEntries) SetSkipCount(v int64) *RequestEntries {
	t.SkipCount = v
	return t
}

// SetPageSize adds or modifies the value of the optional
// parameter `pageSize` in the RequestEntries CDP command.
//
// Number of records to fetch.
func (t *RequestEntries) SetPageSize(v int64) *RequestEntries {
	t.PageSize = v
	return t
}

// SetPathFilter adds or modifies the value of the optional
// parameter `pathFilter` in the RequestEntries CDP command.
//
// If present, only return the entries containing this substring in the path
func (t *RequestEntries) SetPathFilter(v string) *RequestEntries {
	t.PathFilter = v
	return t
}

// RequestEntriesResult contains the browser's response
// to calling the RequestEntries CDP command with Do().
type RequestEntriesResult struct {
	// Array of object store data entries.
	CacheDataEntries []DataEntry `json:"cacheDataEntries"`
	// Count of returned entries from this storage. If pathFilter is empty, it
	// is the count of all entries from this storage.
	ReturnCount float64 `json:"returnCount"`
}

// Do sends the RequestEntries CDP command to a browser,
// and returns the browser's response.
func (t *RequestEntries) Do(ctx context.Context) (*RequestEntriesResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "CacheStorage.requestEntries", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the RequestEntries CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *RequestEntries) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "CacheStorage.requestEntries", b)
}

// ParseResponse parses the browser's response
// to the RequestEntries CDP command.
func (t *RequestEntries) ParseResponse(m *devtools.Message) (*RequestEntriesResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &RequestEntriesResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}
