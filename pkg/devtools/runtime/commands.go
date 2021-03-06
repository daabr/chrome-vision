package runtime

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
)

// AwaitPromise contains the parameters, and acts as
// a Go receiver, for the CDP command `awaitPromise`.
//
// Add handler to promise with given promise object id.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-awaitPromise
type AwaitPromise struct {
	// Identifier of the promise.
	PromiseObjectID string `json:"promiseObjectId"`
	// Whether the result is expected to be a JSON object that should be sent by value.
	ReturnByValue bool `json:"returnByValue,omitempty"`
	// Whether preview should be generated for the result.
	GeneratePreview bool `json:"generatePreview,omitempty"`
}

// NewAwaitPromise constructs a new AwaitPromise struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-awaitPromise
func NewAwaitPromise(promiseObjectID string) *AwaitPromise {
	return &AwaitPromise{
		PromiseObjectID: promiseObjectID,
	}
}

// SetReturnByValue adds or modifies the value of the optional
// parameter `returnByValue` in the AwaitPromise CDP command.
//
// Whether the result is expected to be a JSON object that should be sent by value.
func (t *AwaitPromise) SetReturnByValue(v bool) *AwaitPromise {
	t.ReturnByValue = v
	return t
}

// SetGeneratePreview adds or modifies the value of the optional
// parameter `generatePreview` in the AwaitPromise CDP command.
//
// Whether preview should be generated for the result.
func (t *AwaitPromise) SetGeneratePreview(v bool) *AwaitPromise {
	t.GeneratePreview = v
	return t
}

// AwaitPromiseResult contains the browser's response
// to calling the AwaitPromise CDP command with Do().
type AwaitPromiseResult struct {
	// Promise result. Will contain rejected value if promise was rejected.
	Result RemoteObject `json:"result"`
	// Exception details if stack strace is available.
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"`
}

// Do sends the AwaitPromise CDP command to a browser,
// and returns the browser's response.
func (t *AwaitPromise) Do(ctx context.Context) (*AwaitPromiseResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Runtime.awaitPromise", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the AwaitPromise CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *AwaitPromise) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Runtime.awaitPromise", b)
}

// ParseResponse parses the browser's response
// to the AwaitPromise CDP command.
func (t *AwaitPromise) ParseResponse(m *devtools.Message) (*AwaitPromiseResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &AwaitPromiseResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// CallFunctionOn contains the parameters, and acts as
// a Go receiver, for the CDP command `callFunctionOn`.
//
// Calls function with given declaration on the given object. Object group of the result is
// inherited from the target object.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-callFunctionOn
type CallFunctionOn struct {
	// Declaration of the function to call.
	FunctionDeclaration string `json:"functionDeclaration"`
	// Identifier of the object to call function on. Either objectId or executionContextId should
	// be specified.
	ObjectID string `json:"objectId,omitempty"`
	// Call arguments. All call arguments must belong to the same JavaScript world as the target
	// object.
	Arguments []CallArgument `json:"arguments,omitempty"`
	// In silent mode exceptions thrown during evaluation are not reported and do not pause
	// execution. Overrides `setPauseOnException` state.
	Silent bool `json:"silent,omitempty"`
	// Whether the result is expected to be a JSON object which should be sent by value.
	ReturnByValue bool `json:"returnByValue,omitempty"`
	// Whether preview should be generated for the result.
	//
	// This CDP parameter is experimental.
	GeneratePreview bool `json:"generatePreview,omitempty"`
	// Whether execution should be treated as initiated by user in the UI.
	UserGesture bool `json:"userGesture,omitempty"`
	// Whether execution should `await` for resulting value and return once awaited promise is
	// resolved.
	AwaitPromise bool `json:"awaitPromise,omitempty"`
	// Specifies execution context which global object will be used to call function on. Either
	// executionContextId or objectId should be specified.
	ExecutionContextID int64 `json:"executionContextId,omitempty"`
	// Symbolic group name that can be used to release multiple objects. If objectGroup is not
	// specified and objectId is, objectGroup will be inherited from object.
	ObjectGroup string `json:"objectGroup,omitempty"`
	// Whether to throw an exception if side effect cannot be ruled out during evaluation.
	//
	// This CDP parameter is experimental.
	ThrowOnSideEffect bool `json:"throwOnSideEffect,omitempty"`
}

// NewCallFunctionOn constructs a new CallFunctionOn struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-callFunctionOn
func NewCallFunctionOn(functionDeclaration string) *CallFunctionOn {
	return &CallFunctionOn{
		FunctionDeclaration: functionDeclaration,
	}
}

// SetObjectID adds or modifies the value of the optional
// parameter `objectId` in the CallFunctionOn CDP command.
//
// Identifier of the object to call function on. Either objectId or executionContextId should
// be specified.
func (t *CallFunctionOn) SetObjectID(v string) *CallFunctionOn {
	t.ObjectID = v
	return t
}

// SetArguments adds or modifies the value of the optional
// parameter `arguments` in the CallFunctionOn CDP command.
//
// Call arguments. All call arguments must belong to the same JavaScript world as the target
// object.
func (t *CallFunctionOn) SetArguments(v []CallArgument) *CallFunctionOn {
	t.Arguments = v
	return t
}

// SetSilent adds or modifies the value of the optional
// parameter `silent` in the CallFunctionOn CDP command.
//
// In silent mode exceptions thrown during evaluation are not reported and do not pause
// execution. Overrides `setPauseOnException` state.
func (t *CallFunctionOn) SetSilent(v bool) *CallFunctionOn {
	t.Silent = v
	return t
}

// SetReturnByValue adds or modifies the value of the optional
// parameter `returnByValue` in the CallFunctionOn CDP command.
//
// Whether the result is expected to be a JSON object which should be sent by value.
func (t *CallFunctionOn) SetReturnByValue(v bool) *CallFunctionOn {
	t.ReturnByValue = v
	return t
}

// SetGeneratePreview adds or modifies the value of the optional
// parameter `generatePreview` in the CallFunctionOn CDP command.
//
// Whether preview should be generated for the result.
//
// This CDP parameter is experimental.
func (t *CallFunctionOn) SetGeneratePreview(v bool) *CallFunctionOn {
	t.GeneratePreview = v
	return t
}

// SetUserGesture adds or modifies the value of the optional
// parameter `userGesture` in the CallFunctionOn CDP command.
//
// Whether execution should be treated as initiated by user in the UI.
func (t *CallFunctionOn) SetUserGesture(v bool) *CallFunctionOn {
	t.UserGesture = v
	return t
}

// SetAwaitPromise adds or modifies the value of the optional
// parameter `awaitPromise` in the CallFunctionOn CDP command.
//
// Whether execution should `await` for resulting value and return once awaited promise is
// resolved.
func (t *CallFunctionOn) SetAwaitPromise(v bool) *CallFunctionOn {
	t.AwaitPromise = v
	return t
}

// SetExecutionContextID adds or modifies the value of the optional
// parameter `executionContextId` in the CallFunctionOn CDP command.
//
// Specifies execution context which global object will be used to call function on. Either
// executionContextId or objectId should be specified.
func (t *CallFunctionOn) SetExecutionContextID(v int64) *CallFunctionOn {
	t.ExecutionContextID = v
	return t
}

// SetObjectGroup adds or modifies the value of the optional
// parameter `objectGroup` in the CallFunctionOn CDP command.
//
// Symbolic group name that can be used to release multiple objects. If objectGroup is not
// specified and objectId is, objectGroup will be inherited from object.
func (t *CallFunctionOn) SetObjectGroup(v string) *CallFunctionOn {
	t.ObjectGroup = v
	return t
}

// SetThrowOnSideEffect adds or modifies the value of the optional
// parameter `throwOnSideEffect` in the CallFunctionOn CDP command.
//
// Whether to throw an exception if side effect cannot be ruled out during evaluation.
//
// This CDP parameter is experimental.
func (t *CallFunctionOn) SetThrowOnSideEffect(v bool) *CallFunctionOn {
	t.ThrowOnSideEffect = v
	return t
}

// CallFunctionOnResult contains the browser's response
// to calling the CallFunctionOn CDP command with Do().
type CallFunctionOnResult struct {
	// Call result.
	Result RemoteObject `json:"result"`
	// Exception details.
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"`
}

// Do sends the CallFunctionOn CDP command to a browser,
// and returns the browser's response.
func (t *CallFunctionOn) Do(ctx context.Context) (*CallFunctionOnResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Runtime.callFunctionOn", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the CallFunctionOn CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *CallFunctionOn) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Runtime.callFunctionOn", b)
}

