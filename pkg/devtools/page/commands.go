package page

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
	"github.com/daabr/chrome-vision/pkg/devtools/debugger"
	"github.com/daabr/chrome-vision/pkg/devtools/dom"
	"github.com/daabr/chrome-vision/pkg/devtools/runtime"
)

// AddScriptToEvaluateOnLoad contains the parameters, and acts as
// a Go receiver, for the CDP command `addScriptToEvaluateOnLoad`.
//
// Deprecated, please use addScriptToEvaluateOnNewDocument instead.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-addScriptToEvaluateOnLoad
//
// This CDP method is deprecated.
// This CDP method is experimental.
type AddScriptToEvaluateOnLoad struct {
	ScriptSource string `json:"scriptSource"`
}

// NewAddScriptToEvaluateOnLoad constructs a new AddScriptToEvaluateOnLoad struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-addScriptToEvaluateOnLoad
//
// This CDP method is deprecated.
// This CDP method is experimental.
func NewAddScriptToEvaluateOnLoad(scriptSource string) *AddScriptToEvaluateOnLoad {
	return &AddScriptToEvaluateOnLoad{
		ScriptSource: scriptSource,
	}
}

// AddScriptToEvaluateOnLoadResult contains the browser's response
// to calling the AddScriptToEvaluateOnLoad CDP command with Do().
type AddScriptToEvaluateOnLoadResult struct {
	// Identifier of the added script.
	Identifier string `json:"identifier"`
}

// Do sends the AddScriptToEvaluateOnLoad CDP command to a browser,
// and returns the browser's response.
func (t *AddScriptToEvaluateOnLoad) Do(ctx context.Context) (*AddScriptToEvaluateOnLoadResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Page.addScriptToEvaluateOnLoad", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the AddScriptToEvaluateOnLoad CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *AddScriptToEvaluateOnLoad) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Page.addScriptToEvaluateOnLoad", b)
}

// ParseResponse parses the browser's response
// to the AddScriptToEvaluateOnLoad CDP command.
func (t *AddScriptToEvaluateOnLoad) ParseResponse(m *devtools.Message) (*AddScriptToEvaluateOnLoadResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &AddScriptToEvaluateOnLoadResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// AddScriptToEvaluateOnNewDocument contains the parameters, and acts as
// a Go receiver, for the CDP command `addScriptToEvaluateOnNewDocument`.
//
// Evaluates given script in every frame upon creation (before loading frame's scripts).
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-addScriptToEvaluateOnNewDocument
type AddScriptToEvaluateOnNewDocument struct {
	Source string `json:"source"`
	// If specified, creates an isolated world with the given name and evaluates given script in it.
	// This world name will be used as the ExecutionContextDescription::name when the corresponding
	// event is emitted.
	//
	// This CDP parameter is experimental.
	WorldName string `json:"worldName,omitempty"`
	// Specifies whether command line API should be available to the script, defaults
	// to false.
	//
	// This CDP parameter is experimental.
	IncludeCommandLineAPI bool `json:"includeCommandLineAPI,omitempty"`
}

// NewAddScriptToEvaluateOnNewDocument constructs a new AddScriptToEvaluateOnNewDocument struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-addScriptToEvaluateOnNewDocument
func NewAddScriptToEvaluateOnNewDocument(source string) *AddScriptToEvaluateOnNewDocument {
	return &AddScriptToEvaluateOnNewDocument{
		Source: source,
	}
}

// SetWorldName adds or modifies the value of the optional
// parameter `worldName` in the AddScriptToEvaluateOnNewDocument CDP command.
//
// If specified, creates an isolated world with the given name and evaluates given script in it.
// This world name will be used as the ExecutionContextDescription::name when the corresponding
// event is emitted.
//
// This CDP parameter is experimental.
func (t *AddScriptToEvaluateOnNewDocument) SetWorldName(v string) *AddScriptToEvaluateOnNewDocument {
	t.WorldName = v
	return t
}

// SetIncludeCommandLineAPI adds or modifies the value of the optional
// parameter `includeCommandLineAPI` in the AddScriptToEvaluateOnNewDocument CDP command.
//
// Specifies whether command line API should be available to the script, defaults
// to false.
//
// This CDP parameter is experimental.
func (t *AddScriptToEvaluateOnNewDocument) SetIncludeCommandLineAPI(v bool) *AddScriptToEvaluateOnNewDocument {
	t.IncludeCommandLineAPI = v
	return t
}

// AddScriptToEvaluateOnNewDocumentResult contains the browser's response
// to calling the AddScriptToEvaluateOnNewDocument CDP command with Do().
type AddScriptToEvaluateOnNewDocumentResult struct {
	// Identifier of the added script.
	Identifier string `json:"identifier"`
}

// Do sends the AddScriptToEvaluateOnNewDocument CDP command to a browser,
// and returns the browser's response.
func (t *AddScriptToEvaluateOnNewDocument) Do(ctx context.Context) (*AddScriptToEvaluateOnNewDocumentResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Page.addScriptToEvaluateOnNewDocument", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the AddScriptToEvaluateOnNewDocument CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *AddScriptToEvaluateOnNewDocument) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Page.addScriptToEvaluateOnNewDocument", b)
}

// ParseResponse parses the browser's response
// to the AddScriptToEvaluateOnNewDocument CDP command.
func (t *AddScriptToEvaluateOnNewDocument) ParseResponse(m *devtools.Message) (*AddScriptToEvaluateOnNewDocumentResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &AddScriptToEvaluateOnNewDocumentResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// BringToFront contains the parameters, and acts as
// a Go receiver, for the CDP command `bringToFront`.
//
// Brings page to front (activates tab).
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-bringToFront
type BringToFront struct{}

// NewBringToFront constructs a new BringToFront struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-bringToFront
func NewBringToFront() *BringToFront {
	return &BringToFront{}
}

// Do sends the BringToFront CDP command to a browser,
// and returns the browser's response.
func (t *BringToFront) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Page.bringToFront", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the BringToFront CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *BringToFront) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Page.bringToFront", nil)
}

// ParseResponse parses the browser's response
// to the BringToFront CDP command.
func (t *BringToFront) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// CaptureScreenshot contains the parameters, and acts as
// a Go receiver, for the CDP command `captureScreenshot`.
//
// Capture page screenshot.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-captureScreenshot
type CaptureScreenshot struct {
	// Image compression format (defaults to png).
	Format string `json:"format,omitempty"`
	// Compression quality from range [0..100] (jpeg only).
	Quality int64 `json:"quality,omitempty"`
	// Capture the screenshot of a given region only.
	Clip *Viewport `json:"clip,omitempty"`
	// Capture the screenshot from the surface, rather than the view. Defaults to true.
	//
	// This CDP parameter is experimental.
	FromSurface bool `json:"fromSurface,omitempty"`
	// Capture the screenshot beyond the viewport. Defaults to false.
	//
	// This CDP parameter is experimental.
	CaptureBeyondViewport bool `json:"captureBeyondViewport,omitempty"`
}

// NewCaptureScreenshot constructs a new CaptureScreenshot struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-captureScreenshot
func NewCaptureScreenshot() *CaptureScreenshot {
	return &CaptureScreenshot{}
}

// SetFormat adds or modifies the value of the optional
// parameter `format` in the CaptureScreenshot CDP command.
//
// Image compression format (defaults to png).
func (t *CaptureScreenshot) SetFormat(v string) *CaptureScreenshot {
	t.Format = v
	return t
}

// SetQuality adds or modifies the value of the optional
// parameter `quality` in the CaptureScreenshot CDP command.
//
// Compression quality from range [0..100] (jpeg only).
func (t *CaptureScreenshot) SetQuality(v int64) *CaptureScreenshot {
	t.Quality = v
	return t
}

// SetClip adds or modifies the value of the optional
// parameter `clip` in the CaptureScreenshot CDP command.
//
// Capture the screenshot of a given region only.
func (t *CaptureScreenshot) SetClip(v Viewport) *CaptureScreenshot {
	t.Clip = &v
	return t
}

// SetFromSurface adds or modifies the value of the optional
// parameter `fromSurface` in the CaptureScreenshot CDP command.
//
// Capture the screenshot from the surface, rather than the view. Defaults to true.
//
// This CDP parameter is experimental.
func (t *CaptureScreenshot) SetFromSurface(v bool) *CaptureScreenshot {
	t.FromSurface = v
	return t
}

// SetCaptureBeyondViewport adds or modifies the value of the optional
// parameter `captureBeyondViewport` in the CaptureScreenshot CDP command.
//
// Capture the screenshot beyond the viewport. Defaults to false.
//
// This CDP parameter is experimental.
func (t *CaptureScreenshot) SetCaptureBeyondViewport(v bool) *CaptureScreenshot {
	t.CaptureBeyondViewport = v
	return t
}

// CaptureScreenshotResult contains the browser's response
// to calling the CaptureScreenshot CDP command with Do().
type CaptureScreenshotResult struct {
	// Base64-encoded image data. (Encoded as a base64 string when passed over JSON)
	Data string `json:"data"`
}

// Do sends the CaptureScreenshot CDP command to a browser,
// and returns the browser's response.
func (t *CaptureScreenshot) Do(ctx context.Context) (*CaptureScreenshotResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Page.captureScreenshot", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the CaptureScreenshot CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *CaptureScreenshot) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Page.captureScreenshot", b)
}

