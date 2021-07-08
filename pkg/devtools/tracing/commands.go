package tracing

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
)

// End contains the parameters, and acts as
// a Go receiver, for the CDP command `end`.
//
// Stop trace events collection.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#method-end
type End struct{}

// NewEnd constructs a new End struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#method-end
func NewEnd() *End {
	return &End{}
}

// Do sends the End CDP command to a browser,
// and returns the browser's response.
func (t *End) Do(ctx context.Context) error {
	response, err := devtools.Send(ctx, "Tracing.end", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// GetCategories contains the parameters, and acts as
// a Go receiver, for the CDP command `getCategories`.
//
// Gets supported tracing categories.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#method-getCategories
type GetCategories struct{}

// NewGetCategories constructs a new GetCategories struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#method-getCategories
func NewGetCategories() *GetCategories {
	return &GetCategories{}
}

// GetCategoriesResult contains the browser's response
// to calling the GetCategories CDP command with Do().
type GetCategoriesResult struct {
	// A list of supported tracing categories.
	Categories []string `json:"categories"`
}

// Do sends the GetCategories CDP command to a browser,
// and returns the browser's response.
func (t *GetCategories) Do(ctx context.Context) (*GetCategoriesResult, error) {
	response, err := devtools.Send(ctx, "Tracing.getCategories", nil)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetCategoriesResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// RecordClockSyncMarker contains the parameters, and acts as
// a Go receiver, for the CDP command `recordClockSyncMarker`.
//
// Record a clock sync marker in the trace.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#method-recordClockSyncMarker
type RecordClockSyncMarker struct {
	// The ID of this clock sync marker
	SyncID string `json:"syncId"`
}

// NewRecordClockSyncMarker constructs a new RecordClockSyncMarker struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#method-recordClockSyncMarker
func NewRecordClockSyncMarker(syncID string) *RecordClockSyncMarker {
	return &RecordClockSyncMarker{
		SyncID: syncID,
	}
}

// Do sends the RecordClockSyncMarker CDP command to a browser,
// and returns the browser's response.
func (t *RecordClockSyncMarker) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := devtools.Send(ctx, "Tracing.recordClockSyncMarker", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// RequestMemoryDump contains the parameters, and acts as
// a Go receiver, for the CDP command `requestMemoryDump`.
//
// Request a global memory dump.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#method-requestMemoryDump
type RequestMemoryDump struct {
	// Enables more deterministic results by forcing garbage collection
	Deterministic bool `json:"deterministic,omitempty"`
	// Specifies level of details in memory dump. Defaults to "detailed".
	LevelOfDetail *MemoryDumpLevelOfDetail `json:"levelOfDetail,omitempty"`
}

// NewRequestMemoryDump constructs a new RequestMemoryDump struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#method-requestMemoryDump
func NewRequestMemoryDump() *RequestMemoryDump {
	return &RequestMemoryDump{}
}

// SetDeterministic adds or modifies the value of the optional
// parameter `deterministic` in the RequestMemoryDump CDP command.
//
// Enables more deterministic results by forcing garbage collection
func (t *RequestMemoryDump) SetDeterministic(v bool) *RequestMemoryDump {
	t.Deterministic = v
	return t
}

// SetLevelOfDetail adds or modifies the value of the optional
// parameter `levelOfDetail` in the RequestMemoryDump CDP command.
//
// Specifies level of details in memory dump. Defaults to "detailed".
func (t *RequestMemoryDump) SetLevelOfDetail(v MemoryDumpLevelOfDetail) *RequestMemoryDump {
	t.LevelOfDetail = &v
	return t
}

// RequestMemoryDumpResult contains the browser's response
// to calling the RequestMemoryDump CDP command with Do().
type RequestMemoryDumpResult struct {
	// GUID of the resulting global memory dump.
	DumpGUID string `json:"dumpGuid"`
	// True iff the global memory dump succeeded.
	Success bool `json:"success"`
}

// Do sends the RequestMemoryDump CDP command to a browser,
// and returns the browser's response.
func (t *RequestMemoryDump) Do(ctx context.Context) (*RequestMemoryDumpResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.Send(ctx, "Tracing.requestMemoryDump", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &RequestMemoryDumpResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Start contains the parameters, and acts as
// a Go receiver, for the CDP command `start`.
//
// Start trace events collection.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#method-start
type Start struct {
	// Category/tag filter
	//
	// This CDP parameter is deprecated.
	Categories string `json:"categories,omitempty"`
	// Tracing options
	//
	// This CDP parameter is deprecated.
	Options string `json:"options,omitempty"`
	// If set, the agent will issue bufferUsage events at this interval, specified in milliseconds
	BufferUsageReportingInterval float64 `json:"bufferUsageReportingInterval,omitempty"`
	// Whether to report trace events as series of dataCollected events or to save trace to a
	// stream (defaults to `ReportEvents`).
	TransferMode string `json:"transferMode,omitempty"`
	// Trace data format to use. This only applies when using `ReturnAsStream`
	// transfer mode (defaults to `json`).
	StreamFormat *StreamFormat `json:"streamFormat,omitempty"`
	// Compression format to use. This only applies when using `ReturnAsStream`
	// transfer mode (defaults to `none`)
	StreamCompression *StreamCompression `json:"streamCompression,omitempty"`
	TraceConfig       *TraceConfig       `json:"traceConfig,omitempty"`
	// Base64-encoded serialized perfetto.protos.TraceConfig protobuf message
	// When specified, the parameters `categories`, `options`, `traceConfig`
	// are ignored. (Encoded as a base64 string when passed over JSON)
	PerfettoConfig string `json:"perfettoConfig,omitempty"`
	// Backend type (defaults to `auto`)
	TracingBackend *Backend `json:"tracingBackend,omitempty"`
}

// NewStart constructs a new Start struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#method-start
func NewStart() *Start {
	return &Start{}
}

// SetCategories adds or modifies the value of the optional
// parameter `categories` in the Start CDP command.
//
// Category/tag filter
//
// This CDP parameter is deprecated.
func (t *Start) SetCategories(v string) *Start {
	t.Categories = v
	return t
}

// SetOptions adds or modifies the value of the optional
// parameter `options` in the Start CDP command.
//
// Tracing options
//
// This CDP parameter is deprecated.
func (t *Start) SetOptions(v string) *Start {
	t.Options = v
	return t
}

// SetBufferUsageReportingInterval adds or modifies the value of the optional
// parameter `bufferUsageReportingInterval` in the Start CDP command.
//
// If set, the agent will issue bufferUsage events at this interval, specified in milliseconds
func (t *Start) SetBufferUsageReportingInterval(v float64) *Start {
	t.BufferUsageReportingInterval = v
	return t
}

// SetTransferMode adds or modifies the value of the optional
// parameter `transferMode` in the Start CDP command.
//
// Whether to report trace events as series of dataCollected events or to save trace to a
// stream (defaults to `ReportEvents`).
func (t *Start) SetTransferMode(v string) *Start {
	t.TransferMode = v
	return t
}

// SetStreamFormat adds or modifies the value of the optional
// parameter `streamFormat` in the Start CDP command.
//
// Trace data format to use. This only applies when using `ReturnAsStream`
// transfer mode (defaults to `json`).
func (t *Start) SetStreamFormat(v StreamFormat) *Start {
	t.StreamFormat = &v
	return t
}

// SetStreamCompression adds or modifies the value of the optional
// parameter `streamCompression` in the Start CDP command.
//
// Compression format to use. This only applies when using `ReturnAsStream`
// transfer mode (defaults to `none`)
func (t *Start) SetStreamCompression(v StreamCompression) *Start {
	t.StreamCompression = &v
	return t
}

// SetTraceConfig adds or modifies the value of the optional
// parameter `traceConfig` in the Start CDP command.
func (t *Start) SetTraceConfig(v TraceConfig) *Start {
	t.TraceConfig = &v
	return t
}

// SetPerfettoConfig adds or modifies the value of the optional
// parameter `perfettoConfig` in the Start CDP command.
//
// Base64-encoded serialized perfetto.protos.TraceConfig protobuf message
// When specified, the parameters `categories`, `options`, `traceConfig`
// are ignored. (Encoded as a base64 string when passed over JSON)
func (t *Start) SetPerfettoConfig(v string) *Start {
	t.PerfettoConfig = v
	return t
}

// SetTracingBackend adds or modifies the value of the optional
// parameter `tracingBackend` in the Start CDP command.
//
// Backend type (defaults to `auto`)
func (t *Start) SetTracingBackend(v Backend) *Start {
	t.TracingBackend = &v
	return t
}

// Do sends the Start CDP command to a browser,
// and returns the browser's response.
func (t *Start) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := devtools.Send(ctx, "Tracing.start", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}
