package webaudio

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/cdp"
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
	response, err := cdp.Send(ctx, "WebAudio.enable", nil)
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
	response, err := cdp.Send(ctx, "WebAudio.disable", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
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
	ContextID GraphObjectID `json:"contextId"`
}

// NewGetRealtimeData constructs a new GetRealtimeData struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#method-getRealtimeData
func NewGetRealtimeData(contextID GraphObjectID) *GetRealtimeData {
	return &GetRealtimeData{
		ContextID: contextID,
	}
}

// GetRealtimeDataResponse contains the browser's response
// to calling the GetRealtimeData CDP command with Do().
type GetRealtimeDataResponse struct {
	RealtimeData ContextRealtimeData `json:"realtimeData"`
}

// Do sends the GetRealtimeData CDP command to a browser,
// and returns the browser's response.
func (t *GetRealtimeData) Do(ctx context.Context) (*GetRealtimeDataResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "WebAudio.getRealtimeData", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetRealtimeDataResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}