// ParseResponse parses the browser's response
// to the CallFunctionOn CDP command.
func (t *CallFunctionOn) ParseResponse(m *devtools.Message) (*CallFunctionOnResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &CallFunctionOnResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// CompileScript contains the parameters, and acts as
// a Go receiver, for the CDP command `compileScript`.
//
// Compiles expression.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-compileScript
type CompileScript struct {
	// Expression to compile.
	Expression string `json:"expression"`
	// Source url to be set for the script.
	SourceURL string `json:"sourceURL"`
	// Specifies whether the compiled script should be persisted.
	PersistScript bool `json:"persistScript"`
	// Specifies in which execution context to perform script run. If the parameter is omitted the
	// evaluation will be performed in the context of the inspected page.
	ExecutionContextID int64 `json:"executionContextId,omitempty"`
}

// NewCompileScript constructs a new CompileScript struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-compileScript
func NewCompileScript(expression string, sourceURL string, persistScript bool) *CompileScript {
	return &CompileScript{
		Expression:    expression,
		SourceURL:     sourceURL,
		PersistScript: persistScript,
	}
}

// SetExecutionContextID adds or modifies the value of the optional
// parameter `executionContextId` in the CompileScript CDP command.
//
// Specifies in which execution context to perform script run. If the parameter is omitted the
// evaluation will be performed in the context of the inspected page.
func (t *CompileScript) SetExecutionContextID(v int64) *CompileScript {
	t.ExecutionContextID = v
	return t
}

// CompileScriptResult contains the browser's response
// to calling the CompileScript CDP command with Do().
type CompileScriptResult struct {
	// Id of the script.
	ScriptID string `json:"scriptId,omitempty"`
	// Exception details.
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"`
}

// Do sends the CompileScript CDP command to a browser,
// and returns the browser's response.
func (t *CompileScript) Do(ctx context.Context) (*CompileScriptResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Runtime.compileScript", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the CompileScript CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *CompileScript) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Runtime.compileScript", b)
}

// ParseResponse parses the browser's response
// to the CompileScript CDP command.
func (t *CompileScript) ParseResponse(m *devtools.Message) (*CompileScriptResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &CompileScriptResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Disable contains the parameters, and acts as
// a Go receiver, for the CDP command `disable`.
//
// Disables reporting of execution contexts creation.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-disable
type Disable struct{}

// NewDisable constructs a new Disable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-disable
func NewDisable() *Disable {
	return &Disable{}
}

// Do sends the Disable CDP command to a browser,
// and returns the browser's response.
func (t *Disable) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Runtime.disable", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the Disable CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *Disable) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Runtime.disable", nil)
}

// ParseResponse parses the browser's response
// to the Disable CDP command.
func (t *Disable) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// DiscardConsoleEntries contains the parameters, and acts as
// a Go receiver, for the CDP command `discardConsoleEntries`.
//
// Discards collected exceptions and console API calls.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-discardConsoleEntries
type DiscardConsoleEntries struct{}

// NewDiscardConsoleEntries constructs a new DiscardConsoleEntries struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-discardConsoleEntries
func NewDiscardConsoleEntries() *DiscardConsoleEntries {
	return &DiscardConsoleEntries{}
}

