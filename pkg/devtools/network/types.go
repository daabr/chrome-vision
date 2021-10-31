package network

import (
	"encoding/json"

	"github.com/daabr/chrome-vision/pkg/devtools/runtime"
	"github.com/daabr/chrome-vision/pkg/devtools/security"
)

// ResourceType data type. Resource type as it was perceived by the rendering engine.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-ResourceType
type ResourceType string

// ResourceType valid values.
const (
	ResourceTypeDocument           ResourceType = "Document"
	ResourceTypeStylesheet         ResourceType = "Stylesheet"
	ResourceTypeImage              ResourceType = "Image"
	ResourceTypeMedia              ResourceType = "Media"
	ResourceTypeFont               ResourceType = "Font"
	ResourceTypeScript             ResourceType = "Script"
	ResourceTypeTextTrack          ResourceType = "TextTrack"
	ResourceTypeXHR                ResourceType = "XHR"
	ResourceTypeFetch              ResourceType = "Fetch"
	ResourceTypeEventSource        ResourceType = "EventSource"
	ResourceTypeWebSocket          ResourceType = "WebSocket"
	ResourceTypeManifest           ResourceType = "Manifest"
	ResourceTypeSignedExchange     ResourceType = "SignedExchange"
	ResourceTypePing               ResourceType = "Ping"
	ResourceTypeCSPViolationReport ResourceType = "CSPViolationReport"
	ResourceTypePreflight          ResourceType = "Preflight"
	ResourceTypeOther              ResourceType = "Other"
)

// String returns the ResourceType value as a built-in string.
func (t ResourceType) String() string {
	return string(t)
}

// LoaderID data type. Unique loader identifier.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-LoaderId
type LoaderID string

// RequestID data type. Unique request identifier.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-RequestId
type RequestID string

// InterceptionID data type. Unique intercepted request identifier.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-InterceptionId
type InterceptionID string

// ErrorReason data type. Network level fetch failure reason.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-ErrorReason
type ErrorReason string

// ErrorReason valid values.
const (
	ErrorReasonFailed               ErrorReason = "Failed"
	ErrorReasonAborted              ErrorReason = "Aborted"
	ErrorReasonTimedOut             ErrorReason = "TimedOut"
	ErrorReasonAccessDenied         ErrorReason = "AccessDenied"
	ErrorReasonConnectionClosed     ErrorReason = "ConnectionClosed"
	ErrorReasonConnectionReset      ErrorReason = "ConnectionReset"
	ErrorReasonConnectionRefused    ErrorReason = "ConnectionRefused"
	ErrorReasonConnectionAborted    ErrorReason = "ConnectionAborted"
	ErrorReasonConnectionFailed     ErrorReason = "ConnectionFailed"
	ErrorReasonNameNotResolved      ErrorReason = "NameNotResolved"
	ErrorReasonInternetDisconnected ErrorReason = "InternetDisconnected"
	ErrorReasonAddressUnreachable   ErrorReason = "AddressUnreachable"
	ErrorReasonBlockedByClient      ErrorReason = "BlockedByClient"
	ErrorReasonBlockedByResponse    ErrorReason = "BlockedByResponse"
)

// String returns the ErrorReason value as a built-in string.
func (t ErrorReason) String() string {
	return string(t)
}

// TimeSinceEpoch data type. UTC time in seconds, counted from January 1, 1970.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-TimeSinceEpoch
type TimeSinceEpoch float64

// MonotonicTime data type. Monotonically increasing time in seconds since an arbitrary point in the past.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-MonotonicTime
type MonotonicTime float64

// Headers data type. Request / response headers as keys / values of JSON object.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-Headers
type Headers struct{}

// ConnectionType data type. The underlying connection technology that the browser is supposedly using.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-ConnectionType
type ConnectionType string

// ConnectionType valid values.
const (
	ConnectionTypeNone       ConnectionType = "none"
	ConnectionTypeCellular2g ConnectionType = "cellular2g"
	ConnectionTypeCellular3g ConnectionType = "cellular3g"
	ConnectionTypeCellular4g ConnectionType = "cellular4g"
	ConnectionTypeBluetooth  ConnectionType = "bluetooth"
	ConnectionTypeEthernet   ConnectionType = "ethernet"
	ConnectionTypeWifi       ConnectionType = "wifi"
	ConnectionTypeWimax      ConnectionType = "wimax"
	ConnectionTypeOther      ConnectionType = "other"
)

// String returns the ConnectionType value as a built-in string.
func (t ConnectionType) String() string {
	return string(t)
}

// CookieSameSite data type. Represents the cookie's 'SameSite' status:
// https://tools.ietf.org/html/draft-west-first-party-cookies
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-CookieSameSite
type CookieSameSite string

// CookieSameSite valid values.
const (
	CookieSameSiteStrict CookieSameSite = "Strict"
	CookieSameSiteLax    CookieSameSite = "Lax"
	CookieSameSiteNone   CookieSameSite = "None"
)

// String returns the CookieSameSite value as a built-in string.
func (t CookieSameSite) String() string {
	return string(t)
}

// CookiePriority data type. Represents the cookie's 'Priority' status:
// https://tools.ietf.org/html/draft-west-cookie-priority-00
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-CookiePriority
//
// This CDP type is experimental.
type CookiePriority string

// CookiePriority valid values.
const (
	CookiePriorityLow    CookiePriority = "Low"
	CookiePriorityMedium CookiePriority = "Medium"
	CookiePriorityHigh   CookiePriority = "High"
)

// String returns the CookiePriority value as a built-in string.
func (t CookiePriority) String() string {
	return string(t)
}

// CookieSourceScheme data type. Represents the source scheme of the origin that originally set the cookie.
// A value of "Unset" allows protocol clients to emulate legacy cookie scope for the scheme.
// This is a temporary ability and it will be removed in the future.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-CookieSourceScheme
//
// This CDP type is experimental.
type CookieSourceScheme string

// CookieSourceScheme valid values.
const (
	CookieSourceSchemeUnset     CookieSourceScheme = "Unset"
	CookieSourceSchemeNonSecure CookieSourceScheme = "NonSecure"
	CookieSourceSchemeSecure    CookieSourceScheme = "Secure"
)

