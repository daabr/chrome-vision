package debugger

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
	"github.com/daabr/chrome-vision/pkg/devtools/runtime"
)

// ContinueToLocation contains the parameters, and acts as
// a Go receiver, for the CDP command `continueToLocation`.
//
// Continues execution until specific location is reached.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-continueToLocation
type ContinueToLocation struct {
	// Location to continue to.
	Location         Location `json:"location"`
	TargetCallFrames string   `json:"targetCallFrames,omitempty"`
}

// NewContinueToLocation constructs a new ContinueToLocation struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-continueToLocation
func NewContinueToLocation(location Location) *ContinueToLocation {
	return &ContinueToLocation{
		Location: location,
	}
}

// SetTargetCallFrames adds or modifies the value of the optional
// parameter `targetCallFrames` in the ContinueToLocation CDP command.
func (t *ContinueToLocation) SetTargetCallFrames(v string) *ContinueToLocation {
	t.TargetCallFrames = v
	return t
}

// Do sends the ContinueToLocation CDP command to a browser,
// and returns the browser's response.
func (t *ContinueToLocation) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Debugger.continueToLocation", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the ContinueToLocation CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *ContinueToLocation) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Debugger.continueToLocation", b)
}

// ParseResponse parses the browser's response
// to the ContinueToLocation CDP command.
func (t *ContinueToLocation) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// Disable contains the parameters, and acts as
// a Go receiver, for the CDP command `disable`.
//
// Disables debugger for given page.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-disable
type Disable struct{}

// NewDisable constructs a new Disable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-disable
func NewDisable() *Disable {
	return &Disable{}
}

// Do sends the Disable CDP command to a browser,
// and returns the browser's response.
func (t *Disable) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Debugger.disable", nil)
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
	return devtools.Send(ctx, "Debugger.disable", nil)
}

// ParseResponse parses the browser's response
// to the Disable CDP command.
func (t *Disable) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// Enable contains the parameters, and acts as
// a Go receiver, for the CDP command `enable`.
//
// Enables debugger for the given page. Clients should not assume that the debugging has been
// enabled until the result for this command is received.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-enable
type Enable struct {
	// The maximum size in bytes of collected scripts (not referenced by other heap objects)
	// the debugger can hold. Puts no limit if parameter is omitted.
	//
	// This CDP parameter is experimental.
	MaxScriptsCacheSize float64 `json:"maxScriptsCacheSize,omitempty"`
}

// NewEnable constructs a new Enable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-enable
func NewEnable() *Enable {
	return &Enable{}
}

// SetMaxScriptsCacheSize adds or modifies the value of the optional
// parameter `maxScriptsCacheSize` in the Enable CDP command.
//
// The maximum size in bytes of collected scripts (not referenced by other heap objects)
// the debugger can hold. Puts no limit if parameter is omitted.
//
// This CDP parameter is experimental.
func (t *Enable) SetMaxScriptsCacheSize(v float64) *Enable {
	t.MaxScriptsCacheSize = v
	return t
}

// EnableResult contains the browser's response
// to calling the Enable CDP command with Do().
type EnableResult struct {
	// Unique identifier of the debugger.
	//
	// This CDP parameter is experimental.
	DebuggerID string `json:"debuggerId"`
}

// Do sends the Enable CDP command to a browser,
// and returns the browser's response.
func (t *Enable) Do(ctx context.Context) (*EnableResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Debugger.enable", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the Enable CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *Enable) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Debugger.enable", b)
}

