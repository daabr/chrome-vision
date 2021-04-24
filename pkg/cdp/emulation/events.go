package emulation

// VirtualTimeBudgetExpired asynchronous event.
//
// Notification sent after the virtual time budget for the current VirtualTimePolicy has run out.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#event-virtualTimeBudgetExpired
//
// This CDP event is experimental.
type VirtualTimeBudgetExpired struct{}
