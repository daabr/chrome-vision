package domsnapshot

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
)

// Disable contains the parameters, and acts as
// a Go receiver, for the CDP command `disable`.
//
// Disables DOM snapshot agent for the given page.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/#method-disable
type Disable struct{}

// NewDisable constructs a new Disable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/#method-disable
func NewDisable() *Disable {
	return &Disable{}
}

// Do sends the Disable CDP command to a browser,
// and returns the browser's response.
func (t *Disable) Do(ctx context.Context) error {
	response, err := devtools.SendAndWait(ctx, "DOMSnapshot.disable", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// Enable contains the parameters, and acts as
// a Go receiver, for the CDP command `enable`.
//
// Enables DOM snapshot agent for the given page.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/#method-enable
type Enable struct{}

// NewEnable constructs a new Enable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/#method-enable
func NewEnable() *Enable {
	return &Enable{}
}

// Do sends the Enable CDP command to a browser,
// and returns the browser's response.
func (t *Enable) Do(ctx context.Context) error {
	response, err := devtools.SendAndWait(ctx, "DOMSnapshot.enable", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// GetSnapshot contains the parameters, and acts as
// a Go receiver, for the CDP command `getSnapshot`.
//
// Returns a document snapshot, including the full DOM tree of the root node (including iframes,
// template contents, and imported documents) in a flattened array, as well as layout and
// white-listed computed style information for the nodes. Shadow DOM in the returned DOM tree is
// flattened.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/#method-getSnapshot
//
// This CDP method is deprecated.
type GetSnapshot struct {
	// Whitelist of computed styles to return.
	ComputedStyleWhitelist []string `json:"computedStyleWhitelist"`
	// Whether or not to retrieve details of DOM listeners (default false).
	IncludeEventListeners bool `json:"includeEventListeners,omitempty"`
	// Whether to determine and include the paint order index of LayoutTreeNodes (default false).
	IncludePaintOrder bool `json:"includePaintOrder,omitempty"`
	// Whether to include UA shadow tree in the snapshot (default false).
	IncludeUserAgentShadowTree bool `json:"includeUserAgentShadowTree,omitempty"`
}

// NewGetSnapshot constructs a new GetSnapshot struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/#method-getSnapshot
//
// This CDP method is deprecated.
func NewGetSnapshot(computedStyleWhitelist []string) *GetSnapshot {
	return &GetSnapshot{
		ComputedStyleWhitelist: computedStyleWhitelist,
	}
}

// SetIncludeEventListeners adds or modifies the value of the optional
// parameter `includeEventListeners` in the GetSnapshot CDP command.
//
// Whether or not to retrieve details of DOM listeners (default false).
func (t *GetSnapshot) SetIncludeEventListeners(v bool) *GetSnapshot {
	t.IncludeEventListeners = v
	return t
}

// SetIncludePaintOrder adds or modifies the value of the optional
// parameter `includePaintOrder` in the GetSnapshot CDP command.
//
// Whether to determine and include the paint order index of LayoutTreeNodes (default false).
func (t *GetSnapshot) SetIncludePaintOrder(v bool) *GetSnapshot {
	t.IncludePaintOrder = v
	return t
}

// SetIncludeUserAgentShadowTree adds or modifies the value of the optional
// parameter `includeUserAgentShadowTree` in the GetSnapshot CDP command.
//
// Whether to include UA shadow tree in the snapshot (default false).
func (t *GetSnapshot) SetIncludeUserAgentShadowTree(v bool) *GetSnapshot {
	t.IncludeUserAgentShadowTree = v
	return t
}

// GetSnapshotResult contains the browser's response
// to calling the GetSnapshot CDP command with Do().
type GetSnapshotResult struct {
	// The nodes in the DOM tree. The DOMNode at index 0 corresponds to the root document.
	DomNodes []DOMNode `json:"domNodes"`
	// The nodes in the layout tree.
	LayoutTreeNodes []LayoutTreeNode `json:"layoutTreeNodes"`
	// Whitelisted ComputedStyle properties for each node in the layout tree.
	ComputedStyles []ComputedStyle `json:"computedStyles"`
}

// Do sends the GetSnapshot CDP command to a browser,
// and returns the browser's response.
func (t *GetSnapshot) Do(ctx context.Context) (*GetSnapshotResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.SendAndWait(ctx, "DOMSnapshot.getSnapshot", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetSnapshotResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// CaptureSnapshot contains the parameters, and acts as
// a Go receiver, for the CDP command `captureSnapshot`.
//
// Returns a document snapshot, including the full DOM tree of the root node (including iframes,
// template contents, and imported documents) in a flattened array, as well as layout and
// white-listed computed style information for the nodes. Shadow DOM in the returned DOM tree is
// flattened.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/#method-captureSnapshot
type CaptureSnapshot struct {
	// Whitelist of computed styles to return.
	ComputedStyles []string `json:"computedStyles"`
	// Whether to include layout object paint orders into the snapshot.
	IncludePaintOrder bool `json:"includePaintOrder,omitempty"`
	// Whether to include DOM rectangles (offsetRects, clientRects, scrollRects) into the snapshot
	IncludeDOMRects bool `json:"includeDOMRects,omitempty"`
	// Whether to include blended background colors in the snapshot (default: false).
	// Blended background color is achieved by blending background colors of all elements
	// that overlap with the current element.
	//
	// This CDP parameter is experimental.
	IncludeBlendedBackgroundColors bool `json:"includeBlendedBackgroundColors,omitempty"`
	// Whether to include text color opacity in the snapshot (default: false).
	// An element might have the opacity property set that affects the text color of the element.
	// The final text color opacity is computed based on the opacity of all overlapping elements.
	//
	// This CDP parameter is experimental.
	IncludeTextColorOpacities bool `json:"includeTextColorOpacities,omitempty"`
}

// NewCaptureSnapshot constructs a new CaptureSnapshot struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/#method-captureSnapshot
func NewCaptureSnapshot(computedStyles []string) *CaptureSnapshot {
	return &CaptureSnapshot{
		ComputedStyles: computedStyles,
	}
}

// SetIncludePaintOrder adds or modifies the value of the optional
// parameter `includePaintOrder` in the CaptureSnapshot CDP command.
//
// Whether to include layout object paint orders into the snapshot.
func (t *CaptureSnapshot) SetIncludePaintOrder(v bool) *CaptureSnapshot {
	t.IncludePaintOrder = v
	return t
}

// SetIncludeDOMRects adds or modifies the value of the optional
// parameter `includeDOMRects` in the CaptureSnapshot CDP command.
//
// Whether to include DOM rectangles (offsetRects, clientRects, scrollRects) into the snapshot
func (t *CaptureSnapshot) SetIncludeDOMRects(v bool) *CaptureSnapshot {
	t.IncludeDOMRects = v
	return t
}

// SetIncludeBlendedBackgroundColors adds or modifies the value of the optional
// parameter `includeBlendedBackgroundColors` in the CaptureSnapshot CDP command.
//
// Whether to include blended background colors in the snapshot (default: false).
// Blended background color is achieved by blending background colors of all elements
// that overlap with the current element.
//
// This CDP parameter is experimental.
func (t *CaptureSnapshot) SetIncludeBlendedBackgroundColors(v bool) *CaptureSnapshot {
	t.IncludeBlendedBackgroundColors = v
	return t
}

// SetIncludeTextColorOpacities adds or modifies the value of the optional
// parameter `includeTextColorOpacities` in the CaptureSnapshot CDP command.
//
// Whether to include text color opacity in the snapshot (default: false).
// An element might have the opacity property set that affects the text color of the element.
// The final text color opacity is computed based on the opacity of all overlapping elements.
//
// This CDP parameter is experimental.
func (t *CaptureSnapshot) SetIncludeTextColorOpacities(v bool) *CaptureSnapshot {
	t.IncludeTextColorOpacities = v
	return t
}

// CaptureSnapshotResult contains the browser's response
// to calling the CaptureSnapshot CDP command with Do().
type CaptureSnapshotResult struct {
	// The nodes in the DOM tree. The DOMNode at index 0 corresponds to the root document.
	Documents []DocumentSnapshot `json:"documents"`
	// Shared string table that all string properties refer to with indexes.
	Strings []string `json:"strings"`
}

// Do sends the CaptureSnapshot CDP command to a browser,
// and returns the browser's response.
func (t *CaptureSnapshot) Do(ctx context.Context) (*CaptureSnapshotResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.SendAndWait(ctx, "DOMSnapshot.captureSnapshot", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &CaptureSnapshotResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}
