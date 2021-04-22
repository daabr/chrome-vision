package heapprofiler

// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#event-addHeapSnapshotChunk
type AddHeapSnapshotChunk struct {
	Chunk string `json:"chunk"`
}

// If heap objects tracking has been started then backend may send update for one or more fragments
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#event-heapStatsUpdate
type HeapStatsUpdate struct {
	// An array of triplets. Each triplet describes a fragment. The first integer is the fragment
	// index, the second integer is a total count of objects for the fragment, the third integer is
	// a total size of the objects for the fragment.
	StatsUpdate []int64 `json:"statsUpdate"`
}

// If heap objects tracking has been started then backend regularly sends a current value for last
// seen object id and corresponding timestamp. If the were changes in the heap since last event
// then one or more heapStatsUpdate events will be sent before a new lastSeenObjectId event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#event-lastSeenObjectId
type LastSeenObjectID struct {
	LastSeenObjectID int64   `json:"lastSeenObjectId"`
	Timestamp        float64 `json:"timestamp"`
}

// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#event-reportHeapSnapshotProgress
type ReportHeapSnapshotProgress struct {
	Done     int64 `json:"done"`
	Total    int64 `json:"total"`
	Finished bool  `json:"finished,omitempty"`
}

// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#event-resetProfiles
type ResetProfiles struct{}