// Do sends the DiscardConsoleEntries CDP command to a browser,
// and returns the browser's response.
func (t *DiscardConsoleEntries) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Runtime.discardConsoleEntries", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the DiscardConsoleEntries CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *DiscardConsoleEntries) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Runtime.discardConsoleEntries", nil)
}

// ParseResponse parses the browser's response
// to the DiscardConsoleEntries CDP command.
func (t *DiscardConsoleEntries) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// Enable contains the parameters, and acts as
// a Go receiver, for the CDP command `enable`.
//
// Enables reporting of execution contexts creation by means of `executionContextCreated` event.
// When the reporting gets enabled the event will be sent immediately for each existing execution
// context.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-enable
type Enable struct{}

// NewEnable constructs a new Enable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-enable
func NewEnable() *Enable {
	return &Enable{}
}

// Do sends the Enable CDP command to a browser,
// and returns the browser's response.
func (t *Enable) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Runtime.enable", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the Enable CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *Enable) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Runtime.enable", nil)
}

// ParseResponse parses the browser's response
// to the Enable CDP command.
func (t *Enable) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// Evaluate contains the parameters, and acts as
// a Go receiver, for the CDP command `evaluate`.
//
// Evaluates expression on global object.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-evaluate
type Evaluate struct {
	// Expression to evaluate.
	Expression string `json:"expression"`
	// Symbolic group name that can be used to release multiple objects.
	ObjectGroup string `json:"objectGroup,omitempty"`
	// Determines whether Command Line API should be available during the evaluation.
	IncludeCommandLineAPI bool `json:"includeCommandLineAPI,omitempty"`
	// In silent mode exceptions thrown during evaluation are not reported and do not pause
	// execution. Overrides `setPauseOnException` state.
	Silent bool `json:"silent,omitempty"`
	// Specifies in which execution context to perform evaluation. If the parameter is omitted the
	// evaluation will be performed in the context of the inspected page.
	// This is mutually exclusive with `uniqueContextId`, which offers an
	// alternative way to identify the execution context that is more reliable
	// in a multi-process environment.
	ContextID int64 `json:"contextId,omitempty"`
	// Whether the result is expected to be a JSON object that should be sent by value.
	ReturnByValue bool `json:"returnByValue,omitempty"`
	// Whether preview should be generated for the result.
	//
	// This CDP parameter is experimental.
	GeneratePreview bool `json:"generatePreview,omitempty"`
	// Whether execution should be treated as initiated by user in the UI.
	UserGesture bool `json:"userGesture,omitempty"`
	// Whether execution should `await` for resulting value and return once awaited promise is
	// resolved.
	AwaitPromise bool `json:"awaitPromise,omitempty"`
	// Whether to throw an exception if side effect cannot be ruled out during evaluation.
	// This implies `disableBreaks` below.
	//
	// This CDP parameter is experimental.
	ThrowOnSideEffect bool `json:"throwOnSideEffect,omitempty"`
	// Terminate execution after timing out (number of milliseconds).
	//
	// This CDP parameter is experimental.
	Timeout float64 `json:"timeout,omitempty"`
	// Disable breakpoints during execution.
	//
	// This CDP parameter is experimental.
	DisableBreaks bool `json:"disableBreaks,omitempty"`
	// Setting this flag to true enables `let` re-declaration and top-level `await`.
	// Note that `let` variables can only be re-declared if they originate from
	// `replMode` themselves.
	//
	// This CDP parameter is experimental.
	ReplMode bool `json:"replMode,omitempty"`
	// The Content Security Policy (CSP) for the target might block 'unsafe-eval'
	// which includes eval(), Function(), setTimeout() and setInterval()
	// when called with non-callable arguments. This flag bypasses CSP for this
	// evaluation and allows unsafe-eval. Defaults to true.
	//
	// This CDP parameter is experimental.
	AllowUnsafeEvalBlockedByCSP bool `json:"allowUnsafeEvalBlockedByCSP,omitempty"`
	// An alternative way to specify the execution context to evaluate in.
	// Compared to contextId that may be reused across processes, this is guaranteed to be
	// system-unique, so it can be used to prevent accidental evaluation of the expression
	// in context different than intended (e.g. as a result of navigation across process
	// boundaries).
	// This is mutually exclusive with `contextId`.
	//
	// This CDP parameter is experimental.
	UniqueContextID string `json:"uniqueContextId,omitempty"`
}

// NewEvaluate constructs a new Evaluate struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-evaluate
func NewEvaluate(expression string) *Evaluate {
	return &Evaluate{
		Expression: expression,
	}
}

// SetObjectGroup adds or modifies the value of the optional
// parameter `objectGroup` in the Evaluate CDP command.
//
// Symbolic group name that can be used to release multiple objects.
func (t *Evaluate) SetObjectGroup(v string) *Evaluate {
	t.ObjectGroup = v
	return t
}

// SetIncludeCommandLineAPI adds or modifies the value of the optional
// parameter `includeCommandLineAPI` in the Evaluate CDP command.
//
// Determines whether Command Line API should be available during the evaluation.
func (t *Evaluate) SetIncludeCommandLineAPI(v bool) *Evaluate {
	t.IncludeCommandLineAPI = v
	return t
}

// SetSilent adds or modifies the value of the optional
// parameter `silent` in the Evaluate CDP command.
//
// In silent mode exceptions thrown during evaluation are not reported and do not pause
// execution. Overrides `setPauseOnException` state.
func (t *Evaluate) SetSilent(v bool) *Evaluate {
	t.Silent = v
	return t
}

// SetContextID adds or modifies the value of the optional
// parameter `contextId` in the Evaluate CDP command.
//
// Specifies in which execution context to perform evaluation. If the parameter is omitted the
// evaluation will be performed in the context of the inspected page.
// This is mutually exclusive with `uniqueContextId`, which offers an
// alternative way to identify the execution context that is more reliable
// in a multi-process environment.
func (t *Evaluate) SetContextID(v int64) *Evaluate {
	t.ContextID = v
	return t
}

