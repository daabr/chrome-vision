package domdebugger

import (
	"context"
	"fmt"

	"github.com/daabr/chrome-vision/pkg/cdp/dom"
	"github.com/daabr/chrome-vision/pkg/cdp/runtime"
)

// GetEventListeners contains the parameters, and acts as
// a Go receiver, for the CDP command `getEventListeners`.
//
// Returns event listeners of the given object.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-getEventListeners
type GetEventListeners struct {
	// Identifier of the object to return listeners for.
	ObjectID runtime.RemoteObjectID `json:"objectId"`
	// The maximum depth at which Node children should be retrieved, defaults to 1. Use -1 for the
	// entire subtree or provide an integer larger than 0.
	Depth int64 `json:"depth,omitempty"`
	// Whether or not iframes and shadow roots should be traversed when returning the subtree
	// (default is false). Reports listeners for all contexts if pierce is enabled.
	Pierce bool `json:"pierce,omitempty"`
}

// NewGetEventListeners constructs a new GetEventListeners struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-getEventListeners
func NewGetEventListeners(objectId runtime.RemoteObjectID) *GetEventListeners {
	return &GetEventListeners{
		ObjectID: objectId,
	}
}

// SetDepth adds or modifies the value of the optional
// parameter `depth` in the GetEventListeners CDP command.
//
// The maximum depth at which Node children should be retrieved, defaults to 1. Use -1 for the
// entire subtree or provide an integer larger than 0.
func (t *GetEventListeners) SetDepth(v int64) *GetEventListeners {
	t.Depth = v
	return t
}

// SetPierce adds or modifies the value of the optional
// parameter `pierce` in the GetEventListeners CDP command.
//
// Whether or not iframes and shadow roots should be traversed when returning the subtree
// (default is false). Reports listeners for all contexts if pierce is enabled.
func (t *GetEventListeners) SetPierce(v bool) *GetEventListeners {
	t.Pierce = v
	return t
}

// GetEventListenersResponse contains the browser's response
// to calling the GetEventListeners CDP command with Do().
type GetEventListenersResponse struct {
	// Array of relevant listeners.
	Listeners []EventListener `json:"listeners"`
}

// Do sends the GetEventListeners CDP command to a browser,
// and returns the browser's response.
func (t *GetEventListeners) Do(ctx context.Context) (*GetEventListenersResponse, error) {
	fmt.Println(ctx)
	return new(GetEventListenersResponse), nil
}

// RemoveDOMBreakpoint contains the parameters, and acts as
// a Go receiver, for the CDP command `removeDOMBreakpoint`.
//
// Removes DOM breakpoint that was set using `setDOMBreakpoint`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeDOMBreakpoint
type RemoveDOMBreakpoint struct {
	// Identifier of the node to remove breakpoint from.
	NodeID dom.NodeID `json:"nodeId"`
	// Type of the breakpoint to remove.
	Type DOMBreakpointType `json:"type"`
}

// NewRemoveDOMBreakpoint constructs a new RemoveDOMBreakpoint struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeDOMBreakpoint
func NewRemoveDOMBreakpoint(nodeId dom.NodeID, t DOMBreakpointType) *RemoveDOMBreakpoint {
	return &RemoveDOMBreakpoint{
		NodeID: nodeId,
		Type:   t,
	}
}

// Do sends the RemoveDOMBreakpoint CDP command to a browser,
// and returns the browser's response.
func (t *RemoveDOMBreakpoint) Do(ctx context.Context) error {
	fmt.Println(ctx)
	return nil
}

// RemoveEventListenerBreakpoint contains the parameters, and acts as
// a Go receiver, for the CDP command `removeEventListenerBreakpoint`.
//
// Removes breakpoint on particular DOM event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeEventListenerBreakpoint
type RemoveEventListenerBreakpoint struct {
	// Event name.
	EventName string `json:"eventName"`
	// EventTarget interface name.
	//
	// This CDP parameter is experimental.
	TargetName string `json:"targetName,omitempty"`
}

