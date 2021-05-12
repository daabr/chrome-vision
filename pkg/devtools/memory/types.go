package memory

// PressureLevel data type. Memory pressure level.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Memory/#type-PressureLevel
type PressureLevel string

// PressureLevel valid values.
const (
	PressureLevelModerate PressureLevel = "moderate"
	PressureLevelCritical PressureLevel = "critical"
)

// SamplingProfileNode data type. Heap profile sample.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Memory/#type-SamplingProfileNode
type SamplingProfileNode struct {
	// Size of the sampled allocation.
	Size float64 `json:"size"`
	// Total bytes attributed to this sample.
	Total float64 `json:"total"`
	// Execution stack at the point of allocation.
	Stack []string `json:"stack"`
}

// SamplingProfile data type. Array of heap profile samples.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Memory/#type-SamplingProfile
type SamplingProfile struct {
	Samples []SamplingProfileNode `json:"samples"`
	Modules []Module              `json:"modules"`
}

// Module data type. Executable module information
//
// https://chromedevtools.github.io/devtools-protocol/tot/Memory/#type-Module
type Module struct {
	// Name of the module.
	Name string `json:"name"`
	// UUID of the module.
	Uuid string `json:"uuid"`
	// Base address where the module is loaded into memory. Encoded as a decimal
	// or hexadecimal (0x prefixed) string.
	BaseAddress string `json:"baseAddress"`
	// Size of the module in bytes.
	Size float64 `json:"size"`
}
