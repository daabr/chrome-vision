package console

// Console message.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Console/#type-ConsoleMessage
type ConsoleMessage struct {
	// Message source.
	Source string
	// Message severity.
	Level string
	// Message text.
	Text string
	// URL of the message origin.
	URL string `json:"url,omitempty"`
	// Line number in the resource that generated this message (1-based).
	Line int64 `json:"line,omitempty"`
	// Column number in the resource that generated this message (1-based).
	Column int64 `json:"column,omitempty"`
}
