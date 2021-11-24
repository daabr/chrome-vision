package security

// CertificateError asynchronous event. There is a certificate error. If overriding certificate errors is enabled, then it should be
// handled with the `handleCertificateError` command. Note: this event does not fire if the
// certificate error has been allowed internally. Only one client per target should override
// certificate errors at the same time.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Security/#event-certificateError
//
// This CDP event is deprecated.
type CertificateError struct {
	// The ID of the event.
	EventID int64 `json:"eventId"`
	// The type of the error.
	ErrorType string `json:"errorType"`
	// The url that was requested.
	RequestURL string `json:"requestURL"`
}

// VisibleSecurityStateChanged asynchronous event. The security state of the page changed.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Security/#event-visibleSecurityStateChanged
//
// This CDP event is experimental.
type VisibleSecurityStateChanged struct {
	// Security state information about the page.
	VisibleSecurityState VisibleSecurityState `json:"visibleSecurityState"`
}

// StateChanged asynchronous event. The security state of the page changed. No longer being sent.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Security/#event-securityStateChanged
//
// This CDP event is deprecated.
type StateChanged struct {
	// Security state.
	SecurityState State `json:"securityState"`
	// True if the page was loaded over cryptographic transport such as HTTPS.
	//
	// This CDP parameter is deprecated.
	SchemeIsCryptographic bool `json:"schemeIsCryptographic"`
	// Previously a list of explanations for the security state. Now always
	// empty.
	//
	// This CDP parameter is deprecated.
	Explanations []StateExplanation `json:"explanations"`
	// Information about insecure content on the page.
	//
	// This CDP parameter is deprecated.
	InsecureContentStatus InsecureContentStatus `json:"insecureContentStatus"`
	// Overrides user-visible description of the state. Always omitted.
	//
	// This CDP parameter is deprecated.
	Summary string `json:"summary,omitempty"`
}
