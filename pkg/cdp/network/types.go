package network

import "github.com/daabr/chrome-vision/pkg/cdp/runtime"

// Resource type as it was perceived by the rendering engine.
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

// Unique loader identifier.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-LoaderId
type LoaderID string

// Unique request identifier.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-RequestId
type RequestID string

// Unique intercepted request identifier.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-InterceptionId
type InterceptionID string

// Network level fetch failure reason.
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

// UTC time in seconds, counted from January 1, 1970.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-TimeSinceEpoch
type TimeSinceEpoch float64

// Monotonically increasing time in seconds since an arbitrary point in the past.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-MonotonicTime
type MonotonicTime float64

// Request / response headers as keys / values of JSON object.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-Headers
type Headers struct{}

// The underlying connection technology that the browser is supposedly using.
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

// Represents the cookie's 'SameSite' status:
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

// Represents the cookie's 'Priority' status:
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

// Represents the source scheme of the origin that originally set the cookie.
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

// Timing information for the request.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-ResourceTiming
type ResourceTiming struct {
	// Timing's requestTime is a baseline in seconds, while the other numbers are ticks in
	// milliseconds relatively to this requestTime.
	RequestTime float64
	// Started resolving proxy.
	ProxyStart float64
	// Finished resolving proxy.
	ProxyEnd float64
	// Started DNS address resolve.
	DnsStart float64
	// Finished DNS address resolve.
	DnsEnd float64
	// Started connecting to the remote host.
	ConnectStart float64
	// Connected to the remote host.
	ConnectEnd float64
	// Started SSL handshake.
	SslStart float64
	// Finished SSL handshake.
	SslEnd float64
	// Started running ServiceWorker.
	//
	// This CDP property is experimental.
	WorkerStart float64
	// Finished Starting ServiceWorker.
	//
	// This CDP property is experimental.
	WorkerReady float64
	// Started fetch event.
	//
	// This CDP property is experimental.
	WorkerFetchStart float64
	// Settled fetch event respondWith promise.
	//
	// This CDP property is experimental.
	WorkerRespondWithSettled float64
	// Started sending request.
	SendStart float64
	// Finished sending request.
	SendEnd float64
	// Time the server started pushing request.
	//
	// This CDP property is experimental.
	PushStart float64
	// Time the server finished pushing request.
	//
	// This CDP property is experimental.
	PushEnd float64
	// Finished receiving response headers.
	ReceiveHeadersEnd float64
}

// Loading priority of a resource request.
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

// Post data entry for HTTP request
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-PostDataEntry
type PostDataEntry struct {
	Bytes string `json:"bytes,omitempty"`
}

// HTTP request data.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-Request
type Request struct {
	// Request URL (without fragment).
	URL string
	// Fragment of the requested URL starting with hash, if present.
	URLFragment string `json:"urlFragment,omitempty"`
	// HTTP request method.
	Method string
	// HTTP request headers.
	Headers Headers
	// HTTP POST request data.
	PostData string `json:"postData,omitempty"`
	// True when the request has POST data. Note that postData might still be omitted when this flag is true when the data is too long.
	HasPostData bool `json:"hasPostData,omitempty"`
	// Request body elements. This will be converted from base64 to binary
	//
	// This CDP property is experimental.
	PostDataEntries []PostDataEntry `json:"postDataEntries,omitempty"`
	// The mixed content type of the request.
	MixedContentType string `json:"mixedContentType,omitempty"`
	// Priority of the resource request at the time request is sent.
	InitialPriority string
	// The referrer policy of the request, as defined in https://www.w3.org/TR/referrer-policy/
	ReferrerPolicy string
	// Whether is loaded via link preload.
	IsLinkPreload bool `json:"isLinkPreload,omitempty"`
	// Set for requests when the TrustToken API is used. Contains the parameters
	// passed by the developer (e.g. via "fetch") as understood by the backend.
	//
	// This CDP property is experimental.
	TrustTokenParams *TrustTokenParams `json:"trustTokenParams,omitempty"`
}

