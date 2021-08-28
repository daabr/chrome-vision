package accessibility

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
	"github.com/daabr/chrome-vision/pkg/devtools/runtime"
)

// Disable contains the parameters, and acts as
// a Go receiver, for the CDP command `disable`.
//
// Disables the accessibility domain.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#method-disable
type Disable struct{}

// NewDisable constructs a new Disable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#method-disable
func NewDisable() *Disable {
	return &Disable{}
}

// Do sends the Disable CDP command to a browser,
// and returns the browser's response.
func (t *Disable) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Accessibility.disable", nil)
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
	return devtools.Send(ctx, "Accessibility.disable", nil)
}

// ParseResponse parses the browser's response
// to the Disable CDP command.
func (t *Disable) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// Enable contains the parameters, and acts as
// a Go receiver, for the CDP command `enable`.
//
// Enables the accessibility domain which causes `AXNodeId`s to remain consistent between method calls.
// This turns on accessibility for the page, which can impact performance until accessibility is disabled.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#method-enable
type Enable struct{}

// NewEnable constructs a new Enable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#method-enable
func NewEnable() *Enable {
	return &Enable{}
}

// Do sends the Enable CDP command to a browser,
// and returns the browser's response.
func (t *Enable) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Accessibility.enable", nil)
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
	return devtools.Send(ctx, "Accessibility.enable", nil)
}

// ParseResponse parses the browser's response
// to the Enable CDP command.
func (t *Enable) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// GetPartialAXTree contains the parameters, and acts as
// a Go receiver, for the CDP command `getPartialAXTree`.
//
// Fetches the accessibility node and partial accessibility tree for this DOM node, if it exists.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#method-getPartialAXTree
//
// This CDP method is experimental.
type GetPartialAXTree struct {
	// Identifier of the node to get the partial accessibility tree for.
	NodeID int64 `json:"nodeId,omitempty"`
	// Identifier of the backend node to get the partial accessibility tree for.
	BackendNodeID int64 `json:"backendNodeId,omitempty"`
	// JavaScript object id of the node wrapper to get the partial accessibility tree for.
	ObjectID *runtime.RemoteObjectID `json:"objectId,omitempty"`
	// Whether to fetch this nodes ancestors, siblings and children. Defaults to true.
	FetchRelatives bool `json:"fetchRelatives,omitempty"`
}

// NewGetPartialAXTree constructs a new GetPartialAXTree struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#method-getPartialAXTree
//
// This CDP method is experimental.
func NewGetPartialAXTree() *GetPartialAXTree {
	return &GetPartialAXTree{}
}

// SetNodeID adds or modifies the value of the optional
// parameter `nodeId` in the GetPartialAXTree CDP command.
//
// Identifier of the node to get the partial accessibility tree for.
func (t *GetPartialAXTree) SetNodeID(v int64) *GetPartialAXTree {
	t.NodeID = v
	return t
}

// SetBackendNodeID adds or modifies the value of the optional
// parameter `backendNodeId` in the GetPartialAXTree CDP command.
//
// Identifier of the backend node to get the partial accessibility tree for.
func (t *GetPartialAXTree) SetBackendNodeID(v int64) *GetPartialAXTree {
	t.BackendNodeID = v
	return t
}

// SetObjectID adds or modifies the value of the optional
// parameter `objectId` in the GetPartialAXTree CDP command.
//
// JavaScript object id of the node wrapper to get the partial accessibility tree for.
func (t *GetPartialAXTree) SetObjectID(v runtime.RemoteObjectID) *GetPartialAXTree {
	t.ObjectID = &v
	return t
}

// SetFetchRelatives adds or modifies the value of the optional
// parameter `fetchRelatives` in the GetPartialAXTree CDP command.
//
// Whether to fetch this nodes ancestors, siblings and children. Defaults to true.
func (t *GetPartialAXTree) SetFetchRelatives(v bool) *GetPartialAXTree {
	t.FetchRelatives = v
	return t
}

// GetPartialAXTreeResult contains the browser's response
// to calling the GetPartialAXTree CDP command with Do().
type GetPartialAXTreeResult struct {
	// The `Accessibility.AXNode` for this DOM node, if it exists, plus its ancestors, siblings and
	// children, if requested.
	Nodes []AXNode `json:"nodes"`
}