// ParseResponse parses the browser's response
// to the Enable CDP command.
func (t *Enable) ParseResponse(m *devtools.Message) (*EnableResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &EnableResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// EvaluateOnCallFrame contains the parameters, and acts as
// a Go receiver, for the CDP command `evaluateOnCallFrame`.
//
// Evaluates expression on a given call frame.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-evaluateOnCallFrame
type EvaluateOnCallFrame struct {
	// Call frame identifier to evaluate on.
	CallFrameID string `json:"callFrameId"`
	// Expression to evaluate.
	Expression string `json:"expression"`
	// String object group name to put result into (allows rapid releasing resulting object handles
	// using `releaseObjectGroup`).
	ObjectGroup string `json:"objectGroup,omitempty"`
	// Specifies whether command line API should be available to the evaluated expression, defaults
	// to false.
	IncludeCommandLineAPI bool `json:"includeCommandLineAPI,omitempty"`
	// In silent mode exceptions thrown during evaluation are not reported and do not pause
	// execution. Overrides `setPauseOnException` state.
	Silent bool `json:"silent,omitempty"`
	// Whether the result is expected to be a JSON object that should be sent by value.
	ReturnByValue bool `json:"returnByValue,omitempty"`
	// Whether preview should be generated for the result.
	//
	// This CDP parameter is experimental.
	GeneratePreview bool `json:"generatePreview,omitempty"`
	// Whether to throw an exception if side effect cannot be ruled out during evaluation.
	ThrowOnSideEffect bool `json:"throwOnSideEffect,omitempty"`
	// Terminate execution after timing out (number of milliseconds).
	//
	// This CDP parameter is experimental.
	Timeout float64 `json:"timeout,omitempty"`
}

// NewEvaluateOnCallFrame constructs a new EvaluateOnCallFrame struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-evaluateOnCallFrame
func NewEvaluateOnCallFrame(callFrameID string, expression string) *EvaluateOnCallFrame {
	return &EvaluateOnCallFrame{
		CallFrameID: callFrameID,
		Expression:  expression,
	}
}

// SetObjectGroup adds or modifies the value of the optional
// parameter `objectGroup` in the EvaluateOnCallFrame CDP command.
//
// String object group name to put result into (allows rapid releasing resulting object handles
// using `releaseObjectGroup`).
func (t *EvaluateOnCallFrame) SetObjectGroup(v string) *EvaluateOnCallFrame {
	t.ObjectGroup = v
	return t
}

// SetIncludeCommandLineAPI adds or modifies the value of the optional
// parameter `includeCommandLineAPI` in the EvaluateOnCallFrame CDP command.
//
// Specifies whether command line API should be available to the evaluated expression, defaults
// to false.
func (t *EvaluateOnCallFrame) SetIncludeCommandLineAPI(v bool) *EvaluateOnCallFrame {
	t.IncludeCommandLineAPI = v
	return t
}

// SetSilent adds or modifies the value of the optional
// parameter `silent` in the EvaluateOnCallFrame CDP command.
//
// In silent mode exceptions thrown during evaluation are not reported and do not pause
// execution. Overrides `setPauseOnException` state.
func (t *EvaluateOnCallFrame) SetSilent(v bool) *EvaluateOnCallFrame {
	t.Silent = v
	return t
}

// SetReturnByValue adds or modifies the value of the optional
// parameter `returnByValue` in the EvaluateOnCallFrame CDP command.
//
// Whether the result is expected to be a JSON object that should be sent by value.
func (t *EvaluateOnCallFrame) SetReturnByValue(v bool) *EvaluateOnCallFrame {
	t.ReturnByValue = v
	return t
}

// SetGeneratePreview adds or modifies the value of the optional
// parameter `generatePreview` in the EvaluateOnCallFrame CDP command.
//
// Whether preview should be generated for the result.
//
// This CDP parameter is experimental.
func (t *EvaluateOnCallFrame) SetGeneratePreview(v bool) *EvaluateOnCallFrame {
	t.GeneratePreview = v
	return t
}

// SetThrowOnSideEffect adds or modifies the value of the optional
// parameter `throwOnSideEffect` in the EvaluateOnCallFrame CDP command.
//
// Whether to throw an exception if side effect cannot be ruled out during evaluation.
func (t *EvaluateOnCallFrame) SetThrowOnSideEffect(v bool) *EvaluateOnCallFrame {
	t.ThrowOnSideEffect = v
	return t
}

// SetTimeout adds or modifies the value of the optional
// parameter `timeout` in the EvaluateOnCallFrame CDP command.
//
// Terminate execution after timing out (number of milliseconds).
//
// This CDP parameter is experimental.
func (t *EvaluateOnCallFrame) SetTimeout(v float64) *EvaluateOnCallFrame {
	t.Timeout = v
	return t
}

// EvaluateOnCallFrameResult contains the browser's response
// to calling the EvaluateOnCallFrame CDP command with Do().
type EvaluateOnCallFrameResult struct {
	// Object wrapper for the evaluation result.
	Result runtime.RemoteObject `json:"result"`
	// Exception details.
	ExceptionDetails *runtime.ExceptionDetails `json:"exceptionDetails,omitempty"`
}

// Do sends the EvaluateOnCallFrame CDP command to a browser,
// and returns the browser's response.
func (t *EvaluateOnCallFrame) Do(ctx context.Context) (*EvaluateOnCallFrameResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Debugger.evaluateOnCallFrame", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the EvaluateOnCallFrame CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *EvaluateOnCallFrame) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Debugger.evaluateOnCallFrame", b)
}

