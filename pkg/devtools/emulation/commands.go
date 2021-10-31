package emulation

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
	"github.com/daabr/chrome-vision/pkg/devtools/dom"
	"github.com/daabr/chrome-vision/pkg/devtools/page"
)

// CanEmulate contains the parameters, and acts as
// a Go receiver, for the CDP command `canEmulate`.
//
// Tells whether emulation is supported.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-canEmulate
type CanEmulate struct{}

// NewCanEmulate constructs a new CanEmulate struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-canEmulate
func NewCanEmulate() *CanEmulate {
	return &CanEmulate{}
}

// CanEmulateResult contains the browser's response
// to calling the CanEmulate CDP command with Do().
type CanEmulateResult struct {
	// True if emulation is supported.
	Result bool `json:"result"`
}

// Do sends the CanEmulate CDP command to a browser,
// and returns the browser's response.
func (t *CanEmulate) Do(ctx context.Context) (*CanEmulateResult, error) {
	m, err := devtools.SendAndWait(ctx, "Emulation.canEmulate", nil)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the CanEmulate CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *CanEmulate) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Emulation.canEmulate", nil)
}

// ParseResponse parses the browser's response
// to the CanEmulate CDP command.
func (t *CanEmulate) ParseResponse(m *devtools.Message) (*CanEmulateResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &CanEmulateResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ClearDeviceMetricsOverride contains the parameters, and acts as
// a Go receiver, for the CDP command `clearDeviceMetricsOverride`.
//
// Clears the overridden device metrics.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-clearDeviceMetricsOverride
type ClearDeviceMetricsOverride struct{}

// NewClearDeviceMetricsOverride constructs a new ClearDeviceMetricsOverride struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-clearDeviceMetricsOverride
func NewClearDeviceMetricsOverride() *ClearDeviceMetricsOverride {
	return &ClearDeviceMetricsOverride{}
}

// Do sends the ClearDeviceMetricsOverride CDP command to a browser,
// and returns the browser's response.
func (t *ClearDeviceMetricsOverride) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Emulation.clearDeviceMetricsOverride", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the ClearDeviceMetricsOverride CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *ClearDeviceMetricsOverride) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Emulation.clearDeviceMetricsOverride", nil)
}

// ParseResponse parses the browser's response
// to the ClearDeviceMetricsOverride CDP command.
func (t *ClearDeviceMetricsOverride) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// ClearGeolocationOverride contains the parameters, and acts as
// a Go receiver, for the CDP command `clearGeolocationOverride`.
//
// Clears the overridden Geolocation Position and Error.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-clearGeolocationOverride
type ClearGeolocationOverride struct{}

// NewClearGeolocationOverride constructs a new ClearGeolocationOverride struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-clearGeolocationOverride
func NewClearGeolocationOverride() *ClearGeolocationOverride {
	return &ClearGeolocationOverride{}
}

// Do sends the ClearGeolocationOverride CDP command to a browser,
// and returns the browser's response.
func (t *ClearGeolocationOverride) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Emulation.clearGeolocationOverride", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the ClearGeolocationOverride CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *ClearGeolocationOverride) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Emulation.clearGeolocationOverride", nil)
}

// ParseResponse parses the browser's response
// to the ClearGeolocationOverride CDP command.
func (t *ClearGeolocationOverride) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// ResetPageScaleFactor contains the parameters, and acts as
// a Go receiver, for the CDP command `resetPageScaleFactor`.
//
// Requests that page scale factor is reset to initial values.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-resetPageScaleFactor
//
// This CDP method is experimental.
type ResetPageScaleFactor struct{}

// NewResetPageScaleFactor constructs a new ResetPageScaleFactor struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-resetPageScaleFactor
//
// This CDP method is experimental.
func NewResetPageScaleFactor() *ResetPageScaleFactor {
	return &ResetPageScaleFactor{}
}

// Do sends the ResetPageScaleFactor CDP command to a browser,
// and returns the browser's response.
func (t *ResetPageScaleFactor) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Emulation.resetPageScaleFactor", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the ResetPageScaleFactor CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *ResetPageScaleFactor) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Emulation.resetPageScaleFactor", nil)
}

// ParseResponse parses the browser's response
// to the ResetPageScaleFactor CDP command.
func (t *ResetPageScaleFactor) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetFocusEmulationEnabled contains the parameters, and acts as
// a Go receiver, for the CDP command `setFocusEmulationEnabled`.
//
// Enables or disables simulating a focused and active page.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setFocusEmulationEnabled
//
// This CDP method is experimental.
type SetFocusEmulationEnabled struct {
	// Whether to enable to disable focus emulation.
	Enabled bool `json:"enabled"`
}

// NewSetFocusEmulationEnabled constructs a new SetFocusEmulationEnabled struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setFocusEmulationEnabled
//
// This CDP method is experimental.
func NewSetFocusEmulationEnabled(enabled bool) *SetFocusEmulationEnabled {
	return &SetFocusEmulationEnabled{
		Enabled: enabled,
	}
}

// Do sends the SetFocusEmulationEnabled CDP command to a browser,
// and returns the browser's response.
func (t *SetFocusEmulationEnabled) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Emulation.setFocusEmulationEnabled", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetFocusEmulationEnabled CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetFocusEmulationEnabled) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Emulation.setFocusEmulationEnabled", b)
}

