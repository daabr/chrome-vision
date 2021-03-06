package heapprofiler

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
	"github.com/daabr/chrome-vision/pkg/devtools/runtime"
)

// AddInspectedHeapObject contains the parameters, and acts as
// a Go receiver, for the CDP command `addInspectedHeapObject`.
//
// Enables console to refer to the node with given id via $x (see Command Line API for more details
// $x functions).
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-addInspectedHeapObject
type AddInspectedHeapObject struct {
	// Heap snapshot object id to be accessible by means of $x command line API.
	HeapObjectID string `json:"heapObjectId"`
}

// NewAddInspectedHeapObject constructs a new AddInspectedHeapObject struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-addInspectedHeapObject
func NewAddInspectedHeapObject(heapObjectID string) *AddInspectedHeapObject {
	return &AddInspectedHeapObject{
		HeapObjectID: heapObjectID,
	}
}

// Do sends the AddInspectedHeapObject CDP command to a browser,
// and returns the browser's response.
func (t *AddInspectedHeapObject) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "HeapProfiler.addInspectedHeapObject", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the AddInspectedHeapObject CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *AddInspectedHeapObject) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "HeapProfiler.addInspectedHeapObject", b)
}

// ParseResponse parses the browser's response
// to the AddInspectedHeapObject CDP command.
func (t *AddInspectedHeapObject) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// CollectGarbage contains the parameters, and acts as
// a Go receiver, for the CDP command `collectGarbage`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-collectGarbage
type CollectGarbage struct{}

// NewCollectGarbage constructs a new CollectGarbage struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-collectGarbage
func NewCollectGarbage() *CollectGarbage {
	return &CollectGarbage{}
}

// Do sends the CollectGarbage CDP command to a browser,
// and returns the browser's response.
func (t *CollectGarbage) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "HeapProfiler.collectGarbage", nil)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the CollectGarbage CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *CollectGarbage) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "HeapProfiler.collectGarbage", nil)
}

// ParseResponse parses the browser's response
// to the CollectGarbage CDP command.
func (t *CollectGarbage) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// Disable contains the parameters, and acts as
// a Go receiver, for the CDP command `disable`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-disable
type Disable struct{}

// NewDisable constructs a new Disable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-disable
func NewDisable() *Disable {
	return &Disable{}
}

// Do sends the Disable CDP command to a browser,
// and returns the browser's response.
func (t *Disable) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "HeapProfiler.disable", nil)
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
	return devtools.Send(ctx, "HeapProfiler.disable", nil)
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
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-enable
type Enable struct{}

// NewEnable constructs a new Enable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-enable
func NewEnable() *Enable {
	return &Enable{}
}

// Do sends the Enable CDP command to a browser,
// and returns the browser's response.
func (t *Enable) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "HeapProfiler.enable", nil)
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
	return devtools.Send(ctx, "HeapProfiler.enable", nil)
}

// ParseResponse parses the browser's response
// to the Enable CDP command.
func (t *Enable) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// GetHeapObjectID contains the parameters, and acts as
// a Go receiver, for the CDP command `getHeapObjectId`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getHeapObjectId
type GetHeapObjectID struct {
	// Identifier of the object to get heap object id for.
	ObjectID string `json:"objectId"`
}

// NewGetHeapObjectID constructs a new GetHeapObjectID struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getHeapObjectId
func NewGetHeapObjectID(objectID string) *GetHeapObjectID {
	return &GetHeapObjectID{
		ObjectID: objectID,
	}
}

// GetHeapObjectIDResult contains the browser's response
// to calling the GetHeapObjectID CDP command with Do().
type GetHeapObjectIDResult struct {
	// Id of the heap snapshot object corresponding to the passed remote object id.
	HeapSnapshotObjectID string `json:"heapSnapshotObjectId"`
}

