package applicationcache

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
)

// Enable contains the parameters, and acts as
// a Go receiver, for the CDP command `enable`.
//
// Enables application cache domain notifications.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-enable
type Enable struct{}

// NewEnable constructs a new Enable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-enable
func NewEnable() *Enable {
	return &Enable{}
}

// Do sends the Enable CDP command to a browser,
// and returns the browser's response.
func (t *Enable) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "ApplicationCache.enable", nil)
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
	return devtools.Send(ctx, "ApplicationCache.enable", nil)
}

// ParseResponse parses the browser's response
// to the Enable CDP command.
func (t *Enable) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// GetApplicationCacheForFrame contains the parameters, and acts as
// a Go receiver, for the CDP command `getApplicationCacheForFrame`.
//
// Returns relevant application cache data for the document in given frame.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getApplicationCacheForFrame
type GetApplicationCacheForFrame struct {
	// Identifier of the frame containing document whose application cache is retrieved.
	FrameID string `json:"frameId"`
}

// NewGetApplicationCacheForFrame constructs a new GetApplicationCacheForFrame struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getApplicationCacheForFrame
func NewGetApplicationCacheForFrame(frameID string) *GetApplicationCacheForFrame {
	return &GetApplicationCacheForFrame{
		FrameID: frameID,
	}
}

// GetApplicationCacheForFrameResult contains the browser's response
// to calling the GetApplicationCacheForFrame CDP command with Do().
type GetApplicationCacheForFrameResult struct {
	// Relevant application cache data for the document in given frame.
	ApplicationCache ApplicationCache `json:"applicationCache"`
}

// Do sends the GetApplicationCacheForFrame CDP command to a browser,
// and returns the browser's response.
func (t *GetApplicationCacheForFrame) Do(ctx context.Context) (*GetApplicationCacheForFrameResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "ApplicationCache.getApplicationCacheForFrame", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetApplicationCacheForFrame CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetApplicationCacheForFrame) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "ApplicationCache.getApplicationCacheForFrame", b)
}

// ParseResponse parses the browser's response
// to the GetApplicationCacheForFrame CDP command.
func (t *GetApplicationCacheForFrame) ParseResponse(m *devtools.Message) (*GetApplicationCacheForFrameResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetApplicationCacheForFrameResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetFramesWithManifests contains the parameters, and acts as
// a Go receiver, for the CDP command `getFramesWithManifests`.
//
// Returns array of frame identifiers with manifest urls for each frame containing a document
// associated with some application cache.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getFramesWithManifests
type GetFramesWithManifests struct{}

// NewGetFramesWithManifests constructs a new GetFramesWithManifests struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getFramesWithManifests
func NewGetFramesWithManifests() *GetFramesWithManifests {
	return &GetFramesWithManifests{}
}

// GetFramesWithManifestsResult contains the browser's response
// to calling the GetFramesWithManifests CDP command with Do().
type GetFramesWithManifestsResult struct {
	// Array of frame identifiers with manifest urls for each frame containing a document
	// associated with some application cache.
	FrameIds []FrameWithManifest `json:"frameIds"`
}

// Do sends the GetFramesWithManifests CDP command to a browser,
// and returns the browser's response.
func (t *GetFramesWithManifests) Do(ctx context.Context) (*GetFramesWithManifestsResult, error) {
	m, err := devtools.SendAndWait(ctx, "ApplicationCache.getFramesWithManifests", nil)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetFramesWithManifests CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetFramesWithManifests) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "ApplicationCache.getFramesWithManifests", nil)
}

// ParseResponse parses the browser's response
// to the GetFramesWithManifests CDP command.
func (t *GetFramesWithManifests) ParseResponse(m *devtools.Message) (*GetFramesWithManifestsResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetFramesWithManifestsResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetManifestForFrame contains the parameters, and acts as
// a Go receiver, for the CDP command `getManifestForFrame`.
//
// Returns manifest URL for document in the given frame.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getManifestForFrame
type GetManifestForFrame struct {
	// Identifier of the frame containing document whose manifest is retrieved.
	FrameID string `json:"frameId"`
}

// NewGetManifestForFrame constructs a new GetManifestForFrame struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getManifestForFrame
func NewGetManifestForFrame(frameID string) *GetManifestForFrame {
	return &GetManifestForFrame{
		FrameID: frameID,
	}
}

// GetManifestForFrameResult contains the browser's response
// to calling the GetManifestForFrame CDP command with Do().
type GetManifestForFrameResult struct {
	// Manifest URL for document in the given frame.
	ManifestURL string `json:"manifestURL"`
}

// Do sends the GetManifestForFrame CDP command to a browser,
// and returns the browser's response.
func (t *GetManifestForFrame) Do(ctx context.Context) (*GetManifestForFrameResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "ApplicationCache.getManifestForFrame", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetManifestForFrame CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetManifestForFrame) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "ApplicationCache.getManifestForFrame", b)
}

// ParseResponse parses the browser's response
// to the GetManifestForFrame CDP command.
func (t *GetManifestForFrame) ParseResponse(m *devtools.Message) (*GetManifestForFrameResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetManifestForFrameResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}
