package network

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/cdp"
	"github.com/daabr/chrome-vision/pkg/cdp/debugger"
)

// SetAcceptedEncodings contains the parameters, and acts as
// a Go receiver, for the CDP command `setAcceptedEncodings`.
//
// Sets a list of content encodings that will be accepted. Empty list means no encoding is accepted.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setAcceptedEncodings
//
// This CDP method is experimental.
type SetAcceptedEncodings struct {
	// List of accepted content encodings.
	Encodings []ContentEncoding
}

// NewSetAcceptedEncodings constructs a new SetAcceptedEncodings struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setAcceptedEncodings
//
// This CDP method is experimental.
func NewSetAcceptedEncodings(encodings []ContentEncoding) *SetAcceptedEncodings {
	return &SetAcceptedEncodings{
		Encodings: encodings,
	}
}

// Do sends the SetAcceptedEncodings CDP command to a browser,
// and returns the browser's response.
func (t *SetAcceptedEncodings) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Network.setAcceptedEncodings", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// ClearAcceptedEncodingsOverride contains the parameters, and acts as
// a Go receiver, for the CDP command `clearAcceptedEncodingsOverride`.
//
// Clears accepted encodings set by setAcceptedEncodings
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-clearAcceptedEncodingsOverride
//
// This CDP method is experimental.
type ClearAcceptedEncodingsOverride struct{}

// NewClearAcceptedEncodingsOverride constructs a new ClearAcceptedEncodingsOverride struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-clearAcceptedEncodingsOverride
//
// This CDP method is experimental.
func NewClearAcceptedEncodingsOverride() *ClearAcceptedEncodingsOverride {
	return &ClearAcceptedEncodingsOverride{}
}