// NewRemoveEventListenerBreakpoint constructs a new RemoveEventListenerBreakpoint struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeEventListenerBreakpoint
func NewRemoveEventListenerBreakpoint(eventName string) *RemoveEventListenerBreakpoint {
	return &RemoveEventListenerBreakpoint{
		EventName: eventName,
	}
}

// SetTargetName adds or modifies the value of the optional
// parameter `targetName` in the RemoveEventListenerBreakpoint CDP command.
//
// EventTarget interface name.
//
// This CDP parameter is experimental.
func (t *RemoveEventListenerBreakpoint) SetTargetName(v string) *RemoveEventListenerBreakpoint {
	t.TargetName = v
	return t
}

// Do sends the RemoveEventListenerBreakpoint CDP command to a browser,
// and returns the browser's response.
func (t *RemoveEventListenerBreakpoint) Do(ctx context.Context) error {
	fmt.Println(ctx)
	return nil
}

// RemoveInstrumentationBreakpoint contains the parameters, and acts as
// a Go receiver, for the CDP command `removeInstrumentationBreakpoint`.
//
// Removes breakpoint on particular native event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeInstrumentationBreakpoint
//
// This CDP method is experimental.
type RemoveInstrumentationBreakpoint struct {
	// Instrumentation name to stop on.
	EventName string `json:"eventName"`
}

// NewRemoveInstrumentationBreakpoint constructs a new RemoveInstrumentationBreakpoint struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeInstrumentationBreakpoint
//
// This CDP method is experimental.
func NewRemoveInstrumentationBreakpoint(eventName string) *RemoveInstrumentationBreakpoint {
	return &RemoveInstrumentationBreakpoint{
		EventName: eventName,
	}
}

// Do sends the RemoveInstrumentationBreakpoint CDP command to a browser,
// and returns the browser's response.
func (t *RemoveInstrumentationBreakpoint) Do(ctx context.Context) error {
	fmt.Println(ctx)
	return nil
}

// RemoveXHRBreakpoint contains the parameters, and acts as
// a Go receiver, for the CDP command `removeXHRBreakpoint`.
//
// Removes breakpoint from XMLHttpRequest.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeXHRBreakpoint
type RemoveXHRBreakpoint struct {
	// Resource URL substring.
	URL string `json:"url"`
}

// NewRemoveXHRBreakpoint constructs a new RemoveXHRBreakpoint struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeXHRBreakpoint
func NewRemoveXHRBreakpoint(url string) *RemoveXHRBreakpoint {
	return &RemoveXHRBreakpoint{
		URL: url,
	}
}

// Do sends the RemoveXHRBreakpoint CDP command to a browser,
// and returns the browser's response.
func (t *RemoveXHRBreakpoint) Do(ctx context.Context) error {
	fmt.Println(ctx)
	return nil
}

// SetBreakOnCSPViolation contains the parameters, and acts as
// a Go receiver, for the CDP command `setBreakOnCSPViolation`.
//
// Sets breakpoint on particular CSP violations.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setBreakOnCSPViolation
//
// This CDP method is experimental.
type SetBreakOnCSPViolation struct {
	// CSP Violations to stop upon.
	ViolationTypes []CSPViolationType `json:"violationTypes"`
}

// NewSetBreakOnCSPViolation constructs a new SetBreakOnCSPViolation struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setBreakOnCSPViolation
//
// This CDP method is experimental.
func NewSetBreakOnCSPViolation(violationTypes []CSPViolationType) *SetBreakOnCSPViolation {
	return &SetBreakOnCSPViolation{
		ViolationTypes: violationTypes,
	}
}

// Do sends the SetBreakOnCSPViolation CDP command to a browser,
// and returns the browser's response.
func (t *SetBreakOnCSPViolation) Do(ctx context.Context) error {
	fmt.Println(ctx)
	return nil
}

// SetDOMBreakpoint contains the parameters, and acts as
// a Go receiver, for the CDP command `setDOMBreakpoint`.
//
// Sets breakpoint on particular operation with DOM.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setDOMBreakpoint
type SetDOMBreakpoint struct {
	// Identifier of the node to set breakpoint on.
	NodeID dom.NodeID `json:"nodeId"`
	// Type of the operation to stop upon.
	Type DOMBreakpointType `json:"type"`
}

