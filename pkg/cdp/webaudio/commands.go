package webaudio

import (
	"context"
	"fmt"
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
	fmt.Println(ctx)
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
	fmt.Println(ctx)
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
func NewGetRealtimeData(contextId GraphObjectID) *GetRealtimeData {
	return &GetRealtimeData{
		ContextID: contextId,
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
	fmt.Println(ctx)
	return new(GetRealtimeDataResponse), nil
}
