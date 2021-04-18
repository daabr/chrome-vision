package browser

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/cdp"
)

// SetPermission contains the parameters, and acts as
// a Go receiver, for the CDP command `setPermission`.
//
// Set permission settings for given origin.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-setPermission
//
// This CDP method is experimental.
type SetPermission struct {
	// Descriptor of permission to override.
	Permission PermissionDescriptor `json:"permission"`
	// Setting of the permission.
	Setting PermissionSetting `json:"setting"`
	// Origin the permission applies to, all origins if not specified.
	Origin string `json:"origin,omitempty"`
	// Context to override. When omitted, default browser context is used.
	BrowserContextID *BrowserContextID `json:"browserContextId,omitempty"`
}

// NewSetPermission constructs a new SetPermission struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-setPermission
//
// This CDP method is experimental.
func NewSetPermission(permission PermissionDescriptor, setting PermissionSetting) *SetPermission {
	return &SetPermission{
		Permission: permission,
		Setting:    setting,
	}
}

// SetOrigin adds or modifies the value of the optional
// parameter `origin` in the SetPermission CDP command.
//
// Origin the permission applies to, all origins if not specified.
func (t *SetPermission) SetOrigin(v string) *SetPermission {
	t.Origin = v
	return t
}

// SetBrowserContextID adds or modifies the value of the optional
// parameter `browserContextId` in the SetPermission CDP command.
//
// Context to override. When omitted, default browser context is used.
func (t *SetPermission) SetBrowserContextID(v BrowserContextID) *SetPermission {
	t.BrowserContextID = &v
	return t
}

// Do sends the SetPermission CDP command to a browser,
// and returns the browser's response.
func (t *SetPermission) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Browser.setPermission", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// GrantPermissions contains the parameters, and acts as
// a Go receiver, for the CDP command `grantPermissions`.
//
// Grant specific permissions to the given origin and reject all others.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-grantPermissions
//
// This CDP method is experimental.
type GrantPermissions struct {
	Permissions []PermissionType `json:"permissions"`
	// Origin the permission applies to, all origins if not specified.
	Origin string `json:"origin,omitempty"`
	// BrowserContext to override permissions. When omitted, default browser context is used.
	BrowserContextID *BrowserContextID `json:"browserContextId,omitempty"`
}

// NewGrantPermissions constructs a new GrantPermissions struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-grantPermissions
//
// This CDP method is experimental.
func NewGrantPermissions(permissions []PermissionType) *GrantPermissions {
	return &GrantPermissions{
		Permissions: permissions,
	}
}

// SetOrigin adds or modifies the value of the optional
// parameter `origin` in the GrantPermissions CDP command.
//
// Origin the permission applies to, all origins if not specified.
func (t *GrantPermissions) SetOrigin(v string) *GrantPermissions {
	t.Origin = v
	return t
}

// SetBrowserContextID adds or modifies the value of the optional
// parameter `browserContextId` in the GrantPermissions CDP command.
//
// BrowserContext to override permissions. When omitted, default browser context is used.
func (t *GrantPermissions) SetBrowserContextID(v BrowserContextID) *GrantPermissions {
	t.BrowserContextID = &v
	return t
}

// Do sends the GrantPermissions CDP command to a browser,
// and returns the browser's response.
func (t *GrantPermissions) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Browser.grantPermissions", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// ResetPermissions contains the parameters, and acts as
// a Go receiver, for the CDP command `resetPermissions`.
//
// Reset all permission management for all origins.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-resetPermissions
//
// This CDP method is experimental.
type ResetPermissions struct {
	// BrowserContext to reset permissions. When omitted, default browser context is used.
	BrowserContextID *BrowserContextID `json:"browserContextId,omitempty"`
}

// NewResetPermissions constructs a new ResetPermissions struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-resetPermissions
//
// This CDP method is experimental.
func NewResetPermissions() *ResetPermissions {
	return &ResetPermissions{}
}

// SetBrowserContextID adds or modifies the value of the optional
// parameter `browserContextId` in the ResetPermissions CDP command.
//
// BrowserContext to reset permissions. When omitted, default browser context is used.
func (t *ResetPermissions) SetBrowserContextID(v BrowserContextID) *ResetPermissions {
	t.BrowserContextID = &v
	return t
}