// String returns the CookieSourceScheme value as a built-in string.
func (t CookieSourceScheme) String() string {
	return string(t)
}

// ResourceTiming data type. Timing information for the request.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-ResourceTiming
type ResourceTiming struct {
	// Timing's requestTime is a baseline in seconds, while the other numbers are ticks in
	// milliseconds relatively to this requestTime.
	RequestTime float64 `json:"requestTime"`
	// Started resolving proxy.
	ProxyStart float64 `json:"proxyStart"`
	// Finished resolving proxy.
	ProxyEnd float64 `json:"proxyEnd"`
	// Started DNS address resolve.
	DNSStart float64 `json:"dnsStart"`
	// Finished DNS address resolve.
	DNSEnd float64 `json:"dnsEnd"`
	// Started connecting to the remote host.
	ConnectStart float64 `json:"connectStart"`
	// Connected to the remote host.
	ConnectEnd float64 `json:"connectEnd"`
	// Started SSL handshake.
	SslStart float64 `json:"sslStart"`
	// Finished SSL handshake.
	SslEnd float64 `json:"sslEnd"`
	// Started running ServiceWorker.
	//
	// This CDP property is experimental.
	WorkerStart float64 `json:"workerStart"`
	// Finished Starting ServiceWorker.
	//
	// This CDP property is experimental.
	WorkerReady float64 `json:"workerReady"`
	// Started fetch event.
	//
	// This CDP property is experimental.
	WorkerFetchStart float64 `json:"workerFetchStart"`
	// Settled fetch event respondWith promise.
	//
	// This CDP property is experimental.
	WorkerRespondWithSettled float64 `json:"workerRespondWithSettled"`
	// Started sending request.
	SendStart float64 `json:"sendStart"`
	// Finished sending request.
	SendEnd float64 `json:"sendEnd"`
	// Time the server started pushing request.
	//
	// This CDP property is experimental.
	PushStart float64 `json:"pushStart"`
	// Time the server finished pushing request.
	//
	// This CDP property is experimental.
	PushEnd float64 `json:"pushEnd"`
	// Finished receiving response headers.
	ReceiveHeadersEnd float64 `json:"receiveHeadersEnd"`
}

// ResourcePriority data type. Loading priority of a resource request.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-ResourcePriority
type ResourcePriority string

// ResourcePriority valid values.
const (
	ResourcePriorityVeryLow  ResourcePriority = "VeryLow"
	ResourcePriorityLow      ResourcePriority = "Low"
	ResourcePriorityMedium   ResourcePriority = "Medium"
	ResourcePriorityHigh     ResourcePriority = "High"
	ResourcePriorityVeryHigh ResourcePriority = "VeryHigh"
)

// String returns the ResourcePriority value as a built-in string.
func (t ResourcePriority) String() string {
	return string(t)
}

// PostDataEntry data type. Post data entry for HTTP request
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-PostDataEntry
type PostDataEntry struct {
	Bytes string `json:"bytes,omitempty"`
}

// Request data type. HTTP request data.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-Request
type Request struct {
	// Request URL (without fragment).
	URL string `json:"url"`
	// Fragment of the requested URL starting with hash, if present.
	URLFragment string `json:"urlFragment,omitempty"`
	// HTTP request method.
	Method string `json:"method"`
	// HTTP request headers.
	Headers Headers `json:"headers"`
	// HTTP POST request data.
	PostData string `json:"postData,omitempty"`
	// True when the request has POST data. Note that postData might still be omitted when this flag is true when the data is too long.
	HasPostData bool `json:"hasPostData,omitempty"`
	// Request body elements. This will be converted from base64 to binary
	//
	// This CDP property is experimental.
	PostDataEntries []PostDataEntry `json:"postDataEntries,omitempty"`
	// The mixed content type of the request.
	MixedContentType *security.MixedContentType `json:"mixedContentType,omitempty"`
	// Priority of the resource request at the time request is sent.
	InitialPriority ResourcePriority `json:"initialPriority"`
	// The referrer policy of the request, as defined in https://www.w3.org/TR/referrer-policy/
	ReferrerPolicy string `json:"referrerPolicy"`
	// Whether is loaded via link preload.
	IsLinkPreload bool `json:"isLinkPreload,omitempty"`
	// Set for requests when the TrustToken API is used. Contains the parameters
	// passed by the developer (e.g. via "fetch") as understood by the backend.
	//
	// This CDP property is experimental.
	TrustTokenParams *TrustTokenParams `json:"trustTokenParams,omitempty"`
	// True if this resource request is considered to be the 'same site' as the
	// request correspondinfg to the main frame.
	//
	// This CDP property is experimental.
	IsSameSite bool `json:"isSameSite,omitempty"`
}

// SignedCertificateTimestamp data type. Details of a signed certificate timestamp (SCT).
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-SignedCertificateTimestamp
type SignedCertificateTimestamp struct {
	// Validation status.
	Status string `json:"status"`
	// Origin.
	Origin string `json:"origin"`
	// Log name / description.
	LogDescription string `json:"logDescription"`
	// Log ID.
	LogID string `json:"logId"`
	// Issuance date.
	Timestamp float64 `json:"timestamp"`
	// Hash algorithm.
	HashAlgorithm string `json:"hashAlgorithm"`
	// Signature algorithm.
	SignatureAlgorithm string `json:"signatureAlgorithm"`
	// Signature data.
	SignatureData string `json:"signatureData"`
}

