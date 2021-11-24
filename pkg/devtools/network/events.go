package network

// DataReceived asynchronous event. Fired when data chunk was received over the network.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-dataReceived
type DataReceived struct {
	// Request identifier.
	RequestID string `json:"requestId"`
	// Timestamp.
	Timestamp float64 `json:"timestamp"`
	// Data chunk length.
	DataLength int64 `json:"dataLength"`
	// Actual bytes received (might be less than dataLength for compressed encodings).
	EncodedDataLength int64 `json:"encodedDataLength"`
}

// EventSourceMessageReceived asynchronous event. Fired when EventSource message is received.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-eventSourceMessageReceived
type EventSourceMessageReceived struct {
	// Request identifier.
	RequestID string `json:"requestId"`
	// Timestamp.
	Timestamp float64 `json:"timestamp"`
	// Message type.
	EventName string `json:"eventName"`
	// Message identifier.
	EventID string `json:"eventId"`
	// Message content.
	Data string `json:"data"`
}

// LoadingFailed asynchronous event. Fired when HTTP request has failed to load.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-loadingFailed
type LoadingFailed struct {
	// Request identifier.
	RequestID string `json:"requestId"`
	// Timestamp.
	Timestamp float64 `json:"timestamp"`
	// Resource type.
	Type ResourceType `json:"type"`
	// User friendly error message.
	ErrorText string `json:"errorText"`
	// True if loading was canceled.
	Canceled bool `json:"canceled,omitempty"`
	// The reason why loading was blocked, if any.
	BlockedReason *BlockedReason `json:"blockedReason,omitempty"`
	// The reason why loading was blocked by CORS, if any.
	CorsErrorStatus *CorsErrorStatus `json:"corsErrorStatus,omitempty"`
}

// LoadingFinished asynchronous event. Fired when HTTP request has finished loading.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-loadingFinished
type LoadingFinished struct {
	// Request identifier.
	RequestID string `json:"requestId"`
	// Timestamp.
	Timestamp float64 `json:"timestamp"`
	// Total number of bytes received for this request.
	EncodedDataLength float64 `json:"encodedDataLength"`
	// Set when 1) response was blocked by Cross-Origin Read Blocking and also
	// 2) this needs to be reported to the DevTools console.
	ShouldReportCorbBlocking bool `json:"shouldReportCorbBlocking,omitempty"`
}

// RequestIntercepted asynchronous event. Details of an intercepted HTTP request, which must be either allowed, blocked, modified or
// mocked.
// Deprecated, use Fetch.requestPaused instead.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-requestIntercepted
//
// This CDP event is deprecated.
// This CDP event is experimental.
type RequestIntercepted struct {
	// Each request the page makes will have a unique id, however if any redirects are encountered
	// while processing that fetch, they will be reported with the same id as the original fetch.
	// Likewise if HTTP authentication is needed then the same fetch id will be used.
	InterceptionID string  `json:"interceptionId"`
	Request        Request `json:"request"`
	// The id of the frame that initiated the request.
	FrameID string `json:"frameId"`
	// How the requested resource will be used.
	ResourceType ResourceType `json:"resourceType"`
	// Whether this is a navigation request, which can abort the navigation completely.
	IsNavigationRequest bool `json:"isNavigationRequest"`
	// Set if the request is a navigation that will result in a download.
	// Only present after response is received from the server (i.e. HeadersReceived stage).
	IsDownload bool `json:"isDownload,omitempty"`
	// Redirect location, only sent if a redirect was intercepted.
	RedirectURL string `json:"redirectUrl,omitempty"`
	// Details of the Authorization Challenge encountered. If this is set then
	// continueInterceptedRequest must contain an authChallengeResponse.
	AuthChallenge *AuthChallenge `json:"authChallenge,omitempty"`
	// Response error if intercepted at response stage or if redirect occurred while intercepting
	// request.
	ResponseErrorReason *ErrorReason `json:"responseErrorReason,omitempty"`
	// Response code if intercepted at response stage or if redirect occurred while intercepting
	// request or auth retry occurred.
	ResponseStatusCode int64 `json:"responseStatusCode,omitempty"`
	// Response headers if intercepted at the response stage or if redirect occurred while
	// intercepting request or auth retry occurred.
	ResponseHeaders *Headers `json:"responseHeaders,omitempty"`
	// If the intercepted request had a corresponding requestWillBeSent event fired for it, then
	// this requestId will be the same as the requestId present in the requestWillBeSent event.
	RequestID string `json:"requestId,omitempty"`
}

