package memory

// Memory pressure level.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Memory/#type-PressureLevel
type PressureLevel string

// PressureLevel valid values.
const (
	PressureLevelModerate PressureLevel = "moderate"
	PressureLevelCritical PressureLevel = "critical"
)

// Heap profile sample.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Memory/#type-SamplingProfileNode
type SamplingProfileNode struct {
	// Size of the sampled allocation.
	Size float64
	// Total bytes attributed to this sample.
	Total float64
	// Execution stack at the point of allocation.
	Stack []string
}

// Array of heap profile samples.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Memory/#type-SamplingProfile
type SamplingProfile struct {
	Samples []SamplingProfileNode
	Modules []Module
}

// Executable module information
//
// https://chromedevtools.github.io/devtools-protocol/tot/Memory/#type-Module
type Module struct {
	// Name of the module.
	Name string
	// UUID of the module.
	Uuid string
	// Base address where the module is loaded into memory. Encoded as a decimal
	// or hexadecimal (0x prefixed) string.
	BaseAddress string
	// Size of the module in bytes.
	Size float64
}
