package layertree

import "github.com/daabr/chrome-vision/pkg/cdp/dom"

// Unique Layer identifier.
//
// https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#type-LayerId
type LayerID string

// Unique snapshot identifier.
//
// https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#type-SnapshotId
type SnapshotID string

// Rectangle where scrolling happens on the main thread.
//
// https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#type-ScrollRect
type ScrollRect struct {
	// Rectangle itself.
	Rect dom.Rect
	// Reason for rectangle to force scrolling on the main thread
	Type string
}

// Sticky position constraints.
//
// https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#type-StickyPositionConstraint
type StickyPositionConstraint struct {
	// Layout rectangle of the sticky element before being shifted
	StickyBoxRect dom.Rect
	// Layout rectangle of the containing block of the sticky element
	ContainingBlockRect dom.Rect
	// The nearest sticky layer that shifts the sticky box
	NearestLayerShiftingStickyBox string `json:"nearestLayerShiftingStickyBox,omitempty"`
	// The nearest sticky layer that shifts the containing block
	NearestLayerShiftingContainingBlock string `json:"nearestLayerShiftingContainingBlock,omitempty"`
}

// Serialized fragment of layer picture along with its offset within the layer.
//
// https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#type-PictureTile
type PictureTile struct {
	// Offset from owning layer left boundary
	X float64
	// Offset from owning layer top boundary
	Y float64
	// Base64-encoded snapshot data. (Encoded as a base64 string when passed over JSON)
	Picture string
}

// Information about a compositing layer.
//
// https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#type-Layer
type Layer struct {
	// The unique id for this layer.
	LayerID string
	// The id of parent (not present for root).
	ParentLayerID string `json:"parentLayerId,omitempty"`
	// The backend id for the node associated with this layer.
	BackendNodeID int64 `json:"backendNodeId,omitempty"`
	// Offset from parent layer, X coordinate.
	OffsetX float64
	// Offset from parent layer, Y coordinate.
	OffsetY float64
	// Layer width.
	Width float64
	// Layer height.
	Height float64
	// Transformation matrix for layer, default is identity matrix
	Transform []float64 `json:"transform,omitempty"`
	// Transform anchor point X, absent if no transform specified
	AnchorX float64 `json:"anchorX,omitempty"`
	// Transform anchor point Y, absent if no transform specified
	AnchorY float64 `json:"anchorY,omitempty"`
	// Transform anchor point Z, absent if no transform specified
	AnchorZ float64 `json:"anchorZ,omitempty"`
	// Indicates how many time this layer has painted.
	PaintCount int64
	// Indicates whether this layer hosts any content, rather than being used for
	// transform/scrolling purposes only.
	DrawsContent bool
	// Set if layer is not visible.
	Invisible bool `json:"invisible,omitempty"`
	// Rectangles scrolling on main thread only.
	ScrollRects []ScrollRect `json:"scrollRects,omitempty"`
	// Sticky position constraint information
	StickyPositionConstraint *StickyPositionConstraint `json:"stickyPositionConstraint,omitempty"`
}

// Array of timings, one per paint step.
//
// https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#type-PaintProfile
type PaintProfile []float64
