package backgroundservice

// RecordingStateChanged asynchronous event. Called when the recording state for the service has been updated.
//
// https://chromedevtools.github.io/devtools-protocol/tot/BackgroundService/#event-recordingStateChanged
type RecordingStateChanged struct {
	IsRecording bool        `json:"isRecording"`
	Service     ServiceName `json:"service"`
}

// BackgroundServiceEventReceived asynchronous event. Called with all existing backgroundServiceEvents when enabled, and all new
// events afterwards if enabled and recording.
//
// https://chromedevtools.github.io/devtools-protocol/tot/BackgroundService/#event-backgroundServiceEventReceived
type EventReceived struct {
	BackgroundServiceEvent Event `json:"backgroundServiceEvent"`
}
