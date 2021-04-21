package domsnapshot

import (
	"github.com/daabr/chrome-vision/pkg/cdp/dom"
	"github.com/daabr/chrome-vision/pkg/cdp/domdebugger"
)

// A Node in the DOM tree.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/#type-DOMNode
type DOMNode struct {
	// `Node`'s nodeType.
	NodeType int64 `json:"nodeType"`
	// `Node`'s nodeName.
	NodeName string `json:"nodeName"`
	// `Node`'s nodeValue.
	NodeValue string `json:"nodeValue"`
	// Only set for textarea elements, contains the text value.
	TextValue string `json:"textValue,omitempty"`
	// Only set for input elements, contains the input's associated text value.
	InputValue string `json:"inputValue,omitempty"`
	// Only set for radio and checkbox input elements, indicates if the element has been checked
	InputChecked bool `json:"inputChecked,omitempty"`
	// Only set for option elements, indicates if the element has been selected
	OptionSelected bool `json:"optionSelected,omitempty"`
	// `Node`'s id, corresponds to DOM.Node.backendNodeId.
	BackendNodeID int64 `json:"backendNodeId"`
	// The indexes of the node's child nodes in the `domNodes` array returned by `getSnapshot`, if
	// any.
	ChildNodeIndexes []int64 `json:"childNodeIndexes,omitempty"`
	// Attributes of an `Element` node.
	Attributes []NameValue `json:"attributes,omitempty"`
	// Indexes of pseudo elements associated with this node in the `domNodes` array returned by
	// `getSnapshot`, if any.
	PseudoElementIndexes []int64 `json:"pseudoElementIndexes,omitempty"`
	// The index of the node's related layout tree node in the `layoutTreeNodes` array returned by
	// `getSnapshot`, if any.
	LayoutNodeIndex int64 `json:"layoutNodeIndex,omitempty"`
	// Document URL that `Document` or `FrameOwner` node points to.
	DocumentURL string `json:"documentURL,omitempty"`
	// Base URL that `Document` or `FrameOwner` node uses for URL completion.
	BaseURL string `json:"baseURL,omitempty"`
	// Only set for documents, contains the document's content language.
	ContentLanguage string `json:"contentLanguage,omitempty"`
	// Only set for documents, contains the document's character set encoding.
	DocumentEncoding string `json:"documentEncoding,omitempty"`
	// `DocumentType` node's publicId.
	PublicID string `json:"publicId,omitempty"`
	// `DocumentType` node's systemId.
	SystemID string `json:"systemId,omitempty"`
	// Frame ID for frame owner elements and also for the document node.
	FrameID string `json:"frameId,omitempty"`
	// The index of a frame owner element's content document in the `domNodes` array returned by
	// `getSnapshot`, if any.
	ContentDocumentIndex int64 `json:"contentDocumentIndex,omitempty"`
	// Type of a pseudo element node.
	PseudoType string `json:"pseudoType,omitempty"`
	// Shadow root type.
	ShadowRootType string `json:"shadowRootType,omitempty"`
	// Whether this DOM node responds to mouse clicks. This includes nodes that have had click
	// event listeners attached via JavaScript as well as anchor tags that naturally navigate when
	// clicked.
	IsClickable bool `json:"isClickable,omitempty"`
	// Details of the node's event listeners, if any.
	EventListeners []domdebugger.EventListener `json:"eventListeners,omitempty"`
	// The selected url for nodes with a srcset attribute.
	CurrentSourceURL string `json:"currentSourceURL,omitempty"`
	// The url of the script (if any) that generates this node.
	OriginURL string `json:"originURL,omitempty"`
	// Scroll offsets, set when this node is a Document.
	ScrollOffsetX float64 `json:"scrollOffsetX,omitempty"`
	ScrollOffsetY float64 `json:"scrollOffsetY,omitempty"`
}

// Details of post layout rendered text positions. The exact layout should not be regarded as
// stable and may change between versions.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/#type-InlineTextBox
type InlineTextBox struct {
	// The bounding box in document coordinates. Note that scroll offset of the document is ignored.
	BoundingBox dom.Rect `json:"boundingBox"`
	// The starting index in characters, for this post layout textbox substring. Characters that
	// would be represented as a surrogate pair in UTF-16 have length 2.
	StartCharacterIndex int64 `json:"startCharacterIndex"`
	// The number of characters in this post layout textbox substring. Characters that would be
	// represented as a surrogate pair in UTF-16 have length 2.
	NumCharacters int64 `json:"numCharacters"`
}