// NewSetDOMBreakpoint constructs a new SetDOMBreakpoint struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setDOMBreakpoint
func NewSetDOMBreakpoint(nodeId dom.NodeID, t DOMBreakpointType) *SetDOMBreakpoint {
	return &SetDOMBreakpoint{
		NodeID: nodeId,
		Type:   t,
	}
}

// Do sends the SetDOMBreakpoint CDP command to a browser,
// and returns the browser's response.
func (t *SetDOMBreakpoint) Do(ctx context.Context) error {
	fmt.Println(ctx)
	return nil
}

// SetEventListenerBreakpoint contains the parameters, and acts as
// a Go receiver, for the CDP command `setEventListenerBreakpoint`.
//
// Sets breakpoint on particular DOM event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setEventListenerBreakpoint
type SetEventListenerBreakpoint struct {
	// DOM Event name to stop on (any DOM event will do).
	EventName string `json:"eventName"`
	// EventTarget interface name to stop on. If equal to `"*"` or not provided, will stop on any
	// EventTarget.
	//
	// This CDP parameter is experimental.
	TargetName string `json:"targetName,omitempty"`
}

// NewSetEventListenerBreakpoint constructs a new SetEventListenerBreakpoint struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setEventListenerBreakpoint
func NewSetEventListenerBreakpoint(eventName string) *SetEventListenerBreakpoint {
	return &SetEventListenerBreakpoint{
		EventName: eventName,
	}
}

// SetTargetName adds or modifies the value of the optional
// parameter `targetName` in the SetEventListenerBreakpoint CDP command.
//
// EventTarget interface name to stop on. If equal to `"*"` or not provided, will stop on any
// EventTarget.
//
// This CDP parameter is experimental.
func (t *SetEventListenerBreakpoint) SetTargetName(v string) *SetEventListenerBreakpoint {
	t.TargetName = v
	return t
}

// Do sends the SetEventListenerBreakpoint CDP command to a browser,
// and returns the browser's response.
func (t *SetEventListenerBreakpoint) Do(ctx context.Context) error {
	fmt.Println(ctx)
	return nil
}

// SetInstrumentationBreakpoint contains the parameters, and acts as
// a Go receiver, for the CDP command `setInstrumentationBreakpoint`.
//
// Sets breakpoint on particular native event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setInstrumentationBreakpoint
//
// This CDP method is experimental.
type SetInstrumentationBreakpoint struct {
	// Instrumentation name to stop on.
	EventName string `json:"eventName"`
}

// NewSetInstrumentationBreakpoint constructs a new SetInstrumentationBreakpoint struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setInstrumentationBreakpoint
//
// This CDP method is experimental.
func NewSetInstrumentationBreakpoint(eventName string) *SetInstrumentationBreakpoint {
	return &SetInstrumentationBreakpoint{
		EventName: eventName,
	}
}

// Do sends the SetInstrumentationBreakpoint CDP command to a browser,
// and returns the browser's response.
func (t *SetInstrumentationBreakpoint) Do(ctx context.Context) error {
	fmt.Println(ctx)
	return nil
}

// SetXHRBreakpoint contains the parameters, and acts as
// a Go receiver, for the CDP command `setXHRBreakpoint`.
//
// Sets breakpoint on XMLHttpRequest.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setXHRBreakpoint
type SetXHRBreakpoint struct {
	// Resource URL substring. All XHRs having this substring in the URL will get stopped upon.
	URL string `json:"url"`
}

// NewSetXHRBreakpoint constructs a new SetXHRBreakpoint struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setXHRBreakpoint
func NewSetXHRBreakpoint(url string) *SetXHRBreakpoint {
	return &SetXHRBreakpoint{
		URL: url,
	}
}

// Do sends the SetXHRBreakpoint CDP command to a browser,
// and returns the browser's response.
func (t *SetXHRBreakpoint) Do(ctx context.Context) error {
	fmt.Println(ctx)
	return nil
}
