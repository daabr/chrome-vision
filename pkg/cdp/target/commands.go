package target

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/cdp"
)

// ActivateTarget contains the parameters, and acts as
// a Go receiver, for the CDP command `activateTarget`.
//
// Activates (focuses) the target.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-activateTarget
type ActivateTarget struct {
	TargetID string
}

// NewActivateTarget constructs a new ActivateTarget struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-activateTarget
func NewActivateTarget(targetID string) *ActivateTarget {
	return &ActivateTarget{
		TargetID: targetID,
	}
}

// Do sends the ActivateTarget CDP command to a browser,
// and returns the browser's response.
func (t *ActivateTarget) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Target.activateTarget", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// AttachToTarget contains the parameters, and acts as
// a Go receiver, for the CDP command `attachToTarget`.
//
// Attaches to the target with given id.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-attachToTarget
type AttachToTarget struct {
	TargetID string
	// Enables "flat" access to the session via specifying sessionId attribute in the commands.
	// We plan to make this the default, deprecate non-flattened mode,
	// and eventually retire it. See crbug.com/991325.
	Flatten bool `json:"flatten,omitempty"`
}

// NewAttachToTarget constructs a new AttachToTarget struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-attachToTarget
func NewAttachToTarget(targetID string) *AttachToTarget {
	return &AttachToTarget{
		TargetID: targetID,
	}
}

// SetFlatten adds or modifies the value of the optional
// parameter `flatten` in the AttachToTarget CDP command.
//
// Enables "flat" access to the session via specifying sessionId attribute in the commands.
// We plan to make this the default, deprecate non-flattened mode,
// and eventually retire it. See crbug.com/991325.
func (t *AttachToTarget) SetFlatten(v bool) *AttachToTarget {
	t.Flatten = v
	return t
}

// AttachToTargetResponse contains the browser's response
// to calling the AttachToTarget CDP command with Do().
type AttachToTargetResponse struct {
	// Id assigned to the session.
	SessionID string
}