// Do sends the GetHeapObjectID CDP command to a browser,
// and returns the browser's response.
func (t *GetHeapObjectID) Do(ctx context.Context) (*GetHeapObjectIDResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "HeapProfiler.getHeapObjectId", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetHeapObjectID CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetHeapObjectID) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "HeapProfiler.getHeapObjectId", b)
}

// ParseResponse parses the browser's response
// to the GetHeapObjectID CDP command.
func (t *GetHeapObjectID) ParseResponse(m *devtools.Message) (*GetHeapObjectIDResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetHeapObjectIDResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetObjectByHeapObjectID contains the parameters, and acts as
// a Go receiver, for the CDP command `getObjectByHeapObjectId`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getObjectByHeapObjectId
type GetObjectByHeapObjectID struct {
	ObjectID string `json:"objectId"`
	// Symbolic group name that can be used to release multiple objects.
	ObjectGroup string `json:"objectGroup,omitempty"`
}

// NewGetObjectByHeapObjectID constructs a new GetObjectByHeapObjectID struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getObjectByHeapObjectId
func NewGetObjectByHeapObjectID(objectID string) *GetObjectByHeapObjectID {
	return &GetObjectByHeapObjectID{
		ObjectID: objectID,
	}
}

// SetObjectGroup adds or modifies the value of the optional
// parameter `objectGroup` in the GetObjectByHeapObjectID CDP command.
//
// Symbolic group name that can be used to release multiple objects.
func (t *GetObjectByHeapObjectID) SetObjectGroup(v string) *GetObjectByHeapObjectID {
	t.ObjectGroup = v
	return t
}

// GetObjectByHeapObjectIDResult contains the browser's response
// to calling the GetObjectByHeapObjectID CDP command with Do().
type GetObjectByHeapObjectIDResult struct {
	// Evaluation result.
	Result runtime.RemoteObject `json:"result"`
}

// Do sends the GetObjectByHeapObjectID CDP command to a browser,
// and returns the browser's response.
func (t *GetObjectByHeapObjectID) Do(ctx context.Context) (*GetObjectByHeapObjectIDResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "HeapProfiler.getObjectByHeapObjectId", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetObjectByHeapObjectID CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetObjectByHeapObjectID) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "HeapProfiler.getObjectByHeapObjectId", b)
}

// ParseResponse parses the browser's response
// to the GetObjectByHeapObjectID CDP command.
func (t *GetObjectByHeapObjectID) ParseResponse(m *devtools.Message) (*GetObjectByHeapObjectIDResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetObjectByHeapObjectIDResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetSamplingProfile contains the parameters, and acts as
// a Go receiver, for the CDP command `getSamplingProfile`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getSamplingProfile
type GetSamplingProfile struct{}

// NewGetSamplingProfile constructs a new GetSamplingProfile struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getSamplingProfile
func NewGetSamplingProfile() *GetSamplingProfile {
	return &GetSamplingProfile{}
}

// GetSamplingProfileResult contains the browser's response
// to calling the GetSamplingProfile CDP command with Do().
type GetSamplingProfileResult struct {
	// Return the sampling profile being collected.
	Profile SamplingHeapProfile `json:"profile"`
}

// Do sends the GetSamplingProfile CDP command to a browser,
// and returns the browser's response.
func (t *GetSamplingProfile) Do(ctx context.Context) (*GetSamplingProfileResult, error) {
	m, err := devtools.SendAndWait(ctx, "HeapProfiler.getSamplingProfile", nil)
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
	return devtools.Send(ctx, "HeapProfiler.getSamplingProfile", nil)
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

// StartSampling contains the parameters, and acts as
// a Go receiver, for the CDP command `startSampling`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-startSampling
type StartSampling struct {
	// Average sample interval in bytes. Poisson distribution is used for the intervals. The
	// default value is 32768 bytes.
	SamplingInterval float64 `json:"samplingInterval,omitempty"`
}

// NewStartSampling constructs a new StartSampling struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-startSampling
func NewStartSampling() *StartSampling {
	return &StartSampling{}
}

// SetSamplingInterval adds or modifies the value of the optional
// parameter `samplingInterval` in the StartSampling CDP command.
//
// Average sample interval in bytes. Poisson distribution is used for the intervals. The
// default value is 32768 bytes.
func (t *StartSampling) SetSamplingInterval(v float64) *StartSampling {
	t.SamplingInterval = v
	return t
}

// Do sends the StartSampling CDP command to a browser,
// and returns the browser's response.
func (t *StartSampling) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "HeapProfiler.startSampling", b)
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
	return devtools.Send(ctx, "HeapProfiler.startSampling", b)
}

