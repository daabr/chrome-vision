package log

import "github.com/daabr/chrome-vision/pkg/cdp/runtime"

// Log entry.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Log/#type-LogEntry
type LogEntry struct {
	// Log entry source.
	Source string
	// Log entry severity.
	Level string
	// Logged text.
	Text string
	// Timestamp when this entry was added.
	Timestamp runtime.Timestamp
	// URL of the resource if known.
	URL string `json:"url,omitempty"`
	// Line number in the resource.
	LineNumber int64 `json:"lineNumber,omitempty"`
	// JavaScript stack trace.
	StackTrace *runtime.StackTrace `json:"stackTrace,omitempty"`
	// Identifier of the network request associated with this entry.
	NetworkRequestID string `json:"networkRequestId,omitempty"`
	// Identifier of the worker associated with this entry.
	WorkerID string `json:"workerId,omitempty"`
	// Call arguments.
	Args []runtime.RemoteObject `json:"args,omitempty"`
}

// Violation configuration setting.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Log/#type-ViolationSetting
type ViolationSetting struct {
	// Violation type.
	Name string
	// Time threshold to trigger upon.
	Threshold float64
}