// SecurityDetails data type. Security details about a request.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-SecurityDetails
type SecurityDetails struct {
	// Protocol name (e.g. "TLS 1.2" or "QUIC").
	Protocol string `json:"protocol"`
	// Key Exchange used by the connection, or the empty string if not applicable.
	KeyExchange string `json:"keyExchange"`
	// (EC)DH group used by the connection, if applicable.
	KeyExchangeGroup string `json:"keyExchangeGroup,omitempty"`
	// Cipher name.
	Cipher string `json:"cipher"`
	// TLS MAC. Note that AEAD ciphers do not have separate MACs.
	Mac string `json:"mac,omitempty"`
	// Certificate ID value.
	CertificateID int64 `json:"certificateId"`
	// Certificate subject name.
	SubjectName string `json:"subjectName"`
	// Subject Alternative Name (SAN) DNS names and IP addresses.
	SanList []string `json:"sanList"`
	// Name of the issuing CA.
	Issuer string `json:"issuer"`
	// Certificate valid from date.
	ValidFrom float64 `json:"validFrom"`
	// Certificate valid to (expiration) date
	ValidTo float64 `json:"validTo"`
	// List of signed certificate timestamps (SCTs).
	SignedCertificateTimestampList []SignedCertificateTimestamp `json:"signedCertificateTimestampList"`
	// Whether the request complied with Certificate Transparency policy
	CertificateTransparencyCompliance CertificateTransparencyCompliance `json:"certificateTransparencyCompliance"`
}

// CertificateTransparencyCompliance data type. Whether the request complied with Certificate Transparency policy.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-CertificateTransparencyCompliance
type CertificateTransparencyCompliance string

// CertificateTransparencyCompliance valid values.
const (
	CertificateTransparencyComplianceUnknown      CertificateTransparencyCompliance = "unknown"
	CertificateTransparencyComplianceNotCompliant CertificateTransparencyCompliance = "not-compliant"
	CertificateTransparencyComplianceCompliant    CertificateTransparencyCompliance = "compliant"
)

// String returns the CertificateTransparencyCompliance value as a built-in string.
func (t CertificateTransparencyCompliance) String() string {
	return string(t)
}

// BlockedReason data type. The reason why request was blocked.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-BlockedReason
type BlockedReason string

// BlockedReason valid values.
const (
	BlockedReasonOther                                             BlockedReason = "other"
	BlockedReasonCsp                                               BlockedReason = "csp"
	BlockedReasonMixedContent                                      BlockedReason = "mixed-content"
	BlockedReasonOrigin                                            BlockedReason = "origin"
	BlockedReasonInspector                                         BlockedReason = "inspector"
	BlockedReasonSubresourceFilter                                 BlockedReason = "subresource-filter"
	BlockedReasonContentType                                       BlockedReason = "content-type"
	BlockedReasonCoepFrameResourceNeedsCoepHeader                  BlockedReason = "coep-frame-resource-needs-coep-header"
	BlockedReasonCoopSandboxedIframeCannotNavigateToCoopPage       BlockedReason = "coop-sandboxed-iframe-cannot-navigate-to-coop-page"
	BlockedReasonCorpNotSameOrigin                                 BlockedReason = "corp-not-same-origin"
	BlockedReasonCorpNotSameOriginAfterDefaultedToSameOriginByCoep BlockedReason = "corp-not-same-origin-after-defaulted-to-same-origin-by-coep"
	BlockedReasonCorpNotSameSite                                   BlockedReason = "corp-not-same-site"
)

// String returns the BlockedReason value as a built-in string.
func (t BlockedReason) String() string {
	return string(t)
}

// CorsError data type. The reason why request was blocked.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-CorsError
type CorsError string

// CorsError valid values.
const (
	CorsErrorDisallowedByMode                     CorsError = "DisallowedByMode"
	CorsErrorInvalidResponse                      CorsError = "InvalidResponse"
	CorsErrorWildcardOriginNotAllowed             CorsError = "WildcardOriginNotAllowed"
	CorsErrorMissingAllowOriginHeader             CorsError = "MissingAllowOriginHeader"
	CorsErrorMultipleAllowOriginValues            CorsError = "MultipleAllowOriginValues"
	CorsErrorInvalidAllowOriginValue              CorsError = "InvalidAllowOriginValue"
	CorsErrorAllowOriginMismatch                  CorsError = "AllowOriginMismatch"
	CorsErrorInvalidAllowCredentials              CorsError = "InvalidAllowCredentials"
	CorsErrorCorsDisabledScheme                   CorsError = "CorsDisabledScheme"
	CorsErrorPreflightInvalidStatus               CorsError = "PreflightInvalidStatus"
	CorsErrorPreflightDisallowedRedirect          CorsError = "PreflightDisallowedRedirect"
	CorsErrorPreflightWildcardOriginNotAllowed    CorsError = "PreflightWildcardOriginNotAllowed"
	CorsErrorPreflightMissingAllowOriginHeader    CorsError = "PreflightMissingAllowOriginHeader"
	CorsErrorPreflightMultipleAllowOriginValues   CorsError = "PreflightMultipleAllowOriginValues"
	CorsErrorPreflightInvalidAllowOriginValue     CorsError = "PreflightInvalidAllowOriginValue"
	CorsErrorPreflightAllowOriginMismatch         CorsError = "PreflightAllowOriginMismatch"
	CorsErrorPreflightInvalidAllowCredentials     CorsError = "PreflightInvalidAllowCredentials"
	CorsErrorPreflightMissingAllowExternal        CorsError = "PreflightMissingAllowExternal"
	CorsErrorPreflightInvalidAllowExternal        CorsError = "PreflightInvalidAllowExternal"
	CorsErrorInvalidAllowMethodsPreflightResponse CorsError = "InvalidAllowMethodsPreflightResponse"
	CorsErrorInvalidAllowHeadersPreflightResponse CorsError = "InvalidAllowHeadersPreflightResponse"
	CorsErrorMethodDisallowedByPreflightResponse  CorsError = "MethodDisallowedByPreflightResponse"
	CorsErrorHeaderDisallowedByPreflightResponse  CorsError = "HeaderDisallowedByPreflightResponse"
	CorsErrorRedirectContainsCredentials          CorsError = "RedirectContainsCredentials"
	CorsErrorInsecurePrivateNetwork               CorsError = "InsecurePrivateNetwork"
	CorsErrorInvalidPrivateNetworkAccess          CorsError = "InvalidPrivateNetworkAccess"
	CorsErrorUnexpectedPrivateNetworkAccess       CorsError = "UnexpectedPrivateNetworkAccess"
	CorsErrorNoCorsRedirectModeNotFollow          CorsError = "NoCorsRedirectModeNotFollow"
)

// String returns the CorsError value as a built-in string.
func (t CorsError) String() string {
	return string(t)
}

