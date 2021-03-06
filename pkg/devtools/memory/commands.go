package memory

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
)

// GetDOMCounters contains the parameters, and acts as
// a Go receiver, for the CDP command `getDOMCounters`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-getDOMCounters
type GetDOMCounters struct{}

// NewGetDOMCounters constructs a new GetDOMCounters struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-getDOMCounters
func NewGetDOMCounters() *GetDOMCounters {
	return &GetDOMCounters{}
}

// GetDOMCountersResult contains the browser's response
// to calling the GetDOMCounters CDP command with Do().
type GetDOMCountersResult struct {
	Documents        int64 `json:"documents"`
	Nodes            int64 `json:"nodes"`
	JsEventListeners int64 `json:"jsEventListeners"`
}

// Do sends the GetDOMCounters CDP command to a browser,
// and returns the browser's response.
func (t *GetDOMCounters) Do(ctx context.Context) (*GetDOMCountersResult, error) {
	m, err := devtools.SendAndWait(ctx, "Memory.getDOMCounters", nil)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetDOMCounters CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetDOMCounters) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Memory.getDOMCounters", nil)
}

// ParseResponse parses the browser's response
// to the GetDOMCounters CDP command.
func (t *GetDOMCounters) ParseResponse(m *devtools.Message) (*GetDOMCountersResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetDOMCountersResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// PrepareForLeakDetection contains the parameters, and acts as
// a Go receiver, for the CDP command `prepareForLeakDetection`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-prepareForLeakDetection
type PrepareForLeakDetection struct{}

// NewPrepareForLeakDetection constructs a new PrepareForLeakDetection struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-prepareForLeakDetection
func NewPrepareForLeakDetection() *PrepareForLeakDetection {
	return &PrepareForLeakDetection{}
}

// Do sends the PrepareForLeakDetection CDP command to a browser,
// and returns the browser's response.
func (t *PrepareForLeakDetection) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Memory.prepareForLeakDetection", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the PrepareForLeakDetection CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *PrepareForLeakDetection) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Memory.prepareForLeakDetection", nil)
}

// ParseResponse parses the browser's response
// to the PrepareForLeakDetection CDP command.
func (t *PrepareForLeakDetection) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// ForciblyPurgeJavaScriptMemory contains the parameters, and acts as
// a Go receiver, for the CDP command `forciblyPurgeJavaScriptMemory`.
//
// Simulate OomIntervention by purging V8 memory.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-forciblyPurgeJavaScriptMemory
type ForciblyPurgeJavaScriptMemory struct{}

// NewForciblyPurgeJavaScriptMemory constructs a new ForciblyPurgeJavaScriptMemory struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-forciblyPurgeJavaScriptMemory
func NewForciblyPurgeJavaScriptMemory() *ForciblyPurgeJavaScriptMemory {
	return &ForciblyPurgeJavaScriptMemory{}
}

// Do sends the ForciblyPurgeJavaScriptMemory CDP command to a browser,
// and returns the browser's response.
func (t *ForciblyPurgeJavaScriptMemory) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Memory.forciblyPurgeJavaScriptMemory", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the ForciblyPurgeJavaScriptMemory CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *ForciblyPurgeJavaScriptMemory) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Memory.forciblyPurgeJavaScriptMemory", nil)
}

// ParseResponse parses the browser's response
// to the ForciblyPurgeJavaScriptMemory CDP command.
func (t *ForciblyPurgeJavaScriptMemory) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetPressureNotificationsSuppressed contains the parameters, and acts as
// a Go receiver, for the CDP command `setPressureNotificationsSuppressed`.
//
// Enable/disable suppressing memory pressure notifications in all processes.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-setPressureNotificationsSuppressed
type SetPressureNotificationsSuppressed struct {
	// If true, memory pressure notifications will be suppressed.
	Suppressed bool `json:"suppressed"`
}

// NewSetPressureNotificationsSuppressed constructs a new SetPressureNotificationsSuppressed struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-setPressureNotificationsSuppressed
func NewSetPressureNotificationsSuppressed(suppressed bool) *SetPressureNotificationsSuppressed {
	return &SetPressureNotificationsSuppressed{
		Suppressed: suppressed,
	}
}

