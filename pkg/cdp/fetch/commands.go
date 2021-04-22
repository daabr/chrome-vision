package fetch

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/cdp"
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
	response, err := cdp.Send(ctx, "Fetch.disable", nil)
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
	response, err := cdp.Send(ctx, "Fetch.enable", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
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
	RequestID string
	// Causes the request to fail with the given reason.
	ErrorReason string
}

// NewFailRequest constructs a new FailRequest struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Fetch/#method-failRequest
func NewFailRequest(requestID string, errorReason string) *FailRequest {
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
	response, err := cdp.Send(ctx, "Fetch.failRequest", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
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
	RequestID string
	// An HTTP response code.
	ResponseCode int64
	// Response headers.
	ResponseHeaders []HeaderEntry `json:"responseHeaders,omitempty"`
	// Alternative way of specifying response headers as a \0-separated
	// series of name: value pairs. Prefer the above method unless you
	// need to represent some non-UTF8 values that can't be transmitted
	// over the protocol as text. (Encoded as a base64 string when passed over JSON)
	BinaryResponseHeaders string `json:"binaryResponseHeaders,omitempty"`
	// A response body. (Encoded as a base64 string when passed over JSON)
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
// A response body. (Encoded as a base64 string when passed over JSON)
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
	response, err := cdp.Send(ctx, "Fetch.fulfillRequest", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
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
	RequestID string
	// If set, the request url will be modified in a way that's not observable by page.
	URL string `json:"url,omitempty"`
	// If set, the request method is overridden.
	Method string `json:"method,omitempty"`
	// If set, overrides the post data in the request. (Encoded as a base64 string when passed over JSON)
	PostData string `json:"postData,omitempty"`
	// If set, overrides the request headers.
	Headers []HeaderEntry `json:"headers,omitempty"`
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

// Do sends the ContinueRequest CDP command to a browser,
// and returns the browser's response.
func (t *ContinueRequest) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Fetch.continueRequest", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
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
	RequestID string
	// Response to  with an authChallenge.
	AuthChallengeResponse AuthChallengeResponse
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
	response, err := cdp.Send(ctx, "Fetch.continueWithAuth", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
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
	RequestID string
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
	response, err := cdp.Send(ctx, "Fetch.getResponseBody", b)
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
	RequestID string
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

// TakeResponseBodyAsStreamResponse contains the browser's response
// to calling the TakeResponseBodyAsStream CDP command with Do().
type TakeResponseBodyAsStreamResponse struct {
	Stream string
}

// Do sends the TakeResponseBodyAsStream CDP command to a browser,
// and returns the browser's response.
func (t *TakeResponseBodyAsStream) Do(ctx context.Context) (*TakeResponseBodyAsStreamResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "Fetch.takeResponseBodyAsStream", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &TakeResponseBodyAsStreamResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}
