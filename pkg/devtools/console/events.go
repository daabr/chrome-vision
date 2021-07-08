package console

// MessageAdded asynchronous event. Issued when new console message is added.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Console/#event-messageAdded
type MessageAdded struct {
	// Console message that has been added.
	Message Message `json:"message"`
}
