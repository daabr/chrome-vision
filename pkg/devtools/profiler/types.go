package profiler

import "github.com/daabr/chrome-vision/pkg/devtools/runtime"

// ProfileNode data type. Profile node. Holds callsite information, execution statistics and child nodes.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#type-ProfileNode
type ProfileNode struct {
	// Unique id of the node.
	ID int64 `json:"id"`
	// Function location.
	CallFrame runtime.CallFrame `json:"callFrame"`
	// Number of samples where this node was on top of the call stack.
	HitCount int64 `json:"hitCount,omitempty"`
	// Child node ids.
	Children []int64 `json:"children,omitempty"`
	// The reason of being not optimized. The function may be deoptimized or marked as don't
	// optimize.
	DeoptReason string `json:"deoptReason,omitempty"`
	// An array of source position ticks.
	PositionTicks []PositionTickInfo `json:"positionTicks,omitempty"`
}

// Profile data type. Profile.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#type-Profile
type Profile struct {
	// The list of profile nodes. First item is the root node.
	Nodes []ProfileNode `json:"nodes"`
	// Profiling start timestamp in microseconds.
	StartTime float64 `json:"startTime"`
	// Profiling end timestamp in microseconds.
	EndTime float64 `json:"endTime"`
	// Ids of samples top nodes.
	Samples []int64 `json:"samples,omitempty"`
	// Time intervals between adjacent samples in microseconds. The first delta is relative to the
	// profile startTime.
	TimeDeltas []int64 `json:"timeDeltas,omitempty"`
}

// PositionTickInfo data type. Specifies a number of samples attributed to a certain source position.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#type-PositionTickInfo
type PositionTickInfo struct {
	// Source line number (1-based).
	Line int64 `json:"line"`
	// Number of samples attributed to the source line.
	Ticks int64 `json:"ticks"`
}

// CoverageRange data type. Coverage data for a source range.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#type-CoverageRange
type CoverageRange struct {
	// JavaScript script source offset for the range start.
	StartOffset int64 `json:"startOffset"`
	// JavaScript script source offset for the range end.
	EndOffset int64 `json:"endOffset"`
	// Collected execution count of the source range.
	Count int64 `json:"count"`
}

// FunctionCoverage data type. Coverage data for a JavaScript function.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#type-FunctionCoverage
type FunctionCoverage struct {
	// JavaScript function name.
	FunctionName string `json:"functionName"`
	// Source ranges inside the function with coverage data.
	Ranges []CoverageRange `json:"ranges"`
	// Whether coverage data for this function has block granularity.
	IsBlockCoverage bool `json:"isBlockCoverage"`
}

// ScriptCoverage data type. Coverage data for a JavaScript script.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#type-ScriptCoverage
type ScriptCoverage struct {
	// JavaScript script id.
	ScriptID string `json:"scriptId"`
	// JavaScript script name or url.
	URL string `json:"url"`
	// Functions contained in the script that has coverage data.
	Functions []FunctionCoverage `json:"functions"`
}

// TypeObject data type. Describes a type collected during runtime.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#type-TypeObject
//
// This CDP type is experimental.
type TypeObject struct {
	// Name of a type collected with type profiling.
	Name string `json:"name"`
}

// TypeProfileEntry data type. Source offset and types for a parameter or return value.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#type-TypeProfileEntry
//
// This CDP type is experimental.
type TypeProfileEntry struct {
	// Source offset of the parameter or end of function for return values.
	Offset int64 `json:"offset"`
	// The types for this parameter or return value.
	Types []TypeObject `json:"types"`
}

// ScriptTypeProfile data type. Type profile data collected during runtime for a JavaScript script.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#type-ScriptTypeProfile
//
// This CDP type is experimental.
type ScriptTypeProfile struct {
	// JavaScript script id.
	ScriptID string `json:"scriptId"`
	// JavaScript script name or url.
	URL string `json:"url"`
	// Type profile entries for parameters and return values of the functions in the script.
	Entries []TypeProfileEntry `json:"entries"`
}
