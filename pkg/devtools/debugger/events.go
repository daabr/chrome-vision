package debugger

import (
	"encoding/json"

	"github.com/daabr/chrome-vision/pkg/devtools/runtime"
)

// BreakpointResolved asynchronous event. Fired when breakpoint is resolved to an actual script and location.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#event-breakpointResolved
type BreakpointResolved struct {
	// Breakpoint unique identifier.
	BreakpointID string `json:"breakpointId"`
	// Actual breakpoint location.
	Location Location `json:"location"`
}

// Paused asynchronous event. Fired when the virtual machine stopped on breakpoint or exception or any other stop criteria.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#event-paused
type Paused struct {
	// Call stack the virtual machine stopped on.
	CallFrames []CallFrame `json:"callFrames"`
	// Pause reason.
	Reason string `json:"reason"`
	// Object containing break-specific auxiliary properties.
	Data json.RawMessage `json:"data,omitempty"`
	// Hit breakpoints IDs
	HitBreakpoints []string `json:"hitBreakpoints,omitempty"`
	// Async stack trace, if any.
	AsyncStackTrace *runtime.StackTrace `json:"asyncStackTrace,omitempty"`
	// Async stack trace, if any.
	//
	// This CDP parameter is experimental.
	AsyncStackTraceID *runtime.StackTraceID `json:"asyncStackTraceId,omitempty"`
	// Never present, will be removed.
	//
	// This CDP parameter is deprecated.
	// This CDP parameter is experimental.
	AsyncCallStackTraceID *runtime.StackTraceID `json:"asyncCallStackTraceId,omitempty"`
}

// Resumed asynchronous event. Fired when the virtual machine resumed execution.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#event-resumed
type Resumed struct{}

// ScriptFailedToParse asynchronous event. Fired when virtual machine fails to parse the script.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#event-scriptFailedToParse
type ScriptFailedToParse struct {
	// Identifier of the script parsed.
	ScriptID string `json:"scriptId"`
	// URL or name of the script parsed (if any).
	URL string `json:"url"`
	// Line offset of the script within the resource with given URL (for script tags).
	StartLine int64 `json:"startLine"`
	// Column offset of the script within the resource with given URL.
	StartColumn int64 `json:"startColumn"`
	// Last line of the script.
	EndLine int64 `json:"endLine"`
	// Length of the last line of the script.
	EndColumn int64 `json:"endColumn"`
	// Specifies script creation context.
	ExecutionContextID int64 `json:"executionContextId"`
	// Content hash of the script.
	Hash string `json:"hash"`
	// Embedder-specific auxiliary data.
	ExecutionContextAuxData json.RawMessage `json:"executionContextAuxData,omitempty"`
	// URL of source map associated with script (if any).
	SourceMapURL string `json:"sourceMapURL,omitempty"`
	// True, if this script has sourceURL.
	HasSourceURL bool `json:"hasSourceURL,omitempty"`
	// True, if this script is ES6 module.
	IsModule bool `json:"isModule,omitempty"`
	// This script length.
	Length int64 `json:"length,omitempty"`
	// JavaScript top stack frame of where the script parsed event was triggered if available.
	//
	// This CDP parameter is experimental.
	StackTrace *runtime.StackTrace `json:"stackTrace,omitempty"`
	// If the scriptLanguage is WebAssembly, the code section offset in the module.
	//
	// This CDP parameter is experimental.
	CodeOffset int64 `json:"codeOffset,omitempty"`
	// The language of the script.
	//
	// This CDP parameter is experimental.
	ScriptLanguage *ScriptLanguage `json:"scriptLanguage,omitempty"`
	// The name the embedder supplied for this script.
	//
	// This CDP parameter is experimental.
	EmbedderName string `json:"embedderName,omitempty"`
}

// ScriptParsed asynchronous event. Fired when virtual machine parses script. This event is also fired for all known and uncollected
// scripts upon enabling debugger.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#event-scriptParsed
type ScriptParsed struct {
	// Identifier of the script parsed.
	ScriptID string `json:"scriptId"`
	// URL or name of the script parsed (if any).
	URL string `json:"url"`
	// Line offset of the script within the resource with given URL (for script tags).
	StartLine int64 `json:"startLine"`
	// Column offset of the script within the resource with given URL.
	StartColumn int64 `json:"startColumn"`
	// Last line of the script.
	EndLine int64 `json:"endLine"`
	// Length of the last line of the script.
	EndColumn int64 `json:"endColumn"`
	// Specifies script creation context.
	ExecutionContextID int64 `json:"executionContextId"`
	// Content hash of the script.
	Hash string `json:"hash"`
	// Embedder-specific auxiliary data.
	ExecutionContextAuxData json.RawMessage `json:"executionContextAuxData,omitempty"`
	// True, if this script is generated as a result of the live edit operation.
	//
	// This CDP parameter is experimental.
	IsLiveEdit bool `json:"isLiveEdit,omitempty"`
	// URL of source map associated with script (if any).
	SourceMapURL string `json:"sourceMapURL,omitempty"`
	// True, if this script has sourceURL.
	HasSourceURL bool `json:"hasSourceURL,omitempty"`
	// True, if this script is ES6 module.
	IsModule bool `json:"isModule,omitempty"`
	// This script length.
	Length int64 `json:"length,omitempty"`
	// JavaScript top stack frame of where the script parsed event was triggered if available.
	//
	// This CDP parameter is experimental.
	StackTrace *runtime.StackTrace `json:"stackTrace,omitempty"`
	// If the scriptLanguage is WebAssembly, the code section offset in the module.
	//
	// This CDP parameter is experimental.
	CodeOffset int64 `json:"codeOffset,omitempty"`
	// The language of the script.
	//
	// This CDP parameter is experimental.
	ScriptLanguage *ScriptLanguage `json:"scriptLanguage,omitempty"`
	// If the scriptLanguage is WebASsembly, the source of debug symbols for the module.
	//
	// This CDP parameter is experimental.
	DebugSymbols *DebugSymbols `json:"debugSymbols,omitempty"`
	// The name the embedder supplied for this script.
	//
	// This CDP parameter is experimental.
	EmbedderName string `json:"embedderName,omitempty"`
}
