package animation

// Canceled asynchronous event. Event for when an animation has been cancelled.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Animation/#event-animationCanceled
type Canceled struct {
	// Id of the animation that was cancelled.
	ID string `json:"id"`
}

// Created asynchronous event. Event for each animation that has been created.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Animation/#event-animationCreated
type Created struct {
	// Id of the animation that was created.
	ID string `json:"id"`
}

// Started asynchronous event. Event for animation that has been started.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Animation/#event-animationStarted
type Started struct {
	// Animation that was started.
	Animation Animation `json:"animation"`
}
