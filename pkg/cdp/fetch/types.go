package fetch

import "github.com/daabr/chrome-vision/pkg/cdp/network"

// RequestID data type. Unique request identifier.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Fetch/#type-RequestId
type RequestID string

// RequestStage data type. Stages of the request to handle. Request will intercept before the request is
// sent. Response will intercept after the response is received (but before response
// body is received.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Fetch/#type-RequestStage
type RequestStage string

// RequestStage valid values.
const (
	RequestStageRequest  RequestStage = "Request"
	RequestStageResponse RequestStage = "Response"
)

// RequestPattern data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Fetch/#type-RequestPattern
type RequestPattern struct {
	// Wildcards (`'*'` -> zero or more, `'?'` -> exactly one) are allowed. Escape character is
	// backslash. Omitting is equivalent to `"*"`.
	URLPattern string `json:"urlPattern,omitempty"`
	// If set, only requests for matching resource types will be intercepted.
	ResourceType *network.ResourceType `json:"resourceType,omitempty"`
	// Stage at which to begin intercepting requests. Default is Request.
	RequestStage *RequestStage `json:"requestStage,omitempty"`
}

// HeaderEntry data type. Response HTTP header entry
//
// https://chromedevtools.github.io/devtools-protocol/tot/Fetch/#type-HeaderEntry
type HeaderEntry struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// AuthChallenge data type. Authorization challenge for HTTP status code 401 or 407.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Fetch/#type-AuthChallenge
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
// https://chromedevtools.github.io/devtools-protocol/tot/Fetch/#type-AuthChallengeResponse
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
