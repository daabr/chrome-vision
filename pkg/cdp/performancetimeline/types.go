package performancetimeline

import "github.com/daabr/chrome-vision/pkg/cdp/dom"

// See https://github.com/WICG/LargestContentfulPaint and largest_contentful_paint.idl
//
// https://chromedevtools.github.io/devtools-protocol/tot/PerformanceTimeline/#type-LargestContentfulPaint
type LargestContentfulPaint struct {
	RenderTime float64
	LoadTime   float64
	// The number of pixels being painted.
	Size float64
	// The id attribute of the element, if available.
	ElementID string `json:"elementId,omitempty"`
	// The URL of the image (may be trimmed).
	URL    string `json:"url,omitempty"`
	NodeID int64  `json:"nodeId,omitempty"`
}

// https://chromedevtools.github.io/devtools-protocol/tot/PerformanceTimeline/#type-LayoutShiftAttribution
type LayoutShiftAttribution struct {
	PreviousRect dom.Rect
	CurrentRect  dom.Rect
	NodeID       int64 `json:"nodeId,omitempty"`
}

// See https://wicg.github.io/layout-instability/#sec-layout-shift and layout_shift.idl
//
// https://chromedevtools.github.io/devtools-protocol/tot/PerformanceTimeline/#type-LayoutShift
type LayoutShift struct {
	// Score increment produced by this event.
	Value          float64
	HadRecentInput bool
	LastInputTime  float64
	Sources        []LayoutShiftAttribution
}

// https://chromedevtools.github.io/devtools-protocol/tot/PerformanceTimeline/#type-TimelineEvent
type TimelineEvent struct {
	// Identifies the frame that this event is related to. Empty for non-frame targets.
	FrameID string
	// The event type, as specified in https://w3c.github.io/performance-timeline/#dom-performanceentry-entrytype
	// This determines which of the optional "details" fiedls is present.
	Type string
	// Name may be empty depending on the type.
	Name string
	// Time in seconds since Epoch, monotonically increasing within document lifetime.
	Time float64
	// Event duration, if applicable.
	Duration           float64                 `json:"duration,omitempty"`
	LcpDetails         *LargestContentfulPaint `json:"lcpDetails,omitempty"`
	LayoutShiftDetails *LayoutShift            `json:"layoutShiftDetails,omitempty"`
}