// Do sends the AttachToTarget CDP command to a browser,
// and returns the browser's response.
func (t *AttachToTarget) Do(ctx context.Context) (*AttachToTargetResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "Target.attachToTarget", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &AttachToTargetResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// AttachToBrowserTarget contains the parameters, and acts as
// a Go receiver, for the CDP command `attachToBrowserTarget`.
//
// Attaches to the browser target, only uses flat sessionId mode.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-attachToBrowserTarget
//
// This CDP method is experimental.
type AttachToBrowserTarget struct{}

// NewAttachToBrowserTarget constructs a new AttachToBrowserTarget struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-attachToBrowserTarget
//
// This CDP method is experimental.
func NewAttachToBrowserTarget() *AttachToBrowserTarget {
	return &AttachToBrowserTarget{}
}

// AttachToBrowserTargetResponse contains the browser's response
// to calling the AttachToBrowserTarget CDP command with Do().
type AttachToBrowserTargetResponse struct {
	// Id assigned to the session.
	SessionID string
}

// Do sends the AttachToBrowserTarget CDP command to a browser,
// and returns the browser's response.
func (t *AttachToBrowserTarget) Do(ctx context.Context) (*AttachToBrowserTargetResponse, error) {
	response, err := cdp.Send(ctx, "Target.attachToBrowserTarget", nil)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &AttachToBrowserTargetResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// CloseTarget contains the parameters, and acts as
// a Go receiver, for the CDP command `closeTarget`.
//
// Closes the target. If the target is a page that gets closed too.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-closeTarget
type CloseTarget struct {
	TargetID string
}

// NewCloseTarget constructs a new CloseTarget struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-closeTarget
func NewCloseTarget(targetID string) *CloseTarget {
	return &CloseTarget{
		TargetID: targetID,
	}
}

// CloseTargetResponse contains the browser's response
// to calling the CloseTarget CDP command with Do().
type CloseTargetResponse struct {
	// Always set to true. If an error occurs, the response indicates protocol error.
	//
	// This CDP parameter is deprecated.
	Success bool
}

// Do sends the CloseTarget CDP command to a browser,
// and returns the browser's response.
func (t *CloseTarget) Do(ctx context.Context) (*CloseTargetResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "Target.closeTarget", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &CloseTargetResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ExposeDevToolsProtocol contains the parameters, and acts as
// a Go receiver, for the CDP command `exposeDevToolsProtocol`.
//
// Inject object to the target's main frame that provides a communication
// channel with browser target.
//
// Injected object will be available as `window[bindingName]`.
//
// The object has the follwing API:
// - `binding.send(json)` - a method to send messages over the remote debugging protocol
// - `binding.onmessage = json => handleMessage(json)` - a callback that will be called for the protocol notifications and command responses.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-exposeDevToolsProtocol
//
// This CDP method is experimental.
type ExposeDevToolsProtocol struct {
	TargetID string
	// Binding name, 'cdp' if not specified.
	BindingName string `json:"bindingName,omitempty"`
}

// NewExposeDevToolsProtocol constructs a new ExposeDevToolsProtocol struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-exposeDevToolsProtocol
//
// This CDP method is experimental.
func NewExposeDevToolsProtocol(targetID string) *ExposeDevToolsProtocol {
	return &ExposeDevToolsProtocol{
		TargetID: targetID,
	}
}

// SetBindingName adds or modifies the value of the optional
// parameter `bindingName` in the ExposeDevToolsProtocol CDP command.
//
// Binding name, 'cdp' if not specified.
func (t *ExposeDevToolsProtocol) SetBindingName(v string) *ExposeDevToolsProtocol {
	t.BindingName = v
	return t
}

// Do sends the ExposeDevToolsProtocol CDP command to a browser,
// and returns the browser's response.
func (t *ExposeDevToolsProtocol) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Target.exposeDevToolsProtocol", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// CreateBrowserContext contains the parameters, and acts as
// a Go receiver, for the CDP command `createBrowserContext`.
//
// Creates a new empty BrowserContext. Similar to an incognito profile but you can have more than
// one.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-createBrowserContext
//
// This CDP method is experimental.
type CreateBrowserContext struct {
	// If specified, disposes this context when debugging session disconnects.
	DisposeOnDetach bool `json:"disposeOnDetach,omitempty"`
	// Proxy server, similar to the one passed to --proxy-server
	ProxyServer string `json:"proxyServer,omitempty"`
	// Proxy bypass list, similar to the one passed to --proxy-bypass-list
	ProxyBypassList string `json:"proxyBypassList,omitempty"`
}

// NewCreateBrowserContext constructs a new CreateBrowserContext struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-createBrowserContext
//
// This CDP method is experimental.
func NewCreateBrowserContext() *CreateBrowserContext {
	return &CreateBrowserContext{}
}

// SetDisposeOnDetach adds or modifies the value of the optional
// parameter `disposeOnDetach` in the CreateBrowserContext CDP command.
//
// If specified, disposes this context when debugging session disconnects.
func (t *CreateBrowserContext) SetDisposeOnDetach(v bool) *CreateBrowserContext {
	t.DisposeOnDetach = v
	return t
}

// SetProxyServer adds or modifies the value of the optional
// parameter `proxyServer` in the CreateBrowserContext CDP command.
//
// Proxy server, similar to the one passed to --proxy-server
func (t *CreateBrowserContext) SetProxyServer(v string) *CreateBrowserContext {
	t.ProxyServer = v
	return t
}

// SetProxyBypassList adds or modifies the value of the optional
// parameter `proxyBypassList` in the CreateBrowserContext CDP command.
//
// Proxy bypass list, similar to the one passed to --proxy-bypass-list
func (t *CreateBrowserContext) SetProxyBypassList(v string) *CreateBrowserContext {
	t.ProxyBypassList = v
	return t
}

// CreateBrowserContextResponse contains the browser's response
// to calling the CreateBrowserContext CDP command with Do().
type CreateBrowserContextResponse struct {
	// The id of the context created.
	BrowserContextID string
}

// Do sends the CreateBrowserContext CDP command to a browser,
// and returns the browser's response.
func (t *CreateBrowserContext) Do(ctx context.Context) (*CreateBrowserContextResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "Target.createBrowserContext", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &CreateBrowserContextResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetBrowserContexts contains the parameters, and acts as
// a Go receiver, for the CDP command `getBrowserContexts`.
//
// Returns all browser contexts created with `Target.createBrowserContext` method.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-getBrowserContexts
//
// This CDP method is experimental.
type GetBrowserContexts struct{}

// NewGetBrowserContexts constructs a new GetBrowserContexts struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-getBrowserContexts
//
// This CDP method is experimental.
func NewGetBrowserContexts() *GetBrowserContexts {
	return &GetBrowserContexts{}
}

// GetBrowserContextsResponse contains the browser's response
// to calling the GetBrowserContexts CDP command with Do().
type GetBrowserContextsResponse struct {
	// An array of browser context ids.
	BrowserContextIds []string
}

// Do sends the GetBrowserContexts CDP command to a browser,
// and returns the browser's response.
func (t *GetBrowserContexts) Do(ctx context.Context) (*GetBrowserContextsResponse, error) {
	response, err := cdp.Send(ctx, "Target.getBrowserContexts", nil)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetBrowserContextsResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateTarget contains the parameters, and acts as
// a Go receiver, for the CDP command `createTarget`.
//
// Creates a new page.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-createTarget
type CreateTarget struct {
	// The initial URL the page will be navigated to. An empty string indicates about:blank.
	URL string
	// Frame width in DIP (headless chrome only).
	Width int64 `json:"width,omitempty"`
	// Frame height in DIP (headless chrome only).
	Height int64 `json:"height,omitempty"`
	// The browser context to create the page in.
	BrowserContextID string `json:"browserContextId,omitempty"`
	// Whether BeginFrames for this target will be controlled via DevTools (headless chrome only,
	// not supported on MacOS yet, false by default).
	//
	// This CDP parameter is experimental.
	EnableBeginFrameControl bool `json:"enableBeginFrameControl,omitempty"`
	// Whether to create a new Window or Tab (chrome-only, false by default).
	NewWindow bool `json:"newWindow,omitempty"`
	// Whether to create the target in background or foreground (chrome-only,
	// false by default).
	Background bool `json:"background,omitempty"`
}

// NewCreateTarget constructs a new CreateTarget struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-createTarget
func NewCreateTarget(url string) *CreateTarget {
	return &CreateTarget{
		URL: url,
	}
}

// SetWidth adds or modifies the value of the optional
// parameter `width` in the CreateTarget CDP command.
//
// Frame width in DIP (headless chrome only).
func (t *CreateTarget) SetWidth(v int64) *CreateTarget {
	t.Width = v
	return t
}

// SetHeight adds or modifies the value of the optional
// parameter `height` in the CreateTarget CDP command.
//
// Frame height in DIP (headless chrome only).
func (t *CreateTarget) SetHeight(v int64) *CreateTarget {
	t.Height = v
	return t
}

// SetBrowserContextID adds or modifies the value of the optional
// parameter `browserContextId` in the CreateTarget CDP command.
//
// The browser context to create the page in.
func (t *CreateTarget) SetBrowserContextID(v string) *CreateTarget {
	t.BrowserContextID = v
	return t
}

// SetEnableBeginFrameControl adds or modifies the value of the optional
// parameter `enableBeginFrameControl` in the CreateTarget CDP command.
//
// Whether BeginFrames for this target will be controlled via DevTools (headless chrome only,
// not supported on MacOS yet, false by default).
//
// This CDP parameter is experimental.
func (t *CreateTarget) SetEnableBeginFrameControl(v bool) *CreateTarget {
	t.EnableBeginFrameControl = v
	return t
}

// SetNewWindow adds or modifies the value of the optional
// parameter `newWindow` in the CreateTarget CDP command.
//
// Whether to create a new Window or Tab (chrome-only, false by default).
func (t *CreateTarget) SetNewWindow(v bool) *CreateTarget {
	t.NewWindow = v
	return t
}

// SetBackground adds or modifies the value of the optional
// parameter `background` in the CreateTarget CDP command.
//
// Whether to create the target in background or foreground (chrome-only,
// false by default).
func (t *CreateTarget) SetBackground(v bool) *CreateTarget {
	t.Background = v
	return t
}

// CreateTargetResponse contains the browser's response
// to calling the CreateTarget CDP command with Do().
type CreateTargetResponse struct {
	// The id of the page opened.
	TargetID string
}

// Do sends the CreateTarget CDP command to a browser,
// and returns the browser's response.
func (t *CreateTarget) Do(ctx context.Context) (*CreateTargetResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "Target.createTarget", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &CreateTargetResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// DetachFromTarget contains the parameters, and acts as
// a Go receiver, for the CDP command `detachFromTarget`.
//
// Detaches session with given id.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-detachFromTarget
type DetachFromTarget struct {
	// Session to detach.
	SessionID string `json:"sessionId,omitempty"`
	// Deprecated.
	//
	// This CDP parameter is deprecated.
	TargetID string `json:"targetId,omitempty"`
}

// NewDetachFromTarget constructs a new DetachFromTarget struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-detachFromTarget
func NewDetachFromTarget() *DetachFromTarget {
	return &DetachFromTarget{}
}

// SetSessionID adds or modifies the value of the optional
// parameter `sessionId` in the DetachFromTarget CDP command.
//
// Session to detach.
func (t *DetachFromTarget) SetSessionID(v string) *DetachFromTarget {
	t.SessionID = v
	return t
}

// SetTargetID adds or modifies the value of the optional
// parameter `targetId` in the DetachFromTarget CDP command.
//
// Deprecated.
//
// This CDP parameter is deprecated.
func (t *DetachFromTarget) SetTargetID(v string) *DetachFromTarget {
	t.TargetID = v
	return t
}

// Do sends the DetachFromTarget CDP command to a browser,
// and returns the browser's response.
func (t *DetachFromTarget) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Target.detachFromTarget", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// DisposeBrowserContext contains the parameters, and acts as
// a Go receiver, for the CDP command `disposeBrowserContext`.
//
// Deletes a BrowserContext. All the belonging pages will be closed without calling their
// beforeunload hooks.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-disposeBrowserContext
//
// This CDP method is experimental.
type DisposeBrowserContext struct {
	BrowserContextID string
}

// NewDisposeBrowserContext constructs a new DisposeBrowserContext struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-disposeBrowserContext
//
// This CDP method is experimental.
func NewDisposeBrowserContext(browserContextID string) *DisposeBrowserContext {
	return &DisposeBrowserContext{
		BrowserContextID: browserContextID,
	}
}

// Do sends the DisposeBrowserContext CDP command to a browser,
// and returns the browser's response.
func (t *DisposeBrowserContext) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Target.disposeBrowserContext", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// GetTargetInfo contains the parameters, and acts as
// a Go receiver, for the CDP command `getTargetInfo`.
//
// Returns information about a target.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-getTargetInfo
//
// This CDP method is experimental.
type GetTargetInfo struct {
	TargetID string `json:"targetId,omitempty"`
}

// NewGetTargetInfo constructs a new GetTargetInfo struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-getTargetInfo
//
// This CDP method is experimental.
func NewGetTargetInfo() *GetTargetInfo {
	return &GetTargetInfo{}
}

// SetTargetID adds or modifies the value of the optional
// parameter `targetId` in the GetTargetInfo CDP command.
func (t *GetTargetInfo) SetTargetID(v string) *GetTargetInfo {
	t.TargetID = v
	return t
}

// GetTargetInfoResponse contains the browser's response
// to calling the GetTargetInfo CDP command with Do().
type GetTargetInfoResponse struct {
	TargetInfo TargetInfo
}

// Do sends the GetTargetInfo CDP command to a browser,
// and returns the browser's response.
func (t *GetTargetInfo) Do(ctx context.Context) (*GetTargetInfoResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "Target.getTargetInfo", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetTargetInfoResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetTargets contains the parameters, and acts as
// a Go receiver, for the CDP command `getTargets`.
//
// Retrieves a list of available targets.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-getTargets
type GetTargets struct{}

// NewGetTargets constructs a new GetTargets struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-getTargets
func NewGetTargets() *GetTargets {
	return &GetTargets{}
}

// GetTargetsResponse contains the browser's response
// to calling the GetTargets CDP command with Do().
type GetTargetsResponse struct {
	// The list of targets.
	TargetInfos []TargetInfo
}

// Do sends the GetTargets CDP command to a browser,
// and returns the browser's response.
func (t *GetTargets) Do(ctx context.Context) (*GetTargetsResponse, error) {
	response, err := cdp.Send(ctx, "Target.getTargets", nil)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetTargetsResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SendMessageToTarget contains the parameters, and acts as
// a Go receiver, for the CDP command `sendMessageToTarget`.
//
// Sends protocol message over session with given id.
// Consider using flat mode instead; see commands attachToTarget, setAutoAttach,
// and crbug.com/991325.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-sendMessageToTarget
//
// This CDP method is deprecated.
type SendMessageToTarget struct {
	Message string
	// Identifier of the session.
	SessionID string `json:"sessionId,omitempty"`
	// Deprecated.
	//
	// This CDP parameter is deprecated.
	TargetID string `json:"targetId,omitempty"`
}

// NewSendMessageToTarget constructs a new SendMessageToTarget struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-sendMessageToTarget
//
// This CDP method is deprecated.
func NewSendMessageToTarget(message string) *SendMessageToTarget {
	return &SendMessageToTarget{
		Message: message,
	}
}

// SetSessionID adds or modifies the value of the optional
// parameter `sessionId` in the SendMessageToTarget CDP command.
//
// Identifier of the session.
func (t *SendMessageToTarget) SetSessionID(v string) *SendMessageToTarget {
	t.SessionID = v
	return t
}

// SetTargetID adds or modifies the value of the optional
// parameter `targetId` in the SendMessageToTarget CDP command.
//
// Deprecated.
//
// This CDP parameter is deprecated.
func (t *SendMessageToTarget) SetTargetID(v string) *SendMessageToTarget {
	t.TargetID = v
	return t
}

// Do sends the SendMessageToTarget CDP command to a browser,
// and returns the browser's response.
func (t *SendMessageToTarget) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Target.sendMessageToTarget", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetAutoAttach contains the parameters, and acts as
// a Go receiver, for the CDP command `setAutoAttach`.
//
// Controls whether to automatically attach to new targets which are considered to be related to
// this one. When turned on, attaches to all existing related targets as well. When turned off,
// automatically detaches from all currently attached targets.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-setAutoAttach
//
// This CDP method is experimental.
type SetAutoAttach struct {
	// Whether to auto-attach to related targets.
	AutoAttach bool
	// Whether to pause new targets when attaching to them. Use `Runtime.runIfWaitingForDebugger`
	// to run paused targets.
	WaitForDebuggerOnStart bool
	// Enables "flat" access to the session via specifying sessionId attribute in the commands.
	// We plan to make this the default, deprecate non-flattened mode,
	// and eventually retire it. See crbug.com/991325.
	Flatten bool `json:"flatten,omitempty"`
}

// NewSetAutoAttach constructs a new SetAutoAttach struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-setAutoAttach
//
// This CDP method is experimental.
func NewSetAutoAttach(autoAttach bool, waitForDebuggerOnStart bool) *SetAutoAttach {
	return &SetAutoAttach{
		AutoAttach:             autoAttach,
		WaitForDebuggerOnStart: waitForDebuggerOnStart,
	}
}

// SetFlatten adds or modifies the value of the optional
// parameter `flatten` in the SetAutoAttach CDP command.
//
// Enables "flat" access to the session via specifying sessionId attribute in the commands.
// We plan to make this the default, deprecate non-flattened mode,
// and eventually retire it. See crbug.com/991325.
func (t *SetAutoAttach) SetFlatten(v bool) *SetAutoAttach {
	t.Flatten = v
	return t
}

// Do sends the SetAutoAttach CDP command to a browser,
// and returns the browser's response.
func (t *SetAutoAttach) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Target.setAutoAttach", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetDiscoverTargets contains the parameters, and acts as
// a Go receiver, for the CDP command `setDiscoverTargets`.
//
// Controls whether to discover available targets and notify via
// `targetCreated/targetInfoChanged/targetDestroyed` events.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-setDiscoverTargets
type SetDiscoverTargets struct {
	// Whether to discover available targets.
	Discover bool
}

// NewSetDiscoverTargets constructs a new SetDiscoverTargets struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-setDiscoverTargets
func NewSetDiscoverTargets(discover bool) *SetDiscoverTargets {
	return &SetDiscoverTargets{
		Discover: discover,
	}
}

// Do sends the SetDiscoverTargets CDP command to a browser,
// and returns the browser's response.
func (t *SetDiscoverTargets) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Target.setDiscoverTargets", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetRemoteLocations contains the parameters, and acts as
// a Go receiver, for the CDP command `setRemoteLocations`.
//
// Enables target discovery for the specified locations, when `setDiscoverTargets` was set to
// `true`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-setRemoteLocations
//
// This CDP method is experimental.
type SetRemoteLocations struct {
	// List of remote locations.
	Locations []RemoteLocation
}

// NewSetRemoteLocations constructs a new SetRemoteLocations struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-setRemoteLocations
//
// This CDP method is experimental.
func NewSetRemoteLocations(locations []RemoteLocation) *SetRemoteLocations {
	return &SetRemoteLocations{
		Locations: locations,
	}
}

// Do sends the SetRemoteLocations CDP command to a browser,
// and returns the browser's response.
func (t *SetRemoteLocations) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Target.setRemoteLocations", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}
