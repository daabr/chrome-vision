package cdp

// UTC time in seconds, counted from January 1, 1970.
//
// This is a redefinition of the CDP type `Network.TimeSinceEpoch`
// (https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-TimeSinceEpoch)
// to work-around circular dependencies
// (https://bugs.chromium.org/p/chromium/issues/detail?id=1193242).
type TimeSinceEpoch float64

// Unique frame identifier.
//
// Redefinition of the CDP type `Page.FrameID`
// (https://chromedevtools.github.io/devtools-protocol/tot/Page/#type-FrameId)
// to work-around circular dependencies
// (https://bugs.chromium.org/p/chromium/issues/detail?id=1193242).
type FrameID string

// Redefinition of the CDP type `Target.TargetID`
// (https://chromedevtools.github.io/devtools-protocol/tot/Target/#type-TargetID)
// to work-around circular dependencies
// (https://bugs.chromium.org/p/chromium/issues/detail?id=1193242).
type TargetID string
