package eventbreakpoints

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
)

// SetInstrumentationBreakpoint contains the parameters, and acts as
// a Go receiver, for the CDP command `setInstrumentationBreakpoint`.
//
// Sets breakpoint on particular native event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/EventBreakpoints/#method-setInstrumentationBreakpoint
type SetInstrumentationBreakpoint struct {
	// Instrumentation name to stop on.
	EventName string `json:"eventName"`
}

// NewSetInstrumentationBreakpoint constructs a new SetInstrumentationBreakpoint struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/EventBreakpoints/#method-setInstrumentationBreakpoint
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
	m, err := devtools.SendAndWait(ctx, "EventBreakpoints.setInstrumentationBreakpoint", b)
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
	return devtools.Send(ctx, "EventBreakpoints.setInstrumentationBreakpoint", b)
}

// ParseResponse parses the browser's response
// to the SetInstrumentationBreakpoint CDP command.
func (t *SetInstrumentationBreakpoint) ParseResponse(m *devtools.Message) error {
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
// https://chromedevtools.github.io/devtools-protocol/tot/EventBreakpoints/#method-removeInstrumentationBreakpoint
type RemoveInstrumentationBreakpoint struct {
	// Instrumentation name to stop on.
	EventName string `json:"eventName"`
}

// NewRemoveInstrumentationBreakpoint constructs a new RemoveInstrumentationBreakpoint struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/EventBreakpoints/#method-removeInstrumentationBreakpoint
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
	m, err := devtools.SendAndWait(ctx, "EventBreakpoints.removeInstrumentationBreakpoint", b)
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
	return devtools.Send(ctx, "EventBreakpoints.removeInstrumentationBreakpoint", b)
}

// ParseResponse parses the browser's response
// to the RemoveInstrumentationBreakpoint CDP command.
func (t *RemoveInstrumentationBreakpoint) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}
