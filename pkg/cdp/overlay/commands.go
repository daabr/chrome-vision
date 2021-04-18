package overlay

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/cdp"
	"github.com/daabr/chrome-vision/pkg/cdp/dom"
	"github.com/daabr/chrome-vision/pkg/cdp/runtime"
)

// Disable contains the parameters, and acts as
// a Go receiver, for the CDP command `disable`.
//
// Disables domain notifications.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-disable
type Disable struct{}

// NewDisable constructs a new Disable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-disable
func NewDisable() *Disable {
	return &Disable{}
}

// Do sends the Disable CDP command to a browser,
// and returns the browser's response.
func (t *Disable) Do(ctx context.Context) error {
	response, err := cdp.Send(ctx, "Disable", nil)
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
// Enables domain notifications.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-enable
type Enable struct{}

// NewEnable constructs a new Enable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-enable
func NewEnable() *Enable {
	return &Enable{}
}

// Do sends the Enable CDP command to a browser,
// and returns the browser's response.
func (t *Enable) Do(ctx context.Context) error {
	response, err := cdp.Send(ctx, "Enable", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// GetHighlightObjectForTest contains the parameters, and acts as
// a Go receiver, for the CDP command `getHighlightObjectForTest`.
//
// For testing.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-getHighlightObjectForTest
type GetHighlightObjectForTest struct {
	// Id of the node to get highlight object for.
	NodeID dom.NodeID `json:"nodeId"`
	// Whether to include distance info.
	IncludeDistance bool `json:"includeDistance,omitempty"`
	// Whether to include style info.
	IncludeStyle bool `json:"includeStyle,omitempty"`
	// The color format to get config with (default: hex).
	ColorFormat *ColorFormat `json:"colorFormat,omitempty"`
	// Whether to show accessibility info (default: true).
	ShowAccessibilityInfo bool `json:"showAccessibilityInfo,omitempty"`
}

// NewGetHighlightObjectForTest constructs a new GetHighlightObjectForTest struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-getHighlightObjectForTest
func NewGetHighlightObjectForTest(nodeId dom.NodeID) *GetHighlightObjectForTest {
	return &GetHighlightObjectForTest{
		NodeID: nodeId,
	}
}

// SetIncludeDistance adds or modifies the value of the optional
// parameter `includeDistance` in the GetHighlightObjectForTest CDP command.
//
// Whether to include distance info.
func (t *GetHighlightObjectForTest) SetIncludeDistance(v bool) *GetHighlightObjectForTest {
	t.IncludeDistance = v
	return t
}

// SetIncludeStyle adds or modifies the value of the optional
// parameter `includeStyle` in the GetHighlightObjectForTest CDP command.
//
// Whether to include style info.
func (t *GetHighlightObjectForTest) SetIncludeStyle(v bool) *GetHighlightObjectForTest {
	t.IncludeStyle = v
	return t
}

// SetColorFormat adds or modifies the value of the optional
// parameter `colorFormat` in the GetHighlightObjectForTest CDP command.
//
// The color format to get config with (default: hex).
func (t *GetHighlightObjectForTest) SetColorFormat(v ColorFormat) *GetHighlightObjectForTest {
	t.ColorFormat = &v
	return t
}

// SetShowAccessibilityInfo adds or modifies the value of the optional
// parameter `showAccessibilityInfo` in the GetHighlightObjectForTest CDP command.
//
// Whether to show accessibility info (default: true).
func (t *GetHighlightObjectForTest) SetShowAccessibilityInfo(v bool) *GetHighlightObjectForTest {
	t.ShowAccessibilityInfo = v
	return t
}

// GetHighlightObjectForTestResponse contains the browser's response
// to calling the GetHighlightObjectForTest CDP command with Do().
type GetHighlightObjectForTestResponse struct {
	// Highlight data for the node.
	Highlight json.RawMessage `json:"highlight"`
}

// Do sends the GetHighlightObjectForTest CDP command to a browser,
// and returns the browser's response.
func (t *GetHighlightObjectForTest) Do(ctx context.Context) (*GetHighlightObjectForTestResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "GetHighlightObjectForTest", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetHighlightObjectForTestResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetGridHighlightObjectsForTest contains the parameters, and acts as
// a Go receiver, for the CDP command `getGridHighlightObjectsForTest`.
//
// For Persistent Grid testing.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-getGridHighlightObjectsForTest
type GetGridHighlightObjectsForTest struct {
	// Ids of the node to get highlight object for.
	NodeIds []dom.NodeID `json:"nodeIds"`
}

// NewGetGridHighlightObjectsForTest constructs a new GetGridHighlightObjectsForTest struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-getGridHighlightObjectsForTest
func NewGetGridHighlightObjectsForTest(nodeIds []dom.NodeID) *GetGridHighlightObjectsForTest {
	return &GetGridHighlightObjectsForTest{
		NodeIds: nodeIds,
	}
}

// GetGridHighlightObjectsForTestResponse contains the browser's response
// to calling the GetGridHighlightObjectsForTest CDP command with Do().
type GetGridHighlightObjectsForTestResponse struct {
	// Grid Highlight data for the node ids provided.
	Highlights json.RawMessage `json:"highlights"`
}

// Do sends the GetGridHighlightObjectsForTest CDP command to a browser,
// and returns the browser's response.
func (t *GetGridHighlightObjectsForTest) Do(ctx context.Context) (*GetGridHighlightObjectsForTestResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "GetGridHighlightObjectsForTest", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetGridHighlightObjectsForTestResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetSourceOrderHighlightObjectForTest contains the parameters, and acts as
// a Go receiver, for the CDP command `getSourceOrderHighlightObjectForTest`.
//
// For Source Order Viewer testing.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-getSourceOrderHighlightObjectForTest
type GetSourceOrderHighlightObjectForTest struct {
	// Id of the node to highlight.
	NodeID dom.NodeID `json:"nodeId"`
}

// NewGetSourceOrderHighlightObjectForTest constructs a new GetSourceOrderHighlightObjectForTest struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-getSourceOrderHighlightObjectForTest
func NewGetSourceOrderHighlightObjectForTest(nodeId dom.NodeID) *GetSourceOrderHighlightObjectForTest {
	return &GetSourceOrderHighlightObjectForTest{
		NodeID: nodeId,
	}
}

// GetSourceOrderHighlightObjectForTestResponse contains the browser's response
// to calling the GetSourceOrderHighlightObjectForTest CDP command with Do().
type GetSourceOrderHighlightObjectForTestResponse struct {
	// Source order highlight data for the node id provided.
	Highlight json.RawMessage `json:"highlight"`
}

// Do sends the GetSourceOrderHighlightObjectForTest CDP command to a browser,
// and returns the browser's response.
func (t *GetSourceOrderHighlightObjectForTest) Do(ctx context.Context) (*GetSourceOrderHighlightObjectForTestResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "GetSourceOrderHighlightObjectForTest", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetSourceOrderHighlightObjectForTestResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// HideHighlight contains the parameters, and acts as
// a Go receiver, for the CDP command `hideHighlight`.
//
// Hides any highlight.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-hideHighlight
type HideHighlight struct{}

// NewHideHighlight constructs a new HideHighlight struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-hideHighlight
func NewHideHighlight() *HideHighlight {
	return &HideHighlight{}
}

// Do sends the HideHighlight CDP command to a browser,
// and returns the browser's response.
func (t *HideHighlight) Do(ctx context.Context) error {
	response, err := cdp.Send(ctx, "HideHighlight", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// HighlightFrame contains the parameters, and acts as
// a Go receiver, for the CDP command `highlightFrame`.
//
// Highlights owner element of the frame with given id.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightFrame
type HighlightFrame struct {
	// Identifier of the frame to highlight.
	FrameID cdp.FrameID `json:"frameId"`
	// The content box highlight fill color (default: transparent).
	ContentColor *dom.RGBA `json:"contentColor,omitempty"`
	// The content box highlight outline color (default: transparent).
	ContentOutlineColor *dom.RGBA `json:"contentOutlineColor,omitempty"`
}

// NewHighlightFrame constructs a new HighlightFrame struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightFrame
func NewHighlightFrame(frameId cdp.FrameID) *HighlightFrame {
	return &HighlightFrame{
		FrameID: frameId,
	}
}

// SetContentColor adds or modifies the value of the optional
// parameter `contentColor` in the HighlightFrame CDP command.
//
// The content box highlight fill color (default: transparent).
func (t *HighlightFrame) SetContentColor(v dom.RGBA) *HighlightFrame {
	t.ContentColor = &v
	return t
}

// SetContentOutlineColor adds or modifies the value of the optional
// parameter `contentOutlineColor` in the HighlightFrame CDP command.
//
// The content box highlight outline color (default: transparent).
func (t *HighlightFrame) SetContentOutlineColor(v dom.RGBA) *HighlightFrame {
	t.ContentOutlineColor = &v
	return t
}

// Do sends the HighlightFrame CDP command to a browser,
// and returns the browser's response.
func (t *HighlightFrame) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "HighlightFrame", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// HighlightNode contains the parameters, and acts as
// a Go receiver, for the CDP command `highlightNode`.
//
// Highlights DOM node with given id or with the given JavaScript object wrapper. Either nodeId or
// objectId must be specified.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightNode
type HighlightNode struct {
	// A descriptor for the highlight appearance.
	HighlightConfig HighlightConfig `json:"highlightConfig"`
	// Identifier of the node to highlight.
	NodeID *dom.NodeID `json:"nodeId,omitempty"`
	// Identifier of the backend node to highlight.
	BackendNodeID *dom.BackendNodeID `json:"backendNodeId,omitempty"`
	// JavaScript object id of the node to be highlighted.
	ObjectID *runtime.RemoteObjectID `json:"objectId,omitempty"`
	// Selectors to highlight relevant nodes.
	Selector string `json:"selector,omitempty"`
}

// NewHighlightNode constructs a new HighlightNode struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightNode
func NewHighlightNode(highlightConfig HighlightConfig) *HighlightNode {
	return &HighlightNode{
		HighlightConfig: highlightConfig,
	}
}

// SetNodeID adds or modifies the value of the optional
// parameter `nodeId` in the HighlightNode CDP command.
//
// Identifier of the node to highlight.
func (t *HighlightNode) SetNodeID(v dom.NodeID) *HighlightNode {
	t.NodeID = &v
	return t
}

// SetBackendNodeID adds or modifies the value of the optional
// parameter `backendNodeId` in the HighlightNode CDP command.
//
// Identifier of the backend node to highlight.
func (t *HighlightNode) SetBackendNodeID(v dom.BackendNodeID) *HighlightNode {
	t.BackendNodeID = &v
	return t
}

// SetObjectID adds or modifies the value of the optional
// parameter `objectId` in the HighlightNode CDP command.
//
// JavaScript object id of the node to be highlighted.
func (t *HighlightNode) SetObjectID(v runtime.RemoteObjectID) *HighlightNode {
	t.ObjectID = &v
	return t
}

// SetSelector adds or modifies the value of the optional
// parameter `selector` in the HighlightNode CDP command.
//
// Selectors to highlight relevant nodes.
func (t *HighlightNode) SetSelector(v string) *HighlightNode {
	t.Selector = v
	return t
}

// Do sends the HighlightNode CDP command to a browser,
// and returns the browser's response.
func (t *HighlightNode) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "HighlightNode", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// HighlightQuad contains the parameters, and acts as
// a Go receiver, for the CDP command `highlightQuad`.
//
// Highlights given quad. Coordinates are absolute with respect to the main frame viewport.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightQuad
type HighlightQuad struct {
	// Quad to highlight
	Quad dom.Quad `json:"quad"`
	// The highlight fill color (default: transparent).
	Color *dom.RGBA `json:"color,omitempty"`
	// The highlight outline color (default: transparent).
	OutlineColor *dom.RGBA `json:"outlineColor,omitempty"`
}

// NewHighlightQuad constructs a new HighlightQuad struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightQuad
func NewHighlightQuad(quad dom.Quad) *HighlightQuad {
	return &HighlightQuad{
		Quad: quad,
	}
}

// SetColor adds or modifies the value of the optional
// parameter `color` in the HighlightQuad CDP command.
//
// The highlight fill color (default: transparent).
func (t *HighlightQuad) SetColor(v dom.RGBA) *HighlightQuad {
	t.Color = &v
	return t
}

// SetOutlineColor adds or modifies the value of the optional
// parameter `outlineColor` in the HighlightQuad CDP command.
//
// The highlight outline color (default: transparent).
func (t *HighlightQuad) SetOutlineColor(v dom.RGBA) *HighlightQuad {
	t.OutlineColor = &v
	return t
}

// Do sends the HighlightQuad CDP command to a browser,
// and returns the browser's response.
func (t *HighlightQuad) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "HighlightQuad", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// HighlightRect contains the parameters, and acts as
// a Go receiver, for the CDP command `highlightRect`.
//
// Highlights given rectangle. Coordinates are absolute with respect to the main frame viewport.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightRect
type HighlightRect struct {
	// X coordinate
	X int64 `json:"x"`
	// Y coordinate
	Y int64 `json:"y"`
	// Rectangle width
	Width int64 `json:"width"`
	// Rectangle height
	Height int64 `json:"height"`
	// The highlight fill color (default: transparent).
	Color *dom.RGBA `json:"color,omitempty"`
	// The highlight outline color (default: transparent).
	OutlineColor *dom.RGBA `json:"outlineColor,omitempty"`
}

// NewHighlightRect constructs a new HighlightRect struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightRect
func NewHighlightRect(x int64, y int64, width int64, height int64) *HighlightRect {
	return &HighlightRect{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}

// SetColor adds or modifies the value of the optional
// parameter `color` in the HighlightRect CDP command.
//
// The highlight fill color (default: transparent).
func (t *HighlightRect) SetColor(v dom.RGBA) *HighlightRect {
	t.Color = &v
	return t
}

// SetOutlineColor adds or modifies the value of the optional
// parameter `outlineColor` in the HighlightRect CDP command.
//
// The highlight outline color (default: transparent).
func (t *HighlightRect) SetOutlineColor(v dom.RGBA) *HighlightRect {
	t.OutlineColor = &v
	return t
}

// Do sends the HighlightRect CDP command to a browser,
// and returns the browser's response.
func (t *HighlightRect) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "HighlightRect", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// HighlightSourceOrder contains the parameters, and acts as
// a Go receiver, for the CDP command `highlightSourceOrder`.
//
// Highlights the source order of the children of the DOM node with given id or with the given
// JavaScript object wrapper. Either nodeId or objectId must be specified.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightSourceOrder
type HighlightSourceOrder struct {
	// A descriptor for the appearance of the overlay drawing.
	SourceOrderConfig SourceOrderConfig `json:"sourceOrderConfig"`
	// Identifier of the node to highlight.
	NodeID *dom.NodeID `json:"nodeId,omitempty"`
	// Identifier of the backend node to highlight.
	BackendNodeID *dom.BackendNodeID `json:"backendNodeId,omitempty"`
	// JavaScript object id of the node to be highlighted.
	ObjectID *runtime.RemoteObjectID `json:"objectId,omitempty"`
}

// NewHighlightSourceOrder constructs a new HighlightSourceOrder struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightSourceOrder
func NewHighlightSourceOrder(sourceOrderConfig SourceOrderConfig) *HighlightSourceOrder {
	return &HighlightSourceOrder{
		SourceOrderConfig: sourceOrderConfig,
	}
}

// SetNodeID adds or modifies the value of the optional
// parameter `nodeId` in the HighlightSourceOrder CDP command.
//
// Identifier of the node to highlight.
func (t *HighlightSourceOrder) SetNodeID(v dom.NodeID) *HighlightSourceOrder {
	t.NodeID = &v
	return t
}

// SetBackendNodeID adds or modifies the value of the optional
// parameter `backendNodeId` in the HighlightSourceOrder CDP command.
//
// Identifier of the backend node to highlight.
func (t *HighlightSourceOrder) SetBackendNodeID(v dom.BackendNodeID) *HighlightSourceOrder {
	t.BackendNodeID = &v
	return t
}

// SetObjectID adds or modifies the value of the optional
// parameter `objectId` in the HighlightSourceOrder CDP command.
//
// JavaScript object id of the node to be highlighted.
func (t *HighlightSourceOrder) SetObjectID(v runtime.RemoteObjectID) *HighlightSourceOrder {
	t.ObjectID = &v
	return t
}

// Do sends the HighlightSourceOrder CDP command to a browser,
// and returns the browser's response.
func (t *HighlightSourceOrder) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "HighlightSourceOrder", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetInspectMode contains the parameters, and acts as
// a Go receiver, for the CDP command `setInspectMode`.
//
// Enters the 'inspect' mode. In this mode, elements that user is hovering over are highlighted.
// Backend then generates 'inspectNodeRequested' event upon element selection.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setInspectMode
type SetInspectMode struct {
	// Set an inspection mode.
	Mode InspectMode `json:"mode"`
	// A descriptor for the highlight appearance of hovered-over nodes. May be omitted if `enabled
	// == false`.
	HighlightConfig *HighlightConfig `json:"highlightConfig,omitempty"`
}

// NewSetInspectMode constructs a new SetInspectMode struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setInspectMode
func NewSetInspectMode(mode InspectMode) *SetInspectMode {
	return &SetInspectMode{
		Mode: mode,
	}
}

// SetHighlightConfig adds or modifies the value of the optional
// parameter `highlightConfig` in the SetInspectMode CDP command.
//
// A descriptor for the highlight appearance of hovered-over nodes. May be omitted if `enabled
// == false`.
func (t *SetInspectMode) SetHighlightConfig(v HighlightConfig) *SetInspectMode {
	t.HighlightConfig = &v
	return t
}

// Do sends the SetInspectMode CDP command to a browser,
// and returns the browser's response.
func (t *SetInspectMode) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "SetInspectMode", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetShowAdHighlights contains the parameters, and acts as
// a Go receiver, for the CDP command `setShowAdHighlights`.
//
// Highlights owner element of all frames detected to be ads.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowAdHighlights
type SetShowAdHighlights struct {
	// True for showing ad highlights
	Show bool `json:"show"`
}

// NewSetShowAdHighlights constructs a new SetShowAdHighlights struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowAdHighlights
func NewSetShowAdHighlights(show bool) *SetShowAdHighlights {
	return &SetShowAdHighlights{
		Show: show,
	}
}

// Do sends the SetShowAdHighlights CDP command to a browser,
// and returns the browser's response.
func (t *SetShowAdHighlights) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "SetShowAdHighlights", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetPausedInDebuggerMessage contains the parameters, and acts as
// a Go receiver, for the CDP command `setPausedInDebuggerMessage`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setPausedInDebuggerMessage
type SetPausedInDebuggerMessage struct {
	// The message to display, also triggers resume and step over controls.
	Message string `json:"message,omitempty"`
}

// NewSetPausedInDebuggerMessage constructs a new SetPausedInDebuggerMessage struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setPausedInDebuggerMessage
func NewSetPausedInDebuggerMessage() *SetPausedInDebuggerMessage {
	return &SetPausedInDebuggerMessage{}
}

// SetMessage adds or modifies the value of the optional
// parameter `message` in the SetPausedInDebuggerMessage CDP command.
//
// The message to display, also triggers resume and step over controls.
func (t *SetPausedInDebuggerMessage) SetMessage(v string) *SetPausedInDebuggerMessage {
	t.Message = v
	return t
}

// Do sends the SetPausedInDebuggerMessage CDP command to a browser,
// and returns the browser's response.
func (t *SetPausedInDebuggerMessage) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "SetPausedInDebuggerMessage", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetShowDebugBorders contains the parameters, and acts as
// a Go receiver, for the CDP command `setShowDebugBorders`.
//
// Requests that backend shows debug borders on layers
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowDebugBorders
type SetShowDebugBorders struct {
	// True for showing debug borders
	Show bool `json:"show"`
}

// NewSetShowDebugBorders constructs a new SetShowDebugBorders struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowDebugBorders
func NewSetShowDebugBorders(show bool) *SetShowDebugBorders {
	return &SetShowDebugBorders{
		Show: show,
	}
}

// Do sends the SetShowDebugBorders CDP command to a browser,
// and returns the browser's response.
func (t *SetShowDebugBorders) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "SetShowDebugBorders", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetShowFPSCounter contains the parameters, and acts as
// a Go receiver, for the CDP command `setShowFPSCounter`.
//
// Requests that backend shows the FPS counter
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowFPSCounter
type SetShowFPSCounter struct {
	// True for showing the FPS counter
	Show bool `json:"show"`
}

// NewSetShowFPSCounter constructs a new SetShowFPSCounter struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowFPSCounter
func NewSetShowFPSCounter(show bool) *SetShowFPSCounter {
	return &SetShowFPSCounter{
		Show: show,
	}
}

// Do sends the SetShowFPSCounter CDP command to a browser,
// and returns the browser's response.
func (t *SetShowFPSCounter) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "SetShowFPSCounter", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetShowGridOverlays contains the parameters, and acts as
// a Go receiver, for the CDP command `setShowGridOverlays`.
//
// Highlight multiple elements with the CSS Grid overlay.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowGridOverlays
type SetShowGridOverlays struct {
	// An array of node identifiers and descriptors for the highlight appearance.
	GridNodeHighlightConfigs []GridNodeHighlightConfig `json:"gridNodeHighlightConfigs"`
}

// NewSetShowGridOverlays constructs a new SetShowGridOverlays struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowGridOverlays
func NewSetShowGridOverlays(gridNodeHighlightConfigs []GridNodeHighlightConfig) *SetShowGridOverlays {
	return &SetShowGridOverlays{
		GridNodeHighlightConfigs: gridNodeHighlightConfigs,
	}
}

// Do sends the SetShowGridOverlays CDP command to a browser,
// and returns the browser's response.
func (t *SetShowGridOverlays) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "SetShowGridOverlays", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetShowFlexOverlays contains the parameters, and acts as
// a Go receiver, for the CDP command `setShowFlexOverlays`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowFlexOverlays
type SetShowFlexOverlays struct {
	// An array of node identifiers and descriptors for the highlight appearance.
	FlexNodeHighlightConfigs []FlexNodeHighlightConfig `json:"flexNodeHighlightConfigs"`
}

// NewSetShowFlexOverlays constructs a new SetShowFlexOverlays struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowFlexOverlays
func NewSetShowFlexOverlays(flexNodeHighlightConfigs []FlexNodeHighlightConfig) *SetShowFlexOverlays {
	return &SetShowFlexOverlays{
		FlexNodeHighlightConfigs: flexNodeHighlightConfigs,
	}
}

// Do sends the SetShowFlexOverlays CDP command to a browser,
// and returns the browser's response.
func (t *SetShowFlexOverlays) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "SetShowFlexOverlays", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetShowScrollSnapOverlays contains the parameters, and acts as
// a Go receiver, for the CDP command `setShowScrollSnapOverlays`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowScrollSnapOverlays
type SetShowScrollSnapOverlays struct {
	// An array of node identifiers and descriptors for the highlight appearance.
	ScrollSnapHighlightConfigs []ScrollSnapHighlightConfig `json:"scrollSnapHighlightConfigs"`
}

// NewSetShowScrollSnapOverlays constructs a new SetShowScrollSnapOverlays struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowScrollSnapOverlays
func NewSetShowScrollSnapOverlays(scrollSnapHighlightConfigs []ScrollSnapHighlightConfig) *SetShowScrollSnapOverlays {
	return &SetShowScrollSnapOverlays{
		ScrollSnapHighlightConfigs: scrollSnapHighlightConfigs,
	}
}

// Do sends the SetShowScrollSnapOverlays CDP command to a browser,
// and returns the browser's response.
func (t *SetShowScrollSnapOverlays) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "SetShowScrollSnapOverlays", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetShowPaintRects contains the parameters, and acts as
// a Go receiver, for the CDP command `setShowPaintRects`.
//
// Requests that backend shows paint rectangles
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowPaintRects
type SetShowPaintRects struct {
	// True for showing paint rectangles
	Result bool `json:"result"`
}

// NewSetShowPaintRects constructs a new SetShowPaintRects struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowPaintRects
func NewSetShowPaintRects(result bool) *SetShowPaintRects {
	return &SetShowPaintRects{
		Result: result,
	}
}

// Do sends the SetShowPaintRects CDP command to a browser,
// and returns the browser's response.
func (t *SetShowPaintRects) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "SetShowPaintRects", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetShowLayoutShiftRegions contains the parameters, and acts as
// a Go receiver, for the CDP command `setShowLayoutShiftRegions`.
//
// Requests that backend shows layout shift regions
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowLayoutShiftRegions
type SetShowLayoutShiftRegions struct {
	// True for showing layout shift regions
	Result bool `json:"result"`
}

// NewSetShowLayoutShiftRegions constructs a new SetShowLayoutShiftRegions struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowLayoutShiftRegions
func NewSetShowLayoutShiftRegions(result bool) *SetShowLayoutShiftRegions {
	return &SetShowLayoutShiftRegions{
		Result: result,
	}
}

// Do sends the SetShowLayoutShiftRegions CDP command to a browser,
// and returns the browser's response.
func (t *SetShowLayoutShiftRegions) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "SetShowLayoutShiftRegions", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetShowScrollBottleneckRects contains the parameters, and acts as
// a Go receiver, for the CDP command `setShowScrollBottleneckRects`.
//
// Requests that backend shows scroll bottleneck rects
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowScrollBottleneckRects
type SetShowScrollBottleneckRects struct {
	// True for showing scroll bottleneck rects
	Show bool `json:"show"`
}

// NewSetShowScrollBottleneckRects constructs a new SetShowScrollBottleneckRects struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowScrollBottleneckRects
func NewSetShowScrollBottleneckRects(show bool) *SetShowScrollBottleneckRects {
	return &SetShowScrollBottleneckRects{
		Show: show,
	}
}

// Do sends the SetShowScrollBottleneckRects CDP command to a browser,
// and returns the browser's response.
func (t *SetShowScrollBottleneckRects) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "SetShowScrollBottleneckRects", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetShowHitTestBorders contains the parameters, and acts as
// a Go receiver, for the CDP command `setShowHitTestBorders`.
//
// Requests that backend shows hit-test borders on layers
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowHitTestBorders
type SetShowHitTestBorders struct {
	// True for showing hit-test borders
	Show bool `json:"show"`
}

// NewSetShowHitTestBorders constructs a new SetShowHitTestBorders struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowHitTestBorders
func NewSetShowHitTestBorders(show bool) *SetShowHitTestBorders {
	return &SetShowHitTestBorders{
		Show: show,
	}
}

// Do sends the SetShowHitTestBorders CDP command to a browser,
// and returns the browser's response.
func (t *SetShowHitTestBorders) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "SetShowHitTestBorders", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetShowWebVitals contains the parameters, and acts as
// a Go receiver, for the CDP command `setShowWebVitals`.
//
// Request that backend shows an overlay with web vital metrics.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowWebVitals
type SetShowWebVitals struct {
	Show bool `json:"show"`
}

// NewSetShowWebVitals constructs a new SetShowWebVitals struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowWebVitals
func NewSetShowWebVitals(show bool) *SetShowWebVitals {
	return &SetShowWebVitals{
		Show: show,
	}
}

// Do sends the SetShowWebVitals CDP command to a browser,
// and returns the browser's response.
func (t *SetShowWebVitals) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "SetShowWebVitals", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetShowViewportSizeOnResize contains the parameters, and acts as
// a Go receiver, for the CDP command `setShowViewportSizeOnResize`.
//
// Paints viewport size upon main frame resize.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowViewportSizeOnResize
type SetShowViewportSizeOnResize struct {
	// Whether to paint size or not.
	Show bool `json:"show"`
}

// NewSetShowViewportSizeOnResize constructs a new SetShowViewportSizeOnResize struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowViewportSizeOnResize
func NewSetShowViewportSizeOnResize(show bool) *SetShowViewportSizeOnResize {
	return &SetShowViewportSizeOnResize{
		Show: show,
	}
}

// Do sends the SetShowViewportSizeOnResize CDP command to a browser,
// and returns the browser's response.
func (t *SetShowViewportSizeOnResize) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "SetShowViewportSizeOnResize", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetShowHinge contains the parameters, and acts as
// a Go receiver, for the CDP command `setShowHinge`.
//
// Add a dual screen device hinge
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowHinge
type SetShowHinge struct {
	// hinge data, null means hideHinge
	HingeConfig *HingeConfig `json:"hingeConfig,omitempty"`
}

// NewSetShowHinge constructs a new SetShowHinge struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowHinge
func NewSetShowHinge() *SetShowHinge {
	return &SetShowHinge{}
}

// SetHingeConfig adds or modifies the value of the optional
// parameter `hingeConfig` in the SetShowHinge CDP command.
//
// hinge data, null means hideHinge
func (t *SetShowHinge) SetHingeConfig(v HingeConfig) *SetShowHinge {
	t.HingeConfig = &v
	return t
}

// Do sends the SetShowHinge CDP command to a browser,
// and returns the browser's response.
func (t *SetShowHinge) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "SetShowHinge", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}
