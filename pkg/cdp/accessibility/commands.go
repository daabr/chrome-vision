package accessibility

import (
	"context"
	"fmt"

	"github.com/daabr/chrome-vision/pkg/cdp/dom"
	"github.com/daabr/chrome-vision/pkg/cdp/runtime"
)

// Disable contains the parameters, and acts as
// a Go receiver, for the CDP command `disable`.
//
// Disables the accessibility domain.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#method-disable
type Disable struct{}

// NewDisable constructs a new Disable struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#method-disable
func NewDisable() *Disable {
	return &Disable{}
}

// Do sends the Disable CDP command to a browser,
// and returns the browser's response.
func (t *Disable) Do(ctx context.Context) error {
	fmt.Println(ctx)
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
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#method-enable
func NewEnable() *Enable {
	return &Enable{}
}

// Do sends the Enable CDP command to a browser,
// and returns the browser's response.
func (t *Enable) Do(ctx context.Context) error {
	fmt.Println(ctx)
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
	NodeID *dom.NodeID `json:"nodeId,omitempty"`
	// Identifier of the backend node to get the partial accessibility tree for.
	BackendNodeID *dom.BackendNodeID `json:"backendNodeId,omitempty"`
	// JavaScript object id of the node wrapper to get the partial accessibility tree for.
	ObjectID *runtime.RemoteObjectID `json:"objectId,omitempty"`
	// Whether to fetch this nodes ancestors, siblings and children. Defaults to true.
	FetchRelatives bool `json:"fetchRelatives,omitempty"`
}

// NewGetPartialAXTree constructs a new GetPartialAXTree struct instance, with
// all the required parameters, and only them. Optional parameters
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
func (t *GetPartialAXTree) SetNodeID(v dom.NodeID) *GetPartialAXTree {
	t.NodeID = &v
	return t
}

// SetBackendNodeID adds or modifies the value of the optional
// parameter `backendNodeId` in the GetPartialAXTree CDP command.
//
// Identifier of the backend node to get the partial accessibility tree for.
func (t *GetPartialAXTree) SetBackendNodeID(v dom.BackendNodeID) *GetPartialAXTree {
	t.BackendNodeID = &v
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

// GetPartialAXTreeResponse contains the browser's response
// to calling the GetPartialAXTree CDP command with Do().
type GetPartialAXTreeResponse struct {
	// The `Accessibility.AXNode` for this DOM node, if it exists, plus its ancestors, siblings and
	// children, if requested.
	Nodes []AXNode `json:"nodes"`
}

// Do sends the GetPartialAXTree CDP command to a browser,
// and returns the browser's response.
func (t *GetPartialAXTree) Do(ctx context.Context) (*GetPartialAXTreeResponse, error) {
	fmt.Println(ctx)
	return new(GetPartialAXTreeResponse), nil
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
	MaxDepth int64 `json:"max_depth,omitempty"`
}

// NewGetFullAXTree constructs a new GetFullAXTree struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#method-getFullAXTree
//
// This CDP method is experimental.
func NewGetFullAXTree() *GetFullAXTree {
	return &GetFullAXTree{}
}

// SetMaxDepth adds or modifies the value of the optional
// parameter `max_depth` in the GetFullAXTree CDP command.
//
// The maximum depth at which descendants of the root node should be retrieved.
// If omitted, the full tree is returned.
func (t *GetFullAXTree) SetMaxDepth(v int64) *GetFullAXTree {
	t.MaxDepth = v
	return t
}

// GetFullAXTreeResponse contains the browser's response
// to calling the GetFullAXTree CDP command with Do().
type GetFullAXTreeResponse struct {
	Nodes []AXNode `json:"nodes"`
}

// Do sends the GetFullAXTree CDP command to a browser,
// and returns the browser's response.
func (t *GetFullAXTree) Do(ctx context.Context) (*GetFullAXTreeResponse, error) {
	fmt.Println(ctx)
	return new(GetFullAXTreeResponse), nil
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
	ID AXNodeID `json:"id"`
}

// NewGetChildAXNodes constructs a new GetChildAXNodes struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#method-getChildAXNodes
//
// This CDP method is experimental.
func NewGetChildAXNodes(id AXNodeID) *GetChildAXNodes {
	return &GetChildAXNodes{
		ID: id,
	}
}

// GetChildAXNodesResponse contains the browser's response
// to calling the GetChildAXNodes CDP command with Do().
type GetChildAXNodesResponse struct {
	Nodes []AXNode `json:"nodes"`
}

// Do sends the GetChildAXNodes CDP command to a browser,
// and returns the browser's response.
func (t *GetChildAXNodes) Do(ctx context.Context) (*GetChildAXNodesResponse, error) {
	fmt.Println(ctx)
	return new(GetChildAXNodesResponse), nil
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
	NodeID *dom.NodeID `json:"nodeId,omitempty"`
	// Identifier of the backend node for the root to query.
	BackendNodeID *dom.BackendNodeID `json:"backendNodeId,omitempty"`
	// JavaScript object id of the node wrapper for the root to query.
	ObjectID *runtime.RemoteObjectID `json:"objectId,omitempty"`
	// Find nodes with this computed name.
	AccessibleName string `json:"accessibleName,omitempty"`
	// Find nodes with this computed role.
	Role string `json:"role,omitempty"`
}

// NewQueryAXTree constructs a new QueryAXTree struct instance, with
// all the required parameters, and only them. Optional parameters
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
func (t *QueryAXTree) SetNodeID(v dom.NodeID) *QueryAXTree {
	t.NodeID = &v
	return t
}

// SetBackendNodeID adds or modifies the value of the optional
// parameter `backendNodeId` in the QueryAXTree CDP command.
//
// Identifier of the backend node for the root to query.
func (t *QueryAXTree) SetBackendNodeID(v dom.BackendNodeID) *QueryAXTree {
	t.BackendNodeID = &v
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

// QueryAXTreeResponse contains the browser's response
// to calling the QueryAXTree CDP command with Do().
type QueryAXTreeResponse struct {
	// A list of `Accessibility.AXNode` matching the specified attributes,
	// including nodes that are ignored for accessibility.
	Nodes []AXNode `json:"nodes"`
}

// Do sends the QueryAXTree CDP command to a browser,
// and returns the browser's response.
func (t *QueryAXTree) Do(ctx context.Context) (*QueryAXTreeResponse, error) {
	fmt.Println(ctx)
	return new(QueryAXTreeResponse), nil
}