// Details of a signed certificate timestamp (SCT).
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-SignedCertificateTimestamp
type SignedCertificateTimestamp struct {
	// Validation status.
	Status string
	// Origin.
	Origin string
	// Log name / description.
	LogDescription string
	// Log ID.
	LogID string
	// Issuance date.
	Timestamp float64
	// Hash algorithm.
	HashAlgorithm string
	// Signature algorithm.
	SignatureAlgorithm string
	// Signature data.
	SignatureData string
}

// Security details about a request.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-SecurityDetails
type SecurityDetails struct {
	// Protocol name (e.g. "TLS 1.2" or "QUIC").
	Protocol string
	// Key Exchange used by the connection, or the empty string if not applicable.
	KeyExchange string
	// (EC)DH group used by the connection, if applicable.
	KeyExchangeGroup string `json:"keyExchangeGroup,omitempty"`
	// Cipher name.
	Cipher string
	// TLS MAC. Note that AEAD ciphers do not have separate MACs.
	Mac string `json:"mac,omitempty"`
	// Certificate ID value.
	CertificateID int64
	// Certificate subject name.
	SubjectName string
	// Subject Alternative Name (SAN) DNS names and IP addresses.
	SanList []string
	// Name of the issuing CA.
	Issuer string
	// Certificate valid from date.
	ValidFrom float64
	// Certificate valid to (expiration) date
	ValidTo float64
	// List of signed certificate timestamps (SCTs).
	SignedCertificateTimestampList []SignedCertificateTimestamp
	// Whether the request complied with Certificate Transparency policy
	CertificateTransparencyCompliance string
}

// Whether the request complied with Certificate Transparency policy.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-CertificateTransparencyCompliance
type CertificateTransparencyCompliance string

// CertificateTransparencyCompliance valid values.
const (
	CertificateTransparencyComplianceUnknown      CertificateTransparencyCompliance = "unknown"
	CertificateTransparencyComplianceNotCompliant CertificateTransparencyCompliance = "not-compliant"
	CertificateTransparencyComplianceCompliant    CertificateTransparencyCompliance = "compliant"
)

// The reason why request was blocked.
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

// The reason why request was blocked.
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
)

// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-CorsErrorStatus
type CorsErrorStatus struct {
	CorsError       string
	FailedParameter string
}

// Source of serviceworker response.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-ServiceWorkerResponseSource
type ServiceWorkerResponseSource string

// ServiceWorkerResponseSource valid values.
const (
	ServiceWorkerResponseSourceCacheStorage ServiceWorkerResponseSource = "cache-storage"
	ServiceWorkerResponseSourceHttpCache    ServiceWorkerResponseSource = "http-cache"
	ServiceWorkerResponseSourceFallbackCode ServiceWorkerResponseSource = "fallback-code"
	ServiceWorkerResponseSourceNetwork      ServiceWorkerResponseSource = "network"
)

// Determines what type of Trust Token operation is executed and
// depending on the type, some additional parameters. The values
// are specified in third_party/blink/renderer/core/fetch/trust_token.idl.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-TrustTokenParams
//
// This CDP type is experimental.
type TrustTokenParams struct {
	Type string
	// Only set for "token-redemption" type and determine whether
	// to request a fresh SRR or use a still valid cached SRR.
	RefreshPolicy string
	// Origins of issuers from whom to request tokens or redemption
	// records.
	Issuers []string `json:"issuers,omitempty"`
}

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

// HTTP response data.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-Response
type Response struct {
	// Response URL. This URL can be different from CachedResource.url in case of redirect.
	URL string
	// HTTP response status code.
	Status int64
	// HTTP response status text.
	StatusText string
	// HTTP response headers.
	Headers Headers
	// HTTP response headers text.
	HeadersText string `json:"headersText,omitempty"`
	// Resource mimeType as determined by the browser.
	MimeType string
	// Refined HTTP request headers that were actually transmitted over the network.
	RequestHeaders *Headers `json:"requestHeaders,omitempty"`
	// HTTP request headers text.
	RequestHeadersText string `json:"requestHeadersText,omitempty"`
	// Specifies whether physical connection was actually reused for this request.
	ConnectionReused bool
	// Physical connection id that was actually used for this request.
	ConnectionID float64
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
	EncodedDataLength float64
	// Timing information for the given request.
	Timing *ResourceTiming `json:"timing,omitempty"`
	// Response source of response from ServiceWorker.
	ServiceWorkerResponseSource string `json:"serviceWorkerResponseSource,omitempty"`
	// The time at which the returned response was generated.
	ResponseTime float64 `json:"responseTime,omitempty"`
	// Cache Storage Cache Name.
	CacheStorageCacheName string `json:"cacheStorageCacheName,omitempty"`
	// Protocol used to fetch this request.
	Protocol string `json:"protocol,omitempty"`
	// Security state of the request resource.
	SecurityState string
	// Security details for the request.
	SecurityDetails *SecurityDetails `json:"securityDetails,omitempty"`
}

