package heapprofiler

import "github.com/daabr/chrome-vision/pkg/cdp/runtime"

// HeapSnapshotObjectID data type. Heap snapshot object id.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#type-HeapSnapshotObjectId
type HeapSnapshotObjectID string

// SamplingHeapProfileNode data type. Sampling Heap Profile node. Holds callsite information, allocation statistics and child nodes.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#type-SamplingHeapProfileNode
type SamplingHeapProfileNode struct {
	// Function location.
	CallFrame runtime.CallFrame `json:"callFrame"`
	// Allocations size in bytes for the node excluding children.
	SelfSize float64 `json:"selfSize"`
	// Node id. Ids are unique across all profiles collected between startSampling and stopSampling.
	ID int64 `json:"id"`
	// Child nodes.
	Children []SamplingHeapProfileNode `json:"children"`
}

// SamplingHeapProfileSample data type. A single sample from a sampling profile.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#type-SamplingHeapProfileSample
type SamplingHeapProfileSample struct {
	// Allocation size in bytes attributed to the sample.
	Size float64 `json:"size"`
	// Id of the corresponding profile tree node.
	NodeID int64 `json:"nodeId"`
	// Time-ordered sample ordinal number. It is unique across all profiles retrieved
	// between startSampling and stopSampling.
	Ordinal float64 `json:"ordinal"`
}

// SamplingHeapProfile data type. Sampling profile.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#type-SamplingHeapProfile
type SamplingHeapProfile struct {
	Head    SamplingHeapProfileNode     `json:"head"`
	Samples []SamplingHeapProfileSample `json:"samples"`
}
