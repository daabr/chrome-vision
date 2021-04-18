package dom

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/cdp"
	"github.com/daabr/chrome-vision/pkg/cdp/runtime"
)

// CollectClassNamesFromSubtree contains the parameters, and acts as
// a Go receiver, for the CDP command `collectClassNamesFromSubtree`.
//
// Collects class names for the node with given id and all of it's child nodes.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-collectClassNamesFromSubtree
//
// This CDP method is experimental.
type CollectClassNamesFromSubtree struct {
	// Id of the node to collect class names.
	NodeID NodeID `json:"nodeId"`
}

// NewCollectClassNamesFromSubtree constructs a new CollectClassNamesFromSubtree struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-collectClassNamesFromSubtree
//
// This CDP method is experimental.
func NewCollectClassNamesFromSubtree(nodeId NodeID) *CollectClassNamesFromSubtree {
	return &CollectClassNamesFromSubtree{
		NodeID: nodeId,
	}
}

// CollectClassNamesFromSubtreeResponse contains the browser's response
// to calling the CollectClassNamesFromSubtree CDP command with Do().
type CollectClassNamesFromSubtreeResponse struct {
	// Class name list.
	ClassNames []string `json:"classNames"`
}

// Do sends the CollectClassNamesFromSubtree CDP command to a browser,
// and returns the browser's response.
func (t *CollectClassNamesFromSubtree) Do(ctx context.Context) (*CollectClassNamesFromSubtreeResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "CollectClassNamesFromSubtree", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &CollectClassNamesFromSubtreeResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// CopyTo contains the parameters, and acts as
// a Go receiver, for the CDP command `copyTo`.
//
// Creates a deep copy of the specified node and places it into the target container before the
// given anchor.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-copyTo
//
// This CDP method is experimental.
type CopyTo struct {
	// Id of the node to copy.
	NodeID NodeID `json:"nodeId"`
	// Id of the element to drop the copy into.
	TargetNodeID NodeID `json:"targetNodeId"`
	// Drop the copy before this node (if absent, the copy becomes the last child of
	// `targetNodeId`).
	InsertBeforeNodeID *NodeID `json:"insertBeforeNodeId,omitempty"`
}

// NewCopyTo constructs a new CopyTo struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-copyTo
//
// This CDP method is experimental.
func NewCopyTo(nodeId NodeID, targetNodeId NodeID) *CopyTo {
	return &CopyTo{
		NodeID:       nodeId,
		TargetNodeID: targetNodeId,
	}
}

// SetInsertBeforeNodeID adds or modifies the value of the optional
// parameter `insertBeforeNodeId` in the CopyTo CDP command.
//
// Drop the copy before this node (if absent, the copy becomes the last child of
// `targetNodeId`).
func (t *CopyTo) SetInsertBeforeNodeID(v NodeID) *CopyTo {
	t.InsertBeforeNodeID = &v
	return t
}

// CopyToResponse contains the browser's response
// to calling the CopyTo CDP command with Do().
type CopyToResponse struct {
	// Id of the node clone.
	NodeID NodeID `json:"nodeId"`
}

// Do sends the CopyTo CDP command to a browser,
// and returns the browser's response.
func (t *CopyTo) Do(ctx context.Context) (*CopyToResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "CopyTo", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &CopyToResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeNode contains the parameters, and acts as
// a Go receiver, for the CDP command `describeNode`.
//
// Describes node given its id, does not require domain to be enabled. Does not start tracking any
// objects, can be used for automation.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-describeNode
type DescribeNode struct {
	// Identifier of the node.
	NodeID *NodeID `json:"nodeId,omitempty"`
	// Identifier of the backend node.
	BackendNodeID *BackendNodeID `json:"backendNodeId,omitempty"`
	// JavaScript object id of the node wrapper.
	ObjectID *runtime.RemoteObjectID `json:"objectId,omitempty"`
	// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the
	// entire subtree or provide an integer larger than 0.
	Depth int64 `json:"depth,omitempty"`
	// Whether or not iframes and shadow roots should be traversed when returning the subtree
	// (default is false).
	Pierce bool `json:"pierce,omitempty"`
}

// NewDescribeNode constructs a new DescribeNode struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-describeNode
func NewDescribeNode() *DescribeNode {
	return &DescribeNode{}
}

// SetNodeID adds or modifies the value of the optional
// parameter `nodeId` in the DescribeNode CDP command.
//
// Identifier of the node.
func (t *DescribeNode) SetNodeID(v NodeID) *DescribeNode {
	t.NodeID = &v
	return t
}

// SetBackendNodeID adds or modifies the value of the optional
// parameter `backendNodeId` in the DescribeNode CDP command.
//
// Identifier of the backend node.
func (t *DescribeNode) SetBackendNodeID(v BackendNodeID) *DescribeNode {
	t.BackendNodeID = &v
	return t
}

// SetObjectID adds or modifies the value of the optional
// parameter `objectId` in the DescribeNode CDP command.
//
// JavaScript object id of the node wrapper.
func (t *DescribeNode) SetObjectID(v runtime.RemoteObjectID) *DescribeNode {
	t.ObjectID = &v
	return t
}

// SetDepth adds or modifies the value of the optional
// parameter `depth` in the DescribeNode CDP command.
//
// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the
// entire subtree or provide an integer larger than 0.
func (t *DescribeNode) SetDepth(v int64) *DescribeNode {
	t.Depth = v
	return t
}

// SetPierce adds or modifies the value of the optional
// parameter `pierce` in the DescribeNode CDP command.
//
// Whether or not iframes and shadow roots should be traversed when returning the subtree
// (default is false).
func (t *DescribeNode) SetPierce(v bool) *DescribeNode {
	t.Pierce = v
	return t
}

// DescribeNodeResponse contains the browser's response
// to calling the DescribeNode CDP command with Do().
type DescribeNodeResponse struct {
	// Node description.
	Node Node `json:"node"`
}

// Do sends the DescribeNode CDP command to a browser,
// and returns the browser's response.
func (t *DescribeNode) Do(ctx context.Context) (*DescribeNodeResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "DescribeNode", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &DescribeNodeResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ScrollIntoViewIfNeeded contains the parameters, and acts as
// a Go receiver, for the CDP command `scrollIntoViewIfNeeded`.
//
// Scrolls the specified rect of the given node into view if not already visible.
// Note: exactly one between nodeId, backendNodeId and objectId should be passed
// to identify the node.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-scrollIntoViewIfNeeded
//
// This CDP method is experimental.
type ScrollIntoViewIfNeeded struct {
	// Identifier of the node.
	NodeID *NodeID `json:"nodeId,omitempty"`
	// Identifier of the backend node.
	BackendNodeID *BackendNodeID `json:"backendNodeId,omitempty"`
	// JavaScript object id of the node wrapper.
	ObjectID *runtime.RemoteObjectID `json:"objectId,omitempty"`
	// The rect to be scrolled into view, relative to the node's border box, in CSS pixels.
	// When omitted, center of the node will be used, similar to Element.scrollIntoView.
	Rect *Rect `json:"rect,omitempty"`
}

// NewScrollIntoViewIfNeeded constructs a new ScrollIntoViewIfNeeded struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-scrollIntoViewIfNeeded
//
// This CDP method is experimental.
func NewScrollIntoViewIfNeeded() *ScrollIntoViewIfNeeded {
	return &ScrollIntoViewIfNeeded{}
}

// SetNodeID adds or modifies the value of the optional
// parameter `nodeId` in the ScrollIntoViewIfNeeded CDP command.
//
// Identifier of the node.
func (t *ScrollIntoViewIfNeeded) SetNodeID(v NodeID) *ScrollIntoViewIfNeeded {
	t.NodeID = &v
	return t
}

// SetBackendNodeID adds or modifies the value of the optional
// parameter `backendNodeId` in the ScrollIntoViewIfNeeded CDP command.
//
// Identifier of the backend node.
func (t *ScrollIntoViewIfNeeded) SetBackendNodeID(v BackendNodeID) *ScrollIntoViewIfNeeded {
	t.BackendNodeID = &v
	return t
}

// SetObjectID adds or modifies the value of the optional
// parameter `objectId` in the ScrollIntoViewIfNeeded CDP command.
//
// JavaScript object id of the node wrapper.
func (t *ScrollIntoViewIfNeeded) SetObjectID(v runtime.RemoteObjectID) *ScrollIntoViewIfNeeded {
	t.ObjectID = &v
	return t
}

// SetRect adds or modifies the value of the optional
// parameter `rect` in the ScrollIntoViewIfNeeded CDP command.
//
// The rect to be scrolled into view, relative to the node's border box, in CSS pixels.
// When omitted, center of the node will be used, similar to Element.scrollIntoView.
func (t *ScrollIntoViewIfNeeded) SetRect(v Rect) *ScrollIntoViewIfNeeded {
	t.Rect = &v
	return t
}

// Do sends the ScrollIntoViewIfNeeded CDP command to a browser,
// and returns the browser's response.
func (t *ScrollIntoViewIfNeeded) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "ScrollIntoViewIfNeeded", b)
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
// Disables DOM agent for the given page.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-disable
type Disable struct{}

// NewDisable constructs a new Disable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-disable
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

// DiscardSearchResults contains the parameters, and acts as
// a Go receiver, for the CDP command `discardSearchResults`.
//
// Discards search results from the session with the given id. `getSearchResults` should no longer
// be called for that search.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-discardSearchResults
//
// This CDP method is experimental.
type DiscardSearchResults struct {
	// Unique search session identifier.
	SearchID string `json:"searchId"`
}

// NewDiscardSearchResults constructs a new DiscardSearchResults struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-discardSearchResults
//
// This CDP method is experimental.
func NewDiscardSearchResults(searchId string) *DiscardSearchResults {
	return &DiscardSearchResults{
		SearchID: searchId,
	}
}

// Do sends the DiscardSearchResults CDP command to a browser,
// and returns the browser's response.
func (t *DiscardSearchResults) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "DiscardSearchResults", b)
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
// Enables DOM agent for the given page.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-enable
type Enable struct{}

// NewEnable constructs a new Enable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-enable
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

// Focus contains the parameters, and acts as
// a Go receiver, for the CDP command `focus`.
//
// Focuses the given element.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-focus
type Focus struct {
	// Identifier of the node.
	NodeID *NodeID `json:"nodeId,omitempty"`
	// Identifier of the backend node.
	BackendNodeID *BackendNodeID `json:"backendNodeId,omitempty"`
	// JavaScript object id of the node wrapper.
	ObjectID *runtime.RemoteObjectID `json:"objectId,omitempty"`
}

// NewFocus constructs a new Focus struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-focus
func NewFocus() *Focus {
	return &Focus{}
}

// SetNodeID adds or modifies the value of the optional
// parameter `nodeId` in the Focus CDP command.
//
// Identifier of the node.
func (t *Focus) SetNodeID(v NodeID) *Focus {
	t.NodeID = &v
	return t
}

// SetBackendNodeID adds or modifies the value of the optional
// parameter `backendNodeId` in the Focus CDP command.
//
// Identifier of the backend node.
func (t *Focus) SetBackendNodeID(v BackendNodeID) *Focus {
	t.BackendNodeID = &v
	return t
}

// SetObjectID adds or modifies the value of the optional
// parameter `objectId` in the Focus CDP command.
//
// JavaScript object id of the node wrapper.
func (t *Focus) SetObjectID(v runtime.RemoteObjectID) *Focus {
	t.ObjectID = &v
	return t
}

// Do sends the Focus CDP command to a browser,
// and returns the browser's response.
func (t *Focus) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Focus", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// GetAttributes contains the parameters, and acts as
// a Go receiver, for the CDP command `getAttributes`.
//
// Returns attributes for the specified node.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getAttributes
type GetAttributes struct {
	// Id of the node to retrieve attibutes for.
	NodeID NodeID `json:"nodeId"`
}

// NewGetAttributes constructs a new GetAttributes struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getAttributes
func NewGetAttributes(nodeId NodeID) *GetAttributes {
	return &GetAttributes{
		NodeID: nodeId,
	}
}

// GetAttributesResponse contains the browser's response
// to calling the GetAttributes CDP command with Do().
type GetAttributesResponse struct {
	// An interleaved array of node attribute names and values.
	Attributes []string `json:"attributes"`
}

// Do sends the GetAttributes CDP command to a browser,
// and returns the browser's response.
func (t *GetAttributes) Do(ctx context.Context) (*GetAttributesResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "GetAttributes", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetAttributesResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetBoxModel contains the parameters, and acts as
// a Go receiver, for the CDP command `getBoxModel`.
//
// Returns boxes for the given node.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getBoxModel
type GetBoxModel struct {
	// Identifier of the node.
	NodeID *NodeID `json:"nodeId,omitempty"`
	// Identifier of the backend node.
	BackendNodeID *BackendNodeID `json:"backendNodeId,omitempty"`
	// JavaScript object id of the node wrapper.
	ObjectID *runtime.RemoteObjectID `json:"objectId,omitempty"`
}

// NewGetBoxModel constructs a new GetBoxModel struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getBoxModel
func NewGetBoxModel() *GetBoxModel {
	return &GetBoxModel{}
}

// SetNodeID adds or modifies the value of the optional
// parameter `nodeId` in the GetBoxModel CDP command.
//
// Identifier of the node.
func (t *GetBoxModel) SetNodeID(v NodeID) *GetBoxModel {
	t.NodeID = &v
	return t
}

// SetBackendNodeID adds or modifies the value of the optional
// parameter `backendNodeId` in the GetBoxModel CDP command.
//
// Identifier of the backend node.
func (t *GetBoxModel) SetBackendNodeID(v BackendNodeID) *GetBoxModel {
	t.BackendNodeID = &v
	return t
}

// SetObjectID adds or modifies the value of the optional
// parameter `objectId` in the GetBoxModel CDP command.
//
// JavaScript object id of the node wrapper.
func (t *GetBoxModel) SetObjectID(v runtime.RemoteObjectID) *GetBoxModel {
	t.ObjectID = &v
	return t
}

// GetBoxModelResponse contains the browser's response
// to calling the GetBoxModel CDP command with Do().
type GetBoxModelResponse struct {
	// Box model for the node.
	Model BoxModel `json:"model"`
}

// Do sends the GetBoxModel CDP command to a browser,
// and returns the browser's response.
func (t *GetBoxModel) Do(ctx context.Context) (*GetBoxModelResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "GetBoxModel", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetBoxModelResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetContentQuads contains the parameters, and acts as
// a Go receiver, for the CDP command `getContentQuads`.
//
// Returns quads that describe node position on the page. This method
// might return multiple quads for inline nodes.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getContentQuads
//
// This CDP method is experimental.
type GetContentQuads struct {
	// Identifier of the node.
	NodeID *NodeID `json:"nodeId,omitempty"`
	// Identifier of the backend node.
	BackendNodeID *BackendNodeID `json:"backendNodeId,omitempty"`
	// JavaScript object id of the node wrapper.
	ObjectID *runtime.RemoteObjectID `json:"objectId,omitempty"`
}

// NewGetContentQuads constructs a new GetContentQuads struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getContentQuads
//
// This CDP method is experimental.
func NewGetContentQuads() *GetContentQuads {
	return &GetContentQuads{}
}

// SetNodeID adds or modifies the value of the optional
// parameter `nodeId` in the GetContentQuads CDP command.
//
// Identifier of the node.
func (t *GetContentQuads) SetNodeID(v NodeID) *GetContentQuads {
	t.NodeID = &v
	return t
}

// SetBackendNodeID adds or modifies the value of the optional
// parameter `backendNodeId` in the GetContentQuads CDP command.
//
// Identifier of the backend node.
func (t *GetContentQuads) SetBackendNodeID(v BackendNodeID) *GetContentQuads {
	t.BackendNodeID = &v
	return t
}

// SetObjectID adds or modifies the value of the optional
// parameter `objectId` in the GetContentQuads CDP command.
//
// JavaScript object id of the node wrapper.
func (t *GetContentQuads) SetObjectID(v runtime.RemoteObjectID) *GetContentQuads {
	t.ObjectID = &v
	return t
}

// GetContentQuadsResponse contains the browser's response
// to calling the GetContentQuads CDP command with Do().
type GetContentQuadsResponse struct {
	// Quads that describe node layout relative to viewport.
	Quads []Quad `json:"quads"`
}

// Do sends the GetContentQuads CDP command to a browser,
// and returns the browser's response.
func (t *GetContentQuads) Do(ctx context.Context) (*GetContentQuadsResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "GetContentQuads", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetContentQuadsResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetDocument contains the parameters, and acts as
// a Go receiver, for the CDP command `getDocument`.
//
// Returns the root DOM node (and optionally the subtree) to the caller.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getDocument
type GetDocument struct {
	// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the
	// entire subtree or provide an integer larger than 0.
	Depth int64 `json:"depth,omitempty"`
	// Whether or not iframes and shadow roots should be traversed when returning the subtree
	// (default is false).
	Pierce bool `json:"pierce,omitempty"`
}

// NewGetDocument constructs a new GetDocument struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getDocument
func NewGetDocument() *GetDocument {
	return &GetDocument{}
}

// SetDepth adds or modifies the value of the optional
// parameter `depth` in the GetDocument CDP command.
//
// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the
// entire subtree or provide an integer larger than 0.
func (t *GetDocument) SetDepth(v int64) *GetDocument {
	t.Depth = v
	return t
}

// SetPierce adds or modifies the value of the optional
// parameter `pierce` in the GetDocument CDP command.
//
// Whether or not iframes and shadow roots should be traversed when returning the subtree
// (default is false).
func (t *GetDocument) SetPierce(v bool) *GetDocument {
	t.Pierce = v
	return t
}

// GetDocumentResponse contains the browser's response
// to calling the GetDocument CDP command with Do().
type GetDocumentResponse struct {
	// Resulting node.
	Root Node `json:"root"`
}

// Do sends the GetDocument CDP command to a browser,
// and returns the browser's response.
func (t *GetDocument) Do(ctx context.Context) (*GetDocumentResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "GetDocument", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetDocumentResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetFlattenedDocument contains the parameters, and acts as
// a Go receiver, for the CDP command `getFlattenedDocument`.
//
// Returns the root DOM node (and optionally the subtree) to the caller.
// Deprecated, as it is not designed to work well with the rest of the DOM agent.
// Use DOMSnapshot.captureSnapshot instead.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getFlattenedDocument
//
// This CDP method is deprecated.
type GetFlattenedDocument struct {
	// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the
	// entire subtree or provide an integer larger than 0.
	Depth int64 `json:"depth,omitempty"`
	// Whether or not iframes and shadow roots should be traversed when returning the subtree
	// (default is false).
	Pierce bool `json:"pierce,omitempty"`
}

// NewGetFlattenedDocument constructs a new GetFlattenedDocument struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getFlattenedDocument
//
// This CDP method is deprecated.
func NewGetFlattenedDocument() *GetFlattenedDocument {
	return &GetFlattenedDocument{}
}

// SetDepth adds or modifies the value of the optional
// parameter `depth` in the GetFlattenedDocument CDP command.
//
// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the
// entire subtree or provide an integer larger than 0.
func (t *GetFlattenedDocument) SetDepth(v int64) *GetFlattenedDocument {
	t.Depth = v
	return t
}

// SetPierce adds or modifies the value of the optional
// parameter `pierce` in the GetFlattenedDocument CDP command.
//
// Whether or not iframes and shadow roots should be traversed when returning the subtree
// (default is false).
func (t *GetFlattenedDocument) SetPierce(v bool) *GetFlattenedDocument {
	t.Pierce = v
	return t
}

// GetFlattenedDocumentResponse contains the browser's response
// to calling the GetFlattenedDocument CDP command with Do().
type GetFlattenedDocumentResponse struct {
	// Resulting node.
	Nodes []Node `json:"nodes"`
}

// Do sends the GetFlattenedDocument CDP command to a browser,
// and returns the browser's response.
func (t *GetFlattenedDocument) Do(ctx context.Context) (*GetFlattenedDocumentResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "GetFlattenedDocument", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetFlattenedDocumentResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetNodesForSubtreeByStyle contains the parameters, and acts as
// a Go receiver, for the CDP command `getNodesForSubtreeByStyle`.
//
// Finds nodes with a given computed style in a subtree.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getNodesForSubtreeByStyle
//
// This CDP method is experimental.
type GetNodesForSubtreeByStyle struct {
	// Node ID pointing to the root of a subtree.
	NodeID NodeID `json:"nodeId"`
	// The style to filter nodes by (includes nodes if any of properties matches).
	ComputedStyles []CSSComputedStyleProperty `json:"computedStyles"`
	// Whether or not iframes and shadow roots in the same target should be traversed when returning the
	// results (default is false).
	Pierce bool `json:"pierce,omitempty"`
}

// NewGetNodesForSubtreeByStyle constructs a new GetNodesForSubtreeByStyle struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getNodesForSubtreeByStyle
//
// This CDP method is experimental.
func NewGetNodesForSubtreeByStyle(nodeId NodeID, computedStyles []CSSComputedStyleProperty) *GetNodesForSubtreeByStyle {
	return &GetNodesForSubtreeByStyle{
		NodeID:         nodeId,
		ComputedStyles: computedStyles,
	}
}

// SetPierce adds or modifies the value of the optional
// parameter `pierce` in the GetNodesForSubtreeByStyle CDP command.
//
// Whether or not iframes and shadow roots in the same target should be traversed when returning the
// results (default is false).
func (t *GetNodesForSubtreeByStyle) SetPierce(v bool) *GetNodesForSubtreeByStyle {
	t.Pierce = v
	return t
}

// GetNodesForSubtreeByStyleResponse contains the browser's response
// to calling the GetNodesForSubtreeByStyle CDP command with Do().
type GetNodesForSubtreeByStyleResponse struct {
	// Resulting nodes.
	NodeIds []NodeID `json:"nodeIds"`
}

// Do sends the GetNodesForSubtreeByStyle CDP command to a browser,
// and returns the browser's response.
func (t *GetNodesForSubtreeByStyle) Do(ctx context.Context) (*GetNodesForSubtreeByStyleResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "GetNodesForSubtreeByStyle", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetNodesForSubtreeByStyleResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetNodeForLocation contains the parameters, and acts as
// a Go receiver, for the CDP command `getNodeForLocation`.
//
// Returns node id at given location. Depending on whether DOM domain is enabled, nodeId is
// either returned or not.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getNodeForLocation
type GetNodeForLocation struct {
	// X coordinate.
	X int64 `json:"x"`
	// Y coordinate.
	Y int64 `json:"y"`
	// False to skip to the nearest non-UA shadow root ancestor (default: false).
	IncludeUserAgentShadowDOM bool `json:"includeUserAgentShadowDOM,omitempty"`
	// Whether to ignore pointer-events: none on elements and hit test them.
	IgnorePointerEventsNone bool `json:"ignorePointerEventsNone,omitempty"`
}

// NewGetNodeForLocation constructs a new GetNodeForLocation struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getNodeForLocation
func NewGetNodeForLocation(x int64, y int64) *GetNodeForLocation {
	return &GetNodeForLocation{
		X: x,
		Y: y,
	}
}

// SetIncludeUserAgentShadowDOM adds or modifies the value of the optional
// parameter `includeUserAgentShadowDOM` in the GetNodeForLocation CDP command.
//
// False to skip to the nearest non-UA shadow root ancestor (default: false).
func (t *GetNodeForLocation) SetIncludeUserAgentShadowDOM(v bool) *GetNodeForLocation {
	t.IncludeUserAgentShadowDOM = v
	return t
}

// SetIgnorePointerEventsNone adds or modifies the value of the optional
// parameter `ignorePointerEventsNone` in the GetNodeForLocation CDP command.
//
// Whether to ignore pointer-events: none on elements and hit test them.
func (t *GetNodeForLocation) SetIgnorePointerEventsNone(v bool) *GetNodeForLocation {
	t.IgnorePointerEventsNone = v
	return t
}

// GetNodeForLocationResponse contains the browser's response
// to calling the GetNodeForLocation CDP command with Do().
type GetNodeForLocationResponse struct {
	// Resulting node.
	BackendNodeID BackendNodeID `json:"backendNodeId"`
	// Frame this node belongs to.
	FrameID cdp.FrameID `json:"frameId"`
	// Id of the node at given coordinates, only when enabled and requested document.
	NodeID *NodeID `json:"nodeId,omitempty"`
}

// Do sends the GetNodeForLocation CDP command to a browser,
// and returns the browser's response.
func (t *GetNodeForLocation) Do(ctx context.Context) (*GetNodeForLocationResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "GetNodeForLocation", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetNodeForLocationResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetOuterHTML contains the parameters, and acts as
// a Go receiver, for the CDP command `getOuterHTML`.
//
// Returns node's HTML markup.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getOuterHTML
type GetOuterHTML struct {
	// Identifier of the node.
	NodeID *NodeID `json:"nodeId,omitempty"`
	// Identifier of the backend node.
	BackendNodeID *BackendNodeID `json:"backendNodeId,omitempty"`
	// JavaScript object id of the node wrapper.
	ObjectID *runtime.RemoteObjectID `json:"objectId,omitempty"`
}

// NewGetOuterHTML constructs a new GetOuterHTML struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getOuterHTML
func NewGetOuterHTML() *GetOuterHTML {
	return &GetOuterHTML{}
}

// SetNodeID adds or modifies the value of the optional
// parameter `nodeId` in the GetOuterHTML CDP command.
//
// Identifier of the node.
func (t *GetOuterHTML) SetNodeID(v NodeID) *GetOuterHTML {
	t.NodeID = &v
	return t
}

// SetBackendNodeID adds or modifies the value of the optional
// parameter `backendNodeId` in the GetOuterHTML CDP command.
//
// Identifier of the backend node.
func (t *GetOuterHTML) SetBackendNodeID(v BackendNodeID) *GetOuterHTML {
	t.BackendNodeID = &v
	return t
}

// SetObjectID adds or modifies the value of the optional
// parameter `objectId` in the GetOuterHTML CDP command.
//
// JavaScript object id of the node wrapper.
func (t *GetOuterHTML) SetObjectID(v runtime.RemoteObjectID) *GetOuterHTML {
	t.ObjectID = &v
	return t
}

// GetOuterHTMLResponse contains the browser's response
// to calling the GetOuterHTML CDP command with Do().
type GetOuterHTMLResponse struct {
	// Outer HTML markup.
	OuterHTML string `json:"outerHTML"`
}

// Do sends the GetOuterHTML CDP command to a browser,
// and returns the browser's response.
func (t *GetOuterHTML) Do(ctx context.Context) (*GetOuterHTMLResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "GetOuterHTML", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetOuterHTMLResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetRelayoutBoundary contains the parameters, and acts as
// a Go receiver, for the CDP command `getRelayoutBoundary`.
//
// Returns the id of the nearest ancestor that is a relayout boundary.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getRelayoutBoundary
//
// This CDP method is experimental.
type GetRelayoutBoundary struct {
	// Id of the node.
	NodeID NodeID `json:"nodeId"`
}

// NewGetRelayoutBoundary constructs a new GetRelayoutBoundary struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getRelayoutBoundary
//
// This CDP method is experimental.
func NewGetRelayoutBoundary(nodeId NodeID) *GetRelayoutBoundary {
	return &GetRelayoutBoundary{
		NodeID: nodeId,
	}
}

// GetRelayoutBoundaryResponse contains the browser's response
// to calling the GetRelayoutBoundary CDP command with Do().
type GetRelayoutBoundaryResponse struct {
	// Relayout boundary node id for the given node.
	NodeID NodeID `json:"nodeId"`
}

// Do sends the GetRelayoutBoundary CDP command to a browser,
// and returns the browser's response.
func (t *GetRelayoutBoundary) Do(ctx context.Context) (*GetRelayoutBoundaryResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "GetRelayoutBoundary", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetRelayoutBoundaryResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetSearchResults contains the parameters, and acts as
// a Go receiver, for the CDP command `getSearchResults`.
//
// Returns search results from given `fromIndex` to given `toIndex` from the search with the given
// identifier.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getSearchResults
//
// This CDP method is experimental.
type GetSearchResults struct {
	// Unique search session identifier.
	SearchID string `json:"searchId"`
	// Start index of the search result to be returned.
	FromIndex int64 `json:"fromIndex"`
	// End index of the search result to be returned.
	ToIndex int64 `json:"toIndex"`
}

// NewGetSearchResults constructs a new GetSearchResults struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getSearchResults
//
// This CDP method is experimental.
func NewGetSearchResults(searchId string, fromIndex int64, toIndex int64) *GetSearchResults {
	return &GetSearchResults{
		SearchID:  searchId,
		FromIndex: fromIndex,
		ToIndex:   toIndex,
	}
}

// GetSearchResultsResponse contains the browser's response
// to calling the GetSearchResults CDP command with Do().
type GetSearchResultsResponse struct {
	// Ids of the search result nodes.
	NodeIds []NodeID `json:"nodeIds"`
}

// Do sends the GetSearchResults CDP command to a browser,
// and returns the browser's response.
func (t *GetSearchResults) Do(ctx context.Context) (*GetSearchResultsResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "GetSearchResults", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetSearchResultsResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// MarkUndoableState contains the parameters, and acts as
// a Go receiver, for the CDP command `markUndoableState`.
//
// Marks last undoable state.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-markUndoableState
//
// This CDP method is experimental.
type MarkUndoableState struct{}

// NewMarkUndoableState constructs a new MarkUndoableState struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-markUndoableState
//
// This CDP method is experimental.
func NewMarkUndoableState() *MarkUndoableState {
	return &MarkUndoableState{}
}

// Do sends the MarkUndoableState CDP command to a browser,
// and returns the browser's response.
func (t *MarkUndoableState) Do(ctx context.Context) error {
	response, err := cdp.Send(ctx, "MarkUndoableState", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// MoveTo contains the parameters, and acts as
// a Go receiver, for the CDP command `moveTo`.
//
// Moves node into the new container, places it before the given anchor.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-moveTo
type MoveTo struct {
	// Id of the node to move.
	NodeID NodeID `json:"nodeId"`
	// Id of the element to drop the moved node into.
	TargetNodeID NodeID `json:"targetNodeId"`
	// Drop node before this one (if absent, the moved node becomes the last child of
	// `targetNodeId`).
	InsertBeforeNodeID *NodeID `json:"insertBeforeNodeId,omitempty"`
}

// NewMoveTo constructs a new MoveTo struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-moveTo
func NewMoveTo(nodeId NodeID, targetNodeId NodeID) *MoveTo {
	return &MoveTo{
		NodeID:       nodeId,
		TargetNodeID: targetNodeId,
	}
}

// SetInsertBeforeNodeID adds or modifies the value of the optional
// parameter `insertBeforeNodeId` in the MoveTo CDP command.
//
// Drop node before this one (if absent, the moved node becomes the last child of
// `targetNodeId`).
func (t *MoveTo) SetInsertBeforeNodeID(v NodeID) *MoveTo {
	t.InsertBeforeNodeID = &v
	return t
}

// MoveToResponse contains the browser's response
// to calling the MoveTo CDP command with Do().
type MoveToResponse struct {
	// New id of the moved node.
	NodeID NodeID `json:"nodeId"`
}

// Do sends the MoveTo CDP command to a browser,
// and returns the browser's response.
func (t *MoveTo) Do(ctx context.Context) (*MoveToResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "MoveTo", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &MoveToResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// PerformSearch contains the parameters, and acts as
// a Go receiver, for the CDP command `performSearch`.
//
// Searches for a given string in the DOM tree. Use `getSearchResults` to access search results or
// `cancelSearch` to end this search session.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-performSearch
//
// This CDP method is experimental.
type PerformSearch struct {
	// Plain text or query selector or XPath search query.
	Query string `json:"query"`
	// True to search in user agent shadow DOM.
	IncludeUserAgentShadowDOM bool `json:"includeUserAgentShadowDOM,omitempty"`
}

// NewPerformSearch constructs a new PerformSearch struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-performSearch
//
// This CDP method is experimental.
func NewPerformSearch(query string) *PerformSearch {
	return &PerformSearch{
		Query: query,
	}
}

// SetIncludeUserAgentShadowDOM adds or modifies the value of the optional
// parameter `includeUserAgentShadowDOM` in the PerformSearch CDP command.
//
// True to search in user agent shadow DOM.
func (t *PerformSearch) SetIncludeUserAgentShadowDOM(v bool) *PerformSearch {
	t.IncludeUserAgentShadowDOM = v
	return t
}

// PerformSearchResponse contains the browser's response
// to calling the PerformSearch CDP command with Do().
type PerformSearchResponse struct {
	// Unique search session identifier.
	SearchID string `json:"searchId"`
	// Number of search results.
	ResultCount int64 `json:"resultCount"`
}

// Do sends the PerformSearch CDP command to a browser,
// and returns the browser's response.
func (t *PerformSearch) Do(ctx context.Context) (*PerformSearchResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "PerformSearch", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &PerformSearchResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// PushNodeByPathToFrontend contains the parameters, and acts as
// a Go receiver, for the CDP command `pushNodeByPathToFrontend`.
//
// Requests that the node is sent to the caller given its path. // FIXME, use XPath
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-pushNodeByPathToFrontend
//
// This CDP method is experimental.
type PushNodeByPathToFrontend struct {
	// Path to node in the proprietary format.
	Path string `json:"path"`
}

// NewPushNodeByPathToFrontend constructs a new PushNodeByPathToFrontend struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-pushNodeByPathToFrontend
//
// This CDP method is experimental.
func NewPushNodeByPathToFrontend(path string) *PushNodeByPathToFrontend {
	return &PushNodeByPathToFrontend{
		Path: path,
	}
}

// PushNodeByPathToFrontendResponse contains the browser's response
// to calling the PushNodeByPathToFrontend CDP command with Do().
type PushNodeByPathToFrontendResponse struct {
	// Id of the node for given path.
	NodeID NodeID `json:"nodeId"`
}

// Do sends the PushNodeByPathToFrontend CDP command to a browser,
// and returns the browser's response.
func (t *PushNodeByPathToFrontend) Do(ctx context.Context) (*PushNodeByPathToFrontendResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "PushNodeByPathToFrontend", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &PushNodeByPathToFrontendResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// PushNodesByBackendIdsToFrontend contains the parameters, and acts as
// a Go receiver, for the CDP command `pushNodesByBackendIdsToFrontend`.
//
// Requests that a batch of nodes is sent to the caller given their backend node ids.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-pushNodesByBackendIdsToFrontend
//
// This CDP method is experimental.
type PushNodesByBackendIdsToFrontend struct {
	// The array of backend node ids.
	BackendNodeIds []BackendNodeID `json:"backendNodeIds"`
}

// NewPushNodesByBackendIdsToFrontend constructs a new PushNodesByBackendIdsToFrontend struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-pushNodesByBackendIdsToFrontend
//
// This CDP method is experimental.
func NewPushNodesByBackendIdsToFrontend(backendNodeIds []BackendNodeID) *PushNodesByBackendIdsToFrontend {
	return &PushNodesByBackendIdsToFrontend{
		BackendNodeIds: backendNodeIds,
	}
}

// PushNodesByBackendIdsToFrontendResponse contains the browser's response
// to calling the PushNodesByBackendIdsToFrontend CDP command with Do().
type PushNodesByBackendIdsToFrontendResponse struct {
	// The array of ids of pushed nodes that correspond to the backend ids specified in
	// backendNodeIds.
	NodeIds []NodeID `json:"nodeIds"`
}

// Do sends the PushNodesByBackendIdsToFrontend CDP command to a browser,
// and returns the browser's response.
func (t *PushNodesByBackendIdsToFrontend) Do(ctx context.Context) (*PushNodesByBackendIdsToFrontendResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "PushNodesByBackendIdsToFrontend", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &PushNodesByBackendIdsToFrontendResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// QuerySelector contains the parameters, and acts as
// a Go receiver, for the CDP command `querySelector`.
//
// Executes `querySelector` on a given node.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-querySelector
type QuerySelector struct {
	// Id of the node to query upon.
	NodeID NodeID `json:"nodeId"`
	// Selector string.
	Selector string `json:"selector"`
}

// NewQuerySelector constructs a new QuerySelector struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-querySelector
func NewQuerySelector(nodeId NodeID, selector string) *QuerySelector {
	return &QuerySelector{
		NodeID:   nodeId,
		Selector: selector,
	}
}

// QuerySelectorResponse contains the browser's response
// to calling the QuerySelector CDP command with Do().
type QuerySelectorResponse struct {
	// Query selector result.
	NodeID NodeID `json:"nodeId"`
}

// Do sends the QuerySelector CDP command to a browser,
// and returns the browser's response.
func (t *QuerySelector) Do(ctx context.Context) (*QuerySelectorResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "QuerySelector", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &QuerySelectorResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// QuerySelectorAll contains the parameters, and acts as
// a Go receiver, for the CDP command `querySelectorAll`.
//
// Executes `querySelectorAll` on a given node.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-querySelectorAll
type QuerySelectorAll struct {
	// Id of the node to query upon.
	NodeID NodeID `json:"nodeId"`
	// Selector string.
	Selector string `json:"selector"`
}

// NewQuerySelectorAll constructs a new QuerySelectorAll struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-querySelectorAll
func NewQuerySelectorAll(nodeId NodeID, selector string) *QuerySelectorAll {
	return &QuerySelectorAll{
		NodeID:   nodeId,
		Selector: selector,
	}
}

// QuerySelectorAllResponse contains the browser's response
// to calling the QuerySelectorAll CDP command with Do().
type QuerySelectorAllResponse struct {
	// Query selector result.
	NodeIds []NodeID `json:"nodeIds"`
}

// Do sends the QuerySelectorAll CDP command to a browser,
// and returns the browser's response.
func (t *QuerySelectorAll) Do(ctx context.Context) (*QuerySelectorAllResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "QuerySelectorAll", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &QuerySelectorAllResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Redo contains the parameters, and acts as
// a Go receiver, for the CDP command `redo`.
//
// Re-does the last undone action.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-redo
//
// This CDP method is experimental.
type Redo struct{}

// NewRedo constructs a new Redo struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-redo
//
// This CDP method is experimental.
func NewRedo() *Redo {
	return &Redo{}
}

// Do sends the Redo CDP command to a browser,
// and returns the browser's response.
func (t *Redo) Do(ctx context.Context) error {
	response, err := cdp.Send(ctx, "Redo", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// RemoveAttribute contains the parameters, and acts as
// a Go receiver, for the CDP command `removeAttribute`.
//
// Removes attribute with given name from an element with given id.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-removeAttribute
type RemoveAttribute struct {
	// Id of the element to remove attribute from.
	NodeID NodeID `json:"nodeId"`
	// Name of the attribute to remove.
	Name string `json:"name"`
}

// NewRemoveAttribute constructs a new RemoveAttribute struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-removeAttribute
func NewRemoveAttribute(nodeId NodeID, name string) *RemoveAttribute {
	return &RemoveAttribute{
		NodeID: nodeId,
		Name:   name,
	}
}

// Do sends the RemoveAttribute CDP command to a browser,
// and returns the browser's response.
func (t *RemoveAttribute) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "RemoveAttribute", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// RemoveNode contains the parameters, and acts as
// a Go receiver, for the CDP command `removeNode`.
//
// Removes node with given id.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-removeNode
type RemoveNode struct {
	// Id of the node to remove.
	NodeID NodeID `json:"nodeId"`
}

// NewRemoveNode constructs a new RemoveNode struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-removeNode
func NewRemoveNode(nodeId NodeID) *RemoveNode {
	return &RemoveNode{
		NodeID: nodeId,
	}
}

// Do sends the RemoveNode CDP command to a browser,
// and returns the browser's response.
func (t *RemoveNode) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "RemoveNode", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// RequestChildNodes contains the parameters, and acts as
// a Go receiver, for the CDP command `requestChildNodes`.
//
// Requests that children of the node with given id are returned to the caller in form of
// `setChildNodes` events where not only immediate children are retrieved, but all children down to
// the specified depth.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-requestChildNodes
type RequestChildNodes struct {
	// Id of the node to get children for.
	NodeID NodeID `json:"nodeId"`
	// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the
	// entire subtree or provide an integer larger than 0.
	Depth int64 `json:"depth,omitempty"`
	// Whether or not iframes and shadow roots should be traversed when returning the sub-tree
	// (default is false).
	Pierce bool `json:"pierce,omitempty"`
}

// NewRequestChildNodes constructs a new RequestChildNodes struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-requestChildNodes
func NewRequestChildNodes(nodeId NodeID) *RequestChildNodes {
	return &RequestChildNodes{
		NodeID: nodeId,
	}
}

// SetDepth adds or modifies the value of the optional
// parameter `depth` in the RequestChildNodes CDP command.
//
// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the
// entire subtree or provide an integer larger than 0.
func (t *RequestChildNodes) SetDepth(v int64) *RequestChildNodes {
	t.Depth = v
	return t
}

// SetPierce adds or modifies the value of the optional
// parameter `pierce` in the RequestChildNodes CDP command.
//
// Whether or not iframes and shadow roots should be traversed when returning the sub-tree
// (default is false).
func (t *RequestChildNodes) SetPierce(v bool) *RequestChildNodes {
	t.Pierce = v
	return t
}

// Do sends the RequestChildNodes CDP command to a browser,
// and returns the browser's response.
func (t *RequestChildNodes) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "RequestChildNodes", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// RequestNode contains the parameters, and acts as
// a Go receiver, for the CDP command `requestNode`.
//
// Requests that the node is sent to the caller given the JavaScript node object reference. All
// nodes that form the path from the node to the root are also sent to the client as a series of
// `setChildNodes` notifications.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-requestNode
type RequestNode struct {
	// JavaScript object id to convert into node.
	ObjectID runtime.RemoteObjectID `json:"objectId"`
}

// NewRequestNode constructs a new RequestNode struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-requestNode
func NewRequestNode(objectId runtime.RemoteObjectID) *RequestNode {
	return &RequestNode{
		ObjectID: objectId,
	}
}

// RequestNodeResponse contains the browser's response
// to calling the RequestNode CDP command with Do().
type RequestNodeResponse struct {
	// Node id for given object.
	NodeID NodeID `json:"nodeId"`
}

// Do sends the RequestNode CDP command to a browser,
// and returns the browser's response.
func (t *RequestNode) Do(ctx context.Context) (*RequestNodeResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "RequestNode", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &RequestNodeResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ResolveNode contains the parameters, and acts as
// a Go receiver, for the CDP command `resolveNode`.
//
// Resolves the JavaScript node object for a given NodeId or BackendNodeId.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-resolveNode
type ResolveNode struct {
	// Id of the node to resolve.
	NodeID *NodeID `json:"nodeId,omitempty"`
	// Backend identifier of the node to resolve.
	BackendNodeID *BackendNodeID `json:"backendNodeId,omitempty"`
	// Symbolic group name that can be used to release multiple objects.
	ObjectGroup string `json:"objectGroup,omitempty"`
	// Execution context in which to resolve the node.
	ExecutionContextID *runtime.ExecutionContextID `json:"executionContextId,omitempty"`
}

// NewResolveNode constructs a new ResolveNode struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-resolveNode
func NewResolveNode() *ResolveNode {
	return &ResolveNode{}
}

// SetNodeID adds or modifies the value of the optional
// parameter `nodeId` in the ResolveNode CDP command.
//
// Id of the node to resolve.
func (t *ResolveNode) SetNodeID(v NodeID) *ResolveNode {
	t.NodeID = &v
	return t
}

// SetBackendNodeID adds or modifies the value of the optional
// parameter `backendNodeId` in the ResolveNode CDP command.
//
// Backend identifier of the node to resolve.
func (t *ResolveNode) SetBackendNodeID(v BackendNodeID) *ResolveNode {
	t.BackendNodeID = &v
	return t
}

// SetObjectGroup adds or modifies the value of the optional
// parameter `objectGroup` in the ResolveNode CDP command.
//
// Symbolic group name that can be used to release multiple objects.
func (t *ResolveNode) SetObjectGroup(v string) *ResolveNode {
	t.ObjectGroup = v
	return t
}

// SetExecutionContextID adds or modifies the value of the optional
// parameter `executionContextId` in the ResolveNode CDP command.
//
// Execution context in which to resolve the node.
func (t *ResolveNode) SetExecutionContextID(v runtime.ExecutionContextID) *ResolveNode {
	t.ExecutionContextID = &v
	return t
}

// ResolveNodeResponse contains the browser's response
// to calling the ResolveNode CDP command with Do().
type ResolveNodeResponse struct {
	// JavaScript object wrapper for given node.
	Object runtime.RemoteObject `json:"object"`
}

// Do sends the ResolveNode CDP command to a browser,
// and returns the browser's response.
func (t *ResolveNode) Do(ctx context.Context) (*ResolveNodeResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "ResolveNode", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &ResolveNodeResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetAttributeValue contains the parameters, and acts as
// a Go receiver, for the CDP command `setAttributeValue`.
//
// Sets attribute for an element with given id.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setAttributeValue
type SetAttributeValue struct {
	// Id of the element to set attribute for.
	NodeID NodeID `json:"nodeId"`
	// Attribute name.
	Name string `json:"name"`
	// Attribute value.
	Value string `json:"value"`
}

// NewSetAttributeValue constructs a new SetAttributeValue struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setAttributeValue
func NewSetAttributeValue(nodeId NodeID, name string, value string) *SetAttributeValue {
	return &SetAttributeValue{
		NodeID: nodeId,
		Name:   name,
		Value:  value,
	}
}

// Do sends the SetAttributeValue CDP command to a browser,
// and returns the browser's response.
func (t *SetAttributeValue) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "SetAttributeValue", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetAttributesAsText contains the parameters, and acts as
// a Go receiver, for the CDP command `setAttributesAsText`.
//
// Sets attributes on element with given id. This method is useful when user edits some existing
// attribute value and types in several attribute name/value pairs.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setAttributesAsText
type SetAttributesAsText struct {
	// Id of the element to set attributes for.
	NodeID NodeID `json:"nodeId"`
	// Text with a number of attributes. Will parse this text using HTML parser.
	Text string `json:"text"`
	// Attribute name to replace with new attributes derived from text in case text parsed
	// successfully.
	Name string `json:"name,omitempty"`
}

// NewSetAttributesAsText constructs a new SetAttributesAsText struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setAttributesAsText
func NewSetAttributesAsText(nodeId NodeID, text string) *SetAttributesAsText {
	return &SetAttributesAsText{
		NodeID: nodeId,
		Text:   text,
	}
}

// SetName adds or modifies the value of the optional
// parameter `name` in the SetAttributesAsText CDP command.
//
// Attribute name to replace with new attributes derived from text in case text parsed
// successfully.
func (t *SetAttributesAsText) SetName(v string) *SetAttributesAsText {
	t.Name = v
	return t
}

// Do sends the SetAttributesAsText CDP command to a browser,
// and returns the browser's response.
func (t *SetAttributesAsText) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "SetAttributesAsText", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetFileInputFiles contains the parameters, and acts as
// a Go receiver, for the CDP command `setFileInputFiles`.
//
// Sets files for the given file input element.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setFileInputFiles
type SetFileInputFiles struct {
	// Array of file paths to set.
	Files []string `json:"files"`
	// Identifier of the node.
	NodeID *NodeID `json:"nodeId,omitempty"`
	// Identifier of the backend node.
	BackendNodeID *BackendNodeID `json:"backendNodeId,omitempty"`
	// JavaScript object id of the node wrapper.
	ObjectID *runtime.RemoteObjectID `json:"objectId,omitempty"`
}

// NewSetFileInputFiles constructs a new SetFileInputFiles struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setFileInputFiles
func NewSetFileInputFiles(files []string) *SetFileInputFiles {
	return &SetFileInputFiles{
		Files: files,
	}
}

// SetNodeID adds or modifies the value of the optional
// parameter `nodeId` in the SetFileInputFiles CDP command.
//
// Identifier of the node.
func (t *SetFileInputFiles) SetNodeID(v NodeID) *SetFileInputFiles {
	t.NodeID = &v
	return t
}

// SetBackendNodeID adds or modifies the value of the optional
// parameter `backendNodeId` in the SetFileInputFiles CDP command.
//
// Identifier of the backend node.
func (t *SetFileInputFiles) SetBackendNodeID(v BackendNodeID) *SetFileInputFiles {
	t.BackendNodeID = &v
	return t
}

// SetObjectID adds or modifies the value of the optional
// parameter `objectId` in the SetFileInputFiles CDP command.
//
// JavaScript object id of the node wrapper.
func (t *SetFileInputFiles) SetObjectID(v runtime.RemoteObjectID) *SetFileInputFiles {
	t.ObjectID = &v
	return t
}

// Do sends the SetFileInputFiles CDP command to a browser,
// and returns the browser's response.
func (t *SetFileInputFiles) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "SetFileInputFiles", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetNodeStackTracesEnabled contains the parameters, and acts as
// a Go receiver, for the CDP command `setNodeStackTracesEnabled`.
//
// Sets if stack traces should be captured for Nodes. See `Node.getNodeStackTraces`. Default is disabled.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setNodeStackTracesEnabled
//
// This CDP method is experimental.
type SetNodeStackTracesEnabled struct {
	// Enable or disable.
	Enable bool `json:"enable"`
}

// NewSetNodeStackTracesEnabled constructs a new SetNodeStackTracesEnabled struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setNodeStackTracesEnabled
//
// This CDP method is experimental.
func NewSetNodeStackTracesEnabled(enable bool) *SetNodeStackTracesEnabled {
	return &SetNodeStackTracesEnabled{
		Enable: enable,
	}
}

// Do sends the SetNodeStackTracesEnabled CDP command to a browser,
// and returns the browser's response.
func (t *SetNodeStackTracesEnabled) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "SetNodeStackTracesEnabled", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// GetNodeStackTraces contains the parameters, and acts as
// a Go receiver, for the CDP command `getNodeStackTraces`.
//
// Gets stack traces associated with a Node. As of now, only provides stack trace for Node creation.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getNodeStackTraces
//
// This CDP method is experimental.
type GetNodeStackTraces struct {
	// Id of the node to get stack traces for.
	NodeID NodeID `json:"nodeId"`
}

// NewGetNodeStackTraces constructs a new GetNodeStackTraces struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getNodeStackTraces
//
// This CDP method is experimental.
func NewGetNodeStackTraces(nodeId NodeID) *GetNodeStackTraces {
	return &GetNodeStackTraces{
		NodeID: nodeId,
	}
}

// GetNodeStackTracesResponse contains the browser's response
// to calling the GetNodeStackTraces CDP command with Do().
type GetNodeStackTracesResponse struct {
	// Creation stack trace, if available.
	Creation *runtime.StackTrace `json:"creation,omitempty"`
}

// Do sends the GetNodeStackTraces CDP command to a browser,
// and returns the browser's response.
func (t *GetNodeStackTraces) Do(ctx context.Context) (*GetNodeStackTracesResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "GetNodeStackTraces", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetNodeStackTracesResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetFileInfo contains the parameters, and acts as
// a Go receiver, for the CDP command `getFileInfo`.
//
// Returns file information for the given
// File wrapper.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getFileInfo
//
// This CDP method is experimental.
type GetFileInfo struct {
	// JavaScript object id of the node wrapper.
	ObjectID runtime.RemoteObjectID `json:"objectId"`
}

// NewGetFileInfo constructs a new GetFileInfo struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getFileInfo
//
// This CDP method is experimental.
func NewGetFileInfo(objectId runtime.RemoteObjectID) *GetFileInfo {
	return &GetFileInfo{
		ObjectID: objectId,
	}
}

// GetFileInfoResponse contains the browser's response
// to calling the GetFileInfo CDP command with Do().
type GetFileInfoResponse struct {
	Path string `json:"path"`
}

// Do sends the GetFileInfo CDP command to a browser,
// and returns the browser's response.
func (t *GetFileInfo) Do(ctx context.Context) (*GetFileInfoResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "GetFileInfo", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetFileInfoResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetInspectedNode contains the parameters, and acts as
// a Go receiver, for the CDP command `setInspectedNode`.
//
// Enables console to refer to the node with given id via $x (see Command Line API for more details
// $x functions).
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setInspectedNode
//
// This CDP method is experimental.
type SetInspectedNode struct {
	// DOM node id to be accessible by means of $x command line API.
	NodeID NodeID `json:"nodeId"`
}

// NewSetInspectedNode constructs a new SetInspectedNode struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setInspectedNode
//
// This CDP method is experimental.
func NewSetInspectedNode(nodeId NodeID) *SetInspectedNode {
	return &SetInspectedNode{
		NodeID: nodeId,
	}
}

// Do sends the SetInspectedNode CDP command to a browser,
// and returns the browser's response.
func (t *SetInspectedNode) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "SetInspectedNode", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetNodeName contains the parameters, and acts as
// a Go receiver, for the CDP command `setNodeName`.
//
// Sets node name for a node with given id.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setNodeName
type SetNodeName struct {
	// Id of the node to set name for.
	NodeID NodeID `json:"nodeId"`
	// New node's name.
	Name string `json:"name"`
}

// NewSetNodeName constructs a new SetNodeName struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setNodeName
func NewSetNodeName(nodeId NodeID, name string) *SetNodeName {
	return &SetNodeName{
		NodeID: nodeId,
		Name:   name,
	}
}

// SetNodeNameResponse contains the browser's response
// to calling the SetNodeName CDP command with Do().
type SetNodeNameResponse struct {
	// New node's id.
	NodeID NodeID `json:"nodeId"`
}

// Do sends the SetNodeName CDP command to a browser,
// and returns the browser's response.
func (t *SetNodeName) Do(ctx context.Context) (*SetNodeNameResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "SetNodeName", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &SetNodeNameResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetNodeValue contains the parameters, and acts as
// a Go receiver, for the CDP command `setNodeValue`.
//
// Sets node value for a node with given id.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setNodeValue
type SetNodeValue struct {
	// Id of the node to set value for.
	NodeID NodeID `json:"nodeId"`
	// New node's value.
	Value string `json:"value"`
}

// NewSetNodeValue constructs a new SetNodeValue struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setNodeValue
func NewSetNodeValue(nodeId NodeID, value string) *SetNodeValue {
	return &SetNodeValue{
		NodeID: nodeId,
		Value:  value,
	}
}

// Do sends the SetNodeValue CDP command to a browser,
// and returns the browser's response.
func (t *SetNodeValue) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "SetNodeValue", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetOuterHTML contains the parameters, and acts as
// a Go receiver, for the CDP command `setOuterHTML`.
//
// Sets node HTML markup, returns new node id.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setOuterHTML
type SetOuterHTML struct {
	// Id of the node to set markup for.
	NodeID NodeID `json:"nodeId"`
	// Outer HTML markup to set.
	OuterHTML string `json:"outerHTML"`
}

// NewSetOuterHTML constructs a new SetOuterHTML struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setOuterHTML
func NewSetOuterHTML(nodeId NodeID, outerHTML string) *SetOuterHTML {
	return &SetOuterHTML{
		NodeID:    nodeId,
		OuterHTML: outerHTML,
	}
}

// Do sends the SetOuterHTML CDP command to a browser,
// and returns the browser's response.
func (t *SetOuterHTML) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "SetOuterHTML", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// Undo contains the parameters, and acts as
// a Go receiver, for the CDP command `undo`.
//
// Undoes the last performed action.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-undo
//
// This CDP method is experimental.
type Undo struct{}

// NewUndo constructs a new Undo struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-undo
//
// This CDP method is experimental.
func NewUndo() *Undo {
	return &Undo{}
}

// Do sends the Undo CDP command to a browser,
// and returns the browser's response.
func (t *Undo) Do(ctx context.Context) error {
	response, err := cdp.Send(ctx, "Undo", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// GetFrameOwner contains the parameters, and acts as
// a Go receiver, for the CDP command `getFrameOwner`.
//
// Returns iframe node that owns iframe with the given domain.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getFrameOwner
//
// This CDP method is experimental.
type GetFrameOwner struct {
	FrameID cdp.FrameID `json:"frameId"`
}

// NewGetFrameOwner constructs a new GetFrameOwner struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getFrameOwner
//
// This CDP method is experimental.
func NewGetFrameOwner(frameId cdp.FrameID) *GetFrameOwner {
	return &GetFrameOwner{
		FrameID: frameId,
	}
}

// GetFrameOwnerResponse contains the browser's response
// to calling the GetFrameOwner CDP command with Do().
type GetFrameOwnerResponse struct {
	// Resulting node.
	BackendNodeID BackendNodeID `json:"backendNodeId"`
	// Id of the node at given coordinates, only when enabled and requested document.
	NodeID *NodeID `json:"nodeId,omitempty"`
}

// Do sends the GetFrameOwner CDP command to a browser,
// and returns the browser's response.
func (t *GetFrameOwner) Do(ctx context.Context) (*GetFrameOwnerResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "GetFrameOwner", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetFrameOwnerResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}