// ParseResponse parses the browser's response
// to the CaptureScreenshot CDP command.
func (t *CaptureScreenshot) ParseResponse(m *devtools.Message) (*CaptureScreenshotResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &CaptureScreenshotResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// CaptureSnapshot contains the parameters, and acts as
// a Go receiver, for the CDP command `captureSnapshot`.
//
// Returns a snapshot of the page as a string. For MHTML format, the serialization includes
// iframes, shadow DOM, external resources, and element-inline styles.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-captureSnapshot
//
// This CDP method is experimental.
type CaptureSnapshot struct {
	// Format (defaults to mhtml).
	Format string `json:"format,omitempty"`
}

// NewCaptureSnapshot constructs a new CaptureSnapshot struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-captureSnapshot
//
// This CDP method is experimental.
func NewCaptureSnapshot() *CaptureSnapshot {
	return &CaptureSnapshot{}
}

// SetFormat adds or modifies the value of the optional
// parameter `format` in the CaptureSnapshot CDP command.
//
// Format (defaults to mhtml).
func (t *CaptureSnapshot) SetFormat(v string) *CaptureSnapshot {
	t.Format = v
	return t
}

// CaptureSnapshotResult contains the browser's response
// to calling the CaptureSnapshot CDP command with Do().
type CaptureSnapshotResult struct {
	// Serialized page data.
	Data string `json:"data"`
}

// Do sends the CaptureSnapshot CDP command to a browser,
// and returns the browser's response.
func (t *CaptureSnapshot) Do(ctx context.Context) (*CaptureSnapshotResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Page.captureSnapshot", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the CaptureSnapshot CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *CaptureSnapshot) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Page.captureSnapshot", b)
}

// ParseResponse parses the browser's response
// to the CaptureSnapshot CDP command.
func (t *CaptureSnapshot) ParseResponse(m *devtools.Message) (*CaptureSnapshotResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &CaptureSnapshotResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateIsolatedWorld contains the parameters, and acts as
// a Go receiver, for the CDP command `createIsolatedWorld`.
//
// Creates an isolated world for the given frame.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-createIsolatedWorld
type CreateIsolatedWorld struct {
	// Id of the frame in which the isolated world should be created.
	FrameID string `json:"frameId"`
	// An optional name which is reported in the Execution Context.
	WorldName string `json:"worldName,omitempty"`
	// Whether or not universal access should be granted to the isolated world. This is a powerful
	// option, use with caution.
	GrantUniveralAccess bool `json:"grantUniveralAccess,omitempty"`
}

// NewCreateIsolatedWorld constructs a new CreateIsolatedWorld struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-createIsolatedWorld
func NewCreateIsolatedWorld(frameID string) *CreateIsolatedWorld {
	return &CreateIsolatedWorld{
		FrameID: frameID,
	}
}

// SetWorldName adds or modifies the value of the optional
// parameter `worldName` in the CreateIsolatedWorld CDP command.
//
// An optional name which is reported in the Execution Context.
func (t *CreateIsolatedWorld) SetWorldName(v string) *CreateIsolatedWorld {
	t.WorldName = v
	return t
}

// SetGrantUniveralAccess adds or modifies the value of the optional
// parameter `grantUniveralAccess` in the CreateIsolatedWorld CDP command.
//
// Whether or not universal access should be granted to the isolated world. This is a powerful
// option, use with caution.
func (t *CreateIsolatedWorld) SetGrantUniveralAccess(v bool) *CreateIsolatedWorld {
	t.GrantUniveralAccess = v
	return t
}

// CreateIsolatedWorldResult contains the browser's response
// to calling the CreateIsolatedWorld CDP command with Do().
type CreateIsolatedWorldResult struct {
	// Execution context of the isolated world.
	ExecutionContextID runtime.ExecutionContextID `json:"executionContextId"`
}

// Do sends the CreateIsolatedWorld CDP command to a browser,
// and returns the browser's response.
func (t *CreateIsolatedWorld) Do(ctx context.Context) (*CreateIsolatedWorldResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Page.createIsolatedWorld", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the CreateIsolatedWorld CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *CreateIsolatedWorld) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Page.createIsolatedWorld", b)
}

// ParseResponse parses the browser's response
// to the CreateIsolatedWorld CDP command.
func (t *CreateIsolatedWorld) ParseResponse(m *devtools.Message) (*CreateIsolatedWorldResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &CreateIsolatedWorldResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Disable contains the parameters, and acts as
// a Go receiver, for the CDP command `disable`.
//
// Disables page domain notifications.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-disable
type Disable struct{}

// NewDisable constructs a new Disable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-disable
func NewDisable() *Disable {
	return &Disable{}
}

// Do sends the Disable CDP command to a browser,
// and returns the browser's response.
func (t *Disable) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Page.disable", nil)
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
	return devtools.Send(ctx, "Page.disable", nil)
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
// Enables page domain notifications.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-enable
type Enable struct{}

// NewEnable constructs a new Enable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-enable
func NewEnable() *Enable {
	return &Enable{}
}

// Do sends the Enable CDP command to a browser,
// and returns the browser's response.
func (t *Enable) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Page.enable", nil)
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
	return devtools.Send(ctx, "Page.enable", nil)
}

// ParseResponse parses the browser's response
// to the Enable CDP command.
func (t *Enable) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// GetAppManifest contains the parameters, and acts as
// a Go receiver, for the CDP command `getAppManifest`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getAppManifest
type GetAppManifest struct{}

// NewGetAppManifest constructs a new GetAppManifest struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getAppManifest
func NewGetAppManifest() *GetAppManifest {
	return &GetAppManifest{}
}

// GetAppManifestResult contains the browser's response
// to calling the GetAppManifest CDP command with Do().
type GetAppManifestResult struct {
	// Manifest location.
	URL    string             `json:"url"`
	Errors []AppManifestError `json:"errors"`
	// Manifest content.
	Data string `json:"data,omitempty"`
	// Parsed manifest properties
	//
	// This CDP parameter is experimental.
	Parsed *AppManifestParsedProperties `json:"parsed,omitempty"`
}

// Do sends the GetAppManifest CDP command to a browser,
// and returns the browser's response.
func (t *GetAppManifest) Do(ctx context.Context) (*GetAppManifestResult, error) {
	m, err := devtools.SendAndWait(ctx, "Page.getAppManifest", nil)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetAppManifest CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetAppManifest) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Page.getAppManifest", nil)
}

// ParseResponse parses the browser's response
// to the GetAppManifest CDP command.
func (t *GetAppManifest) ParseResponse(m *devtools.Message) (*GetAppManifestResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetAppManifestResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetInstallabilityErrors contains the parameters, and acts as
// a Go receiver, for the CDP command `getInstallabilityErrors`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getInstallabilityErrors
//
// This CDP method is experimental.
type GetInstallabilityErrors struct{}

// NewGetInstallabilityErrors constructs a new GetInstallabilityErrors struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getInstallabilityErrors
//
// This CDP method is experimental.
func NewGetInstallabilityErrors() *GetInstallabilityErrors {
	return &GetInstallabilityErrors{}
}

// GetInstallabilityErrorsResult contains the browser's response
// to calling the GetInstallabilityErrors CDP command with Do().
type GetInstallabilityErrorsResult struct {
	InstallabilityErrors []InstallabilityError `json:"installabilityErrors"`
}

// Do sends the GetInstallabilityErrors CDP command to a browser,
// and returns the browser's response.
func (t *GetInstallabilityErrors) Do(ctx context.Context) (*GetInstallabilityErrorsResult, error) {
	m, err := devtools.SendAndWait(ctx, "Page.getInstallabilityErrors", nil)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetInstallabilityErrors CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetInstallabilityErrors) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Page.getInstallabilityErrors", nil)
}

// ParseResponse parses the browser's response
// to the GetInstallabilityErrors CDP command.
func (t *GetInstallabilityErrors) ParseResponse(m *devtools.Message) (*GetInstallabilityErrorsResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetInstallabilityErrorsResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetManifestIcons contains the parameters, and acts as
// a Go receiver, for the CDP command `getManifestIcons`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getManifestIcons
//
// This CDP method is experimental.
type GetManifestIcons struct{}

// NewGetManifestIcons constructs a new GetManifestIcons struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getManifestIcons
//
// This CDP method is experimental.
func NewGetManifestIcons() *GetManifestIcons {
	return &GetManifestIcons{}
}

// GetManifestIconsResult contains the browser's response
// to calling the GetManifestIcons CDP command with Do().
type GetManifestIconsResult struct {
	PrimaryIcon string `json:"primaryIcon,omitempty"`
}

// Do sends the GetManifestIcons CDP command to a browser,
// and returns the browser's response.
func (t *GetManifestIcons) Do(ctx context.Context) (*GetManifestIconsResult, error) {
	m, err := devtools.SendAndWait(ctx, "Page.getManifestIcons", nil)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetManifestIcons CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetManifestIcons) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Page.getManifestIcons", nil)
}