// ParseResponse parses the browser's response
// to the SetFocusEmulationEnabled CDP command.
func (t *SetFocusEmulationEnabled) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetAutoDarkModeOverride contains the parameters, and acts as
// a Go receiver, for the CDP command `setAutoDarkModeOverride`.
//
// Automatically render all web contents using a dark theme.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setAutoDarkModeOverride
//
// This CDP method is experimental.
type SetAutoDarkModeOverride struct {
	// Whether to enable or disable automatic dark mode.
	// If not specified, any existing override will be cleared.
	Enabled bool `json:"enabled,omitempty"`
}

// NewSetAutoDarkModeOverride constructs a new SetAutoDarkModeOverride struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setAutoDarkModeOverride
//
// This CDP method is experimental.
func NewSetAutoDarkModeOverride() *SetAutoDarkModeOverride {
	return &SetAutoDarkModeOverride{}
}

// SetEnabled adds or modifies the value of the optional
// parameter `enabled` in the SetAutoDarkModeOverride CDP command.
//
// Whether to enable or disable automatic dark mode.
// If not specified, any existing override will be cleared.
func (t *SetAutoDarkModeOverride) SetEnabled(v bool) *SetAutoDarkModeOverride {
	t.Enabled = v
	return t
}

// Do sends the SetAutoDarkModeOverride CDP command to a browser,
// and returns the browser's response.
func (t *SetAutoDarkModeOverride) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Emulation.setAutoDarkModeOverride", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetAutoDarkModeOverride CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetAutoDarkModeOverride) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Emulation.setAutoDarkModeOverride", b)
}

// ParseResponse parses the browser's response
// to the SetAutoDarkModeOverride CDP command.
func (t *SetAutoDarkModeOverride) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetCPUThrottlingRate contains the parameters, and acts as
// a Go receiver, for the CDP command `setCPUThrottlingRate`.
//
// Enables CPU throttling to emulate slow CPUs.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setCPUThrottlingRate
//
// This CDP method is experimental.
type SetCPUThrottlingRate struct {
	// Throttling rate as a slowdown factor (1 is no throttle, 2 is 2x slowdown, etc).
	Rate float64 `json:"rate"`
}

// NewSetCPUThrottlingRate constructs a new SetCPUThrottlingRate struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setCPUThrottlingRate
//
// This CDP method is experimental.
func NewSetCPUThrottlingRate(rate float64) *SetCPUThrottlingRate {
	return &SetCPUThrottlingRate{
		Rate: rate,
	}
}

// Do sends the SetCPUThrottlingRate CDP command to a browser,
// and returns the browser's response.
func (t *SetCPUThrottlingRate) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Emulation.setCPUThrottlingRate", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetCPUThrottlingRate CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetCPUThrottlingRate) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Emulation.setCPUThrottlingRate", b)
}

// ParseResponse parses the browser's response
// to the SetCPUThrottlingRate CDP command.
func (t *SetCPUThrottlingRate) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetDefaultBackgroundColorOverride contains the parameters, and acts as
// a Go receiver, for the CDP command `setDefaultBackgroundColorOverride`.
//
// Sets or clears an override of the default background color of the frame. This override is used
// if the content does not specify one.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setDefaultBackgroundColorOverride
type SetDefaultBackgroundColorOverride struct {
	// RGBA of the default background color. If not specified, any existing override will be
	// cleared.
	Color *dom.RGBA `json:"color,omitempty"`
}

// NewSetDefaultBackgroundColorOverride constructs a new SetDefaultBackgroundColorOverride struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setDefaultBackgroundColorOverride
func NewSetDefaultBackgroundColorOverride() *SetDefaultBackgroundColorOverride {
	return &SetDefaultBackgroundColorOverride{}
}

// SetColor adds or modifies the value of the optional
// parameter `color` in the SetDefaultBackgroundColorOverride CDP command.
//
// RGBA of the default background color. If not specified, any existing override will be
// cleared.
func (t *SetDefaultBackgroundColorOverride) SetColor(v dom.RGBA) *SetDefaultBackgroundColorOverride {
	t.Color = &v
	return t
}

// Do sends the SetDefaultBackgroundColorOverride CDP command to a browser,
// and returns the browser's response.
func (t *SetDefaultBackgroundColorOverride) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Emulation.setDefaultBackgroundColorOverride", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetDefaultBackgroundColorOverride CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetDefaultBackgroundColorOverride) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Emulation.setDefaultBackgroundColorOverride", b)
}

