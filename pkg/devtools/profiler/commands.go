package profiler

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
)

// Disable contains the parameters, and acts as
// a Go receiver, for the CDP command `disable`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-disable
type Disable struct{}

// NewDisable constructs a new Disable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-disable
func NewDisable() *Disable {
	return &Disable{}
}

// Do sends the Disable CDP command to a browser,
// and returns the browser's response.
func (t *Disable) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Profiler.disable", nil)
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
	return devtools.Send(ctx, "Profiler.disable", nil)
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
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-enable
type Enable struct{}

// NewEnable constructs a new Enable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-enable
func NewEnable() *Enable {
	return &Enable{}
}

// Do sends the Enable CDP command to a browser,
// and returns the browser's response.
func (t *Enable) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Profiler.enable", nil)
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
	return devtools.Send(ctx, "Profiler.enable", nil)
}

// ParseResponse parses the browser's response
// to the Enable CDP command.
func (t *Enable) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// GetBestEffortCoverage contains the parameters, and acts as
// a Go receiver, for the CDP command `getBestEffortCoverage`.
//
// Collect coverage data for the current isolate. The coverage data may be incomplete due to
// garbage collection.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-getBestEffortCoverage
type GetBestEffortCoverage struct{}

// NewGetBestEffortCoverage constructs a new GetBestEffortCoverage struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-getBestEffortCoverage
func NewGetBestEffortCoverage() *GetBestEffortCoverage {
	return &GetBestEffortCoverage{}
}

// GetBestEffortCoverageResult contains the browser's response
// to calling the GetBestEffortCoverage CDP command with Do().
type GetBestEffortCoverageResult struct {
	// Coverage data for the current isolate.
	Result []ScriptCoverage `json:"result"`
}

// Do sends the GetBestEffortCoverage CDP command to a browser,
// and returns the browser's response.
func (t *GetBestEffortCoverage) Do(ctx context.Context) (*GetBestEffortCoverageResult, error) {
	m, err := devtools.SendAndWait(ctx, "Profiler.getBestEffortCoverage", nil)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetBestEffortCoverage CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetBestEffortCoverage) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Profiler.getBestEffortCoverage", nil)
}

// ParseResponse parses the browser's response
// to the GetBestEffortCoverage CDP command.
func (t *GetBestEffortCoverage) ParseResponse(m *devtools.Message) (*GetBestEffortCoverageResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetBestEffortCoverageResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetSamplingInterval contains the parameters, and acts as
// a Go receiver, for the CDP command `setSamplingInterval`.
//
// Changes CPU profiler sampling interval. Must be called before CPU profiles recording started.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-setSamplingInterval
type SetSamplingInterval struct {
	// New sampling interval in microseconds.
	Interval int64 `json:"interval"`
}

// NewSetSamplingInterval constructs a new SetSamplingInterval struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-setSamplingInterval
func NewSetSamplingInterval(interval int64) *SetSamplingInterval {
	return &SetSamplingInterval{
		Interval: interval,
	}
}

// Do sends the SetSamplingInterval CDP command to a browser,
// and returns the browser's response.
func (t *SetSamplingInterval) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Profiler.setSamplingInterval", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetSamplingInterval CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetSamplingInterval) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Profiler.setSamplingInterval", b)
}

// ParseResponse parses the browser's response
// to the SetSamplingInterval CDP command.
func (t *SetSamplingInterval) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// Start contains the parameters, and acts as
// a Go receiver, for the CDP command `start`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-start
type Start struct{}

// NewStart constructs a new Start struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-start
func NewStart() *Start {
	return &Start{}
}

// Do sends the Start CDP command to a browser,
// and returns the browser's response.
func (t *Start) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Profiler.start", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the Start CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *Start) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Profiler.start", nil)
}

// ParseResponse parses the browser's response
// to the Start CDP command.
func (t *Start) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// StartPreciseCoverage contains the parameters, and acts as
// a Go receiver, for the CDP command `startPreciseCoverage`.
//
// Enable precise code coverage. Coverage data for JavaScript executed before enabling precise code
// coverage may be incomplete. Enabling prevents running optimized code and resets execution
// counters.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-startPreciseCoverage
type StartPreciseCoverage struct {
	// Collect accurate call counts beyond simple 'covered' or 'not covered'.
	CallCount bool `json:"callCount,omitempty"`
	// Collect block-based coverage.
	Detailed bool `json:"detailed,omitempty"`
	// Allow the backend to send updates on its own initiative
	AllowTriggeredUpdates bool `json:"allowTriggeredUpdates,omitempty"`
}

