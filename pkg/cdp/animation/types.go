package animation

// Animation instance.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Animation/#type-Animation
type Animation struct {
	// `Animation`'s id.
	ID string
	// `Animation`'s name.
	Name string
	// `Animation`'s internal paused state.
	PausedState bool
	// `Animation`'s play state.
	PlayState string
	// `Animation`'s playback rate.
	PlaybackRate float64
	// `Animation`'s start time.
	StartTime float64
	// `Animation`'s current time.
	CurrentTime float64
	// Animation type of `Animation`.
	Type string
	// `Animation`'s source animation node.
	Source *AnimationEffect `json:"source,omitempty"`
	// A unique ID for `Animation` representing the sources that triggered this CSS
	// animation/transition.
	CssID string `json:"cssId,omitempty"`
}

// AnimationEffect instance
//
// https://chromedevtools.github.io/devtools-protocol/tot/Animation/#type-AnimationEffect
type AnimationEffect struct {
	// `AnimationEffect`'s delay.
	Delay float64
	// `AnimationEffect`'s end delay.
	EndDelay float64
	// `AnimationEffect`'s iteration start.
	IterationStart float64
	// `AnimationEffect`'s iterations.
	Iterations float64
	// `AnimationEffect`'s iteration duration.
	Duration float64
	// `AnimationEffect`'s playback direction.
	Direction string
	// `AnimationEffect`'s fill mode.
	Fill string
	// `AnimationEffect`'s target node.
	BackendNodeID int64 `json:"backendNodeId,omitempty"`
	// `AnimationEffect`'s keyframes.
	KeyframesRule *KeyframesRule `json:"keyframesRule,omitempty"`
	// `AnimationEffect`'s timing function.
	Easing string
}

// Keyframes Rule
//
// https://chromedevtools.github.io/devtools-protocol/tot/Animation/#type-KeyframesRule
type KeyframesRule struct {
	// CSS keyframed animation's name.
	Name string `json:"name,omitempty"`
	// List of animation keyframes.
	Keyframes []KeyframeStyle
}

// Keyframe Style
//
// https://chromedevtools.github.io/devtools-protocol/tot/Animation/#type-KeyframeStyle
type KeyframeStyle struct {
	// Keyframe's time offset.
	Offset string
	// `AnimationEffect`'s timing function.
	Easing string
}
