package input

// TouchPoint data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Input/#type-TouchPoint
type TouchPoint struct {
	// X coordinate of the event relative to the main frame's viewport in CSS pixels.
	X float64 `json:"x"`
	// Y coordinate of the event relative to the main frame's viewport in CSS pixels. 0 refers to
	// the top of the viewport and Y increases as it proceeds towards the bottom of the viewport.
	Y float64 `json:"y"`
	// X radius of the touch area (default: 1.0).
	RadiusX float64 `json:"radiusX,omitempty"`
	// Y radius of the touch area (default: 1.0).
	RadiusY float64 `json:"radiusY,omitempty"`
	// Rotation angle (default: 0.0).
	RotationAngle float64 `json:"rotationAngle,omitempty"`
	// Force (default: 1.0).
	Force float64 `json:"force,omitempty"`
	// The normalized tangential pressure, which has a range of [-1,1] (default: 0).
	//
	// This CDP property is experimental.
	TangentialPressure float64 `json:"tangentialPressure,omitempty"`
	// The plane angle between the Y-Z plane and the plane containing both the stylus axis and the Y axis, in degrees of the range [-90,90], a positive tiltX is to the right (default: 0)
	//
	// This CDP property is experimental.
	TiltX int64 `json:"tiltX,omitempty"`
	// The plane angle between the X-Z plane and the plane containing both the stylus axis and the X axis, in degrees of the range [-90,90], a positive tiltY is towards the user (default: 0).
	//
	// This CDP property is experimental.
	TiltY int64 `json:"tiltY,omitempty"`
	// The clockwise rotation of a pen stylus around its own major axis, in degrees in the range [0,359] (default: 0).
	//
	// This CDP property is experimental.
	Twist int64 `json:"twist,omitempty"`
	// Identifier used to track touch sources between events, must be unique within an event.
	ID float64 `json:"id,omitempty"`
}

// GestureSourceType data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Input/#type-GestureSourceType
//
// This CDP type is experimental.
type GestureSourceType string

// GestureSourceType valid values.
const (
	GestureSourceTypeDefault GestureSourceType = "default"
	GestureSourceTypeTouch   GestureSourceType = "touch"
	GestureSourceTypeMouse   GestureSourceType = "mouse"
)

// MouseButton data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Input/#type-MouseButton
type MouseButton string

// MouseButton valid values.
const (
	MouseButtonNone    MouseButton = "none"
	MouseButtonLeft    MouseButton = "left"
	MouseButtonMiddle  MouseButton = "middle"
	MouseButtonRight   MouseButton = "right"
	MouseButtonBack    MouseButton = "back"
	MouseButtonForward MouseButton = "forward"
)

// TimeSinceEpoch data type. UTC time in seconds, counted from January 1, 1970.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Input/#type-TimeSinceEpoch
type TimeSinceEpoch float64

// DragDataItem data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Input/#type-DragDataItem
//
// This CDP type is experimental.
type DragDataItem struct {
	// Mime type of the dragged data.
	MimeType string `json:"mimeType"`
	// Depending of the value of `mimeType`, it contains the dragged link,
	// text, HTML markup or any other data.
	Data string `json:"data"`
	// Title associated with a link. Only valid when `mimeType` == "text/uri-list".
	Title string `json:"title,omitempty"`
	// Stores the base URL for the contained markup. Only valid when `mimeType`
	// == "text/html".
	BaseURL string `json:"baseURL,omitempty"`
}

// DragData data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Input/#type-DragData
//
// This CDP type is experimental.
type DragData struct {
	Items []DragDataItem `json:"items"`
	// Bit field representing allowed drag operations. Copy = 1, Link = 2, Move = 16
	DragOperationsMask int64 `json:"dragOperationsMask"`
}