// ParseResponse parses the browser's response
// to the GetManifestIcons CDP command.
func (t *GetManifestIcons) ParseResponse(m *devtools.Message) (*GetManifestIconsResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetManifestIconsResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetAppID contains the parameters, and acts as
// a Go receiver, for the CDP command `getAppId`.
//
// Returns the unique (PWA) app id.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getAppId
//
// This CDP method is experimental.
type GetAppID struct{}

// NewGetAppID constructs a new GetAppID struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getAppId
//
// This CDP method is experimental.
func NewGetAppID() *GetAppID {
	return &GetAppID{}
}

// GetAppIDResult contains the browser's response
// to calling the GetAppID CDP command with Do().
type GetAppIDResult struct {
	// Only returns a value if the feature flag 'WebAppEnableManifestId' is enabled
	AppID string `json:"appId,omitempty"`
}

// Do sends the GetAppID CDP command to a browser,
// and returns the browser's response.
func (t *GetAppID) Do(ctx context.Context) (*GetAppIDResult, error) {
	m, err := devtools.SendAndWait(ctx, "Page.getAppId", nil)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetAppID CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetAppID) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Page.getAppId", nil)
}

// ParseResponse parses the browser's response
// to the GetAppID CDP command.
func (t *GetAppID) ParseResponse(m *devtools.Message) (*GetAppIDResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetAppIDResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetFrameTree contains the parameters, and acts as
// a Go receiver, for the CDP command `getFrameTree`.
//
// Returns present frame tree structure.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getFrameTree
type GetFrameTree struct{}

// NewGetFrameTree constructs a new GetFrameTree struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getFrameTree
func NewGetFrameTree() *GetFrameTree {
	return &GetFrameTree{}
}

// GetFrameTreeResult contains the browser's response
// to calling the GetFrameTree CDP command with Do().
type GetFrameTreeResult struct {
	// Present frame tree structure.
	FrameTree FrameTree `json:"frameTree"`
}

// Do sends the GetFrameTree CDP command to a browser,
// and returns the browser's response.
func (t *GetFrameTree) Do(ctx context.Context) (*GetFrameTreeResult, error) {
	m, err := devtools.SendAndWait(ctx, "Page.getFrameTree", nil)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetFrameTree CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetFrameTree) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Page.getFrameTree", nil)
}

// ParseResponse parses the browser's response
// to the GetFrameTree CDP command.
func (t *GetFrameTree) ParseResponse(m *devtools.Message) (*GetFrameTreeResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetFrameTreeResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetLayoutMetrics contains the parameters, and acts as
// a Go receiver, for the CDP command `getLayoutMetrics`.
//
// Returns metrics relating to the layouting of the page, such as viewport bounds/scale.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getLayoutMetrics
type GetLayoutMetrics struct{}

// NewGetLayoutMetrics constructs a new GetLayoutMetrics struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getLayoutMetrics
func NewGetLayoutMetrics() *GetLayoutMetrics {
	return &GetLayoutMetrics{}
}

// GetLayoutMetricsResult contains the browser's response
// to calling the GetLayoutMetrics CDP command with Do().
type GetLayoutMetricsResult struct {
	// Deprecated metrics relating to the layout viewport. Can be in DP or in CSS pixels depending on the `enable-use-zoom-for-dsf` flag. Use `cssLayoutViewport` instead.
	//
	// This CDP parameter is deprecated.
	LayoutViewport LayoutViewport `json:"layoutViewport"`
	// Deprecated metrics relating to the visual viewport. Can be in DP or in CSS pixels depending on the `enable-use-zoom-for-dsf` flag. Use `cssVisualViewport` instead.
	//
	// This CDP parameter is deprecated.
	VisualViewport VisualViewport `json:"visualViewport"`
	// Deprecated size of scrollable area. Can be in DP or in CSS pixels depending on the `enable-use-zoom-for-dsf` flag. Use `cssContentSize` instead.
	//
	// This CDP parameter is deprecated.
	ContentSize dom.Rect `json:"contentSize"`
	// Metrics relating to the layout viewport in CSS pixels.
	CSSLayoutViewport LayoutViewport `json:"cssLayoutViewport"`
	// Metrics relating to the visual viewport in CSS pixels.
	CSSVisualViewport VisualViewport `json:"cssVisualViewport"`
	// Size of scrollable area in CSS pixels.
	CSSContentSize dom.Rect `json:"cssContentSize"`
}

// Do sends the GetLayoutMetrics CDP command to a browser,
// and returns the browser's response.
func (t *GetLayoutMetrics) Do(ctx context.Context) (*GetLayoutMetricsResult, error) {
	m, err := devtools.SendAndWait(ctx, "Page.getLayoutMetrics", nil)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetLayoutMetrics CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetLayoutMetrics) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Page.getLayoutMetrics", nil)
}

// ParseResponse parses the browser's response
// to the GetLayoutMetrics CDP command.
func (t *GetLayoutMetrics) ParseResponse(m *devtools.Message) (*GetLayoutMetricsResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetLayoutMetricsResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetNavigationHistory contains the parameters, and acts as
// a Go receiver, for the CDP command `getNavigationHistory`.
//
// Returns navigation history for the current page.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getNavigationHistory
type GetNavigationHistory struct{}

// NewGetNavigationHistory constructs a new GetNavigationHistory struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getNavigationHistory
func NewGetNavigationHistory() *GetNavigationHistory {
	return &GetNavigationHistory{}
}

// GetNavigationHistoryResult contains the browser's response
// to calling the GetNavigationHistory CDP command with Do().
type GetNavigationHistoryResult struct {
	// Index of the current navigation history entry.
	CurrentIndex int64 `json:"currentIndex"`
	// Array of navigation history entries.
	Entries []NavigationEntry `json:"entries"`
}

// Do sends the GetNavigationHistory CDP command to a browser,
// and returns the browser's response.
func (t *GetNavigationHistory) Do(ctx context.Context) (*GetNavigationHistoryResult, error) {
	m, err := devtools.SendAndWait(ctx, "Page.getNavigationHistory", nil)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetNavigationHistory CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetNavigationHistory) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Page.getNavigationHistory", nil)
}

// ParseResponse parses the browser's response
// to the GetNavigationHistory CDP command.
func (t *GetNavigationHistory) ParseResponse(m *devtools.Message) (*GetNavigationHistoryResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetNavigationHistoryResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ResetNavigationHistory contains the parameters, and acts as
// a Go receiver, for the CDP command `resetNavigationHistory`.
//
// Resets navigation history for the current page.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-resetNavigationHistory
type ResetNavigationHistory struct{}

// NewResetNavigationHistory constructs a new ResetNavigationHistory struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-resetNavigationHistory
func NewResetNavigationHistory() *ResetNavigationHistory {
	return &ResetNavigationHistory{}
}

// Do sends the ResetNavigationHistory CDP command to a browser,
// and returns the browser's response.
func (t *ResetNavigationHistory) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Page.resetNavigationHistory", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the ResetNavigationHistory CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *ResetNavigationHistory) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Page.resetNavigationHistory", nil)
}

// ParseResponse parses the browser's response
// to the ResetNavigationHistory CDP command.
func (t *ResetNavigationHistory) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// GetResourceContent contains the parameters, and acts as
// a Go receiver, for the CDP command `getResourceContent`.
//
// Returns content of the given resource.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getResourceContent
//
// This CDP method is experimental.
type GetResourceContent struct {
	// Frame id to get resource for.
	FrameID string `json:"frameId"`
	// URL of the resource to get content for.
	URL string `json:"url"`
}

// NewGetResourceContent constructs a new GetResourceContent struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getResourceContent
//
// This CDP method is experimental.
func NewGetResourceContent(frameID string, url string) *GetResourceContent {
	return &GetResourceContent{
		FrameID: frameID,
		URL:     url,
	}
}

// GetResourceContentResult contains the browser's response
// to calling the GetResourceContent CDP command with Do().
type GetResourceContentResult struct {
	// Resource content.
	Content string `json:"content"`
	// True, if content was served as base64.
	Base64Encoded bool `json:"base64Encoded"`
}

// Do sends the GetResourceContent CDP command to a browser,
// and returns the browser's response.
func (t *GetResourceContent) Do(ctx context.Context) (*GetResourceContentResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Page.getResourceContent", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetResourceContent CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetResourceContent) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Page.getResourceContent", b)
}

// ParseResponse parses the browser's response
// to the GetResourceContent CDP command.
func (t *GetResourceContent) ParseResponse(m *devtools.Message) (*GetResourceContentResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetResourceContentResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetResourceTree contains the parameters, and acts as
// a Go receiver, for the CDP command `getResourceTree`.
//
// Returns present frame / resource tree structure.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getResourceTree
//
// This CDP method is experimental.
type GetResourceTree struct{}

// NewGetResourceTree constructs a new GetResourceTree struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getResourceTree
//
// This CDP method is experimental.
func NewGetResourceTree() *GetResourceTree {
	return &GetResourceTree{}
}

// GetResourceTreeResult contains the browser's response
// to calling the GetResourceTree CDP command with Do().
type GetResourceTreeResult struct {
	// Present frame / resource tree structure.
	FrameTree FrameResourceTree `json:"frameTree"`
}

// Do sends the GetResourceTree CDP command to a browser,
// and returns the browser's response.
func (t *GetResourceTree) Do(ctx context.Context) (*GetResourceTreeResult, error) {
	m, err := devtools.SendAndWait(ctx, "Page.getResourceTree", nil)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetResourceTree CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetResourceTree) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Page.getResourceTree", nil)
}