// NewStartPreciseCoverage constructs a new StartPreciseCoverage struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-startPreciseCoverage
func NewStartPreciseCoverage() *StartPreciseCoverage {
	return &StartPreciseCoverage{}
}

// SetCallCount adds or modifies the value of the optional
// parameter `callCount` in the StartPreciseCoverage CDP command.
//
// Collect accurate call counts beyond simple 'covered' or 'not covered'.
func (t *StartPreciseCoverage) SetCallCount(v bool) *StartPreciseCoverage {
	t.CallCount = v
	return t
}

// SetDetailed adds or modifies the value of the optional
// parameter `detailed` in the StartPreciseCoverage CDP command.
//
// Collect block-based coverage.
func (t *StartPreciseCoverage) SetDetailed(v bool) *StartPreciseCoverage {
	t.Detailed = v
	return t
}

// SetAllowTriggeredUpdates adds or modifies the value of the optional
// parameter `allowTriggeredUpdates` in the StartPreciseCoverage CDP command.
//
// Allow the backend to send updates on its own initiative
func (t *StartPreciseCoverage) SetAllowTriggeredUpdates(v bool) *StartPreciseCoverage {
	t.AllowTriggeredUpdates = v
	return t
}

// StartPreciseCoverageResult contains the browser's response
// to calling the StartPreciseCoverage CDP command with Do().
type StartPreciseCoverageResult struct {
	// Monotonically increasing time (in seconds) when the coverage update was taken in the backend.
	Timestamp float64 `json:"timestamp"`
}

// Do sends the StartPreciseCoverage CDP command to a browser,
// and returns the browser's response.
func (t *StartPreciseCoverage) Do(ctx context.Context) (*StartPreciseCoverageResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Profiler.startPreciseCoverage", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the StartPreciseCoverage CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *StartPreciseCoverage) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Profiler.startPreciseCoverage", b)
}

// ParseResponse parses the browser's response
// to the StartPreciseCoverage CDP command.
func (t *StartPreciseCoverage) ParseResponse(m *devtools.Message) (*StartPreciseCoverageResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &StartPreciseCoverageResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// StartTypeProfile contains the parameters, and acts as
// a Go receiver, for the CDP command `startTypeProfile`.
//
// Enable type profile.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-startTypeProfile
//
// This CDP method is experimental.
type StartTypeProfile struct{}

// NewStartTypeProfile constructs a new StartTypeProfile struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-startTypeProfile
//
// This CDP method is experimental.
func NewStartTypeProfile() *StartTypeProfile {
	return &StartTypeProfile{}
}

// Do sends the StartTypeProfile CDP command to a browser,
// and returns the browser's response.
func (t *StartTypeProfile) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Profiler.startTypeProfile", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the StartTypeProfile CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *StartTypeProfile) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Profiler.startTypeProfile", nil)
}

// ParseResponse parses the browser's response
// to the StartTypeProfile CDP command.
func (t *StartTypeProfile) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// Stop contains the parameters, and acts as
// a Go receiver, for the CDP command `stop`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-stop
type Stop struct{}

// NewStop constructs a new Stop struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-stop
func NewStop() *Stop {
	return &Stop{}
}

// StopResult contains the browser's response
// to calling the Stop CDP command with Do().
type StopResult struct {
	// Recorded profile.
	Profile Profile `json:"profile"`
}

// Do sends the Stop CDP command to a browser,
// and returns the browser's response.
func (t *Stop) Do(ctx context.Context) (*StopResult, error) {
	m, err := devtools.SendAndWait(ctx, "Profiler.stop", nil)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the Stop CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *Stop) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Profiler.stop", nil)
}

// ParseResponse parses the browser's response
// to the Stop CDP command.
func (t *Stop) ParseResponse(m *devtools.Message) (*StopResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &StopResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// StopPreciseCoverage contains the parameters, and acts as
// a Go receiver, for the CDP command `stopPreciseCoverage`.
//
// Disable precise code coverage. Disabling releases unnecessary execution count records and allows
// executing optimized code.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-stopPreciseCoverage
type StopPreciseCoverage struct{}

// NewStopPreciseCoverage constructs a new StopPreciseCoverage struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-stopPreciseCoverage
func NewStopPreciseCoverage() *StopPreciseCoverage {
	return &StopPreciseCoverage{}
}

// Do sends the StopPreciseCoverage CDP command to a browser,
// and returns the browser's response.
func (t *StopPreciseCoverage) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Profiler.stopPreciseCoverage", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the StopPreciseCoverage CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *StopPreciseCoverage) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Profiler.stopPreciseCoverage", nil)
}

