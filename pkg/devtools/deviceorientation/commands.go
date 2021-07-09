package deviceorientation

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
)

// ClearDeviceOrientationOverride contains the parameters, and acts as
// a Go receiver, for the CDP command `clearDeviceOrientationOverride`.
//
// Clears the overridden Device Orientation.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DeviceOrientation/#method-clearDeviceOrientationOverride
type ClearDeviceOrientationOverride struct{}

// NewClearDeviceOrientationOverride constructs a new ClearDeviceOrientationOverride struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DeviceOrientation/#method-clearDeviceOrientationOverride
func NewClearDeviceOrientationOverride() *ClearDeviceOrientationOverride {
	return &ClearDeviceOrientationOverride{}
}

// Do sends the ClearDeviceOrientationOverride CDP command to a browser,
// and returns the browser's response.
func (t *ClearDeviceOrientationOverride) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "DeviceOrientation.clearDeviceOrientationOverride", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the ClearDeviceOrientationOverride CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *ClearDeviceOrientationOverride) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "DeviceOrientation.clearDeviceOrientationOverride", nil)
}

// ParseResponse parses the browser's response
// to the ClearDeviceOrientationOverride CDP command.
func (t *ClearDeviceOrientationOverride) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetDeviceOrientationOverride contains the parameters, and acts as
// a Go receiver, for the CDP command `setDeviceOrientationOverride`.
//
// Overrides the Device Orientation.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DeviceOrientation/#method-setDeviceOrientationOverride
type SetDeviceOrientationOverride struct {
	// Mock alpha
	Alpha float64 `json:"alpha"`
	// Mock beta
	Beta float64 `json:"beta"`
	// Mock gamma
	Gamma float64 `json:"gamma"`
}

// NewSetDeviceOrientationOverride constructs a new SetDeviceOrientationOverride struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DeviceOrientation/#method-setDeviceOrientationOverride
func NewSetDeviceOrientationOverride(alpha float64, beta float64, gamma float64) *SetDeviceOrientationOverride {
	return &SetDeviceOrientationOverride{
		Alpha: alpha,
		Beta:  beta,
		Gamma: gamma,
	}
}

// Do sends the SetDeviceOrientationOverride CDP command to a browser,
// and returns the browser's response.
func (t *SetDeviceOrientationOverride) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "DeviceOrientation.setDeviceOrientationOverride", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetDeviceOrientationOverride CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetDeviceOrientationOverride) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "DeviceOrientation.setDeviceOrientationOverride", b)
}

// ParseResponse parses the browser's response
// to the SetDeviceOrientationOverride CDP command.
func (t *SetDeviceOrientationOverride) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}