// SetReturnByValue adds or modifies the value of the optional
// parameter `returnByValue` in the Evaluate CDP command.
//
// Whether the result is expected to be a JSON object that should be sent by value.
func (t *Evaluate) SetReturnByValue(v bool) *Evaluate {
	t.ReturnByValue = v
	return t
}

// SetGeneratePreview adds or modifies the value of the optional
// parameter `generatePreview` in the Evaluate CDP command.
//
// Whether preview should be generated for the result.
//
// This CDP parameter is experimental.
func (t *Evaluate) SetGeneratePreview(v bool) *Evaluate {
	t.GeneratePreview = v
	return t
}

// SetUserGesture adds or modifies the value of the optional
// parameter `userGesture` in the Evaluate CDP command.
//
// Whether execution should be treated as initiated by user in the UI.
func (t *Evaluate) SetUserGesture(v bool) *Evaluate {
	t.UserGesture = v
	return t
}

// SetAwaitPromise adds or modifies the value of the optional
// parameter `awaitPromise` in the Evaluate CDP command.
//
// Whether execution should `await` for resulting value and return once awaited promise is
// resolved.
func (t *Evaluate) SetAwaitPromise(v bool) *Evaluate {
	t.AwaitPromise = v
	return t
}

// SetThrowOnSideEffect adds or modifies the value of the optional
// parameter `throwOnSideEffect` in the Evaluate CDP command.
//
// Whether to throw an exception if side effect cannot be ruled out during evaluation.
// This implies `disableBreaks` below.
//
// This CDP parameter is experimental.
func (t *Evaluate) SetThrowOnSideEffect(v bool) *Evaluate {
	t.ThrowOnSideEffect = v
	return t
}

// SetTimeout adds or modifies the value of the optional
// parameter `timeout` in the Evaluate CDP command.
//
// Terminate execution after timing out (number of milliseconds).
//
// This CDP parameter is experimental.
func (t *Evaluate) SetTimeout(v float64) *Evaluate {
	t.Timeout = v
	return t
}

// SetDisableBreaks adds or modifies the value of the optional
// parameter `disableBreaks` in the Evaluate CDP command.
//
// Disable breakpoints during execution.
//
// This CDP parameter is experimental.
func (t *Evaluate) SetDisableBreaks(v bool) *Evaluate {
	t.DisableBreaks = v
	return t
}

// SetReplMode adds or modifies the value of the optional
// parameter `replMode` in the Evaluate CDP command.
//
// Setting this flag to true enables `let` re-declaration and top-level `await`.
// Note that `let` variables can only be re-declared if they originate from
// `replMode` themselves.
//
// This CDP parameter is experimental.
func (t *Evaluate) SetReplMode(v bool) *Evaluate {
	t.ReplMode = v
	return t
}

// SetAllowUnsafeEvalBlockedByCSP adds or modifies the value of the optional
// parameter `allowUnsafeEvalBlockedByCSP` in the Evaluate CDP command.
//
// The Content Security Policy (CSP) for the target might block 'unsafe-eval'
// which includes eval(), Function(), setTimeout() and setInterval()
// when called with non-callable arguments. This flag bypasses CSP for this
// evaluation and allows unsafe-eval. Defaults to true.
//
// This CDP parameter is experimental.
func (t *Evaluate) SetAllowUnsafeEvalBlockedByCSP(v bool) *Evaluate {
	t.AllowUnsafeEvalBlockedByCSP = v
	return t
}

// SetUniqueContextID adds or modifies the value of the optional
// parameter `uniqueContextId` in the Evaluate CDP command.
//
// An alternative way to specify the execution context to evaluate in.
// Compared to contextId that may be reused across processes, this is guaranteed to be
// system-unique, so it can be used to prevent accidental evaluation of the expression
// in context different than intended (e.g. as a result of navigation across process
// boundaries).
// This is mutually exclusive with `contextId`.
//
// This CDP parameter is experimental.
func (t *Evaluate) SetUniqueContextID(v string) *Evaluate {
	t.UniqueContextID = v
	return t
}

// EvaluateResult contains the browser's response
// to calling the Evaluate CDP command with Do().
type EvaluateResult struct {
	// Evaluation result.
	Result RemoteObject `json:"result"`
	// Exception details.
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"`
}

// Do sends the Evaluate CDP command to a browser,
// and returns the browser's response.
func (t *Evaluate) Do(ctx context.Context) (*EvaluateResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Runtime.evaluate", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the Evaluate CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *Evaluate) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Runtime.evaluate", b)
}

// ParseResponse parses the browser's response
// to the Evaluate CDP command.
func (t *Evaluate) ParseResponse(m *devtools.Message) (*EvaluateResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &EvaluateResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetIsolateID contains the parameters, and acts as
// a Go receiver, for the CDP command `getIsolateId`.
//
// Returns the isolate id.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-getIsolateId
//
// This CDP method is experimental.
type GetIsolateID struct{}

// NewGetIsolateID constructs a new GetIsolateID struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-getIsolateId
//
// This CDP method is experimental.
func NewGetIsolateID() *GetIsolateID {
	return &GetIsolateID{}
}

// GetIsolateIDResult contains the browser's response
// to calling the GetIsolateID CDP command with Do().
type GetIsolateIDResult struct {
	// The isolate id.
	ID string `json:"id"`
}

// Do sends the GetIsolateID CDP command to a browser,
// and returns the browser's response.
func (t *GetIsolateID) Do(ctx context.Context) (*GetIsolateIDResult, error) {
	m, err := devtools.SendAndWait(ctx, "Runtime.getIsolateId", nil)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetIsolateID CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetIsolateID) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Runtime.getIsolateId", nil)
}