// Do sends the SetPressureNotificationsSuppressed CDP command to a browser,
// and returns the browser's response.
func (t *SetPressureNotificationsSuppressed) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Memory.setPressureNotificationsSuppressed", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetPressureNotificationsSuppressed CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetPressureNotificationsSuppressed) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Memory.setPressureNotificationsSuppressed", b)
}

// ParseResponse parses the browser's response
// to the SetPressureNotificationsSuppressed CDP command.
func (t *SetPressureNotificationsSuppressed) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SimulatePressureNotification contains the parameters, and acts as
// a Go receiver, for the CDP command `simulatePressureNotification`.
//
// Simulate a memory pressure notification in all processes.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-simulatePressureNotification
type SimulatePressureNotification struct {
	// Memory pressure level of the notification.
	Level PressureLevel `json:"level"`
}

// NewSimulatePressureNotification constructs a new SimulatePressureNotification struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-simulatePressureNotification
func NewSimulatePressureNotification(level PressureLevel) *SimulatePressureNotification {
	return &SimulatePressureNotification{
		Level: level,
	}
}

// Do sends the SimulatePressureNotification CDP command to a browser,
// and returns the browser's response.
func (t *SimulatePressureNotification) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Memory.simulatePressureNotification", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SimulatePressureNotification CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SimulatePressureNotification) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Memory.simulatePressureNotification", b)
}

// ParseResponse parses the browser's response
// to the SimulatePressureNotification CDP command.
func (t *SimulatePressureNotification) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// StartSampling contains the parameters, and acts as
// a Go receiver, for the CDP command `startSampling`.
//
// Start collecting native memory profile.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-startSampling
type StartSampling struct {
	// Average number of bytes between samples.
	SamplingInterval int64 `json:"samplingInterval,omitempty"`
	// Do not randomize intervals between samples.
	SuppressRandomness bool `json:"suppressRandomness,omitempty"`
}

// NewStartSampling constructs a new StartSampling struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-startSampling
func NewStartSampling() *StartSampling {
	return &StartSampling{}
}

// SetSamplingInterval adds or modifies the value of the optional
// parameter `samplingInterval` in the StartSampling CDP command.
//
// Average number of bytes between samples.
func (t *StartSampling) SetSamplingInterval(v int64) *StartSampling {
	t.SamplingInterval = v
	return t
}

// SetSuppressRandomness adds or modifies the value of the optional
// parameter `suppressRandomness` in the StartSampling CDP command.
//
// Do not randomize intervals between samples.
func (t *StartSampling) SetSuppressRandomness(v bool) *StartSampling {
	t.SuppressRandomness = v
	return t
}

// Do sends the StartSampling CDP command to a browser,
// and returns the browser's response.
func (t *StartSampling) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Memory.startSampling", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the StartSampling CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *StartSampling) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Memory.startSampling", b)
}

// ParseResponse parses the browser's response
// to the StartSampling CDP command.
func (t *StartSampling) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// StopSampling contains the parameters, and acts as
// a Go receiver, for the CDP command `stopSampling`.
//
// Stop collecting native memory profile.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-stopSampling
type StopSampling struct{}

// NewStopSampling constructs a new StopSampling struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-stopSampling
func NewStopSampling() *StopSampling {
	return &StopSampling{}
}

// Do sends the StopSampling CDP command to a browser,
// and returns the browser's response.
func (t *StopSampling) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Memory.stopSampling", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the StopSampling CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *StopSampling) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Memory.stopSampling", nil)
}

// ParseResponse parses the browser's response
// to the StopSampling CDP command.
func (t *StopSampling) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// GetAllTimeSamplingProfile contains the parameters, and acts as
// a Go receiver, for the CDP command `getAllTimeSamplingProfile`.
//
// Retrieve native memory allocations profile
// collected since renderer process startup.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-getAllTimeSamplingProfile
type GetAllTimeSamplingProfile struct{}

