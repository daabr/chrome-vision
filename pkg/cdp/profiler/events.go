package profiler

import "github.com/daabr/chrome-vision/pkg/cdp/debugger"

// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#event-consoleProfileFinished
type ConsoleProfileFinished struct {
	ID string
	// Location of console.profileEnd().
	Location debugger.Location
	Profile  Profile
	// Profile title passed as an argument to console.profile().
	Title string `json:"title,omitempty"`
}

// Sent when new profile recording is started using console.profile() call.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#event-consoleProfileStarted
type ConsoleProfileStarted struct {
	ID string
	// Location of console.profile().
	Location debugger.Location
	// Profile title passed as an argument to console.profile().
	Title string `json:"title,omitempty"`
}

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
	Timestamp float64
	// Identifier for distinguishing coverage events.
	Occassion string
	// Coverage data for the current isolate.
	Result []ScriptCoverage
}