// ParseResponse parses the browser's response
// to the GetIsolateID CDP command.
func (t *GetIsolateID) ParseResponse(m *devtools.Message) (*GetIsolateIDResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetIsolateIDResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetHeapUsage contains the parameters, and acts as
// a Go receiver, for the CDP command `getHeapUsage`.
//
// Returns the JavaScript heap usage.
// It is the total usage of the corresponding isolate not scoped to a particular Runtime.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-getHeapUsage
//
// This CDP method is experimental.
type GetHeapUsage struct{}

// NewGetHeapUsage constructs a new GetHeapUsage struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-getHeapUsage
//
// This CDP method is experimental.
func NewGetHeapUsage() *GetHeapUsage {
	return &GetHeapUsage{}
}

// GetHeapUsageResult contains the browser's response
// to calling the GetHeapUsage CDP command with Do().
type GetHeapUsageResult struct {
	// Used heap size in bytes.
	UsedSize float64 `json:"usedSize"`
	// Allocated heap size in bytes.
	TotalSize float64 `json:"totalSize"`
}

// Do sends the GetHeapUsage CDP command to a browser,
// and returns the browser's response.
func (t *GetHeapUsage) Do(ctx context.Context) (*GetHeapUsageResult, error) {
	m, err := devtools.SendAndWait(ctx, "Runtime.getHeapUsage", nil)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetHeapUsage CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetHeapUsage) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Runtime.getHeapUsage", nil)
}

// ParseResponse parses the browser's response
// to the GetHeapUsage CDP command.
func (t *GetHeapUsage) ParseResponse(m *devtools.Message) (*GetHeapUsageResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetHeapUsageResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetProperties contains the parameters, and acts as
// a Go receiver, for the CDP command `getProperties`.
//
// Returns properties of a given object. Object group of the result is inherited from the target
// object.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-getProperties
type GetProperties struct {
	// Identifier of the object to return properties for.
	ObjectID string `json:"objectId"`
	// If true, returns properties belonging only to the element itself, not to its prototype
	// chain.
	OwnProperties bool `json:"ownProperties,omitempty"`
	// If true, returns accessor properties (with getter/setter) only; internal properties are not
	// returned either.
	//
	// This CDP parameter is experimental.
	AccessorPropertiesOnly bool `json:"accessorPropertiesOnly,omitempty"`
	// Whether preview should be generated for the results.
	//
	// This CDP parameter is experimental.
	GeneratePreview bool `json:"generatePreview,omitempty"`
	// If true, returns non-indexed properties only.
	//
	// This CDP parameter is experimental.
	NonIndexedPropertiesOnly bool `json:"nonIndexedPropertiesOnly,omitempty"`
}

// NewGetProperties constructs a new GetProperties struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-getProperties
func NewGetProperties(objectID string) *GetProperties {
	return &GetProperties{
		ObjectID: objectID,
	}
}

// SetOwnProperties adds or modifies the value of the optional
// parameter `ownProperties` in the GetProperties CDP command.
//
// If true, returns properties belonging only to the element itself, not to its prototype
// chain.
func (t *GetProperties) SetOwnProperties(v bool) *GetProperties {
	t.OwnProperties = v
	return t
}

// SetAccessorPropertiesOnly adds or modifies the value of the optional
// parameter `accessorPropertiesOnly` in the GetProperties CDP command.
//
// If true, returns accessor properties (with getter/setter) only; internal properties are not
// returned either.
//
// This CDP parameter is experimental.
func (t *GetProperties) SetAccessorPropertiesOnly(v bool) *GetProperties {
	t.AccessorPropertiesOnly = v
	return t
}

// SetGeneratePreview adds or modifies the value of the optional
// parameter `generatePreview` in the GetProperties CDP command.
//
// Whether preview should be generated for the results.
//
// This CDP parameter is experimental.
func (t *GetProperties) SetGeneratePreview(v bool) *GetProperties {
	t.GeneratePreview = v
	return t
}

// SetNonIndexedPropertiesOnly adds or modifies the value of the optional
// parameter `nonIndexedPropertiesOnly` in the GetProperties CDP command.
//
// If true, returns non-indexed properties only.
//
// This CDP parameter is experimental.
func (t *GetProperties) SetNonIndexedPropertiesOnly(v bool) *GetProperties {
	t.NonIndexedPropertiesOnly = v
	return t
}

// GetPropertiesResult contains the browser's response
// to calling the GetProperties CDP command with Do().
type GetPropertiesResult struct {
	// Object properties.
	Result []PropertyDescriptor `json:"result"`
	// Internal object properties (only of the element itself).
	InternalProperties []InternalPropertyDescriptor `json:"internalProperties,omitempty"`
	// Object private properties.
	//
	// This CDP parameter is experimental.
	PrivateProperties []PrivatePropertyDescriptor `json:"privateProperties,omitempty"`
	// Exception details.
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"`
}

// Do sends the GetProperties CDP command to a browser,
// and returns the browser's response.
func (t *GetProperties) Do(ctx context.Context) (*GetPropertiesResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Runtime.getProperties", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetProperties CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetProperties) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Runtime.getProperties", b)
}