// WebSocket request data.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-WebSocketRequest
type WebSocketRequest struct {
	// HTTP request headers.
	Headers Headers
}

// WebSocket response data.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-WebSocketResponse
type WebSocketResponse struct {
	// HTTP response status code.
	Status int64
	// HTTP response status text.
	StatusText string
	// HTTP response headers.
	Headers Headers
	// HTTP response headers text.
	HeadersText string `json:"headersText,omitempty"`
	// HTTP request headers.
	RequestHeaders *Headers `json:"requestHeaders,omitempty"`
	// HTTP request headers text.
	RequestHeadersText string `json:"requestHeadersText,omitempty"`
}

// WebSocket message data. This represents an entire WebSocket message, not just a fragmented frame as the name suggests.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-WebSocketFrame
type WebSocketFrame struct {
	// WebSocket message opcode.
	Opcode float64
	// WebSocket message mask.
	Mask bool
	// WebSocket message payload data.
	// If the opcode is 1, this is a text message and payloadData is a UTF-8 string.
	// If the opcode isn't 1, then payloadData is a base64 encoded string representing binary data.
	PayloadData string
}

// Information about the cached resource.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-CachedResource
type CachedResource struct {
	// Resource URL. This is the url of the original network request.
	URL string
	// Type of this resource.
	Type string
	// Cached response data.
	Response *Response `json:"response,omitempty"`
	// Cached response body size.
	BodySize float64
}

// Information about the request initiator.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-Initiator
type Initiator struct {
	// Type of this initiator.
	Type string
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
	Name string
	// Cookie value.
	Value string
	// Cookie domain.
	Domain string
	// Cookie path.
	Path string
	// Cookie expiration date as the number of seconds since the UNIX epoch.
	Expires float64
	// Cookie size.
	Size int64
	// True if cookie is http-only.
	HttpOnly bool
	// True if cookie is secure.
	Secure bool
	// True in case of session cookie.
	Session bool
	// Cookie SameSite type.
	SameSite string `json:"sameSite,omitempty"`
	// Cookie Priority
	//
	// This CDP property is experimental.
	Priority string
	// True if cookie is SameParty.
	//
	// This CDP property is experimental.
	SameParty bool
	// Cookie source scheme type.
	//
	// This CDP property is experimental.
	SourceScheme string
	// Cookie source port. Valid values are {-1, [1, 65535]}, -1 indicates an unspecified port.
	// An unspecified port value allows protocol clients to emulate legacy cookie scope for the port.
	// This is a temporary ability and it will be removed in the future.
	//
	// This CDP property is experimental.
	SourcePort int64
}

// Types of reasons why a cookie may not be stored from a response.
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
)

// Types of reasons why a cookie may not be sent with a request.
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
)

// A cookie which was not stored from a response with the corresponding reason.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-BlockedSetCookieWithReason
//
// This CDP type is experimental.
type BlockedSetCookieWithReason struct {
	// The reason(s) this cookie was blocked.
	BlockedReasons []SetCookieBlockedReason
	// The string representing this individual cookie as it would appear in the header.
	// This is not the entire "cookie" or "set-cookie" header which could have multiple cookies.
	CookieLine string
	// The cookie object which represents the cookie which was not stored. It is optional because
	// sometimes complete cookie information is not available, such as in the case of parsing
	// errors.
	Cookie *Cookie `json:"cookie,omitempty"`
}

// A cookie with was not sent with a request with the corresponding reason.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-BlockedCookieWithReason
//
// This CDP type is experimental.
type BlockedCookieWithReason struct {
	// The reason(s) the cookie was blocked.
	BlockedReasons []CookieBlockedReason
	// The cookie object representing the cookie which was not sent.
	Cookie Cookie
}

