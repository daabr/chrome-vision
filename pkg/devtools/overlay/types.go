package overlay

import "github.com/daabr/chrome-vision/pkg/devtools/dom"

// SourceOrderConfig data type. Configuration data for drawing the source order of an elements children.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#type-SourceOrderConfig
type SourceOrderConfig struct {
	// the color to outline the givent element in.
	ParentOutlineColor dom.RGBA `json:"parentOutlineColor"`
	// the color to outline the child elements in.
	ChildOutlineColor dom.RGBA `json:"childOutlineColor"`
}

// GridHighlightConfig data type. Configuration data for the highlighting of Grid elements.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#type-GridHighlightConfig
type GridHighlightConfig struct {
	// Whether the extension lines from grid cells to the rulers should be shown (default: false).
	ShowGridExtensionLines bool `json:"showGridExtensionLines,omitempty"`
	// Show Positive line number labels (default: false).
	ShowPositiveLineNumbers bool `json:"showPositiveLineNumbers,omitempty"`
	// Show Negative line number labels (default: false).
	ShowNegativeLineNumbers bool `json:"showNegativeLineNumbers,omitempty"`
	// Show area name labels (default: false).
	ShowAreaNames bool `json:"showAreaNames,omitempty"`
	// Show line name labels (default: false).
	ShowLineNames bool `json:"showLineNames,omitempty"`
	// Show track size labels (default: false).
	ShowTrackSizes bool `json:"showTrackSizes,omitempty"`
	// The grid container border highlight color (default: transparent).
	GridBorderColor *dom.RGBA `json:"gridBorderColor,omitempty"`
	// The cell border color (default: transparent). Deprecated, please use rowLineColor and columnLineColor instead.
	//
	// This CDP property is deprecated.
	CellBorderColor *dom.RGBA `json:"cellBorderColor,omitempty"`
	// The row line color (default: transparent).
	RowLineColor *dom.RGBA `json:"rowLineColor,omitempty"`
	// The column line color (default: transparent).
	ColumnLineColor *dom.RGBA `json:"columnLineColor,omitempty"`
	// Whether the grid border is dashed (default: false).
	GridBorderDash bool `json:"gridBorderDash,omitempty"`
	// Whether the cell border is dashed (default: false). Deprecated, please us rowLineDash and columnLineDash instead.
	//
	// This CDP property is deprecated.
	CellBorderDash bool `json:"cellBorderDash,omitempty"`
	// Whether row lines are dashed (default: false).
	RowLineDash bool `json:"rowLineDash,omitempty"`
	// Whether column lines are dashed (default: false).
	ColumnLineDash bool `json:"columnLineDash,omitempty"`
	// The row gap highlight fill color (default: transparent).
	RowGapColor *dom.RGBA `json:"rowGapColor,omitempty"`
	// The row gap hatching fill color (default: transparent).
	RowHatchColor *dom.RGBA `json:"rowHatchColor,omitempty"`
	// The column gap highlight fill color (default: transparent).
	ColumnGapColor *dom.RGBA `json:"columnGapColor,omitempty"`
	// The column gap hatching fill color (default: transparent).
	ColumnHatchColor *dom.RGBA `json:"columnHatchColor,omitempty"`
	// The named grid areas border color (Default: transparent).
	AreaBorderColor *dom.RGBA `json:"areaBorderColor,omitempty"`
	// The grid container background color (Default: transparent).
	GridBackgroundColor *dom.RGBA `json:"gridBackgroundColor,omitempty"`
}

// FlexContainerHighlightConfig data type. Configuration data for the highlighting of Flex container elements.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#type-FlexContainerHighlightConfig
type FlexContainerHighlightConfig struct {
	// The style of the container border
	ContainerBorder *LineStyle `json:"containerBorder,omitempty"`
	// The style of the separator between lines
	LineSeparator *LineStyle `json:"lineSeparator,omitempty"`
	// The style of the separator between items
	ItemSeparator *LineStyle `json:"itemSeparator,omitempty"`
	// Style of content-distribution space on the main axis (justify-content).
	MainDistributedSpace *BoxStyle `json:"mainDistributedSpace,omitempty"`
	// Style of content-distribution space on the cross axis (align-content).
	CrossDistributedSpace *BoxStyle `json:"crossDistributedSpace,omitempty"`
	// Style of empty space caused by row gaps (gap/row-gap).
	RowGapSpace *BoxStyle `json:"rowGapSpace,omitempty"`
	// Style of empty space caused by columns gaps (gap/column-gap).
	ColumnGapSpace *BoxStyle `json:"columnGapSpace,omitempty"`
	// Style of the self-alignment line (align-items).
	CrossAlignment *LineStyle `json:"crossAlignment,omitempty"`
}