// ParseResponse parses the browser's response
// to the EvaluateOnCallFrame CDP command.
func (t *EvaluateOnCallFrame) ParseResponse(m *devtools.Message) (*EvaluateOnCallFrameResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &EvaluateOnCallFrameResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetPossibleBreakpoints contains the parameters, and acts as
// a Go receiver, for the CDP command `getPossibleBreakpoints`.
//
// Returns possible locations for breakpoint. scriptId in start and end range locations should be
// the same.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-getPossibleBreakpoints
type GetPossibleBreakpoints struct {
	// Start of range to search possible breakpoint locations in.
	Start Location `json:"start"`
	// End of range to search possible breakpoint locations in (excluding). When not specified, end
	// of scripts is used as end of range.
	End *Location `json:"end,omitempty"`
	// Only consider locations which are in the same (non-nested) function as start.
	RestrictToFunction bool `json:"restrictToFunction,omitempty"`
}

// NewGetPossibleBreakpoints constructs a new GetPossibleBreakpoints struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-getPossibleBreakpoints
func NewGetPossibleBreakpoints(start Location) *GetPossibleBreakpoints {
	return &GetPossibleBreakpoints{
		Start: start,
	}
}

// SetEnd adds or modifies the value of the optional
// parameter `end` in the GetPossibleBreakpoints CDP command.
//
// End of range to search possible breakpoint locations in (excluding). When not specified, end
// of scripts is used as end of range.
func (t *GetPossibleBreakpoints) SetEnd(v Location) *GetPossibleBreakpoints {
	t.End = &v
	return t
}

// SetRestrictToFunction adds or modifies the value of the optional
// parameter `restrictToFunction` in the GetPossibleBreakpoints CDP command.
//
// Only consider locations which are in the same (non-nested) function as start.
func (t *GetPossibleBreakpoints) SetRestrictToFunction(v bool) *GetPossibleBreakpoints {
	t.RestrictToFunction = v
	return t
}

// GetPossibleBreakpointsResult contains the browser's response
// to calling the GetPossibleBreakpoints CDP command with Do().
type GetPossibleBreakpointsResult struct {
	// List of the possible breakpoint locations.
	Locations []BreakLocation `json:"locations"`
}

// Do sends the GetPossibleBreakpoints CDP command to a browser,
// and returns the browser's response.
func (t *GetPossibleBreakpoints) Do(ctx context.Context) (*GetPossibleBreakpointsResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Debugger.getPossibleBreakpoints", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// ParseResponse parses the browser's response
// to the GetPossibleBreakpoints CDP command.
func (t *GetPossibleBreakpoints) ParseResponse(m *devtools.Message) (*GetPossibleBreakpointsResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetPossibleBreakpointsResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetScriptSource contains the parameters, and acts as
// a Go receiver, for the CDP command `getScriptSource`.
//
// Returns source for the script with given id.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-getScriptSource
type GetScriptSource struct {
	// Id of the script to get source for.
	ScriptID string `json:"scriptId"`
}

// NewGetScriptSource constructs a new GetScriptSource struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-getScriptSource
func NewGetScriptSource(scriptID string) *GetScriptSource {
	return &GetScriptSource{
		ScriptID: scriptID,
	}
}

// GetScriptSourceResult contains the browser's response
// to calling the GetScriptSource CDP command with Do().
type GetScriptSourceResult struct {
	// Script source (empty in case of Wasm bytecode).
	ScriptSource string `json:"scriptSource"`
	// Wasm bytecode. (Encoded as a base64 string when passed over JSON)
	Bytecode string `json:"bytecode,omitempty"`
}

// Do sends the GetScriptSource CDP command to a browser,
// and returns the browser's response.
func (t *GetScriptSource) Do(ctx context.Context) (*GetScriptSourceResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Debugger.getScriptSource", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetScriptSource CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetScriptSource) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Debugger.getScriptSource", b)
}

// ParseResponse parses the browser's response
// to the GetScriptSource CDP command.
func (t *GetScriptSource) ParseResponse(m *devtools.Message) (*GetScriptSourceResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetScriptSourceResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetWasmBytecode contains the parameters, and acts as
// a Go receiver, for the CDP command `getWasmBytecode`.
//
// This command is deprecated. Use getScriptSource instead.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-getWasmBytecode
//
// This CDP method is deprecated.
type GetWasmBytecode struct {
	// Id of the Wasm script to get source for.
	ScriptID string `json:"scriptId"`
}

// NewGetWasmBytecode constructs a new GetWasmBytecode struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-getWasmBytecode
//
// This CDP method is deprecated.
func NewGetWasmBytecode(scriptID string) *GetWasmBytecode {
	return &GetWasmBytecode{
		ScriptID: scriptID,
	}
}

// GetWasmBytecodeResult contains the browser's response
// to calling the GetWasmBytecode CDP command with Do().
type GetWasmBytecodeResult struct {
	// Script source. (Encoded as a base64 string when passed over JSON)
	Bytecode string `json:"bytecode"`
}

// Do sends the GetWasmBytecode CDP command to a browser,
// and returns the browser's response.
func (t *GetWasmBytecode) Do(ctx context.Context) (*GetWasmBytecodeResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Debugger.getWasmBytecode", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetWasmBytecode CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetWasmBytecode) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Debugger.getWasmBytecode", b)
}

// ParseResponse parses the browser's response
// to the GetWasmBytecode CDP command.
func (t *GetWasmBytecode) ParseResponse(m *devtools.Message) (*GetWasmBytecodeResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetWasmBytecodeResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetStackTrace contains the parameters, and acts as
// a Go receiver, for the CDP command `getStackTrace`.
//
// Returns stack trace with given `stackTraceId`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-getStackTrace
//
// This CDP method is experimental.
type GetStackTrace struct {
	StackTraceID runtime.StackTraceID `json:"stackTraceId"`
}

// NewGetStackTrace constructs a new GetStackTrace struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-getStackTrace
//
// This CDP method is experimental.
func NewGetStackTrace(stackTraceID runtime.StackTraceID) *GetStackTrace {
	return &GetStackTrace{
		StackTraceID: stackTraceID,
	}
}

// GetStackTraceResult contains the browser's response
// to calling the GetStackTrace CDP command with Do().
type GetStackTraceResult struct {
	StackTrace runtime.StackTrace `json:"stackTrace"`
}

// Do sends the GetStackTrace CDP command to a browser,
// and returns the browser's response.
func (t *GetStackTrace) Do(ctx context.Context) (*GetStackTraceResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Debugger.getStackTrace", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetStackTrace CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetStackTrace) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Debugger.getStackTrace", b)
}

// ParseResponse parses the browser's response
// to the GetStackTrace CDP command.
func (t *GetStackTrace) ParseResponse(m *devtools.Message) (*GetStackTraceResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetStackTraceResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Pause contains the parameters, and acts as
// a Go receiver, for the CDP command `pause`.
//
// Stops on the next JavaScript statement.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-pause
type Pause struct{}

// NewPause constructs a new Pause struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-pause
func NewPause() *Pause {
	return &Pause{}
}

// Do sends the Pause CDP command to a browser,
// and returns the browser's response.
func (t *Pause) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Debugger.pause", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the Pause CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *Pause) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Debugger.pause", nil)
}

// ParseResponse parses the browser's response
// to the Pause CDP command.
func (t *Pause) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// PauseOnAsyncCall contains the parameters, and acts as
// a Go receiver, for the CDP command `pauseOnAsyncCall`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-pauseOnAsyncCall
//
// This CDP method is deprecated.
// This CDP method is experimental.
type PauseOnAsyncCall struct {
	// Debugger will pause when async call with given stack trace is started.
	ParentStackTraceID runtime.StackTraceID `json:"parentStackTraceId"`
}

// NewPauseOnAsyncCall constructs a new PauseOnAsyncCall struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-pauseOnAsyncCall
//
// This CDP method is deprecated.
// This CDP method is experimental.
func NewPauseOnAsyncCall(parentStackTraceID runtime.StackTraceID) *PauseOnAsyncCall {
	return &PauseOnAsyncCall{
		ParentStackTraceID: parentStackTraceID,
	}
}

// Do sends the PauseOnAsyncCall CDP command to a browser,
// and returns the browser's response.
func (t *PauseOnAsyncCall) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Debugger.pauseOnAsyncCall", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the PauseOnAsyncCall CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *PauseOnAsyncCall) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Debugger.pauseOnAsyncCall", b)
}

// ParseResponse parses the browser's response
// to the PauseOnAsyncCall CDP command.
func (t *PauseOnAsyncCall) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// RemoveBreakpoint contains the parameters, and acts as
// a Go receiver, for the CDP command `removeBreakpoint`.
//
// Removes JavaScript breakpoint.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-removeBreakpoint
type RemoveBreakpoint struct {
	BreakpointID string `json:"breakpointId"`
}

// NewRemoveBreakpoint constructs a new RemoveBreakpoint struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-removeBreakpoint
func NewRemoveBreakpoint(breakpointID string) *RemoveBreakpoint {
	return &RemoveBreakpoint{
		BreakpointID: breakpointID,
	}
}