// Do sends the ClearAcceptedEncodingsOverride CDP command to a browser,
// and returns the browser's response.
func (t *ClearAcceptedEncodingsOverride) Do(ctx context.Context) error {
	response, err := cdp.Send(ctx, "Network.clearAcceptedEncodingsOverride", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// CanClearBrowserCache contains the parameters, and acts as
// a Go receiver, for the CDP command `canClearBrowserCache`.
//
// Tells whether clearing browser cache is supported.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-canClearBrowserCache
//
// This CDP method is deprecated.
type CanClearBrowserCache struct{}

// NewCanClearBrowserCache constructs a new CanClearBrowserCache struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-canClearBrowserCache
//
// This CDP method is deprecated.
func NewCanClearBrowserCache() *CanClearBrowserCache {
	return &CanClearBrowserCache{}
}

// CanClearBrowserCacheResponse contains the browser's response
// to calling the CanClearBrowserCache CDP command with Do().
type CanClearBrowserCacheResponse struct {
	// True if browser cache can be cleared.
	Result bool
}

// Do sends the CanClearBrowserCache CDP command to a browser,
// and returns the browser's response.
func (t *CanClearBrowserCache) Do(ctx context.Context) (*CanClearBrowserCacheResponse, error) {
	response, err := cdp.Send(ctx, "Network.canClearBrowserCache", nil)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &CanClearBrowserCacheResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// CanClearBrowserCookies contains the parameters, and acts as
// a Go receiver, for the CDP command `canClearBrowserCookies`.
//
// Tells whether clearing browser cookies is supported.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-canClearBrowserCookies
//
// This CDP method is deprecated.
type CanClearBrowserCookies struct{}

// NewCanClearBrowserCookies constructs a new CanClearBrowserCookies struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-canClearBrowserCookies
//
// This CDP method is deprecated.
func NewCanClearBrowserCookies() *CanClearBrowserCookies {
	return &CanClearBrowserCookies{}
}

// CanClearBrowserCookiesResponse contains the browser's response
// to calling the CanClearBrowserCookies CDP command with Do().
type CanClearBrowserCookiesResponse struct {
	// True if browser cookies can be cleared.
	Result bool
}

// Do sends the CanClearBrowserCookies CDP command to a browser,
// and returns the browser's response.
func (t *CanClearBrowserCookies) Do(ctx context.Context) (*CanClearBrowserCookiesResponse, error) {
	response, err := cdp.Send(ctx, "Network.canClearBrowserCookies", nil)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &CanClearBrowserCookiesResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// CanEmulateNetworkConditions contains the parameters, and acts as
// a Go receiver, for the CDP command `canEmulateNetworkConditions`.
//
// Tells whether emulation of network conditions is supported.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-canEmulateNetworkConditions
//
// This CDP method is deprecated.
type CanEmulateNetworkConditions struct{}

// NewCanEmulateNetworkConditions constructs a new CanEmulateNetworkConditions struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-canEmulateNetworkConditions
//
// This CDP method is deprecated.
func NewCanEmulateNetworkConditions() *CanEmulateNetworkConditions {
	return &CanEmulateNetworkConditions{}
}

// CanEmulateNetworkConditionsResponse contains the browser's response
// to calling the CanEmulateNetworkConditions CDP command with Do().
type CanEmulateNetworkConditionsResponse struct {
	// True if emulation of network conditions is supported.
	Result bool
}

// Do sends the CanEmulateNetworkConditions CDP command to a browser,
// and returns the browser's response.
func (t *CanEmulateNetworkConditions) Do(ctx context.Context) (*CanEmulateNetworkConditionsResponse, error) {
	response, err := cdp.Send(ctx, "Network.canEmulateNetworkConditions", nil)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &CanEmulateNetworkConditionsResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ClearBrowserCache contains the parameters, and acts as
// a Go receiver, for the CDP command `clearBrowserCache`.
//
// Clears browser cache.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-clearBrowserCache
type ClearBrowserCache struct{}

// NewClearBrowserCache constructs a new ClearBrowserCache struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-clearBrowserCache
func NewClearBrowserCache() *ClearBrowserCache {
	return &ClearBrowserCache{}
}

// Do sends the ClearBrowserCache CDP command to a browser,
// and returns the browser's response.
func (t *ClearBrowserCache) Do(ctx context.Context) error {
	response, err := cdp.Send(ctx, "Network.clearBrowserCache", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// ClearBrowserCookies contains the parameters, and acts as
// a Go receiver, for the CDP command `clearBrowserCookies`.
//
// Clears browser cookies.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-clearBrowserCookies
type ClearBrowserCookies struct{}

// NewClearBrowserCookies constructs a new ClearBrowserCookies struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-clearBrowserCookies
func NewClearBrowserCookies() *ClearBrowserCookies {
	return &ClearBrowserCookies{}
}

// Do sends the ClearBrowserCookies CDP command to a browser,
// and returns the browser's response.
func (t *ClearBrowserCookies) Do(ctx context.Context) error {
	response, err := cdp.Send(ctx, "Network.clearBrowserCookies", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// ContinueInterceptedRequest contains the parameters, and acts as
// a Go receiver, for the CDP command `continueInterceptedRequest`.
//
// Response to Network.requestIntercepted which either modifies the request to continue with any
// modifications, or blocks it, or completes it with the provided response bytes. If a network
// fetch occurs as a result which encounters a redirect an additional Network.requestIntercepted
// event will be sent with the same InterceptionId.
// Deprecated, use Fetch.continueRequest, Fetch.fulfillRequest and Fetch.failRequest instead.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-continueInterceptedRequest
//
// This CDP method is deprecated.
// This CDP method is experimental.
type ContinueInterceptedRequest struct {
	InterceptionID string
	// If set this causes the request to fail with the given reason. Passing `Aborted` for requests
	// marked with `isNavigationRequest` also cancels the navigation. Must not be set in response
	// to an authChallenge.
	ErrorReason string `json:"errorReason,omitempty"`
	// If set the requests completes using with the provided base64 encoded raw response, including
	// HTTP status line and headers etc... Must not be set in response to an authChallenge. (Encoded as a base64 string when passed over JSON)
	RawResponse string `json:"rawResponse,omitempty"`
	// If set the request url will be modified in a way that's not observable by page. Must not be
	// set in response to an authChallenge.
	URL string `json:"url,omitempty"`
	// If set this allows the request method to be overridden. Must not be set in response to an
	// authChallenge.
	Method string `json:"method,omitempty"`
	// If set this allows postData to be set. Must not be set in response to an authChallenge.
	PostData string `json:"postData,omitempty"`
	// If set this allows the request headers to be changed. Must not be set in response to an
	// authChallenge.
	Headers *Headers `json:"headers,omitempty"`
	// Response to a requestIntercepted with an authChallenge. Must not be set otherwise.
	AuthChallengeResponse *AuthChallengeResponse `json:"authChallengeResponse,omitempty"`
}

// NewContinueInterceptedRequest constructs a new ContinueInterceptedRequest struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-continueInterceptedRequest
//
// This CDP method is deprecated.
// This CDP method is experimental.
func NewContinueInterceptedRequest(interceptionID string) *ContinueInterceptedRequest {
	return &ContinueInterceptedRequest{
		InterceptionID: interceptionID,
	}
}

// SetErrorReason adds or modifies the value of the optional
// parameter `errorReason` in the ContinueInterceptedRequest CDP command.
//
// If set this causes the request to fail with the given reason. Passing `Aborted` for requests
// marked with `isNavigationRequest` also cancels the navigation. Must not be set in response
// to an authChallenge.
func (t *ContinueInterceptedRequest) SetErrorReason(v string) *ContinueInterceptedRequest {
	t.ErrorReason = v
	return t
}

// SetRawResponse adds or modifies the value of the optional
// parameter `rawResponse` in the ContinueInterceptedRequest CDP command.
//
// If set the requests completes using with the provided base64 encoded raw response, including
// HTTP status line and headers etc... Must not be set in response to an authChallenge. (Encoded as a base64 string when passed over JSON)
func (t *ContinueInterceptedRequest) SetRawResponse(v string) *ContinueInterceptedRequest {
	t.RawResponse = v
	return t
}

// SetURL adds or modifies the value of the optional
// parameter `url` in the ContinueInterceptedRequest CDP command.
//
// If set the request url will be modified in a way that's not observable by page. Must not be
// set in response to an authChallenge.
func (t *ContinueInterceptedRequest) SetURL(v string) *ContinueInterceptedRequest {
	t.URL = v
	return t
}

// SetMethod adds or modifies the value of the optional
// parameter `method` in the ContinueInterceptedRequest CDP command.
//
// If set this allows the request method to be overridden. Must not be set in response to an
// authChallenge.
func (t *ContinueInterceptedRequest) SetMethod(v string) *ContinueInterceptedRequest {
	t.Method = v
	return t
}

// SetPostData adds or modifies the value of the optional
// parameter `postData` in the ContinueInterceptedRequest CDP command.
//
// If set this allows postData to be set. Must not be set in response to an authChallenge.
func (t *ContinueInterceptedRequest) SetPostData(v string) *ContinueInterceptedRequest {
	t.PostData = v
	return t
}

// SetHeaders adds or modifies the value of the optional
// parameter `headers` in the ContinueInterceptedRequest CDP command.
//
// If set this allows the request headers to be changed. Must not be set in response to an
// authChallenge.
func (t *ContinueInterceptedRequest) SetHeaders(v Headers) *ContinueInterceptedRequest {
	t.Headers = &v
	return t
}

// SetAuthChallengeResponse adds or modifies the value of the optional
// parameter `authChallengeResponse` in the ContinueInterceptedRequest CDP command.
//
// Response to a requestIntercepted with an authChallenge. Must not be set otherwise.
func (t *ContinueInterceptedRequest) SetAuthChallengeResponse(v AuthChallengeResponse) *ContinueInterceptedRequest {
	t.AuthChallengeResponse = &v
	return t
}

// Do sends the ContinueInterceptedRequest CDP command to a browser,
// and returns the browser's response.
func (t *ContinueInterceptedRequest) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Network.continueInterceptedRequest", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// DeleteCookies contains the parameters, and acts as
// a Go receiver, for the CDP command `deleteCookies`.
//
// Deletes browser cookies with matching name and url or domain/path pair.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-deleteCookies
type DeleteCookies struct {
	// Name of the cookies to remove.
	Name string
	// If specified, deletes all the cookies with the given name where domain and path match
	// provided URL.
	URL string `json:"url,omitempty"`
	// If specified, deletes only cookies with the exact domain.
	Domain string `json:"domain,omitempty"`
	// If specified, deletes only cookies with the exact path.
	Path string `json:"path,omitempty"`
}

// NewDeleteCookies constructs a new DeleteCookies struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-deleteCookies
func NewDeleteCookies(name string) *DeleteCookies {
	return &DeleteCookies{
		Name: name,
	}
}

// SetURL adds or modifies the value of the optional
// parameter `url` in the DeleteCookies CDP command.
//
// If specified, deletes all the cookies with the given name where domain and path match
// provided URL.
func (t *DeleteCookies) SetURL(v string) *DeleteCookies {
	t.URL = v
	return t
}

// SetDomain adds or modifies the value of the optional
// parameter `domain` in the DeleteCookies CDP command.
//
// If specified, deletes only cookies with the exact domain.
func (t *DeleteCookies) SetDomain(v string) *DeleteCookies {
	t.Domain = v
	return t
}

// SetPath adds or modifies the value of the optional
// parameter `path` in the DeleteCookies CDP command.
//
// If specified, deletes only cookies with the exact path.
func (t *DeleteCookies) SetPath(v string) *DeleteCookies {
	t.Path = v
	return t
}

// Do sends the DeleteCookies CDP command to a browser,
// and returns the browser's response.
func (t *DeleteCookies) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Network.deleteCookies", b)
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
// Disables network tracking, prevents network events from being sent to the client.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-disable
type Disable struct{}

// NewDisable constructs a new Disable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-disable
func NewDisable() *Disable {
	return &Disable{}
}

// Do sends the Disable CDP command to a browser,
// and returns the browser's response.
func (t *Disable) Do(ctx context.Context) error {
	response, err := cdp.Send(ctx, "Network.disable", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// EmulateNetworkConditions contains the parameters, and acts as
// a Go receiver, for the CDP command `emulateNetworkConditions`.
//
// Activates emulation of network conditions.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-emulateNetworkConditions
type EmulateNetworkConditions struct {
	// True to emulate internet disconnection.
	Offline bool
	// Minimum latency from request sent to response headers received (ms).
	Latency float64
	// Maximal aggregated download throughput (bytes/sec). -1 disables download throttling.
	DownloadThroughput float64
	// Maximal aggregated upload throughput (bytes/sec).  -1 disables upload throttling.
	UploadThroughput float64
	// Connection type if known.
	ConnectionType string `json:"connectionType,omitempty"`
}

// NewEmulateNetworkConditions constructs a new EmulateNetworkConditions struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-emulateNetworkConditions
func NewEmulateNetworkConditions(offline bool, latency float64, downloadThroughput float64, uploadThroughput float64) *EmulateNetworkConditions {
	return &EmulateNetworkConditions{
		Offline:            offline,
		Latency:            latency,
		DownloadThroughput: downloadThroughput,
		UploadThroughput:   uploadThroughput,
	}
}

// SetConnectionType adds or modifies the value of the optional
// parameter `connectionType` in the EmulateNetworkConditions CDP command.
//
// Connection type if known.
func (t *EmulateNetworkConditions) SetConnectionType(v string) *EmulateNetworkConditions {
	t.ConnectionType = v
	return t
}

// Do sends the EmulateNetworkConditions CDP command to a browser,
// and returns the browser's response.
func (t *EmulateNetworkConditions) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Network.emulateNetworkConditions", b)
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
// Enables network tracking, network events will now be delivered to the client.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-enable
type Enable struct {
	// Buffer size in bytes to use when preserving network payloads (XHRs, etc).
	//
	// This CDP parameter is experimental.
	MaxTotalBufferSize int64 `json:"maxTotalBufferSize,omitempty"`
	// Per-resource buffer size in bytes to use when preserving network payloads (XHRs, etc).
	//
	// This CDP parameter is experimental.
	MaxResourceBufferSize int64 `json:"maxResourceBufferSize,omitempty"`
	// Longest post body size (in bytes) that would be included in requestWillBeSent notification
	MaxPostDataSize int64 `json:"maxPostDataSize,omitempty"`
}

// NewEnable constructs a new Enable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-enable
func NewEnable() *Enable {
	return &Enable{}
}

// SetMaxTotalBufferSize adds or modifies the value of the optional
// parameter `maxTotalBufferSize` in the Enable CDP command.
//
// Buffer size in bytes to use when preserving network payloads (XHRs, etc).
//
// This CDP parameter is experimental.
func (t *Enable) SetMaxTotalBufferSize(v int64) *Enable {
	t.MaxTotalBufferSize = v
	return t
}

// SetMaxResourceBufferSize adds or modifies the value of the optional
// parameter `maxResourceBufferSize` in the Enable CDP command.
//
// Per-resource buffer size in bytes to use when preserving network payloads (XHRs, etc).
//
// This CDP parameter is experimental.
func (t *Enable) SetMaxResourceBufferSize(v int64) *Enable {
	t.MaxResourceBufferSize = v
	return t
}

// SetMaxPostDataSize adds or modifies the value of the optional
// parameter `maxPostDataSize` in the Enable CDP command.
//
// Longest post body size (in bytes) that would be included in requestWillBeSent notification
func (t *Enable) SetMaxPostDataSize(v int64) *Enable {
	t.MaxPostDataSize = v
	return t
}

// Do sends the Enable CDP command to a browser,
// and returns the browser's response.
func (t *Enable) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Network.enable", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// GetAllCookies contains the parameters, and acts as
// a Go receiver, for the CDP command `getAllCookies`.
//
// Returns all browser cookies. Depending on the backend support, will return detailed cookie
// information in the `cookies` field.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getAllCookies
type GetAllCookies struct{}

// NewGetAllCookies constructs a new GetAllCookies struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getAllCookies
func NewGetAllCookies() *GetAllCookies {
	return &GetAllCookies{}
}

// GetAllCookiesResponse contains the browser's response
// to calling the GetAllCookies CDP command with Do().
type GetAllCookiesResponse struct {
	// Array of cookie objects.
	Cookies []Cookie
}

// Do sends the GetAllCookies CDP command to a browser,
// and returns the browser's response.
func (t *GetAllCookies) Do(ctx context.Context) (*GetAllCookiesResponse, error) {
	response, err := cdp.Send(ctx, "Network.getAllCookies", nil)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetAllCookiesResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetCertificate contains the parameters, and acts as
// a Go receiver, for the CDP command `getCertificate`.
//
// Returns the DER-encoded certificate.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getCertificate
//
// This CDP method is experimental.
type GetCertificate struct {
	// Origin to get certificate for.
	Origin string
}

// NewGetCertificate constructs a new GetCertificate struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getCertificate
//
// This CDP method is experimental.
func NewGetCertificate(origin string) *GetCertificate {
	return &GetCertificate{
		Origin: origin,
	}
}

// GetCertificateResponse contains the browser's response
// to calling the GetCertificate CDP command with Do().
type GetCertificateResponse struct {
	TableNames []string
}

// Do sends the GetCertificate CDP command to a browser,
// and returns the browser's response.
func (t *GetCertificate) Do(ctx context.Context) (*GetCertificateResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "Network.getCertificate", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetCertificateResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetCookies contains the parameters, and acts as
// a Go receiver, for the CDP command `getCookies`.
//
// Returns all browser cookies for the current URL. Depending on the backend support, will return
// detailed cookie information in the `cookies` field.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getCookies
type GetCookies struct {
	// The list of URLs for which applicable cookies will be fetched.
	// If not specified, it's assumed to be set to the list containing
	// the URLs of the page and all of its subframes.
	URLs []string `json:"urls,omitempty"`
}

// NewGetCookies constructs a new GetCookies struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getCookies
func NewGetCookies() *GetCookies {
	return &GetCookies{}
}

// SetURLs adds or modifies the value of the optional
// parameter `urls` in the GetCookies CDP command.
//
// The list of URLs for which applicable cookies will be fetched.
// If not specified, it's assumed to be set to the list containing
// the URLs of the page and all of its subframes.
func (t *GetCookies) SetURLs(v []string) *GetCookies {
	t.URLs = v
	return t
}

// GetCookiesResponse contains the browser's response
// to calling the GetCookies CDP command with Do().
type GetCookiesResponse struct {
	// Array of cookie objects.
	Cookies []Cookie
}

// Do sends the GetCookies CDP command to a browser,
// and returns the browser's response.
func (t *GetCookies) Do(ctx context.Context) (*GetCookiesResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "Network.getCookies", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetCookiesResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetResponseBody contains the parameters, and acts as
// a Go receiver, for the CDP command `getResponseBody`.
//
// Returns content served for the given request.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getResponseBody
type GetResponseBody struct {
	// Identifier of the network request to get content for.
	RequestID string
}

// NewGetResponseBody constructs a new GetResponseBody struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getResponseBody
func NewGetResponseBody(requestID string) *GetResponseBody {
	return &GetResponseBody{
		RequestID: requestID,
	}
}

// GetResponseBodyResponse contains the browser's response
// to calling the GetResponseBody CDP command with Do().
type GetResponseBodyResponse struct {
	// Response body.
	Body string
	// True, if content was sent as base64.
	Base64Encoded bool
}

// Do sends the GetResponseBody CDP command to a browser,
// and returns the browser's response.
func (t *GetResponseBody) Do(ctx context.Context) (*GetResponseBodyResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "Network.getResponseBody", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetResponseBodyResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetRequestPostData contains the parameters, and acts as
// a Go receiver, for the CDP command `getRequestPostData`.
//
// Returns post data sent with the request. Returns an error when no data was sent with the request.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getRequestPostData
type GetRequestPostData struct {
	// Identifier of the network request to get content for.
	RequestID string
}

// NewGetRequestPostData constructs a new GetRequestPostData struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getRequestPostData
func NewGetRequestPostData(requestID string) *GetRequestPostData {
	return &GetRequestPostData{
		RequestID: requestID,
	}
}

// GetRequestPostDataResponse contains the browser's response
// to calling the GetRequestPostData CDP command with Do().
type GetRequestPostDataResponse struct {
	// Request body string, omitting files from multipart requests
	PostData string
}

// Do sends the GetRequestPostData CDP command to a browser,
// and returns the browser's response.
func (t *GetRequestPostData) Do(ctx context.Context) (*GetRequestPostDataResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "Network.getRequestPostData", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetRequestPostDataResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetResponseBodyForInterception contains the parameters, and acts as
// a Go receiver, for the CDP command `getResponseBodyForInterception`.
//
// Returns content served for the given currently intercepted request.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getResponseBodyForInterception
//
// This CDP method is experimental.
type GetResponseBodyForInterception struct {
	// Identifier for the intercepted request to get body for.
	InterceptionID string
}

// NewGetResponseBodyForInterception constructs a new GetResponseBodyForInterception struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getResponseBodyForInterception
//
// This CDP method is experimental.
func NewGetResponseBodyForInterception(interceptionID string) *GetResponseBodyForInterception {
	return &GetResponseBodyForInterception{
		InterceptionID: interceptionID,
	}
}

// GetResponseBodyForInterceptionResponse contains the browser's response
// to calling the GetResponseBodyForInterception CDP command with Do().
type GetResponseBodyForInterceptionResponse struct {
	// Response body.
	Body string
	// True, if content was sent as base64.
	Base64Encoded bool
}

// Do sends the GetResponseBodyForInterception CDP command to a browser,
// and returns the browser's response.
func (t *GetResponseBodyForInterception) Do(ctx context.Context) (*GetResponseBodyForInterceptionResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "Network.getResponseBodyForInterception", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetResponseBodyForInterceptionResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// TakeResponseBodyForInterceptionAsStream contains the parameters, and acts as
// a Go receiver, for the CDP command `takeResponseBodyForInterceptionAsStream`.
//
// Returns a handle to the stream representing the response body. Note that after this command,
// the intercepted request can't be continued as is -- you either need to cancel it or to provide
// the response body. The stream only supports sequential read, IO.read will fail if the position
// is specified.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-takeResponseBodyForInterceptionAsStream
//
// This CDP method is experimental.
type TakeResponseBodyForInterceptionAsStream struct {
	InterceptionID string
}

// NewTakeResponseBodyForInterceptionAsStream constructs a new TakeResponseBodyForInterceptionAsStream struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-takeResponseBodyForInterceptionAsStream
//
// This CDP method is experimental.
func NewTakeResponseBodyForInterceptionAsStream(interceptionID string) *TakeResponseBodyForInterceptionAsStream {
	return &TakeResponseBodyForInterceptionAsStream{
		InterceptionID: interceptionID,
	}
}

// TakeResponseBodyForInterceptionAsStreamResponse contains the browser's response
// to calling the TakeResponseBodyForInterceptionAsStream CDP command with Do().
type TakeResponseBodyForInterceptionAsStreamResponse struct {
	Stream string
}

// Do sends the TakeResponseBodyForInterceptionAsStream CDP command to a browser,
// and returns the browser's response.
func (t *TakeResponseBodyForInterceptionAsStream) Do(ctx context.Context) (*TakeResponseBodyForInterceptionAsStreamResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "Network.takeResponseBodyForInterceptionAsStream", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &TakeResponseBodyForInterceptionAsStreamResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ReplayXHR contains the parameters, and acts as
// a Go receiver, for the CDP command `replayXHR`.
//
// This method sends a new XMLHttpRequest which is identical to the original one. The following
// parameters should be identical: method, url, async, request body, extra headers, withCredentials
// attribute, user, password.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-replayXHR
//
// This CDP method is experimental.
type ReplayXHR struct {
	// Identifier of XHR to replay.
	RequestID string
}

// NewReplayXHR constructs a new ReplayXHR struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-replayXHR
//
// This CDP method is experimental.
func NewReplayXHR(requestID string) *ReplayXHR {
	return &ReplayXHR{
		RequestID: requestID,
	}
}

// Do sends the ReplayXHR CDP command to a browser,
// and returns the browser's response.
func (t *ReplayXHR) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Network.replayXHR", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SearchInResponseBody contains the parameters, and acts as
// a Go receiver, for the CDP command `searchInResponseBody`.
//
// Searches for given string in response content.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-searchInResponseBody
//
// This CDP method is experimental.
type SearchInResponseBody struct {
	// Identifier of the network response to search.
	RequestID string
	// String to search for.
	Query string
	// If true, search is case sensitive.
	CaseSensitive bool `json:"caseSensitive,omitempty"`
	// If true, treats string parameter as regex.
	IsRegex bool `json:"isRegex,omitempty"`
}

// NewSearchInResponseBody constructs a new SearchInResponseBody struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-searchInResponseBody
//
// This CDP method is experimental.
func NewSearchInResponseBody(requestID string, query string) *SearchInResponseBody {
	return &SearchInResponseBody{
		RequestID: requestID,
		Query:     query,
	}
}

// SetCaseSensitive adds or modifies the value of the optional
// parameter `caseSensitive` in the SearchInResponseBody CDP command.
//
// If true, search is case sensitive.
func (t *SearchInResponseBody) SetCaseSensitive(v bool) *SearchInResponseBody {
	t.CaseSensitive = v
	return t
}

// SetIsRegex adds or modifies the value of the optional
// parameter `isRegex` in the SearchInResponseBody CDP command.
//
// If true, treats string parameter as regex.
func (t *SearchInResponseBody) SetIsRegex(v bool) *SearchInResponseBody {
	t.IsRegex = v
	return t
}

// SearchInResponseBodyResponse contains the browser's response
// to calling the SearchInResponseBody CDP command with Do().
type SearchInResponseBodyResponse struct {
	// List of search matches.
	Result []debugger.SearchMatch
}

// Do sends the SearchInResponseBody CDP command to a browser,
// and returns the browser's response.
func (t *SearchInResponseBody) Do(ctx context.Context) (*SearchInResponseBodyResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "Network.searchInResponseBody", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &SearchInResponseBodyResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetBlockedURLs contains the parameters, and acts as
// a Go receiver, for the CDP command `setBlockedURLs`.
//
// Blocks URLs from loading.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setBlockedURLs
//
// This CDP method is experimental.
type SetBlockedURLs struct {
	// URL patterns to block. Wildcards ('*') are allowed.
	URLs []string
}

// NewSetBlockedURLs constructs a new SetBlockedURLs struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setBlockedURLs
//
// This CDP method is experimental.
func NewSetBlockedURLs(urls []string) *SetBlockedURLs {
	return &SetBlockedURLs{
		URLs: urls,
	}
}

// Do sends the SetBlockedURLs CDP command to a browser,
// and returns the browser's response.
func (t *SetBlockedURLs) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Network.setBlockedURLs", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetBypassServiceWorker contains the parameters, and acts as
// a Go receiver, for the CDP command `setBypassServiceWorker`.
//
// Toggles ignoring of service worker for each request.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setBypassServiceWorker
//
// This CDP method is experimental.
type SetBypassServiceWorker struct {
	// Bypass service worker and load from network.
	Bypass bool
}

// NewSetBypassServiceWorker constructs a new SetBypassServiceWorker struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setBypassServiceWorker
//
// This CDP method is experimental.
func NewSetBypassServiceWorker(bypass bool) *SetBypassServiceWorker {
	return &SetBypassServiceWorker{
		Bypass: bypass,
	}
}

// Do sends the SetBypassServiceWorker CDP command to a browser,
// and returns the browser's response.
func (t *SetBypassServiceWorker) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Network.setBypassServiceWorker", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetCacheDisabled contains the parameters, and acts as
// a Go receiver, for the CDP command `setCacheDisabled`.
//
// Toggles ignoring cache for each request. If `true`, cache will not be used.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setCacheDisabled
type SetCacheDisabled struct {
	// Cache disabled state.
	CacheDisabled bool
}

// NewSetCacheDisabled constructs a new SetCacheDisabled struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setCacheDisabled
func NewSetCacheDisabled(cacheDisabled bool) *SetCacheDisabled {
	return &SetCacheDisabled{
		CacheDisabled: cacheDisabled,
	}
}

// Do sends the SetCacheDisabled CDP command to a browser,
// and returns the browser's response.
func (t *SetCacheDisabled) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Network.setCacheDisabled", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetCookie contains the parameters, and acts as
// a Go receiver, for the CDP command `setCookie`.
//
// Sets a cookie with the given cookie data; may overwrite equivalent cookies if they exist.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setCookie
type SetCookie struct {
	// Cookie name.
	Name string
	// Cookie value.
	Value string
	// The request-URI to associate with the setting of the cookie. This value can affect the
	// default domain, path, source port, and source scheme values of the created cookie.
	URL string `json:"url,omitempty"`
	// Cookie domain.
	Domain string `json:"domain,omitempty"`
	// Cookie path.
	Path string `json:"path,omitempty"`
	// True if cookie is secure.
	Secure bool `json:"secure,omitempty"`
	// True if cookie is http-only.
	HttpOnly bool `json:"httpOnly,omitempty"`
	// Cookie SameSite type.
	SameSite string `json:"sameSite,omitempty"`
	// Cookie expiration date, session cookie if not set
	Expires float64 `json:"expires,omitempty"`
	// Cookie Priority type.
	//
	// This CDP parameter is experimental.
	Priority string `json:"priority,omitempty"`
	// True if cookie is SameParty.
	//
	// This CDP parameter is experimental.
	SameParty bool `json:"sameParty,omitempty"`
	// Cookie source scheme type.
	//
	// This CDP parameter is experimental.
	SourceScheme string `json:"sourceScheme,omitempty"`
	// Cookie source port. Valid values are {-1, [1, 65535]}, -1 indicates an unspecified port.
	// An unspecified port value allows protocol clients to emulate legacy cookie scope for the port.
	// This is a temporary ability and it will be removed in the future.
	//
	// This CDP parameter is experimental.
	SourcePort int64 `json:"sourcePort,omitempty"`
}

// NewSetCookie constructs a new SetCookie struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setCookie
func NewSetCookie(name string, value string) *SetCookie {
	return &SetCookie{
		Name:  name,
		Value: value,
	}
}

// SetURL adds or modifies the value of the optional
// parameter `url` in the SetCookie CDP command.
//
// The request-URI to associate with the setting of the cookie. This value can affect the
// default domain, path, source port, and source scheme values of the created cookie.
func (t *SetCookie) SetURL(v string) *SetCookie {
	t.URL = v
	return t
}

// SetDomain adds or modifies the value of the optional
// parameter `domain` in the SetCookie CDP command.
//
// Cookie domain.
func (t *SetCookie) SetDomain(v string) *SetCookie {
	t.Domain = v
	return t
}

// SetPath adds or modifies the value of the optional
// parameter `path` in the SetCookie CDP command.
//
// Cookie path.
func (t *SetCookie) SetPath(v string) *SetCookie {
	t.Path = v
	return t
}

// SetSecure adds or modifies the value of the optional
// parameter `secure` in the SetCookie CDP command.
//
// True if cookie is secure.
func (t *SetCookie) SetSecure(v bool) *SetCookie {
	t.Secure = v
	return t
}

// SetHttpOnly adds or modifies the value of the optional
// parameter `httpOnly` in the SetCookie CDP command.
//
// True if cookie is http-only.
func (t *SetCookie) SetHttpOnly(v bool) *SetCookie {
	t.HttpOnly = v
	return t
}

// SetSameSite adds or modifies the value of the optional
// parameter `sameSite` in the SetCookie CDP command.
//
// Cookie SameSite type.
func (t *SetCookie) SetSameSite(v string) *SetCookie {
	t.SameSite = v
	return t
}

// SetExpires adds or modifies the value of the optional
// parameter `expires` in the SetCookie CDP command.
//
// Cookie expiration date, session cookie if not set
func (t *SetCookie) SetExpires(v float64) *SetCookie {
	t.Expires = v
	return t
}

// SetPriority adds or modifies the value of the optional
// parameter `priority` in the SetCookie CDP command.
//
// Cookie Priority type.
//
// This CDP parameter is experimental.
func (t *SetCookie) SetPriority(v string) *SetCookie {
	t.Priority = v
	return t
}

// SetSameParty adds or modifies the value of the optional
// parameter `sameParty` in the SetCookie CDP command.
//
// True if cookie is SameParty.
//
// This CDP parameter is experimental.
func (t *SetCookie) SetSameParty(v bool) *SetCookie {
	t.SameParty = v
	return t
}

// SetSourceScheme adds or modifies the value of the optional
// parameter `sourceScheme` in the SetCookie CDP command.
//
// Cookie source scheme type.
//
// This CDP parameter is experimental.
func (t *SetCookie) SetSourceScheme(v string) *SetCookie {
	t.SourceScheme = v
	return t
}

// SetSourcePort adds or modifies the value of the optional
// parameter `sourcePort` in the SetCookie CDP command.
//
// Cookie source port. Valid values are {-1, [1, 65535]}, -1 indicates an unspecified port.
// An unspecified port value allows protocol clients to emulate legacy cookie scope for the port.
// This is a temporary ability and it will be removed in the future.
//
// This CDP parameter is experimental.
func (t *SetCookie) SetSourcePort(v int64) *SetCookie {
	t.SourcePort = v
	return t
}

// SetCookieResponse contains the browser's response
// to calling the SetCookie CDP command with Do().
type SetCookieResponse struct {
	// Always set to true. If an error occurs, the response indicates protocol error.
	//
	// This CDP parameter is deprecated.
	Success bool
}

// Do sends the SetCookie CDP command to a browser,
// and returns the browser's response.
func (t *SetCookie) Do(ctx context.Context) (*SetCookieResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "Network.setCookie", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &SetCookieResponse{}
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
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setCookies
type SetCookies struct {
	// Cookies to be set.
	Cookies []CookieParam
}

// NewSetCookies constructs a new SetCookies struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setCookies
func NewSetCookies(cookies []CookieParam) *SetCookies {
	return &SetCookies{
		Cookies: cookies,
	}
}

// Do sends the SetCookies CDP command to a browser,
// and returns the browser's response.
func (t *SetCookies) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Network.setCookies", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetDataSizeLimitsForTest contains the parameters, and acts as
// a Go receiver, for the CDP command `setDataSizeLimitsForTest`.
//
// For testing.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setDataSizeLimitsForTest
//
// This CDP method is experimental.
type SetDataSizeLimitsForTest struct {
	// Maximum total buffer size.
	MaxTotalSize int64
	// Maximum per-resource size.
	MaxResourceSize int64
}

// NewSetDataSizeLimitsForTest constructs a new SetDataSizeLimitsForTest struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setDataSizeLimitsForTest
//
// This CDP method is experimental.
func NewSetDataSizeLimitsForTest(maxTotalSize int64, maxResourceSize int64) *SetDataSizeLimitsForTest {
	return &SetDataSizeLimitsForTest{
		MaxTotalSize:    maxTotalSize,
		MaxResourceSize: maxResourceSize,
	}
}

// Do sends the SetDataSizeLimitsForTest CDP command to a browser,
// and returns the browser's response.
func (t *SetDataSizeLimitsForTest) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Network.setDataSizeLimitsForTest", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetExtraHTTPHeaders contains the parameters, and acts as
// a Go receiver, for the CDP command `setExtraHTTPHeaders`.
//
// Specifies whether to always send extra HTTP headers with the requests from this page.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setExtraHTTPHeaders
type SetExtraHTTPHeaders struct {
	// Map with extra HTTP headers.
	Headers Headers
}

// NewSetExtraHTTPHeaders constructs a new SetExtraHTTPHeaders struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setExtraHTTPHeaders
func NewSetExtraHTTPHeaders(headers Headers) *SetExtraHTTPHeaders {
	return &SetExtraHTTPHeaders{
		Headers: headers,
	}
}

// Do sends the SetExtraHTTPHeaders CDP command to a browser,
// and returns the browser's response.
func (t *SetExtraHTTPHeaders) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Network.setExtraHTTPHeaders", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetAttachDebugStack contains the parameters, and acts as
// a Go receiver, for the CDP command `setAttachDebugStack`.
//
// Specifies whether to attach a page script stack id in requests
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setAttachDebugStack
//
// This CDP method is experimental.
type SetAttachDebugStack struct {
	// Whether to attach a page script stack for debugging purpose.
	Enabled bool
}

// NewSetAttachDebugStack constructs a new SetAttachDebugStack struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setAttachDebugStack
//
// This CDP method is experimental.
func NewSetAttachDebugStack(enabled bool) *SetAttachDebugStack {
	return &SetAttachDebugStack{
		Enabled: enabled,
	}
}

// Do sends the SetAttachDebugStack CDP command to a browser,
// and returns the browser's response.
func (t *SetAttachDebugStack) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Network.setAttachDebugStack", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetRequestInterception contains the parameters, and acts as
// a Go receiver, for the CDP command `setRequestInterception`.
//
// Sets the requests to intercept that match the provided patterns and optionally resource types.
// Deprecated, please use Fetch.enable instead.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setRequestInterception
//
// This CDP method is deprecated.
// This CDP method is experimental.
type SetRequestInterception struct {
	// Requests matching any of these patterns will be forwarded and wait for the corresponding
	// continueInterceptedRequest call.
	Patterns []RequestPattern
}

// NewSetRequestInterception constructs a new SetRequestInterception struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-setRequestInterception
//
// This CDP method is deprecated.
// This CDP method is experimental.
func NewSetRequestInterception(patterns []RequestPattern) *SetRequestInterception {
	return &SetRequestInterception{
		Patterns: patterns,
	}
}

// Do sends the SetRequestInterception CDP command to a browser,
// and returns the browser's response.
func (t *SetRequestInterception) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Network.setRequestInterception", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// GetSecurityIsolationStatus contains the parameters, and acts as
// a Go receiver, for the CDP command `getSecurityIsolationStatus`.
//
// Returns information about the COEP/COOP isolation status.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getSecurityIsolationStatus
//
// This CDP method is experimental.
type GetSecurityIsolationStatus struct {
	// If no frameId is provided, the status of the target is provided.
	FrameID string `json:"frameId,omitempty"`
}

// NewGetSecurityIsolationStatus constructs a new GetSecurityIsolationStatus struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getSecurityIsolationStatus
//
// This CDP method is experimental.
func NewGetSecurityIsolationStatus() *GetSecurityIsolationStatus {
	return &GetSecurityIsolationStatus{}
}

// SetFrameID adds or modifies the value of the optional
// parameter `frameId` in the GetSecurityIsolationStatus CDP command.
//
// If no frameId is provided, the status of the target is provided.
func (t *GetSecurityIsolationStatus) SetFrameID(v string) *GetSecurityIsolationStatus {
	t.FrameID = v
	return t
}

// GetSecurityIsolationStatusResponse contains the browser's response
// to calling the GetSecurityIsolationStatus CDP command with Do().
type GetSecurityIsolationStatusResponse struct {
	Status SecurityIsolationStatus
}

// Do sends the GetSecurityIsolationStatus CDP command to a browser,
// and returns the browser's response.
func (t *GetSecurityIsolationStatus) Do(ctx context.Context) (*GetSecurityIsolationStatusResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "Network.getSecurityIsolationStatus", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetSecurityIsolationStatusResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// LoadNetworkResource contains the parameters, and acts as
// a Go receiver, for the CDP command `loadNetworkResource`.
//
// Fetches the resource and returns the content.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-loadNetworkResource
//
// This CDP method is experimental.
type LoadNetworkResource struct {
	// Frame id to get the resource for.
	FrameID string
	// URL of the resource to get content for.
	URL string
	// Options for the request.
	Options LoadNetworkResourceOptions
}

// NewLoadNetworkResource constructs a new LoadNetworkResource struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-loadNetworkResource
//
// This CDP method is experimental.
func NewLoadNetworkResource(frameID string, url string, options LoadNetworkResourceOptions) *LoadNetworkResource {
	return &LoadNetworkResource{
		FrameID: frameID,
		URL:     url,
		Options: options,
	}
}

// LoadNetworkResourceResponse contains the browser's response
// to calling the LoadNetworkResource CDP command with Do().
type LoadNetworkResourceResponse struct {
	Resource LoadNetworkResourcePageResult
}

// Do sends the LoadNetworkResource CDP command to a browser,
// and returns the browser's response.
func (t *LoadNetworkResource) Do(ctx context.Context) (*LoadNetworkResourceResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "Network.loadNetworkResource", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &LoadNetworkResourceResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}