// Details of an element in the DOM tree with a LayoutObject.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/#type-LayoutTreeNode
type LayoutTreeNode struct {
	// The index of the related DOM node in the `domNodes` array returned by `getSnapshot`.
	DomNodeIndex int64 `json:"domNodeIndex"`
	// The bounding box in document coordinates. Note that scroll offset of the document is ignored.
	BoundingBox dom.Rect `json:"boundingBox"`
	// Contents of the LayoutText, if any.
	LayoutText string `json:"layoutText,omitempty"`
	// The post-layout inline text nodes, if any.
	InlineTextNodes []InlineTextBox `json:"inlineTextNodes,omitempty"`
	// Index into the `computedStyles` array returned by `getSnapshot`.
	StyleIndex int64 `json:"styleIndex,omitempty"`
	// Global paint order index, which is determined by the stacking order of the nodes. Nodes
	// that are painted together will have the same index. Only provided if includePaintOrder in
	// getSnapshot was true.
	PaintOrder int64 `json:"paintOrder,omitempty"`
	// Set to true to indicate the element begins a new stacking context.
	IsStackingContext bool `json:"isStackingContext,omitempty"`
}

// A subset of the full ComputedStyle as defined by the request whitelist.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/#type-ComputedStyle
type ComputedStyle struct {
	// Name/value pairs of computed style properties.
	Properties []NameValue `json:"properties"`
}

// A name/value pair.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/#type-NameValue
type NameValue struct {
	// Attribute/property name.
	Name string `json:"name"`
	// Attribute/property value.
	Value string `json:"value"`
}

// Index of the string in the strings table.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/#type-StringIndex
type StringIndex int64

// Index of the string in the strings table.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/#type-ArrayOfStrings
type ArrayOfStrings []StringIndex

// Data that is only present on rare nodes.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/#type-RareStringData
type RareStringData struct {
	Index []int64       `json:"index"`
	Value []StringIndex `json:"value"`
}

// https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/#type-RareBooleanData
type RareBooleanData struct {
	Index []int64 `json:"index"`
}

// https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/#type-RareIntegerData
type RareIntegerData struct {
	Index []int64 `json:"index"`
	Value []int64 `json:"value"`
}

// https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/#type-Rectangle
type Rectangle []float64

// Document snapshot.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/#type-DocumentSnapshot
type DocumentSnapshot struct {
	// Document URL that `Document` or `FrameOwner` node points to.
	DocumentURL int64 `json:"documentURL"`
	// Document title.
	Title int64 `json:"title"`
	// Base URL that `Document` or `FrameOwner` node uses for URL completion.
	BaseURL int64 `json:"baseURL"`
	// Contains the document's content language.
	ContentLanguage int64 `json:"contentLanguage"`
	// Contains the document's character set encoding.
	EncodingName int64 `json:"encodingName"`
	// `DocumentType` node's publicId.
	PublicID int64 `json:"publicId"`
	// `DocumentType` node's systemId.
	SystemID int64 `json:"systemId"`
	// Frame ID for frame owner elements and also for the document node.
	FrameID int64 `json:"frameId"`
	// A table with dom nodes.
	Nodes NodeTreeSnapshot `json:"nodes"`
	// The nodes in the layout tree.
	Layout LayoutTreeSnapshot `json:"layout"`
	// The post-layout inline text nodes.
	TextBoxes TextBoxSnapshot `json:"textBoxes"`
	// Horizontal scroll offset.
	ScrollOffsetX float64 `json:"scrollOffsetX,omitempty"`
	// Vertical scroll offset.
	ScrollOffsetY float64 `json:"scrollOffsetY,omitempty"`
	// Document content width.
	ContentWidth float64 `json:"contentWidth,omitempty"`
	// Document content height.
	ContentHeight float64 `json:"contentHeight,omitempty"`
}

