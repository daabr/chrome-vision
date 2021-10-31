package target

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
)

// ActivateTarget contains the parameters, and acts as
// a Go receiver, for the CDP command `activateTarget`.
//
// Activates (focuses) the target.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-activateTarget
type ActivateTarget struct {
	TargetID string `json:"targetId"`
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
	m, err := devtools.SendAndWait(ctx, "Target.activateTarget", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the ActivateTarget CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *ActivateTarget) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Target.activateTarget", b)
}

// ParseResponse parses the browser's response
// to the ActivateTarget CDP command.
func (t *ActivateTarget) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
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
	TargetID string `json:"targetId"`
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

// AttachToTargetResult contains the browser's response
// to calling the AttachToTarget CDP command with Do().
type AttachToTargetResult struct {
	// Id assigned to the session.
	SessionID string `json:"sessionId"`
}

// Do sends the AttachToTarget CDP command to a browser,
// and returns the browser's response.
func (t *AttachToTarget) Do(ctx context.Context) (*AttachToTargetResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Target.attachToTarget", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the AttachToTarget CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *AttachToTarget) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Target.attachToTarget", b)
}

// ParseResponse parses the browser's response
// to the AttachToTarget CDP command.
func (t *AttachToTarget) ParseResponse(m *devtools.Message) (*AttachToTargetResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &AttachToTargetResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
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

// AttachToBrowserTargetResult contains the browser's response
// to calling the AttachToBrowserTarget CDP command with Do().
type AttachToBrowserTargetResult struct {
	// Id assigned to the session.
	SessionID string `json:"sessionId"`
}

// Do sends the AttachToBrowserTarget CDP command to a browser,
// and returns the browser's response.
func (t *AttachToBrowserTarget) Do(ctx context.Context) (*AttachToBrowserTargetResult, error) {
	m, err := devtools.SendAndWait(ctx, "Target.attachToBrowserTarget", nil)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the AttachToBrowserTarget CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *AttachToBrowserTarget) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Target.attachToBrowserTarget", nil)
}

// ParseResponse parses the browser's response
// to the AttachToBrowserTarget CDP command.
func (t *AttachToBrowserTarget) ParseResponse(m *devtools.Message) (*AttachToBrowserTargetResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &AttachToBrowserTargetResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
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
	TargetID string `json:"targetId"`
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

// CloseTargetResult contains the browser's response
// to calling the CloseTarget CDP command with Do().
type CloseTargetResult struct {
	// Always set to true. If an error occurs, the response indicates protocol error.
	//
	// This CDP parameter is deprecated.
	Success bool `json:"success"`
}

// Do sends the CloseTarget CDP command to a browser,
// and returns the browser's response.
func (t *CloseTarget) Do(ctx context.Context) (*CloseTargetResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Target.closeTarget", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the CloseTarget CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *CloseTarget) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Target.closeTarget", b)
}

// ParseResponse parses the browser's response
// to the CloseTarget CDP command.
func (t *CloseTarget) ParseResponse(m *devtools.Message) (*CloseTargetResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &CloseTargetResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
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
	TargetID string `json:"targetId"`
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
	m, err := devtools.SendAndWait(ctx, "Target.exposeDevToolsProtocol", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the ExposeDevToolsProtocol CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *ExposeDevToolsProtocol) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Target.exposeDevToolsProtocol", b)
}

// ParseResponse parses the browser's response
// to the ExposeDevToolsProtocol CDP command.
func (t *ExposeDevToolsProtocol) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
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

// CreateBrowserContextResult contains the browser's response
// to calling the CreateBrowserContext CDP command with Do().
type CreateBrowserContextResult struct {
	// The id of the context created.
	BrowserContextID string `json:"browserContextId"`
}

