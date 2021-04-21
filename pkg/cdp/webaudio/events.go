package webaudio

// Notifies that a new BaseAudioContext has been created.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#event-contextCreated
type ContextCreated struct {
	Context BaseAudioContext `json:"context"`
}

// Notifies that an existing BaseAudioContext will be destroyed.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#event-contextWillBeDestroyed
type ContextWillBeDestroyed struct {
	ContextID string `json:"contextId"`
}

// Notifies that existing BaseAudioContext has changed some properties (id stays the same)..
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#event-contextChanged
type ContextChanged struct {
	Context BaseAudioContext `json:"context"`
}

// Notifies that the construction of an AudioListener has finished.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#event-audioListenerCreated
type AudioListenerCreated struct {
	Listener AudioListener `json:"listener"`
}

// Notifies that a new AudioListener has been created.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#event-audioListenerWillBeDestroyed
type AudioListenerWillBeDestroyed struct {
	ContextID  string `json:"contextId"`
	ListenerID string `json:"listenerId"`
}

// Notifies that a new AudioNode has been created.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#event-audioNodeCreated
type AudioNodeCreated struct {
	Node AudioNode `json:"node"`
}

// Notifies that an existing AudioNode has been destroyed.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#event-audioNodeWillBeDestroyed
type AudioNodeWillBeDestroyed struct {
	ContextID string `json:"contextId"`
	NodeID    string `json:"nodeId"`
}

// Notifies that a new AudioParam has been created.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#event-audioParamCreated
type AudioParamCreated struct {
	Param AudioParam `json:"param"`
}

// Notifies that an existing AudioParam has been destroyed.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#event-audioParamWillBeDestroyed
type AudioParamWillBeDestroyed struct {
	ContextID string `json:"contextId"`
	NodeID    string `json:"nodeId"`
	ParamID   string `json:"paramId"`
}

// Notifies that two AudioNodes are connected.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#event-nodesConnected
type NodesConnected struct {
	ContextID             string  `json:"contextId"`
	SourceID              string  `json:"sourceId"`
	DestinationID         string  `json:"destinationId"`
	SourceOutputIndex     float64 `json:"sourceOutputIndex,omitempty"`
	DestinationInputIndex float64 `json:"destinationInputIndex,omitempty"`
}

// Notifies that AudioNodes are disconnected. The destination can be null, and it means all the outgoing connections from the source are disconnected.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#event-nodesDisconnected
type NodesDisconnected struct {
	ContextID             string  `json:"contextId"`
	SourceID              string  `json:"sourceId"`
	DestinationID         string  `json:"destinationId"`
	SourceOutputIndex     float64 `json:"sourceOutputIndex,omitempty"`
	DestinationInputIndex float64 `json:"destinationInputIndex,omitempty"`
}

// Notifies that an AudioNode is connected to an AudioParam.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#event-nodeParamConnected
type NodeParamConnected struct {
	ContextID         string  `json:"contextId"`
	SourceID          string  `json:"sourceId"`
	DestinationID     string  `json:"destinationId"`
	SourceOutputIndex float64 `json:"sourceOutputIndex,omitempty"`
}

// Notifies that an AudioNode is disconnected to an AudioParam.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#event-nodeParamDisconnected
type NodeParamDisconnected struct {
	ContextID         string  `json:"contextId"`
	SourceID          string  `json:"sourceId"`
	DestinationID     string  `json:"destinationId"`
	SourceOutputIndex float64 `json:"sourceOutputIndex,omitempty"`
}