// RequestServedFromCache asynchronous event. Fired if request ended up loading from cache.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-requestServedFromCache
type RequestServedFromCache struct {
	// Request identifier.
	RequestID string `json:"requestId"`
}

// RequestWillBeSent asynchronous event. Fired when page is about to send HTTP request.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-requestWillBeSent
type RequestWillBeSent struct {
	// Request identifier.
	RequestID string `json:"requestId"`
	// Loader identifier. Empty string if the request is fetched from worker.
	LoaderID string `json:"loaderId"`
	// URL of the document this request is loaded for.
	DocumentURL string `json:"documentURL"`
	// Request data.
	Request Request `json:"request"`
	// Timestamp.
	Timestamp float64 `json:"timestamp"`
	// Timestamp.
	WallTime float64 `json:"wallTime"`
	// Request initiator.
	Initiator Initiator `json:"initiator"`
	// In the case that redirectResponse is populated, this flag indicates whether
	// requestWillBeSentExtraInfo and responseReceivedExtraInfo events will be or were emitted
	// for the request which was just redirected.
	//
	// This CDP parameter is experimental.
	RedirectHasExtraInfo bool `json:"redirectHasExtraInfo"`
	// Redirect response data.
	RedirectResponse *Response `json:"redirectResponse,omitempty"`
	// Type of this resource.
	Type *ResourceType `json:"type,omitempty"`
	// Frame identifier.
	FrameID string `json:"frameId,omitempty"`
	// Whether the request is initiated by a user gesture. Defaults to false.
	HasUserGesture bool `json:"hasUserGesture,omitempty"`
}

// ResourceChangedPriority asynchronous event. Fired when resource loading priority is changed
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-resourceChangedPriority
//
// This CDP event is experimental.
type ResourceChangedPriority struct {
	// Request identifier.
	RequestID string `json:"requestId"`
	// New priority
	NewPriority ResourcePriority `json:"newPriority"`
	// Timestamp.
	Timestamp float64 `json:"timestamp"`
}

// SignedExchangeReceived asynchronous event. Fired when a signed exchange was received over the network
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-signedExchangeReceived
//
// This CDP event is experimental.
type SignedExchangeReceived struct {
	// Request identifier.
	RequestID string `json:"requestId"`
	// Information about the signed exchange response.
	Info SignedExchangeInfo `json:"info"`
}

// ResponseReceived asynchronous event. Fired when HTTP response is available.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-responseReceived
type ResponseReceived struct {
	// Request identifier.
	RequestID string `json:"requestId"`
	// Loader identifier. Empty string if the request is fetched from worker.
	LoaderID string `json:"loaderId"`
	// Timestamp.
	Timestamp float64 `json:"timestamp"`
	// Resource type.
	Type ResourceType `json:"type"`
	// Response data.
	Response Response `json:"response"`
	// Indicates whether requestWillBeSentExtraInfo and responseReceivedExtraInfo events will be
	// or were emitted for this request.
	//
	// This CDP parameter is experimental.
	HasExtraInfo bool `json:"hasExtraInfo"`
	// Frame identifier.
	FrameID string `json:"frameId,omitempty"`
}

// WebSocketClosed asynchronous event. Fired when WebSocket is closed.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketClosed
type WebSocketClosed struct {
	// Request identifier.
	RequestID string `json:"requestId"`
	// Timestamp.
	Timestamp float64 `json:"timestamp"`
}

// WebSocketCreated asynchronous event. Fired upon WebSocket creation.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketCreated
type WebSocketCreated struct {
	// Request identifier.
	RequestID string `json:"requestId"`
	// WebSocket request URL.
	URL string `json:"url"`
	// Request initiator.
	Initiator *Initiator `json:"initiator,omitempty"`
}

// WebSocketFrameError asynchronous event. Fired when WebSocket message error occurs.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketFrameError
type WebSocketFrameError struct {
	// Request identifier.
	RequestID string `json:"requestId"`
	// Timestamp.
	Timestamp float64 `json:"timestamp"`
	// WebSocket error message.
	ErrorMessage string `json:"errorMessage"`
}

// WebSocketFrameReceived asynchronous event. Fired when WebSocket message is received.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketFrameReceived
type WebSocketFrameReceived struct {
	// Request identifier.
	RequestID string `json:"requestId"`
	// Timestamp.
	Timestamp float64 `json:"timestamp"`
	// WebSocket response data.
	Response WebSocketFrame `json:"response"`
}

// WebSocketFrameSent asynchronous event. Fired when WebSocket message is sent.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketFrameSent
type WebSocketFrameSent struct {
	// Request identifier.
	RequestID string `json:"requestId"`
	// Timestamp.
	Timestamp float64 `json:"timestamp"`
	// WebSocket response data.
	Response WebSocketFrame `json:"response"`
}

