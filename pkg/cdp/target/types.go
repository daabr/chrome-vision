package target

// https://chromedevtools.github.io/devtools-protocol/tot/Target/#type-TargetID
type TargetID string

// Unique identifier of attached debugging session.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#type-SessionID
type SessionID string

// https://chromedevtools.github.io/devtools-protocol/tot/Target/#type-TargetInfo
type TargetInfo struct {
	TargetID string
	Type     string
	Title    string
	URL      string
	// Whether the target has an attached client.
	Attached bool
	// Opener target Id
	OpenerID string `json:"openerId,omitempty"`
	// Whether the target has access to the originating window.
	//
	// This CDP property is experimental.
	CanAccessOpener bool
	// Frame id of originating window (is only set if target has an opener).
	//
	// This CDP property is experimental.
	OpenerFrameID string `json:"openerFrameId,omitempty"`
	// This CDP property is experimental.
	BrowserContextID string `json:"browserContextId,omitempty"`
}

// https://chromedevtools.github.io/devtools-protocol/tot/Target/#type-RemoteLocation
//
// This CDP type is experimental.
type RemoteLocation struct {
	Host string
	Port int64
}
