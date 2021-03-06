package tethering

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
)

// Bind contains the parameters, and acts as
// a Go receiver, for the CDP command `bind`.
//
// Request browser port binding.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Tethering/#method-bind
type Bind struct {
	// Port number to bind.
	Port int64 `json:"port"`
}

// NewBind constructs a new Bind struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Tethering/#method-bind
func NewBind(port int64) *Bind {
	return &Bind{
		Port: port,
	}
}

// Do sends the Bind CDP command to a browser,
// and returns the browser's response.
func (t *Bind) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Tethering.bind", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the Bind CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *Bind) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Tethering.bind", b)
}

// ParseResponse parses the browser's response
// to the Bind CDP command.
func (t *Bind) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// Unbind contains the parameters, and acts as
// a Go receiver, for the CDP command `unbind`.
//
// Request browser port unbinding.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Tethering/#method-unbind
type Unbind struct {
	// Port number to unbind.
	Port int64 `json:"port"`
}

// NewUnbind constructs a new Unbind struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Tethering/#method-unbind
func NewUnbind(port int64) *Unbind {
	return &Unbind{
		Port: port,
	}
}

// Do sends the Unbind CDP command to a browser,
// and returns the browser's response.
func (t *Unbind) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Tethering.unbind", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the Unbind CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *Unbind) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Tethering.unbind", b)
}

// ParseResponse parses the browser's response
// to the Unbind CDP command.
func (t *Unbind) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}