// Do sends the RemoveBreakpoint CDP command to a browser,
// and returns the browser's response.
func (t *RemoveBreakpoint) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Debugger.removeBreakpoint", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the RemoveBreakpoint CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *RemoveBreakpoint) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Debugger.removeBreakpoint", b)
}

// ParseResponse parses the browser's response
// to the RemoveBreakpoint CDP command.
func (t *RemoveBreakpoint) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// RestartFrame contains the parameters, and acts as
// a Go receiver, for the CDP command `restartFrame`.
//
// Restarts particular call frame from the beginning.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-restartFrame
//
// This CDP method is deprecated.
type RestartFrame struct {
	// Call frame identifier to evaluate on.
	CallFrameID string `json:"callFrameId"`
}

// NewRestartFrame constructs a new RestartFrame struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-restartFrame
//
// This CDP method is deprecated.
func NewRestartFrame(callFrameID string) *RestartFrame {
	return &RestartFrame{
		CallFrameID: callFrameID,
	}
}

// RestartFrameResult contains the browser's response
// to calling the RestartFrame CDP command with Do().
type RestartFrameResult struct {
	// New stack trace.
	CallFrames []CallFrame `json:"callFrames"`
	// Async stack trace, if any.
	AsyncStackTrace *runtime.StackTrace `json:"asyncStackTrace,omitempty"`
	// Async stack trace, if any.
	//
	// This CDP parameter is experimental.
	AsyncStackTraceID *runtime.StackTraceID `json:"asyncStackTraceId,omitempty"`
}

// Do sends the RestartFrame CDP command to a browser,
// and returns the browser's response.
func (t *RestartFrame) Do(ctx context.Context) (*RestartFrameResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Debugger.restartFrame", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the RestartFrame CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *RestartFrame) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Debugger.restartFrame", b)
}

// ParseResponse parses the browser's response
// to the RestartFrame CDP command.
func (t *RestartFrame) ParseResponse(m *devtools.Message) (*RestartFrameResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &RestartFrameResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Resume contains the parameters, and acts as
// a Go receiver, for the CDP command `resume`.
//
// Resumes JavaScript execution.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-resume
type Resume struct {
	// Set to true to terminate execution upon resuming execution. In contrast
	// to Runtime.terminateExecution, this will allows to execute further
	// JavaScript (i.e. via evaluation) until execution of the paused code
	// is actually resumed, at which point termination is triggered.
	// If execution is currently not paused, this parameter has no effect.
	TerminateOnResume bool `json:"terminateOnResume,omitempty"`
}

// NewResume constructs a new Resume struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-resume
func NewResume() *Resume {
	return &Resume{}
}

// SetTerminateOnResume adds or modifies the value of the optional
// parameter `terminateOnResume` in the Resume CDP command.
//
// Set to true to terminate execution upon resuming execution. In contrast
// to Runtime.terminateExecution, this will allows to execute further
// JavaScript (i.e. via evaluation) until execution of the paused code
// is actually resumed, at which point termination is triggered.
// If execution is currently not paused, this parameter has no effect.
func (t *Resume) SetTerminateOnResume(v bool) *Resume {
	t.TerminateOnResume = v
	return t
}

// Do sends the Resume CDP command to a browser,
// and returns the browser's response.
func (t *Resume) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Debugger.resume", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the Resume CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *Resume) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Debugger.resume", b)
}

// ParseResponse parses the browser's response
// to the Resume CDP command.
func (t *Resume) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SearchInContent contains the parameters, and acts as
// a Go receiver, for the CDP command `searchInContent`.
//
// Searches for given string in script content.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-searchInContent
type SearchInContent struct {
	// Id of the script to search in.
	ScriptID string `json:"scriptId"`
	// String to search for.
	Query string `json:"query"`
	// If true, search is case sensitive.
	CaseSensitive bool `json:"caseSensitive,omitempty"`
	// If true, treats string parameter as regex.
	IsRegex bool `json:"isRegex,omitempty"`
}

// NewSearchInContent constructs a new SearchInContent struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-searchInContent
func NewSearchInContent(scriptID string, query string) *SearchInContent {
	return &SearchInContent{
		ScriptID: scriptID,
		Query:    query,
	}
}

// SetCaseSensitive adds or modifies the value of the optional
// parameter `caseSensitive` in the SearchInContent CDP command.
//
// If true, search is case sensitive.
func (t *SearchInContent) SetCaseSensitive(v bool) *SearchInContent {
	t.CaseSensitive = v
	return t
}

// SetIsRegex adds or modifies the value of the optional
// parameter `isRegex` in the SearchInContent CDP command.
//
// If true, treats string parameter as regex.
func (t *SearchInContent) SetIsRegex(v bool) *SearchInContent {
	t.IsRegex = v
	return t
}

// SearchInContentResult contains the browser's response
// to calling the SearchInContent CDP command with Do().
type SearchInContentResult struct {
	// List of search matches.
	Result []SearchMatch `json:"result"`
}

// Do sends the SearchInContent CDP command to a browser,
// and returns the browser's response.
func (t *SearchInContent) Do(ctx context.Context) (*SearchInContentResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Debugger.searchInContent", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the SearchInContent CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SearchInContent) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Debugger.searchInContent", b)
}

// ParseResponse parses the browser's response
// to the SearchInContent CDP command.
func (t *SearchInContent) ParseResponse(m *devtools.Message) (*SearchInContentResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &SearchInContentResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetAsyncCallStackDepth contains the parameters, and acts as
// a Go receiver, for the CDP command `setAsyncCallStackDepth`.
//
// Enables or disables async call stacks tracking.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setAsyncCallStackDepth
type SetAsyncCallStackDepth struct {
	// Maximum depth of async call stacks. Setting to `0` will effectively disable collecting async
	// call stacks (default).
	MaxDepth int64 `json:"maxDepth"`
}

// NewSetAsyncCallStackDepth constructs a new SetAsyncCallStackDepth struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setAsyncCallStackDepth
func NewSetAsyncCallStackDepth(maxDepth int64) *SetAsyncCallStackDepth {
	return &SetAsyncCallStackDepth{
		MaxDepth: maxDepth,
	}
}

// Do sends the SetAsyncCallStackDepth CDP command to a browser,
// and returns the browser's response.
func (t *SetAsyncCallStackDepth) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Debugger.setAsyncCallStackDepth", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetAsyncCallStackDepth CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetAsyncCallStackDepth) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Debugger.setAsyncCallStackDepth", b)
}

