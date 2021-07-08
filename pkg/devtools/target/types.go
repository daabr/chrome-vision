package target

// TargetID data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#type-TargetID
type ID string

// SessionID data type. Unique identifier of attached debugging session.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#type-SessionID
type SessionID string

// TargetInfo data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#type-TargetInfo
type Info struct {
	TargetID string `json:"targetId"`
	Type     string `json:"type"`
	Title    string `json:"title"`
	URL      string `json:"url"`
	// Whether the target has an attached client.
	Attached bool `json:"attached"`
	// Opener target Id
	OpenerID string `json:"openerId,omitempty"`
	// Whether the target has access to the originating window.
	//
	// This CDP property is experimental.
	CanAccessOpener bool `json:"canAccessOpener"`
	// Frame id of originating window (is only set if target has an opener).
	//
	// This CDP property is experimental.
	OpenerFrameID string `json:"openerFrameId,omitempty"`
	// This CDP property is experimental.
	BrowserContextID string `json:"browserContextId,omitempty"`
}

// RemoteLocation data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#type-RemoteLocation
//
// This CDP type is experimental.
type RemoteLocation struct {
	Host string `json:"host"`
	Port int64  `json:"port"`
}
