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
	response, err := devtools.SendAndWait(ctx, "Cast.enable", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
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
	response, err := devtools.SendAndWait(ctx, "Cast.disable", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
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
	response, err := devtools.SendAndWait(ctx, "Cast.setSinkToUse", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
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
	response, err := devtools.SendAndWait(ctx, "Cast.startTabMirroring", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
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
	response, err := devtools.SendAndWait(ctx, "Cast.stopCasting", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}
