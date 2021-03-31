package applicationcache

import (
	"context"
	"fmt"

	"github.com/daabr/chrome-vision/pkg/cdp"
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
	fmt.Println(ctx)
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
	FrameID cdp.FrameID `json:"frameId"`
}

// NewGetApplicationCacheForFrame constructs a new GetApplicationCacheForFrame struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getApplicationCacheForFrame
func NewGetApplicationCacheForFrame(frameId cdp.FrameID) *GetApplicationCacheForFrame {
	return &GetApplicationCacheForFrame{
		FrameID: frameId,
	}
}

// GetApplicationCacheForFrameResponse contains the browser's response
// to calling the GetApplicationCacheForFrame CDP command with Do().
type GetApplicationCacheForFrameResponse struct {
	// Relevant application cache data for the document in given frame.
	ApplicationCache ApplicationCache `json:"applicationCache"`
}

// Do sends the GetApplicationCacheForFrame CDP command to a browser,
// and returns the browser's response.
func (t *GetApplicationCacheForFrame) Do(ctx context.Context) (*GetApplicationCacheForFrameResponse, error) {
	fmt.Println(ctx)
	return new(GetApplicationCacheForFrameResponse), nil
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

// GetFramesWithManifestsResponse contains the browser's response
// to calling the GetFramesWithManifests CDP command with Do().
type GetFramesWithManifestsResponse struct {
	// Array of frame identifiers with manifest urls for each frame containing a document
	// associated with some application cache.
	FrameIds []FrameWithManifest `json:"frameIds"`
}

// Do sends the GetFramesWithManifests CDP command to a browser,
// and returns the browser's response.
func (t *GetFramesWithManifests) Do(ctx context.Context) (*GetFramesWithManifestsResponse, error) {
	fmt.Println(ctx)
	return new(GetFramesWithManifestsResponse), nil
}

// GetManifestForFrame contains the parameters, and acts as
// a Go receiver, for the CDP command `getManifestForFrame`.
//
// Returns manifest URL for document in the given frame.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getManifestForFrame
type GetManifestForFrame struct {
	// Identifier of the frame containing document whose manifest is retrieved.
	FrameID cdp.FrameID `json:"frameId"`
}

// NewGetManifestForFrame constructs a new GetManifestForFrame struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getManifestForFrame
func NewGetManifestForFrame(frameId cdp.FrameID) *GetManifestForFrame {
	return &GetManifestForFrame{
		FrameID: frameId,
	}
}

// GetManifestForFrameResponse contains the browser's response
// to calling the GetManifestForFrame CDP command with Do().
type GetManifestForFrameResponse struct {
	// Manifest URL for document in the given frame.
	ManifestURL string `json:"manifestURL"`
}

// Do sends the GetManifestForFrame CDP command to a browser,
// and returns the browser's response.
func (t *GetManifestForFrame) Do(ctx context.Context) (*GetManifestForFrameResponse, error) {
	fmt.Println(ctx)
	return new(GetManifestForFrameResponse), nil
}