// Do sends the GetPartialAXTree CDP command to a browser,
// and returns the browser's response.
func (t *GetPartialAXTree) Do(ctx context.Context) (*GetPartialAXTreeResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Accessibility.getPartialAXTree", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetPartialAXTree CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetPartialAXTree) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Accessibility.getPartialAXTree", b)
}

// ParseResponse parses the browser's response
// to the GetPartialAXTree CDP command.
func (t *GetPartialAXTree) ParseResponse(m *devtools.Message) (*GetPartialAXTreeResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetPartialAXTreeResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetFullAXTree contains the parameters, and acts as
// a Go receiver, for the CDP command `getFullAXTree`.
//
// Fetches the entire accessibility tree for the root Document
//
// https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#method-getFullAXTree
//
// This CDP method is experimental.
type GetFullAXTree struct {
	// The maximum depth at which descendants of the root node should be retrieved.
	// If omitted, the full tree is returned.
	Depth int64 `json:"depth,omitempty"`
	// Deprecated. This parameter has been renamed to `depth`. If depth is not provided, max_depth will be used.
	//
	// This CDP parameter is deprecated.
	MaxDepth int64 `json:"max_depth,omitempty"`
	// The frame for whose document the AX tree should be retrieved.
	// If omited, the root frame is used.
	FrameID string `json:"frameId,omitempty"`
}

// NewGetFullAXTree constructs a new GetFullAXTree struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#method-getFullAXTree
//
// This CDP method is experimental.
func NewGetFullAXTree() *GetFullAXTree {
	return &GetFullAXTree{}
}

// SetDepth adds or modifies the value of the optional
// parameter `depth` in the GetFullAXTree CDP command.
//
// The maximum depth at which descendants of the root node should be retrieved.
// If omitted, the full tree is returned.
func (t *GetFullAXTree) SetDepth(v int64) *GetFullAXTree {
	t.Depth = v
	return t
}

// SetMaxDepth adds or modifies the value of the optional
// parameter `max_depth` in the GetFullAXTree CDP command.
//
// Deprecated. This parameter has been renamed to `depth`. If depth is not provided, max_depth will be used.
//
// This CDP parameter is deprecated.
func (t *GetFullAXTree) SetMaxDepth(v int64) *GetFullAXTree {
	t.MaxDepth = v
	return t
}

// SetFrameID adds or modifies the value of the optional
// parameter `frameId` in the GetFullAXTree CDP command.
//
// The frame for whose document the AX tree should be retrieved.
// If omited, the root frame is used.
func (t *GetFullAXTree) SetFrameID(v string) *GetFullAXTree {
	t.FrameID = v
	return t
}

// GetFullAXTreeResult contains the browser's response
// to calling the GetFullAXTree CDP command with Do().
type GetFullAXTreeResult struct {
	Nodes []AXNode `json:"nodes"`
}

// Do sends the GetFullAXTree CDP command to a browser,
// and returns the browser's response.
func (t *GetFullAXTree) Do(ctx context.Context) (*GetFullAXTreeResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Accessibility.getFullAXTree", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetFullAXTree CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetFullAXTree) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Accessibility.getFullAXTree", b)
}

// ParseResponse parses the browser's response
// to the GetFullAXTree CDP command.
func (t *GetFullAXTree) ParseResponse(m *devtools.Message) (*GetFullAXTreeResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetFullAXTreeResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetChildAXNodes contains the parameters, and acts as
// a Go receiver, for the CDP command `getChildAXNodes`.
//
// Fetches a particular accessibility node by AXNodeId.
// Requires `enable()` to have been called previously.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#method-getChildAXNodes
//
// This CDP method is experimental.
type GetChildAXNodes struct {
	ID string `json:"id"`
	// The frame in whose document the node resides.
	// If omitted, the root frame is used.
	FrameID string `json:"frameId,omitempty"`
}

// NewGetChildAXNodes constructs a new GetChildAXNodes struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#method-getChildAXNodes
//
// This CDP method is experimental.
func NewGetChildAXNodes(id string) *GetChildAXNodes {
	return &GetChildAXNodes{
		ID: id,
	}
}

// SetFrameID adds or modifies the value of the optional
// parameter `frameId` in the GetChildAXNodes CDP command.
//
// The frame in whose document the node resides.
// If omitted, the root frame is used.
func (t *GetChildAXNodes) SetFrameID(v string) *GetChildAXNodes {
	t.FrameID = v
	return t
}

// GetChildAXNodesResult contains the browser's response
// to calling the GetChildAXNodes CDP command with Do().
type GetChildAXNodesResult struct {
	Nodes []AXNode `json:"nodes"`
}

// Do sends the GetChildAXNodes CDP command to a browser,
// and returns the browser's response.
func (t *GetChildAXNodes) Do(ctx context.Context) (*GetChildAXNodesResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Accessibility.getChildAXNodes", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetChildAXNodes CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetChildAXNodes) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Accessibility.getChildAXNodes", b)
}