// ParseResponse parses the browser's response
// to the SetDefaultBackgroundColorOverride CDP command.
func (t *SetDefaultBackgroundColorOverride) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetDeviceMetricsOverride contains the parameters, and acts as
// a Go receiver, for the CDP command `setDeviceMetricsOverride`.
//
// Overrides the values of device screen dimensions (window.screen.width, window.screen.height,
// window.innerWidth, window.innerHeight, and "device-width"/"device-height"-related CSS media
// query results).
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setDeviceMetricsOverride
type SetDeviceMetricsOverride struct {
	// Overriding width value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	Width int64 `json:"width"`
	// Overriding height value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	Height int64 `json:"height"`
	// Overriding device scale factor value. 0 disables the override.
	DeviceScaleFactor float64 `json:"deviceScaleFactor"`
	// Whether to emulate mobile device. This includes viewport meta tag, overlay scrollbars, text
	// autosizing and more.
	Mobile bool `json:"mobile"`
	// Scale to apply to resulting view image.
	//
	// This CDP parameter is experimental.
	Scale float64 `json:"scale,omitempty"`
	// Overriding screen width value in pixels (minimum 0, maximum 10000000).
	//
	// This CDP parameter is experimental.
	ScreenWidth int64 `json:"screenWidth,omitempty"`
	// Overriding screen height value in pixels (minimum 0, maximum 10000000).
	//
	// This CDP parameter is experimental.
	ScreenHeight int64 `json:"screenHeight,omitempty"`
	// Overriding view X position on screen in pixels (minimum 0, maximum 10000000).
	//
	// This CDP parameter is experimental.
	PositionX int64 `json:"positionX,omitempty"`
	// Overriding view Y position on screen in pixels (minimum 0, maximum 10000000).
	//
	// This CDP parameter is experimental.
	PositionY int64 `json:"positionY,omitempty"`
	// Do not set visible view size, rely upon explicit setVisibleSize call.
	//
	// This CDP parameter is experimental.
	DontSetVisibleSize bool `json:"dontSetVisibleSize,omitempty"`
	// Screen orientation override.
	ScreenOrientation *ScreenOrientation `json:"screenOrientation,omitempty"`
	// If set, the visible area of the page will be overridden to this viewport. This viewport
	// change is not observed by the page, e.g. viewport-relative elements do not change positions.
	//
	// This CDP parameter is experimental.
	Viewport *page.Viewport `json:"viewport,omitempty"`
	// If set, the display feature of a multi-segment screen. If not set, multi-segment support
	// is turned-off.
	//
	// This CDP parameter is experimental.
	DisplayFeature *DisplayFeature `json:"displayFeature,omitempty"`
}

// NewSetDeviceMetricsOverride constructs a new SetDeviceMetricsOverride struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setDeviceMetricsOverride
func NewSetDeviceMetricsOverride(width int64, height int64, deviceScaleFactor float64, mobile bool) *SetDeviceMetricsOverride {
	return &SetDeviceMetricsOverride{
		Width:             width,
		Height:            height,
		DeviceScaleFactor: deviceScaleFactor,
		Mobile:            mobile,
	}
}

// SetScale adds or modifies the value of the optional
// parameter `scale` in the SetDeviceMetricsOverride CDP command.
//
// Scale to apply to resulting view image.
//
// This CDP parameter is experimental.
func (t *SetDeviceMetricsOverride) SetScale(v float64) *SetDeviceMetricsOverride {
	t.Scale = v
	return t
}

// SetScreenWidth adds or modifies the value of the optional
// parameter `screenWidth` in the SetDeviceMetricsOverride CDP command.
//
// Overriding screen width value in pixels (minimum 0, maximum 10000000).
//
// This CDP parameter is experimental.
func (t *SetDeviceMetricsOverride) SetScreenWidth(v int64) *SetDeviceMetricsOverride {
	t.ScreenWidth = v
	return t
}

// SetScreenHeight adds or modifies the value of the optional
// parameter `screenHeight` in the SetDeviceMetricsOverride CDP command.
//
// Overriding screen height value in pixels (minimum 0, maximum 10000000).
//
// This CDP parameter is experimental.
func (t *SetDeviceMetricsOverride) SetScreenHeight(v int64) *SetDeviceMetricsOverride {
	t.ScreenHeight = v
	return t
}

// SetPositionX adds or modifies the value of the optional
// parameter `positionX` in the SetDeviceMetricsOverride CDP command.
//
// Overriding view X position on screen in pixels (minimum 0, maximum 10000000).
//
// This CDP parameter is experimental.
func (t *SetDeviceMetricsOverride) SetPositionX(v int64) *SetDeviceMetricsOverride {
	t.PositionX = v
	return t
}

// SetPositionY adds or modifies the value of the optional
// parameter `positionY` in the SetDeviceMetricsOverride CDP command.
//
// Overriding view Y position on screen in pixels (minimum 0, maximum 10000000).
//
// This CDP parameter is experimental.
func (t *SetDeviceMetricsOverride) SetPositionY(v int64) *SetDeviceMetricsOverride {
	t.PositionY = v
	return t
}

// SetDontSetVisibleSize adds or modifies the value of the optional
// parameter `dontSetVisibleSize` in the SetDeviceMetricsOverride CDP command.
//
// Do not set visible view size, rely upon explicit setVisibleSize call.
//
// This CDP parameter is experimental.
func (t *SetDeviceMetricsOverride) SetDontSetVisibleSize(v bool) *SetDeviceMetricsOverride {
	t.DontSetVisibleSize = v
	return t
}

// SetScreenOrientation adds or modifies the value of the optional
// parameter `screenOrientation` in the SetDeviceMetricsOverride CDP command.
//
// Screen orientation override.
func (t *SetDeviceMetricsOverride) SetScreenOrientation(v ScreenOrientation) *SetDeviceMetricsOverride {
	t.ScreenOrientation = &v
	return t
}

// SetViewport adds or modifies the value of the optional
// parameter `viewport` in the SetDeviceMetricsOverride CDP command.
//
// If set, the visible area of the page will be overridden to this viewport. This viewport
// change is not observed by the page, e.g. viewport-relative elements do not change positions.
//
// This CDP parameter is experimental.
func (t *SetDeviceMetricsOverride) SetViewport(v page.Viewport) *SetDeviceMetricsOverride {
	t.Viewport = &v
	return t
}

// SetDisplayFeature adds or modifies the value of the optional
// parameter `displayFeature` in the SetDeviceMetricsOverride CDP command.
//
// If set, the display feature of a multi-segment screen. If not set, multi-segment support
// is turned-off.
//
// This CDP parameter is experimental.
func (t *SetDeviceMetricsOverride) SetDisplayFeature(v DisplayFeature) *SetDeviceMetricsOverride {
	t.DisplayFeature = &v
	return t
}

