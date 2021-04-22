package headlessexperimental

// Issued when the target starts or stops needing BeginFrames.
// Deprecated. Issue beginFrame unconditionally instead and use result from
// beginFrame to detect whether the frames were suppressed.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#event-needsBeginFramesChanged
//
// This CDP event is deprecated.
type NeedsBeginFramesChanged struct {
	// True if BeginFrames are needed, false otherwise.
	NeedsBeginFrames bool
}