// NewGetAllTimeSamplingProfile constructs a new GetAllTimeSamplingProfile struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-getAllTimeSamplingProfile
func NewGetAllTimeSamplingProfile() *GetAllTimeSamplingProfile {
	return &GetAllTimeSamplingProfile{}
}

// GetAllTimeSamplingProfileResult contains the browser's response
// to calling the GetAllTimeSamplingProfile CDP command with Do().
type GetAllTimeSamplingProfileResult struct {
	Profile SamplingProfile `json:"profile"`
}

// Do sends the GetAllTimeSamplingProfile CDP command to a browser,
// and returns the browser's response.
func (t *GetAllTimeSamplingProfile) Do(ctx context.Context) (*GetAllTimeSamplingProfileResult, error) {
	m, err := devtools.SendAndWait(ctx, "Memory.getAllTimeSamplingProfile", nil)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetAllTimeSamplingProfile CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetAllTimeSamplingProfile) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Memory.getAllTimeSamplingProfile", nil)
}

// ParseResponse parses the browser's response
// to the GetAllTimeSamplingProfile CDP command.
func (t *GetAllTimeSamplingProfile) ParseResponse(m *devtools.Message) (*GetAllTimeSamplingProfileResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetAllTimeSamplingProfileResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetBrowserSamplingProfile contains the parameters, and acts as
// a Go receiver, for the CDP command `getBrowserSamplingProfile`.
//
// Retrieve native memory allocations profile
// collected since browser process startup.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-getBrowserSamplingProfile
type GetBrowserSamplingProfile struct{}

// NewGetBrowserSamplingProfile constructs a new GetBrowserSamplingProfile struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-getBrowserSamplingProfile
func NewGetBrowserSamplingProfile() *GetBrowserSamplingProfile {
	return &GetBrowserSamplingProfile{}
}

// GetBrowserSamplingProfileResult contains the browser's response
// to calling the GetBrowserSamplingProfile CDP command with Do().
type GetBrowserSamplingProfileResult struct {
	Profile SamplingProfile `json:"profile"`
}

// Do sends the GetBrowserSamplingProfile CDP command to a browser,
// and returns the browser's response.
func (t *GetBrowserSamplingProfile) Do(ctx context.Context) (*GetBrowserSamplingProfileResult, error) {
	m, err := devtools.SendAndWait(ctx, "Memory.getBrowserSamplingProfile", nil)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetBrowserSamplingProfile CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetBrowserSamplingProfile) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Memory.getBrowserSamplingProfile", nil)
}

// ParseResponse parses the browser's response
// to the GetBrowserSamplingProfile CDP command.
func (t *GetBrowserSamplingProfile) ParseResponse(m *devtools.Message) (*GetBrowserSamplingProfileResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetBrowserSamplingProfileResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetSamplingProfile contains the parameters, and acts as
// a Go receiver, for the CDP command `getSamplingProfile`.
//
// Retrieve native memory allocations profile collected since last
// `startSampling` call.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-getSamplingProfile
type GetSamplingProfile struct{}

// NewGetSamplingProfile constructs a new GetSamplingProfile struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Memory/#method-getSamplingProfile
func NewGetSamplingProfile() *GetSamplingProfile {
	return &GetSamplingProfile{}
}

// GetSamplingProfileResult contains the browser's response
// to calling the GetSamplingProfile CDP command with Do().
type GetSamplingProfileResult struct {
	Profile SamplingProfile `json:"profile"`
}

// Do sends the GetSamplingProfile CDP command to a browser,
// and returns the browser's response.
func (t *GetSamplingProfile) Do(ctx context.Context) (*GetSamplingProfileResult, error) {
	m, err := devtools.SendAndWait(ctx, "Memory.getSamplingProfile", nil)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetSamplingProfile CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetSamplingProfile) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Memory.getSamplingProfile", nil)
}

// ParseResponse parses the browser's response
// to the GetSamplingProfile CDP command.
func (t *GetSamplingProfile) ParseResponse(m *devtools.Message) (*GetSamplingProfileResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetSamplingProfileResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}
