package profiler

import "github.com/daabr/chrome-vision/pkg/cdp/debugger"

// ConsoleProfileFinished asynchronous event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#event-consoleProfileFinished
type ConsoleProfileFinished struct {
	ID string `json:"id"`
	// Location of console.profileEnd().
	Location debugger.Location `json:"location"`
	Profile  Profile           `json:"profile"`
	// Profile title passed as an argument to console.profile().
	Title string `json:"title,omitempty"`
}

// ConsoleProfileStarted asynchronous event.
//
// Sent when new profile recording is started using console.profile() call.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#event-consoleProfileStarted
type ConsoleProfileStarted struct {
	ID string `json:"id"`
	// Location of console.profile().
	Location debugger.Location `json:"location"`
	// Profile title passed as an argument to console.profile().
	Title string `json:"title,omitempty"`
}

// PreciseCoverageDeltaUpdate asynchronous event.
//
// Reports coverage delta since the last poll (either from an event like this, or from
// `takePreciseCoverage` for the current isolate. May only be sent if precise code
// coverage has been started. This event can be trigged by the embedder to, for example,
// trigger collection of coverage data immediatelly at a certain point in time.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#event-preciseCoverageDeltaUpdate
//
// This CDP event is experimental.
type PreciseCoverageDeltaUpdate struct {
	// Monotonically increasing time (in seconds) when the coverage update was taken in the backend.
	Timestamp float64 `json:"timestamp"`
	// Identifier for distinguishing coverage events.
	Occassion string `json:"occassion"`
	// Coverage data for the current isolate.
	Result []ScriptCoverage `json:"result"`
}
