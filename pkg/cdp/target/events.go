package target

// AttachedToTarget asynchronous event.
//
// Issued when attached to target because of auto-attach or `attachToTarget` command.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-attachedToTarget
//
// This CDP event is experimental.
type AttachedToTarget struct {
	// Identifier assigned to the session used to send/receive messages.
	SessionID          string     `json:"sessionId"`
	TargetInfo         TargetInfo `json:"targetInfo"`
	WaitingForDebugger bool       `json:"waitingForDebugger"`
}

// DetachedFromTarget asynchronous event.
//
// Issued when detached from target for any reason (including `detachFromTarget` command). Can be
// issued multiple times per target if multiple sessions have been attached to it.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-detachedFromTarget
//
// This CDP event is experimental.
type DetachedFromTarget struct {
	// Detached session identifier.
	SessionID string `json:"sessionId"`
	// Deprecated.
	//
	// This CDP parameter is deprecated.
	TargetID string `json:"targetId,omitempty"`
}

// ReceivedMessageFromTarget asynchronous event.
//
// Notifies about a new protocol message received from the session (as reported in
// `attachedToTarget` event).
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-receivedMessageFromTarget
type ReceivedMessageFromTarget struct {
	// Identifier of a session which sends a message.
	SessionID string `json:"sessionId"`
	Message   string `json:"message"`
	// Deprecated.
	//
	// This CDP parameter is deprecated.
	TargetID string `json:"targetId,omitempty"`
}

// TargetCreated asynchronous event.
//
// Issued when a possible inspection target is created.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-targetCreated
type TargetCreated struct {
	TargetInfo TargetInfo `json:"targetInfo"`
}

// TargetDestroyed asynchronous event.
//
// Issued when a target is destroyed.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-targetDestroyed
type TargetDestroyed struct {
	TargetID string `json:"targetId"`
}

// TargetCrashed asynchronous event.
//
// Issued when a target has crashed.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-targetCrashed
type TargetCrashed struct {
	TargetID string `json:"targetId"`
	// Termination status type.
	Status string `json:"status"`
	// Termination error code.
	ErrorCode int64 `json:"errorCode"`
}

// TargetInfoChanged asynchronous event.
//
// Issued when some information about a target has changed. This only happens between
// `targetCreated` and `targetDestroyed`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-targetInfoChanged
type TargetInfoChanged struct {
	TargetInfo TargetInfo `json:"targetInfo"`
}
