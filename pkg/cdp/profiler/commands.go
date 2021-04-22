package profiler

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/cdp"
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
	response, err := cdp.Send(ctx, "Profiler.disable", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
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
	response, err := cdp.Send(ctx, "Profiler.enable", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
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

// GetBestEffortCoverageResponse contains the browser's response
// to calling the GetBestEffortCoverage CDP command with Do().
type GetBestEffortCoverageResponse struct {
	// Coverage data for the current isolate.
	Result []ScriptCoverage `json:"result"`
}

// Do sends the GetBestEffortCoverage CDP command to a browser,
// and returns the browser's response.
func (t *GetBestEffortCoverage) Do(ctx context.Context) (*GetBestEffortCoverageResponse, error) {
	response, err := cdp.Send(ctx, "Profiler.getBestEffortCoverage", nil)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetBestEffortCoverageResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
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
	response, err := cdp.Send(ctx, "Profiler.setSamplingInterval", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
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
	response, err := cdp.Send(ctx, "Profiler.start", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
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

// StartPreciseCoverageResponse contains the browser's response
// to calling the StartPreciseCoverage CDP command with Do().
type StartPreciseCoverageResponse struct {
	// Monotonically increasing time (in seconds) when the coverage update was taken in the backend.
	Timestamp float64 `json:"timestamp"`
}

// Do sends the StartPreciseCoverage CDP command to a browser,
// and returns the browser's response.
func (t *StartPreciseCoverage) Do(ctx context.Context) (*StartPreciseCoverageResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "Profiler.startPreciseCoverage", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &StartPreciseCoverageResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
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
	response, err := cdp.Send(ctx, "Profiler.startTypeProfile", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
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

// StopResponse contains the browser's response
// to calling the Stop CDP command with Do().
type StopResponse struct {
	// Recorded profile.
	Profile Profile `json:"profile"`
}

// Do sends the Stop CDP command to a browser,
// and returns the browser's response.
func (t *Stop) Do(ctx context.Context) (*StopResponse, error) {
	response, err := cdp.Send(ctx, "Profiler.stop", nil)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &StopResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
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
	response, err := cdp.Send(ctx, "Profiler.stopPreciseCoverage", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
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
	response, err := cdp.Send(ctx, "Profiler.stopTypeProfile", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
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

// TakePreciseCoverageResponse contains the browser's response
// to calling the TakePreciseCoverage CDP command with Do().
type TakePreciseCoverageResponse struct {
	// Coverage data for the current isolate.
	Result []ScriptCoverage `json:"result"`
	// Monotonically increasing time (in seconds) when the coverage update was taken in the backend.
	Timestamp float64 `json:"timestamp"`
}

// Do sends the TakePreciseCoverage CDP command to a browser,
// and returns the browser's response.
func (t *TakePreciseCoverage) Do(ctx context.Context) (*TakePreciseCoverageResponse, error) {
	response, err := cdp.Send(ctx, "Profiler.takePreciseCoverage", nil)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &TakePreciseCoverageResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
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

// TakeTypeProfileResponse contains the browser's response
// to calling the TakeTypeProfile CDP command with Do().
type TakeTypeProfileResponse struct {
	// Type profile for all scripts since startTypeProfile() was turned on.
	Result []ScriptTypeProfile `json:"result"`
}

// Do sends the TakeTypeProfile CDP command to a browser,
// and returns the browser's response.
func (t *TakeTypeProfile) Do(ctx context.Context) (*TakeTypeProfileResponse, error) {
	response, err := cdp.Send(ctx, "Profiler.takeTypeProfile", nil)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &TakeTypeProfileResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// EnableCounters contains the parameters, and acts as
// a Go receiver, for the CDP command `enableCounters`.
//
// Enable counters collection.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-enableCounters
//
// This CDP method is experimental.
type EnableCounters struct{}

// NewEnableCounters constructs a new EnableCounters struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-enableCounters
//
// This CDP method is experimental.
func NewEnableCounters() *EnableCounters {
	return &EnableCounters{}
}

// Do sends the EnableCounters CDP command to a browser,
// and returns the browser's response.
func (t *EnableCounters) Do(ctx context.Context) error {
	response, err := cdp.Send(ctx, "Profiler.enableCounters", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// DisableCounters contains the parameters, and acts as
// a Go receiver, for the CDP command `disableCounters`.
//
// Disable counters collection.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-disableCounters
//
// This CDP method is experimental.
type DisableCounters struct{}

// NewDisableCounters constructs a new DisableCounters struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-disableCounters
//
// This CDP method is experimental.
func NewDisableCounters() *DisableCounters {
	return &DisableCounters{}
}

// Do sends the DisableCounters CDP command to a browser,
// and returns the browser's response.
func (t *DisableCounters) Do(ctx context.Context) error {
	response, err := cdp.Send(ctx, "Profiler.disableCounters", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// GetCounters contains the parameters, and acts as
// a Go receiver, for the CDP command `getCounters`.
//
// Retrieve counters.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-getCounters
//
// This CDP method is experimental.
type GetCounters struct{}

// NewGetCounters constructs a new GetCounters struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-getCounters
//
// This CDP method is experimental.
func NewGetCounters() *GetCounters {
	return &GetCounters{}
}

// GetCountersResponse contains the browser's response
// to calling the GetCounters CDP command with Do().
type GetCountersResponse struct {
	// Collected counters information.
	Result []CounterInfo `json:"result"`
}

// Do sends the GetCounters CDP command to a browser,
// and returns the browser's response.
func (t *GetCounters) Do(ctx context.Context) (*GetCountersResponse, error) {
	response, err := cdp.Send(ctx, "Profiler.getCounters", nil)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetCountersResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// EnableRuntimeCallStats contains the parameters, and acts as
// a Go receiver, for the CDP command `enableRuntimeCallStats`.
//
// Enable run time call stats collection.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-enableRuntimeCallStats
//
// This CDP method is experimental.
type EnableRuntimeCallStats struct{}

// NewEnableRuntimeCallStats constructs a new EnableRuntimeCallStats struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-enableRuntimeCallStats
//
// This CDP method is experimental.
func NewEnableRuntimeCallStats() *EnableRuntimeCallStats {
	return &EnableRuntimeCallStats{}
}

// Do sends the EnableRuntimeCallStats CDP command to a browser,
// and returns the browser's response.
func (t *EnableRuntimeCallStats) Do(ctx context.Context) error {
	response, err := cdp.Send(ctx, "Profiler.enableRuntimeCallStats", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// DisableRuntimeCallStats contains the parameters, and acts as
// a Go receiver, for the CDP command `disableRuntimeCallStats`.
//
// Disable run time call stats collection.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-disableRuntimeCallStats
//
// This CDP method is experimental.
type DisableRuntimeCallStats struct{}

// NewDisableRuntimeCallStats constructs a new DisableRuntimeCallStats struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-disableRuntimeCallStats
//
// This CDP method is experimental.
func NewDisableRuntimeCallStats() *DisableRuntimeCallStats {
	return &DisableRuntimeCallStats{}
}

// Do sends the DisableRuntimeCallStats CDP command to a browser,
// and returns the browser's response.
func (t *DisableRuntimeCallStats) Do(ctx context.Context) error {
	response, err := cdp.Send(ctx, "Profiler.disableRuntimeCallStats", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// GetRuntimeCallStats contains the parameters, and acts as
// a Go receiver, for the CDP command `getRuntimeCallStats`.
//
// Retrieve run time call stats.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-getRuntimeCallStats
//
// This CDP method is experimental.
type GetRuntimeCallStats struct{}

// NewGetRuntimeCallStats constructs a new GetRuntimeCallStats struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-getRuntimeCallStats
//
// This CDP method is experimental.
func NewGetRuntimeCallStats() *GetRuntimeCallStats {
	return &GetRuntimeCallStats{}
}

// GetRuntimeCallStatsResponse contains the browser's response
// to calling the GetRuntimeCallStats CDP command with Do().
type GetRuntimeCallStatsResponse struct {
	// Collected runtime call counter information.
	Result []RuntimeCallCounterInfo `json:"result"`
}

// Do sends the GetRuntimeCallStats CDP command to a browser,
// and returns the browser's response.
func (t *GetRuntimeCallStats) Do(ctx context.Context) (*GetRuntimeCallStatsResponse, error) {
	response, err := cdp.Send(ctx, "Profiler.getRuntimeCallStats", nil)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetRuntimeCallStatsResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}