// ParseResponse parses the browser's response
// to the GetChildAXNodes CDP command.
func (t *GetChildAXNodes) ParseResponse(m *devtools.Message) (*GetChildAXNodesResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetChildAXNodesResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// QueryAXTree contains the parameters, and acts as
// a Go receiver, for the CDP command `queryAXTree`.
//
// Query a DOM node's accessibility subtree for accessible name and role.
// This command computes the name and role for all nodes in the subtree, including those that are
// ignored for accessibility, and returns those that mactch the specified name and role. If no DOM
// node is specified, or the DOM node does not exist, the command returns an error. If neither
// `accessibleName` or `role` is specified, it returns all the accessibility nodes in the subtree.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#method-queryAXTree
//
// This CDP method is experimental.
type QueryAXTree struct {
	// Identifier of the node for the root to query.
	NodeID int64 `json:"nodeId,omitempty"`
	// Identifier of the backend node for the root to query.
	BackendNodeID int64 `json:"backendNodeId,omitempty"`
	// JavaScript object id of the node wrapper for the root to query.
	ObjectID *runtime.RemoteObjectID `json:"objectId,omitempty"`
	// Find nodes with this computed name.
	AccessibleName string `json:"accessibleName,omitempty"`
	// Find nodes with this computed role.
	Role string `json:"role,omitempty"`
}

// NewQueryAXTree constructs a new QueryAXTree struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#method-queryAXTree
//
// This CDP method is experimental.
func NewQueryAXTree() *QueryAXTree {
	return &QueryAXTree{}
}

// SetNodeID adds or modifies the value of the optional
// parameter `nodeId` in the QueryAXTree CDP command.
//
// Identifier of the node for the root to query.
func (t *QueryAXTree) SetNodeID(v int64) *QueryAXTree {
	t.NodeID = v
	return t
}

// SetBackendNodeID adds or modifies the value of the optional
// parameter `backendNodeId` in the QueryAXTree CDP command.
//
// Identifier of the backend node for the root to query.
func (t *QueryAXTree) SetBackendNodeID(v int64) *QueryAXTree {
	t.BackendNodeID = v
	return t
}

// SetObjectID adds or modifies the value of the optional
// parameter `objectId` in the QueryAXTree CDP command.
//
// JavaScript object id of the node wrapper for the root to query.
func (t *QueryAXTree) SetObjectID(v runtime.RemoteObjectID) *QueryAXTree {
	t.ObjectID = &v
	return t
}

// SetAccessibleName adds or modifies the value of the optional
// parameter `accessibleName` in the QueryAXTree CDP command.
//
// Find nodes with this computed name.
func (t *QueryAXTree) SetAccessibleName(v string) *QueryAXTree {
	t.AccessibleName = v
	return t
}

// SetRole adds or modifies the value of the optional
// parameter `role` in the QueryAXTree CDP command.
//
// Find nodes with this computed role.
func (t *QueryAXTree) SetRole(v string) *QueryAXTree {
	t.Role = v
	return t
}

// QueryAXTreeResult contains the browser's response
// to calling the QueryAXTree CDP command with Do().
type QueryAXTreeResult struct {
	// A list of `Accessibility.AXNode` matching the specified attributes,
	// including nodes that are ignored for accessibility.
	Nodes []AXNode `json:"nodes"`
}

// Do sends the QueryAXTree CDP command to a browser,
// and returns the browser's response.
func (t *QueryAXTree) Do(ctx context.Context) (*QueryAXTreeResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Accessibility.queryAXTree", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the QueryAXTree CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *QueryAXTree) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Accessibility.queryAXTree", b)
}

// ParseResponse parses the browser's response
// to the QueryAXTree CDP command.
func (t *QueryAXTree) ParseResponse(m *devtools.Message) (*QueryAXTreeResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &QueryAXTreeResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}
