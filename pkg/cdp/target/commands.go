package target

import (
	"context"
	"fmt"

	"github.com/daabr/chrome-vision/pkg/cdp/browser"
)

// ActivateTarget contains the parameters, and acts as
// a Go receiver, for the CDP command `activateTarget`.
//
// Activates (focuses) the target.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-activateTarget
type ActivateTarget struct {
	TargetID TargetID `json:"targetId"`
}

// NewActivateTarget constructs a new ActivateTarget struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-activateTarget
func NewActivateTarget(targetId TargetID) *ActivateTarget {
	return &ActivateTarget{
		TargetID: targetId,
	}
}

// Do sends the ActivateTarget CDP command to a browser,
// and returns the browser's response.
func (t *ActivateTarget) Do(ctx context.Context) error {
	fmt.Println(ctx)
	return nil
}

// AttachToTarget contains the parameters, and acts as
// a Go receiver, for the CDP command `attachToTarget`.
//
// Attaches to the target with given id.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-attachToTarget
type AttachToTarget struct {
	TargetID TargetID `json:"targetId"`
	// Enables "flat" access to the session via specifying sessionId attribute in the commands.
	// We plan to make this the default, deprecate non-flattened mode,
	// and eventually retire it. See crbug.com/991325.
	Flatten bool `json:"flatten,omitempty"`
}

