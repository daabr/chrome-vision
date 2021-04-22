package webaudio

// Notifies that a new BaseAudioContext has been created.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#event-contextCreated
type ContextCreated struct {
	Context BaseAudioContext
}

// Notifies that an existing BaseAudioContext will be destroyed.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#event-contextWillBeDestroyed
type ContextWillBeDestroyed struct {
	ContextID string
}

// Notifies that existing BaseAudioContext has changed some properties (id stays the same)..
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#event-contextChanged
type ContextChanged struct {
	Context BaseAudioContext
}

// Notifies that the construction of an AudioListener has finished.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#event-audioListenerCreated
type AudioListenerCreated struct {
	Listener AudioListener
}

// Notifies that a new AudioListener has been created.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#event-audioListenerWillBeDestroyed
type AudioListenerWillBeDestroyed struct {
	ContextID  string
	ListenerID string
}

// Notifies that a new AudioNode has been created.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#event-audioNodeCreated
type AudioNodeCreated struct {
	Node AudioNode
}

// Notifies that an existing AudioNode has been destroyed.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#event-audioNodeWillBeDestroyed
type AudioNodeWillBeDestroyed struct {
	ContextID string
	NodeID    string
}

// Notifies that a new AudioParam has been created.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#event-audioParamCreated
type AudioParamCreated struct {
	Param AudioParam
}

// Notifies that an existing AudioParam has been destroyed.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#event-audioParamWillBeDestroyed
type AudioParamWillBeDestroyed struct {
	ContextID string
	NodeID    string
	ParamID   string
}

// Notifies that two AudioNodes are connected.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#event-nodesConnected
type NodesConnected struct {
	ContextID             string
	SourceID              string
	DestinationID         string
	SourceOutputIndex     float64 `json:"sourceOutputIndex,omitempty"`
	DestinationInputIndex float64 `json:"destinationInputIndex,omitempty"`
}

// Notifies that AudioNodes are disconnected. The destination can be null, and it means all the outgoing connections from the source are disconnected.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#event-nodesDisconnected
type NodesDisconnected struct {
	ContextID             string
	SourceID              string
	DestinationID         string
	SourceOutputIndex     float64 `json:"sourceOutputIndex,omitempty"`
	DestinationInputIndex float64 `json:"destinationInputIndex,omitempty"`
}

// Notifies that an AudioNode is connected to an AudioParam.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#event-nodeParamConnected
type NodeParamConnected struct {
	ContextID         string
	SourceID          string
	DestinationID     string
	SourceOutputIndex float64 `json:"sourceOutputIndex,omitempty"`
}

// Notifies that an AudioNode is disconnected to an AudioParam.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#event-nodeParamDisconnected
type NodeParamDisconnected struct {
	ContextID         string
	SourceID          string
	DestinationID     string
	SourceOutputIndex float64 `json:"sourceOutputIndex,omitempty"`
}