// Do sends the SetDeviceMetricsOverride CDP command to a browser,
// and returns the browser's response.
func (t *SetDeviceMetricsOverride) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Emulation.setDeviceMetricsOverride", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetDeviceMetricsOverride CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetDeviceMetricsOverride) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Emulation.setDeviceMetricsOverride", b)
}

// ParseResponse parses the browser's response
// to the SetDeviceMetricsOverride CDP command.
func (t *SetDeviceMetricsOverride) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetScrollbarsHidden contains the parameters, and acts as
// a Go receiver, for the CDP command `setScrollbarsHidden`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setScrollbarsHidden
//
// This CDP method is experimental.
type SetScrollbarsHidden struct {
	// Whether scrollbars should be always hidden.
	Hidden bool `json:"hidden"`
}

// NewSetScrollbarsHidden constructs a new SetScrollbarsHidden struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setScrollbarsHidden
//
// This CDP method is experimental.
func NewSetScrollbarsHidden(hidden bool) *SetScrollbarsHidden {
	return &SetScrollbarsHidden{
		Hidden: hidden,
	}
}

// Do sends the SetScrollbarsHidden CDP command to a browser,
// and returns the browser's response.
func (t *SetScrollbarsHidden) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Emulation.setScrollbarsHidden", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetScrollbarsHidden CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetScrollbarsHidden) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Emulation.setScrollbarsHidden", b)
}

// ParseResponse parses the browser's response
// to the SetScrollbarsHidden CDP command.
func (t *SetScrollbarsHidden) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetDocumentCookieDisabled contains the parameters, and acts as
// a Go receiver, for the CDP command `setDocumentCookieDisabled`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setDocumentCookieDisabled
//
// This CDP method is experimental.
type SetDocumentCookieDisabled struct {
	// Whether document.coookie API should be disabled.
	Disabled bool `json:"disabled"`
}

// NewSetDocumentCookieDisabled constructs a new SetDocumentCookieDisabled struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setDocumentCookieDisabled
//
// This CDP method is experimental.
func NewSetDocumentCookieDisabled(disabled bool) *SetDocumentCookieDisabled {
	return &SetDocumentCookieDisabled{
		Disabled: disabled,
	}
}

// Do sends the SetDocumentCookieDisabled CDP command to a browser,
// and returns the browser's response.
func (t *SetDocumentCookieDisabled) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Emulation.setDocumentCookieDisabled", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetDocumentCookieDisabled CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetDocumentCookieDisabled) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Emulation.setDocumentCookieDisabled", b)
}

// ParseResponse parses the browser's response
// to the SetDocumentCookieDisabled CDP command.
func (t *SetDocumentCookieDisabled) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetEmitTouchEventsForMouse contains the parameters, and acts as
// a Go receiver, for the CDP command `setEmitTouchEventsForMouse`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setEmitTouchEventsForMouse
//
// This CDP method is experimental.
type SetEmitTouchEventsForMouse struct {
	// Whether touch emulation based on mouse input should be enabled.
	Enabled bool `json:"enabled"`
	// Touch/gesture events configuration. Default: current platform.
	Configuration string `json:"configuration,omitempty"`
}

// NewSetEmitTouchEventsForMouse constructs a new SetEmitTouchEventsForMouse struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setEmitTouchEventsForMouse
//
// This CDP method is experimental.
func NewSetEmitTouchEventsForMouse(enabled bool) *SetEmitTouchEventsForMouse {
	return &SetEmitTouchEventsForMouse{
		Enabled: enabled,
	}
}

// SetConfiguration adds or modifies the value of the optional
// parameter `configuration` in the SetEmitTouchEventsForMouse CDP command.
//
// Touch/gesture events configuration. Default: current platform.
func (t *SetEmitTouchEventsForMouse) SetConfiguration(v string) *SetEmitTouchEventsForMouse {
	t.Configuration = v
	return t
}

// Do sends the SetEmitTouchEventsForMouse CDP command to a browser,
// and returns the browser's response.
func (t *SetEmitTouchEventsForMouse) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Emulation.setEmitTouchEventsForMouse", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetEmitTouchEventsForMouse CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetEmitTouchEventsForMouse) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Emulation.setEmitTouchEventsForMouse", b)
}

// ParseResponse parses the browser's response
// to the SetEmitTouchEventsForMouse CDP command.
func (t *SetEmitTouchEventsForMouse) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetEmulatedMedia contains the parameters, and acts as
// a Go receiver, for the CDP command `setEmulatedMedia`.
//
// Emulates the given media type or media feature for CSS media queries.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setEmulatedMedia
type SetEmulatedMedia struct {
	// Media type to emulate. Empty string disables the override.
	Media string `json:"media,omitempty"`
	// Media features to emulate.
	Features []MediaFeature `json:"features,omitempty"`
}

// NewSetEmulatedMedia constructs a new SetEmulatedMedia struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setEmulatedMedia
func NewSetEmulatedMedia() *SetEmulatedMedia {
	return &SetEmulatedMedia{}
}

// SetMedia adds or modifies the value of the optional
// parameter `media` in the SetEmulatedMedia CDP command.
//
// Media type to emulate. Empty string disables the override.
func (t *SetEmulatedMedia) SetMedia(v string) *SetEmulatedMedia {
	t.Media = v
	return t
}

// SetFeatures adds or modifies the value of the optional
// parameter `features` in the SetEmulatedMedia CDP command.
//
// Media features to emulate.
func (t *SetEmulatedMedia) SetFeatures(v []MediaFeature) *SetEmulatedMedia {
	t.Features = v
	return t
}