// Cookie parameter object
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-CookieParam
type CookieParam struct {
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
	// Cookie Priority.
	//
	// This CDP property is experimental.
	Priority string `json:"priority,omitempty"`
	// True if cookie is SameParty.
	//
	// This CDP property is experimental.
	SameParty bool `json:"sameParty,omitempty"`
	// Cookie source scheme type.
	//
	// This CDP property is experimental.
	SourceScheme string `json:"sourceScheme,omitempty"`
	// Cookie source port. Valid values are {-1, [1, 65535]}, -1 indicates an unspecified port.
	// An unspecified port value allows protocol clients to emulate legacy cookie scope for the port.
	// This is a temporary ability and it will be removed in the future.
	//
	// This CDP property is experimental.
	SourcePort int64 `json:"sourcePort,omitempty"`
}

// Authorization challenge for HTTP status code 401 or 407.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-AuthChallenge
//
// This CDP type is experimental.
type AuthChallenge struct {
	// Source of the authentication challenge.
	Source string `json:"source,omitempty"`
	// Origin of the challenger.
	Origin string
	// The authentication scheme used, such as basic or digest
	Scheme string
	// The realm of the challenge. May be empty.
	Realm string
}

// Response to an AuthChallenge.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-AuthChallengeResponse
//
// This CDP type is experimental.
type AuthChallengeResponse struct {
	// The decision on what to do in response to the authorization challenge.  Default means
	// deferring to the default behavior of the net stack, which will likely either the Cancel
	// authentication or display a popup dialog box.
	Response string
	// The username to provide, possibly empty. Should only be set if response is
	// ProvideCredentials.
	Username string `json:"username,omitempty"`
	// The password to provide, possibly empty. Should only be set if response is
	// ProvideCredentials.
	Password string `json:"password,omitempty"`
}

// Stages of the interception to begin intercepting. Request will intercept before the request is
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

// Request pattern for interception.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-RequestPattern
//
// This CDP type is experimental.
type RequestPattern struct {
	// Wildcards ('*' -> zero or more, '?' -> exactly one) are allowed. Escape character is
	// backslash. Omitting is equivalent to "*".
	URLPattern string `json:"urlPattern,omitempty"`
	// If set, only requests for matching resource types will be intercepted.
	ResourceType string `json:"resourceType,omitempty"`
	// Stage at wich to begin intercepting requests. Default is Request.
	InterceptionStage string `json:"interceptionStage,omitempty"`
}

// Information about a signed exchange signature.
// https://wicg.github.io/webpackage/draft-yasskin-httpbis-origin-signed-exchanges-impl.html#rfc.section.3.1
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-SignedExchangeSignature
//
// This CDP type is experimental.
type SignedExchangeSignature struct {
	// Signed exchange signature label.
	Label string
	// The hex string of signed exchange signature.
	Signature string
	// Signed exchange signature integrity.
	Integrity string
	// Signed exchange signature cert Url.
	CertURL string `json:"certUrl,omitempty"`
	// The hex string of signed exchange signature cert sha256.
	CertSha256 string `json:"certSha256,omitempty"`
	// Signed exchange signature validity Url.
	ValidityURL string
	// Signed exchange signature date.
	Date int64
	// Signed exchange signature expires.
	Expires int64
	// The encoded certificates.
	Certificates []string `json:"certificates,omitempty"`
}

// Information about a signed exchange header.
// https://wicg.github.io/webpackage/draft-yasskin-httpbis-origin-signed-exchanges-impl.html#cbor-representation
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-SignedExchangeHeader
//
// This CDP type is experimental.
type SignedExchangeHeader struct {
	// Signed exchange request URL.
	RequestURL string
	// Signed exchange response code.
	ResponseCode int64
	// Signed exchange response headers.
	ResponseHeaders Headers
	// Signed exchange response signature.
	Signatures []SignedExchangeSignature
	// Signed exchange header integrity hash in the form of "sha256-<base64-hash-value>".
	HeaderIntegrity string
}

// Field type for a signed exchange related error.
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

