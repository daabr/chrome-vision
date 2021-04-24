package log

// EntryAdded asynchronous event.
//
// Issued when new message was logged.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Log/#event-entryAdded
type EntryAdded struct {
	// The entry.
	Entry LogEntry `json:"entry"`
}