// Do sends the CreateBrowserContext CDP command to a browser,
// and returns the browser's response.
func (t *CreateBrowserContext) Do(ctx context.Context) (*CreateBrowserContextResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Target.createBrowserContext", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the CreateBrowserContext CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *CreateBrowserContext) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Target.createBrowserContext", b)
}

// ParseResponse parses the browser's response
// to the CreateBrowserContext CDP command.
func (t *CreateBrowserContext) ParseResponse(m *devtools.Message) (*CreateBrowserContextResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &CreateBrowserContextResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
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

// GetBrowserContextsResult contains the browser's response
// to calling the GetBrowserContexts CDP command with Do().
type GetBrowserContextsResult struct {
	// An array of browser context ids.
	BrowserContextIds []string `json:"browserContextIds"`
}

// Do sends the GetBrowserContexts CDP command to a browser,
// and returns the browser's response.
func (t *GetBrowserContexts) Do(ctx context.Context) (*GetBrowserContextsResult, error) {
	m, err := devtools.SendAndWait(ctx, "Target.getBrowserContexts", nil)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetBrowserContexts CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetBrowserContexts) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Target.getBrowserContexts", nil)
}

// ParseResponse parses the browser's response
// to the GetBrowserContexts CDP command.
func (t *GetBrowserContexts) ParseResponse(m *devtools.Message) (*GetBrowserContextsResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetBrowserContextsResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
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
	URL string `json:"url"`
	// Frame width in DIP (headless chrome only).
	Width int64 `json:"width,omitempty"`
	// Frame height in DIP (headless chrome only).
	Height int64 `json:"height,omitempty"`
	// The browser context to create the page in.
	//
	// This CDP parameter is experimental.
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
//
// This CDP parameter is experimental.
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

// CreateTargetResult contains the browser's response
// to calling the CreateTarget CDP command with Do().
type CreateTargetResult struct {
	// The id of the page opened.
	TargetID string `json:"targetId"`
}

// Do sends the CreateTarget CDP command to a browser,
// and returns the browser's response.
func (t *CreateTarget) Do(ctx context.Context) (*CreateTargetResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Target.createTarget", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the CreateTarget CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *CreateTarget) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Target.createTarget", b)
}

// ParseResponse parses the browser's response
// to the CreateTarget CDP command.
func (t *CreateTarget) ParseResponse(m *devtools.Message) (*CreateTargetResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &CreateTargetResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
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
	m, err := devtools.SendAndWait(ctx, "Target.detachFromTarget", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the DetachFromTarget CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *DetachFromTarget) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Target.detachFromTarget", b)
}

// ParseResponse parses the browser's response
// to the DetachFromTarget CDP command.
func (t *DetachFromTarget) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
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
	BrowserContextID string `json:"browserContextId"`
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
	m, err := devtools.SendAndWait(ctx, "Target.disposeBrowserContext", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the DisposeBrowserContext CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *DisposeBrowserContext) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Target.disposeBrowserContext", b)
}

// ParseResponse parses the browser's response
// to the DisposeBrowserContext CDP command.
func (t *DisposeBrowserContext) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
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

// GetTargetInfoResult contains the browser's response
// to calling the GetTargetInfo CDP command with Do().
type GetTargetInfoResult struct {
	TargetInfo Info `json:"targetInfo"`
}

// Do sends the GetTargetInfo CDP command to a browser,
// and returns the browser's response.
func (t *GetTargetInfo) Do(ctx context.Context) (*GetTargetInfoResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Target.getTargetInfo", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetTargetInfo CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetTargetInfo) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Target.getTargetInfo", b)
}

// ParseResponse parses the browser's response
// to the GetTargetInfo CDP command.
func (t *GetTargetInfo) ParseResponse(m *devtools.Message) (*GetTargetInfoResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetTargetInfoResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
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

// GetTargetsResult contains the browser's response
// to calling the GetTargets CDP command with Do().
type GetTargetsResult struct {
	// The list of targets.
	TargetInfos []Info `json:"targetInfos"`
}

// Do sends the GetTargets CDP command to a browser,
// and returns the browser's response.
func (t *GetTargets) Do(ctx context.Context) (*GetTargetsResult, error) {
	m, err := devtools.SendAndWait(ctx, "Target.getTargets", nil)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetTargets CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetTargets) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Target.getTargets", nil)
}

