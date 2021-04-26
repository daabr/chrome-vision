package storage

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/cdp"
	"github.com/daabr/chrome-vision/pkg/cdp/network"
)

// ClearDataForOrigin contains the parameters, and acts as
// a Go receiver, for the CDP command `clearDataForOrigin`.
//
// Clears storage for origin.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-clearDataForOrigin
type ClearDataForOrigin struct {
	// Security origin.
	Origin string `json:"origin"`
	// Comma separated list of StorageType to clear.
	StorageTypes string `json:"storageTypes"`
}

// NewClearDataForOrigin constructs a new ClearDataForOrigin struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-clearDataForOrigin
func NewClearDataForOrigin(origin string, storageTypes string) *ClearDataForOrigin {
	return &ClearDataForOrigin{
		Origin:       origin,
		StorageTypes: storageTypes,
	}
}

// Do sends the ClearDataForOrigin CDP command to a browser,
// and returns the browser's response.
func (t *ClearDataForOrigin) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Storage.clearDataForOrigin", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// GetCookies contains the parameters, and acts as
// a Go receiver, for the CDP command `getCookies`.
//
// Returns all browser cookies.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-getCookies
type GetCookies struct {
	// Browser context to use when called on the browser endpoint.
	BrowserContextID string `json:"browserContextId,omitempty"`
}

// NewGetCookies constructs a new GetCookies struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-getCookies
func NewGetCookies() *GetCookies {
	return &GetCookies{}
}

// SetBrowserContextID adds or modifies the value of the optional
// parameter `browserContextId` in the GetCookies CDP command.
//
// Browser context to use when called on the browser endpoint.
func (t *GetCookies) SetBrowserContextID(v string) *GetCookies {
	t.BrowserContextID = v
	return t
}

// GetCookiesResult contains the browser's response
// to calling the GetCookies CDP command with Do().
type GetCookiesResult struct {
	// Array of cookie objects.
	Cookies []network.Cookie `json:"cookies"`
}

