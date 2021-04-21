package backgroundservice

// Called when the recording state for the service has been updated.
//
// https://chromedevtools.github.io/devtools-protocol/tot/BackgroundService/#event-recordingStateChanged
type RecordingStateChanged struct {
	IsRecording bool   `json:"isRecording"`
	Service     string `json:"service"`
}

// Called with all existing backgroundServiceEvents when enabled, and all new
// events afterwards if enabled and recording.
//
// https://chromedevtools.github.io/devtools-protocol/tot/BackgroundService/#event-backgroundServiceEventReceived
type BackgroundServiceEventReceived struct {
	BackgroundServiceEvent BackgroundServiceEvent `json:"backgroundServiceEvent"`
}
