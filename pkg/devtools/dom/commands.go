package dom

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
	"github.com/daabr/chrome-vision/pkg/devtools/runtime"
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
	NodeID int64 `json:"nodeId"`
}

// NewCollectClassNamesFromSubtree constructs a new CollectClassNamesFromSubtree struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-collectClassNamesFromSubtree
//
// This CDP method is experimental.
func NewCollectClassNamesFromSubtree(nodeID int64) *CollectClassNamesFromSubtree {
	return &CollectClassNamesFromSubtree{
		NodeID: nodeID,
	}
}

// CollectClassNamesFromSubtreeResult contains the browser's response
// to calling the CollectClassNamesFromSubtree CDP command with Do().
type CollectClassNamesFromSubtreeResult struct {
	// Class name list.
	ClassNames []string `json:"classNames"`
}

// Do sends the CollectClassNamesFromSubtree CDP command to a browser,
// and returns the browser's response.
func (t *CollectClassNamesFromSubtree) Do(ctx context.Context) (*CollectClassNamesFromSubtreeResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.Send(ctx, "DOM.collectClassNamesFromSubtree", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &CollectClassNamesFromSubtreeResult{}
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
	NodeID int64 `json:"nodeId"`
	// Id of the element to drop the copy into.
	TargetNodeID int64 `json:"targetNodeId"`
	// Drop the copy before this node (if absent, the copy becomes the last child of
	// `targetNodeId`).
	InsertBeforeNodeID int64 `json:"insertBeforeNodeId,omitempty"`
}

// NewCopyTo constructs a new CopyTo struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-copyTo
//
// This CDP method is experimental.
func NewCopyTo(nodeID int64, targetNodeID int64) *CopyTo {
	return &CopyTo{
		NodeID:       nodeID,
		TargetNodeID: targetNodeID,
	}
}

// SetInsertBeforeNodeID adds or modifies the value of the optional
// parameter `insertBeforeNodeId` in the CopyTo CDP command.
//
// Drop the copy before this node (if absent, the copy becomes the last child of
// `targetNodeId`).
func (t *CopyTo) SetInsertBeforeNodeID(v int64) *CopyTo {
	t.InsertBeforeNodeID = v
	return t
}

// CopyToResult contains the browser's response
// to calling the CopyTo CDP command with Do().
type CopyToResult struct {
	// Id of the node clone.
	NodeID int64 `json:"nodeId"`
}

// Do sends the CopyTo CDP command to a browser,
// and returns the browser's response.
func (t *CopyTo) Do(ctx context.Context) (*CopyToResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.Send(ctx, "DOM.copyTo", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &CopyToResult{}
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
	NodeID int64 `json:"nodeId,omitempty"`
	// Identifier of the backend node.
	BackendNodeID int64 `json:"backendNodeId,omitempty"`
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
func (t *DescribeNode) SetNodeID(v int64) *DescribeNode {
	t.NodeID = v
	return t
}

// SetBackendNodeID adds or modifies the value of the optional
// parameter `backendNodeId` in the DescribeNode CDP command.
//
// Identifier of the backend node.
func (t *DescribeNode) SetBackendNodeID(v int64) *DescribeNode {
	t.BackendNodeID = v
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

// DescribeNodeResult contains the browser's response
// to calling the DescribeNode CDP command with Do().
type DescribeNodeResult struct {
	// Node description.
	Node Node `json:"node"`
}

// Do sends the DescribeNode CDP command to a browser,
// and returns the browser's response.
func (t *DescribeNode) Do(ctx context.Context) (*DescribeNodeResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.Send(ctx, "DOM.describeNode", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &DescribeNodeResult{}
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
	NodeID int64 `json:"nodeId,omitempty"`
	// Identifier of the backend node.
	BackendNodeID int64 `json:"backendNodeId,omitempty"`
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
func (t *ScrollIntoViewIfNeeded) SetNodeID(v int64) *ScrollIntoViewIfNeeded {
	t.NodeID = v
	return t
}

// SetBackendNodeID adds or modifies the value of the optional
// parameter `backendNodeId` in the ScrollIntoViewIfNeeded CDP command.
//
// Identifier of the backend node.
func (t *ScrollIntoViewIfNeeded) SetBackendNodeID(v int64) *ScrollIntoViewIfNeeded {
	t.BackendNodeID = v
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
	response, err := devtools.Send(ctx, "DOM.scrollIntoViewIfNeeded", b)
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
	response, err := devtools.Send(ctx, "DOM.disable", nil)
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
func NewDiscardSearchResults(searchID string) *DiscardSearchResults {
	return &DiscardSearchResults{
		SearchID: searchID,
	}
}

// Do sends the DiscardSearchResults CDP command to a browser,
// and returns the browser's response.
func (t *DiscardSearchResults) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := devtools.Send(ctx, "DOM.discardSearchResults", b)
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
	response, err := devtools.Send(ctx, "DOM.enable", nil)
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
	NodeID int64 `json:"nodeId,omitempty"`
	// Identifier of the backend node.
	BackendNodeID int64 `json:"backendNodeId,omitempty"`
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
func (t *Focus) SetNodeID(v int64) *Focus {
	t.NodeID = v
	return t
}

// SetBackendNodeID adds or modifies the value of the optional
// parameter `backendNodeId` in the Focus CDP command.
//
// Identifier of the backend node.
func (t *Focus) SetBackendNodeID(v int64) *Focus {
	t.BackendNodeID = v
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
	response, err := devtools.Send(ctx, "DOM.focus", b)
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
	NodeID int64 `json:"nodeId"`
}

// NewGetAttributes constructs a new GetAttributes struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getAttributes
func NewGetAttributes(nodeID int64) *GetAttributes {
	return &GetAttributes{
		NodeID: nodeID,
	}
}

// GetAttributesResult contains the browser's response
// to calling the GetAttributes CDP command with Do().
type GetAttributesResult struct {
	// An interleaved array of node attribute names and values.
	Attributes []string `json:"attributes"`
}

// Do sends the GetAttributes CDP command to a browser,
// and returns the browser's response.
func (t *GetAttributes) Do(ctx context.Context) (*GetAttributesResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.Send(ctx, "DOM.getAttributes", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetAttributesResult{}
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
	NodeID int64 `json:"nodeId,omitempty"`
	// Identifier of the backend node.
	BackendNodeID int64 `json:"backendNodeId,omitempty"`
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
func (t *GetBoxModel) SetNodeID(v int64) *GetBoxModel {
	t.NodeID = v
	return t
}

// SetBackendNodeID adds or modifies the value of the optional
// parameter `backendNodeId` in the GetBoxModel CDP command.
//
// Identifier of the backend node.
func (t *GetBoxModel) SetBackendNodeID(v int64) *GetBoxModel {
	t.BackendNodeID = v
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

// GetBoxModelResult contains the browser's response
// to calling the GetBoxModel CDP command with Do().
type GetBoxModelResult struct {
	// Box model for the node.
	Model BoxModel `json:"model"`
}

// Do sends the GetBoxModel CDP command to a browser,
// and returns the browser's response.
func (t *GetBoxModel) Do(ctx context.Context) (*GetBoxModelResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.Send(ctx, "DOM.getBoxModel", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetBoxModelResult{}
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
	NodeID int64 `json:"nodeId,omitempty"`
	// Identifier of the backend node.
	BackendNodeID int64 `json:"backendNodeId,omitempty"`
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
func (t *GetContentQuads) SetNodeID(v int64) *GetContentQuads {
	t.NodeID = v
	return t
}

// SetBackendNodeID adds or modifies the value of the optional
// parameter `backendNodeId` in the GetContentQuads CDP command.
//
// Identifier of the backend node.
func (t *GetContentQuads) SetBackendNodeID(v int64) *GetContentQuads {
	t.BackendNodeID = v
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

// GetContentQuadsResult contains the browser's response
// to calling the GetContentQuads CDP command with Do().
type GetContentQuadsResult struct {
	// Quads that describe node layout relative to viewport.
	Quads []Quad `json:"quads"`
}

// Do sends the GetContentQuads CDP command to a browser,
// and returns the browser's response.
func (t *GetContentQuads) Do(ctx context.Context) (*GetContentQuadsResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.Send(ctx, "DOM.getContentQuads", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetContentQuadsResult{}
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

// GetDocumentResult contains the browser's response
// to calling the GetDocument CDP command with Do().
type GetDocumentResult struct {
	// Resulting node.
	Root Node `json:"root"`
}

// Do sends the GetDocument CDP command to a browser,
// and returns the browser's response.
func (t *GetDocument) Do(ctx context.Context) (*GetDocumentResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.Send(ctx, "DOM.getDocument", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetDocumentResult{}
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

// GetFlattenedDocumentResult contains the browser's response
// to calling the GetFlattenedDocument CDP command with Do().
type GetFlattenedDocumentResult struct {
	// Resulting node.
	Nodes []Node `json:"nodes"`
}

// Do sends the GetFlattenedDocument CDP command to a browser,
// and returns the browser's response.
func (t *GetFlattenedDocument) Do(ctx context.Context) (*GetFlattenedDocumentResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.Send(ctx, "DOM.getFlattenedDocument", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetFlattenedDocumentResult{}
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
	NodeID int64 `json:"nodeId"`
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
func NewGetNodesForSubtreeByStyle(nodeID int64, computedStyles []CSSComputedStyleProperty) *GetNodesForSubtreeByStyle {
	return &GetNodesForSubtreeByStyle{
		NodeID:         nodeID,
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

// GetNodesForSubtreeByStyleResult contains the browser's response
// to calling the GetNodesForSubtreeByStyle CDP command with Do().
type GetNodesForSubtreeByStyleResult struct {
	// Resulting nodes.
	NodeIds []NodeID `json:"nodeIds"`
}

// Do sends the GetNodesForSubtreeByStyle CDP command to a browser,
// and returns the browser's response.
func (t *GetNodesForSubtreeByStyle) Do(ctx context.Context) (*GetNodesForSubtreeByStyleResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.Send(ctx, "DOM.getNodesForSubtreeByStyle", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetNodesForSubtreeByStyleResult{}
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

// GetNodeForLocationResult contains the browser's response
// to calling the GetNodeForLocation CDP command with Do().
type GetNodeForLocationResult struct {
	// Resulting node.
	BackendNodeID int64 `json:"backendNodeId"`
	// Frame this node belongs to.
	FrameID string `json:"frameId"`
	// Id of the node at given coordinates, only when enabled and requested document.
	NodeID int64 `json:"nodeId,omitempty"`
}

// Do sends the GetNodeForLocation CDP command to a browser,
// and returns the browser's response.
func (t *GetNodeForLocation) Do(ctx context.Context) (*GetNodeForLocationResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.Send(ctx, "DOM.getNodeForLocation", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetNodeForLocationResult{}
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
	NodeID int64 `json:"nodeId,omitempty"`
	// Identifier of the backend node.
	BackendNodeID int64 `json:"backendNodeId,omitempty"`
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
func (t *GetOuterHTML) SetNodeID(v int64) *GetOuterHTML {
	t.NodeID = v
	return t
}

// SetBackendNodeID adds or modifies the value of the optional
// parameter `backendNodeId` in the GetOuterHTML CDP command.
//
// Identifier of the backend node.
func (t *GetOuterHTML) SetBackendNodeID(v int64) *GetOuterHTML {
	t.BackendNodeID = v
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

// GetOuterHTMLResult contains the browser's response
// to calling the GetOuterHTML CDP command with Do().
type GetOuterHTMLResult struct {
	// Outer HTML markup.
	OuterHTML string `json:"outerHTML"`
}

// Do sends the GetOuterHTML CDP command to a browser,
// and returns the browser's response.
func (t *GetOuterHTML) Do(ctx context.Context) (*GetOuterHTMLResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.Send(ctx, "DOM.getOuterHTML", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetOuterHTMLResult{}
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
	NodeID int64 `json:"nodeId"`
}

// NewGetRelayoutBoundary constructs a new GetRelayoutBoundary struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getRelayoutBoundary
//
// This CDP method is experimental.
func NewGetRelayoutBoundary(nodeID int64) *GetRelayoutBoundary {
	return &GetRelayoutBoundary{
		NodeID: nodeID,
	}
}

// GetRelayoutBoundaryResult contains the browser's response
// to calling the GetRelayoutBoundary CDP command with Do().
type GetRelayoutBoundaryResult struct {
	// Relayout boundary node id for the given node.
	NodeID int64 `json:"nodeId"`
}

// Do sends the GetRelayoutBoundary CDP command to a browser,
// and returns the browser's response.
func (t *GetRelayoutBoundary) Do(ctx context.Context) (*GetRelayoutBoundaryResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.Send(ctx, "DOM.getRelayoutBoundary", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetRelayoutBoundaryResult{}
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
func NewGetSearchResults(searchID string, fromIndex int64, toIndex int64) *GetSearchResults {
	return &GetSearchResults{
		SearchID:  searchID,
		FromIndex: fromIndex,
		ToIndex:   toIndex,
	}
}

// GetSearchResultsResult contains the browser's response
// to calling the GetSearchResults CDP command with Do().
type GetSearchResultsResult struct {
	// Ids of the search result nodes.
	NodeIds []NodeID `json:"nodeIds"`
}

// Do sends the GetSearchResults CDP command to a browser,
// and returns the browser's response.
func (t *GetSearchResults) Do(ctx context.Context) (*GetSearchResultsResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.Send(ctx, "DOM.getSearchResults", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetSearchResultsResult{}
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
	response, err := devtools.Send(ctx, "DOM.markUndoableState", nil)
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
	NodeID int64 `json:"nodeId"`
	// Id of the element to drop the moved node into.
	TargetNodeID int64 `json:"targetNodeId"`
	// Drop node before this one (if absent, the moved node becomes the last child of
	// `targetNodeId`).
	InsertBeforeNodeID int64 `json:"insertBeforeNodeId,omitempty"`
}

// NewMoveTo constructs a new MoveTo struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-moveTo
func NewMoveTo(nodeID int64, targetNodeID int64) *MoveTo {
	return &MoveTo{
		NodeID:       nodeID,
		TargetNodeID: targetNodeID,
	}
}

// SetInsertBeforeNodeID adds or modifies the value of the optional
// parameter `insertBeforeNodeId` in the MoveTo CDP command.
//
// Drop node before this one (if absent, the moved node becomes the last child of
// `targetNodeId`).
func (t *MoveTo) SetInsertBeforeNodeID(v int64) *MoveTo {
	t.InsertBeforeNodeID = v
	return t
}

// MoveToResult contains the browser's response
// to calling the MoveTo CDP command with Do().
type MoveToResult struct {
	// New id of the moved node.
	NodeID int64 `json:"nodeId"`
}

// Do sends the MoveTo CDP command to a browser,
// and returns the browser's response.
func (t *MoveTo) Do(ctx context.Context) (*MoveToResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.Send(ctx, "DOM.moveTo", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &MoveToResult{}
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

// PerformSearchResult contains the browser's response
// to calling the PerformSearch CDP command with Do().
type PerformSearchResult struct {
	// Unique search session identifier.
	SearchID string `json:"searchId"`
	// Number of search results.
	ResultCount int64 `json:"resultCount"`
}

// Do sends the PerformSearch CDP command to a browser,
// and returns the browser's response.
func (t *PerformSearch) Do(ctx context.Context) (*PerformSearchResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.Send(ctx, "DOM.performSearch", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &PerformSearchResult{}
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

// PushNodeByPathToFrontendResult contains the browser's response
// to calling the PushNodeByPathToFrontend CDP command with Do().
type PushNodeByPathToFrontendResult struct {
	// Id of the node for given path.
	NodeID int64 `json:"nodeId"`
}

// Do sends the PushNodeByPathToFrontend CDP command to a browser,
// and returns the browser's response.
func (t *PushNodeByPathToFrontend) Do(ctx context.Context) (*PushNodeByPathToFrontendResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.Send(ctx, "DOM.pushNodeByPathToFrontend", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &PushNodeByPathToFrontendResult{}
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

// PushNodesByBackendIdsToFrontendResult contains the browser's response
// to calling the PushNodesByBackendIdsToFrontend CDP command with Do().
type PushNodesByBackendIdsToFrontendResult struct {
	// The array of ids of pushed nodes that correspond to the backend ids specified in
	// backendNodeIds.
	NodeIds []NodeID `json:"nodeIds"`
}

// Do sends the PushNodesByBackendIdsToFrontend CDP command to a browser,
// and returns the browser's response.
func (t *PushNodesByBackendIdsToFrontend) Do(ctx context.Context) (*PushNodesByBackendIdsToFrontendResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.Send(ctx, "DOM.pushNodesByBackendIdsToFrontend", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &PushNodesByBackendIdsToFrontendResult{}
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
	NodeID int64 `json:"nodeId"`
	// Selector string.
	Selector string `json:"selector"`
}

// NewQuerySelector constructs a new QuerySelector struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-querySelector
func NewQuerySelector(nodeID int64, selector string) *QuerySelector {
	return &QuerySelector{
		NodeID:   nodeID,
		Selector: selector,
	}
}

// QuerySelectorResult contains the browser's response
// to calling the QuerySelector CDP command with Do().
type QuerySelectorResult struct {
	// Query selector result.
	NodeID int64 `json:"nodeId"`
}

// Do sends the QuerySelector CDP command to a browser,
// and returns the browser's response.
func (t *QuerySelector) Do(ctx context.Context) (*QuerySelectorResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.Send(ctx, "DOM.querySelector", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &QuerySelectorResult{}
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
	NodeID int64 `json:"nodeId"`
	// Selector string.
	Selector string `json:"selector"`
}

// NewQuerySelectorAll constructs a new QuerySelectorAll struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-querySelectorAll
func NewQuerySelectorAll(nodeID int64, selector string) *QuerySelectorAll {
	return &QuerySelectorAll{
		NodeID:   nodeID,
		Selector: selector,
	}
}

// QuerySelectorAllResult contains the browser's response
// to calling the QuerySelectorAll CDP command with Do().
type QuerySelectorAllResult struct {
	// Query selector result.
	NodeIds []NodeID `json:"nodeIds"`
}

// Do sends the QuerySelectorAll CDP command to a browser,
// and returns the browser's response.
func (t *QuerySelectorAll) Do(ctx context.Context) (*QuerySelectorAllResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.Send(ctx, "DOM.querySelectorAll", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &QuerySelectorAllResult{}
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
	response, err := devtools.Send(ctx, "DOM.redo", nil)
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
	NodeID int64 `json:"nodeId"`
	// Name of the attribute to remove.
	Name string `json:"name"`
}

// NewRemoveAttribute constructs a new RemoveAttribute struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-removeAttribute
func NewRemoveAttribute(nodeID int64, name string) *RemoveAttribute {
	return &RemoveAttribute{
		NodeID: nodeID,
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
	response, err := devtools.Send(ctx, "DOM.removeAttribute", b)
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
	NodeID int64 `json:"nodeId"`
}

// NewRemoveNode constructs a new RemoveNode struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-removeNode
func NewRemoveNode(nodeID int64) *RemoveNode {
	return &RemoveNode{
		NodeID: nodeID,
	}
}

// Do sends the RemoveNode CDP command to a browser,
// and returns the browser's response.
func (t *RemoveNode) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := devtools.Send(ctx, "DOM.removeNode", b)
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
	NodeID int64 `json:"nodeId"`
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
func NewRequestChildNodes(nodeID int64) *RequestChildNodes {
	return &RequestChildNodes{
		NodeID: nodeID,
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
	response, err := devtools.Send(ctx, "DOM.requestChildNodes", b)
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
func NewRequestNode(objectID runtime.RemoteObjectID) *RequestNode {
	return &RequestNode{
		ObjectID: objectID,
	}
}

// RequestNodeResult contains the browser's response
// to calling the RequestNode CDP command with Do().
type RequestNodeResult struct {
	// Node id for given object.
	NodeID int64 `json:"nodeId"`
}

// Do sends the RequestNode CDP command to a browser,
// and returns the browser's response.
func (t *RequestNode) Do(ctx context.Context) (*RequestNodeResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.Send(ctx, "DOM.requestNode", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &RequestNodeResult{}
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
	NodeID int64 `json:"nodeId,omitempty"`
	// Backend identifier of the node to resolve.
	BackendNodeID int64 `json:"backendNodeId,omitempty"`
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
func (t *ResolveNode) SetNodeID(v int64) *ResolveNode {
	t.NodeID = v
	return t
}

// SetBackendNodeID adds or modifies the value of the optional
// parameter `backendNodeId` in the ResolveNode CDP command.
//
// Backend identifier of the node to resolve.
func (t *ResolveNode) SetBackendNodeID(v int64) *ResolveNode {
	t.BackendNodeID = v
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

// ResolveNodeResult contains the browser's response
// to calling the ResolveNode CDP command with Do().
type ResolveNodeResult struct {
	// JavaScript object wrapper for given node.
	Object runtime.RemoteObject `json:"object"`
}

// Do sends the ResolveNode CDP command to a browser,
// and returns the browser's response.
func (t *ResolveNode) Do(ctx context.Context) (*ResolveNodeResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.Send(ctx, "DOM.resolveNode", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &ResolveNodeResult{}
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
	NodeID int64 `json:"nodeId"`
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
func NewSetAttributeValue(nodeID int64, name string, value string) *SetAttributeValue {
	return &SetAttributeValue{
		NodeID: nodeID,
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
	response, err := devtools.Send(ctx, "DOM.setAttributeValue", b)
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
	NodeID int64 `json:"nodeId"`
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
func NewSetAttributesAsText(nodeID int64, text string) *SetAttributesAsText {
	return &SetAttributesAsText{
		NodeID: nodeID,
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
	response, err := devtools.Send(ctx, "DOM.setAttributesAsText", b)
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
	NodeID int64 `json:"nodeId,omitempty"`
	// Identifier of the backend node.
	BackendNodeID int64 `json:"backendNodeId,omitempty"`
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
func (t *SetFileInputFiles) SetNodeID(v int64) *SetFileInputFiles {
	t.NodeID = v
	return t
}

// SetBackendNodeID adds or modifies the value of the optional
// parameter `backendNodeId` in the SetFileInputFiles CDP command.
//
// Identifier of the backend node.
func (t *SetFileInputFiles) SetBackendNodeID(v int64) *SetFileInputFiles {
	t.BackendNodeID = v
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
	response, err := devtools.Send(ctx, "DOM.setFileInputFiles", b)
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
	response, err := devtools.Send(ctx, "DOM.setNodeStackTracesEnabled", b)
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
	NodeID int64 `json:"nodeId"`
}

// NewGetNodeStackTraces constructs a new GetNodeStackTraces struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getNodeStackTraces
//
// This CDP method is experimental.
func NewGetNodeStackTraces(nodeID int64) *GetNodeStackTraces {
	return &GetNodeStackTraces{
		NodeID: nodeID,
	}
}

// GetNodeStackTracesResult contains the browser's response
// to calling the GetNodeStackTraces CDP command with Do().
type GetNodeStackTracesResult struct {
	// Creation stack trace, if available.
	Creation *runtime.StackTrace `json:"creation,omitempty"`
}

// Do sends the GetNodeStackTraces CDP command to a browser,
// and returns the browser's response.
func (t *GetNodeStackTraces) Do(ctx context.Context) (*GetNodeStackTracesResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.Send(ctx, "DOM.getNodeStackTraces", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetNodeStackTracesResult{}
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
func NewGetFileInfo(objectID runtime.RemoteObjectID) *GetFileInfo {
	return &GetFileInfo{
		ObjectID: objectID,
	}
}

// GetFileInfoResult contains the browser's response
// to calling the GetFileInfo CDP command with Do().
type GetFileInfoResult struct {
	Path string `json:"path"`
}

// Do sends the GetFileInfo CDP command to a browser,
// and returns the browser's response.
func (t *GetFileInfo) Do(ctx context.Context) (*GetFileInfoResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.Send(ctx, "DOM.getFileInfo", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetFileInfoResult{}
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
	NodeID int64 `json:"nodeId"`
}

// NewSetInspectedNode constructs a new SetInspectedNode struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setInspectedNode
//
// This CDP method is experimental.
func NewSetInspectedNode(nodeID int64) *SetInspectedNode {
	return &SetInspectedNode{
		NodeID: nodeID,
	}
}

// Do sends the SetInspectedNode CDP command to a browser,
// and returns the browser's response.
func (t *SetInspectedNode) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := devtools.Send(ctx, "DOM.setInspectedNode", b)
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
	NodeID int64 `json:"nodeId"`
	// New node's name.
	Name string `json:"name"`
}

// NewSetNodeName constructs a new SetNodeName struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setNodeName
func NewSetNodeName(nodeID int64, name string) *SetNodeName {
	return &SetNodeName{
		NodeID: nodeID,
		Name:   name,
	}
}

// SetNodeNameResult contains the browser's response
// to calling the SetNodeName CDP command with Do().
type SetNodeNameResult struct {
	// New node's id.
	NodeID int64 `json:"nodeId"`
}

// Do sends the SetNodeName CDP command to a browser,
// and returns the browser's response.
func (t *SetNodeName) Do(ctx context.Context) (*SetNodeNameResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.Send(ctx, "DOM.setNodeName", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &SetNodeNameResult{}
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
	NodeID int64 `json:"nodeId"`
	// New node's value.
	Value string `json:"value"`
}

// NewSetNodeValue constructs a new SetNodeValue struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setNodeValue
func NewSetNodeValue(nodeID int64, value string) *SetNodeValue {
	return &SetNodeValue{
		NodeID: nodeID,
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
	response, err := devtools.Send(ctx, "DOM.setNodeValue", b)
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
	NodeID int64 `json:"nodeId"`
	// Outer HTML markup to set.
	OuterHTML string `json:"outerHTML"`
}

// NewSetOuterHTML constructs a new SetOuterHTML struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setOuterHTML
func NewSetOuterHTML(nodeID int64, outerHTML string) *SetOuterHTML {
	return &SetOuterHTML{
		NodeID:    nodeID,
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
	response, err := devtools.Send(ctx, "DOM.setOuterHTML", b)
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
	response, err := devtools.Send(ctx, "DOM.undo", nil)
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
	FrameID string `json:"frameId"`
}

// NewGetFrameOwner constructs a new GetFrameOwner struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getFrameOwner
//
// This CDP method is experimental.
func NewGetFrameOwner(frameID string) *GetFrameOwner {
	return &GetFrameOwner{
		FrameID: frameID,
	}
}

// GetFrameOwnerResult contains the browser's response
// to calling the GetFrameOwner CDP command with Do().
type GetFrameOwnerResult struct {
	// Resulting node.
	BackendNodeID int64 `json:"backendNodeId"`
	// Id of the node at given coordinates, only when enabled and requested document.
	NodeID int64 `json:"nodeId,omitempty"`
}

// Do sends the GetFrameOwner CDP command to a browser,
// and returns the browser's response.
func (t *GetFrameOwner) Do(ctx context.Context) (*GetFrameOwnerResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.Send(ctx, "DOM.getFrameOwner", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetFrameOwnerResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}