// Do sends the SetEmulatedMedia CDP command to a browser,
// and returns the browser's response.
func (t *SetEmulatedMedia) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Emulation.setEmulatedMedia", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetEmulatedMedia CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetEmulatedMedia) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Emulation.setEmulatedMedia", b)
}

// ParseResponse parses the browser's response
// to the SetEmulatedMedia CDP command.
func (t *SetEmulatedMedia) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetEmulatedVisionDeficiency contains the parameters, and acts as
// a Go receiver, for the CDP command `setEmulatedVisionDeficiency`.
//
// Emulates the given vision deficiency.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setEmulatedVisionDeficiency
//
// This CDP method is experimental.
type SetEmulatedVisionDeficiency struct {
	// Vision deficiency to emulate.
	Type string `json:"type"`
}

// NewSetEmulatedVisionDeficiency constructs a new SetEmulatedVisionDeficiency struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setEmulatedVisionDeficiency
//
// This CDP method is experimental.
func NewSetEmulatedVisionDeficiency(t string) *SetEmulatedVisionDeficiency {
	return &SetEmulatedVisionDeficiency{
		Type: t,
	}
}

// Do sends the SetEmulatedVisionDeficiency CDP command to a browser,
// and returns the browser's response.
func (t *SetEmulatedVisionDeficiency) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Emulation.setEmulatedVisionDeficiency", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetEmulatedVisionDeficiency CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetEmulatedVisionDeficiency) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Emulation.setEmulatedVisionDeficiency", b)
}

// ParseResponse parses the browser's response
// to the SetEmulatedVisionDeficiency CDP command.
func (t *SetEmulatedVisionDeficiency) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetGeolocationOverride contains the parameters, and acts as
// a Go receiver, for the CDP command `setGeolocationOverride`.
//
// Overrides the Geolocation Position or Error. Omitting any of the parameters emulates position
// unavailable.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setGeolocationOverride
type SetGeolocationOverride struct {
	// Mock latitude
	Latitude float64 `json:"latitude,omitempty"`
	// Mock longitude
	Longitude float64 `json:"longitude,omitempty"`
	// Mock accuracy
	Accuracy float64 `json:"accuracy,omitempty"`
}

// NewSetGeolocationOverride constructs a new SetGeolocationOverride struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setGeolocationOverride
func NewSetGeolocationOverride() *SetGeolocationOverride {
	return &SetGeolocationOverride{}
}

// SetLatitude adds or modifies the value of the optional
// parameter `latitude` in the SetGeolocationOverride CDP command.
//
// Mock latitude
func (t *SetGeolocationOverride) SetLatitude(v float64) *SetGeolocationOverride {
	t.Latitude = v
	return t
}

// SetLongitude adds or modifies the value of the optional
// parameter `longitude` in the SetGeolocationOverride CDP command.
//
// Mock longitude
func (t *SetGeolocationOverride) SetLongitude(v float64) *SetGeolocationOverride {
	t.Longitude = v
	return t
}

// SetAccuracy adds or modifies the value of the optional
// parameter `accuracy` in the SetGeolocationOverride CDP command.
//
// Mock accuracy
func (t *SetGeolocationOverride) SetAccuracy(v float64) *SetGeolocationOverride {
	t.Accuracy = v
	return t
}

// Do sends the SetGeolocationOverride CDP command to a browser,
// and returns the browser's response.
func (t *SetGeolocationOverride) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Emulation.setGeolocationOverride", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetGeolocationOverride CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetGeolocationOverride) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Emulation.setGeolocationOverride", b)
}

// ParseResponse parses the browser's response
// to the SetGeolocationOverride CDP command.
func (t *SetGeolocationOverride) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetIdleOverride contains the parameters, and acts as
// a Go receiver, for the CDP command `setIdleOverride`.
//
// Overrides the Idle state.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setIdleOverride
//
// This CDP method is experimental.
type SetIdleOverride struct {
	// Mock isUserActive
	IsUserActive bool `json:"isUserActive"`
	// Mock isScreenUnlocked
	IsScreenUnlocked bool `json:"isScreenUnlocked"`
}

// NewSetIdleOverride constructs a new SetIdleOverride struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setIdleOverride
//
// This CDP method is experimental.
func NewSetIdleOverride(isUserActive bool, isScreenUnlocked bool) *SetIdleOverride {
	return &SetIdleOverride{
		IsUserActive:     isUserActive,
		IsScreenUnlocked: isScreenUnlocked,
	}
}

// Do sends the SetIdleOverride CDP command to a browser,
// and returns the browser's response.
func (t *SetIdleOverride) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Emulation.setIdleOverride", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetIdleOverride CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetIdleOverride) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Emulation.setIdleOverride", b)
}

// ParseResponse parses the browser's response
// to the SetIdleOverride CDP command.
func (t *SetIdleOverride) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// ClearIdleOverride contains the parameters, and acts as
// a Go receiver, for the CDP command `clearIdleOverride`.
//
// Clears Idle state overrides.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-clearIdleOverride
//
// This CDP method is experimental.
type ClearIdleOverride struct{}

// NewClearIdleOverride constructs a new ClearIdleOverride struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-clearIdleOverride
//
// This CDP method is experimental.
func NewClearIdleOverride() *ClearIdleOverride {
	return &ClearIdleOverride{}
}

// Do sends the ClearIdleOverride CDP command to a browser,
// and returns the browser's response.
func (t *ClearIdleOverride) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Emulation.clearIdleOverride", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the ClearIdleOverride CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *ClearIdleOverride) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Emulation.clearIdleOverride", nil)
}

