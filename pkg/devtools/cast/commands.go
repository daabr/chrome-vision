package cast

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
)

// Enable contains the parameters, and acts as
// a Go receiver, for the CDP command `enable`.
//
// Starts observing for sinks that can be used for tab mirroring, and if set,
// sinks compatible with |presentationUrl| as well. When sinks are found, a
// |sinksUpdated| event is fired.
// Also starts observing for issue messages. When an issue is added or removed,
// an |issueUpdated| event is fired.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Cast/#method-enable
type Enable struct {
	PresentationURL string `json:"presentationUrl,omitempty"`
}

// NewEnable constructs a new Enable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Cast/#method-enable
func NewEnable() *Enable {
	return &Enable{}
}

// SetPresentationURL adds or modifies the value of the optional
// parameter `presentationUrl` in the Enable CDP command.
func (t *Enable) SetPresentationURL(v string) *Enable {
	t.PresentationURL = v
	return t
}

// Do sends the Enable CDP command to a browser,
// and returns the browser's response.
func (t *Enable) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Cast.enable", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the Enable CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *Enable) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Cast.enable", b)
}

// ParseResponse parses the browser's response
// to the Enable CDP command.
func (t *Enable) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// Disable contains the parameters, and acts as
// a Go receiver, for the CDP command `disable`.
//
// Stops observing for sinks and issues.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Cast/#method-disable
type Disable struct{}

// NewDisable constructs a new Disable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Cast/#method-disable
func NewDisable() *Disable {
	return &Disable{}
}

// Do sends the Disable CDP command to a browser,
// and returns the browser's response.
func (t *Disable) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Cast.disable", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the Disable CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *Disable) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Cast.disable", nil)
}

// ParseResponse parses the browser's response
// to the Disable CDP command.
func (t *Disable) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetSinkToUse contains the parameters, and acts as
// a Go receiver, for the CDP command `setSinkToUse`.
//
// Sets a sink to be used when the web page requests the browser to choose a
// sink via Presentation API, Remote Playback API, or Cast SDK.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Cast/#method-setSinkToUse
type SetSinkToUse struct {
	SinkName string `json:"sinkName"`
}

// NewSetSinkToUse constructs a new SetSinkToUse struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Cast/#method-setSinkToUse
func NewSetSinkToUse(sinkName string) *SetSinkToUse {
	return &SetSinkToUse{
		SinkName: sinkName,
	}
}

// Do sends the SetSinkToUse CDP command to a browser,
// and returns the browser's response.
func (t *SetSinkToUse) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Cast.setSinkToUse", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetSinkToUse CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetSinkToUse) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Cast.setSinkToUse", b)
}

// ParseResponse parses the browser's response
// to the SetSinkToUse CDP command.
func (t *SetSinkToUse) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// StartDesktopMirroring contains the parameters, and acts as
// a Go receiver, for the CDP command `startDesktopMirroring`.
//
// Starts mirroring the desktop to the sink.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Cast/#method-startDesktopMirroring
type StartDesktopMirroring struct {
	SinkName string `json:"sinkName"`
}

// NewStartDesktopMirroring constructs a new StartDesktopMirroring struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Cast/#method-startDesktopMirroring
func NewStartDesktopMirroring(sinkName string) *StartDesktopMirroring {
	return &StartDesktopMirroring{
		SinkName: sinkName,
	}
}

// Do sends the StartDesktopMirroring CDP command to a browser,
// and returns the browser's response.
func (t *StartDesktopMirroring) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Cast.startDesktopMirroring", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the StartDesktopMirroring CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *StartDesktopMirroring) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Cast.startDesktopMirroring", b)
}

// ParseResponse parses the browser's response
// to the StartDesktopMirroring CDP command.
func (t *StartDesktopMirroring) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// StartTabMirroring contains the parameters, and acts as
// a Go receiver, for the CDP command `startTabMirroring`.
//
// Starts mirroring the tab to the sink.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Cast/#method-startTabMirroring
type StartTabMirroring struct {
	SinkName string `json:"sinkName"`
}

// NewStartTabMirroring constructs a new StartTabMirroring struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Cast/#method-startTabMirroring
func NewStartTabMirroring(sinkName string) *StartTabMirroring {
	return &StartTabMirroring{
		SinkName: sinkName,
	}
}

// Do sends the StartTabMirroring CDP command to a browser,
// and returns the browser's response.
func (t *StartTabMirroring) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Cast.startTabMirroring", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the StartTabMirroring CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *StartTabMirroring) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Cast.startTabMirroring", b)
}

// ParseResponse parses the browser's response
// to the StartTabMirroring CDP command.
func (t *StartTabMirroring) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// StopCasting contains the parameters, and acts as
// a Go receiver, for the CDP command `stopCasting`.
//
// Stops the active Cast session on the sink.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Cast/#method-stopCasting
type StopCasting struct {
	SinkName string `json:"sinkName"`
}

// NewStopCasting constructs a new StopCasting struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Cast/#method-stopCasting
func NewStopCasting(sinkName string) *StopCasting {
	return &StopCasting{
		SinkName: sinkName,
	}
}

// Do sends the StopCasting CDP command to a browser,
// and returns the browser's response.
func (t *StopCasting) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Cast.stopCasting", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the StopCasting CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *StopCasting) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Cast.stopCasting", b)
}

// ParseResponse parses the browser's response
// to the StopCasting CDP command.
func (t *StopCasting) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}