// Table containing nodes.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/#type-NodeTreeSnapshot
type NodeTreeSnapshot struct {
	// Parent node index.
	ParentIndex []int64 `json:"parentIndex,omitempty"`
	// `Node`'s nodeType.
	NodeType []int64 `json:"nodeType,omitempty"`
	// `Node`'s nodeName.
	NodeName []StringIndex `json:"nodeName,omitempty"`
	// `Node`'s nodeValue.
	NodeValue []StringIndex `json:"nodeValue,omitempty"`
	// `Node`'s id, corresponds to DOM.Node.backendNodeId.
	BackendNodeID []int64 `json:"backendNodeId,omitempty"`
	// Attributes of an `Element` node. Flatten name, value pairs.
	Attributes []ArrayOfStrings `json:"attributes,omitempty"`
	// Only set for textarea elements, contains the text value.
	TextValue *RareStringData `json:"textValue,omitempty"`
	// Only set for input elements, contains the input's associated text value.
	InputValue *RareStringData `json:"inputValue,omitempty"`
	// Only set for radio and checkbox input elements, indicates if the element has been checked
	InputChecked *RareBooleanData `json:"inputChecked,omitempty"`
	// Only set for option elements, indicates if the element has been selected
	OptionSelected *RareBooleanData `json:"optionSelected,omitempty"`
	// The index of the document in the list of the snapshot documents.
	ContentDocumentIndex *RareIntegerData `json:"contentDocumentIndex,omitempty"`
	// Type of a pseudo element node.
	PseudoType *RareStringData `json:"pseudoType,omitempty"`
	// Whether this DOM node responds to mouse clicks. This includes nodes that have had click
	// event listeners attached via JavaScript as well as anchor tags that naturally navigate when
	// clicked.
	IsClickable *RareBooleanData `json:"isClickable,omitempty"`
	// The selected url for nodes with a srcset attribute.
	CurrentSourceURL *RareStringData `json:"currentSourceURL,omitempty"`
	// The url of the script (if any) that generates this node.
	OriginURL *RareStringData `json:"originURL,omitempty"`
}

// Table of details of an element in the DOM tree with a LayoutObject.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/#type-LayoutTreeSnapshot
type LayoutTreeSnapshot struct {
	// Index of the corresponding node in the `NodeTreeSnapshot` array returned by `captureSnapshot`.
	NodeIndex []int64 `json:"nodeIndex"`
	// Array of indexes specifying computed style strings, filtered according to the `computedStyles` parameter passed to `captureSnapshot`.
	Styles []ArrayOfStrings `json:"styles"`
	// The absolute position bounding box.
	Bounds []Rectangle `json:"bounds"`
	// Contents of the LayoutText, if any.
	Text []StringIndex `json:"text"`
	// Stacking context information.
	StackingContexts RareBooleanData `json:"stackingContexts"`
	// Global paint order index, which is determined by the stacking order of the nodes. Nodes
	// that are painted together will have the same index. Only provided if includePaintOrder in
	// captureSnapshot was true.
	PaintOrders []int64 `json:"paintOrders,omitempty"`
	// The offset rect of nodes. Only available when includeDOMRects is set to true
	OffsetRects []Rectangle `json:"offsetRects,omitempty"`
	// The scroll rect of nodes. Only available when includeDOMRects is set to true
	ScrollRects []Rectangle `json:"scrollRects,omitempty"`
	// The client rect of nodes. Only available when includeDOMRects is set to true
	ClientRects []Rectangle `json:"clientRects,omitempty"`
	// The list of background colors that are blended with colors of overlapping elements.
	//
	// This CDP property is experimental.
	BlendedBackgroundColors []StringIndex `json:"blendedBackgroundColors,omitempty"`
	// The list of computed text opacities.
	//
	// This CDP property is experimental.
	TextColorOpacities []float64 `json:"textColorOpacities,omitempty"`
}

// Table of details of the post layout rendered text positions. The exact layout should not be regarded as
// stable and may change between versions.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot/#type-TextBoxSnapshot
type TextBoxSnapshot struct {
	// Index of the layout tree node that owns this box collection.
	LayoutIndex []int64 `json:"layoutIndex"`
	// The absolute position bounding box.
	Bounds []Rectangle `json:"bounds"`
	// The starting index in characters, for this post layout textbox substring. Characters that
	// would be represented as a surrogate pair in UTF-16 have length 2.
	Start []int64 `json:"start"`
	// The number of characters in this post layout textbox substring. Characters that would be
	// represented as a surrogate pair in UTF-16 have length 2.
	Length []int64 `json:"length"`
}