// FlexItemHighlightConfig data type. Configuration data for the highlighting of Flex item elements.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#type-FlexItemHighlightConfig
type FlexItemHighlightConfig struct {
	// Style of the box representing the item's base size
	BaseSizeBox *BoxStyle `json:"baseSizeBox,omitempty"`
	// Style of the border around the box representing the item's base size
	BaseSizeBorder *LineStyle `json:"baseSizeBorder,omitempty"`
	// Style of the arrow representing if the item grew or shrank
	FlexibilityArrow *LineStyle `json:"flexibilityArrow,omitempty"`
}

// LineStyle data type. Style information for drawing a line.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#type-LineStyle
type LineStyle struct {
	// The color of the line (default: transparent)
	Color *dom.RGBA `json:"color,omitempty"`
	// The line pattern (default: solid)
	Pattern string `json:"pattern,omitempty"`
}

// BoxStyle data type. Style information for drawing a box.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#type-BoxStyle
type BoxStyle struct {
	// The background color for the box (default: transparent)
	FillColor *dom.RGBA `json:"fillColor,omitempty"`
	// The hatching color for the box (default: transparent)
	HatchColor *dom.RGBA `json:"hatchColor,omitempty"`
}

// ContrastAlgorithm data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#type-ContrastAlgorithm
type ContrastAlgorithm string

// ContrastAlgorithm valid values.
const (
	ContrastAlgorithmAa   ContrastAlgorithm = "aa"
	ContrastAlgorithmAaa  ContrastAlgorithm = "aaa"
	ContrastAlgorithmApca ContrastAlgorithm = "apca"
)

// String returns the ContrastAlgorithm value as a built-in string.
func (t ContrastAlgorithm) String() string {
	return string(t)
}

// HighlightConfig data type. Configuration data for the highlighting of page elements.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#type-HighlightConfig
type HighlightConfig struct {
	// Whether the node info tooltip should be shown (default: false).
	ShowInfo bool `json:"showInfo,omitempty"`
	// Whether the node styles in the tooltip (default: false).
	ShowStyles bool `json:"showStyles,omitempty"`
	// Whether the rulers should be shown (default: false).
	ShowRulers bool `json:"showRulers,omitempty"`
	// Whether the a11y info should be shown (default: true).
	ShowAccessibilityInfo bool `json:"showAccessibilityInfo,omitempty"`
	// Whether the extension lines from node to the rulers should be shown (default: false).
	ShowExtensionLines bool `json:"showExtensionLines,omitempty"`
	// The content box highlight fill color (default: transparent).
	ContentColor *dom.RGBA `json:"contentColor,omitempty"`
	// The padding highlight fill color (default: transparent).
	PaddingColor *dom.RGBA `json:"paddingColor,omitempty"`
	// The border highlight fill color (default: transparent).
	BorderColor *dom.RGBA `json:"borderColor,omitempty"`
	// The margin highlight fill color (default: transparent).
	MarginColor *dom.RGBA `json:"marginColor,omitempty"`
	// The event target element highlight fill color (default: transparent).
	EventTargetColor *dom.RGBA `json:"eventTargetColor,omitempty"`
	// The shape outside fill color (default: transparent).
	ShapeColor *dom.RGBA `json:"shapeColor,omitempty"`
	// The shape margin fill color (default: transparent).
	ShapeMarginColor *dom.RGBA `json:"shapeMarginColor,omitempty"`
	// The grid layout color (default: transparent).
	CSSGridColor *dom.RGBA `json:"cssGridColor,omitempty"`
	// The color format used to format color styles (default: hex).
	ColorFormat *ColorFormat `json:"colorFormat,omitempty"`
	// The grid layout highlight configuration (default: all transparent).
	GridHighlightConfig *GridHighlightConfig `json:"gridHighlightConfig,omitempty"`
	// The flex container highlight configuration (default: all transparent).
	FlexContainerHighlightConfig *FlexContainerHighlightConfig `json:"flexContainerHighlightConfig,omitempty"`
	// The flex item highlight configuration (default: all transparent).
	FlexItemHighlightConfig *FlexItemHighlightConfig `json:"flexItemHighlightConfig,omitempty"`
	// The contrast algorithm to use for the contrast ratio (default: aa).
	ContrastAlgorithm *ContrastAlgorithm `json:"contrastAlgorithm,omitempty"`
	// The container query container highlight configuration (default: all transparent).
	ContainerQueryContainerHighlightConfig *ContainerQueryContainerHighlightConfig `json:"containerQueryContainerHighlightConfig,omitempty"`
}

