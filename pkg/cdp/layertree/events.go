package layertree

import "github.com/daabr/chrome-vision/pkg/cdp/dom"

// LayerPainted asynchronous event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#event-layerPainted
type LayerPainted struct {
	// The id of the painted layer.
	LayerID string `json:"layerId"`
	// Clip rectangle.
	Clip dom.Rect `json:"clip"`
}

// LayerTreeDidChange asynchronous event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#event-layerTreeDidChange
type LayerTreeDidChange struct {
	// Layer tree, absent if not in the comspositing mode.
	Layers []Layer `json:"layers,omitempty"`
}
