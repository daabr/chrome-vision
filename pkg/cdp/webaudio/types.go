package webaudio

// An unique ID for a graph object (AudioContext, AudioNode, AudioParam) in Web Audio API
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#type-GraphObjectId
type GraphObjectID string

// Enum of BaseAudioContext types
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#type-ContextType
type ContextType string

// ContextType valid values.
const (
	ContextTypeRealtime ContextType = "realtime"
	ContextTypeOffline  ContextType = "offline"
)

// Enum of AudioContextState from the spec
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#type-ContextState
type ContextState string

// ContextState valid values.
const (
	ContextStateSuspended ContextState = "suspended"
	ContextStateRunning   ContextState = "running"
	ContextStateClosed    ContextState = "closed"
)

// Enum of AudioNode types
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#type-NodeType
type NodeType string

// Enum of AudioNode::ChannelCountMode from the spec
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#type-ChannelCountMode
type ChannelCountMode string

// ChannelCountMode valid values.
const (
	ChannelCountModeClampedMax ChannelCountMode = "clamped-max"
	ChannelCountModeExplicit   ChannelCountMode = "explicit"
	ChannelCountModeMax        ChannelCountMode = "max"
)

// Enum of AudioNode::ChannelInterpretation from the spec
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#type-ChannelInterpretation
type ChannelInterpretation string

// ChannelInterpretation valid values.
const (
	ChannelInterpretationDiscrete ChannelInterpretation = "discrete"
	ChannelInterpretationSpeakers ChannelInterpretation = "speakers"
)

// Enum of AudioParam types
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#type-ParamType
type ParamType string

// Enum of AudioParam::AutomationRate from the spec
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#type-AutomationRate
type AutomationRate string

// AutomationRate valid values.
const (
	AutomationRateARate AutomationRate = "a-rate"
	AutomationRateKRate AutomationRate = "k-rate"
)

// Fields in AudioContext that change in real-time.
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#type-ContextRealtimeData
type ContextRealtimeData struct {
	// The current context time in second in BaseAudioContext.
	CurrentTime float64
	// The time spent on rendering graph divided by render qunatum duration,
	// and multiplied by 100. 100 means the audio renderer reached the full
	// capacity and glitch may occur.
	RenderCapacity float64
	// A running mean of callback interval.
	CallbackIntervalMean float64
	// A running variance of callback interval.
	CallbackIntervalVariance float64
}

// Protocol object for BaseAudioContext
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#type-BaseAudioContext
type BaseAudioContext struct {
	ContextID    string
	ContextType  string
	ContextState string
	RealtimeData *ContextRealtimeData `json:"realtimeData,omitempty"`
	// Platform-dependent callback buffer size.
	CallbackBufferSize float64
	// Number of output channels supported by audio hardware in use.
	MaxOutputChannelCount float64
	// Context sample rate.
	SampleRate float64
}

// Protocol object for AudioListener
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#type-AudioListener
type AudioListener struct {
	ListenerID string
	ContextID  string
}

// Protocol object for AudioNode
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#type-AudioNode
type AudioNode struct {
	NodeID                string
	ContextID             string
	NodeType              string
	NumberOfInputs        float64
	NumberOfOutputs       float64
	ChannelCount          float64
	ChannelCountMode      string
	ChannelInterpretation string
}

// Protocol object for AudioParam
//
// https://chromedevtools.github.io/devtools-protocol/tot/WebAudio/#type-AudioParam
type AudioParam struct {
	ParamID      string
	NodeID       string
	ContextID    string
	ParamType    string
	Rate         string
	DefaultValue float64
	MinValue     float64
	MaxValue     float64
}
