package domdebugger

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
	"github.com/daabr/chrome-vision/pkg/devtools/runtime"
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
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-getEventListeners
func NewGetEventListeners(objectID runtime.RemoteObjectID) *GetEventListeners {
	return &GetEventListeners{
		ObjectID: objectID,
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

// GetEventListenersResult contains the browser's response
// to calling the GetEventListeners CDP command with Do().
type GetEventListenersResult struct {
	// Array of relevant listeners.
	Listeners []EventListener `json:"listeners"`
}

// Do sends the GetEventListeners CDP command to a browser,
// and returns the browser's response.
func (t *GetEventListeners) Do(ctx context.Context) (*GetEventListenersResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "DOMDebugger.getEventListeners", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetEventListeners CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetEventListeners) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "DOMDebugger.getEventListeners", b)
}

// ParseResponse parses the browser's response
// to the GetEventListeners CDP command.
func (t *GetEventListeners) ParseResponse(m *devtools.Message) (*GetEventListenersResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetEventListenersResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// RemoveDOMBreakpoint contains the parameters, and acts as
// a Go receiver, for the CDP command `removeDOMBreakpoint`.
//
// Removes DOM breakpoint that was set using `setDOMBreakpoint`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeDOMBreakpoint
type RemoveDOMBreakpoint struct {
	// Identifier of the node to remove breakpoint from.
	NodeID int64 `json:"nodeId"`
	// Type of the breakpoint to remove.
	Type DOMBreakpointType `json:"type"`
}

// NewRemoveDOMBreakpoint constructs a new RemoveDOMBreakpoint struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-removeDOMBreakpoint
func NewRemoveDOMBreakpoint(nodeID int64, t DOMBreakpointType) *RemoveDOMBreakpoint {
	return &RemoveDOMBreakpoint{
		NodeID: nodeID,
		Type:   t,
	}
}

// Do sends the RemoveDOMBreakpoint CDP command to a browser,
// and returns the browser's response.
func (t *RemoveDOMBreakpoint) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "DOMDebugger.removeDOMBreakpoint", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the RemoveDOMBreakpoint CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *RemoveDOMBreakpoint) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "DOMDebugger.removeDOMBreakpoint", b)
}

// ParseResponse parses the browser's response
// to the RemoveDOMBreakpoint CDP command.
func (t *RemoveDOMBreakpoint) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
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
// all (but only) the required parameters. Optional parameters
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
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "DOMDebugger.removeEventListenerBreakpoint", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the RemoveEventListenerBreakpoint CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *RemoveEventListenerBreakpoint) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "DOMDebugger.removeEventListenerBreakpoint", b)
}

// ParseResponse parses the browser's response
// to the RemoveEventListenerBreakpoint CDP command.
func (t *RemoveEventListenerBreakpoint) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
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
// all (but only) the required parameters. Optional parameters
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
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "DOMDebugger.removeInstrumentationBreakpoint", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the RemoveInstrumentationBreakpoint CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *RemoveInstrumentationBreakpoint) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "DOMDebugger.removeInstrumentationBreakpoint", b)
}

// ParseResponse parses the browser's response
// to the RemoveInstrumentationBreakpoint CDP command.
func (t *RemoveInstrumentationBreakpoint) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
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
// all (but only) the required parameters. Optional parameters
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
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "DOMDebugger.removeXHRBreakpoint", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the RemoveXHRBreakpoint CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *RemoveXHRBreakpoint) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "DOMDebugger.removeXHRBreakpoint", b)
}

// ParseResponse parses the browser's response
// to the RemoveXHRBreakpoint CDP command.
func (t *RemoveXHRBreakpoint) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
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
// all (but only) the required parameters. Optional parameters
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
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "DOMDebugger.setBreakOnCSPViolation", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetBreakOnCSPViolation CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetBreakOnCSPViolation) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "DOMDebugger.setBreakOnCSPViolation", b)
}

// ParseResponse parses the browser's response
// to the SetBreakOnCSPViolation CDP command.
func (t *SetBreakOnCSPViolation) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
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
	NodeID int64 `json:"nodeId"`
	// Type of the operation to stop upon.
	Type DOMBreakpointType `json:"type"`
}

// NewSetDOMBreakpoint constructs a new SetDOMBreakpoint struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#method-setDOMBreakpoint
func NewSetDOMBreakpoint(nodeID int64, t DOMBreakpointType) *SetDOMBreakpoint {
	return &SetDOMBreakpoint{
		NodeID: nodeID,
		Type:   t,
	}
}

// Do sends the SetDOMBreakpoint CDP command to a browser,
// and returns the browser's response.
func (t *SetDOMBreakpoint) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "DOMDebugger.setDOMBreakpoint", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetDOMBreakpoint CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetDOMBreakpoint) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "DOMDebugger.setDOMBreakpoint", b)
}

// ParseResponse parses the browser's response
// to the SetDOMBreakpoint CDP command.
func (t *SetDOMBreakpoint) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
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
// all (but only) the required parameters. Optional parameters
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
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "DOMDebugger.setEventListenerBreakpoint", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetEventListenerBreakpoint CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetEventListenerBreakpoint) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "DOMDebugger.setEventListenerBreakpoint", b)
}

// ParseResponse parses the browser's response
// to the SetEventListenerBreakpoint CDP command.
func (t *SetEventListenerBreakpoint) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
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
// all (but only) the required parameters. Optional parameters
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
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "DOMDebugger.setInstrumentationBreakpoint", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetInstrumentationBreakpoint CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetInstrumentationBreakpoint) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "DOMDebugger.setInstrumentationBreakpoint", b)
}

// ParseResponse parses the browser's response
// to the SetInstrumentationBreakpoint CDP command.
func (t *SetInstrumentationBreakpoint) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
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
// all (but only) the required parameters. Optional parameters
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
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "DOMDebugger.setXHRBreakpoint", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetXHRBreakpoint CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetXHRBreakpoint) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "DOMDebugger.setXHRBreakpoint", b)
}

// ParseResponse parses the browser's response
// to the SetXHRBreakpoint CDP command.
func (t *SetXHRBreakpoint) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}