// ParseResponse parses the browser's response
// to the GetProperties CDP command.
func (t *GetProperties) ParseResponse(m *devtools.Message) (*GetPropertiesResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetPropertiesResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GlobalLexicalScopeNames contains the parameters, and acts as
// a Go receiver, for the CDP command `globalLexicalScopeNames`.
//
// Returns all let, const and class variables from global scope.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-globalLexicalScopeNames
type GlobalLexicalScopeNames struct {
	// Specifies in which execution context to lookup global scope variables.
	ExecutionContextID int64 `json:"executionContextId,omitempty"`
}

// NewGlobalLexicalScopeNames constructs a new GlobalLexicalScopeNames struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-globalLexicalScopeNames
func NewGlobalLexicalScopeNames() *GlobalLexicalScopeNames {
	return &GlobalLexicalScopeNames{}
}

// SetExecutionContextID adds or modifies the value of the optional
// parameter `executionContextId` in the GlobalLexicalScopeNames CDP command.
//
// Specifies in which execution context to lookup global scope variables.
func (t *GlobalLexicalScopeNames) SetExecutionContextID(v int64) *GlobalLexicalScopeNames {
	t.ExecutionContextID = v
	return t
}

// GlobalLexicalScopeNamesResult contains the browser's response
// to calling the GlobalLexicalScopeNames CDP command with Do().
type GlobalLexicalScopeNamesResult struct {
	Names []string `json:"names"`
}

// Do sends the GlobalLexicalScopeNames CDP command to a browser,
// and returns the browser's response.
func (t *GlobalLexicalScopeNames) Do(ctx context.Context) (*GlobalLexicalScopeNamesResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Runtime.globalLexicalScopeNames", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GlobalLexicalScopeNames CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GlobalLexicalScopeNames) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Runtime.globalLexicalScopeNames", b)
}

// ParseResponse parses the browser's response
// to the GlobalLexicalScopeNames CDP command.
func (t *GlobalLexicalScopeNames) ParseResponse(m *devtools.Message) (*GlobalLexicalScopeNamesResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GlobalLexicalScopeNamesResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// QueryObjects contains the parameters, and acts as
// a Go receiver, for the CDP command `queryObjects`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-queryObjects
type QueryObjects struct {
	// Identifier of the prototype to return objects for.
	PrototypeObjectID string `json:"prototypeObjectId"`
	// Symbolic group name that can be used to release the results.
	ObjectGroup string `json:"objectGroup,omitempty"`
}

// NewQueryObjects constructs a new QueryObjects struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-queryObjects
func NewQueryObjects(prototypeObjectID string) *QueryObjects {
	return &QueryObjects{
		PrototypeObjectID: prototypeObjectID,
	}
}

// SetObjectGroup adds or modifies the value of the optional
// parameter `objectGroup` in the QueryObjects CDP command.
//
// Symbolic group name that can be used to release the results.
func (t *QueryObjects) SetObjectGroup(v string) *QueryObjects {
	t.ObjectGroup = v
	return t
}

// QueryObjectsResult contains the browser's response
// to calling the QueryObjects CDP command with Do().
type QueryObjectsResult struct {
	// Array with objects.
	Objects RemoteObject `json:"objects"`
}

// Do sends the QueryObjects CDP command to a browser,
// and returns the browser's response.
func (t *QueryObjects) Do(ctx context.Context) (*QueryObjectsResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Runtime.queryObjects", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the QueryObjects CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *QueryObjects) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Runtime.queryObjects", b)
}

// ParseResponse parses the browser's response
// to the QueryObjects CDP command.
func (t *QueryObjects) ParseResponse(m *devtools.Message) (*QueryObjectsResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &QueryObjectsResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ReleaseObject contains the parameters, and acts as
// a Go receiver, for the CDP command `releaseObject`.
//
// Releases remote object with given id.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-releaseObject
type ReleaseObject struct {
	// Identifier of the object to release.
	ObjectID string `json:"objectId"`
}

// NewReleaseObject constructs a new ReleaseObject struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-releaseObject
func NewReleaseObject(objectID string) *ReleaseObject {
	return &ReleaseObject{
		ObjectID: objectID,
	}
}

// Do sends the ReleaseObject CDP command to a browser,
// and returns the browser's response.
func (t *ReleaseObject) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Runtime.releaseObject", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the ReleaseObject CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *ReleaseObject) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Runtime.releaseObject", b)
}

// ParseResponse parses the browser's response
// to the ReleaseObject CDP command.
func (t *ReleaseObject) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// ReleaseObjectGroup contains the parameters, and acts as
// a Go receiver, for the CDP command `releaseObjectGroup`.
//
// Releases all remote objects that belong to a given group.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-releaseObjectGroup
type ReleaseObjectGroup struct {
	// Symbolic object group name.
	ObjectGroup string `json:"objectGroup"`
}

// NewReleaseObjectGroup constructs a new ReleaseObjectGroup struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-releaseObjectGroup
func NewReleaseObjectGroup(objectGroup string) *ReleaseObjectGroup {
	return &ReleaseObjectGroup{
		ObjectGroup: objectGroup,
	}
}

// Do sends the ReleaseObjectGroup CDP command to a browser,
// and returns the browser's response.
func (t *ReleaseObjectGroup) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Runtime.releaseObjectGroup", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the ReleaseObjectGroup CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *ReleaseObjectGroup) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Runtime.releaseObjectGroup", b)
}

// ParseResponse parses the browser's response
// to the ReleaseObjectGroup CDP command.
func (t *ReleaseObjectGroup) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// RunIfWaitingForDebugger contains the parameters, and acts as
// a Go receiver, for the CDP command `runIfWaitingForDebugger`.
//
// Tells inspected instance to run if it was waiting for debugger to attach.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-runIfWaitingForDebugger
type RunIfWaitingForDebugger struct{}

// NewRunIfWaitingForDebugger constructs a new RunIfWaitingForDebugger struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-runIfWaitingForDebugger
func NewRunIfWaitingForDebugger() *RunIfWaitingForDebugger {
	return &RunIfWaitingForDebugger{}
}

// Do sends the RunIfWaitingForDebugger CDP command to a browser,
// and returns the browser's response.
func (t *RunIfWaitingForDebugger) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Runtime.runIfWaitingForDebugger", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the RunIfWaitingForDebugger CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *RunIfWaitingForDebugger) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Runtime.runIfWaitingForDebugger", nil)
}

