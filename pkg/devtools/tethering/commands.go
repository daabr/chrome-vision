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
	response, err := devtools.Send(ctx, "Tethering.bind", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
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
	response, err := devtools.Send(ctx, "Tethering.unbind", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}
