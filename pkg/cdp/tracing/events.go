package tracing

import "encoding/json"

// BufferUsage asynchronous event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#event-bufferUsage
type BufferUsage struct {
	// A number in range [0..1] that indicates the used size of event buffer as a fraction of its
	// total size.
	PercentFull float64 `json:"percentFull,omitempty"`
	// An approximate number of events in the trace log.
	EventCount float64 `json:"eventCount,omitempty"`
	// A number in range [0..1] that indicates the used size of event buffer as a fraction of its
	// total size.
	Value float64 `json:"value,omitempty"`
}

// DataCollected asynchronous event.
//
// Contains an bucket of collected trace events. When tracing is stopped collected events will be
// send as a sequence of dataCollected events followed by tracingComplete event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#event-dataCollected
type DataCollected struct {
	Value []json.RawMessage `json:"value"`
}

// TracingComplete asynchronous event.
//
// Signals that tracing is stopped and there is no trace buffers pending flush, all data were
// delivered via dataCollected events.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#event-tracingComplete
type TracingComplete struct {
	// Indicates whether some trace data is known to have been lost, e.g. because the trace ring
	// buffer wrapped around.
	DataLossOccurred bool `json:"dataLossOccurred"`
	// A handle of the stream that holds resulting trace data.
	Stream string `json:"stream,omitempty"`
	// Trace data format of returned stream.
	TraceFormat string `json:"traceFormat,omitempty"`
	// Compression format of returned stream.
	StreamCompression string `json:"streamCompression,omitempty"`
}
