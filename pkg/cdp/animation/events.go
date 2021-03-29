package animation

// Event for when an animation has been cancelled.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Animation/#event-animationCanceled
type AnimationCanceled struct {
	// Id of the animation that was cancelled.
	ID string `json:"id"`
}

// Event for each animation that has been created.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Animation/#event-animationCreated
type AnimationCreated struct {
	// Id of the animation that was created.
	ID string `json:"id"`
}

// Event for animation that has been started.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Animation/#event-animationStarted
type AnimationStarted struct {
	// Animation that was started.
	Animation Animation `json:"animation"`
}
