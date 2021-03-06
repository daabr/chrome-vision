package fetch

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
	"github.com/daabr/chrome-vision/pkg/devtools/network"
)

// Disable contains the parameters, and acts as
// a Go receiver, for the CDP command `disable`.
//
// Disables the fetch domain.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Fetch/#method-disable
type Disable struct{}

// NewDisable constructs a new Disable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Fetch/#method-disable
func NewDisable() *Disable {
	return &Disable{}
}

// Do sends the Disable CDP command to a browser,
// and returns the browser's response.
func (t *Disable) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Fetch.disable", nil)
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
	return devtools.Send(ctx, "Fetch.disable", nil)
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
// Enables issuing of requestPaused events. A request will be paused until client
// calls one of failRequest, fulfillRequest or continueRequest/continueWithAuth.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Fetch/#method-enable
type Enable struct {
	// If specified, only requests matching any of these patterns will produce
	// fetchRequested event and will be paused until clients response. If not set,
	// all requests will be affected.
	Patterns []RequestPattern `json:"patterns,omitempty"`
	// If true, authRequired events will be issued and requests will be paused
	// expecting a call to continueWithAuth.
	HandleAuthRequests bool `json:"handleAuthRequests,omitempty"`
}

// NewEnable constructs a new Enable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Fetch/#method-enable
func NewEnable() *Enable {
	return &Enable{}
}

// SetPatterns adds or modifies the value of the optional
// parameter `patterns` in the Enable CDP command.
//
// If specified, only requests matching any of these patterns will produce
// fetchRequested event and will be paused until clients response. If not set,
// all requests will be affected.
func (t *Enable) SetPatterns(v []RequestPattern) *Enable {
	t.Patterns = v
	return t
}

// SetHandleAuthRequests adds or modifies the value of the optional
// parameter `handleAuthRequests` in the Enable CDP command.
//
// If true, authRequired events will be issued and requests will be paused
// expecting a call to continueWithAuth.
func (t *Enable) SetHandleAuthRequests(v bool) *Enable {
	t.HandleAuthRequests = v
	return t
}

// Do sends the Enable CDP command to a browser,
// and returns the browser's response.
func (t *Enable) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Fetch.enable", b)
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
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Fetch.enable", b)
}

// ParseResponse parses the browser's response
// to the Enable CDP command.
func (t *Enable) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// FailRequest contains the parameters, and acts as
// a Go receiver, for the CDP command `failRequest`.
//
// Causes the request to fail with specified reason.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Fetch/#method-failRequest
type FailRequest struct {
	// An id the client received in requestPaused event.
	RequestID string `json:"requestId"`
	// Causes the request to fail with the given reason.
	ErrorReason network.ErrorReason `json:"errorReason"`
}

// NewFailRequest constructs a new FailRequest struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Fetch/#method-failRequest
func NewFailRequest(requestID string, errorReason network.ErrorReason) *FailRequest {
	return &FailRequest{
		RequestID:   requestID,
		ErrorReason: errorReason,
	}
}

// Do sends the FailRequest CDP command to a browser,
// and returns the browser's response.
func (t *FailRequest) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Fetch.failRequest", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the FailRequest CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *FailRequest) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Fetch.failRequest", b)
}

// ParseResponse parses the browser's response
// to the FailRequest CDP command.
func (t *FailRequest) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// FulfillRequest contains the parameters, and acts as
// a Go receiver, for the CDP command `fulfillRequest`.
//
// Provides response to the request.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Fetch/#method-fulfillRequest
type FulfillRequest struct {
	// An id the client received in requestPaused event.
	RequestID string `json:"requestId"`
	// An HTTP response code.
	ResponseCode int64 `json:"responseCode"`
	// Response headers.
	ResponseHeaders []HeaderEntry `json:"responseHeaders,omitempty"`
	// Alternative way of specifying response headers as a \0-separated
	// series of name: value pairs. Prefer the above method unless you
	// need to represent some non-UTF8 values that can't be transmitted
	// over the protocol as text. (Encoded as a base64 string when passed over JSON)
	BinaryResponseHeaders string `json:"binaryResponseHeaders,omitempty"`
	// A response body. If absent, original response body will be used if
	// the request is intercepted at the response stage and empty body
	// will be used if the request is intercepted at the request stage. (Encoded as a base64 string when passed over JSON)
	Body string `json:"body,omitempty"`
	// A textual representation of responseCode.
	// If absent, a standard phrase matching responseCode is used.
	ResponsePhrase string `json:"responsePhrase,omitempty"`
}