// ParseResponse parses the browser's response
// to the RunIfWaitingForDebugger CDP command.
func (t *RunIfWaitingForDebugger) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// RunScript contains the parameters, and acts as
// a Go receiver, for the CDP command `runScript`.
//
// Runs script with given id in a given context.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-runScript
type RunScript struct {
	// Id of the script to run.
	ScriptID string `json:"scriptId"`
	// Specifies in which execution context to perform script run. If the parameter is omitted the
	// evaluation will be performed in the context of the inspected page.
	ExecutionContextID int64 `json:"executionContextId,omitempty"`
	// Symbolic group name that can be used to release multiple objects.
	ObjectGroup string `json:"objectGroup,omitempty"`
	// In silent mode exceptions thrown during evaluation are not reported and do not pause
	// execution. Overrides `setPauseOnException` state.
	Silent bool `json:"silent,omitempty"`
	// Determines whether Command Line API should be available during the evaluation.
	IncludeCommandLineAPI bool `json:"includeCommandLineAPI,omitempty"`
	// Whether the result is expected to be a JSON object which should be sent by value.
	ReturnByValue bool `json:"returnByValue,omitempty"`
	// Whether preview should be generated for the result.
	GeneratePreview bool `json:"generatePreview,omitempty"`
	// Whether execution should `await` for resulting value and return once awaited promise is
	// resolved.
	AwaitPromise bool `json:"awaitPromise,omitempty"`
}

// NewRunScript constructs a new RunScript struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-runScript
func NewRunScript(scriptID string) *RunScript {
	return &RunScript{
		ScriptID: scriptID,
	}
}

// SetExecutionContextID adds or modifies the value of the optional
// parameter `executionContextId` in the RunScript CDP command.
//
// Specifies in which execution context to perform script run. If the parameter is omitted the
// evaluation will be performed in the context of the inspected page.
func (t *RunScript) SetExecutionContextID(v int64) *RunScript {
	t.ExecutionContextID = v
	return t
}

// SetObjectGroup adds or modifies the value of the optional
// parameter `objectGroup` in the RunScript CDP command.
//
// Symbolic group name that can be used to release multiple objects.
func (t *RunScript) SetObjectGroup(v string) *RunScript {
	t.ObjectGroup = v
	return t
}

// SetSilent adds or modifies the value of the optional
// parameter `silent` in the RunScript CDP command.
//
// In silent mode exceptions thrown during evaluation are not reported and do not pause
// execution. Overrides `setPauseOnException` state.
func (t *RunScript) SetSilent(v bool) *RunScript {
	t.Silent = v
	return t
}

// SetIncludeCommandLineAPI adds or modifies the value of the optional
// parameter `includeCommandLineAPI` in the RunScript CDP command.
//
// Determines whether Command Line API should be available during the evaluation.
func (t *RunScript) SetIncludeCommandLineAPI(v bool) *RunScript {
	t.IncludeCommandLineAPI = v
	return t
}

// SetReturnByValue adds or modifies the value of the optional
// parameter `returnByValue` in the RunScript CDP command.
//
// Whether the result is expected to be a JSON object which should be sent by value.
func (t *RunScript) SetReturnByValue(v bool) *RunScript {
	t.ReturnByValue = v
	return t
}

// SetGeneratePreview adds or modifies the value of the optional
// parameter `generatePreview` in the RunScript CDP command.
//
// Whether preview should be generated for the result.
func (t *RunScript) SetGeneratePreview(v bool) *RunScript {
	t.GeneratePreview = v
	return t
}

// SetAwaitPromise adds or modifies the value of the optional
// parameter `awaitPromise` in the RunScript CDP command.
//
// Whether execution should `await` for resulting value and return once awaited promise is
// resolved.
func (t *RunScript) SetAwaitPromise(v bool) *RunScript {
	t.AwaitPromise = v
	return t
}

// RunScriptResult contains the browser's response
// to calling the RunScript CDP command with Do().
type RunScriptResult struct {
	// Run result.
	Result RemoteObject `json:"result"`
	// Exception details.
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"`
}

// Do sends the RunScript CDP command to a browser,
// and returns the browser's response.
func (t *RunScript) Do(ctx context.Context) (*RunScriptResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Runtime.runScript", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the RunScript CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *RunScript) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Runtime.runScript", b)
}

// ParseResponse parses the browser's response
// to the RunScript CDP command.
func (t *RunScript) ParseResponse(m *devtools.Message) (*RunScriptResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &RunScriptResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetCustomObjectFormatterEnabled contains the parameters, and acts as
// a Go receiver, for the CDP command `setCustomObjectFormatterEnabled`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-setCustomObjectFormatterEnabled
//
// This CDP method is experimental.
type SetCustomObjectFormatterEnabled struct {
	Enabled bool `json:"enabled"`
}

// NewSetCustomObjectFormatterEnabled constructs a new SetCustomObjectFormatterEnabled struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-setCustomObjectFormatterEnabled
//
// This CDP method is experimental.
func NewSetCustomObjectFormatterEnabled(enabled bool) *SetCustomObjectFormatterEnabled {
	return &SetCustomObjectFormatterEnabled{
		Enabled: enabled,
	}
}

// Do sends the SetCustomObjectFormatterEnabled CDP command to a browser,
// and returns the browser's response.
func (t *SetCustomObjectFormatterEnabled) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Runtime.setCustomObjectFormatterEnabled", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetCustomObjectFormatterEnabled CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetCustomObjectFormatterEnabled) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Runtime.setCustomObjectFormatterEnabled", b)
}

// ParseResponse parses the browser's response
// to the SetCustomObjectFormatterEnabled CDP command.
func (t *SetCustomObjectFormatterEnabled) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetMaxCallStackSizeToCapture contains the parameters, and acts as
// a Go receiver, for the CDP command `setMaxCallStackSizeToCapture`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-setMaxCallStackSizeToCapture
//
// This CDP method is experimental.
type SetMaxCallStackSizeToCapture struct {
	Size int64 `json:"size"`
}

