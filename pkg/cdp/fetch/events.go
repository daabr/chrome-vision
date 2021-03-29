package fetch

import (
	"github.com/daabr/chrome-vision/pkg/cdp"
	"github.com/daabr/chrome-vision/pkg/cdp/network"
)

// Issued when the domain is enabled and the request URL matches the
// specified filter. The request is paused until the client responds
// with one of continueRequest, failRequest or fulfillRequest.
// The stage of the request can be determined by presence of responseErrorReason
// and responseStatusCode -- the request is at the response stage if either
// of these fields is present and in the request stage otherwise.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Fetch/#event-requestPaused
type RequestPaused struct {
	// Each request the page makes will have a unique id.
	RequestID RequestID `json:"requestId"`
	// The details of the request.
	Request network.Request `json:"request"`
	// The id of the frame that initiated the request.
	FrameID cdp.FrameID `json:"frameId"`
	// How the requested resource will be used.
	ResourceType network.ResourceType `json:"resourceType"`
	// Response error if intercepted at response stage.
	ResponseErrorReason *network.ErrorReason `json:"responseErrorReason,omitempty"`
	// Response code if intercepted at response stage.
	ResponseStatusCode int64 `json:"responseStatusCode,omitempty"`
	// Response headers if intercepted at the response stage.
	ResponseHeaders []HeaderEntry `json:"responseHeaders,omitempty"`
	// If the intercepted request had a corresponding Network.requestWillBeSent event fired for it,
	// then this networkId will be the same as the requestId present in the requestWillBeSent event.
	NetworkID *RequestID `json:"networkId,omitempty"`
}

// Issued when the domain is enabled with handleAuthRequests set to true.
// The request is paused until client responds with continueWithAuth.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Fetch/#event-authRequired
type AuthRequired struct {
	// Each request the page makes will have a unique id.
	RequestID RequestID `json:"requestId"`
	// The details of the request.
	Request network.Request `json:"request"`
	// The id of the frame that initiated the request.
	FrameID cdp.FrameID `json:"frameId"`
	// How the requested resource will be used.
	ResourceType network.ResourceType `json:"resourceType"`
	// Details of the Authorization Challenge encountered.
	// If this is set, client should respond with continueRequest that
	// contains AuthChallengeResponse.
	AuthChallenge AuthChallenge `json:"authChallenge"`
}