// CorsErrorStatus data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-CorsErrorStatus
type CorsErrorStatus struct {
	CorsError       CorsError `json:"corsError"`
	FailedParameter string    `json:"failedParameter"`
}

// ServiceWorkerResponseSource data type. Source of serviceworker response.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-ServiceWorkerResponseSource
type ServiceWorkerResponseSource string

// ServiceWorkerResponseSource valid values.
const (
	ServiceWorkerResponseSourceCacheStorage ServiceWorkerResponseSource = "cache-storage"
	ServiceWorkerResponseSourceHTTPCache    ServiceWorkerResponseSource = "http-cache"
	ServiceWorkerResponseSourceFallbackCode ServiceWorkerResponseSource = "fallback-code"
	ServiceWorkerResponseSourceNetwork      ServiceWorkerResponseSource = "network"
)

// String returns the ServiceWorkerResponseSource value as a built-in string.
func (t ServiceWorkerResponseSource) String() string {
	return string(t)
}

// TrustTokenParams data type. Determines what type of Trust Token operation is executed and
// depending on the type, some additional parameters. The values
// are specified in third_party/blink/renderer/core/fetch/trust_token.idl.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-TrustTokenParams
//
// This CDP type is experimental.
type TrustTokenParams struct {
	Type TrustTokenOperationType `json:"type"`
	// Only set for "token-redemption" type and determine whether
	// to request a fresh SRR or use a still valid cached SRR.
	RefreshPolicy string `json:"refreshPolicy"`
	// Origins of issuers from whom to request tokens or redemption
	// records.
	Issuers []string `json:"issuers,omitempty"`
}

// TrustTokenOperationType data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-TrustTokenOperationType
//
// This CDP type is experimental.
type TrustTokenOperationType string

// TrustTokenOperationType valid values.
const (
	TrustTokenOperationTypeIssuance   TrustTokenOperationType = "Issuance"
	TrustTokenOperationTypeRedemption TrustTokenOperationType = "Redemption"
	TrustTokenOperationTypeSigning    TrustTokenOperationType = "Signing"
)

// String returns the TrustTokenOperationType value as a built-in string.
func (t TrustTokenOperationType) String() string {
	return string(t)
}

// Response data type. HTTP response data.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-Response
type Response struct {
	// Response URL. This URL can be different from CachedResource.url in case of redirect.
	URL string `json:"url"`
	// HTTP response status code.
	Status int64 `json:"status"`
	// HTTP response status text.
	StatusText string `json:"statusText"`
	// HTTP response headers.
	Headers Headers `json:"headers"`
	// HTTP response headers text. This has been replaced by the headers in Network.responseReceivedExtraInfo.
	//
	// This CDP property is deprecated.
	HeadersText string `json:"headersText,omitempty"`
	// Resource mimeType as determined by the browser.
	MimeType string `json:"mimeType"`
	// Refined HTTP request headers that were actually transmitted over the network.
	RequestHeaders *Headers `json:"requestHeaders,omitempty"`
	// HTTP request headers text. This has been replaced by the headers in Network.requestWillBeSentExtraInfo.
	//
	// This CDP property is deprecated.
	RequestHeadersText string `json:"requestHeadersText,omitempty"`
	// Specifies whether physical connection was actually reused for this request.
	ConnectionReused bool `json:"connectionReused"`
	// Physical connection id that was actually used for this request.
	ConnectionID float64 `json:"connectionId"`
	// Remote IP address.
	RemoteIPAddress string `json:"remoteIPAddress,omitempty"`
	// Remote port.
	RemotePort int64 `json:"remotePort,omitempty"`
	// Specifies that the request was served from the disk cache.
	FromDiskCache bool `json:"fromDiskCache,omitempty"`
	// Specifies that the request was served from the ServiceWorker.
	FromServiceWorker bool `json:"fromServiceWorker,omitempty"`
	// Specifies that the request was served from the prefetch cache.
	FromPrefetchCache bool `json:"fromPrefetchCache,omitempty"`
	// Total number of bytes received for this request so far.
	EncodedDataLength float64 `json:"encodedDataLength"`
	// Timing information for the given request.
	Timing *ResourceTiming `json:"timing,omitempty"`
	// Response source of response from ServiceWorker.
	ServiceWorkerResponseSource *ServiceWorkerResponseSource `json:"serviceWorkerResponseSource,omitempty"`
	// The time at which the returned response was generated.
	ResponseTime float64 `json:"responseTime,omitempty"`
	// Cache Storage Cache Name.
	CacheStorageCacheName string `json:"cacheStorageCacheName,omitempty"`
	// Protocol used to fetch this request.
	Protocol string `json:"protocol,omitempty"`
	// Security state of the request resource.
	SecurityState security.State `json:"securityState"`
	// Security details for the request.
	SecurityDetails *SecurityDetails `json:"securityDetails,omitempty"`
}

// WebSocketRequest data type. WebSocket request data.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-WebSocketRequest
type WebSocketRequest struct {
	// HTTP request headers.
	Headers Headers `json:"headers"`
}

// WebSocketResponse data type. WebSocket response data.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-WebSocketResponse
type WebSocketResponse struct {
	// HTTP response status code.
	Status int64 `json:"status"`
	// HTTP response status text.
	StatusText string `json:"statusText"`
	// HTTP response headers.
	Headers Headers `json:"headers"`
	// HTTP response headers text.
	HeadersText string `json:"headersText,omitempty"`
	// HTTP request headers.
	RequestHeaders *Headers `json:"requestHeaders,omitempty"`
	// HTTP request headers text.
	RequestHeadersText string `json:"requestHeadersText,omitempty"`
}

// WebSocketFrame data type. WebSocket message data. This represents an entire WebSocket message, not just a fragmented frame as the name suggests.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-WebSocketFrame
type WebSocketFrame struct {
	// WebSocket message opcode.
	Opcode float64 `json:"opcode"`
	// WebSocket message mask.
	Mask bool `json:"mask"`
	// WebSocket message payload data.
	// If the opcode is 1, this is a text message and payloadData is a UTF-8 string.
	// If the opcode isn't 1, then payloadData is a base64 encoded string representing binary data.
	PayloadData string `json:"payloadData"`
}