// ParseResponse parses the browser's response
// to the StartSampling CDP command.
func (t *StartSampling) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// StartTrackingHeapObjects contains the parameters, and acts as
// a Go receiver, for the CDP command `startTrackingHeapObjects`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-startTrackingHeapObjects
type StartTrackingHeapObjects struct {
	TrackAllocations bool `json:"trackAllocations,omitempty"`
}

// NewStartTrackingHeapObjects constructs a new StartTrackingHeapObjects struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-startTrackingHeapObjects
func NewStartTrackingHeapObjects() *StartTrackingHeapObjects {
	return &StartTrackingHeapObjects{}
}

// SetTrackAllocations adds or modifies the value of the optional
// parameter `trackAllocations` in the StartTrackingHeapObjects CDP command.
func (t *StartTrackingHeapObjects) SetTrackAllocations(v bool) *StartTrackingHeapObjects {
	t.TrackAllocations = v
	return t
}

// Do sends the StartTrackingHeapObjects CDP command to a browser,
// and returns the browser's response.
func (t *StartTrackingHeapObjects) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "HeapProfiler.startTrackingHeapObjects", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the StartTrackingHeapObjects CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *StartTrackingHeapObjects) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "HeapProfiler.startTrackingHeapObjects", b)
}

// ParseResponse parses the browser's response
// to the StartTrackingHeapObjects CDP command.
func (t *StartTrackingHeapObjects) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// StopSampling contains the parameters, and acts as
// a Go receiver, for the CDP command `stopSampling`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-stopSampling
type StopSampling struct{}

// NewStopSampling constructs a new StopSampling struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-stopSampling
func NewStopSampling() *StopSampling {
	return &StopSampling{}
}

// StopSamplingResult contains the browser's response
// to calling the StopSampling CDP command with Do().
type StopSamplingResult struct {
	// Recorded sampling heap profile.
	Profile SamplingHeapProfile `json:"profile"`
}

// Do sends the StopSampling CDP command to a browser,
// and returns the browser's response.
func (t *StopSampling) Do(ctx context.Context) (*StopSamplingResult, error) {
	m, err := devtools.SendAndWait(ctx, "HeapProfiler.stopSampling", nil)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the StopSampling CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *StopSampling) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "HeapProfiler.stopSampling", nil)
}

