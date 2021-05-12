package domdebugger

import "github.com/daabr/chrome-vision/pkg/devtools/runtime"

// DOMBreakpointType data type. DOM breakpoint type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#type-DOMBreakpointType
type DOMBreakpointType string

// DOMBreakpointType valid values.
const (
	DOMBreakpointTypeSubtreeModified   DOMBreakpointType = "subtree-modified"
	DOMBreakpointTypeAttributeModified DOMBreakpointType = "attribute-modified"
	DOMBreakpointTypeNodeRemoved       DOMBreakpointType = "node-removed"
)

// String returns the DOMBreakpointType value as a built-in string.
func (t DOMBreakpointType) String() string {
	return string(t)
}

// CSPViolationType data type. CSP Violation type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#type-CSPViolationType
//
// This CDP type is experimental.
type CSPViolationType string

// CSPViolationType valid values.
const (
	CSPViolationTypeTrustedtypeSinkViolation   CSPViolationType = "trustedtype-sink-violation"
	CSPViolationTypeTrustedtypePolicyViolation CSPViolationType = "trustedtype-policy-violation"
)

// String returns the CSPViolationType value as a built-in string.
func (t CSPViolationType) String() string {
	return string(t)
}

// EventListener data type. Object event listener.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger/#type-EventListener
type EventListener struct {
	// `EventListener`'s type.
	Type string `json:"type"`
	// `EventListener`'s useCapture.
	UseCapture bool `json:"useCapture"`
	// `EventListener`'s passive flag.
	Passive bool `json:"passive"`
	// `EventListener`'s once flag.
	Once bool `json:"once"`
	// Script id of the handler code.
	ScriptID runtime.ScriptID `json:"scriptId"`
	// Line number in the script (0-based).
	LineNumber int64 `json:"lineNumber"`
	// Column number in the script (0-based).
	ColumnNumber int64 `json:"columnNumber"`
	// Event handler function value.
	Handler *runtime.RemoteObject `json:"handler,omitempty"`
	// Event original handler function value.
	OriginalHandler *runtime.RemoteObject `json:"originalHandler,omitempty"`
	// Node the listener is added to (if any).
	BackendNodeID int64 `json:"backendNodeId,omitempty"`
}