// CachedResource data type. Information about the cached resource.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-CachedResource
type CachedResource struct {
	// Resource URL. This is the url of the original network request.
	URL string `json:"url"`
	// Type of this resource.
	Type ResourceType `json:"type"`
	// Cached response data.
	Response *Response `json:"response,omitempty"`
	// Cached response body size.
	BodySize float64 `json:"bodySize"`
}

// Initiator data type. Information about the request initiator.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-Initiator
type Initiator struct {
	// Type of this initiator.
	Type string `json:"type"`
	// Initiator JavaScript stack trace, set for Script only.
	Stack *runtime.StackTrace `json:"stack,omitempty"`
	// Initiator URL, set for Parser type or for Script type (when script is importing module) or for SignedExchange type.
	URL string `json:"url,omitempty"`
	// Initiator line number, set for Parser type or for Script type (when script is importing
	// module) (0-based).
	LineNumber float64 `json:"lineNumber,omitempty"`
	// Initiator column number, set for Parser type or for Script type (when script is importing
	// module) (0-based).
	ColumnNumber float64 `json:"columnNumber,omitempty"`
	// Set if another request triggered this request (e.g. preflight).
	RequestID string `json:"requestId,omitempty"`
}

// Cookie object
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-Cookie
type Cookie struct {
	// Cookie name.
	Name string `json:"name"`
	// Cookie value.
	Value string `json:"value"`
	// Cookie domain.
	Domain string `json:"domain"`
	// Cookie path.
	Path string `json:"path"`
	// Cookie expiration date as the number of seconds since the UNIX epoch.
	Expires float64 `json:"expires"`
	// Cookie size.
	Size int64 `json:"size"`
	// True if cookie is http-only.
	HTTPOnly bool `json:"httpOnly"`
	// True if cookie is secure.
	Secure bool `json:"secure"`
	// True in case of session cookie.
	Session bool `json:"session"`
	// Cookie SameSite type.
	SameSite *CookieSameSite `json:"sameSite,omitempty"`
	// Cookie Priority
	//
	// This CDP property is experimental.
	Priority CookiePriority `json:"priority"`
	// True if cookie is SameParty.
	//
	// This CDP property is experimental.
	SameParty bool `json:"sameParty"`
	// Cookie source scheme type.
	//
	// This CDP property is experimental.
	SourceScheme CookieSourceScheme `json:"sourceScheme"`
	// Cookie source port. Valid values are {-1, [1, 65535]}, -1 indicates an unspecified port.
	// An unspecified port value allows protocol clients to emulate legacy cookie scope for the port.
	// This is a temporary ability and it will be removed in the future.
	//
	// This CDP property is experimental.
	SourcePort int64 `json:"sourcePort"`
}

// SetCookieBlockedReason data type. Types of reasons why a cookie may not be stored from a response.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-SetCookieBlockedReason
//
// This CDP type is experimental.
type SetCookieBlockedReason string

// SetCookieBlockedReason valid values.
const (
	SetCookieBlockedReasonSecureOnly                               SetCookieBlockedReason = "SecureOnly"
	SetCookieBlockedReasonSameSiteStrict                           SetCookieBlockedReason = "SameSiteStrict"
	SetCookieBlockedReasonSameSiteLax                              SetCookieBlockedReason = "SameSiteLax"
	SetCookieBlockedReasonSameSiteUnspecifiedTreatedAsLax          SetCookieBlockedReason = "SameSiteUnspecifiedTreatedAsLax"
	SetCookieBlockedReasonSameSiteNoneInsecure                     SetCookieBlockedReason = "SameSiteNoneInsecure"
	SetCookieBlockedReasonUserPreferences                          SetCookieBlockedReason = "UserPreferences"
	SetCookieBlockedReasonSyntaxError                              SetCookieBlockedReason = "SyntaxError"
	SetCookieBlockedReasonSchemeNotSupported                       SetCookieBlockedReason = "SchemeNotSupported"
	SetCookieBlockedReasonOverwriteSecure                          SetCookieBlockedReason = "OverwriteSecure"
	SetCookieBlockedReasonInvalidDomain                            SetCookieBlockedReason = "InvalidDomain"
	SetCookieBlockedReasonInvalidPrefix                            SetCookieBlockedReason = "InvalidPrefix"
	SetCookieBlockedReasonUnknownError                             SetCookieBlockedReason = "UnknownError"
	SetCookieBlockedReasonSchemefulSameSiteStrict                  SetCookieBlockedReason = "SchemefulSameSiteStrict"
	SetCookieBlockedReasonSchemefulSameSiteLax                     SetCookieBlockedReason = "SchemefulSameSiteLax"
	SetCookieBlockedReasonSchemefulSameSiteUnspecifiedTreatedAsLax SetCookieBlockedReason = "SchemefulSameSiteUnspecifiedTreatedAsLax"
	SetCookieBlockedReasonSamePartyFromCrossPartyContext           SetCookieBlockedReason = "SamePartyFromCrossPartyContext"
	SetCookieBlockedReasonSamePartyConflictsWithOtherAttributes    SetCookieBlockedReason = "SamePartyConflictsWithOtherAttributes"
	SetCookieBlockedReasonNameValuePairExceedsMaxSize              SetCookieBlockedReason = "NameValuePairExceedsMaxSize"
)

// String returns the SetCookieBlockedReason value as a built-in string.
func (t SetCookieBlockedReason) String() string {
	return string(t)
}

// CookieBlockedReason data type. Types of reasons why a cookie may not be sent with a request.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-CookieBlockedReason
//
// This CDP type is experimental.
type CookieBlockedReason string