// ParseResponse parses the browser's response
// to the SetAsyncCallStackDepth CDP command.
func (t *SetAsyncCallStackDepth) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetBlackboxPatterns contains the parameters, and acts as
// a Go receiver, for the CDP command `setBlackboxPatterns`.
//
// Replace previous blackbox patterns with passed ones. Forces backend to skip stepping/pausing in
// scripts with url matching one of the patterns. VM will try to leave blackboxed script by
// performing 'step in' several times, finally resorting to 'step out' if unsuccessful.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBlackboxPatterns
//
// This CDP method is experimental.
type SetBlackboxPatterns struct {
	// Array of regexps that will be used to check script url for blackbox state.
	Patterns []string `json:"patterns"`
}

// NewSetBlackboxPatterns constructs a new SetBlackboxPatterns struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBlackboxPatterns
//
// This CDP method is experimental.
func NewSetBlackboxPatterns(patterns []string) *SetBlackboxPatterns {
	return &SetBlackboxPatterns{
		Patterns: patterns,
	}
}

// Do sends the SetBlackboxPatterns CDP command to a browser,
// and returns the browser's response.
func (t *SetBlackboxPatterns) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Debugger.setBlackboxPatterns", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetBlackboxPatterns CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetBlackboxPatterns) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Debugger.setBlackboxPatterns", b)
}

// ParseResponse parses the browser's response
// to the SetBlackboxPatterns CDP command.
func (t *SetBlackboxPatterns) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetBlackboxedRanges contains the parameters, and acts as
// a Go receiver, for the CDP command `setBlackboxedRanges`.
//
// Makes backend skip steps in the script in blackboxed ranges. VM will try leave blacklisted
// scripts by performing 'step in' several times, finally resorting to 'step out' if unsuccessful.
// Positions array contains positions where blackbox state is changed. First interval isn't
// blackboxed. Array should be sorted.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBlackboxedRanges
//
// This CDP method is experimental.
type SetBlackboxedRanges struct {
	// Id of the script.
	ScriptID  string           `json:"scriptId"`
	Positions []ScriptPosition `json:"positions"`
}

// NewSetBlackboxedRanges constructs a new SetBlackboxedRanges struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBlackboxedRanges
//
// This CDP method is experimental.
func NewSetBlackboxedRanges(scriptID string, positions []ScriptPosition) *SetBlackboxedRanges {
	return &SetBlackboxedRanges{
		ScriptID:  scriptID,
		Positions: positions,
	}
}

// Do sends the SetBlackboxedRanges CDP command to a browser,
// and returns the browser's response.
func (t *SetBlackboxedRanges) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Debugger.setBlackboxedRanges", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetBlackboxedRanges CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetBlackboxedRanges) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Debugger.setBlackboxedRanges", b)
}

// ParseResponse parses the browser's response
// to the SetBlackboxedRanges CDP command.
func (t *SetBlackboxedRanges) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetBreakpoint contains the parameters, and acts as
// a Go receiver, for the CDP command `setBreakpoint`.
//
// Sets JavaScript breakpoint at a given location.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBreakpoint
type SetBreakpoint struct {
	// Location to set breakpoint in.
	Location Location `json:"location"`
	// Expression to use as a breakpoint condition. When specified, debugger will only stop on the
	// breakpoint if this expression evaluates to true.
	Condition string `json:"condition,omitempty"`
}

// NewSetBreakpoint constructs a new SetBreakpoint struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBreakpoint
func NewSetBreakpoint(location Location) *SetBreakpoint {
	return &SetBreakpoint{
		Location: location,
	}
}

// SetCondition adds or modifies the value of the optional
// parameter `condition` in the SetBreakpoint CDP command.
//
// Expression to use as a breakpoint condition. When specified, debugger will only stop on the
// breakpoint if this expression evaluates to true.
func (t *SetBreakpoint) SetCondition(v string) *SetBreakpoint {
	t.Condition = v
	return t
}

// SetBreakpointResult contains the browser's response
// to calling the SetBreakpoint CDP command with Do().
type SetBreakpointResult struct {
	// Id of the created breakpoint for further reference.
	BreakpointID string `json:"breakpointId"`
	// Location this breakpoint resolved into.
	ActualLocation Location `json:"actualLocation"`
}

// Do sends the SetBreakpoint CDP command to a browser,
// and returns the browser's response.
func (t *SetBreakpoint) Do(ctx context.Context) (*SetBreakpointResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Debugger.setBreakpoint", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the SetBreakpoint CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetBreakpoint) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Debugger.setBreakpoint", b)
}

// ParseResponse parses the browser's response
// to the SetBreakpoint CDP command.
func (t *SetBreakpoint) ParseResponse(m *devtools.Message) (*SetBreakpointResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &SetBreakpointResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetInstrumentationBreakpoint contains the parameters, and acts as
// a Go receiver, for the CDP command `setInstrumentationBreakpoint`.
//
// Sets instrumentation breakpoint.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setInstrumentationBreakpoint
type SetInstrumentationBreakpoint struct {
	// Instrumentation name.
	Instrumentation string `json:"instrumentation"`
}

// NewSetInstrumentationBreakpoint constructs a new SetInstrumentationBreakpoint struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setInstrumentationBreakpoint
func NewSetInstrumentationBreakpoint(instrumentation string) *SetInstrumentationBreakpoint {
	return &SetInstrumentationBreakpoint{
		Instrumentation: instrumentation,
	}
}

// SetInstrumentationBreakpointResult contains the browser's response
// to calling the SetInstrumentationBreakpoint CDP command with Do().
type SetInstrumentationBreakpointResult struct {
	// Id of the created breakpoint for further reference.
	BreakpointID string `json:"breakpointId"`
}

// Do sends the SetInstrumentationBreakpoint CDP command to a browser,
// and returns the browser's response.
func (t *SetInstrumentationBreakpoint) Do(ctx context.Context) (*SetInstrumentationBreakpointResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Debugger.setInstrumentationBreakpoint", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the SetInstrumentationBreakpoint CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetInstrumentationBreakpoint) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Debugger.setInstrumentationBreakpoint", b)
}

