package tracing

// MemoryDumpConfig data type. Configuration for memory dump. Used only when "memory-infra" category is enabled.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#type-MemoryDumpConfig
type MemoryDumpConfig struct{}

// TraceConfig data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#type-TraceConfig
type TraceConfig struct {
	// Controls how the trace buffer stores data.
	RecordMode string `json:"recordMode,omitempty"`
	// Turns on JavaScript stack sampling.
	EnableSampling bool `json:"enableSampling,omitempty"`
	// Turns on system tracing.
	EnableSystrace bool `json:"enableSystrace,omitempty"`
	// Turns on argument filter.
	EnableArgumentFilter bool `json:"enableArgumentFilter,omitempty"`
	// Included category filters.
	IncludedCategories []string `json:"includedCategories,omitempty"`
	// Excluded category filters.
	ExcludedCategories []string `json:"excludedCategories,omitempty"`
	// Configuration to synthesize the delays in tracing.
	SyntheticDelays []string `json:"syntheticDelays,omitempty"`
	// Configuration for memory dump triggers. Used only when "memory-infra" category is enabled.
	MemoryDumpConfig *MemoryDumpConfig `json:"memoryDumpConfig,omitempty"`
}

// StreamFormat data type. Data format of a trace. Can be either the legacy JSON format or the
// protocol buffer format. Note that the JSON format will be deprecated soon.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#type-StreamFormat
type StreamFormat string

// StreamFormat valid values.
const (
	StreamFormatJSON  StreamFormat = "json"
	StreamFormatProto StreamFormat = "proto"
)

// String returns the StreamFormat value as a built-in string.
func (t StreamFormat) String() string {
	return string(t)
}

// StreamCompression data type. Compression type to use for traces returned via streams.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#type-StreamCompression
type StreamCompression string

// StreamCompression valid values.
const (
	StreamCompressionNone StreamCompression = "none"
	StreamCompressionGzip StreamCompression = "gzip"
)

// String returns the StreamCompression value as a built-in string.
func (t StreamCompression) String() string {
	return string(t)
}

// MemoryDumpLevelOfDetail data type. Details exposed when memory request explicitly declared.
// Keep consistent with memory_dump_request_args.h and
// memory_instrumentation.mojom
//
// https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#type-MemoryDumpLevelOfDetail
type MemoryDumpLevelOfDetail string

// MemoryDumpLevelOfDetail valid values.
const (
	MemoryDumpLevelOfDetailBackground MemoryDumpLevelOfDetail = "background"
	MemoryDumpLevelOfDetailLight      MemoryDumpLevelOfDetail = "light"
	MemoryDumpLevelOfDetailDetailed   MemoryDumpLevelOfDetail = "detailed"
)

// String returns the MemoryDumpLevelOfDetail value as a built-in string.
func (t MemoryDumpLevelOfDetail) String() string {
	return string(t)
}

// Backend data type. Backend type to use for tracing. `chrome` uses the Chrome-integrated
// tracing service and is supported on all platforms. `system` is only
// supported on Chrome OS and uses the Perfetto system tracing service.
// `auto` chooses `system` when the perfettoConfig provided to Tracing.start
// specifies at least one non-Chrome data source; otherwise uses `chrome`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#type-TracingBackend
type Backend string

// Backend valid values.
const (
	BackendAuto   Backend = "auto"
	BackendChrome Backend = "chrome"
	BackendSystem Backend = "system"
)

// String returns the Backend value as a built-in string.
func (t Backend) String() string {
	return string(t)
}