// CookieBlockedReason valid values.
const (
	CookieBlockedReasonSecureOnly                               CookieBlockedReason = "SecureOnly"
	CookieBlockedReasonNotOnPath                                CookieBlockedReason = "NotOnPath"
	CookieBlockedReasonDomainMismatch                           CookieBlockedReason = "DomainMismatch"
	CookieBlockedReasonSameSiteStrict                           CookieBlockedReason = "SameSiteStrict"
	CookieBlockedReasonSameSiteLax                              CookieBlockedReason = "SameSiteLax"
	CookieBlockedReasonSameSiteUnspecifiedTreatedAsLax          CookieBlockedReason = "SameSiteUnspecifiedTreatedAsLax"
	CookieBlockedReasonSameSiteNoneInsecure                     CookieBlockedReason = "SameSiteNoneInsecure"
	CookieBlockedReasonUserPreferences                          CookieBlockedReason = "UserPreferences"
	CookieBlockedReasonUnknownError                             CookieBlockedReason = "UnknownError"
	CookieBlockedReasonSchemefulSameSiteStrict                  CookieBlockedReason = "SchemefulSameSiteStrict"
	CookieBlockedReasonSchemefulSameSiteLax                     CookieBlockedReason = "SchemefulSameSiteLax"
	CookieBlockedReasonSchemefulSameSiteUnspecifiedTreatedAsLax CookieBlockedReason = "SchemefulSameSiteUnspecifiedTreatedAsLax"
	CookieBlockedReasonSamePartyFromCrossPartyContext           CookieBlockedReason = "SamePartyFromCrossPartyContext"
	CookieBlockedReasonNameValuePairExceedsMaxSize              CookieBlockedReason = "NameValuePairExceedsMaxSize"
)

// String returns the CookieBlockedReason value as a built-in string.
func (t CookieBlockedReason) String() string {
	return string(t)
}

// BlockedSetCookieWithReason data type. A cookie which was not stored from a response with the corresponding reason.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-BlockedSetCookieWithReason
//
// This CDP type is experimental.
type BlockedSetCookieWithReason struct {
	// The reason(s) this cookie was blocked.
	BlockedReasons []SetCookieBlockedReason `json:"blockedReasons"`
	// The string representing this individual cookie as it would appear in the header.
	// This is not the entire "cookie" or "set-cookie" header which could have multiple cookies.
	CookieLine string `json:"cookieLine"`
	// The cookie object which represents the cookie which was not stored. It is optional because
	// sometimes complete cookie information is not available, such as in the case of parsing
	// errors.
	Cookie *Cookie `json:"cookie,omitempty"`
}

// BlockedCookieWithReason data type. A cookie with was not sent with a request with the corresponding reason.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-BlockedCookieWithReason
//
// This CDP type is experimental.
type BlockedCookieWithReason struct {
	// The reason(s) the cookie was blocked.
	BlockedReasons []CookieBlockedReason `json:"blockedReasons"`
	// The cookie object representing the cookie which was not sent.
	Cookie Cookie `json:"cookie"`
}

// CookieParam data type. Cookie parameter object
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-CookieParam
type CookieParam struct {
	// Cookie name.
	Name string `json:"name"`
	// Cookie value.
	Value string `json:"value"`
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
	HTTPOnly bool `json:"httpOnly,omitempty"`
	// Cookie SameSite type.
	SameSite *CookieSameSite `json:"sameSite,omitempty"`
	// Cookie expiration date, session cookie if not set
	Expires float64 `json:"expires,omitempty"`
	// Cookie Priority.
	//
	// This CDP property is experimental.
	Priority *CookiePriority `json:"priority,omitempty"`
	// True if cookie is SameParty.
	//
	// This CDP property is experimental.
	SameParty bool `json:"sameParty,omitempty"`
	// Cookie source scheme type.
	//
	// This CDP property is experimental.
	SourceScheme *CookieSourceScheme `json:"sourceScheme,omitempty"`
	// Cookie source port. Valid values are {-1, [1, 65535]}, -1 indicates an unspecified port.
	// An unspecified port value allows protocol clients to emulate legacy cookie scope for the port.
	// This is a temporary ability and it will be removed in the future.
	//
	// This CDP property is experimental.
	SourcePort int64 `json:"sourcePort,omitempty"`
}

// AuthChallenge data type. Authorization challenge for HTTP status code 401 or 407.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-AuthChallenge
//
// This CDP type is experimental.
type AuthChallenge struct {
	// Source of the authentication challenge.
	Source string `json:"source,omitempty"`
	// Origin of the challenger.
	Origin string `json:"origin"`
	// The authentication scheme used, such as basic or digest
	Scheme string `json:"scheme"`
	// The realm of the challenge. May be empty.
	Realm string `json:"realm"`
}

// AuthChallengeResponse data type. Response to an AuthChallenge.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-AuthChallengeResponse
//
// This CDP type is experimental.
type AuthChallengeResponse struct {
	// The decision on what to do in response to the authorization challenge.  Default means
	// deferring to the default behavior of the net stack, which will likely either the Cancel
	// authentication or display a popup dialog box.
	Response string `json:"response"`
	// The username to provide, possibly empty. Should only be set if response is
	// ProvideCredentials.
	Username string `json:"username,omitempty"`
	// The password to provide, possibly empty. Should only be set if response is
	// ProvideCredentials.
	Password string `json:"password,omitempty"`
}

// InterceptionStage data type. Stages of the interception to begin intercepting. Request will intercept before the request is
// sent. Response will intercept after the response is received.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-InterceptionStage
//
// This CDP type is experimental.
type InterceptionStage string

// InterceptionStage valid values.
const (
	InterceptionStageRequest         InterceptionStage = "Request"
	InterceptionStageHeadersReceived InterceptionStage = "HeadersReceived"
)

// String returns the InterceptionStage value as a built-in string.
func (t InterceptionStage) String() string {
	return string(t)
}

// RequestPattern data type. Request pattern for interception.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-RequestPattern
//
// This CDP type is experimental.
type RequestPattern struct {
	// Wildcards (`'*'` -> zero or more, `'?'` -> exactly one) are allowed. Escape character is
	// backslash. Omitting is equivalent to `"*"`.
	URLPattern string `json:"urlPattern,omitempty"`
	// If set, only requests for matching resource types will be intercepted.
	ResourceType *ResourceType `json:"resourceType,omitempty"`
	// Stage at which to begin intercepting requests. Default is Request.
	InterceptionStage *InterceptionStage `json:"interceptionStage,omitempty"`
}