// ParseResponse parses the browser's response
// to the GetResourceTree CDP command.
func (t *GetResourceTree) ParseResponse(m *devtools.Message) (*GetResourceTreeResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetResourceTreeResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// HandleJavaScriptDialog contains the parameters, and acts as
// a Go receiver, for the CDP command `handleJavaScriptDialog`.
//
// Accepts or dismisses a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload).
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-handleJavaScriptDialog
type HandleJavaScriptDialog struct {
	// Whether to accept or dismiss the dialog.
	Accept bool `json:"accept"`
	// The text to enter into the dialog prompt before accepting. Used only if this is a prompt
	// dialog.
	PromptText string `json:"promptText,omitempty"`
}

// NewHandleJavaScriptDialog constructs a new HandleJavaScriptDialog struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-handleJavaScriptDialog
func NewHandleJavaScriptDialog(accept bool) *HandleJavaScriptDialog {
	return &HandleJavaScriptDialog{
		Accept: accept,
	}
}

// SetPromptText adds or modifies the value of the optional
// parameter `promptText` in the HandleJavaScriptDialog CDP command.
//
// The text to enter into the dialog prompt before accepting. Used only if this is a prompt
// dialog.
func (t *HandleJavaScriptDialog) SetPromptText(v string) *HandleJavaScriptDialog {
	t.PromptText = v
	return t
}

// Do sends the HandleJavaScriptDialog CDP command to a browser,
// and returns the browser's response.
func (t *HandleJavaScriptDialog) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Page.handleJavaScriptDialog", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the HandleJavaScriptDialog CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *HandleJavaScriptDialog) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Page.handleJavaScriptDialog", b)
}

// ParseResponse parses the browser's response
// to the HandleJavaScriptDialog CDP command.
func (t *HandleJavaScriptDialog) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// Navigate contains the parameters, and acts as
// a Go receiver, for the CDP command `navigate`.
//
// Navigates current page to the given URL.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-navigate
type Navigate struct {
	// URL to navigate the page to.
	URL string `json:"url"`
	// Referrer URL.
	Referrer string `json:"referrer,omitempty"`
	// Intended transition type.
	TransitionType *TransitionType `json:"transitionType,omitempty"`
	// Frame id to navigate, if not specified navigates the top frame.
	FrameID string `json:"frameId,omitempty"`
	// Referrer-policy used for the navigation.
	//
	// This CDP parameter is experimental.
	ReferrerPolicy *ReferrerPolicy `json:"referrerPolicy,omitempty"`
}

// NewNavigate constructs a new Navigate struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-navigate
func NewNavigate(url string) *Navigate {
	return &Navigate{
		URL: url,
	}
}

// SetReferrer adds or modifies the value of the optional
// parameter `referrer` in the Navigate CDP command.
//
// Referrer URL.
func (t *Navigate) SetReferrer(v string) *Navigate {
	t.Referrer = v
	return t
}

// SetTransitionType adds or modifies the value of the optional
// parameter `transitionType` in the Navigate CDP command.
//
// Intended transition type.
func (t *Navigate) SetTransitionType(v TransitionType) *Navigate {
	t.TransitionType = &v
	return t
}

// SetFrameID adds or modifies the value of the optional
// parameter `frameId` in the Navigate CDP command.
//
// Frame id to navigate, if not specified navigates the top frame.
func (t *Navigate) SetFrameID(v string) *Navigate {
	t.FrameID = v
	return t
}

// SetReferrerPolicy adds or modifies the value of the optional
// parameter `referrerPolicy` in the Navigate CDP command.
//
// Referrer-policy used for the navigation.
//
// This CDP parameter is experimental.
func (t *Navigate) SetReferrerPolicy(v ReferrerPolicy) *Navigate {
	t.ReferrerPolicy = &v
	return t
}

// NavigateResult contains the browser's response
// to calling the Navigate CDP command with Do().
type NavigateResult struct {
	// Frame id that has navigated (or failed to navigate)
	FrameID string `json:"frameId"`
	// Loader identifier.
	LoaderID string `json:"loaderId,omitempty"`
	// User friendly error message, present if and only if navigation has failed.
	ErrorText string `json:"errorText,omitempty"`
}

// Do sends the Navigate CDP command to a browser,
// and returns the browser's response.
func (t *Navigate) Do(ctx context.Context) (*NavigateResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Page.navigate", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the Navigate CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *Navigate) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Page.navigate", b)
}

// ParseResponse parses the browser's response
// to the Navigate CDP command.
func (t *Navigate) ParseResponse(m *devtools.Message) (*NavigateResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &NavigateResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	if result.ErrorText != "" {
		return nil, errors.New(result.ErrorText)
	}
	return result, nil
}

// NavigateToHistoryEntry contains the parameters, and acts as
// a Go receiver, for the CDP command `navigateToHistoryEntry`.
//
// Navigates current page to the given history entry.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-navigateToHistoryEntry
type NavigateToHistoryEntry struct {
	// Unique id of the entry to navigate to.
	EntryID int64 `json:"entryId"`
}

// NewNavigateToHistoryEntry constructs a new NavigateToHistoryEntry struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-navigateToHistoryEntry
func NewNavigateToHistoryEntry(entryID int64) *NavigateToHistoryEntry {
	return &NavigateToHistoryEntry{
		EntryID: entryID,
	}
}

// Do sends the NavigateToHistoryEntry CDP command to a browser,
// and returns the browser's response.
func (t *NavigateToHistoryEntry) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Page.navigateToHistoryEntry", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the NavigateToHistoryEntry CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *NavigateToHistoryEntry) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Page.navigateToHistoryEntry", b)
}

// ParseResponse parses the browser's response
// to the NavigateToHistoryEntry CDP command.
func (t *NavigateToHistoryEntry) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// PrintToPDF contains the parameters, and acts as
// a Go receiver, for the CDP command `printToPDF`.
//
// Print page as PDF.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-printToPDF
type PrintToPDF struct {
	// Paper orientation. Defaults to false.
	Landscape bool `json:"landscape,omitempty"`
	// Display header and footer. Defaults to false.
	DisplayHeaderFooter bool `json:"displayHeaderFooter,omitempty"`
	// Print background graphics. Defaults to false.
	PrintBackground bool `json:"printBackground,omitempty"`
	// Scale of the webpage rendering. Defaults to 1.
	Scale float64 `json:"scale,omitempty"`
	// Paper width in inches. Defaults to 8.5 inches.
	PaperWidth float64 `json:"paperWidth,omitempty"`
	// Paper height in inches. Defaults to 11 inches.
	PaperHeight float64 `json:"paperHeight,omitempty"`
	// Top margin in inches. Defaults to 1cm (~0.4 inches).
	MarginTop float64 `json:"marginTop,omitempty"`
	// Bottom margin in inches. Defaults to 1cm (~0.4 inches).
	MarginBottom float64 `json:"marginBottom,omitempty"`
	// Left margin in inches. Defaults to 1cm (~0.4 inches).
	MarginLeft float64 `json:"marginLeft,omitempty"`
	// Right margin in inches. Defaults to 1cm (~0.4 inches).
	MarginRight float64 `json:"marginRight,omitempty"`
	// Paper ranges to print, e.g., '1-5, 8, 11-13'. Defaults to the empty string, which means
	// print all pages.
	PageRanges string `json:"pageRanges,omitempty"`
	// Whether to silently ignore invalid but successfully parsed page ranges, such as '3-2'.
	// Defaults to false.
	IgnoreInvalidPageRanges bool `json:"ignoreInvalidPageRanges,omitempty"`
	// HTML template for the print header. Should be valid HTML markup with following
	// classes used to inject printing values into them:
	// - `date`: formatted print date
	// - `title`: document title
	// - `url`: document location
	// - `pageNumber`: current page number
	// - `totalPages`: total pages in the document
	//
	// For example, `<span class=title></span>` would generate span containing the title.
	HeaderTemplate string `json:"headerTemplate,omitempty"`
	// HTML template for the print footer. Should use the same format as the `headerTemplate`.
	FooterTemplate string `json:"footerTemplate,omitempty"`
	// Whether or not to prefer page size as defined by css. Defaults to false,
	// in which case the content will be scaled to fit the paper size.
	PreferCSSPageSize bool `json:"preferCSSPageSize,omitempty"`
	// return as stream
	//
	// This CDP parameter is experimental.
	TransferMode string `json:"transferMode,omitempty"`
}

// NewPrintToPDF constructs a new PrintToPDF struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-printToPDF
func NewPrintToPDF() *PrintToPDF {
	return &PrintToPDF{}
}

// SetLandscape adds or modifies the value of the optional
// parameter `landscape` in the PrintToPDF CDP command.
//
// Paper orientation. Defaults to false.
func (t *PrintToPDF) SetLandscape(v bool) *PrintToPDF {
	t.Landscape = v
	return t
}

// SetDisplayHeaderFooter adds or modifies the value of the optional
// parameter `displayHeaderFooter` in the PrintToPDF CDP command.
//
// Display header and footer. Defaults to false.
func (t *PrintToPDF) SetDisplayHeaderFooter(v bool) *PrintToPDF {
	t.DisplayHeaderFooter = v
	return t
}

// SetPrintBackground adds or modifies the value of the optional
// parameter `printBackground` in the PrintToPDF CDP command.
//
// Print background graphics. Defaults to false.
func (t *PrintToPDF) SetPrintBackground(v bool) *PrintToPDF {
	t.PrintBackground = v
	return t
}

// SetScale adds or modifies the value of the optional
// parameter `scale` in the PrintToPDF CDP command.
//
// Scale of the webpage rendering. Defaults to 1.
func (t *PrintToPDF) SetScale(v float64) *PrintToPDF {
	t.Scale = v
	return t
}