// ParseResponse parses the browser's response
// to the ClearIdleOverride CDP command.
func (t *ClearIdleOverride) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetNavigatorOverrides contains the parameters, and acts as
// a Go receiver, for the CDP command `setNavigatorOverrides`.
//
// Overrides value returned by the javascript navigator object.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setNavigatorOverrides
//
// This CDP method is deprecated.
// This CDP method is experimental.
type SetNavigatorOverrides struct {
	// The platform navigator.platform should return.
	Platform string `json:"platform"`
}

// NewSetNavigatorOverrides constructs a new SetNavigatorOverrides struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setNavigatorOverrides
//
// This CDP method is deprecated.
// This CDP method is experimental.
func NewSetNavigatorOverrides(platform string) *SetNavigatorOverrides {
	return &SetNavigatorOverrides{
		Platform: platform,
	}
}

// Do sends the SetNavigatorOverrides CDP command to a browser,
// and returns the browser's response.
func (t *SetNavigatorOverrides) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Emulation.setNavigatorOverrides", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetNavigatorOverrides CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetNavigatorOverrides) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Emulation.setNavigatorOverrides", b)
}

// ParseResponse parses the browser's response
// to the SetNavigatorOverrides CDP command.
func (t *SetNavigatorOverrides) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetPageScaleFactor contains the parameters, and acts as
// a Go receiver, for the CDP command `setPageScaleFactor`.
//
// Sets a specified page scale factor.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setPageScaleFactor
//
// This CDP method is experimental.
type SetPageScaleFactor struct {
	// Page scale factor.
	PageScaleFactor float64 `json:"pageScaleFactor"`
}

// NewSetPageScaleFactor constructs a new SetPageScaleFactor struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setPageScaleFactor
//
// This CDP method is experimental.
func NewSetPageScaleFactor(pageScaleFactor float64) *SetPageScaleFactor {
	return &SetPageScaleFactor{
		PageScaleFactor: pageScaleFactor,
	}
}

// Do sends the SetPageScaleFactor CDP command to a browser,
// and returns the browser's response.
func (t *SetPageScaleFactor) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Emulation.setPageScaleFactor", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetPageScaleFactor CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetPageScaleFactor) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Emulation.setPageScaleFactor", b)
}

// ParseResponse parses the browser's response
// to the SetPageScaleFactor CDP command.
func (t *SetPageScaleFactor) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetScriptExecutionDisabled contains the parameters, and acts as
// a Go receiver, for the CDP command `setScriptExecutionDisabled`.
//
// Switches script execution in the page.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setScriptExecutionDisabled
type SetScriptExecutionDisabled struct {
	// Whether script execution should be disabled in the page.
	Value bool `json:"value"`
}

// NewSetScriptExecutionDisabled constructs a new SetScriptExecutionDisabled struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setScriptExecutionDisabled
func NewSetScriptExecutionDisabled(value bool) *SetScriptExecutionDisabled {
	return &SetScriptExecutionDisabled{
		Value: value,
	}
}

// Do sends the SetScriptExecutionDisabled CDP command to a browser,
// and returns the browser's response.
func (t *SetScriptExecutionDisabled) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Emulation.setScriptExecutionDisabled", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetScriptExecutionDisabled CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetScriptExecutionDisabled) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Emulation.setScriptExecutionDisabled", b)
}

// ParseResponse parses the browser's response
// to the SetScriptExecutionDisabled CDP command.
func (t *SetScriptExecutionDisabled) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetTouchEmulationEnabled contains the parameters, and acts as
// a Go receiver, for the CDP command `setTouchEmulationEnabled`.
//
// Enables touch on platforms which do not support them.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setTouchEmulationEnabled
type SetTouchEmulationEnabled struct {
	// Whether the touch event emulation should be enabled.
	Enabled bool `json:"enabled"`
	// Maximum touch points supported. Defaults to one.
	MaxTouchPoints int64 `json:"maxTouchPoints,omitempty"`
}

// NewSetTouchEmulationEnabled constructs a new SetTouchEmulationEnabled struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setTouchEmulationEnabled
func NewSetTouchEmulationEnabled(enabled bool) *SetTouchEmulationEnabled {
	return &SetTouchEmulationEnabled{
		Enabled: enabled,
	}
}

// SetMaxTouchPoints adds or modifies the value of the optional
// parameter `maxTouchPoints` in the SetTouchEmulationEnabled CDP command.
//
// Maximum touch points supported. Defaults to one.
func (t *SetTouchEmulationEnabled) SetMaxTouchPoints(v int64) *SetTouchEmulationEnabled {
	t.MaxTouchPoints = v
	return t
}

// Do sends the SetTouchEmulationEnabled CDP command to a browser,
// and returns the browser's response.
func (t *SetTouchEmulationEnabled) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Emulation.setTouchEmulationEnabled", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetTouchEmulationEnabled CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetTouchEmulationEnabled) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Emulation.setTouchEmulationEnabled", b)
}