// SignedExchangeSignature data type. Information about a signed exchange signature.
// https://wicg.github.io/webpackage/draft-yasskin-httpbis-origin-signed-exchanges-impl.html#rfc.section.3.1
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-SignedExchangeSignature
//
// This CDP type is experimental.
type SignedExchangeSignature struct {
	// Signed exchange signature label.
	Label string `json:"label"`
	// The hex string of signed exchange signature.
	Signature string `json:"signature"`
	// Signed exchange signature integrity.
	Integrity string `json:"integrity"`
	// Signed exchange signature cert Url.
	CertURL string `json:"certUrl,omitempty"`
	// The hex string of signed exchange signature cert sha256.
	CertSha256 string `json:"certSha256,omitempty"`
	// Signed exchange signature validity Url.
	ValidityURL string `json:"validityUrl"`
	// Signed exchange signature date.
	Date int64 `json:"date"`
	// Signed exchange signature expires.
	Expires int64 `json:"expires"`
	// The encoded certificates.
	Certificates []string `json:"certificates,omitempty"`
}

// SignedExchangeHeader data type. Information about a signed exchange header.
// https://wicg.github.io/webpackage/draft-yasskin-httpbis-origin-signed-exchanges-impl.html#cbor-representation
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-SignedExchangeHeader
//
// This CDP type is experimental.
type SignedExchangeHeader struct {
	// Signed exchange request URL.
	RequestURL string `json:"requestUrl"`
	// Signed exchange response code.
	ResponseCode int64 `json:"responseCode"`
	// Signed exchange response headers.
	ResponseHeaders Headers `json:"responseHeaders"`
	// Signed exchange response signature.
	Signatures []SignedExchangeSignature `json:"signatures"`
	// Signed exchange header integrity hash in the form of "sha256-<base64-hash-value>".
	HeaderIntegrity string `json:"headerIntegrity"`
}

// SignedExchangeErrorField data type. Field type for a signed exchange related error.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-SignedExchangeErrorField
//
// This CDP type is experimental.
type SignedExchangeErrorField string

// SignedExchangeErrorField valid values.
const (
	SignedExchangeErrorFieldSignatureSig         SignedExchangeErrorField = "signatureSig"
	SignedExchangeErrorFieldSignatureIntegrity   SignedExchangeErrorField = "signatureIntegrity"
	SignedExchangeErrorFieldSignatureCertURL     SignedExchangeErrorField = "signatureCertUrl"
	SignedExchangeErrorFieldSignatureCertSha256  SignedExchangeErrorField = "signatureCertSha256"
	SignedExchangeErrorFieldSignatureValidityURL SignedExchangeErrorField = "signatureValidityUrl"
	SignedExchangeErrorFieldSignatureTimestamps  SignedExchangeErrorField = "signatureTimestamps"
)

// String returns the SignedExchangeErrorField value as a built-in string.
func (t SignedExchangeErrorField) String() string {
	return string(t)
}

// SignedExchangeError data type. Information about a signed exchange response.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-SignedExchangeError
//
// This CDP type is experimental.
type SignedExchangeError struct {
	// Error message.
	Message string `json:"message"`
	// The index of the signature which caused the error.
	SignatureIndex int64 `json:"signatureIndex,omitempty"`
	// The field which caused the error.
	ErrorField *SignedExchangeErrorField `json:"errorField,omitempty"`
}

// SignedExchangeInfo data type. Information about a signed exchange response.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-SignedExchangeInfo
//
// This CDP type is experimental.
type SignedExchangeInfo struct {
	// The outer response of signed HTTP exchange which was received from network.
	OuterResponse Response `json:"outerResponse"`
	// Information about the signed exchange header.
	Header *SignedExchangeHeader `json:"header,omitempty"`
	// Security details for the signed exchange header.
	SecurityDetails *SecurityDetails `json:"securityDetails,omitempty"`
	// Errors occurred while handling the signed exchagne.
	Errors []SignedExchangeError `json:"errors,omitempty"`
}

// ContentEncoding data type. List of content encodings supported by the backend.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-ContentEncoding
//
// This CDP type is experimental.
type ContentEncoding string

// ContentEncoding valid values.
const (
	ContentEncodingDeflate ContentEncoding = "deflate"
	ContentEncodingGzip    ContentEncoding = "gzip"
	ContentEncodingBr      ContentEncoding = "br"
)

// String returns the ContentEncoding value as a built-in string.
func (t ContentEncoding) String() string {
	return string(t)
}

// PrivateNetworkRequestPolicy data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-PrivateNetworkRequestPolicy
//
// This CDP type is experimental.
type PrivateNetworkRequestPolicy string

// PrivateNetworkRequestPolicy valid values.
const (
	PrivateNetworkRequestPolicyAllow                          PrivateNetworkRequestPolicy = "Allow"
	PrivateNetworkRequestPolicyBlockFromInsecureToMorePrivate PrivateNetworkRequestPolicy = "BlockFromInsecureToMorePrivate"
	PrivateNetworkRequestPolicyWarnFromInsecureToMorePrivate  PrivateNetworkRequestPolicy = "WarnFromInsecureToMorePrivate"
	PrivateNetworkRequestPolicyPreflightBlock                 PrivateNetworkRequestPolicy = "PreflightBlock"
	PrivateNetworkRequestPolicyPreflightWarn                  PrivateNetworkRequestPolicy = "PreflightWarn"
)

// String returns the PrivateNetworkRequestPolicy value as a built-in string.
func (t PrivateNetworkRequestPolicy) String() string {
	return string(t)
}

// IPAddressSpace data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-IPAddressSpace
//
// This CDP type is experimental.
type IPAddressSpace string

// IPAddressSpace valid values.
const (
	IPAddressSpaceLocal   IPAddressSpace = "Local"
	IPAddressSpacePrivate IPAddressSpace = "Private"
	IPAddressSpacePublic  IPAddressSpace = "Public"
	IPAddressSpaceUnknown IPAddressSpace = "Unknown"
)

// String returns the IPAddressSpace value as a built-in string.
func (t IPAddressSpace) String() string {
	return string(t)
}

// ConnectTiming data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-ConnectTiming
//
// This CDP type is experimental.
type ConnectTiming struct {
	// Timing's requestTime is a baseline in seconds, while the other numbers are ticks in
	// milliseconds relatively to this requestTime. Matches ResourceTiming's requestTime for
	// the same request (but not for redirected requests).
	RequestTime float64 `json:"requestTime"`
}

