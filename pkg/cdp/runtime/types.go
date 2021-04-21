package runtime

import "encoding/json"

// Unique script identifier.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-ScriptId
type ScriptID string

// Unique object identifier.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-RemoteObjectId
type RemoteObjectID string

// Primitive value which cannot be JSON-stringified. Includes values `-0`, `NaN`, `Infinity`,
// `-Infinity`, and bigint literals.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-UnserializableValue
type UnserializableValue string

// Mirror object referencing original JavaScript object.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-RemoteObject
type RemoteObject struct {
	// Object type.
	Type string `json:"type"`
	// Object subtype hint. Specified for `object` type values only.
	// NOTE: If you change anything here, make sure to also update
	// `subtype` in `ObjectPreview` and `PropertyPreview` below.
	Subtype string `json:"subtype,omitempty"`
	// Object class (constructor) name. Specified for `object` type values only.
	ClassName string `json:"className,omitempty"`
	// Remote object value in case of primitive values or JSON values (if it was requested).
	Value json.RawMessage `json:"value,omitempty"`
	// Primitive value which can not be JSON-stringified does not have `value`, but gets this
	// property.
	UnserializableValue string `json:"unserializableValue,omitempty"`
	// String representation of the object.
	Description string `json:"description,omitempty"`
	// Unique object identifier (for non-primitive values).
	ObjectID string `json:"objectId,omitempty"`
	// Preview containing abbreviated property values. Specified for `object` type values only.
	//
	// This CDP property is experimental.
	Preview *ObjectPreview `json:"preview,omitempty"`
	// This CDP property is experimental.
	CustomPreview *CustomPreview `json:"customPreview,omitempty"`
}

// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-CustomPreview
//
// This CDP type is experimental.
type CustomPreview struct {
	// The JSON-stringified result of formatter.header(object, config) call.
	// It contains json ML array that represents RemoteObject.
	Header string `json:"header"`
	// If formatter returns true as a result of formatter.hasBody call then bodyGetterId will
	// contain RemoteObjectId for the function that returns result of formatter.body(object, config) call.
	// The result value is json ML array.
	BodyGetterID string `json:"bodyGetterId,omitempty"`
}

// Object containing abbreviated remote object value.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-ObjectPreview
//
// This CDP type is experimental.
type ObjectPreview struct {
	// Object type.
	Type string `json:"type"`
	// Object subtype hint. Specified for `object` type values only.
	Subtype string `json:"subtype,omitempty"`
	// String representation of the object.
	Description string `json:"description,omitempty"`
	// True iff some of the properties or entries of the original object did not fit.
	Overflow bool `json:"overflow"`
	// List of the properties.
	Properties []PropertyPreview `json:"properties"`
	// List of the entries. Specified for `map` and `set` subtype values only.
	Entries []EntryPreview `json:"entries,omitempty"`
}

// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-PropertyPreview
//
// This CDP type is experimental.
type PropertyPreview struct {
	// Property name.
	Name string `json:"name"`
	// Object type. Accessor means that the property itself is an accessor property.
	Type string `json:"type"`
	// User-friendly property value string.
	Value string `json:"value,omitempty"`
	// Nested value preview.
	ValuePreview *ObjectPreview `json:"valuePreview,omitempty"`
	// Object subtype hint. Specified for `object` type values only.
	Subtype string `json:"subtype,omitempty"`
}

// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-EntryPreview
//
// This CDP type is experimental.
type EntryPreview struct {
	// Preview of the key. Specified for map-like collection entries.
	Key *ObjectPreview `json:"key,omitempty"`
	// Preview of the value.
	Value ObjectPreview `json:"value"`
}

// Object property descriptor.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-PropertyDescriptor
type PropertyDescriptor struct {
	// Property name or symbol description.
	Name string `json:"name"`
	// The value associated with the property.
	Value *RemoteObject `json:"value,omitempty"`
	// True if the value associated with the property may be changed (data descriptors only).
	Writable bool `json:"writable,omitempty"`
	// A function which serves as a getter for the property, or `undefined` if there is no getter
	// (accessor descriptors only).
	Get *RemoteObject `json:"get,omitempty"`
	// A function which serves as a setter for the property, or `undefined` if there is no setter
	// (accessor descriptors only).
	Set *RemoteObject `json:"set,omitempty"`
	// True if the type of this property descriptor may be changed and if the property may be
	// deleted from the corresponding object.
	Configurable bool `json:"configurable"`
	// True if this property shows up during enumeration of the properties on the corresponding
	// object.
	Enumerable bool `json:"enumerable"`
	// True if the result was thrown during the evaluation.
	WasThrown bool `json:"wasThrown,omitempty"`
	// True if the property is owned for the object.
	IsOwn bool `json:"isOwn,omitempty"`
	// Property symbol object, if the property is of the `symbol` type.
	Symbol *RemoteObject `json:"symbol,omitempty"`
}

