package runtime

import "encoding/json"

// BindingCalled asynchronous event.
//
// Notification is issued every time when binding is called.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-bindingCalled
//
// This CDP event is experimental.
type BindingCalled struct {
	Name    string `json:"name"`
	Payload string `json:"payload"`
	// Identifier of the context where the call was made.
	ExecutionContextID int64 `json:"executionContextId"`
}

// ConsoleAPICalled asynchronous event.
//
// Issued when console API was called.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-consoleAPICalled
type ConsoleAPICalled struct {
	// Type of the call.
	Type string `json:"type"`
	// Call arguments.
	Args []RemoteObject `json:"args"`
	// Identifier of the context where the call was made.
	ExecutionContextID int64 `json:"executionContextId"`
	// Call timestamp.
	Timestamp float64 `json:"timestamp"`
	// Stack trace captured when the call was made. The async stack chain is automatically reported for
	// the following call types: `assert`, `error`, `trace`, `warning`. For other types the async call
	// chain can be retrieved using `Debugger.getStackTrace` and `stackTrace.parentId` field.
	StackTrace *StackTrace `json:"stackTrace,omitempty"`
	// Console context descriptor for calls on non-default console context (not console.*):
	// 'anonymous#unique-logger-id' for call on unnamed context, 'name#unique-logger-id' for call
	// on named context.
	//
	// This CDP parameter is experimental.
	Context string `json:"context,omitempty"`
}

// ExceptionRevoked asynchronous event.
//
// Issued when unhandled exception was revoked.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-exceptionRevoked
type ExceptionRevoked struct {
	// Reason describing why exception was revoked.
	Reason string `json:"reason"`
	// The id of revoked exception, as reported in `exceptionThrown`.
	ExceptionID int64 `json:"exceptionId"`
}

// ExceptionThrown asynchronous event.
//
// Issued when exception was thrown and unhandled.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-exceptionThrown
type ExceptionThrown struct {
	// Timestamp of the exception.
	Timestamp        float64          `json:"timestamp"`
	ExceptionDetails ExceptionDetails `json:"exceptionDetails"`
}

// ExecutionContextCreated asynchronous event.
//
// Issued when new execution context is created.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-executionContextCreated
type ExecutionContextCreated struct {
	// A newly created execution context.
	Context ExecutionContextDescription `json:"context"`
}

// ExecutionContextDestroyed asynchronous event.
//
// Issued when execution context is destroyed.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-executionContextDestroyed
type ExecutionContextDestroyed struct {
	// Id of the destroyed context
	ExecutionContextID int64 `json:"executionContextId"`
}

// ExecutionContextsCleared asynchronous event.
//
// Issued when all executionContexts were cleared in browser
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-executionContextsCleared
type ExecutionContextsCleared struct{}

// InspectRequested asynchronous event.
//
// Issued when object should be inspected (for example, as a result of inspect() command line API
// call).
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-inspectRequested
type InspectRequested struct {
	Object RemoteObject    `json:"object"`
	Hints  json.RawMessage `json:"hints"`
}
