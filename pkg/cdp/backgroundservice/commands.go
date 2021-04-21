package backgroundservice

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/cdp"
)

// StartObserving contains the parameters, and acts as
// a Go receiver, for the CDP command `startObserving`.
//
// Enables event updates for the service.
//
// https://chromedevtools.github.io/devtools-protocol/tot/BackgroundService/#method-startObserving
type StartObserving struct {
	Service string `json:"service"`
}

// NewStartObserving constructs a new StartObserving struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/BackgroundService/#method-startObserving
func NewStartObserving(service string) *StartObserving {
	return &StartObserving{
		Service: service,
	}
}

// Do sends the StartObserving CDP command to a browser,
// and returns the browser's response.
func (t *StartObserving) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "BackgroundService.startObserving", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// StopObserving contains the parameters, and acts as
// a Go receiver, for the CDP command `stopObserving`.
//
// Disables event updates for the service.
//
// https://chromedevtools.github.io/devtools-protocol/tot/BackgroundService/#method-stopObserving
type StopObserving struct {
	Service string `json:"service"`
}

// NewStopObserving constructs a new StopObserving struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/BackgroundService/#method-stopObserving
func NewStopObserving(service string) *StopObserving {
	return &StopObserving{
		Service: service,
	}
}

// Do sends the StopObserving CDP command to a browser,
// and returns the browser's response.
func (t *StopObserving) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "BackgroundService.stopObserving", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetRecording contains the parameters, and acts as
// a Go receiver, for the CDP command `setRecording`.
//
// Set the recording state for the service.
//
// https://chromedevtools.github.io/devtools-protocol/tot/BackgroundService/#method-setRecording
type SetRecording struct {
	ShouldRecord bool   `json:"shouldRecord"`
	Service      string `json:"service"`
}

// NewSetRecording constructs a new SetRecording struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/BackgroundService/#method-setRecording
func NewSetRecording(shouldRecord bool, service string) *SetRecording {
	return &SetRecording{
		ShouldRecord: shouldRecord,
		Service:      service,
	}
}

// Do sends the SetRecording CDP command to a browser,
// and returns the browser's response.
func (t *SetRecording) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "BackgroundService.setRecording", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// ClearEvents contains the parameters, and acts as
// a Go receiver, for the CDP command `clearEvents`.
//
// Clears all stored data for the service.
//
// https://chromedevtools.github.io/devtools-protocol/tot/BackgroundService/#method-clearEvents
type ClearEvents struct {
	Service string `json:"service"`
}

// NewClearEvents constructs a new ClearEvents struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/BackgroundService/#method-clearEvents
func NewClearEvents(service string) *ClearEvents {
	return &ClearEvents{
		Service: service,
	}
}

// Do sends the ClearEvents CDP command to a browser,
// and returns the browser's response.
func (t *ClearEvents) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "BackgroundService.clearEvents", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}
