package tethering

// Accepted asynchronous event. Informs that port was successfully bound and got a specified connection id.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Tethering/#event-accepted
type Accepted struct {
	// Port number that was successfully bound.
	Port int64 `json:"port"`
	// Connection id to be used.
	ConnectionID string `json:"connectionId"`
}