// Information about a signed exchange response.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-SignedExchangeError
//
// This CDP type is experimental.
type SignedExchangeError struct {
	// Error message.
	Message string
	// The index of the signature which caused the error.
	SignatureIndex int64 `json:"signatureIndex,omitempty"`
	// The field which caused the error.
	ErrorField string `json:"errorField,omitempty"`
}

// Information about a signed exchange response.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-SignedExchangeInfo
//
// This CDP type is experimental.
type SignedExchangeInfo struct {
	// The outer response of signed HTTP exchange which was received from network.
	OuterResponse Response
	// Information about the signed exchange header.
	Header *SignedExchangeHeader `json:"header,omitempty"`
	// Security details for the signed exchange header.
	SecurityDetails *SecurityDetails `json:"securityDetails,omitempty"`
	// Errors occurred while handling the signed exchagne.
	Errors []SignedExchangeError `json:"errors,omitempty"`
}

// List of content encodings supported by the backend.
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

// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-PrivateNetworkRequestPolicy
//
// This CDP type is experimental.
type PrivateNetworkRequestPolicy string

// PrivateNetworkRequestPolicy valid values.
const (
	PrivateNetworkRequestPolicyAllow                          PrivateNetworkRequestPolicy = "Allow"
	PrivateNetworkRequestPolicyBlockFromInsecureToMorePrivate PrivateNetworkRequestPolicy = "BlockFromInsecureToMorePrivate"
	PrivateNetworkRequestPolicyWarnFromInsecureToMorePrivate  PrivateNetworkRequestPolicy = "WarnFromInsecureToMorePrivate"
)

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

// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-ClientSecurityState
//
// This CDP type is experimental.
type ClientSecurityState struct {
	InitiatorIsSecureContext    bool
	InitiatorIPAddressSpace     string
	PrivateNetworkRequestPolicy string
}

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

// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-CrossOriginOpenerPolicyStatus
//
// This CDP type is experimental.
type CrossOriginOpenerPolicyStatus struct {
	Value                       string
	ReportOnlyValue             string
	ReportingEndpoint           string `json:"reportingEndpoint,omitempty"`
	ReportOnlyReportingEndpoint string `json:"reportOnlyReportingEndpoint,omitempty"`
}

// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-CrossOriginEmbedderPolicyValue
//
// This CDP type is experimental.
type CrossOriginEmbedderPolicyValue string

// CrossOriginEmbedderPolicyValue valid values.
const (
	CrossOriginEmbedderPolicyValueNone                 CrossOriginEmbedderPolicyValue = "None"
	CrossOriginEmbedderPolicyValueCorsOrCredentialless CrossOriginEmbedderPolicyValue = "CorsOrCredentialless"
	CrossOriginEmbedderPolicyValueRequireCorp          CrossOriginEmbedderPolicyValue = "RequireCorp"
)

// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-CrossOriginEmbedderPolicyStatus
//
// This CDP type is experimental.
type CrossOriginEmbedderPolicyStatus struct {
	Value                       string
	ReportOnlyValue             string
	ReportingEndpoint           string `json:"reportingEndpoint,omitempty"`
	ReportOnlyReportingEndpoint string `json:"reportOnlyReportingEndpoint,omitempty"`
}

// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-SecurityIsolationStatus
//
// This CDP type is experimental.
type SecurityIsolationStatus struct {
	Coop *CrossOriginOpenerPolicyStatus   `json:"coop,omitempty"`
	Coep *CrossOriginEmbedderPolicyStatus `json:"coep,omitempty"`
}

// An object providing the result of a network resource load.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-LoadNetworkResourcePageResult
//
// This CDP type is experimental.
type LoadNetworkResourcePageResult struct {
	Success bool
	// Optional values used for error reporting.
	NetError       float64 `json:"netError,omitempty"`
	NetErrorName   string  `json:"netErrorName,omitempty"`
	HttpStatusCode float64 `json:"httpStatusCode,omitempty"`
	// If successful, one of the following two fields holds the result.
	Stream string `json:"stream,omitempty"`
	// Response headers.
	Headers *Headers `json:"headers,omitempty"`
}

// An options object that may be extended later to better support CORS,
// CORB and streaming.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-LoadNetworkResourceOptions
//
// This CDP type is experimental.
type LoadNetworkResourceOptions struct {
	DisableCache       bool
	IncludeCredentials bool
}