// Do sends the GetCookies CDP command to a browser,
// and returns the browser's response.
func (t *GetCookies) Do(ctx context.Context) (*GetCookiesResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "Storage.getCookies", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetCookiesResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetCookies contains the parameters, and acts as
// a Go receiver, for the CDP command `setCookies`.
//
// Sets given cookies.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-setCookies
type SetCookies struct {
	// Cookies to be set.
	Cookies []network.CookieParam `json:"cookies"`
	// Browser context to use when called on the browser endpoint.
	BrowserContextID string `json:"browserContextId,omitempty"`
}

// NewSetCookies constructs a new SetCookies struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-setCookies
func NewSetCookies(cookies []network.CookieParam) *SetCookies {
	return &SetCookies{
		Cookies: cookies,
	}
}

// SetBrowserContextID adds or modifies the value of the optional
// parameter `browserContextId` in the SetCookies CDP command.
//
// Browser context to use when called on the browser endpoint.
func (t *SetCookies) SetBrowserContextID(v string) *SetCookies {
	t.BrowserContextID = v
	return t
}

// Do sends the SetCookies CDP command to a browser,
// and returns the browser's response.
func (t *SetCookies) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Storage.setCookies", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// ClearCookies contains the parameters, and acts as
// a Go receiver, for the CDP command `clearCookies`.
//
// Clears cookies.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-clearCookies
type ClearCookies struct {
	// Browser context to use when called on the browser endpoint.
	BrowserContextID string `json:"browserContextId,omitempty"`
}

// NewClearCookies constructs a new ClearCookies struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-clearCookies
func NewClearCookies() *ClearCookies {
	return &ClearCookies{}
}

// SetBrowserContextID adds or modifies the value of the optional
// parameter `browserContextId` in the ClearCookies CDP command.
//
// Browser context to use when called on the browser endpoint.
func (t *ClearCookies) SetBrowserContextID(v string) *ClearCookies {
	t.BrowserContextID = v
	return t
}

// Do sends the ClearCookies CDP command to a browser,
// and returns the browser's response.
func (t *ClearCookies) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Storage.clearCookies", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// GetUsageAndQuota contains the parameters, and acts as
// a Go receiver, for the CDP command `getUsageAndQuota`.
//
// Returns usage and quota in bytes.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-getUsageAndQuota
type GetUsageAndQuota struct {
	// Security origin.
	Origin string `json:"origin"`
}

// NewGetUsageAndQuota constructs a new GetUsageAndQuota struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-getUsageAndQuota
func NewGetUsageAndQuota(origin string) *GetUsageAndQuota {
	return &GetUsageAndQuota{
		Origin: origin,
	}
}

// GetUsageAndQuotaResult contains the browser's response
// to calling the GetUsageAndQuota CDP command with Do().
type GetUsageAndQuotaResult struct {
	// Storage usage (bytes).
	Usage float64 `json:"usage"`
	// Storage quota (bytes).
	Quota float64 `json:"quota"`
	// Whether or not the origin has an active storage quota override
	OverrideActive bool `json:"overrideActive"`
	// Storage usage per type (bytes).
	UsageBreakdown []UsageForType `json:"usageBreakdown"`
}

// Do sends the GetUsageAndQuota CDP command to a browser,
// and returns the browser's response.
func (t *GetUsageAndQuota) Do(ctx context.Context) (*GetUsageAndQuotaResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "Storage.getUsageAndQuota", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetUsageAndQuotaResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// OverrideQuotaForOrigin contains the parameters, and acts as
// a Go receiver, for the CDP command `overrideQuotaForOrigin`.
//
// Override quota for the specified origin
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-overrideQuotaForOrigin
//
// This CDP method is experimental.
type OverrideQuotaForOrigin struct {
	// Security origin.
	Origin string `json:"origin"`
	// The quota size (in bytes) to override the original quota with.
	// If this is called multiple times, the overridden quota will be equal to
	// the quotaSize provided in the final call. If this is called without
	// specifying a quotaSize, the quota will be reset to the default value for
	// the specified origin. If this is called multiple times with different
	// origins, the override will be maintained for each origin until it is
	// disabled (called without a quotaSize).
	QuotaSize float64 `json:"quotaSize,omitempty"`
}

// NewOverrideQuotaForOrigin constructs a new OverrideQuotaForOrigin struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-overrideQuotaForOrigin
//
// This CDP method is experimental.
func NewOverrideQuotaForOrigin(origin string) *OverrideQuotaForOrigin {
	return &OverrideQuotaForOrigin{
		Origin: origin,
	}
}

// SetQuotaSize adds or modifies the value of the optional
// parameter `quotaSize` in the OverrideQuotaForOrigin CDP command.
//
// The quota size (in bytes) to override the original quota with.
// If this is called multiple times, the overridden quota will be equal to
// the quotaSize provided in the final call. If this is called without
// specifying a quotaSize, the quota will be reset to the default value for
// the specified origin. If this is called multiple times with different
// origins, the override will be maintained for each origin until it is
// disabled (called without a quotaSize).
func (t *OverrideQuotaForOrigin) SetQuotaSize(v float64) *OverrideQuotaForOrigin {
	t.QuotaSize = v
	return t
}

// Do sends the OverrideQuotaForOrigin CDP command to a browser,
// and returns the browser's response.
func (t *OverrideQuotaForOrigin) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Storage.overrideQuotaForOrigin", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// TrackCacheStorageForOrigin contains the parameters, and acts as
// a Go receiver, for the CDP command `trackCacheStorageForOrigin`.
//
// Registers origin to be notified when an update occurs to its cache storage list.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-trackCacheStorageForOrigin
type TrackCacheStorageForOrigin struct {
	// Security origin.
	Origin string `json:"origin"`
}

// NewTrackCacheStorageForOrigin constructs a new TrackCacheStorageForOrigin struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-trackCacheStorageForOrigin
func NewTrackCacheStorageForOrigin(origin string) *TrackCacheStorageForOrigin {
	return &TrackCacheStorageForOrigin{
		Origin: origin,
	}
}

// Do sends the TrackCacheStorageForOrigin CDP command to a browser,
// and returns the browser's response.
func (t *TrackCacheStorageForOrigin) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Storage.trackCacheStorageForOrigin", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// TrackIndexedDBForOrigin contains the parameters, and acts as
// a Go receiver, for the CDP command `trackIndexedDBForOrigin`.
//
// Registers origin to be notified when an update occurs to its IndexedDB.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-trackIndexedDBForOrigin
type TrackIndexedDBForOrigin struct {
	// Security origin.
	Origin string `json:"origin"`
}

// NewTrackIndexedDBForOrigin constructs a new TrackIndexedDBForOrigin struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-trackIndexedDBForOrigin
func NewTrackIndexedDBForOrigin(origin string) *TrackIndexedDBForOrigin {
	return &TrackIndexedDBForOrigin{
		Origin: origin,
	}
}

// Do sends the TrackIndexedDBForOrigin CDP command to a browser,
// and returns the browser's response.
func (t *TrackIndexedDBForOrigin) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Storage.trackIndexedDBForOrigin", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// UntrackCacheStorageForOrigin contains the parameters, and acts as
// a Go receiver, for the CDP command `untrackCacheStorageForOrigin`.
//
// Unregisters origin from receiving notifications for cache storage.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-untrackCacheStorageForOrigin
type UntrackCacheStorageForOrigin struct {
	// Security origin.
	Origin string `json:"origin"`
}

// NewUntrackCacheStorageForOrigin constructs a new UntrackCacheStorageForOrigin struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-untrackCacheStorageForOrigin
func NewUntrackCacheStorageForOrigin(origin string) *UntrackCacheStorageForOrigin {
	return &UntrackCacheStorageForOrigin{
		Origin: origin,
	}
}

// Do sends the UntrackCacheStorageForOrigin CDP command to a browser,
// and returns the browser's response.
func (t *UntrackCacheStorageForOrigin) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Storage.untrackCacheStorageForOrigin", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// UntrackIndexedDBForOrigin contains the parameters, and acts as
// a Go receiver, for the CDP command `untrackIndexedDBForOrigin`.
//
// Unregisters origin from receiving notifications for IndexedDB.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-untrackIndexedDBForOrigin
type UntrackIndexedDBForOrigin struct {
	// Security origin.
	Origin string `json:"origin"`
}

// NewUntrackIndexedDBForOrigin constructs a new UntrackIndexedDBForOrigin struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-untrackIndexedDBForOrigin
func NewUntrackIndexedDBForOrigin(origin string) *UntrackIndexedDBForOrigin {
	return &UntrackIndexedDBForOrigin{
		Origin: origin,
	}
}

// Do sends the UntrackIndexedDBForOrigin CDP command to a browser,
// and returns the browser's response.
func (t *UntrackIndexedDBForOrigin) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Storage.untrackIndexedDBForOrigin", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// GetTrustTokens contains the parameters, and acts as
// a Go receiver, for the CDP command `getTrustTokens`.
//
// Returns the number of stored Trust Tokens per issuer for the
// current browsing context.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-getTrustTokens
//
// This CDP method is experimental.
type GetTrustTokens struct{}

// NewGetTrustTokens constructs a new GetTrustTokens struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-getTrustTokens
//
// This CDP method is experimental.
func NewGetTrustTokens() *GetTrustTokens {
	return &GetTrustTokens{}
}

// GetTrustTokensResult contains the browser's response
// to calling the GetTrustTokens CDP command with Do().
type GetTrustTokensResult struct {
	Tokens []TrustTokens `json:"tokens"`
}

// Do sends the GetTrustTokens CDP command to a browser,
// and returns the browser's response.
func (t *GetTrustTokens) Do(ctx context.Context) (*GetTrustTokensResult, error) {
	response, err := cdp.Send(ctx, "Storage.getTrustTokens", nil)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetTrustTokensResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ClearTrustTokens contains the parameters, and acts as
// a Go receiver, for the CDP command `clearTrustTokens`.
//
// Removes all Trust Tokens issued by the provided issuerOrigin.
// Leaves other stored data, including the issuer's Redemption Records, intact.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-clearTrustTokens
//
// This CDP method is experimental.
type ClearTrustTokens struct {
	IssuerOrigin string `json:"issuerOrigin"`
}

// NewClearTrustTokens constructs a new ClearTrustTokens struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-clearTrustTokens
//
// This CDP method is experimental.
func NewClearTrustTokens(issuerOrigin string) *ClearTrustTokens {
	return &ClearTrustTokens{
		IssuerOrigin: issuerOrigin,
	}
}

// ClearTrustTokensResult contains the browser's response
// to calling the ClearTrustTokens CDP command with Do().
type ClearTrustTokensResult struct {
	// True if any tokens were deleted, false otherwise.
	DidDeleteTokens bool `json:"didDeleteTokens"`
}

// Do sends the ClearTrustTokens CDP command to a browser,
// and returns the browser's response.
func (t *ClearTrustTokens) Do(ctx context.Context) (*ClearTrustTokensResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "Storage.clearTrustTokens", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &ClearTrustTokensResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}