// ParseResponse parses the browser's response
// to the SetInstrumentationBreakpoint CDP command.
func (t *SetInstrumentationBreakpoint) ParseResponse(m *devtools.Message) (*SetInstrumentationBreakpointResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &SetInstrumentationBreakpointResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetBreakpointByURL contains the parameters, and acts as
// a Go receiver, for the CDP command `setBreakpointByUrl`.
//
// Sets JavaScript breakpoint at given location specified either by URL or URL regex. Once this
// command is issued, all existing parsed scripts will have breakpoints resolved and returned in
// `locations` property. Further matching script parsing will result in subsequent
// `breakpointResolved` events issued. This logical breakpoint will survive page reloads.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBreakpointByUrl
type SetBreakpointByURL struct {
	// Line number to set breakpoint at.
	LineNumber int64 `json:"lineNumber"`
	// URL of the resources to set breakpoint on.
	URL string `json:"url,omitempty"`
	// Regex pattern for the URLs of the resources to set breakpoints on. Either `url` or
	// `urlRegex` must be specified.
	URLRegex string `json:"urlRegex,omitempty"`
	// Script hash of the resources to set breakpoint on.
	ScriptHash string `json:"scriptHash,omitempty"`
	// Offset in the line to set breakpoint at.
	ColumnNumber int64 `json:"columnNumber,omitempty"`
	// Expression to use as a breakpoint condition. When specified, debugger will only stop on the
	// breakpoint if this expression evaluates to true.
	Condition string `json:"condition,omitempty"`
}

// NewSetBreakpointByURL constructs a new SetBreakpointByURL struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBreakpointByUrl
func NewSetBreakpointByURL(lineNumber int64) *SetBreakpointByURL {
	return &SetBreakpointByURL{
		LineNumber: lineNumber,
	}
}

// SetURL adds or modifies the value of the optional
// parameter `url` in the SetBreakpointByURL CDP command.
//
// URL of the resources to set breakpoint on.
func (t *SetBreakpointByURL) SetURL(v string) *SetBreakpointByURL {
	t.URL = v
	return t
}

// SetURLRegex adds or modifies the value of the optional
// parameter `urlRegex` in the SetBreakpointByURL CDP command.
//
// Regex pattern for the URLs of the resources to set breakpoints on. Either `url` or
// `urlRegex` must be specified.
func (t *SetBreakpointByURL) SetURLRegex(v string) *SetBreakpointByURL {
	t.URLRegex = v
	return t
}

// SetScriptHash adds or modifies the value of the optional
// parameter `scriptHash` in the SetBreakpointByURL CDP command.
//
// Script hash of the resources to set breakpoint on.
func (t *SetBreakpointByURL) SetScriptHash(v string) *SetBreakpointByURL {
	t.ScriptHash = v
	return t
}

// SetColumnNumber adds or modifies the value of the optional
// parameter `columnNumber` in the SetBreakpointByURL CDP command.
//
// Offset in the line to set breakpoint at.
func (t *SetBreakpointByURL) SetColumnNumber(v int64) *SetBreakpointByURL {
	t.ColumnNumber = v
	return t
}

// SetCondition adds or modifies the value of the optional
// parameter `condition` in the SetBreakpointByURL CDP command.
//
// Expression to use as a breakpoint condition. When specified, debugger will only stop on the
// breakpoint if this expression evaluates to true.
func (t *SetBreakpointByURL) SetCondition(v string) *SetBreakpointByURL {
	t.Condition = v
	return t
}

// SetBreakpointByURLResult contains the browser's response
// to calling the SetBreakpointByURL CDP command with Do().
type SetBreakpointByURLResult struct {
	// Id of the created breakpoint for further reference.
	BreakpointID string `json:"breakpointId"`
	// List of the locations this breakpoint resolved into upon addition.
	Locations []Location `json:"locations"`
}

// Do sends the SetBreakpointByURL CDP command to a browser,
// and returns the browser's response.
func (t *SetBreakpointByURL) Do(ctx context.Context) (*SetBreakpointByURLResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Debugger.setBreakpointByUrl", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the SetBreakpointByURL CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetBreakpointByURL) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Debugger.setBreakpointByUrl", b)
}

// ParseResponse parses the browser's response
// to the SetBreakpointByURL CDP command.
func (t *SetBreakpointByURL) ParseResponse(m *devtools.Message) (*SetBreakpointByURLResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &SetBreakpointByURLResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetBreakpointOnFunctionCall contains the parameters, and acts as
// a Go receiver, for the CDP command `setBreakpointOnFunctionCall`.
//
// Sets JavaScript breakpoint before each call to the given function.
// If another function was created from the same source as a given one,
// calling it will also trigger the breakpoint.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBreakpointOnFunctionCall
//
// This CDP method is experimental.
type SetBreakpointOnFunctionCall struct {
	// Function object id.
	ObjectID string `json:"objectId"`
	// Expression to use as a breakpoint condition. When specified, debugger will
	// stop on the breakpoint if this expression evaluates to true.
	Condition string `json:"condition,omitempty"`
}

// NewSetBreakpointOnFunctionCall constructs a new SetBreakpointOnFunctionCall struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBreakpointOnFunctionCall
//
// This CDP method is experimental.
func NewSetBreakpointOnFunctionCall(objectID string) *SetBreakpointOnFunctionCall {
	return &SetBreakpointOnFunctionCall{
		ObjectID: objectID,
	}
}

// SetCondition adds or modifies the value of the optional
// parameter `condition` in the SetBreakpointOnFunctionCall CDP command.
//
// Expression to use as a breakpoint condition. When specified, debugger will
// stop on the breakpoint if this expression evaluates to true.
func (t *SetBreakpointOnFunctionCall) SetCondition(v string) *SetBreakpointOnFunctionCall {
	t.Condition = v
	return t
}

// SetBreakpointOnFunctionCallResult contains the browser's response
// to calling the SetBreakpointOnFunctionCall CDP command with Do().
type SetBreakpointOnFunctionCallResult struct {
	// Id of the created breakpoint for further reference.
	BreakpointID string `json:"breakpointId"`
}

// Do sends the SetBreakpointOnFunctionCall CDP command to a browser,
// and returns the browser's response.
func (t *SetBreakpointOnFunctionCall) Do(ctx context.Context) (*SetBreakpointOnFunctionCallResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Debugger.setBreakpointOnFunctionCall", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the SetBreakpointOnFunctionCall CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetBreakpointOnFunctionCall) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Debugger.setBreakpointOnFunctionCall", b)
}

