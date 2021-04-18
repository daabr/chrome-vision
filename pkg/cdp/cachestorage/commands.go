package cachestorage

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/cdp"
)

// DeleteCache contains the parameters, and acts as
// a Go receiver, for the CDP command `deleteCache`.
//
// Deletes a cache.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-deleteCache
type DeleteCache struct {
	// Id of cache for deletion.
	CacheID CacheID `json:"cacheId"`
}

// NewDeleteCache constructs a new DeleteCache struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-deleteCache
func NewDeleteCache(cacheId CacheID) *DeleteCache {
	return &DeleteCache{
		CacheID: cacheId,
	}
}

// Do sends the DeleteCache CDP command to a browser,
// and returns the browser's response.
func (t *DeleteCache) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "CacheStorage.deleteCache", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
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
	CacheID CacheID `json:"cacheId"`
	// URL spec of the request.
	Request string `json:"request"`
}

// NewDeleteEntry constructs a new DeleteEntry struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/CacheStorage/#method-deleteEntry
func NewDeleteEntry(cacheId CacheID, request string) *DeleteEntry {
	return &DeleteEntry{
		CacheID: cacheId,
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
	response, err := cdp.Send(ctx, "CacheStorage.deleteEntry", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
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

// RequestCacheNamesResponse contains the browser's response
// to calling the RequestCacheNames CDP command with Do().
type RequestCacheNamesResponse struct {
	// Caches for the security origin.
	Caches []Cache `json:"caches"`
}

// Do sends the RequestCacheNames CDP command to a browser,
// and returns the browser's response.
func (t *RequestCacheNames) Do(ctx context.Context) (*RequestCacheNamesResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "CacheStorage.requestCacheNames", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &RequestCacheNamesResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
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
	CacheID CacheID `json:"cacheId"`
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
func NewRequestCachedResponse(cacheId CacheID, requestURL string, requestHeaders []Header) *RequestCachedResponse {
	return &RequestCachedResponse{
		CacheID:        cacheId,
		RequestURL:     requestURL,
		RequestHeaders: requestHeaders,
	}
}

// RequestCachedResponseResponse contains the browser's response
// to calling the RequestCachedResponse CDP command with Do().
type RequestCachedResponseResponse struct {
	// Response read from the cache.
	Response CachedResponse `json:"response"`
}

// Do sends the RequestCachedResponse CDP command to a browser,
// and returns the browser's response.
func (t *RequestCachedResponse) Do(ctx context.Context) (*RequestCachedResponseResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "CacheStorage.requestCachedResponse", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &RequestCachedResponseResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
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
	CacheID CacheID `json:"cacheId"`
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
func NewRequestEntries(cacheId CacheID) *RequestEntries {
	return &RequestEntries{
		CacheID: cacheId,
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

// RequestEntriesResponse contains the browser's response
// to calling the RequestEntries CDP command with Do().
type RequestEntriesResponse struct {
	// Array of object store data entries.
	CacheDataEntries []DataEntry `json:"cacheDataEntries"`
	// Count of returned entries from this storage. If pathFilter is empty, it
	// is the count of all entries from this storage.
	ReturnCount float64 `json:"returnCount"`
}

// Do sends the RequestEntries CDP command to a browser,
// and returns the browser's response.
func (t *RequestEntries) Do(ctx context.Context) (*RequestEntriesResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "CacheStorage.requestEntries", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &RequestEntriesResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}
