package dom

// AttributeModified asynchronous event.
//
// Fired when `Element`'s attribute is modified.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-attributeModified
type AttributeModified struct {
	// Id of the node that has changed.
	NodeID int64 `json:"nodeId"`
	// Attribute name.
	Name string `json:"name"`
	// Attribute value.
	Value string `json:"value"`
}

// AttributeRemoved asynchronous event.
//
// Fired when `Element`'s attribute is removed.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-attributeRemoved
type AttributeRemoved struct {
	// Id of the node that has changed.
	NodeID int64 `json:"nodeId"`
	// A ttribute name.
	Name string `json:"name"`
}

// CharacterDataModified asynchronous event.
//
// Mirrors `DOMCharacterDataModified` event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-characterDataModified
type CharacterDataModified struct {
	// Id of the node that has changed.
	NodeID int64 `json:"nodeId"`
	// New text value.
	CharacterData string `json:"characterData"`
}

// ChildNodeCountUpdated asynchronous event.
//
// Fired when `Container`'s child node count has changed.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-childNodeCountUpdated
type ChildNodeCountUpdated struct {
	// Id of the node that has changed.
	NodeID int64 `json:"nodeId"`
	// New node count.
	ChildNodeCount int64 `json:"childNodeCount"`
}

// ChildNodeInserted asynchronous event.
//
// Mirrors `DOMNodeInserted` event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-childNodeInserted
type ChildNodeInserted struct {
	// Id of the node that has changed.
	ParentNodeID int64 `json:"parentNodeId"`
	// If of the previous siblint.
	PreviousNodeID int64 `json:"previousNodeId"`
	// Inserted node data.
	Node Node `json:"node"`
}

// ChildNodeRemoved asynchronous event.
//
// Mirrors `DOMNodeRemoved` event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-childNodeRemoved
type ChildNodeRemoved struct {
	// Parent id.
	ParentNodeID int64 `json:"parentNodeId"`
	// Id of the node that has been removed.
	NodeID int64 `json:"nodeId"`
}

// DistributedNodesUpdated asynchronous event.
//
// Called when distrubution is changed.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-distributedNodesUpdated
//
// This CDP event is experimental.
type DistributedNodesUpdated struct {
	// Insertion point where distrubuted nodes were updated.
	InsertionPointID int64 `json:"insertionPointId"`
	// Distributed nodes for given insertion point.
	DistributedNodes []BackendNode `json:"distributedNodes"`
}

// DocumentUpdated asynchronous event.
//
// Fired when `Document` has been totally updated. Node ids are no longer valid.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-documentUpdated
type DocumentUpdated struct{}

// InlineStyleInvalidated asynchronous event.
//
// Fired when `Element`'s inline style is modified via a CSS property modification.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-inlineStyleInvalidated
//
// This CDP event is experimental.
type InlineStyleInvalidated struct {
	// Ids of the nodes for which the inline styles have been invalidated.
	NodeIds []NodeID `json:"nodeIds"`
}

// PseudoElementAdded asynchronous event.
//
// Called when a pseudo element is added to an element.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-pseudoElementAdded
//
// This CDP event is experimental.
type PseudoElementAdded struct {
	// Pseudo element's parent element id.
	ParentID int64 `json:"parentId"`
	// The added pseudo element.
	PseudoElement Node `json:"pseudoElement"`
}

// PseudoElementRemoved asynchronous event.
//
// Called when a pseudo element is removed from an element.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-pseudoElementRemoved
//
// This CDP event is experimental.
type PseudoElementRemoved struct {
	// Pseudo element's parent element id.
	ParentID int64 `json:"parentId"`
	// The removed pseudo element id.
	PseudoElementID int64 `json:"pseudoElementId"`
}

// SetChildNodes asynchronous event.
//
// Fired when backend wants to provide client with the missing DOM structure. This happens upon
// most of the calls requesting node ids.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-setChildNodes
type SetChildNodes struct {
	// Parent node id to populate with children.
	ParentID int64 `json:"parentId"`
	// Child nodes array.
	Nodes []Node `json:"nodes"`
}

// ShadowRootPopped asynchronous event.
//
// Called when shadow root is popped from the element.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-shadowRootPopped
//
// This CDP event is experimental.
type ShadowRootPopped struct {
	// Host element id.
	HostID int64 `json:"hostId"`
	// Shadow root id.
	RootID int64 `json:"rootId"`
}

// ShadowRootPushed asynchronous event.
//
// Called when shadow root is pushed into the element.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-shadowRootPushed
//
// This CDP event is experimental.
type ShadowRootPushed struct {
	// Host element id.
	HostID int64 `json:"hostId"`
	// Shadow root.
	Root Node `json:"root"`
}
