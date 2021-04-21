package dom

import (
	"encoding/json"
)

// Unique DOM node identifier.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-NodeId
type NodeID int64

// Unique DOM node identifier used to reference a node that may not have been pushed to the
// front-end.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-BackendNodeId
type BackendNodeID int64

// Backend node with a friendly name.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-BackendNode
type BackendNode struct {
	// `Node`'s nodeType.
	NodeType int64 `json:"nodeType"`
	// `Node`'s nodeName.
	NodeName      string `json:"nodeName"`
	BackendNodeID int64  `json:"backendNodeId"`
}

// Pseudo element type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-PseudoType
type PseudoType string

// PseudoType valid values.
const (
	PseudoTypeFirstLine           PseudoType = "first-line"
	PseudoTypeFirstLetter         PseudoType = "first-letter"
	PseudoTypeBefore              PseudoType = "before"
	PseudoTypeAfter               PseudoType = "after"
	PseudoTypeMarker              PseudoType = "marker"
	PseudoTypeBackdrop            PseudoType = "backdrop"
	PseudoTypeSelection           PseudoType = "selection"
	PseudoTypeTargetText          PseudoType = "target-text"
	PseudoTypeSpellingError       PseudoType = "spelling-error"
	PseudoTypeGrammarError        PseudoType = "grammar-error"
	PseudoTypeFirstLineInherited  PseudoType = "first-line-inherited"
	PseudoTypeScrollbar           PseudoType = "scrollbar"
	PseudoTypeScrollbarThumb      PseudoType = "scrollbar-thumb"
	PseudoTypeScrollbarButton     PseudoType = "scrollbar-button"
	PseudoTypeScrollbarTrack      PseudoType = "scrollbar-track"
	PseudoTypeScrollbarTrackPiece PseudoType = "scrollbar-track-piece"
	PseudoTypeScrollbarCorner     PseudoType = "scrollbar-corner"
	PseudoTypeResizer             PseudoType = "resizer"
	PseudoTypeInputListButton     PseudoType = "input-list-button"
)

// Shadow root type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-ShadowRootType
type ShadowRootType string

// ShadowRootType valid values.
const (
	ShadowRootTypeUserAgent ShadowRootType = "user-agent"
	ShadowRootTypeOpen      ShadowRootType = "open"
	ShadowRootTypeClosed    ShadowRootType = "closed"
)

// DOM interaction is implemented in terms of mirror objects that represent the actual DOM nodes.
// DOMNode is a base node mirror type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-Node
type Node struct {
	// Node identifier that is passed into the rest of the DOM messages as the `nodeId`. Backend
	// will only push node with given `id` once. It is aware of all requested nodes and will only
	// fire DOM events for nodes known to the client.
	NodeID int64 `json:"nodeId"`
	// The id of the parent node if any.
	ParentID int64 `json:"parentId,omitempty"`
	// The BackendNodeId for this node.
	BackendNodeID int64 `json:"backendNodeId"`
	// `Node`'s nodeType.
	NodeType int64 `json:"nodeType"`
	// `Node`'s nodeName.
	NodeName string `json:"nodeName"`
	// `Node`'s localName.
	LocalName string `json:"localName"`
	// `Node`'s nodeValue.
	NodeValue string `json:"nodeValue"`
	// Child count for `Container` nodes.
	ChildNodeCount int64 `json:"childNodeCount,omitempty"`
	// Child nodes of this node when requested with children.
	Children []Node `json:"children,omitempty"`
	// Attributes of the `Element` node in the form of flat array `[name1, value1, name2, value2]`.
	Attributes []string `json:"attributes,omitempty"`
	// Document URL that `Document` or `FrameOwner` node points to.
	DocumentURL string `json:"documentURL,omitempty"`
	// Base URL that `Document` or `FrameOwner` node uses for URL completion.
	BaseURL string `json:"baseURL,omitempty"`
	// `DocumentType`'s publicId.
	PublicID string `json:"publicId,omitempty"`
	// `DocumentType`'s systemId.
	SystemID string `json:"systemId,omitempty"`
	// `DocumentType`'s internalSubset.
	InternalSubset string `json:"internalSubset,omitempty"`
	// `Document`'s XML version in case of XML documents.
	XmlVersion string `json:"xmlVersion,omitempty"`
	// `Attr`'s name.
	Name string `json:"name,omitempty"`
	// `Attr`'s value.
	Value string `json:"value,omitempty"`
	// Pseudo element type for this node.
	PseudoType string `json:"pseudoType,omitempty"`
	// Shadow root type.
	ShadowRootType string `json:"shadowRootType,omitempty"`
	// Frame ID for frame owner elements.
	FrameID string `json:"frameId,omitempty"`
	// Content document for frame owner elements.
	ContentDocument *Node `json:"contentDocument,omitempty"`
	// Shadow root list for given element host.
	ShadowRoots []Node `json:"shadowRoots,omitempty"`
	// Content document fragment for template elements.
	TemplateContent *Node `json:"templateContent,omitempty"`
	// Pseudo elements associated with this node.
	PseudoElements []Node `json:"pseudoElements,omitempty"`
	// Deprecated, as the HTML Imports API has been removed (crbug.com/937746).
	// This property used to return the imported document for the HTMLImport links.
	// The property is always undefined now.
	//
	// This CDP property is deprecated.
	ImportedDocument *Node `json:"importedDocument,omitempty"`
	// Distributed nodes for given insertion point.
	DistributedNodes []BackendNode `json:"distributedNodes,omitempty"`
	// Whether the node is SVG.
	IsSVG bool `json:"isSVG,omitempty"`
}

// A structure holding an RGBA color.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-RGBA
type RGBA struct {
	// The red component, in the [0-255] range.
	R int64 `json:"r"`
	// The green component, in the [0-255] range.
	G int64 `json:"g"`
	// The blue component, in the [0-255] range.
	B int64 `json:"b"`
	// The alpha component, in the [0-1] range (default: 1).
	A float64 `json:"a,omitempty"`
}

// An array of quad vertices, x immediately followed by y for each point, points clock-wise.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-Quad
type Quad []float64

// Box model.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-BoxModel
type BoxModel struct {
	// Content box
	Content Quad `json:"content"`
	// Padding box
	Padding Quad `json:"padding"`
	// Border box
	Border Quad `json:"border"`
	// Margin box
	Margin Quad `json:"margin"`
	// Node width
	Width int64 `json:"width"`
	// Node height
	Height int64 `json:"height"`
	// Shape outside coordinates
	ShapeOutside *ShapeOutsideInfo `json:"shapeOutside,omitempty"`
}

// CSS Shape Outside details.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-ShapeOutsideInfo
type ShapeOutsideInfo struct {
	// Shape bounds
	Bounds Quad `json:"bounds"`
	// Shape coordinate details
	Shape []json.RawMessage `json:"shape"`
	// Margin shape bounds
	MarginShape []json.RawMessage `json:"marginShape"`
}

// Rectangle.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-Rect
type Rect struct {
	// X coordinate
	X float64 `json:"x"`
	// Y coordinate
	Y float64 `json:"y"`
	// Rectangle width
	Width float64 `json:"width"`
	// Rectangle height
	Height float64 `json:"height"`
}

// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-CSSComputedStyleProperty
type CSSComputedStyleProperty struct {
	// Computed style property name.
	Name string `json:"name"`
	// Computed style property value.
	Value string `json:"value"`
}