// ParseResponse parses the browser's response
// to the SetBreakpointOnFunctionCall CDP command.
func (t *SetBreakpointOnFunctionCall) ParseResponse(m *devtools.Message) (*SetBreakpointOnFunctionCallResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &SetBreakpointOnFunctionCallResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetBreakpointsActive contains the parameters, and acts as
// a Go receiver, for the CDP command `setBreakpointsActive`.
//
// Activates / deactivates all breakpoints on the page.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBreakpointsActive
type SetBreakpointsActive struct {
	// New value for breakpoints active state.
	Active bool `json:"active"`
}

// NewSetBreakpointsActive constructs a new SetBreakpointsActive struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setBreakpointsActive
func NewSetBreakpointsActive(active bool) *SetBreakpointsActive {
	return &SetBreakpointsActive{
		Active: active,
	}
}

// Do sends the SetBreakpointsActive CDP command to a browser,
// and returns the browser's response.
func (t *SetBreakpointsActive) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Debugger.setBreakpointsActive", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetBreakpointsActive CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetBreakpointsActive) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Debugger.setBreakpointsActive", b)
}

// ParseResponse parses the browser's response
// to the SetBreakpointsActive CDP command.
func (t *SetBreakpointsActive) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetPauseOnExceptions contains the parameters, and acts as
// a Go receiver, for the CDP command `setPauseOnExceptions`.
//
// Defines pause on exceptions state. Can be set to stop on all exceptions, uncaught exceptions or
// no exceptions. Initial pause on exceptions state is `none`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setPauseOnExceptions
type SetPauseOnExceptions struct {
	// Pause on exceptions mode.
	State string `json:"state"`
}

// NewSetPauseOnExceptions constructs a new SetPauseOnExceptions struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setPauseOnExceptions
func NewSetPauseOnExceptions(state string) *SetPauseOnExceptions {
	return &SetPauseOnExceptions{
		State: state,
	}
}

// Do sends the SetPauseOnExceptions CDP command to a browser,
// and returns the browser's response.
func (t *SetPauseOnExceptions) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Debugger.setPauseOnExceptions", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetPauseOnExceptions CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetPauseOnExceptions) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Debugger.setPauseOnExceptions", b)
}

// ParseResponse parses the browser's response
// to the SetPauseOnExceptions CDP command.
func (t *SetPauseOnExceptions) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetReturnValue contains the parameters, and acts as
// a Go receiver, for the CDP command `setReturnValue`.
//
// Changes return value in top frame. Available only at return break position.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setReturnValue
//
// This CDP method is experimental.
type SetReturnValue struct {
	// New return value.
	NewValue runtime.CallArgument `json:"newValue"`
}

// NewSetReturnValue constructs a new SetReturnValue struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setReturnValue
//
// This CDP method is experimental.
func NewSetReturnValue(newValue runtime.CallArgument) *SetReturnValue {
	return &SetReturnValue{
		NewValue: newValue,
	}
}

// Do sends the SetReturnValue CDP command to a browser,
// and returns the browser's response.
func (t *SetReturnValue) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Debugger.setReturnValue", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetReturnValue CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetReturnValue) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Debugger.setReturnValue", b)
}

// ParseResponse parses the browser's response
// to the SetReturnValue CDP command.
func (t *SetReturnValue) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetScriptSource contains the parameters, and acts as
// a Go receiver, for the CDP command `setScriptSource`.
//
// Edits JavaScript source live.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setScriptSource
type SetScriptSource struct {
	// Id of the script to edit.
	ScriptID string `json:"scriptId"`
	// New content of the script.
	ScriptSource string `json:"scriptSource"`
	// If true the change will not actually be applied. Dry run may be used to get result
	// description without actually modifying the code.
	DryRun bool `json:"dryRun,omitempty"`
}

// NewSetScriptSource constructs a new SetScriptSource struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setScriptSource
func NewSetScriptSource(scriptID string, scriptSource string) *SetScriptSource {
	return &SetScriptSource{
		ScriptID:     scriptID,
		ScriptSource: scriptSource,
	}
}

// SetDryRun adds or modifies the value of the optional
// parameter `dryRun` in the SetScriptSource CDP command.
//
// If true the change will not actually be applied. Dry run may be used to get result
// description without actually modifying the code.
func (t *SetScriptSource) SetDryRun(v bool) *SetScriptSource {
	t.DryRun = v
	return t
}

// SetScriptSourceResult contains the browser's response
// to calling the SetScriptSource CDP command with Do().
type SetScriptSourceResult struct {
	// New stack trace in case editing has happened while VM was stopped.
	CallFrames []CallFrame `json:"callFrames,omitempty"`
	// Whether current call stack  was modified after applying the changes.
	StackChanged bool `json:"stackChanged,omitempty"`
	// Async stack trace, if any.
	AsyncStackTrace *runtime.StackTrace `json:"asyncStackTrace,omitempty"`
	// Async stack trace, if any.
	//
	// This CDP parameter is experimental.
	AsyncStackTraceID *runtime.StackTraceID `json:"asyncStackTraceId,omitempty"`
	// Exception details if any.
	ExceptionDetails *runtime.ExceptionDetails `json:"exceptionDetails,omitempty"`
}

// Do sends the SetScriptSource CDP command to a browser,
// and returns the browser's response.
func (t *SetScriptSource) Do(ctx context.Context) (*SetScriptSourceResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Debugger.setScriptSource", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the SetScriptSource CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetScriptSource) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Debugger.setScriptSource", b)
}