// ParseResponse parses the browser's response
// to the StopPreciseCoverage CDP command.
func (t *StopPreciseCoverage) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// StopTypeProfile contains the parameters, and acts as
// a Go receiver, for the CDP command `stopTypeProfile`.
//
// Disable type profile. Disabling releases type profile data collected so far.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-stopTypeProfile
//
// This CDP method is experimental.
type StopTypeProfile struct{}

// NewStopTypeProfile constructs a new StopTypeProfile struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-stopTypeProfile
//
// This CDP method is experimental.
func NewStopTypeProfile() *StopTypeProfile {
	return &StopTypeProfile{}
}

// Do sends the StopTypeProfile CDP command to a browser,
// and returns the browser's response.
func (t *StopTypeProfile) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Profiler.stopTypeProfile", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the StopTypeProfile CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *StopTypeProfile) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Profiler.stopTypeProfile", nil)
}

// ParseResponse parses the browser's response
// to the StopTypeProfile CDP command.
func (t *StopTypeProfile) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// TakePreciseCoverage contains the parameters, and acts as
// a Go receiver, for the CDP command `takePreciseCoverage`.
//
// Collect coverage data for the current isolate, and resets execution counters. Precise code
// coverage needs to have started.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-takePreciseCoverage
type TakePreciseCoverage struct{}

// NewTakePreciseCoverage constructs a new TakePreciseCoverage struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-takePreciseCoverage
func NewTakePreciseCoverage() *TakePreciseCoverage {
	return &TakePreciseCoverage{}
}

// TakePreciseCoverageResult contains the browser's response
// to calling the TakePreciseCoverage CDP command with Do().
type TakePreciseCoverageResult struct {
	// Coverage data for the current isolate.
	Result []ScriptCoverage `json:"result"`
	// Monotonically increasing time (in seconds) when the coverage update was taken in the backend.
	Timestamp float64 `json:"timestamp"`
}

// Do sends the TakePreciseCoverage CDP command to a browser,
// and returns the browser's response.
func (t *TakePreciseCoverage) Do(ctx context.Context) (*TakePreciseCoverageResult, error) {
	m, err := devtools.SendAndWait(ctx, "Profiler.takePreciseCoverage", nil)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the TakePreciseCoverage CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *TakePreciseCoverage) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Profiler.takePreciseCoverage", nil)
}

// ParseResponse parses the browser's response
// to the TakePreciseCoverage CDP command.
func (t *TakePreciseCoverage) ParseResponse(m *devtools.Message) (*TakePreciseCoverageResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &TakePreciseCoverageResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// TakeTypeProfile contains the parameters, and acts as
// a Go receiver, for the CDP command `takeTypeProfile`.
//
// Collect type profile.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-takeTypeProfile
//
// This CDP method is experimental.
type TakeTypeProfile struct{}

// NewTakeTypeProfile constructs a new TakeTypeProfile struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-takeTypeProfile
//
// This CDP method is experimental.
func NewTakeTypeProfile() *TakeTypeProfile {
	return &TakeTypeProfile{}
}

// TakeTypeProfileResult contains the browser's response
// to calling the TakeTypeProfile CDP command with Do().
type TakeTypeProfileResult struct {
	// Type profile for all scripts since startTypeProfile() was turned on.
	Result []ScriptTypeProfile `json:"result"`
}

// Do sends the TakeTypeProfile CDP command to a browser,
// and returns the browser's response.
func (t *TakeTypeProfile) Do(ctx context.Context) (*TakeTypeProfileResult, error) {
	m, err := devtools.SendAndWait(ctx, "Profiler.takeTypeProfile", nil)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the TakeTypeProfile CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *TakeTypeProfile) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Profiler.takeTypeProfile", nil)
}

// ParseResponse parses the browser's response
// to the TakeTypeProfile CDP command.
func (t *TakeTypeProfile) ParseResponse(m *devtools.Message) (*TakeTypeProfileResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &TakeTypeProfileResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}