// ParseResponse parses the browser's response
// to the SetTouchEmulationEnabled CDP command.
func (t *SetTouchEmulationEnabled) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetVirtualTimePolicy contains the parameters, and acts as
// a Go receiver, for the CDP command `setVirtualTimePolicy`.
//
// Turns on virtual time for all frames (replacing real-time with a synthetic time source) and sets
// the current virtual time policy.  Note this supersedes any previous time budget.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setVirtualTimePolicy
//
// This CDP method is experimental.
type SetVirtualTimePolicy struct {
	Policy VirtualTimePolicy `json:"policy"`
	// If set, after this many virtual milliseconds have elapsed virtual time will be paused and a
	// virtualTimeBudgetExpired event is sent.
	Budget float64 `json:"budget,omitempty"`
	// If set this specifies the maximum number of tasks that can be run before virtual is forced
	// forwards to prevent deadlock.
	MaxVirtualTimeTaskStarvationCount int64 `json:"maxVirtualTimeTaskStarvationCount,omitempty"`
	// If set the virtual time policy change should be deferred until any frame starts navigating.
	// Note any previous deferred policy change is superseded.
	WaitForNavigation bool `json:"waitForNavigation,omitempty"`
	// If set, base::Time::Now will be overridden to initially return this value.
	InitialVirtualTime float64 `json:"initialVirtualTime,omitempty"`
}

// NewSetVirtualTimePolicy constructs a new SetVirtualTimePolicy struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setVirtualTimePolicy
//
// This CDP method is experimental.
func NewSetVirtualTimePolicy(policy VirtualTimePolicy) *SetVirtualTimePolicy {
	return &SetVirtualTimePolicy{
		Policy: policy,
	}
}

// SetBudget adds or modifies the value of the optional
// parameter `budget` in the SetVirtualTimePolicy CDP command.
//
// If set, after this many virtual milliseconds have elapsed virtual time will be paused and a
// virtualTimeBudgetExpired event is sent.
func (t *SetVirtualTimePolicy) SetBudget(v float64) *SetVirtualTimePolicy {
	t.Budget = v
	return t
}

// SetMaxVirtualTimeTaskStarvationCount adds or modifies the value of the optional
// parameter `maxVirtualTimeTaskStarvationCount` in the SetVirtualTimePolicy CDP command.
//
// If set this specifies the maximum number of tasks that can be run before virtual is forced
// forwards to prevent deadlock.
func (t *SetVirtualTimePolicy) SetMaxVirtualTimeTaskStarvationCount(v int64) *SetVirtualTimePolicy {
	t.MaxVirtualTimeTaskStarvationCount = v
	return t
}

// SetWaitForNavigation adds or modifies the value of the optional
// parameter `waitForNavigation` in the SetVirtualTimePolicy CDP command.
//
// If set the virtual time policy change should be deferred until any frame starts navigating.
// Note any previous deferred policy change is superseded.
func (t *SetVirtualTimePolicy) SetWaitForNavigation(v bool) *SetVirtualTimePolicy {
	t.WaitForNavigation = v
	return t
}

// SetInitialVirtualTime adds or modifies the value of the optional
// parameter `initialVirtualTime` in the SetVirtualTimePolicy CDP command.
//
// If set, base::Time::Now will be overridden to initially return this value.
func (t *SetVirtualTimePolicy) SetInitialVirtualTime(v float64) *SetVirtualTimePolicy {
	t.InitialVirtualTime = v
	return t
}

// SetVirtualTimePolicyResult contains the browser's response
// to calling the SetVirtualTimePolicy CDP command with Do().
type SetVirtualTimePolicyResult struct {
	// Absolute timestamp at which virtual time was first enabled (up time in milliseconds).
	VirtualTimeTicksBase float64 `json:"virtualTimeTicksBase"`
}

// Do sends the SetVirtualTimePolicy CDP command to a browser,
// and returns the browser's response.
func (t *SetVirtualTimePolicy) Do(ctx context.Context) (*SetVirtualTimePolicyResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Emulation.setVirtualTimePolicy", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the SetVirtualTimePolicy CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetVirtualTimePolicy) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Emulation.setVirtualTimePolicy", b)
}

// ParseResponse parses the browser's response
// to the SetVirtualTimePolicy CDP command.
func (t *SetVirtualTimePolicy) ParseResponse(m *devtools.Message) (*SetVirtualTimePolicyResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &SetVirtualTimePolicyResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetLocaleOverride contains the parameters, and acts as
// a Go receiver, for the CDP command `setLocaleOverride`.
//
// Overrides default host system locale with the specified one.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setLocaleOverride
//
// This CDP method is experimental.
type SetLocaleOverride struct {
	// ICU style C locale (e.g. "en_US"). If not specified or empty, disables the override and
	// restores default host system locale.
	Locale string `json:"locale,omitempty"`
}

// NewSetLocaleOverride constructs a new SetLocaleOverride struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setLocaleOverride
//
// This CDP method is experimental.
func NewSetLocaleOverride() *SetLocaleOverride {
	return &SetLocaleOverride{}
}

// SetLocale adds or modifies the value of the optional
// parameter `locale` in the SetLocaleOverride CDP command.
//
// ICU style C locale (e.g. "en_US"). If not specified or empty, disables the override and
// restores default host system locale.
func (t *SetLocaleOverride) SetLocale(v string) *SetLocaleOverride {
	t.Locale = v
	return t
}

// Do sends the SetLocaleOverride CDP command to a browser,
// and returns the browser's response.
func (t *SetLocaleOverride) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Emulation.setLocaleOverride", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetLocaleOverride CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetLocaleOverride) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Emulation.setLocaleOverride", b)
}

// ParseResponse parses the browser's response
// to the SetLocaleOverride CDP command.
func (t *SetLocaleOverride) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetTimezoneOverride contains the parameters, and acts as
// a Go receiver, for the CDP command `setTimezoneOverride`.
//
// Overrides default host system timezone with the specified one.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setTimezoneOverride
//
// This CDP method is experimental.
type SetTimezoneOverride struct {
	// The timezone identifier. If empty, disables the override and
	// restores default host system timezone.
	TimezoneID string `json:"timezoneId"`
}