// NewFulfillRequest constructs a new FulfillRequest struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Fetch/#method-fulfillRequest
func NewFulfillRequest(requestID string, responseCode int64) *FulfillRequest {
	return &FulfillRequest{
		RequestID:    requestID,
		ResponseCode: responseCode,
	}
}

// SetResponseHeaders adds or modifies the value of the optional
// parameter `responseHeaders` in the FulfillRequest CDP command.
//
// Response headers.
func (t *FulfillRequest) SetResponseHeaders(v []HeaderEntry) *FulfillRequest {
	t.ResponseHeaders = v
	return t
}

// SetBinaryResponseHeaders adds or modifies the value of the optional
// parameter `binaryResponseHeaders` in the FulfillRequest CDP command.
//
// Alternative way of specifying response headers as a \0-separated
// series of name: value pairs. Prefer the above method unless you
// need to represent some non-UTF8 values that can't be transmitted
// over the protocol as text. (Encoded as a base64 string when passed over JSON)
func (t *FulfillRequest) SetBinaryResponseHeaders(v string) *FulfillRequest {
	t.BinaryResponseHeaders = v
	return t
}

// SetBody adds or modifies the value of the optional
// parameter `body` in the FulfillRequest CDP command.
//
// A response body. If absent, original response body will be used if
// the request is intercepted at the response stage and empty body
// will be used if the request is intercepted at the request stage. (Encoded as a base64 string when passed over JSON)
func (t *FulfillRequest) SetBody(v string) *FulfillRequest {
	t.Body = v
	return t
}

// SetResponsePhrase adds or modifies the value of the optional
// parameter `responsePhrase` in the FulfillRequest CDP command.
//
// A textual representation of responseCode.
// If absent, a standard phrase matching responseCode is used.
func (t *FulfillRequest) SetResponsePhrase(v string) *FulfillRequest {
	t.ResponsePhrase = v
	return t
}

// Do sends the FulfillRequest CDP command to a browser,
// and returns the browser's response.
func (t *FulfillRequest) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Fetch.fulfillRequest", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the FulfillRequest CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *FulfillRequest) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Fetch.fulfillRequest", b)
}

// ParseResponse parses the browser's response
// to the FulfillRequest CDP command.
func (t *FulfillRequest) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// ContinueRequest contains the parameters, and acts as
// a Go receiver, for the CDP command `continueRequest`.
//
// Continues the request, optionally modifying some of its parameters.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Fetch/#method-continueRequest
type ContinueRequest struct {
	// An id the client received in requestPaused event.
	RequestID string `json:"requestId"`
	// If set, the request url will be modified in a way that's not observable by page.
	URL string `json:"url,omitempty"`
	// If set, the request method is overridden.
	Method string `json:"method,omitempty"`
	// If set, overrides the post data in the request. (Encoded as a base64 string when passed over JSON)
	PostData string `json:"postData,omitempty"`
	// If set, overrides the request headers.
	Headers []HeaderEntry `json:"headers,omitempty"`
	// If set, overrides response interception behavior for this request.
	//
	// This CDP parameter is experimental.
	InterceptResponse bool `json:"interceptResponse,omitempty"`
}

// NewContinueRequest constructs a new ContinueRequest struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Fetch/#method-continueRequest
func NewContinueRequest(requestID string) *ContinueRequest {
	return &ContinueRequest{
		RequestID: requestID,
	}
}

// SetURL adds or modifies the value of the optional
// parameter `url` in the ContinueRequest CDP command.
//
// If set, the request url will be modified in a way that's not observable by page.
func (t *ContinueRequest) SetURL(v string) *ContinueRequest {
	t.URL = v
	return t
}

// SetMethod adds or modifies the value of the optional
// parameter `method` in the ContinueRequest CDP command.
//
// If set, the request method is overridden.
func (t *ContinueRequest) SetMethod(v string) *ContinueRequest {
	t.Method = v
	return t
}

// SetPostData adds or modifies the value of the optional
// parameter `postData` in the ContinueRequest CDP command.
//
// If set, overrides the post data in the request. (Encoded as a base64 string when passed over JSON)
func (t *ContinueRequest) SetPostData(v string) *ContinueRequest {
	t.PostData = v
	return t
}

// SetHeaders adds or modifies the value of the optional
// parameter `headers` in the ContinueRequest CDP command.
//
// If set, overrides the request headers.
func (t *ContinueRequest) SetHeaders(v []HeaderEntry) *ContinueRequest {
	t.Headers = v
	return t
}