// SetPaperWidth adds or modifies the value of the optional
// parameter `paperWidth` in the PrintToPDF CDP command.
//
// Paper width in inches. Defaults to 8.5 inches.
func (t *PrintToPDF) SetPaperWidth(v float64) *PrintToPDF {
	t.PaperWidth = v
	return t
}

// SetPaperHeight adds or modifies the value of the optional
// parameter `paperHeight` in the PrintToPDF CDP command.
//
// Paper height in inches. Defaults to 11 inches.
func (t *PrintToPDF) SetPaperHeight(v float64) *PrintToPDF {
	t.PaperHeight = v
	return t
}

// SetMarginTop adds or modifies the value of the optional
// parameter `marginTop` in the PrintToPDF CDP command.
//
// Top margin in inches. Defaults to 1cm (~0.4 inches).
func (t *PrintToPDF) SetMarginTop(v float64) *PrintToPDF {
	t.MarginTop = v
	return t
}

// SetMarginBottom adds or modifies the value of the optional
// parameter `marginBottom` in the PrintToPDF CDP command.
//
// Bottom margin in inches. Defaults to 1cm (~0.4 inches).
func (t *PrintToPDF) SetMarginBottom(v float64) *PrintToPDF {
	t.MarginBottom = v
	return t
}

// SetMarginLeft adds or modifies the value of the optional
// parameter `marginLeft` in the PrintToPDF CDP command.
//
// Left margin in inches. Defaults to 1cm (~0.4 inches).
func (t *PrintToPDF) SetMarginLeft(v float64) *PrintToPDF {
	t.MarginLeft = v
	return t
}

// SetMarginRight adds or modifies the value of the optional
// parameter `marginRight` in the PrintToPDF CDP command.
//
// Right margin in inches. Defaults to 1cm (~0.4 inches).
func (t *PrintToPDF) SetMarginRight(v float64) *PrintToPDF {
	t.MarginRight = v
	return t
}

// SetPageRanges adds or modifies the value of the optional
// parameter `pageRanges` in the PrintToPDF CDP command.
//
// Paper ranges to print, e.g., '1-5, 8, 11-13'. Defaults to the empty string, which means
// print all pages.
func (t *PrintToPDF) SetPageRanges(v string) *PrintToPDF {
	t.PageRanges = v
	return t
}

// SetIgnoreInvalidPageRanges adds or modifies the value of the optional
// parameter `ignoreInvalidPageRanges` in the PrintToPDF CDP command.
//
// Whether to silently ignore invalid but successfully parsed page ranges, such as '3-2'.
// Defaults to false.
func (t *PrintToPDF) SetIgnoreInvalidPageRanges(v bool) *PrintToPDF {
	t.IgnoreInvalidPageRanges = v
	return t
}

// SetHeaderTemplate adds or modifies the value of the optional
// parameter `headerTemplate` in the PrintToPDF CDP command.
//
// HTML template for the print header. Should be valid HTML markup with following
// classes used to inject printing values into them:
// - `date`: formatted print date
// - `title`: document title
// - `url`: document location
// - `pageNumber`: current page number
// - `totalPages`: total pages in the document
//
// For example, `<span class=title></span>` would generate span containing the title.
func (t *PrintToPDF) SetHeaderTemplate(v string) *PrintToPDF {
	t.HeaderTemplate = v
	return t
}

// SetFooterTemplate adds or modifies the value of the optional
// parameter `footerTemplate` in the PrintToPDF CDP command.
//
// HTML template for the print footer. Should use the same format as the `headerTemplate`.
func (t *PrintToPDF) SetFooterTemplate(v string) *PrintToPDF {
	t.FooterTemplate = v
	return t
}

// SetPreferCSSPageSize adds or modifies the value of the optional
// parameter `preferCSSPageSize` in the PrintToPDF CDP command.
//
// Whether or not to prefer page size as defined by css. Defaults to false,
// in which case the content will be scaled to fit the paper size.
func (t *PrintToPDF) SetPreferCSSPageSize(v bool) *PrintToPDF {
	t.PreferCSSPageSize = v
	return t
}

// SetTransferMode adds or modifies the value of the optional
// parameter `transferMode` in the PrintToPDF CDP command.
//
// return as stream
//
// This CDP parameter is experimental.
func (t *PrintToPDF) SetTransferMode(v string) *PrintToPDF {
	t.TransferMode = v
	return t
}

// PrintToPDFResult contains the browser's response
// to calling the PrintToPDF CDP command with Do().
type PrintToPDFResult struct {
	// Base64-encoded pdf data. Empty if |returnAsStream| is specified. (Encoded as a base64 string when passed over JSON)
	Data string `json:"data"`
	// A handle of the stream that holds resulting PDF data.
	//
	// This CDP parameter is experimental.
	Stream string `json:"stream,omitempty"`
}

// Do sends the PrintToPDF CDP command to a browser,
// and returns the browser's response.
func (t *PrintToPDF) Do(ctx context.Context) (*PrintToPDFResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Page.printToPDF", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the PrintToPDF CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *PrintToPDF) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Page.printToPDF", b)
}

// ParseResponse parses the browser's response
// to the PrintToPDF CDP command.
func (t *PrintToPDF) ParseResponse(m *devtools.Message) (*PrintToPDFResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &PrintToPDFResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Reload contains the parameters, and acts as
// a Go receiver, for the CDP command `reload`.
//
// Reloads given page optionally ignoring the cache.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-reload
type Reload struct {
	// If true, browser cache is ignored (as if the user pressed Shift+refresh).
	IgnoreCache bool `json:"ignoreCache,omitempty"`
	// If set, the script will be injected into all frames of the inspected page after reload.
	// Argument will be ignored if reloading dataURL origin.
	ScriptToEvaluateOnLoad string `json:"scriptToEvaluateOnLoad,omitempty"`
}

// NewReload constructs a new Reload struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-reload
func NewReload() *Reload {
	return &Reload{}
}

// SetIgnoreCache adds or modifies the value of the optional
// parameter `ignoreCache` in the Reload CDP command.
//
// If true, browser cache is ignored (as if the user pressed Shift+refresh).
func (t *Reload) SetIgnoreCache(v bool) *Reload {
	t.IgnoreCache = v
	return t
}

// SetScriptToEvaluateOnLoad adds or modifies the value of the optional
// parameter `scriptToEvaluateOnLoad` in the Reload CDP command.
//
// If set, the script will be injected into all frames of the inspected page after reload.
// Argument will be ignored if reloading dataURL origin.
func (t *Reload) SetScriptToEvaluateOnLoad(v string) *Reload {
	t.ScriptToEvaluateOnLoad = v
	return t
}

// Do sends the Reload CDP command to a browser,
// and returns the browser's response.
func (t *Reload) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Page.reload", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the Reload CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *Reload) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Page.reload", b)
}

// ParseResponse parses the browser's response
// to the Reload CDP command.
func (t *Reload) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// RemoveScriptToEvaluateOnLoad contains the parameters, and acts as
// a Go receiver, for the CDP command `removeScriptToEvaluateOnLoad`.
//
// Deprecated, please use removeScriptToEvaluateOnNewDocument instead.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-removeScriptToEvaluateOnLoad
//
// This CDP method is deprecated.
// This CDP method is experimental.
type RemoveScriptToEvaluateOnLoad struct {
	Identifier string `json:"identifier"`
}

// NewRemoveScriptToEvaluateOnLoad constructs a new RemoveScriptToEvaluateOnLoad struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-removeScriptToEvaluateOnLoad
//
// This CDP method is deprecated.
// This CDP method is experimental.
func NewRemoveScriptToEvaluateOnLoad(identifier string) *RemoveScriptToEvaluateOnLoad {
	return &RemoveScriptToEvaluateOnLoad{
		Identifier: identifier,
	}
}

// Do sends the RemoveScriptToEvaluateOnLoad CDP command to a browser,
// and returns the browser's response.
func (t *RemoveScriptToEvaluateOnLoad) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Page.removeScriptToEvaluateOnLoad", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the RemoveScriptToEvaluateOnLoad CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *RemoveScriptToEvaluateOnLoad) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Page.removeScriptToEvaluateOnLoad", b)
}

// ParseResponse parses the browser's response
// to the RemoveScriptToEvaluateOnLoad CDP command.
func (t *RemoveScriptToEvaluateOnLoad) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// RemoveScriptToEvaluateOnNewDocument contains the parameters, and acts as
// a Go receiver, for the CDP command `removeScriptToEvaluateOnNewDocument`.
//
// Removes given script from the list.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-removeScriptToEvaluateOnNewDocument
type RemoveScriptToEvaluateOnNewDocument struct {
	Identifier string `json:"identifier"`
}

// NewRemoveScriptToEvaluateOnNewDocument constructs a new RemoveScriptToEvaluateOnNewDocument struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-removeScriptToEvaluateOnNewDocument
func NewRemoveScriptToEvaluateOnNewDocument(identifier string) *RemoveScriptToEvaluateOnNewDocument {
	return &RemoveScriptToEvaluateOnNewDocument{
		Identifier: identifier,
	}
}

// Do sends the RemoveScriptToEvaluateOnNewDocument CDP command to a browser,
// and returns the browser's response.
func (t *RemoveScriptToEvaluateOnNewDocument) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Page.removeScriptToEvaluateOnNewDocument", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the RemoveScriptToEvaluateOnNewDocument CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *RemoveScriptToEvaluateOnNewDocument) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Page.removeScriptToEvaluateOnNewDocument", b)
}