// NewSetTimezoneOverride constructs a new SetTimezoneOverride struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setTimezoneOverride
//
// This CDP method is experimental.
func NewSetTimezoneOverride(timezoneID string) *SetTimezoneOverride {
	return &SetTimezoneOverride{
		TimezoneID: timezoneID,
	}
}

// Do sends the SetTimezoneOverride CDP command to a browser,
// and returns the browser's response.
func (t *SetTimezoneOverride) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Emulation.setTimezoneOverride", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetTimezoneOverride CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetTimezoneOverride) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Emulation.setTimezoneOverride", b)
}

// ParseResponse parses the browser's response
// to the SetTimezoneOverride CDP command.
func (t *SetTimezoneOverride) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetVisibleSize contains the parameters, and acts as
// a Go receiver, for the CDP command `setVisibleSize`.
//
// Resizes the frame/viewport of the page. Note that this does not affect the frame's container
// (e.g. browser window). Can be used to produce screenshots of the specified size. Not supported
// on Android.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setVisibleSize
//
// This CDP method is deprecated.
// This CDP method is experimental.
type SetVisibleSize struct {
	// Frame width (DIP).
	Width int64 `json:"width"`
	// Frame height (DIP).
	Height int64 `json:"height"`
}

// NewSetVisibleSize constructs a new SetVisibleSize struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setVisibleSize
//
// This CDP method is deprecated.
// This CDP method is experimental.
func NewSetVisibleSize(width int64, height int64) *SetVisibleSize {
	return &SetVisibleSize{
		Width:  width,
		Height: height,
	}
}

// Do sends the SetVisibleSize CDP command to a browser,
// and returns the browser's response.
func (t *SetVisibleSize) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Emulation.setVisibleSize", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetVisibleSize CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetVisibleSize) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Emulation.setVisibleSize", b)
}

// ParseResponse parses the browser's response
// to the SetVisibleSize CDP command.
func (t *SetVisibleSize) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetDisabledImageTypes contains the parameters, and acts as
// a Go receiver, for the CDP command `setDisabledImageTypes`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setDisabledImageTypes
//
// This CDP method is experimental.
type SetDisabledImageTypes struct {
	// Image types to disable.
	ImageTypes []DisabledImageType `json:"imageTypes"`
}

// NewSetDisabledImageTypes constructs a new SetDisabledImageTypes struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setDisabledImageTypes
//
// This CDP method is experimental.
func NewSetDisabledImageTypes(imageTypes []DisabledImageType) *SetDisabledImageTypes {
	return &SetDisabledImageTypes{
		ImageTypes: imageTypes,
	}
}

// Do sends the SetDisabledImageTypes CDP command to a browser,
// and returns the browser's response.
func (t *SetDisabledImageTypes) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Emulation.setDisabledImageTypes", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetDisabledImageTypes CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetDisabledImageTypes) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Emulation.setDisabledImageTypes", b)
}

// ParseResponse parses the browser's response
// to the SetDisabledImageTypes CDP command.
func (t *SetDisabledImageTypes) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetUserAgentOverride contains the parameters, and acts as
// a Go receiver, for the CDP command `setUserAgentOverride`.
//
// Allows overriding user agent with the given string.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setUserAgentOverride
type SetUserAgentOverride struct {
	// User agent to use.
	UserAgent string `json:"userAgent"`
	// Browser langugage to emulate.
	AcceptLanguage string `json:"acceptLanguage,omitempty"`
	// The platform navigator.platform should return.
	Platform string `json:"platform,omitempty"`
	// To be sent in Sec-CH-UA-* headers and returned in navigator.userAgentData
	//
	// This CDP parameter is experimental.
	UserAgentMetadata *UserAgentMetadata `json:"userAgentMetadata,omitempty"`
}

// NewSetUserAgentOverride constructs a new SetUserAgentOverride struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setUserAgentOverride
func NewSetUserAgentOverride(userAgent string) *SetUserAgentOverride {
	return &SetUserAgentOverride{
		UserAgent: userAgent,
	}
}

// SetAcceptLanguage adds or modifies the value of the optional
// parameter `acceptLanguage` in the SetUserAgentOverride CDP command.
//
// Browser langugage to emulate.
func (t *SetUserAgentOverride) SetAcceptLanguage(v string) *SetUserAgentOverride {
	t.AcceptLanguage = v
	return t
}

// SetPlatform adds or modifies the value of the optional
// parameter `platform` in the SetUserAgentOverride CDP command.
//
// The platform navigator.platform should return.
func (t *SetUserAgentOverride) SetPlatform(v string) *SetUserAgentOverride {
	t.Platform = v
	return t
}

// SetUserAgentMetadata adds or modifies the value of the optional
// parameter `userAgentMetadata` in the SetUserAgentOverride CDP command.
//
// To be sent in Sec-CH-UA-* headers and returned in navigator.userAgentData
//
// This CDP parameter is experimental.
func (t *SetUserAgentOverride) SetUserAgentMetadata(v UserAgentMetadata) *SetUserAgentOverride {
	t.UserAgentMetadata = &v
	return t
}

// Do sends the SetUserAgentOverride CDP command to a browser,
// and returns the browser's response.
func (t *SetUserAgentOverride) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Emulation.setUserAgentOverride", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetUserAgentOverride CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetUserAgentOverride) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Emulation.setUserAgentOverride", b)
}

// ParseResponse parses the browser's response
// to the SetUserAgentOverride CDP command.
func (t *SetUserAgentOverride) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}