// ParseResponse parses the browser's response
// to the StopSampling CDP command.
func (t *StopSampling) ParseResponse(m *devtools.Message) (*StopSamplingResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &StopSamplingResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// StopTrackingHeapObjects contains the parameters, and acts as
// a Go receiver, for the CDP command `stopTrackingHeapObjects`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-stopTrackingHeapObjects
type StopTrackingHeapObjects struct {
	// If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken
	// when the tracking is stopped.
	ReportProgress            bool `json:"reportProgress,omitempty"`
	TreatGlobalObjectsAsRoots bool `json:"treatGlobalObjectsAsRoots,omitempty"`
	// If true, numerical values are included in the snapshot
	CaptureNumericValue bool `json:"captureNumericValue,omitempty"`
}

// NewStopTrackingHeapObjects constructs a new StopTrackingHeapObjects struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-stopTrackingHeapObjects
func NewStopTrackingHeapObjects() *StopTrackingHeapObjects {
	return &StopTrackingHeapObjects{}
}

// SetReportProgress adds or modifies the value of the optional
// parameter `reportProgress` in the StopTrackingHeapObjects CDP command.
//
// If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken
// when the tracking is stopped.
func (t *StopTrackingHeapObjects) SetReportProgress(v bool) *StopTrackingHeapObjects {
	t.ReportProgress = v
	return t
}

// SetTreatGlobalObjectsAsRoots adds or modifies the value of the optional
// parameter `treatGlobalObjectsAsRoots` in the StopTrackingHeapObjects CDP command.
func (t *StopTrackingHeapObjects) SetTreatGlobalObjectsAsRoots(v bool) *StopTrackingHeapObjects {
	t.TreatGlobalObjectsAsRoots = v
	return t
}

// SetCaptureNumericValue adds or modifies the value of the optional
// parameter `captureNumericValue` in the StopTrackingHeapObjects CDP command.
//
// If true, numerical values are included in the snapshot
func (t *StopTrackingHeapObjects) SetCaptureNumericValue(v bool) *StopTrackingHeapObjects {
	t.CaptureNumericValue = v
	return t
}

// Do sends the StopTrackingHeapObjects CDP command to a browser,
// and returns the browser's response.
func (t *StopTrackingHeapObjects) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "HeapProfiler.stopTrackingHeapObjects", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the StopTrackingHeapObjects CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *StopTrackingHeapObjects) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "HeapProfiler.stopTrackingHeapObjects", b)
}

// ParseResponse parses the browser's response
// to the StopTrackingHeapObjects CDP command.
func (t *StopTrackingHeapObjects) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// TakeHeapSnapshot contains the parameters, and acts as
// a Go receiver, for the CDP command `takeHeapSnapshot`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-takeHeapSnapshot
type TakeHeapSnapshot struct {
	// If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken.
	ReportProgress bool `json:"reportProgress,omitempty"`
	// If true, a raw snapshot without artificial roots will be generated
	TreatGlobalObjectsAsRoots bool `json:"treatGlobalObjectsAsRoots,omitempty"`
	// If true, numerical values are included in the snapshot
	CaptureNumericValue bool `json:"captureNumericValue,omitempty"`
}

// NewTakeHeapSnapshot constructs a new TakeHeapSnapshot struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-takeHeapSnapshot
func NewTakeHeapSnapshot() *TakeHeapSnapshot {
	return &TakeHeapSnapshot{}
}

// SetReportProgress adds or modifies the value of the optional
// parameter `reportProgress` in the TakeHeapSnapshot CDP command.
//
// If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken.
func (t *TakeHeapSnapshot) SetReportProgress(v bool) *TakeHeapSnapshot {
	t.ReportProgress = v
	return t
}

// SetTreatGlobalObjectsAsRoots adds or modifies the value of the optional
// parameter `treatGlobalObjectsAsRoots` in the TakeHeapSnapshot CDP command.
//
// If true, a raw snapshot without artificial roots will be generated
func (t *TakeHeapSnapshot) SetTreatGlobalObjectsAsRoots(v bool) *TakeHeapSnapshot {
	t.TreatGlobalObjectsAsRoots = v
	return t
}

// SetCaptureNumericValue adds or modifies the value of the optional
// parameter `captureNumericValue` in the TakeHeapSnapshot CDP command.
//
// If true, numerical values are included in the snapshot
func (t *TakeHeapSnapshot) SetCaptureNumericValue(v bool) *TakeHeapSnapshot {
	t.CaptureNumericValue = v
	return t
}

// Do sends the TakeHeapSnapshot CDP command to a browser,
// and returns the browser's response.
func (t *TakeHeapSnapshot) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "HeapProfiler.takeHeapSnapshot", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the TakeHeapSnapshot CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *TakeHeapSnapshot) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "HeapProfiler.takeHeapSnapshot", b)
}

// ParseResponse parses the browser's response
// to the TakeHeapSnapshot CDP command.
func (t *TakeHeapSnapshot) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}
