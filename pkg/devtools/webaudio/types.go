package webaudio

// GraphObjectID data type. An unique ID for a graph object (AudioContext, AudioNode, AudioParam) in Web Audio API
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#type-GraphObjectId
type GraphObjectID string

// ContextType data type. Enum of BaseAudioContext types
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#type-ContextType
type ContextType string

// ContextType valid values.
const (
	ContextTypeRealtime ContextType = "realtime"
	ContextTypeOffline  ContextType = "offline"
)

// String returns the ContextType value as a built-in string.
func (t ContextType) String() string {
	return string(t)
}

// ContextState data type. Enum of AudioContextState from the spec
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#type-ContextState
type ContextState string

// ContextState valid values.
const (
	ContextStateSuspended ContextState = "suspended"
	ContextStateRunning   ContextState = "running"
	ContextStateClosed    ContextState = "closed"
)

// String returns the ContextState value as a built-in string.
func (t ContextState) String() string {
	return string(t)
}

// NodeType data type. Enum of AudioNode types
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#type-NodeType
type NodeType string

// ChannelCountMode data type. Enum of AudioNode::ChannelCountMode from the spec
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#type-ChannelCountMode
type ChannelCountMode string

// ChannelCountMode valid values.
const (
	ChannelCountModeClampedMax ChannelCountMode = "clamped-max"
	ChannelCountModeExplicit   ChannelCountMode = "explicit"
	ChannelCountModeMax        ChannelCountMode = "max"
)

// String returns the ChannelCountMode value as a built-in string.
func (t ChannelCountMode) String() string {
	return string(t)
}

// ChannelInterpretation data type. Enum of AudioNode::ChannelInterpretation from the spec
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#type-ChannelInterpretation
type ChannelInterpretation string

// ChannelInterpretation valid values.
const (
	ChannelInterpretationDiscrete ChannelInterpretation = "discrete"
	ChannelInterpretationSpeakers ChannelInterpretation = "speakers"
)

// String returns the ChannelInterpretation value as a built-in string.
func (t ChannelInterpretation) String() string {
	return string(t)
}

// ParamType data type. Enum of AudioParam types
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#type-ParamType
type ParamType string

// AutomationRate data type. Enum of AudioParam::AutomationRate from the spec
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#type-AutomationRate
type AutomationRate string

// AutomationRate valid values.
const (
	AutomationRateARate AutomationRate = "a-rate"
	AutomationRateKRate AutomationRate = "k-rate"
)

// String returns the AutomationRate value as a built-in string.
func (t AutomationRate) String() string {
	return string(t)
}

// ContextRealtimeData data type. Fields in AudioContext that change in real-time.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#type-ContextRealtimeData
type ContextRealtimeData struct {
	// The current context time in second in BaseAudioContext.
	CurrentTime float64 `json:"currentTime"`
	// The time spent on rendering graph divided by render quantum duration,
	// and multiplied by 100. 100 means the audio renderer reached the full
	// capacity and glitch may occur.
	RenderCapacity float64 `json:"renderCapacity"`
	// A running mean of callback interval.
	CallbackIntervalMean float64 `json:"callbackIntervalMean"`
	// A running variance of callback interval.
	CallbackIntervalVariance float64 `json:"callbackIntervalVariance"`
}

// BaseAudioContext data type. Protocol object for BaseAudioContext
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#type-BaseAudioContext
type BaseAudioContext struct {
	ContextID    string               `json:"contextId"`
	ContextType  ContextType          `json:"contextType"`
	ContextState ContextState         `json:"contextState"`
	RealtimeData *ContextRealtimeData `json:"realtimeData,omitempty"`
	// Platform-dependent callback buffer size.
	CallbackBufferSize float64 `json:"callbackBufferSize"`
	// Number of output channels supported by audio hardware in use.
	MaxOutputChannelCount float64 `json:"maxOutputChannelCount"`
	// Context sample rate.
	SampleRate float64 `json:"sampleRate"`
}

// AudioListener data type. Protocol object for AudioListener
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#type-AudioListener
type AudioListener struct {
	ListenerID string `json:"listenerId"`
	ContextID  string `json:"contextId"`
}

// AudioNode data type. Protocol object for AudioNode
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#type-AudioNode
type AudioNode struct {
	NodeID                string                `json:"nodeId"`
	ContextID             string                `json:"contextId"`
	NodeType              string                `json:"nodeType"`
	NumberOfInputs        float64               `json:"numberOfInputs"`
	NumberOfOutputs       float64               `json:"numberOfOutputs"`
	ChannelCount          float64               `json:"channelCount"`
	ChannelCountMode      ChannelCountMode      `json:"channelCountMode"`
	ChannelInterpretation ChannelInterpretation `json:"channelInterpretation"`
}

// AudioParam data type. Protocol object for AudioParam
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#type-AudioParam
type AudioParam struct {
	ParamID      string         `json:"paramId"`
	NodeID       string         `json:"nodeId"`
	ContextID    string         `json:"contextId"`
	ParamType    string         `json:"paramType"`
	Rate         AutomationRate `json:"rate"`
	DefaultValue float64        `json:"defaultValue"`
	MinValue     float64        `json:"minValue"`
	MaxValue     float64        `json:"maxValue"`
}
