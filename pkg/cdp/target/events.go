package target

// Issued when attached to target because of auto-attach or `attachToTarget` command.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-attachedToTarget
//
// This CDP event is experimental.
type AttachedToTarget struct {
	// Identifier assigned to the session used to send/receive messages.
	SessionID          string
	TargetInfo         TargetInfo
	WaitingForDebugger bool
}

// Issued when detached from target for any reason (including `detachFromTarget` command). Can be
// issued multiple times per target if multiple sessions have been attached to it.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-detachedFromTarget
//
// This CDP event is experimental.
type DetachedFromTarget struct {
	// Detached session identifier.
	SessionID string
	// Deprecated.
	//
	// This CDP parameter is deprecated.
	TargetID string `json:"targetId,omitempty"`
}

// Notifies about a new protocol message received from the session (as reported in
// `attachedToTarget` event).
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-receivedMessageFromTarget
type ReceivedMessageFromTarget struct {
	// Identifier of a session which sends a message.
	SessionID string
	Message   string
	// Deprecated.
	//
	// This CDP parameter is deprecated.
	TargetID string `json:"targetId,omitempty"`
}

// Issued when a possible inspection target is created.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-targetCreated
type TargetCreated struct {
	TargetInfo TargetInfo
}

// Issued when a target is destroyed.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-targetDestroyed
type TargetDestroyed struct {
	TargetID string
}

// Issued when a target has crashed.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-targetCrashed
type TargetCrashed struct {
	TargetID string
	// Termination status type.
	Status string
	// Termination error code.
	ErrorCode int64
}

// Issued when some information about a target has changed. This only happens between
// `targetCreated` and `targetDestroyed`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-targetInfoChanged
type TargetInfoChanged struct {
	TargetInfo TargetInfo
}