// NewSetMaxCallStackSizeToCapture constructs a new SetMaxCallStackSizeToCapture struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-setMaxCallStackSizeToCapture
//
// This CDP method is experimental.
func NewSetMaxCallStackSizeToCapture(size int64) *SetMaxCallStackSizeToCapture {
	return &SetMaxCallStackSizeToCapture{
		Size: size,
	}
}

// Do sends the SetMaxCallStackSizeToCapture CDP command to a browser,
// and returns the browser's response.
func (t *SetMaxCallStackSizeToCapture) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Runtime.setMaxCallStackSizeToCapture", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetMaxCallStackSizeToCapture CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetMaxCallStackSizeToCapture) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Runtime.setMaxCallStackSizeToCapture", b)
}

// ParseResponse parses the browser's response
// to the SetMaxCallStackSizeToCapture CDP command.
func (t *SetMaxCallStackSizeToCapture) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// TerminateExecution contains the parameters, and acts as
// a Go receiver, for the CDP command `terminateExecution`.
//
// Terminate current or next JavaScript execution.
// Will cancel the termination when the outer-most script execution ends.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-terminateExecution
//
// This CDP method is experimental.
type TerminateExecution struct{}

// NewTerminateExecution constructs a new TerminateExecution struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-terminateExecution
//
// This CDP method is experimental.
func NewTerminateExecution() *TerminateExecution {
	return &TerminateExecution{}
}

// Do sends the TerminateExecution CDP command to a browser,
// and returns the browser's response.
func (t *TerminateExecution) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Runtime.terminateExecution", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the TerminateExecution CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *TerminateExecution) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Runtime.terminateExecution", nil)
}

// ParseResponse parses the browser's response
// to the TerminateExecution CDP command.
func (t *TerminateExecution) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// AddBinding contains the parameters, and acts as
// a Go receiver, for the CDP command `addBinding`.
//
// If executionContextId is empty, adds binding with the given name on the
// global objects of all inspected contexts, including those created later,
// bindings survive reloads.
// Binding function takes exactly one argument, this argument should be string,
// in case of any other input, function throws an exception.
// Each binding function call produces Runtime.bindingCalled notification.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-addBinding
//
// This CDP method is experimental.
type AddBinding struct {
	Name string `json:"name"`
	// If specified, the binding would only be exposed to the specified
	// execution context. If omitted and `executionContextName` is not set,
	// the binding is exposed to all execution contexts of the target.
	// This parameter is mutually exclusive with `executionContextName`.
	// Deprecated in favor of `executionContextName` due to an unclear use case
	// and bugs in implementation (crbug.com/1169639). `executionContextId` will be
	// removed in the future.
	//
	// This CDP parameter is deprecated.
	ExecutionContextID int64 `json:"executionContextId,omitempty"`
	// If specified, the binding is exposed to the executionContext with
	// matching name, even for contexts created after the binding is added.
	// See also `ExecutionContext.name` and `worldName` parameter to
	// `Page.addScriptToEvaluateOnNewDocument`.
	// This parameter is mutually exclusive with `executionContextId`.
	//
	// This CDP parameter is experimental.
	ExecutionContextName string `json:"executionContextName,omitempty"`
}

// NewAddBinding constructs a new AddBinding struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-addBinding
//
// This CDP method is experimental.
func NewAddBinding(name string) *AddBinding {
	return &AddBinding{
		Name: name,
	}
}

// SetExecutionContextID adds or modifies the value of the optional
// parameter `executionContextId` in the AddBinding CDP command.
//
// If specified, the binding would only be exposed to the specified
// execution context. If omitted and `executionContextName` is not set,
// the binding is exposed to all execution contexts of the target.
// This parameter is mutually exclusive with `executionContextName`.
// Deprecated in favor of `executionContextName` due to an unclear use case
// and bugs in implementation (crbug.com/1169639). `executionContextId` will be
// removed in the future.
//
// This CDP parameter is deprecated.
func (t *AddBinding) SetExecutionContextID(v int64) *AddBinding {
	t.ExecutionContextID = v
	return t
}

// SetExecutionContextName adds or modifies the value of the optional
// parameter `executionContextName` in the AddBinding CDP command.
//
// If specified, the binding is exposed to the executionContext with
// matching name, even for contexts created after the binding is added.
// See also `ExecutionContext.name` and `worldName` parameter to
// `Page.addScriptToEvaluateOnNewDocument`.
// This parameter is mutually exclusive with `executionContextId`.
//
// This CDP parameter is experimental.
func (t *AddBinding) SetExecutionContextName(v string) *AddBinding {
	t.ExecutionContextName = v
	return t
}

// Do sends the AddBinding CDP command to a browser,
// and returns the browser's response.
func (t *AddBinding) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Runtime.addBinding", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the AddBinding CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *AddBinding) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Runtime.addBinding", b)
}

// ParseResponse parses the browser's response
// to the AddBinding CDP command.
func (t *AddBinding) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// RemoveBinding contains the parameters, and acts as
// a Go receiver, for the CDP command `removeBinding`.
//
// This method does not remove binding function from global object but
// unsubscribes current runtime agent from Runtime.bindingCalled notifications.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-removeBinding
//
// This CDP method is experimental.
type RemoveBinding struct {
	Name string `json:"name"`
}

// NewRemoveBinding constructs a new RemoveBinding struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-removeBinding
//
// This CDP method is experimental.
func NewRemoveBinding(name string) *RemoveBinding {
	return &RemoveBinding{
		Name: name,
	}
}

// Do sends the RemoveBinding CDP command to a browser,
// and returns the browser's response.
func (t *RemoveBinding) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Runtime.removeBinding", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the RemoveBinding CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *RemoveBinding) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Runtime.removeBinding", b)
}

// ParseResponse parses the browser's response
// to the RemoveBinding CDP command.
func (t *RemoveBinding) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}
