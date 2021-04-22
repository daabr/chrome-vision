package cast

// https://chromedevtools.github.io/devtools-protocol/tot/Cast/#type-Sink
type Sink struct {
	Name string
	ID   string
	// Text describing the current session. Present only if there is an active
	// session on the sink.
	Session string `json:"session,omitempty"`
}
