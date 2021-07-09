package webaudio

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
)

// Enable contains the parameters, and acts as
// a Go receiver, for the CDP command `enable`.
//
// Enables the WebAudio domain and starts sending context lifetime events.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#method-enable
type Enable struct{}

// NewEnable constructs a new Enable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#method-enable
func NewEnable() *Enable {
	return &Enable{}
}

// Do sends the Enable CDP command to a browser,
// and returns the browser's response.
func (t *Enable) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "WebAudio.enable", nil)
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
	return devtools.Send(ctx, "WebAudio.enable", nil)
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
// Disables the WebAudio domain.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#method-disable
type Disable struct{}

// NewDisable constructs a new Disable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#method-disable
func NewDisable() *Disable {
	return &Disable{}
}

// Do sends the Disable CDP command to a browser,
// and returns the browser's response.
func (t *Disable) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "WebAudio.disable", nil)
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
	return devtools.Send(ctx, "WebAudio.disable", nil)
}

// ParseResponse parses the browser's response
// to the Disable CDP command.
func (t *Disable) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// GetRealtimeData contains the parameters, and acts as
// a Go receiver, for the CDP command `getRealtimeData`.
//
// Fetch the realtime data from the registered contexts.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#method-getRealtimeData
type GetRealtimeData struct {
	ContextID string `json:"contextId"`
}

// NewGetRealtimeData constructs a new GetRealtimeData struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#method-getRealtimeData
func NewGetRealtimeData(contextID string) *GetRealtimeData {
	return &GetRealtimeData{
		ContextID: contextID,
	}
}

// GetRealtimeDataResult contains the browser's response
// to calling the GetRealtimeData CDP command with Do().
type GetRealtimeDataResult struct {
	RealtimeData ContextRealtimeData `json:"realtimeData"`
}

// Do sends the GetRealtimeData CDP command to a browser,
// and returns the browser's response.
func (t *GetRealtimeData) Do(ctx context.Context) (*GetRealtimeDataResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "WebAudio.getRealtimeData", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetRealtimeData CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetRealtimeData) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "WebAudio.getRealtimeData", b)
}

// ParseResponse parses the browser's response
// to the GetRealtimeData CDP command.
func (t *GetRealtimeData) ParseResponse(m *devtools.Message) (*GetRealtimeDataResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetRealtimeDataResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}