// WebSocketHandshakeResponseReceived asynchronous event. Fired when WebSocket handshake response becomes available.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketHandshakeResponseReceived
type WebSocketHandshakeResponseReceived struct {
	// Request identifier.
	RequestID string `json:"requestId"`
	// Timestamp.
	Timestamp float64 `json:"timestamp"`
	// WebSocket response data.
	Response WebSocketResponse `json:"response"`
}

// WebSocketWillSendHandshakeRequest asynchronous event. Fired when WebSocket is about to initiate handshake.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webSocketWillSendHandshakeRequest
type WebSocketWillSendHandshakeRequest struct {
	// Request identifier.
	RequestID string `json:"requestId"`
	// Timestamp.
	Timestamp float64 `json:"timestamp"`
	// UTC Timestamp.
	WallTime float64 `json:"wallTime"`
	// WebSocket request data.
	Request WebSocketRequest `json:"request"`
}

// WebTransportCreated asynchronous event. Fired upon WebTransport creation.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webTransportCreated
type WebTransportCreated struct {
	// WebTransport identifier.
	TransportID string `json:"transportId"`
	// WebTransport request URL.
	URL string `json:"url"`
	// Timestamp.
	Timestamp float64 `json:"timestamp"`
	// Request initiator.
	Initiator *Initiator `json:"initiator,omitempty"`
}

// WebTransportConnectionEstablished asynchronous event. Fired when WebTransport handshake is finished.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webTransportConnectionEstablished
type WebTransportConnectionEstablished struct {
	// WebTransport identifier.
	TransportID string `json:"transportId"`
	// Timestamp.
	Timestamp float64 `json:"timestamp"`
}

// WebTransportClosed asynchronous event. Fired when WebTransport is disposed.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-webTransportClosed
type WebTransportClosed struct {
	// WebTransport identifier.
	TransportID string `json:"transportId"`
	// Timestamp.
	Timestamp float64 `json:"timestamp"`
}

// RequestWillBeSentExtraInfo asynchronous event. Fired when additional information about a requestWillBeSent event is available from the
// network stack. Not every requestWillBeSent event will have an additional
// requestWillBeSentExtraInfo fired for it, and there is no guarantee whether requestWillBeSent
// or requestWillBeSentExtraInfo will be fired first for the same request.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-requestWillBeSentExtraInfo
//
// This CDP event is experimental.
type RequestWillBeSentExtraInfo struct {
	// Request identifier. Used to match this information to an existing requestWillBeSent event.
	RequestID string `json:"requestId"`
	// A list of cookies potentially associated to the requested URL. This includes both cookies sent with
	// the request and the ones not sent; the latter are distinguished by having blockedReason field set.
	AssociatedCookies []BlockedCookieWithReason `json:"associatedCookies"`
	// Raw request headers as they will be sent over the wire.
	Headers Headers `json:"headers"`
	// Connection timing information for the request.
	//
	// This CDP parameter is experimental.
	ConnectTiming ConnectTiming `json:"connectTiming"`
	// The client security state set for the request.
	ClientSecurityState *ClientSecurityState `json:"clientSecurityState,omitempty"`
}

// ResponseReceivedExtraInfo asynchronous event. Fired when additional information about a responseReceived event is available from the network
// stack. Not every responseReceived event will have an additional responseReceivedExtraInfo for
// it, and responseReceivedExtraInfo may be fired before or after responseReceived.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-responseReceivedExtraInfo
//
// This CDP event is experimental.
type ResponseReceivedExtraInfo struct {
	// Request identifier. Used to match this information to another responseReceived event.
	RequestID string `json:"requestId"`
	// A list of cookies which were not stored from the response along with the corresponding
	// reasons for blocking. The cookies here may not be valid due to syntax errors, which
	// are represented by the invalid cookie line string instead of a proper cookie.
	BlockedCookies []BlockedSetCookieWithReason `json:"blockedCookies"`
	// Raw response headers as they were received over the wire.
	Headers Headers `json:"headers"`
	// The IP address space of the resource. The address space can only be determined once the transport
	// established the connection, so we can't send it in `requestWillBeSentExtraInfo`.
	ResourceIPAddressSpace IPAddressSpace `json:"resourceIPAddressSpace"`
	// The status code of the response. This is useful in cases the request failed and no responseReceived
	// event is triggered, which is the case for, e.g., CORS errors. This is also the correct status code
	// for cached requests, where the status in responseReceived is a 200 and this will be 304.
	StatusCode int64 `json:"statusCode"`
	// Raw response header text as it was received over the wire. The raw text may not always be
	// available, such as in the case of HTTP/2 or QUIC.
	HeadersText string `json:"headersText,omitempty"`
}

