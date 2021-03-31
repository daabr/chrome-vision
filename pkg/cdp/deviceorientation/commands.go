package deviceorientation

import (
	"context"
	"fmt"
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
	fmt.Println(ctx)
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
	fmt.Println(ctx)
	return nil
}
