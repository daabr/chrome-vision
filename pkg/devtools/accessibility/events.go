package accessibility

// LoadComplete asynchronous event. The loadComplete event mirrors the load complete event sent by the browser to assistive
// technology when the web page has finished loading.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#event-loadComplete
//
// This CDP event is experimental.
type LoadComplete struct {
	// New document root node.
	Root AXNode `json:"root"`
}

// NodesUpdated asynchronous event. The nodesUpdated event is sent every time a previously requested node has changed the in tree.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Accessibility/#event-nodesUpdated
//
// This CDP event is experimental.
type NodesUpdated struct {
	// Updated node data.
	Nodes []AXNode `json:"nodes"`
}
