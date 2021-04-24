package cast

// Sink data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Cast/#type-Sink
type Sink struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	// Text describing the current session. Present only if there is an active
	// session on the sink.
	Session string `json:"session,omitempty"`
}