// SetInterceptResponse adds or modifies the value of the optional
// parameter `interceptResponse` in the ContinueRequest CDP command.
//
// If set, overrides response interception behavior for this request.
//
// This CDP parameter is experimental.
func (t *ContinueRequest) SetInterceptResponse(v bool) *ContinueRequest {
	t.InterceptResponse = v
	return t
}

// Do sends the ContinueRequest CDP command to a browser,
// and returns the browser's response.
func (t *ContinueRequest) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Fetch.continueRequest", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the ContinueRequest CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *ContinueRequest) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Fetch.continueRequest", b)
}

// ParseResponse parses the browser's response
// to the ContinueRequest CDP command.
func (t *ContinueRequest) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// ContinueWithAuth contains the parameters, and acts as
// a Go receiver, for the CDP command `continueWithAuth`.
//
// Continues a request supplying authChallengeResponse following authRequired event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Fetch/#method-continueWithAuth
type ContinueWithAuth struct {
	// An id the client received in authRequired event.
	RequestID string `json:"requestId"`
	// Response to  with an authChallenge.
	AuthChallengeResponse AuthChallengeResponse `json:"authChallengeResponse"`
}

// NewContinueWithAuth constructs a new ContinueWithAuth struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Fetch/#method-continueWithAuth
func NewContinueWithAuth(requestID string, authChallengeResponse AuthChallengeResponse) *ContinueWithAuth {
	return &ContinueWithAuth{
		RequestID:             requestID,
		AuthChallengeResponse: authChallengeResponse,
	}
}

// Do sends the ContinueWithAuth CDP command to a browser,
// and returns the browser's response.
func (t *ContinueWithAuth) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Fetch.continueWithAuth", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the ContinueWithAuth CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *ContinueWithAuth) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Fetch.continueWithAuth", b)
}

// ParseResponse parses the browser's response
// to the ContinueWithAuth CDP command.
func (t *ContinueWithAuth) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// ContinueResponse contains the parameters, and acts as
// a Go receiver, for the CDP command `continueResponse`.
//
// Continues loading of the paused response, optionally modifying the
// response headers. If either responseCode or headers are modified, all of them
// must be present.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Fetch/#method-continueResponse
//
// This CDP method is experimental.
type ContinueResponse struct {
	// An id the client received in requestPaused event.
	RequestID string `json:"requestId"`
	// An HTTP response code. If absent, original response code will be used.
	ResponseCode int64 `json:"responseCode,omitempty"`
	// A textual representation of responseCode.
	// If absent, a standard phrase matching responseCode is used.
	ResponsePhrase string `json:"responsePhrase,omitempty"`
	// Response headers. If absent, original response headers will be used.
	ResponseHeaders []HeaderEntry `json:"responseHeaders,omitempty"`
	// Alternative way of specifying response headers as a \0-separated
	// series of name: value pairs. Prefer the above method unless you
	// need to represent some non-UTF8 values that can't be transmitted
	// over the protocol as text. (Encoded as a base64 string when passed over JSON)
	BinaryResponseHeaders string `json:"binaryResponseHeaders,omitempty"`
}

// NewContinueResponse constructs a new ContinueResponse struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Fetch/#method-continueResponse
//
// This CDP method is experimental.
func NewContinueResponse(requestID string) *ContinueResponse {
	return &ContinueResponse{
		RequestID: requestID,
	}
}

// SetResponseCode adds or modifies the value of the optional
// parameter `responseCode` in the ContinueResponse CDP command.
//
// An HTTP response code. If absent, original response code will be used.
func (t *ContinueResponse) SetResponseCode(v int64) *ContinueResponse {
	t.ResponseCode = v
	return t
}

// SetResponsePhrase adds or modifies the value of the optional
// parameter `responsePhrase` in the ContinueResponse CDP command.
//
// A textual representation of responseCode.
// If absent, a standard phrase matching responseCode is used.
func (t *ContinueResponse) SetResponsePhrase(v string) *ContinueResponse {
	t.ResponsePhrase = v
	return t
}

// SetResponseHeaders adds or modifies the value of the optional
// parameter `responseHeaders` in the ContinueResponse CDP command.
//
// Response headers. If absent, original response headers will be used.
func (t *ContinueResponse) SetResponseHeaders(v []HeaderEntry) *ContinueResponse {
	t.ResponseHeaders = v
	return t
}