// Do sends the ResetPermissions CDP command to a browser,
// and returns the browser's response.
func (t *ResetPermissions) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Browser.resetPermissions", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetDownloadBehavior contains the parameters, and acts as
// a Go receiver, for the CDP command `setDownloadBehavior`.
//
// Set the behavior when downloading a file.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-setDownloadBehavior
//
// This CDP method is experimental.
type SetDownloadBehavior struct {
	// Whether to allow all or deny all download requests, or use default Chrome behavior if
	// available (otherwise deny). |allowAndName| allows download and names files according to
	// their dowmload guids.
	Behavior string `json:"behavior"`
	// BrowserContext to set download behavior. When omitted, default browser context is used.
	BrowserContextID *BrowserContextID `json:"browserContextId,omitempty"`
	// The default path to save downloaded files to. This is requred if behavior is set to 'allow'
	// or 'allowAndName'.
	DownloadPath string `json:"downloadPath,omitempty"`
	// Whether to emit download events (defaults to false).
	EventsEnabled bool `json:"eventsEnabled,omitempty"`
}

// NewSetDownloadBehavior constructs a new SetDownloadBehavior struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-setDownloadBehavior
//
// This CDP method is experimental.
func NewSetDownloadBehavior(behavior string) *SetDownloadBehavior {
	return &SetDownloadBehavior{
		Behavior: behavior,
	}
}

// SetBrowserContextID adds or modifies the value of the optional
// parameter `browserContextId` in the SetDownloadBehavior CDP command.
//
// BrowserContext to set download behavior. When omitted, default browser context is used.
func (t *SetDownloadBehavior) SetBrowserContextID(v BrowserContextID) *SetDownloadBehavior {
	t.BrowserContextID = &v
	return t
}

// SetDownloadPath adds or modifies the value of the optional
// parameter `downloadPath` in the SetDownloadBehavior CDP command.
//
// The default path to save downloaded files to. This is requred if behavior is set to 'allow'
// or 'allowAndName'.
func (t *SetDownloadBehavior) SetDownloadPath(v string) *SetDownloadBehavior {
	t.DownloadPath = v
	return t
}

// SetEventsEnabled adds or modifies the value of the optional
// parameter `eventsEnabled` in the SetDownloadBehavior CDP command.
//
// Whether to emit download events (defaults to false).
func (t *SetDownloadBehavior) SetEventsEnabled(v bool) *SetDownloadBehavior {
	t.EventsEnabled = v
	return t
}

// Do sends the SetDownloadBehavior CDP command to a browser,
// and returns the browser's response.
func (t *SetDownloadBehavior) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Browser.setDownloadBehavior", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// CancelDownload contains the parameters, and acts as
// a Go receiver, for the CDP command `cancelDownload`.
//
// Cancel a download if in progress
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-cancelDownload
//
// This CDP method is experimental.
type CancelDownload struct {
	// Global unique identifier of the download.
	Guid string `json:"guid"`
	// BrowserContext to perform the action in. When omitted, default browser context is used.
	BrowserContextID *BrowserContextID `json:"browserContextId,omitempty"`
}

// NewCancelDownload constructs a new CancelDownload struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-cancelDownload
//
// This CDP method is experimental.
func NewCancelDownload(guid string) *CancelDownload {
	return &CancelDownload{
		Guid: guid,
	}
}

// SetBrowserContextID adds or modifies the value of the optional
// parameter `browserContextId` in the CancelDownload CDP command.
//
// BrowserContext to perform the action in. When omitted, default browser context is used.
func (t *CancelDownload) SetBrowserContextID(v BrowserContextID) *CancelDownload {
	t.BrowserContextID = &v
	return t
}

// Do sends the CancelDownload CDP command to a browser,
// and returns the browser's response.
func (t *CancelDownload) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Browser.cancelDownload", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// Close contains the parameters, and acts as
// a Go receiver, for the CDP command `close`.
//
// Close browser gracefully.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-close
type Close struct{}

// NewClose constructs a new Close struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-close
func NewClose() *Close {
	return &Close{}
}