// Object internal property descriptor. This property isn't normally visible in JavaScript code.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-InternalPropertyDescriptor
type InternalPropertyDescriptor struct {
	// Conventional property name.
	Name string `json:"name"`
	// The value associated with the property.
	Value *RemoteObject `json:"value,omitempty"`
}

// Object private field descriptor.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-PrivatePropertyDescriptor
//
// This CDP type is experimental.
type PrivatePropertyDescriptor struct {
	// Private property name.
	Name string `json:"name"`
	// The value associated with the private property.
	Value *RemoteObject `json:"value,omitempty"`
	// A function which serves as a getter for the private property,
	// or `undefined` if there is no getter (accessor descriptors only).
	Get *RemoteObject `json:"get,omitempty"`
	// A function which serves as a setter for the private property,
	// or `undefined` if there is no setter (accessor descriptors only).
	Set *RemoteObject `json:"set,omitempty"`
}

// Represents function call argument. Either remote object id `objectId`, primitive `value`,
// unserializable primitive value or neither of (for undefined) them should be specified.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-CallArgument
type CallArgument struct {
	// Primitive value or serializable javascript object.
	Value json.RawMessage `json:"value,omitempty"`
	// Primitive value which can not be JSON-stringified.
	UnserializableValue string `json:"unserializableValue,omitempty"`
	// Remote object handle.
	ObjectID string `json:"objectId,omitempty"`
}

// Id of an execution context.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-ExecutionContextId
type ExecutionContextID int64

// Description of an isolated world.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-ExecutionContextDescription
type ExecutionContextDescription struct {
	// Unique id of the execution context. It can be used to specify in which execution context
	// script evaluation should be performed.
	ID int64 `json:"id"`
	// Execution context origin.
	Origin string `json:"origin"`
	// Human readable name describing given context.
	Name string `json:"name"`
	// A system-unique execution context identifier. Unlike the id, this is unique accross
	// multiple processes, so can be reliably used to identify specific context while backend
	// performs a cross-process navigation.
	//
	// This CDP property is experimental.
	UniqueID string `json:"uniqueId"`
	// Embedder-specific auxiliary data.
	AuxData json.RawMessage `json:"auxData,omitempty"`
}

// Detailed information about exception (or error) that was thrown during script compilation or
// execution.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-ExceptionDetails
type ExceptionDetails struct {
	// Exception id.
	ExceptionID int64 `json:"exceptionId"`
	// Exception text, which should be used together with exception object when available.
	Text string `json:"text"`
	// Line number of the exception location (0-based).
	LineNumber int64 `json:"lineNumber"`
	// Column number of the exception location (0-based).
	ColumnNumber int64 `json:"columnNumber"`
	// Script ID of the exception location.
	ScriptID string `json:"scriptId,omitempty"`
	// URL of the exception location, to be used when the script was not reported.
	URL string `json:"url,omitempty"`
	// JavaScript stack trace if available.
	StackTrace *StackTrace `json:"stackTrace,omitempty"`
	// Exception object if available.
	Exception *RemoteObject `json:"exception,omitempty"`
	// Identifier of the context where exception happened.
	ExecutionContextID int64 `json:"executionContextId,omitempty"`
}

// Number of milliseconds since epoch.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-Timestamp
type Timestamp float64

// Number of milliseconds.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-TimeDelta
type TimeDelta float64

// Stack entry for runtime errors and assertions.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-CallFrame
type CallFrame struct {
	// JavaScript function name.
	FunctionName string `json:"functionName"`
	// JavaScript script id.
	ScriptID string `json:"scriptId"`
	// JavaScript script name or url.
	URL string `json:"url"`
	// JavaScript script line number (0-based).
	LineNumber int64 `json:"lineNumber"`
	// JavaScript script column number (0-based).
	ColumnNumber int64 `json:"columnNumber"`
}

// Call frames for assertions or error messages.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-StackTrace
type StackTrace struct {
	// String label of this stack trace. For async traces this may be a name of the function that
	// initiated the async call.
	Description string `json:"description,omitempty"`
	// JavaScript function name.
	CallFrames []CallFrame `json:"callFrames"`
	// Asynchronous JavaScript stack trace that preceded this stack, if available.
	Parent *StackTrace `json:"parent,omitempty"`
	// Asynchronous JavaScript stack trace that preceded this stack, if available.
	//
	// This CDP property is experimental.
	ParentID *StackTraceID `json:"parentId,omitempty"`
}

// Unique identifier of current debugger.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-UniqueDebuggerId
//
// This CDP type is experimental.
type UniqueDebuggerID string

// If `debuggerId` is set stack trace comes from another debugger and can be resolved there. This
// allows to track cross-debugger calls. See `Runtime.StackTrace` and `Debugger.paused` for usages.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-StackTraceId
//
// This CDP type is experimental.
type StackTraceID struct {
	ID         string `json:"id"`
	DebuggerID string `json:"debuggerId,omitempty"`
}