// ParseResponse parses the browser's response
// to the SetScriptSource CDP command.
func (t *SetScriptSource) ParseResponse(m *devtools.Message) (*SetScriptSourceResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &SetScriptSourceResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetSkipAllPauses contains the parameters, and acts as
// a Go receiver, for the CDP command `setSkipAllPauses`.
//
// Makes page not interrupt on any pauses (breakpoint, exception, dom exception etc).
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setSkipAllPauses
type SetSkipAllPauses struct {
	// New value for skip pauses state.
	Skip bool `json:"skip"`
}

// NewSetSkipAllPauses constructs a new SetSkipAllPauses struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setSkipAllPauses
func NewSetSkipAllPauses(skip bool) *SetSkipAllPauses {
	return &SetSkipAllPauses{
		Skip: skip,
	}
}

// Do sends the SetSkipAllPauses CDP command to a browser,
// and returns the browser's response.
func (t *SetSkipAllPauses) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Debugger.setSkipAllPauses", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetSkipAllPauses CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetSkipAllPauses) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Debugger.setSkipAllPauses", b)
}

// ParseResponse parses the browser's response
// to the SetSkipAllPauses CDP command.
func (t *SetSkipAllPauses) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetVariableValue contains the parameters, and acts as
// a Go receiver, for the CDP command `setVariableValue`.
//
// Changes value of variable in a callframe. Object-based scopes are not supported and must be
// mutated manually.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setVariableValue
type SetVariableValue struct {
	// 0-based number of scope as was listed in scope chain. Only 'local', 'closure' and 'catch'
	// scope types are allowed. Other scopes could be manipulated manually.
	ScopeNumber int64 `json:"scopeNumber"`
	// Variable name.
	VariableName string `json:"variableName"`
	// New variable value.
	NewValue runtime.CallArgument `json:"newValue"`
	// Id of callframe that holds variable.
	CallFrameID string `json:"callFrameId"`
}

// NewSetVariableValue constructs a new SetVariableValue struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-setVariableValue
func NewSetVariableValue(scopeNumber int64, variableName string, newValue runtime.CallArgument, callFrameID string) *SetVariableValue {
	return &SetVariableValue{
		ScopeNumber:  scopeNumber,
		VariableName: variableName,
		NewValue:     newValue,
		CallFrameID:  callFrameID,
	}
}

// Do sends the SetVariableValue CDP command to a browser,
// and returns the browser's response.
func (t *SetVariableValue) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Debugger.setVariableValue", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetVariableValue CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetVariableValue) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Debugger.setVariableValue", b)
}

// ParseResponse parses the browser's response
// to the SetVariableValue CDP command.
func (t *SetVariableValue) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// StepInto contains the parameters, and acts as
// a Go receiver, for the CDP command `stepInto`.
//
// Steps into the function call.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-stepInto
type StepInto struct {
	// Debugger will pause on the execution of the first async task which was scheduled
	// before next pause.
	//
	// This CDP parameter is experimental.
	BreakOnAsyncCall bool `json:"breakOnAsyncCall,omitempty"`
	// The skipList specifies location ranges that should be skipped on step into.
	//
	// This CDP parameter is experimental.
	SkipList []LocationRange `json:"skipList,omitempty"`
}

// NewStepInto constructs a new StepInto struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-stepInto
func NewStepInto() *StepInto {
	return &StepInto{}
}

// SetBreakOnAsyncCall adds or modifies the value of the optional
// parameter `breakOnAsyncCall` in the StepInto CDP command.
//
// Debugger will pause on the execution of the first async task which was scheduled
// before next pause.
//
// This CDP parameter is experimental.
func (t *StepInto) SetBreakOnAsyncCall(v bool) *StepInto {
	t.BreakOnAsyncCall = v
	return t
}

// SetSkipList adds or modifies the value of the optional
// parameter `skipList` in the StepInto CDP command.
//
// The skipList specifies location ranges that should be skipped on step into.
//
// This CDP parameter is experimental.
func (t *StepInto) SetSkipList(v []LocationRange) *StepInto {
	t.SkipList = v
	return t
}

// Do sends the StepInto CDP command to a browser,
// and returns the browser's response.
func (t *StepInto) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Debugger.stepInto", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the StepInto CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *StepInto) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Debugger.stepInto", b)
}

// ParseResponse parses the browser's response
// to the StepInto CDP command.
func (t *StepInto) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// StepOut contains the parameters, and acts as
// a Go receiver, for the CDP command `stepOut`.
//
// Steps out of the function call.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-stepOut
type StepOut struct{}

// NewStepOut constructs a new StepOut struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-stepOut
func NewStepOut() *StepOut {
	return &StepOut{}
}

// Do sends the StepOut CDP command to a browser,
// and returns the browser's response.
func (t *StepOut) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Debugger.stepOut", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the StepOut CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *StepOut) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Debugger.stepOut", nil)
}

// ParseResponse parses the browser's response
// to the StepOut CDP command.
func (t *StepOut) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// StepOver contains the parameters, and acts as
// a Go receiver, for the CDP command `stepOver`.
//
// Steps over the statement.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-stepOver
type StepOver struct {
	// The skipList specifies location ranges that should be skipped on step over.
	//
	// This CDP parameter is experimental.
	SkipList []LocationRange `json:"skipList,omitempty"`
}

// NewStepOver constructs a new StepOver struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#method-stepOver
func NewStepOver() *StepOver {
	return &StepOver{}
}

// SetSkipList adds or modifies the value of the optional
// parameter `skipList` in the StepOver CDP command.
//
// The skipList specifies location ranges that should be skipped on step over.
//
// This CDP parameter is experimental.
func (t *StepOver) SetSkipList(v []LocationRange) *StepOver {
	t.SkipList = v
	return t
}

// Do sends the StepOver CDP command to a browser,
// and returns the browser's response.
func (t *StepOver) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Debugger.stepOver", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the StepOver CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *StepOver) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Debugger.stepOver", b)
}

// ParseResponse parses the browser's response
// to the StepOver CDP command.
func (t *StepOver) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}
