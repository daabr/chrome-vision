package heapprofiler

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/cdp"
	"github.com/daabr/chrome-vision/pkg/cdp/runtime"
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
	HeapObjectID HeapSnapshotObjectID `json:"heapObjectId"`
}

// NewAddInspectedHeapObject constructs a new AddInspectedHeapObject struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-addInspectedHeapObject
func NewAddInspectedHeapObject(heapObjectId HeapSnapshotObjectID) *AddInspectedHeapObject {
	return &AddInspectedHeapObject{
		HeapObjectID: heapObjectId,
	}
}

// Do sends the AddInspectedHeapObject CDP command to a browser,
// and returns the browser's response.
func (t *AddInspectedHeapObject) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "HeapProfiler.addInspectedHeapObject", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
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
	response, err := cdp.Send(ctx, "HeapProfiler.collectGarbage", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
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
	response, err := cdp.Send(ctx, "HeapProfiler.disable", nil)
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
	response, err := cdp.Send(ctx, "HeapProfiler.enable", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// GetHeapObjectID contains the parameters, and acts as
// a Go receiver, for the CDP command `getHeapObjectId`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getHeapObjectId
type GetHeapObjectID struct {
	// Identifier of the object to get heap object id for.
	ObjectID runtime.RemoteObjectID `json:"objectId"`
}

// NewGetHeapObjectID constructs a new GetHeapObjectID struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getHeapObjectId
func NewGetHeapObjectID(objectId runtime.RemoteObjectID) *GetHeapObjectID {
	return &GetHeapObjectID{
		ObjectID: objectId,
	}
}

// GetHeapObjectIDResponse contains the browser's response
// to calling the GetHeapObjectID CDP command with Do().
type GetHeapObjectIDResponse struct {
	// Id of the heap snapshot object corresponding to the passed remote object id.
	HeapSnapshotObjectID HeapSnapshotObjectID `json:"heapSnapshotObjectId"`
}

// Do sends the GetHeapObjectID CDP command to a browser,
// and returns the browser's response.
func (t *GetHeapObjectID) Do(ctx context.Context) (*GetHeapObjectIDResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "HeapProfiler.getHeapObjectId", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetHeapObjectIDResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetObjectByHeapObjectID contains the parameters, and acts as
// a Go receiver, for the CDP command `getObjectByHeapObjectId`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getObjectByHeapObjectId
type GetObjectByHeapObjectID struct {
	ObjectID HeapSnapshotObjectID `json:"objectId"`
	// Symbolic group name that can be used to release multiple objects.
	ObjectGroup string `json:"objectGroup,omitempty"`
}

// NewGetObjectByHeapObjectID constructs a new GetObjectByHeapObjectID struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getObjectByHeapObjectId
func NewGetObjectByHeapObjectID(objectId HeapSnapshotObjectID) *GetObjectByHeapObjectID {
	return &GetObjectByHeapObjectID{
		ObjectID: objectId,
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

// GetObjectByHeapObjectIDResponse contains the browser's response
// to calling the GetObjectByHeapObjectID CDP command with Do().
type GetObjectByHeapObjectIDResponse struct {
	// Evaluation result.
	Result runtime.RemoteObject `json:"result"`
}

// Do sends the GetObjectByHeapObjectID CDP command to a browser,
// and returns the browser's response.
func (t *GetObjectByHeapObjectID) Do(ctx context.Context) (*GetObjectByHeapObjectIDResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "HeapProfiler.getObjectByHeapObjectId", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetObjectByHeapObjectIDResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
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

// GetSamplingProfileResponse contains the browser's response
// to calling the GetSamplingProfile CDP command with Do().
type GetSamplingProfileResponse struct {
	// Return the sampling profile being collected.
	Profile SamplingHeapProfile `json:"profile"`
}

// Do sends the GetSamplingProfile CDP command to a browser,
// and returns the browser's response.
func (t *GetSamplingProfile) Do(ctx context.Context) (*GetSamplingProfileResponse, error) {
	response, err := cdp.Send(ctx, "HeapProfiler.getSamplingProfile", nil)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetSamplingProfileResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
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
	response, err := cdp.Send(ctx, "HeapProfiler.startSampling", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
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
	response, err := cdp.Send(ctx, "HeapProfiler.startTrackingHeapObjects", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
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

// StopSamplingResponse contains the browser's response
// to calling the StopSampling CDP command with Do().
type StopSamplingResponse struct {
	// Recorded sampling heap profile.
	Profile SamplingHeapProfile `json:"profile"`
}

// Do sends the StopSampling CDP command to a browser,
// and returns the browser's response.
func (t *StopSampling) Do(ctx context.Context) (*StopSamplingResponse, error) {
	response, err := cdp.Send(ctx, "HeapProfiler.stopSampling", nil)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &StopSamplingResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
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

// Do sends the StopTrackingHeapObjects CDP command to a browser,
// and returns the browser's response.
func (t *StopTrackingHeapObjects) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "HeapProfiler.stopTrackingHeapObjects", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
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
	// If true, a raw snapshot without artifical roots will be generated
	TreatGlobalObjectsAsRoots bool `json:"treatGlobalObjectsAsRoots,omitempty"`
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
// If true, a raw snapshot without artifical roots will be generated
func (t *TakeHeapSnapshot) SetTreatGlobalObjectsAsRoots(v bool) *TakeHeapSnapshot {
	t.TreatGlobalObjectsAsRoots = v
	return t
}

// Do sends the TakeHeapSnapshot CDP command to a browser,
// and returns the browser's response.
func (t *TakeHeapSnapshot) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "HeapProfiler.takeHeapSnapshot", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}