// NewAttachToTarget constructs a new AttachToTarget struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-attachToTarget
func NewAttachToTarget(targetId TargetID) *AttachToTarget {
	return &AttachToTarget{
		TargetID: targetId,
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
	SessionID SessionID `json:"sessionId"`
}

// Do sends the AttachToTarget CDP command to a browser,
// and returns the browser's response.
func (t *AttachToTarget) Do(ctx context.Context) (*AttachToTargetResponse, error) {
	fmt.Println(ctx)
	return new(AttachToTargetResponse), nil
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
// all the required parameters, and only them. Optional parameters
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
	SessionID SessionID `json:"sessionId"`
}

// Do sends the AttachToBrowserTarget CDP command to a browser,
// and returns the browser's response.
func (t *AttachToBrowserTarget) Do(ctx context.Context) (*AttachToBrowserTargetResponse, error) {
	fmt.Println(ctx)
	return new(AttachToBrowserTargetResponse), nil
}

// CloseTarget contains the parameters, and acts as
// a Go receiver, for the CDP command `closeTarget`.
//
// Closes the target. If the target is a page that gets closed too.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-closeTarget
type CloseTarget struct {
	TargetID TargetID `json:"targetId"`
}

// NewCloseTarget constructs a new CloseTarget struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-closeTarget
func NewCloseTarget(targetId TargetID) *CloseTarget {
	return &CloseTarget{
		TargetID: targetId,
	}
}

// CloseTargetResponse contains the browser's response
// to calling the CloseTarget CDP command with Do().
type CloseTargetResponse struct {
	// Always set to true. If an error occurs, the response indicates protocol error.
	//
	// This CDP parameter is deprecated.
	Success bool `json:"success"`
}

// Do sends the CloseTarget CDP command to a browser,
// and returns the browser's response.
func (t *CloseTarget) Do(ctx context.Context) (*CloseTargetResponse, error) {
	fmt.Println(ctx)
	return new(CloseTargetResponse), nil
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
	TargetID TargetID `json:"targetId"`
	// Binding name, 'cdp' if not specified.
	BindingName string `json:"bindingName,omitempty"`
}

// NewExposeDevToolsProtocol constructs a new ExposeDevToolsProtocol struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-exposeDevToolsProtocol
//
// This CDP method is experimental.
func NewExposeDevToolsProtocol(targetId TargetID) *ExposeDevToolsProtocol {
	return &ExposeDevToolsProtocol{
		TargetID: targetId,
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
	fmt.Println(ctx)
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
// all the required parameters, and only them. Optional parameters
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
	BrowserContextID browser.BrowserContextID `json:"browserContextId"`
}

// Do sends the CreateBrowserContext CDP command to a browser,
// and returns the browser's response.
func (t *CreateBrowserContext) Do(ctx context.Context) (*CreateBrowserContextResponse, error) {
	fmt.Println(ctx)
	return new(CreateBrowserContextResponse), nil
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
// all the required parameters, and only them. Optional parameters
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
	BrowserContextIds []browser.BrowserContextID `json:"browserContextIds"`
}

// Do sends the GetBrowserContexts CDP command to a browser,
// and returns the browser's response.
func (t *GetBrowserContexts) Do(ctx context.Context) (*GetBrowserContextsResponse, error) {
	fmt.Println(ctx)
	return new(GetBrowserContextsResponse), nil
}

// CreateTarget contains the parameters, and acts as
// a Go receiver, for the CDP command `createTarget`.
//
// Creates a new page.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-createTarget
type CreateTarget struct {
	// The initial URL the page will be navigated to. An empty string indicates about:blank.
	URL string `json:"url"`
	// Frame width in DIP (headless chrome only).
	Width int64 `json:"width,omitempty"`
	// Frame height in DIP (headless chrome only).
	Height int64 `json:"height,omitempty"`
	// The browser context to create the page in.
	BrowserContextID *browser.BrowserContextID `json:"browserContextId,omitempty"`
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
// all the required parameters, and only them. Optional parameters
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
func (t *CreateTarget) SetBrowserContextID(v browser.BrowserContextID) *CreateTarget {
	t.BrowserContextID = &v
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
	TargetID TargetID `json:"targetId"`
}

// Do sends the CreateTarget CDP command to a browser,
// and returns the browser's response.
func (t *CreateTarget) Do(ctx context.Context) (*CreateTargetResponse, error) {
	fmt.Println(ctx)
	return new(CreateTargetResponse), nil
}

// DetachFromTarget contains the parameters, and acts as
// a Go receiver, for the CDP command `detachFromTarget`.
//
// Detaches session with given id.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-detachFromTarget
type DetachFromTarget struct {
	// Session to detach.
	SessionID *SessionID `json:"sessionId,omitempty"`
	// Deprecated.
	//
	// This CDP parameter is deprecated.
	TargetID *TargetID `json:"targetId,omitempty"`
}

// NewDetachFromTarget constructs a new DetachFromTarget struct instance, with
// all the required parameters, and only them. Optional parameters
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
func (t *DetachFromTarget) SetSessionID(v SessionID) *DetachFromTarget {
	t.SessionID = &v
	return t
}

// SetTargetID adds or modifies the value of the optional
// parameter `targetId` in the DetachFromTarget CDP command.
//
// Deprecated.
//
// This CDP parameter is deprecated.
func (t *DetachFromTarget) SetTargetID(v TargetID) *DetachFromTarget {
	t.TargetID = &v
	return t
}

// Do sends the DetachFromTarget CDP command to a browser,
// and returns the browser's response.
func (t *DetachFromTarget) Do(ctx context.Context) error {
	fmt.Println(ctx)
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
	BrowserContextID browser.BrowserContextID `json:"browserContextId"`
}

// NewDisposeBrowserContext constructs a new DisposeBrowserContext struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-disposeBrowserContext
//
// This CDP method is experimental.
func NewDisposeBrowserContext(browserContextId browser.BrowserContextID) *DisposeBrowserContext {
	return &DisposeBrowserContext{
		BrowserContextID: browserContextId,
	}
}

// Do sends the DisposeBrowserContext CDP command to a browser,
// and returns the browser's response.
func (t *DisposeBrowserContext) Do(ctx context.Context) error {
	fmt.Println(ctx)
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
	TargetID *TargetID `json:"targetId,omitempty"`
}

// NewGetTargetInfo constructs a new GetTargetInfo struct instance, with
// all the required parameters, and only them. Optional parameters
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
func (t *GetTargetInfo) SetTargetID(v TargetID) *GetTargetInfo {
	t.TargetID = &v
	return t
}

// GetTargetInfoResponse contains the browser's response
// to calling the GetTargetInfo CDP command with Do().
type GetTargetInfoResponse struct {
	TargetInfo TargetInfo `json:"targetInfo"`
}

// Do sends the GetTargetInfo CDP command to a browser,
// and returns the browser's response.
func (t *GetTargetInfo) Do(ctx context.Context) (*GetTargetInfoResponse, error) {
	fmt.Println(ctx)
	return new(GetTargetInfoResponse), nil
}

// GetTargets contains the parameters, and acts as
// a Go receiver, for the CDP command `getTargets`.
//
// Retrieves a list of available targets.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-getTargets
type GetTargets struct{}

// NewGetTargets constructs a new GetTargets struct instance, with
// all the required parameters, and only them. Optional parameters
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
	TargetInfos []TargetInfo `json:"targetInfos"`
}

// Do sends the GetTargets CDP command to a browser,
// and returns the browser's response.
func (t *GetTargets) Do(ctx context.Context) (*GetTargetsResponse, error) {
	fmt.Println(ctx)
	return new(GetTargetsResponse), nil
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
	Message string `json:"message"`
	// Identifier of the session.
	SessionID *SessionID `json:"sessionId,omitempty"`
	// Deprecated.
	//
	// This CDP parameter is deprecated.
	TargetID *TargetID `json:"targetId,omitempty"`
}

// NewSendMessageToTarget constructs a new SendMessageToTarget struct instance, with
// all the required parameters, and only them. Optional parameters
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
func (t *SendMessageToTarget) SetSessionID(v SessionID) *SendMessageToTarget {
	t.SessionID = &v
	return t
}

// SetTargetID adds or modifies the value of the optional
// parameter `targetId` in the SendMessageToTarget CDP command.
//
// Deprecated.
//
// This CDP parameter is deprecated.
func (t *SendMessageToTarget) SetTargetID(v TargetID) *SendMessageToTarget {
	t.TargetID = &v
	return t
}

// Do sends the SendMessageToTarget CDP command to a browser,
// and returns the browser's response.
func (t *SendMessageToTarget) Do(ctx context.Context) error {
	fmt.Println(ctx)
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
	AutoAttach bool `json:"autoAttach"`
	// Whether to pause new targets when attaching to them. Use `Runtime.runIfWaitingForDebugger`
	// to run paused targets.
	WaitForDebuggerOnStart bool `json:"waitForDebuggerOnStart"`
	// Enables "flat" access to the session via specifying sessionId attribute in the commands.
	// We plan to make this the default, deprecate non-flattened mode,
	// and eventually retire it. See crbug.com/991325.
	Flatten bool `json:"flatten,omitempty"`
}

// NewSetAutoAttach constructs a new SetAutoAttach struct instance, with
// all the required parameters, and only them. Optional parameters
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
	fmt.Println(ctx)
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
	Discover bool `json:"discover"`
}

// NewSetDiscoverTargets constructs a new SetDiscoverTargets struct instance, with
// all the required parameters, and only them. Optional parameters
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
	fmt.Println(ctx)
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
	Locations []RemoteLocation `json:"locations"`
}

// NewSetRemoteLocations constructs a new SetRemoteLocations struct instance, with
// all the required parameters, and only them. Optional parameters
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
	fmt.Println(ctx)
	return nil
}
