package debugger

import (
	"github.com/daabr/chrome-vision/pkg/cdp/runtime"
)

// Breakpoint identifier.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#type-BreakpointId
type BreakpointID string

// Call frame identifier.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#type-CallFrameId
type CallFrameID string

// Location in the source code.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#type-Location
type Location struct {
	// Script identifier as reported in the `Debugger.scriptParsed`.
	ScriptID string `json:"scriptId"`
	// Line number in the script (0-based).
	LineNumber int64 `json:"lineNumber"`
	// Column number in the script (0-based).
	ColumnNumber int64 `json:"columnNumber,omitempty"`
}

// Location in the source code.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#type-ScriptPosition
//
// This CDP type is experimental.
type ScriptPosition struct {
	LineNumber   int64 `json:"lineNumber"`
	ColumnNumber int64 `json:"columnNumber"`
}

// Location range within one script.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#type-LocationRange
//
// This CDP type is experimental.
type LocationRange struct {
	ScriptID string         `json:"scriptId"`
	Start    ScriptPosition `json:"start"`
	End      ScriptPosition `json:"end"`
}

// JavaScript call frame. Array of call frames form the call stack.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#type-CallFrame
type CallFrame struct {
	// Call frame identifier. This identifier is only valid while the virtual machine is paused.
	CallFrameID string `json:"callFrameId"`
	// Name of the JavaScript function called on this call frame.
	FunctionName string `json:"functionName"`
	// Location in the source code.
	FunctionLocation *Location `json:"functionLocation,omitempty"`
	// Location in the source code.
	Location Location `json:"location"`
	// JavaScript script name or url.
	URL string `json:"url"`
	// Scope chain for this call frame.
	ScopeChain []Scope `json:"scopeChain"`
	// `this` object for this call frame.
	This runtime.RemoteObject `json:"this"`
	// The value being returned, if the function is at return point.
	ReturnValue *runtime.RemoteObject `json:"returnValue,omitempty"`
}

// Scope description.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#type-Scope
type Scope struct {
	// Scope type.
	Type string `json:"type"`
	// Object representing the scope. For `global` and `with` scopes it represents the actual
	// object; for the rest of the scopes, it is artificial transient object enumerating scope
	// variables as its properties.
	Object runtime.RemoteObject `json:"object"`
	Name   string               `json:"name,omitempty"`
	// Location in the source code where scope starts
	StartLocation *Location `json:"startLocation,omitempty"`
	// Location in the source code where scope ends
	EndLocation *Location `json:"endLocation,omitempty"`
}

// Search match for resource.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#type-SearchMatch
type SearchMatch struct {
	// Line number in resource content.
	LineNumber float64 `json:"lineNumber"`
	// Line with match content.
	LineContent string `json:"lineContent"`
}

// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#type-BreakLocation
type BreakLocation struct {
	// Script identifier as reported in the `Debugger.scriptParsed`.
	ScriptID string `json:"scriptId"`
	// Line number in the script (0-based).
	LineNumber int64 `json:"lineNumber"`
	// Column number in the script (0-based).
	ColumnNumber int64  `json:"columnNumber,omitempty"`
	Type         string `json:"type,omitempty"`
}

// Enum of possible script languages.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#type-ScriptLanguage
type ScriptLanguage string

// ScriptLanguage valid values.
const (
	ScriptLanguageJavaScript  ScriptLanguage = "JavaScript"
	ScriptLanguageWebAssembly ScriptLanguage = "WebAssembly"
)

// Debug symbols available for a wasm script.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#type-DebugSymbols
type DebugSymbols struct {
	// Type of the debug symbols.
	Type string `json:"type"`
	// URL of the external symbol source.
	ExternalURL string `json:"externalURL,omitempty"`
}