// ClientSecurityState data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-ClientSecurityState
//
// This CDP type is experimental.
type ClientSecurityState struct {
	InitiatorIsSecureContext    bool                        `json:"initiatorIsSecureContext"`
	InitiatorIPAddressSpace     IPAddressSpace              `json:"initiatorIPAddressSpace"`
	PrivateNetworkRequestPolicy PrivateNetworkRequestPolicy `json:"privateNetworkRequestPolicy"`
}

// CrossOriginOpenerPolicyValue data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-CrossOriginOpenerPolicyValue
//
// This CDP type is experimental.
type CrossOriginOpenerPolicyValue string

// CrossOriginOpenerPolicyValue valid values.
const (
	CrossOriginOpenerPolicyValueSameOrigin            CrossOriginOpenerPolicyValue = "SameOrigin"
	CrossOriginOpenerPolicyValueSameOriginAllowPopups CrossOriginOpenerPolicyValue = "SameOriginAllowPopups"
	CrossOriginOpenerPolicyValueUnsafeNone            CrossOriginOpenerPolicyValue = "UnsafeNone"
	CrossOriginOpenerPolicyValueSameOriginPlusCoep    CrossOriginOpenerPolicyValue = "SameOriginPlusCoep"
)

// String returns the CrossOriginOpenerPolicyValue value as a built-in string.
func (t CrossOriginOpenerPolicyValue) String() string {
	return string(t)
}

// CrossOriginOpenerPolicyStatus data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-CrossOriginOpenerPolicyStatus
//
// This CDP type is experimental.
type CrossOriginOpenerPolicyStatus struct {
	Value                       CrossOriginOpenerPolicyValue `json:"value"`
	ReportOnlyValue             CrossOriginOpenerPolicyValue `json:"reportOnlyValue"`
	ReportingEndpoint           string                       `json:"reportingEndpoint,omitempty"`
	ReportOnlyReportingEndpoint string                       `json:"reportOnlyReportingEndpoint,omitempty"`
}

// CrossOriginEmbedderPolicyValue data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-CrossOriginEmbedderPolicyValue
//
// This CDP type is experimental.
type CrossOriginEmbedderPolicyValue string

// CrossOriginEmbedderPolicyValue valid values.
const (
	CrossOriginEmbedderPolicyValueNone           CrossOriginEmbedderPolicyValue = "None"
	CrossOriginEmbedderPolicyValueCredentialless CrossOriginEmbedderPolicyValue = "Credentialless"
	CrossOriginEmbedderPolicyValueRequireCorp    CrossOriginEmbedderPolicyValue = "RequireCorp"
)

// String returns the CrossOriginEmbedderPolicyValue value as a built-in string.
func (t CrossOriginEmbedderPolicyValue) String() string {
	return string(t)
}

// CrossOriginEmbedderPolicyStatus data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-CrossOriginEmbedderPolicyStatus
//
// This CDP type is experimental.
type CrossOriginEmbedderPolicyStatus struct {
	Value                       CrossOriginEmbedderPolicyValue `json:"value"`
	ReportOnlyValue             CrossOriginEmbedderPolicyValue `json:"reportOnlyValue"`
	ReportingEndpoint           string                         `json:"reportingEndpoint,omitempty"`
	ReportOnlyReportingEndpoint string                         `json:"reportOnlyReportingEndpoint,omitempty"`
}

// SecurityIsolationStatus data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-SecurityIsolationStatus
//
// This CDP type is experimental.
type SecurityIsolationStatus struct {
	Coop *CrossOriginOpenerPolicyStatus   `json:"coop,omitempty"`
	Coep *CrossOriginEmbedderPolicyStatus `json:"coep,omitempty"`
}

// ReportStatus data type. The status of a Reporting API report.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-ReportStatus
//
// This CDP type is experimental.
type ReportStatus string

// ReportStatus valid values.
const (
	ReportStatusQueued           ReportStatus = "Queued"
	ReportStatusPending          ReportStatus = "Pending"
	ReportStatusMarkedForRemoval ReportStatus = "MarkedForRemoval"
	ReportStatusSuccess          ReportStatus = "Success"
)

// String returns the ReportStatus value as a built-in string.
func (t ReportStatus) String() string {
	return string(t)
}

// ReportID data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-ReportId
//
// This CDP type is experimental.
type ReportID string

// ReportingAPIReport data type. An object representing a report generated by the Reporting API.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-ReportingApiReport
//
// This CDP type is experimental.
type ReportingAPIReport struct {
	ID string `json:"id"`
	// The URL of the document that triggered the report.
	InitiatorURL string `json:"initiatorUrl"`
	// The name of the endpoint group that should be used to deliver the report.
	Destination string `json:"destination"`
	// The type of the report (specifies the set of data that is contained in the report body).
	Type string `json:"type"`
	// When the report was generated.
	Timestamp float64 `json:"timestamp"`
	// How many uploads deep the related request was.
	Depth int64 `json:"depth"`
	// The number of delivery attempts made so far, not including an active attempt.
	CompletedAttempts int64           `json:"completedAttempts"`
	Body              json.RawMessage `json:"body"`
	Status            ReportStatus    `json:"status"`
}

// LoadNetworkResourcePageResult data type. An object providing the result of a network resource load.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-LoadNetworkResourcePageResult
//
// This CDP type is experimental.
type LoadNetworkResourcePageResult struct {
	Success bool `json:"success"`
	// Optional values used for error reporting.
	NetError       float64 `json:"netError,omitempty"`
	NetErrorName   string  `json:"netErrorName,omitempty"`
	HTTPStatusCode float64 `json:"httpStatusCode,omitempty"`
	// If successful, one of the following two fields holds the result.
	Stream string `json:"stream,omitempty"`
	// Response headers.
	Headers *Headers `json:"headers,omitempty"`
}

// LoadNetworkResourceOptions data type. An options object that may be extended later to better support CORS,
// CORB and streaming.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-LoadNetworkResourceOptions
//
// This CDP type is experimental.
type LoadNetworkResourceOptions struct {
	DisableCache       bool `json:"disableCache"`
	IncludeCredentials bool `json:"includeCredentials"`
}