// SetBinaryResponseHeaders adds or modifies the value of the optional
// parameter `binaryResponseHeaders` in the ContinueResponse CDP command.
//
// Alternative way of specifying response headers as a \0-separated
// series of name: value pairs. Prefer the above method unless you
// need to represent some non-UTF8 values that can't be transmitted
// over the protocol as text. (Encoded as a base64 string when passed over JSON)
func (t *ContinueResponse) SetBinaryResponseHeaders(v string) *ContinueResponse {
	t.BinaryResponseHeaders = v
	return t
}

// Do sends the ContinueResponse CDP command to a browser,
// and returns the browser's response.
func (t *ContinueResponse) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Fetch.continueResponse", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the ContinueResponse CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *ContinueResponse) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Fetch.continueResponse", b)
}

// ParseResponse parses the browser's response
// to the ContinueResponse CDP command.
func (t *ContinueResponse) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// GetResponseBody contains the parameters, and acts as
// a Go receiver, for the CDP command `getResponseBody`.
//
// Causes the body of the response to be received from the server and
// returned as a single string. May only be issued for a request that
// is paused in the Response stage and is mutually exclusive with
// takeResponseBodyForInterceptionAsStream. Calling other methods that
// affect the request or disabling fetch domain before body is received
// results in an undefined behavior.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Fetch/#method-getResponseBody
type GetResponseBody struct {
	// Identifier for the intercepted request to get body for.
	RequestID string `json:"requestId"`
}

// NewGetResponseBody constructs a new GetResponseBody struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Fetch/#method-getResponseBody
func NewGetResponseBody(requestID string) *GetResponseBody {
	return &GetResponseBody{
		RequestID: requestID,
	}
}

// GetResponseBodyResult contains the browser's response
// to calling the GetResponseBody CDP command with Do().
type GetResponseBodyResult struct {
	// Response body.
	Body string `json:"body"`
	// True, if content was sent as base64.
	Base64Encoded bool `json:"base64Encoded"`
}

// Do sends the GetResponseBody CDP command to a browser,
// and returns the browser's response.
func (t *GetResponseBody) Do(ctx context.Context) (*GetResponseBodyResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Fetch.getResponseBody", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetResponseBody CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetResponseBody) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Fetch.getResponseBody", b)
}

// ParseResponse parses the browser's response
// to the GetResponseBody CDP command.
func (t *GetResponseBody) ParseResponse(m *devtools.Message) (*GetResponseBodyResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetResponseBodyResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// TakeResponseBodyAsStream contains the parameters, and acts as
// a Go receiver, for the CDP command `takeResponseBodyAsStream`.
//
// Returns a handle to the stream representing the response body.
// The request must be paused in the HeadersReceived stage.
// Note that after this command the request can't be continued
// as is -- client either needs to cancel it or to provide the
// response body.
// The stream only supports sequential read, IO.read will fail if the position
// is specified.
// This method is mutually exclusive with getResponseBody.
// Calling other methods that affect the request or disabling fetch
// domain before body is received results in an undefined behavior.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Fetch/#method-takeResponseBodyAsStream
type TakeResponseBodyAsStream struct {
	RequestID string `json:"requestId"`
}

// NewTakeResponseBodyAsStream constructs a new TakeResponseBodyAsStream struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Fetch/#method-takeResponseBodyAsStream
func NewTakeResponseBodyAsStream(requestID string) *TakeResponseBodyAsStream {
	return &TakeResponseBodyAsStream{
		RequestID: requestID,
	}
}

// TakeResponseBodyAsStreamResult contains the browser's response
// to calling the TakeResponseBodyAsStream CDP command with Do().
type TakeResponseBodyAsStreamResult struct {
	Stream string `json:"stream"`
}

// Do sends the TakeResponseBodyAsStream CDP command to a browser,
// and returns the browser's response.
func (t *TakeResponseBodyAsStream) Do(ctx context.Context) (*TakeResponseBodyAsStreamResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Fetch.takeResponseBodyAsStream", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the TakeResponseBodyAsStream CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *TakeResponseBodyAsStream) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Fetch.takeResponseBodyAsStream", b)
}

// ParseResponse parses the browser's response
// to the TakeResponseBodyAsStream CDP command.
func (t *TakeResponseBodyAsStream) ParseResponse(m *devtools.Message) (*TakeResponseBodyAsStreamResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &TakeResponseBodyAsStreamResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}