// ColorFormat data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#type-ColorFormat
type ColorFormat string

// ColorFormat valid values.
const (
	ColorFormatRgb ColorFormat = "rgb"
	ColorFormatHsl ColorFormat = "hsl"
	ColorFormatHex ColorFormat = "hex"
)

// String returns the ColorFormat value as a built-in string.
func (t ColorFormat) String() string {
	return string(t)
}

// GridNodeHighlightConfig data type. Configurations for Persistent Grid Highlight
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#type-GridNodeHighlightConfig
type GridNodeHighlightConfig struct {
	// A descriptor for the highlight appearance.
	GridHighlightConfig GridHighlightConfig `json:"gridHighlightConfig"`
	// Identifier of the node to highlight.
	NodeID int64 `json:"nodeId"`
}

// FlexNodeHighlightConfig data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#type-FlexNodeHighlightConfig
type FlexNodeHighlightConfig struct {
	// A descriptor for the highlight appearance of flex containers.
	FlexContainerHighlightConfig FlexContainerHighlightConfig `json:"flexContainerHighlightConfig"`
	// Identifier of the node to highlight.
	NodeID int64 `json:"nodeId"`
}

// ScrollSnapContainerHighlightConfig data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#type-ScrollSnapContainerHighlightConfig
type ScrollSnapContainerHighlightConfig struct {
	// The style of the snapport border (default: transparent)
	SnapportBorder *LineStyle `json:"snapportBorder,omitempty"`
	// The style of the snap area border (default: transparent)
	SnapAreaBorder *LineStyle `json:"snapAreaBorder,omitempty"`
	// The margin highlight fill color (default: transparent).
	ScrollMarginColor *dom.RGBA `json:"scrollMarginColor,omitempty"`
	// The padding highlight fill color (default: transparent).
	ScrollPaddingColor *dom.RGBA `json:"scrollPaddingColor,omitempty"`
}

// ScrollSnapHighlightConfig data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#type-ScrollSnapHighlightConfig
type ScrollSnapHighlightConfig struct {
	// A descriptor for the highlight appearance of scroll snap containers.
	ScrollSnapContainerHighlightConfig ScrollSnapContainerHighlightConfig `json:"scrollSnapContainerHighlightConfig"`
	// Identifier of the node to highlight.
	NodeID int64 `json:"nodeId"`
}

// HingeConfig data type. Configuration for dual screen hinge
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#type-HingeConfig
type HingeConfig struct {
	// A rectangle represent hinge
	Rect dom.Rect `json:"rect"`
	// The content box highlight fill color (default: a dark color).
	ContentColor *dom.RGBA `json:"contentColor,omitempty"`
	// The content box highlight outline color (default: transparent).
	OutlineColor *dom.RGBA `json:"outlineColor,omitempty"`
}

// ContainerQueryHighlightConfig data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#type-ContainerQueryHighlightConfig
type ContainerQueryHighlightConfig struct {
	// A descriptor for the highlight appearance of container query containers.
	ContainerQueryContainerHighlightConfig ContainerQueryContainerHighlightConfig `json:"containerQueryContainerHighlightConfig"`
	// Identifier of the container node to highlight.
	NodeID int64 `json:"nodeId"`
}

// ContainerQueryContainerHighlightConfig data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#type-ContainerQueryContainerHighlightConfig
type ContainerQueryContainerHighlightConfig struct {
	// The style of the container border
	ContainerBorder *LineStyle `json:"containerBorder,omitempty"`
}

// InspectMode data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#type-InspectMode
type InspectMode string

// InspectMode valid values.
const (
	InspectModeSearchForNode         InspectMode = "searchForNode"
	InspectModeSearchForUAShadowDOM  InspectMode = "searchForUAShadowDOM"
	InspectModeCaptureAreaScreenshot InspectMode = "captureAreaScreenshot"
	InspectModeShowDistances         InspectMode = "showDistances"
	InspectModeNone                  InspectMode = "none"
)

// String returns the InspectMode value as a built-in string.
func (t InspectMode) String() string {
	return string(t)
}