// Do sends the Close CDP command to a browser,
// and returns the browser's response.
func (t *Close) Do(ctx context.Context) error {
	response, err := cdp.Send(ctx, "Browser.close", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// Crash contains the parameters, and acts as
// a Go receiver, for the CDP command `crash`.
//
// Crashes browser on the main thread.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-crash
//
// This CDP method is experimental.
type Crash struct{}

// NewCrash constructs a new Crash struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-crash
//
// This CDP method is experimental.
func NewCrash() *Crash {
	return &Crash{}
}

// Do sends the Crash CDP command to a browser,
// and returns the browser's response.
func (t *Crash) Do(ctx context.Context) error {
	response, err := cdp.Send(ctx, "Browser.crash", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// CrashGpuProcess contains the parameters, and acts as
// a Go receiver, for the CDP command `crashGpuProcess`.
//
// Crashes GPU process.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-crashGpuProcess
//
// This CDP method is experimental.
type CrashGpuProcess struct{}

// NewCrashGpuProcess constructs a new CrashGpuProcess struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-crashGpuProcess
//
// This CDP method is experimental.
func NewCrashGpuProcess() *CrashGpuProcess {
	return &CrashGpuProcess{}
}

// Do sends the CrashGpuProcess CDP command to a browser,
// and returns the browser's response.
func (t *CrashGpuProcess) Do(ctx context.Context) error {
	response, err := cdp.Send(ctx, "Browser.crashGpuProcess", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// GetVersion contains the parameters, and acts as
// a Go receiver, for the CDP command `getVersion`.
//
// Returns version information.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-getVersion
type GetVersion struct{}

// NewGetVersion constructs a new GetVersion struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-getVersion
func NewGetVersion() *GetVersion {
	return &GetVersion{}
}

// GetVersionResponse contains the browser's response
// to calling the GetVersion CDP command with Do().
type GetVersionResponse struct {
	// Protocol version.
	ProtocolVersion string `json:"protocolVersion"`
	// Product name.
	Product string `json:"product"`
	// Product revision.
	Revision string `json:"revision"`
	// User-Agent.
	UserAgent string `json:"userAgent"`
	// V8 version.
	JsVersion string `json:"jsVersion"`
}

// Do sends the GetVersion CDP command to a browser,
// and returns the browser's response.
func (t *GetVersion) Do(ctx context.Context) (*GetVersionResponse, error) {
	response, err := cdp.Send(ctx, "Browser.getVersion", nil)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetVersionResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetBrowserCommandLine contains the parameters, and acts as
// a Go receiver, for the CDP command `getBrowserCommandLine`.
//
// Returns the command line switches for the browser process if, and only if
// --enable-automation is on the commandline.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-getBrowserCommandLine
//
// This CDP method is experimental.
type GetBrowserCommandLine struct{}

// NewGetBrowserCommandLine constructs a new GetBrowserCommandLine struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-getBrowserCommandLine
//
// This CDP method is experimental.
func NewGetBrowserCommandLine() *GetBrowserCommandLine {
	return &GetBrowserCommandLine{}
}

// GetBrowserCommandLineResponse contains the browser's response
// to calling the GetBrowserCommandLine CDP command with Do().
type GetBrowserCommandLineResponse struct {
	// Commandline parameters
	Arguments []string `json:"arguments"`
}

// Do sends the GetBrowserCommandLine CDP command to a browser,
// and returns the browser's response.
func (t *GetBrowserCommandLine) Do(ctx context.Context) (*GetBrowserCommandLineResponse, error) {
	response, err := cdp.Send(ctx, "Browser.getBrowserCommandLine", nil)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetBrowserCommandLineResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetHistograms contains the parameters, and acts as
// a Go receiver, for the CDP command `getHistograms`.
//
// Get Chrome histograms.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-getHistograms
//
// This CDP method is experimental.
type GetHistograms struct {
	// Requested substring in name. Only histograms which have query as a
	// substring in their name are extracted. An empty or absent query returns
	// all histograms.
	Query string `json:"query,omitempty"`
	// If true, retrieve delta since last call.
	Delta bool `json:"delta,omitempty"`
}

// NewGetHistograms constructs a new GetHistograms struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-getHistograms
//
// This CDP method is experimental.
func NewGetHistograms() *GetHistograms {
	return &GetHistograms{}
}

// SetQuery adds or modifies the value of the optional
// parameter `query` in the GetHistograms CDP command.
//
// Requested substring in name. Only histograms which have query as a
// substring in their name are extracted. An empty or absent query returns
// all histograms.
func (t *GetHistograms) SetQuery(v string) *GetHistograms {
	t.Query = v
	return t
}

// SetDelta adds or modifies the value of the optional
// parameter `delta` in the GetHistograms CDP command.
//
// If true, retrieve delta since last call.
func (t *GetHistograms) SetDelta(v bool) *GetHistograms {
	t.Delta = v
	return t
}

// GetHistogramsResponse contains the browser's response
// to calling the GetHistograms CDP command with Do().
type GetHistogramsResponse struct {
	// Histograms.
	Histograms []Histogram `json:"histograms"`
}

// Do sends the GetHistograms CDP command to a browser,
// and returns the browser's response.
func (t *GetHistograms) Do(ctx context.Context) (*GetHistogramsResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "Browser.getHistograms", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetHistogramsResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetHistogram contains the parameters, and acts as
// a Go receiver, for the CDP command `getHistogram`.
//
// Get a Chrome histogram by name.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-getHistogram
//
// This CDP method is experimental.
type GetHistogram struct {
	// Requested histogram name.
	Name string `json:"name"`
	// If true, retrieve delta since last call.
	Delta bool `json:"delta,omitempty"`
}

// NewGetHistogram constructs a new GetHistogram struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-getHistogram
//
// This CDP method is experimental.
func NewGetHistogram(name string) *GetHistogram {
	return &GetHistogram{
		Name: name,
	}
}

// SetDelta adds or modifies the value of the optional
// parameter `delta` in the GetHistogram CDP command.
//
// If true, retrieve delta since last call.
func (t *GetHistogram) SetDelta(v bool) *GetHistogram {
	t.Delta = v
	return t
}

// GetHistogramResponse contains the browser's response
// to calling the GetHistogram CDP command with Do().
type GetHistogramResponse struct {
	// Histogram.
	Histogram Histogram `json:"histogram"`
}

// Do sends the GetHistogram CDP command to a browser,
// and returns the browser's response.
func (t *GetHistogram) Do(ctx context.Context) (*GetHistogramResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "Browser.getHistogram", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetHistogramResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetWindowBounds contains the parameters, and acts as
// a Go receiver, for the CDP command `getWindowBounds`.
//
// Get position and size of the browser window.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-getWindowBounds
//
// This CDP method is experimental.
type GetWindowBounds struct {
	// Browser window id.
	WindowID WindowID `json:"windowId"`
}

// NewGetWindowBounds constructs a new GetWindowBounds struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-getWindowBounds
//
// This CDP method is experimental.
func NewGetWindowBounds(windowId WindowID) *GetWindowBounds {
	return &GetWindowBounds{
		WindowID: windowId,
	}
}

// GetWindowBoundsResponse contains the browser's response
// to calling the GetWindowBounds CDP command with Do().
type GetWindowBoundsResponse struct {
	// Bounds information of the window. When window state is 'minimized', the restored window
	// position and size are returned.
	Bounds Bounds `json:"bounds"`
}

// Do sends the GetWindowBounds CDP command to a browser,
// and returns the browser's response.
func (t *GetWindowBounds) Do(ctx context.Context) (*GetWindowBoundsResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "Browser.getWindowBounds", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetWindowBoundsResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetWindowForTarget contains the parameters, and acts as
// a Go receiver, for the CDP command `getWindowForTarget`.
//
// Get the browser window that contains the devtools target.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-getWindowForTarget
//
// This CDP method is experimental.
type GetWindowForTarget struct {
	// Devtools agent host id. If called as a part of the session, associated targetId is used.
	TargetID *cdp.TargetID `json:"targetId,omitempty"`
}

// NewGetWindowForTarget constructs a new GetWindowForTarget struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-getWindowForTarget
//
// This CDP method is experimental.
func NewGetWindowForTarget() *GetWindowForTarget {
	return &GetWindowForTarget{}
}

// SetTargetID adds or modifies the value of the optional
// parameter `targetId` in the GetWindowForTarget CDP command.
//
// Devtools agent host id. If called as a part of the session, associated targetId is used.
func (t *GetWindowForTarget) SetTargetID(v cdp.TargetID) *GetWindowForTarget {
	t.TargetID = &v
	return t
}

// GetWindowForTargetResponse contains the browser's response
// to calling the GetWindowForTarget CDP command with Do().
type GetWindowForTargetResponse struct {
	// Browser window id.
	WindowID WindowID `json:"windowId"`
	// Bounds information of the window. When window state is 'minimized', the restored window
	// position and size are returned.
	Bounds Bounds `json:"bounds"`
}

// Do sends the GetWindowForTarget CDP command to a browser,
// and returns the browser's response.
func (t *GetWindowForTarget) Do(ctx context.Context) (*GetWindowForTargetResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "Browser.getWindowForTarget", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetWindowForTargetResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetWindowBounds contains the parameters, and acts as
// a Go receiver, for the CDP command `setWindowBounds`.
//
// Set position and/or size of the browser window.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-setWindowBounds
//
// This CDP method is experimental.
type SetWindowBounds struct {
	// Browser window id.
	WindowID WindowID `json:"windowId"`
	// New window bounds. The 'minimized', 'maximized' and 'fullscreen' states cannot be combined
	// with 'left', 'top', 'width' or 'height'. Leaves unspecified fields unchanged.
	Bounds Bounds `json:"bounds"`
}

// NewSetWindowBounds constructs a new SetWindowBounds struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-setWindowBounds
//
// This CDP method is experimental.
func NewSetWindowBounds(windowId WindowID, bounds Bounds) *SetWindowBounds {
	return &SetWindowBounds{
		WindowID: windowId,
		Bounds:   bounds,
	}
}

// Do sends the SetWindowBounds CDP command to a browser,
// and returns the browser's response.
func (t *SetWindowBounds) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Browser.setWindowBounds", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetDockTile contains the parameters, and acts as
// a Go receiver, for the CDP command `setDockTile`.
//
// Set dock tile details, platform-specific.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-setDockTile
//
// This CDP method is experimental.
type SetDockTile struct {
	BadgeLabel string `json:"badgeLabel,omitempty"`
	// Png encoded image. (Encoded as a base64 string when passed over JSON)
	Image string `json:"image,omitempty"`
}

// NewSetDockTile constructs a new SetDockTile struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-setDockTile
//
// This CDP method is experimental.
func NewSetDockTile() *SetDockTile {
	return &SetDockTile{}
}

// SetBadgeLabel adds or modifies the value of the optional
// parameter `badgeLabel` in the SetDockTile CDP command.
func (t *SetDockTile) SetBadgeLabel(v string) *SetDockTile {
	t.BadgeLabel = v
	return t
}

// SetImage adds or modifies the value of the optional
// parameter `image` in the SetDockTile CDP command.
//
// Png encoded image. (Encoded as a base64 string when passed over JSON)
func (t *SetDockTile) SetImage(v string) *SetDockTile {
	t.Image = v
	return t
}

// Do sends the SetDockTile CDP command to a browser,
// and returns the browser's response.
func (t *SetDockTile) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Browser.setDockTile", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// ExecuteBrowserCommand contains the parameters, and acts as
// a Go receiver, for the CDP command `executeBrowserCommand`.
//
// Invoke custom browser commands used by telemetry.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-executeBrowserCommand
//
// This CDP method is experimental.
type ExecuteBrowserCommand struct {
	CommandID BrowserCommandID `json:"commandId"`
}

// NewExecuteBrowserCommand constructs a new ExecuteBrowserCommand struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-executeBrowserCommand
//
// This CDP method is experimental.
func NewExecuteBrowserCommand(commandId BrowserCommandID) *ExecuteBrowserCommand {
	return &ExecuteBrowserCommand{
		CommandID: commandId,
	}
}

// Do sends the ExecuteBrowserCommand CDP command to a browser,
// and returns the browser's response.
func (t *ExecuteBrowserCommand) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Browser.executeBrowserCommand", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}