// TrustTokenOperationDone asynchronous event. Fired exactly once for each Trust Token operation. Depending on
// the type of the operation and whether the operation succeeded or
// failed, the event is fired before the corresponding request was sent
// or after the response was received.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-trustTokenOperationDone
//
// This CDP event is experimental.
type TrustTokenOperationDone struct {
	// Detailed success or error status of the operation.
	// 'AlreadyExists' also signifies a successful operation, as the result
	// of the operation already exists und thus, the operation was abort
	// preemptively (e.g. a cache hit).
	Status    string                  `json:"status"`
	Type      TrustTokenOperationType `json:"type"`
	RequestID string                  `json:"requestId"`
	// Top level origin. The context in which the operation was attempted.
	TopLevelOrigin string `json:"topLevelOrigin,omitempty"`
	// Origin of the issuer in case of a "Issuance" or "Redemption" operation.
	IssuerOrigin string `json:"issuerOrigin,omitempty"`
	// The number of obtained Trust Tokens on a successful "Issuance" operation.
	IssuedTokenCount int64 `json:"issuedTokenCount,omitempty"`
}

// SubresourceWebBundleMetadataReceived asynchronous event. Fired once when parsing the .wbn file has succeeded.
// The event contains the information about the web bundle contents.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-subresourceWebBundleMetadataReceived
//
// This CDP event is experimental.
type SubresourceWebBundleMetadataReceived struct {
	// Request identifier. Used to match this information to another event.
	RequestID string `json:"requestId"`
	// A list of URLs of resources in the subresource Web Bundle.
	URLs []string `json:"urls"`
}

// SubresourceWebBundleMetadataError asynchronous event. Fired once when parsing the .wbn file has failed.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-subresourceWebBundleMetadataError
//
// This CDP event is experimental.
type SubresourceWebBundleMetadataError struct {
	// Request identifier. Used to match this information to another event.
	RequestID string `json:"requestId"`
	// Error message
	ErrorMessage string `json:"errorMessage"`
}

// SubresourceWebBundleInnerResponseParsed asynchronous event. Fired when handling requests for resources within a .wbn file.
// Note: this will only be fired for resources that are requested by the webpage.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-subresourceWebBundleInnerResponseParsed
//
// This CDP event is experimental.
type SubresourceWebBundleInnerResponseParsed struct {
	// Request identifier of the subresource request
	InnerRequestID string `json:"innerRequestId"`
	// URL of the subresource resource.
	InnerRequestURL string `json:"innerRequestURL"`
	// Bundle request identifier. Used to match this information to another event.
	// This made be absent in case when the instrumentation was enabled only
	// after webbundle was parsed.
	BundleRequestID string `json:"bundleRequestId,omitempty"`
}

// SubresourceWebBundleInnerResponseError asynchronous event. Fired when request for resources within a .wbn file failed.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-subresourceWebBundleInnerResponseError
//
// This CDP event is experimental.
type SubresourceWebBundleInnerResponseError struct {
	// Request identifier of the subresource request
	InnerRequestID string `json:"innerRequestId"`
	// URL of the subresource resource.
	InnerRequestURL string `json:"innerRequestURL"`
	// Error message
	ErrorMessage string `json:"errorMessage"`
	// Bundle request identifier. Used to match this information to another event.
	// This made be absent in case when the instrumentation was enabled only
	// after webbundle was parsed.
	BundleRequestID string `json:"bundleRequestId,omitempty"`
}

// ReportingAPIReportAdded asynchronous event. Is sent whenever a new report is added.
// And after 'enableReportingApi' for all existing reports.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-reportingApiReportAdded
//
// This CDP event is experimental.
type ReportingAPIReportAdded struct {
	Report ReportingAPIReport `json:"report"`
}

// ReportingAPIReportUpdated asynchronous event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-reportingApiReportUpdated
//
// This CDP event is experimental.
type ReportingAPIReportUpdated struct {
	Report ReportingAPIReport `json:"report"`
}

// ReportingAPIEndpointsChangedForOrigin asynchronous event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#event-reportingApiEndpointsChangedForOrigin
//
// This CDP event is experimental.
type ReportingAPIEndpointsChangedForOrigin struct {
	// Origin of the document(s) which configured the endpoints.
	Origin    string                 `json:"origin"`
	Endpoints []ReportingAPIEndpoint `json:"endpoints"`
}