// ParseResponse parses the browser's response
// to the RemoveScriptToEvaluateOnNewDocument CDP command.
func (t *RemoveScriptToEvaluateOnNewDocument) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// ScreencastFrameAck contains the parameters, and acts as
// a Go receiver, for the CDP command `screencastFrameAck`.
//
// Acknowledges that a screencast frame has been received by the frontend.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-screencastFrameAck
//
// This CDP method is experimental.
type ScreencastFrameAck struct {
	// Frame number.
	SessionID int64 `json:"sessionId"`
}

// NewScreencastFrameAck constructs a new ScreencastFrameAck struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-screencastFrameAck
//
// This CDP method is experimental.
func NewScreencastFrameAck(sessionID int64) *ScreencastFrameAck {
	return &ScreencastFrameAck{
		SessionID: sessionID,
	}
}

// Do sends the ScreencastFrameAck CDP command to a browser,
// and returns the browser's response.
func (t *ScreencastFrameAck) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Page.screencastFrameAck", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the ScreencastFrameAck CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *ScreencastFrameAck) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Page.screencastFrameAck", b)
}

// ParseResponse parses the browser's response
// to the ScreencastFrameAck CDP command.
func (t *ScreencastFrameAck) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SearchInResource contains the parameters, and acts as
// a Go receiver, for the CDP command `searchInResource`.
//
// Searches for given string in resource content.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-searchInResource
//
// This CDP method is experimental.
type SearchInResource struct {
	// Frame id for resource to search in.
	FrameID string `json:"frameId"`
	// URL of the resource to search in.
	URL string `json:"url"`
	// String to search for.
	Query string `json:"query"`
	// If true, search is case sensitive.
	CaseSensitive bool `json:"caseSensitive,omitempty"`
	// If true, treats string parameter as regex.
	IsRegex bool `json:"isRegex,omitempty"`
}

// NewSearchInResource constructs a new SearchInResource struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-searchInResource
//
// This CDP method is experimental.
func NewSearchInResource(frameID string, url string, query string) *SearchInResource {
	return &SearchInResource{
		FrameID: frameID,
		URL:     url,
		Query:   query,
	}
}

// SetCaseSensitive adds or modifies the value of the optional
// parameter `caseSensitive` in the SearchInResource CDP command.
//
// If true, search is case sensitive.
func (t *SearchInResource) SetCaseSensitive(v bool) *SearchInResource {
	t.CaseSensitive = v
	return t
}

// SetIsRegex adds or modifies the value of the optional
// parameter `isRegex` in the SearchInResource CDP command.
//
// If true, treats string parameter as regex.
func (t *SearchInResource) SetIsRegex(v bool) *SearchInResource {
	t.IsRegex = v
	return t
}

// SearchInResourceResult contains the browser's response
// to calling the SearchInResource CDP command with Do().
type SearchInResourceResult struct {
	// List of search matches.
	Result []debugger.SearchMatch `json:"result"`
}

// Do sends the SearchInResource CDP command to a browser,
// and returns the browser's response.
func (t *SearchInResource) Do(ctx context.Context) (*SearchInResourceResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Page.searchInResource", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the SearchInResource CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SearchInResource) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Page.searchInResource", b)
}

// ParseResponse parses the browser's response
// to the SearchInResource CDP command.
func (t *SearchInResource) ParseResponse(m *devtools.Message) (*SearchInResourceResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &SearchInResourceResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetAdBlockingEnabled contains the parameters, and acts as
// a Go receiver, for the CDP command `setAdBlockingEnabled`.
//
// Enable Chrome's experimental ad filter on all sites.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setAdBlockingEnabled
//
// This CDP method is experimental.
type SetAdBlockingEnabled struct {
	// Whether to block ads.
	Enabled bool `json:"enabled"`
}

// NewSetAdBlockingEnabled constructs a new SetAdBlockingEnabled struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setAdBlockingEnabled
//
// This CDP method is experimental.
func NewSetAdBlockingEnabled(enabled bool) *SetAdBlockingEnabled {
	return &SetAdBlockingEnabled{
		Enabled: enabled,
	}
}

// Do sends the SetAdBlockingEnabled CDP command to a browser,
// and returns the browser's response.
func (t *SetAdBlockingEnabled) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Page.setAdBlockingEnabled", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetAdBlockingEnabled CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetAdBlockingEnabled) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Page.setAdBlockingEnabled", b)
}

// ParseResponse parses the browser's response
// to the SetAdBlockingEnabled CDP command.
func (t *SetAdBlockingEnabled) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetBypassCSP contains the parameters, and acts as
// a Go receiver, for the CDP command `setBypassCSP`.
//
// Enable page Content Security Policy by-passing.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setBypassCSP
//
// This CDP method is experimental.
type SetBypassCSP struct {
	// Whether to bypass page CSP.
	Enabled bool `json:"enabled"`
}

// NewSetBypassCSP constructs a new SetBypassCSP struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setBypassCSP
//
// This CDP method is experimental.
func NewSetBypassCSP(enabled bool) *SetBypassCSP {
	return &SetBypassCSP{
		Enabled: enabled,
	}
}

// Do sends the SetBypassCSP CDP command to a browser,
// and returns the browser's response.
func (t *SetBypassCSP) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Page.setBypassCSP", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetBypassCSP CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetBypassCSP) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Page.setBypassCSP", b)
}

// ParseResponse parses the browser's response
// to the SetBypassCSP CDP command.
func (t *SetBypassCSP) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// GetPermissionsPolicyState contains the parameters, and acts as
// a Go receiver, for the CDP command `getPermissionsPolicyState`.
//
// Get Permissions Policy state on given frame.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getPermissionsPolicyState
//
// This CDP method is experimental.
type GetPermissionsPolicyState struct {
	FrameID string `json:"frameId"`
}

// NewGetPermissionsPolicyState constructs a new GetPermissionsPolicyState struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getPermissionsPolicyState
//
// This CDP method is experimental.
func NewGetPermissionsPolicyState(frameID string) *GetPermissionsPolicyState {
	return &GetPermissionsPolicyState{
		FrameID: frameID,
	}
}

// GetPermissionsPolicyStateResult contains the browser's response
// to calling the GetPermissionsPolicyState CDP command with Do().
type GetPermissionsPolicyStateResult struct {
	States []PermissionsPolicyFeatureState `json:"states"`
}

// Do sends the GetPermissionsPolicyState CDP command to a browser,
// and returns the browser's response.
func (t *GetPermissionsPolicyState) Do(ctx context.Context) (*GetPermissionsPolicyStateResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Page.getPermissionsPolicyState", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetPermissionsPolicyState CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetPermissionsPolicyState) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Page.getPermissionsPolicyState", b)
}

// ParseResponse parses the browser's response
// to the GetPermissionsPolicyState CDP command.
func (t *GetPermissionsPolicyState) ParseResponse(m *devtools.Message) (*GetPermissionsPolicyStateResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetPermissionsPolicyStateResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetOriginTrials contains the parameters, and acts as
// a Go receiver, for the CDP command `getOriginTrials`.
//
// Get Origin Trials on given frame.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getOriginTrials
//
// This CDP method is experimental.
type GetOriginTrials struct {
	FrameID string `json:"frameId"`
}

// NewGetOriginTrials constructs a new GetOriginTrials struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-getOriginTrials
//
// This CDP method is experimental.
func NewGetOriginTrials(frameID string) *GetOriginTrials {
	return &GetOriginTrials{
		FrameID: frameID,
	}
}

// GetOriginTrialsResult contains the browser's response
// to calling the GetOriginTrials CDP command with Do().
type GetOriginTrialsResult struct {
	OriginTrials []OriginTrial `json:"originTrials"`
}

// Do sends the GetOriginTrials CDP command to a browser,
// and returns the browser's response.
func (t *GetOriginTrials) Do(ctx context.Context) (*GetOriginTrialsResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Page.getOriginTrials", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetOriginTrials CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetOriginTrials) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Page.getOriginTrials", b)
}

// ParseResponse parses the browser's response
// to the GetOriginTrials CDP command.
func (t *GetOriginTrials) ParseResponse(m *devtools.Message) (*GetOriginTrialsResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetOriginTrialsResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetFontFamilies contains the parameters, and acts as
// a Go receiver, for the CDP command `setFontFamilies`.
//
// Set generic font families.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setFontFamilies
//
// This CDP method is experimental.
type SetFontFamilies struct {
	// Specifies font families to set. If a font family is not specified, it won't be changed.
	FontFamilies FontFamilies `json:"fontFamilies"`
}

// NewSetFontFamilies constructs a new SetFontFamilies struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setFontFamilies
//
// This CDP method is experimental.
func NewSetFontFamilies(fontFamilies FontFamilies) *SetFontFamilies {
	return &SetFontFamilies{
		FontFamilies: fontFamilies,
	}
}

// Do sends the SetFontFamilies CDP command to a browser,
// and returns the browser's response.
func (t *SetFontFamilies) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Page.setFontFamilies", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetFontFamilies CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetFontFamilies) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Page.setFontFamilies", b)
}

// ParseResponse parses the browser's response
// to the SetFontFamilies CDP command.
func (t *SetFontFamilies) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetFontSizes contains the parameters, and acts as
// a Go receiver, for the CDP command `setFontSizes`.
//
// Set default font sizes.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setFontSizes
//
// This CDP method is experimental.
type SetFontSizes struct {
	// Specifies font sizes to set. If a font size is not specified, it won't be changed.
	FontSizes FontSizes `json:"fontSizes"`
}

