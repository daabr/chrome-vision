package media

// Players will get an ID that is unique within the agent context.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Media/#type-PlayerId
type PlayerID string

// https://chromedevtools.github.io/devtools-protocol/tot/Media/#type-Timestamp
type Timestamp float64

// Have one type per entry in MediaLogRecord::Type
// Corresponds to kMessage
//
// https://chromedevtools.github.io/devtools-protocol/tot/Media/#type-PlayerMessage
type PlayerMessage struct {
	// Keep in sync with MediaLogMessageLevel
	// We are currently keeping the message level 'error' separate from the
	// PlayerError type because right now they represent different things,
	// this one being a DVLOG(ERROR) style log message that gets printed
	// based on what log level is selected in the UI, and the other is a
	// representation of a media::PipelineStatus object. Soon however we're
	// going to be moving away from using PipelineStatus for errors and
	// introducing a new error type which should hopefully let us integrate
	// the error log level into the PlayerError type.
	Level   string `json:"level"`
	Message string `json:"message"`
}

// Corresponds to kMediaPropertyChange
//
// https://chromedevtools.github.io/devtools-protocol/tot/Media/#type-PlayerProperty
type PlayerProperty struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Corresponds to kMediaEventTriggered
//
// https://chromedevtools.github.io/devtools-protocol/tot/Media/#type-PlayerEvent
type PlayerEvent struct {
	Timestamp float64 `json:"timestamp"`
	Value     string  `json:"value"`
}

// Corresponds to kMediaError
//
// https://chromedevtools.github.io/devtools-protocol/tot/Media/#type-PlayerError
type PlayerError struct {
	Type string `json:"type"`
	// When this switches to using media::Status instead of PipelineStatus
	// we can remove "errorCode" and replace it with the fields from
	// a Status instance. This also seems like a duplicate of the error
	// level enum - there is a todo bug to have that level removed and
	// use this instead. (crbug.com/1068454)
	ErrorCode string `json:"errorCode"`
}
