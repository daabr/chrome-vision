package headlessexperimental

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
)

// BeginFrame contains the parameters, and acts as
// a Go receiver, for the CDP command `beginFrame`.
//
// Sends a BeginFrame to the target and returns when the frame was completed. Optionally captures a
// screenshot from the resulting frame. Requires that the target was created with enabled
// BeginFrameControl. Designed for use with --run-all-compositor-stages-before-draw, see also
// https://goo.gl/3zHXhB for more background.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#method-beginFrame
type BeginFrame struct {
	// Timestamp of this BeginFrame in Renderer TimeTicks (milliseconds of uptime). If not set,
	// the current time will be used.
	FrameTimeTicks float64 `json:"frameTimeTicks,omitempty"`
	// The interval between BeginFrames that is reported to the compositor, in milliseconds.
	// Defaults to a 60 frames/second interval, i.e. about 16.666 milliseconds.
	Interval float64 `json:"interval,omitempty"`
	// Whether updates should not be committed and drawn onto the display. False by default. If
	// true, only side effects of the BeginFrame will be run, such as layout and animations, but
	// any visual updates may not be visible on the display or in screenshots.
	NoDisplayUpdates bool `json:"noDisplayUpdates,omitempty"`
	// If set, a screenshot of the frame will be captured and returned in the response. Otherwise,
	// no screenshot will be captured. Note that capturing a screenshot can fail, for example,
	// during renderer initialization. In such a case, no screenshot data will be returned.
	Screenshot *ScreenshotParams `json:"screenshot,omitempty"`
}

// NewBeginFrame constructs a new BeginFrame struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#method-beginFrame
func NewBeginFrame() *BeginFrame {
	return &BeginFrame{}
}

// SetFrameTimeTicks adds or modifies the value of the optional
// parameter `frameTimeTicks` in the BeginFrame CDP command.
//
// Timestamp of this BeginFrame in Renderer TimeTicks (milliseconds of uptime). If not set,
// the current time will be used.
func (t *BeginFrame) SetFrameTimeTicks(v float64) *BeginFrame {
	t.FrameTimeTicks = v
	return t
}

// SetInterval adds or modifies the value of the optional
// parameter `interval` in the BeginFrame CDP command.
//
// The interval between BeginFrames that is reported to the compositor, in milliseconds.
// Defaults to a 60 frames/second interval, i.e. about 16.666 milliseconds.
func (t *BeginFrame) SetInterval(v float64) *BeginFrame {
	t.Interval = v
	return t
}

// SetNoDisplayUpdates adds or modifies the value of the optional
// parameter `noDisplayUpdates` in the BeginFrame CDP command.
//
// Whether updates should not be committed and drawn onto the display. False by default. If
// true, only side effects of the BeginFrame will be run, such as layout and animations, but
// any visual updates may not be visible on the display or in screenshots.
func (t *BeginFrame) SetNoDisplayUpdates(v bool) *BeginFrame {
	t.NoDisplayUpdates = v
	return t
}

// SetScreenshot adds or modifies the value of the optional
// parameter `screenshot` in the BeginFrame CDP command.
//
// If set, a screenshot of the frame will be captured and returned in the response. Otherwise,
// no screenshot will be captured. Note that capturing a screenshot can fail, for example,
// during renderer initialization. In such a case, no screenshot data will be returned.
func (t *BeginFrame) SetScreenshot(v ScreenshotParams) *BeginFrame {
	t.Screenshot = &v
	return t
}

// BeginFrameResult contains the browser's response
// to calling the BeginFrame CDP command with Do().
type BeginFrameResult struct {
	// Whether the BeginFrame resulted in damage and, thus, a new frame was committed to the
	// display. Reported for diagnostic uses, may be removed in the future.
	HasDamage bool `json:"hasDamage"`
	// Base64-encoded image data of the screenshot, if one was requested and successfully taken. (Encoded as a base64 string when passed over JSON)
	ScreenshotData string `json:"screenshotData,omitempty"`
}

// Do sends the BeginFrame CDP command to a browser,
// and returns the browser's response.
func (t *BeginFrame) Do(ctx context.Context) (*BeginFrameResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "HeadlessExperimental.beginFrame", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the BeginFrame CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *BeginFrame) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "HeadlessExperimental.beginFrame", b)
}

// ParseResponse parses the browser's response
// to the BeginFrame CDP command.
func (t *BeginFrame) ParseResponse(m *devtools.Message) (*BeginFrameResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &BeginFrameResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Disable contains the parameters, and acts as
// a Go receiver, for the CDP command `disable`.
//
// Disables headless events for the target.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#method-disable
type Disable struct{}

// NewDisable constructs a new Disable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#method-disable
func NewDisable() *Disable {
	return &Disable{}
}

// Do sends the Disable CDP command to a browser,
// and returns the browser's response.
func (t *Disable) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "HeadlessExperimental.disable", nil)
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
	return devtools.Send(ctx, "HeadlessExperimental.disable", nil)
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
// Enables headless events for the target.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#method-enable
type Enable struct{}

// NewEnable constructs a new Enable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#method-enable
func NewEnable() *Enable {
	return &Enable{}
}

// Do sends the Enable CDP command to a browser,
// and returns the browser's response.
func (t *Enable) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "HeadlessExperimental.enable", nil)
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
	return devtools.Send(ctx, "HeadlessExperimental.enable", nil)
}

// ParseResponse parses the browser's response
// to the Enable CDP command.
func (t *Enable) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}