// NewSetFontSizes constructs a new SetFontSizes struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setFontSizes
//
// This CDP method is experimental.
func NewSetFontSizes(fontSizes FontSizes) *SetFontSizes {
	return &SetFontSizes{
		FontSizes: fontSizes,
	}
}

// Do sends the SetFontSizes CDP command to a browser,
// and returns the browser's response.
func (t *SetFontSizes) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Page.setFontSizes", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetFontSizes CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetFontSizes) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Page.setFontSizes", b)
}

// ParseResponse parses the browser's response
// to the SetFontSizes CDP command.
func (t *SetFontSizes) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetDocumentContent contains the parameters, and acts as
// a Go receiver, for the CDP command `setDocumentContent`.
//
// Sets given markup as the document's HTML.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setDocumentContent
type SetDocumentContent struct {
	// Frame id to set HTML for.
	FrameID string `json:"frameId"`
	// HTML content to set.
	HTML string `json:"html"`
}

// NewSetDocumentContent constructs a new SetDocumentContent struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setDocumentContent
func NewSetDocumentContent(frameID string, html string) *SetDocumentContent {
	return &SetDocumentContent{
		FrameID: frameID,
		HTML:    html,
	}
}

// Do sends the SetDocumentContent CDP command to a browser,
// and returns the browser's response.
func (t *SetDocumentContent) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Page.setDocumentContent", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetDocumentContent CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetDocumentContent) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Page.setDocumentContent", b)
}

// ParseResponse parses the browser's response
// to the SetDocumentContent CDP command.
func (t *SetDocumentContent) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetDownloadBehavior contains the parameters, and acts as
// a Go receiver, for the CDP command `setDownloadBehavior`.
//
// Set the behavior when downloading a file.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setDownloadBehavior
//
// This CDP method is deprecated.
// This CDP method is experimental.
type SetDownloadBehavior struct {
	// Whether to allow all or deny all download requests, or use default Chrome behavior if
	// available (otherwise deny).
	Behavior string `json:"behavior"`
	// The default path to save downloaded files to. This is required if behavior is set to 'allow'
	DownloadPath string `json:"downloadPath,omitempty"`
}

// NewSetDownloadBehavior constructs a new SetDownloadBehavior struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setDownloadBehavior
//
// This CDP method is deprecated.
// This CDP method is experimental.
func NewSetDownloadBehavior(behavior string) *SetDownloadBehavior {
	return &SetDownloadBehavior{
		Behavior: behavior,
	}
}

// SetDownloadPath adds or modifies the value of the optional
// parameter `downloadPath` in the SetDownloadBehavior CDP command.
//
// The default path to save downloaded files to. This is required if behavior is set to 'allow'
func (t *SetDownloadBehavior) SetDownloadPath(v string) *SetDownloadBehavior {
	t.DownloadPath = v
	return t
}

// Do sends the SetDownloadBehavior CDP command to a browser,
// and returns the browser's response.
func (t *SetDownloadBehavior) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Page.setDownloadBehavior", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetDownloadBehavior CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetDownloadBehavior) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Page.setDownloadBehavior", b)
}

// ParseResponse parses the browser's response
// to the SetDownloadBehavior CDP command.
func (t *SetDownloadBehavior) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetLifecycleEventsEnabled contains the parameters, and acts as
// a Go receiver, for the CDP command `setLifecycleEventsEnabled`.
//
// Controls whether page will emit lifecycle events.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setLifecycleEventsEnabled
//
// This CDP method is experimental.
type SetLifecycleEventsEnabled struct {
	// If true, starts emitting lifecycle events.
	Enabled bool `json:"enabled"`
}

// NewSetLifecycleEventsEnabled constructs a new SetLifecycleEventsEnabled struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setLifecycleEventsEnabled
//
// This CDP method is experimental.
func NewSetLifecycleEventsEnabled(enabled bool) *SetLifecycleEventsEnabled {
	return &SetLifecycleEventsEnabled{
		Enabled: enabled,
	}
}

// Do sends the SetLifecycleEventsEnabled CDP command to a browser,
// and returns the browser's response.
func (t *SetLifecycleEventsEnabled) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Page.setLifecycleEventsEnabled", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetLifecycleEventsEnabled CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetLifecycleEventsEnabled) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Page.setLifecycleEventsEnabled", b)
}

// ParseResponse parses the browser's response
// to the SetLifecycleEventsEnabled CDP command.
func (t *SetLifecycleEventsEnabled) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// StartScreencast contains the parameters, and acts as
// a Go receiver, for the CDP command `startScreencast`.
//
// Starts sending each frame using the `screencastFrame` event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-startScreencast
//
// This CDP method is experimental.
type StartScreencast struct {
	// Image compression format.
	Format string `json:"format,omitempty"`
	// Compression quality from range [0..100].
	Quality int64 `json:"quality,omitempty"`
	// Maximum screenshot width.
	MaxWidth int64 `json:"maxWidth,omitempty"`
	// Maximum screenshot height.
	MaxHeight int64 `json:"maxHeight,omitempty"`
	// Send every n-th frame.
	EveryNthFrame int64 `json:"everyNthFrame,omitempty"`
}

// NewStartScreencast constructs a new StartScreencast struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-startScreencast
//
// This CDP method is experimental.
func NewStartScreencast() *StartScreencast {
	return &StartScreencast{}
}

// SetFormat adds or modifies the value of the optional
// parameter `format` in the StartScreencast CDP command.
//
// Image compression format.
func (t *StartScreencast) SetFormat(v string) *StartScreencast {
	t.Format = v
	return t
}

// SetQuality adds or modifies the value of the optional
// parameter `quality` in the StartScreencast CDP command.
//
// Compression quality from range [0..100].
func (t *StartScreencast) SetQuality(v int64) *StartScreencast {
	t.Quality = v
	return t
}

// SetMaxWidth adds or modifies the value of the optional
// parameter `maxWidth` in the StartScreencast CDP command.
//
// Maximum screenshot width.
func (t *StartScreencast) SetMaxWidth(v int64) *StartScreencast {
	t.MaxWidth = v
	return t
}

// SetMaxHeight adds or modifies the value of the optional
// parameter `maxHeight` in the StartScreencast CDP command.
//
// Maximum screenshot height.
func (t *StartScreencast) SetMaxHeight(v int64) *StartScreencast {
	t.MaxHeight = v
	return t
}

// SetEveryNthFrame adds or modifies the value of the optional
// parameter `everyNthFrame` in the StartScreencast CDP command.
//
// Send every n-th frame.
func (t *StartScreencast) SetEveryNthFrame(v int64) *StartScreencast {
	t.EveryNthFrame = v
	return t
}

// Do sends the StartScreencast CDP command to a browser,
// and returns the browser's response.
func (t *StartScreencast) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Page.startScreencast", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the StartScreencast CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *StartScreencast) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Page.startScreencast", b)
}

// ParseResponse parses the browser's response
// to the StartScreencast CDP command.
func (t *StartScreencast) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// StopLoading contains the parameters, and acts as
// a Go receiver, for the CDP command `stopLoading`.
//
// Force the page stop all navigations and pending resource fetches.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-stopLoading
type StopLoading struct{}

// NewStopLoading constructs a new StopLoading struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-stopLoading
func NewStopLoading() *StopLoading {
	return &StopLoading{}
}

// Do sends the StopLoading CDP command to a browser,
// and returns the browser's response.
func (t *StopLoading) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Page.stopLoading", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the StopLoading CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *StopLoading) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Page.stopLoading", nil)
}

// ParseResponse parses the browser's response
// to the StopLoading CDP command.
func (t *StopLoading) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// Crash contains the parameters, and acts as
// a Go receiver, for the CDP command `crash`.
//
// Crashes renderer on the IO thread, generates minidumps.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-crash
//
// This CDP method is experimental.
type Crash struct{}

// NewCrash constructs a new Crash struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-crash
//
// This CDP method is experimental.
func NewCrash() *Crash {
	return &Crash{}
}

// Do sends the Crash CDP command to a browser,
// and returns the browser's response.
func (t *Crash) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Page.crash", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the Crash CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *Crash) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Page.crash", nil)
}

// ParseResponse parses the browser's response
// to the Crash CDP command.
func (t *Crash) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// Close contains the parameters, and acts as
// a Go receiver, for the CDP command `close`.
//
// Tries to close page, running its beforeunload hooks, if any.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-close
//
// This CDP method is experimental.
type Close struct{}

// NewClose constructs a new Close struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-close
//
// This CDP method is experimental.
func NewClose() *Close {
	return &Close{}
}

// Do sends the Close CDP command to a browser,
// and returns the browser's response.
func (t *Close) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Page.close", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the Close CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *Close) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Page.close", nil)
}

// ParseResponse parses the browser's response
// to the Close CDP command.
func (t *Close) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetWebLifecycleState contains the parameters, and acts as
// a Go receiver, for the CDP command `setWebLifecycleState`.
//
// Tries to update the web lifecycle state of the page.
// It will transition the page to the given state according to:
// https://github.com/WICG/web-lifecycle/
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setWebLifecycleState
//
// This CDP method is experimental.
type SetWebLifecycleState struct {
	// Target lifecycle state
	State string `json:"state"`
}

