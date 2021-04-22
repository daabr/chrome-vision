package console

// Console message.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Console/#type-ConsoleMessage
type ConsoleMessage struct {
	// Message source.
	Source string `json:"source"`
	// Message severity.
	Level string `json:"level"`
	// Message text.
	Text string `json:"text"`
	// URL of the message origin.
	URL string `json:"url,omitempty"`
	// Line number in the resource that generated this message (1-based).
	Line int64 `json:"line,omitempty"`
	// Column number in the resource that generated this message (1-based).
	Column int64 `json:"column,omitempty"`
}
