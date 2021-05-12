package log

import "github.com/daabr/chrome-vision/pkg/devtools/runtime"

// LogEntry data type. Log entry.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Log/#type-LogEntry
type LogEntry struct {
	// Log entry source.
	Source string `json:"source"`
	// Log entry severity.
	Level string `json:"level"`
	// Logged text.
	Text string `json:"text"`
	// Timestamp when this entry was added.
	Timestamp runtime.Timestamp `json:"timestamp"`
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

// ViolationSetting data type. Violation configuration setting.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Log/#type-ViolationSetting
type ViolationSetting struct {
	// Violation type.
	Name string `json:"name"`
	// Time threshold to trigger upon.
	Threshold float64 `json:"threshold"`
}
