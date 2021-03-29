package inspector

// Fired when remote debugging connection is about to be terminated. Contains detach reason.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Inspector/#event-detached
type Detached struct {
	// The reason why connection has been terminated.
	Reason string `json:"reason"`
}

// Fired when debugging target has crashed
//
// https://chromedevtools.github.io/devtools-protocol/tot/Inspector/#event-targetCrashed
type TargetCrashed struct{}

// Fired when debugging target has reloaded after crash
//
// https://chromedevtools.github.io/devtools-protocol/tot/Inspector/#event-targetReloadedAfterCrash
type TargetReloadedAfterCrash struct{}