// NewSetWebLifecycleState constructs a new SetWebLifecycleState struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setWebLifecycleState
//
// This CDP method is experimental.
func NewSetWebLifecycleState(state string) *SetWebLifecycleState {
	return &SetWebLifecycleState{
		State: state,
	}
}

// Do sends the SetWebLifecycleState CDP command to a browser,
// and returns the browser's response.
func (t *SetWebLifecycleState) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Page.setWebLifecycleState", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetWebLifecycleState CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetWebLifecycleState) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Page.setWebLifecycleState", b)
}

// ParseResponse parses the browser's response
// to the SetWebLifecycleState CDP command.
func (t *SetWebLifecycleState) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// StopScreencast contains the parameters, and acts as
// a Go receiver, for the CDP command `stopScreencast`.
//
// Stops sending each frame in the `screencastFrame`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-stopScreencast
//
// This CDP method is experimental.
type StopScreencast struct{}

// NewStopScreencast constructs a new StopScreencast struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-stopScreencast
//
// This CDP method is experimental.
func NewStopScreencast() *StopScreencast {
	return &StopScreencast{}
}

// Do sends the StopScreencast CDP command to a browser,
// and returns the browser's response.
func (t *StopScreencast) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Page.stopScreencast", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the StopScreencast CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *StopScreencast) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Page.stopScreencast", nil)
}

// ParseResponse parses the browser's response
// to the StopScreencast CDP command.
func (t *StopScreencast) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetProduceCompilationCache contains the parameters, and acts as
// a Go receiver, for the CDP command `setProduceCompilationCache`.
//
// Forces compilation cache to be generated for every subresource script.
// See also: `Page.produceCompilationCache`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setProduceCompilationCache
//
// This CDP method is experimental.
type SetProduceCompilationCache struct {
	Enabled bool `json:"enabled"`
}

// NewSetProduceCompilationCache constructs a new SetProduceCompilationCache struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setProduceCompilationCache
//
// This CDP method is experimental.
func NewSetProduceCompilationCache(enabled bool) *SetProduceCompilationCache {
	return &SetProduceCompilationCache{
		Enabled: enabled,
	}
}

// Do sends the SetProduceCompilationCache CDP command to a browser,
// and returns the browser's response.
func (t *SetProduceCompilationCache) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Page.setProduceCompilationCache", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetProduceCompilationCache CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetProduceCompilationCache) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Page.setProduceCompilationCache", b)
}

// ParseResponse parses the browser's response
// to the SetProduceCompilationCache CDP command.
func (t *SetProduceCompilationCache) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// ProduceCompilationCache contains the parameters, and acts as
// a Go receiver, for the CDP command `produceCompilationCache`.
//
// Requests backend to produce compilation cache for the specified scripts.
// Unlike setProduceCompilationCache, this allows client to only produce cache
// for specific scripts. `scripts` are appeneded to the list of scripts
// for which the cache for would produced. Disabling compilation cache with
// `setProduceCompilationCache` would reset all pending cache requests.
// The list may also be reset during page navigation.
// When script with a matching URL is encountered, the cache is optionally
// produced upon backend discretion, based on internal heuristics.
// See also: `Page.compilationCacheProduced`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-produceCompilationCache
//
// This CDP method is experimental.
type ProduceCompilationCache struct {
	Scripts []CompilationCacheParams `json:"scripts"`
}

// NewProduceCompilationCache constructs a new ProduceCompilationCache struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-produceCompilationCache
//
// This CDP method is experimental.
func NewProduceCompilationCache(scripts []CompilationCacheParams) *ProduceCompilationCache {
	return &ProduceCompilationCache{
		Scripts: scripts,
	}
}

// Do sends the ProduceCompilationCache CDP command to a browser,
// and returns the browser's response.
func (t *ProduceCompilationCache) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Page.produceCompilationCache", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the ProduceCompilationCache CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *ProduceCompilationCache) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Page.produceCompilationCache", b)
}

// ParseResponse parses the browser's response
// to the ProduceCompilationCache CDP command.
func (t *ProduceCompilationCache) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// AddCompilationCache contains the parameters, and acts as
// a Go receiver, for the CDP command `addCompilationCache`.
//
// Seeds compilation cache for given url. Compilation cache does not survive
// cross-process navigation.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-addCompilationCache
//
// This CDP method is experimental.
type AddCompilationCache struct {
	URL string `json:"url"`
	// Base64-encoded data (Encoded as a base64 string when passed over JSON)
	Data string `json:"data"`
}

// NewAddCompilationCache constructs a new AddCompilationCache struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-addCompilationCache
//
// This CDP method is experimental.
func NewAddCompilationCache(url string, data string) *AddCompilationCache {
	return &AddCompilationCache{
		URL:  url,
		Data: data,
	}
}

// Do sends the AddCompilationCache CDP command to a browser,
// and returns the browser's response.
func (t *AddCompilationCache) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Page.addCompilationCache", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the AddCompilationCache CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *AddCompilationCache) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Page.addCompilationCache", b)
}

// ParseResponse parses the browser's response
// to the AddCompilationCache CDP command.
func (t *AddCompilationCache) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// ClearCompilationCache contains the parameters, and acts as
// a Go receiver, for the CDP command `clearCompilationCache`.
//
// Clears seeded compilation cache.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-clearCompilationCache
//
// This CDP method is experimental.
type ClearCompilationCache struct{}

// NewClearCompilationCache constructs a new ClearCompilationCache struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-clearCompilationCache
//
// This CDP method is experimental.
func NewClearCompilationCache() *ClearCompilationCache {
	return &ClearCompilationCache{}
}

// Do sends the ClearCompilationCache CDP command to a browser,
// and returns the browser's response.
func (t *ClearCompilationCache) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Page.clearCompilationCache", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the ClearCompilationCache CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *ClearCompilationCache) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Page.clearCompilationCache", nil)
}

// ParseResponse parses the browser's response
// to the ClearCompilationCache CDP command.
func (t *ClearCompilationCache) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// GenerateTestReport contains the parameters, and acts as
// a Go receiver, for the CDP command `generateTestReport`.
//
// Generates a report for testing.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-generateTestReport
//
// This CDP method is experimental.
type GenerateTestReport struct {
	// Message to be displayed in the report.
	Message string `json:"message"`
	// Specifies the endpoint group to deliver the report to.
	Group string `json:"group,omitempty"`
}

// NewGenerateTestReport constructs a new GenerateTestReport struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-generateTestReport
//
// This CDP method is experimental.
func NewGenerateTestReport(message string) *GenerateTestReport {
	return &GenerateTestReport{
		Message: message,
	}
}

// SetGroup adds or modifies the value of the optional
// parameter `group` in the GenerateTestReport CDP command.
//
// Specifies the endpoint group to deliver the report to.
func (t *GenerateTestReport) SetGroup(v string) *GenerateTestReport {
	t.Group = v
	return t
}

// Do sends the GenerateTestReport CDP command to a browser,
// and returns the browser's response.
func (t *GenerateTestReport) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Page.generateTestReport", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the GenerateTestReport CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GenerateTestReport) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Page.generateTestReport", b)
}

// ParseResponse parses the browser's response
// to the GenerateTestReport CDP command.
func (t *GenerateTestReport) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// WaitForDebugger contains the parameters, and acts as
// a Go receiver, for the CDP command `waitForDebugger`.
//
// Pauses page execution. Can be resumed using generic Runtime.runIfWaitingForDebugger.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-waitForDebugger
//
// This CDP method is experimental.
type WaitForDebugger struct{}

// NewWaitForDebugger constructs a new WaitForDebugger struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-waitForDebugger
//
// This CDP method is experimental.
func NewWaitForDebugger() *WaitForDebugger {
	return &WaitForDebugger{}
}

// Do sends the WaitForDebugger CDP command to a browser,
// and returns the browser's response.
func (t *WaitForDebugger) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Page.waitForDebugger", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the WaitForDebugger CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *WaitForDebugger) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Page.waitForDebugger", nil)
}

// ParseResponse parses the browser's response
// to the WaitForDebugger CDP command.
func (t *WaitForDebugger) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetInterceptFileChooserDialog contains the parameters, and acts as
// a Go receiver, for the CDP command `setInterceptFileChooserDialog`.
//
// Intercept file chooser requests and transfer control to protocol clients.
// When file chooser interception is enabled, native file chooser dialog is not shown.
// Instead, a protocol event `Page.fileChooserOpened` is emitted.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setInterceptFileChooserDialog
//
// This CDP method is experimental.
type SetInterceptFileChooserDialog struct {
	Enabled bool `json:"enabled"`
}

// NewSetInterceptFileChooserDialog constructs a new SetInterceptFileChooserDialog struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-setInterceptFileChooserDialog
//
// This CDP method is experimental.
func NewSetInterceptFileChooserDialog(enabled bool) *SetInterceptFileChooserDialog {
	return &SetInterceptFileChooserDialog{
		Enabled: enabled,
	}
}

// Do sends the SetInterceptFileChooserDialog CDP command to a browser,
// and returns the browser's response.
func (t *SetInterceptFileChooserDialog) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Page.setInterceptFileChooserDialog", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetInterceptFileChooserDialog CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetInterceptFileChooserDialog) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Page.setInterceptFileChooserDialog", b)
}

// ParseResponse parses the browser's response
// to the SetInterceptFileChooserDialog CDP command.
func (t *SetInterceptFileChooserDialog) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}