// ParseResponse parses the browser's response
// to the GetTargets CDP command.
func (t *GetTargets) ParseResponse(m *devtools.Message) (*GetTargetsResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetTargetsResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
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
	Message string `json:"message"`
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
	m, err := devtools.SendAndWait(ctx, "Target.sendMessageToTarget", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SendMessageToTarget CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SendMessageToTarget) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Target.sendMessageToTarget", b)
}

// ParseResponse parses the browser's response
// to the SendMessageToTarget CDP command.
func (t *SendMessageToTarget) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetAutoAttach contains the parameters, and acts as
// a Go receiver, for the CDP command `setAutoAttach`.
//
// Controls whether to automatically attach to new targets which are considered to be related to
// this one. When turned on, attaches to all existing related targets as well. When turned off,
// automatically detaches from all currently attached targets.
// This also clears all targets added by `autoAttachRelated` from the list of targets to watch
// for creation of related targets.
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
	m, err := devtools.SendAndWait(ctx, "Target.setAutoAttach", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetAutoAttach CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetAutoAttach) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Target.setAutoAttach", b)
}

// ParseResponse parses the browser's response
// to the SetAutoAttach CDP command.
func (t *SetAutoAttach) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// AutoAttachRelated contains the parameters, and acts as
// a Go receiver, for the CDP command `autoAttachRelated`.
//
// Adds the specified target to the list of targets that will be monitored for any related target
// creation (such as child frames, child workers and new versions of service worker) and reported
// through `attachedToTarget`. The specified target is also auto-attached.
// This cancels the effect of any previous `setAutoAttach` and is also cancelled by subsequent
// `setAutoAttach`. Only available at the Browser target.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-autoAttachRelated
//
// This CDP method is experimental.
type AutoAttachRelated struct {
	TargetID string `json:"targetId"`
	// Whether to pause new targets when attaching to them. Use `Runtime.runIfWaitingForDebugger`
	// to run paused targets.
	WaitForDebuggerOnStart bool `json:"waitForDebuggerOnStart"`
}

// NewAutoAttachRelated constructs a new AutoAttachRelated struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-autoAttachRelated
//
// This CDP method is experimental.
func NewAutoAttachRelated(targetID string, waitForDebuggerOnStart bool) *AutoAttachRelated {
	return &AutoAttachRelated{
		TargetID:               targetID,
		WaitForDebuggerOnStart: waitForDebuggerOnStart,
	}
}

// Do sends the AutoAttachRelated CDP command to a browser,
// and returns the browser's response.
func (t *AutoAttachRelated) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Target.autoAttachRelated", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the AutoAttachRelated CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *AutoAttachRelated) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Target.autoAttachRelated", b)
}

// ParseResponse parses the browser's response
// to the AutoAttachRelated CDP command.
func (t *AutoAttachRelated) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
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
	Discover bool `json:"discover"`
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
	m, err := devtools.SendAndWait(ctx, "Target.setDiscoverTargets", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetDiscoverTargets CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetDiscoverTargets) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Target.setDiscoverTargets", b)
}

// ParseResponse parses the browser's response
// to the SetDiscoverTargets CDP command.
func (t *SetDiscoverTargets) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
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
	Locations []RemoteLocation `json:"locations"`
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
	m, err := devtools.SendAndWait(ctx, "Target.setRemoteLocations", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetRemoteLocations CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetRemoteLocations) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Target.setRemoteLocations", b)
}

// ParseResponse parses the browser's response
// to the SetRemoteLocations CDP command.
func (t *SetRemoteLocations) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}
