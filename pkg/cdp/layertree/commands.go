package layertree

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/cdp"
	"github.com/daabr/chrome-vision/pkg/cdp/dom"
)

// CompositingReasons contains the parameters, and acts as
// a Go receiver, for the CDP command `compositingReasons`.
//
// Provides the reasons why the given layer was composited.
//
// https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-compositingReasons
type CompositingReasons struct {
	// The id of the layer for which we want to get the reasons it was composited.
	LayerID string
}

// NewCompositingReasons constructs a new CompositingReasons struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-compositingReasons
func NewCompositingReasons(layerID string) *CompositingReasons {
	return &CompositingReasons{
		LayerID: layerID,
	}
}

// CompositingReasonsResponse contains the browser's response
// to calling the CompositingReasons CDP command with Do().
type CompositingReasonsResponse struct {
	// A list of strings specifying reasons for the given layer to become composited.
	//
	// This CDP parameter is deprecated.
	CompositingReasons []string
	// A list of strings specifying reason IDs for the given layer to become composited.
	CompositingReasonIds []string
}

// Do sends the CompositingReasons CDP command to a browser,
// and returns the browser's response.
func (t *CompositingReasons) Do(ctx context.Context) (*CompositingReasonsResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "LayerTree.compositingReasons", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &CompositingReasonsResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Disable contains the parameters, and acts as
// a Go receiver, for the CDP command `disable`.
//
// Disables compositing tree inspection.
//
// https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-disable
type Disable struct{}

// NewDisable constructs a new Disable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-disable
func NewDisable() *Disable {
	return &Disable{}
}

// Do sends the Disable CDP command to a browser,
// and returns the browser's response.
func (t *Disable) Do(ctx context.Context) error {
	response, err := cdp.Send(ctx, "LayerTree.disable", nil)
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
// Enables compositing tree inspection.
//
// https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-enable
type Enable struct{}

// NewEnable constructs a new Enable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-enable
func NewEnable() *Enable {
	return &Enable{}
}

// Do sends the Enable CDP command to a browser,
// and returns the browser's response.
func (t *Enable) Do(ctx context.Context) error {
	response, err := cdp.Send(ctx, "LayerTree.enable", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// LoadSnapshot contains the parameters, and acts as
// a Go receiver, for the CDP command `loadSnapshot`.
//
// Returns the snapshot identifier.
//
// https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-loadSnapshot
type LoadSnapshot struct {
	// An array of tiles composing the snapshot.
	Tiles []PictureTile
}

// NewLoadSnapshot constructs a new LoadSnapshot struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-loadSnapshot
func NewLoadSnapshot(tiles []PictureTile) *LoadSnapshot {
	return &LoadSnapshot{
		Tiles: tiles,
	}
}

// LoadSnapshotResponse contains the browser's response
// to calling the LoadSnapshot CDP command with Do().
type LoadSnapshotResponse struct {
	// The id of the snapshot.
	SnapshotID string
}

// Do sends the LoadSnapshot CDP command to a browser,
// and returns the browser's response.
func (t *LoadSnapshot) Do(ctx context.Context) (*LoadSnapshotResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "LayerTree.loadSnapshot", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &LoadSnapshotResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// MakeSnapshot contains the parameters, and acts as
// a Go receiver, for the CDP command `makeSnapshot`.
//
// Returns the layer snapshot identifier.
//
// https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-makeSnapshot
type MakeSnapshot struct {
	// The id of the layer.
	LayerID string
}

// NewMakeSnapshot constructs a new MakeSnapshot struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-makeSnapshot
func NewMakeSnapshot(layerID string) *MakeSnapshot {
	return &MakeSnapshot{
		LayerID: layerID,
	}
}

// MakeSnapshotResponse contains the browser's response
// to calling the MakeSnapshot CDP command with Do().
type MakeSnapshotResponse struct {
	// The id of the layer snapshot.
	SnapshotID string
}

// Do sends the MakeSnapshot CDP command to a browser,
// and returns the browser's response.
func (t *MakeSnapshot) Do(ctx context.Context) (*MakeSnapshotResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "LayerTree.makeSnapshot", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &MakeSnapshotResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ProfileSnapshot contains the parameters, and acts as
// a Go receiver, for the CDP command `profileSnapshot`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-profileSnapshot
type ProfileSnapshot struct {
	// The id of the layer snapshot.
	SnapshotID string
	// The maximum number of times to replay the snapshot (1, if not specified).
	MinRepeatCount int64 `json:"minRepeatCount,omitempty"`
	// The minimum duration (in seconds) to replay the snapshot.
	MinDuration float64 `json:"minDuration,omitempty"`
	// The clip rectangle to apply when replaying the snapshot.
	ClipRect *dom.Rect `json:"clipRect,omitempty"`
}

// NewProfileSnapshot constructs a new ProfileSnapshot struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-profileSnapshot
func NewProfileSnapshot(snapshotID string) *ProfileSnapshot {
	return &ProfileSnapshot{
		SnapshotID: snapshotID,
	}
}

// SetMinRepeatCount adds or modifies the value of the optional
// parameter `minRepeatCount` in the ProfileSnapshot CDP command.
//
// The maximum number of times to replay the snapshot (1, if not specified).
func (t *ProfileSnapshot) SetMinRepeatCount(v int64) *ProfileSnapshot {
	t.MinRepeatCount = v
	return t
}

// SetMinDuration adds or modifies the value of the optional
// parameter `minDuration` in the ProfileSnapshot CDP command.
//
// The minimum duration (in seconds) to replay the snapshot.
func (t *ProfileSnapshot) SetMinDuration(v float64) *ProfileSnapshot {
	t.MinDuration = v
	return t
}

// SetClipRect adds or modifies the value of the optional
// parameter `clipRect` in the ProfileSnapshot CDP command.
//
// The clip rectangle to apply when replaying the snapshot.
func (t *ProfileSnapshot) SetClipRect(v dom.Rect) *ProfileSnapshot {
	t.ClipRect = &v
	return t
}

// ProfileSnapshotResponse contains the browser's response
// to calling the ProfileSnapshot CDP command with Do().
type ProfileSnapshotResponse struct {
	// The array of paint profiles, one per run.
	Timings []PaintProfile
}

// Do sends the ProfileSnapshot CDP command to a browser,
// and returns the browser's response.
func (t *ProfileSnapshot) Do(ctx context.Context) (*ProfileSnapshotResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "LayerTree.profileSnapshot", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &ProfileSnapshotResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ReleaseSnapshot contains the parameters, and acts as
// a Go receiver, for the CDP command `releaseSnapshot`.
//
// Releases layer snapshot captured by the back-end.
//
// https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-releaseSnapshot
type ReleaseSnapshot struct {
	// The id of the layer snapshot.
	SnapshotID string
}

// NewReleaseSnapshot constructs a new ReleaseSnapshot struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-releaseSnapshot
func NewReleaseSnapshot(snapshotID string) *ReleaseSnapshot {
	return &ReleaseSnapshot{
		SnapshotID: snapshotID,
	}
}

// Do sends the ReleaseSnapshot CDP command to a browser,
// and returns the browser's response.
func (t *ReleaseSnapshot) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "LayerTree.releaseSnapshot", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// ReplaySnapshot contains the parameters, and acts as
// a Go receiver, for the CDP command `replaySnapshot`.
//
// Replays the layer snapshot and returns the resulting bitmap.
//
// https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-replaySnapshot
type ReplaySnapshot struct {
	// The id of the layer snapshot.
	SnapshotID string
	// The first step to replay from (replay from the very start if not specified).
	FromStep int64 `json:"fromStep,omitempty"`
	// The last step to replay to (replay till the end if not specified).
	ToStep int64 `json:"toStep,omitempty"`
	// The scale to apply while replaying (defaults to 1).
	Scale float64 `json:"scale,omitempty"`
}

// NewReplaySnapshot constructs a new ReplaySnapshot struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-replaySnapshot
func NewReplaySnapshot(snapshotID string) *ReplaySnapshot {
	return &ReplaySnapshot{
		SnapshotID: snapshotID,
	}
}

// SetFromStep adds or modifies the value of the optional
// parameter `fromStep` in the ReplaySnapshot CDP command.
//
// The first step to replay from (replay from the very start if not specified).
func (t *ReplaySnapshot) SetFromStep(v int64) *ReplaySnapshot {
	t.FromStep = v
	return t
}

// SetToStep adds or modifies the value of the optional
// parameter `toStep` in the ReplaySnapshot CDP command.
//
// The last step to replay to (replay till the end if not specified).
func (t *ReplaySnapshot) SetToStep(v int64) *ReplaySnapshot {
	t.ToStep = v
	return t
}

// SetScale adds or modifies the value of the optional
// parameter `scale` in the ReplaySnapshot CDP command.
//
// The scale to apply while replaying (defaults to 1).
func (t *ReplaySnapshot) SetScale(v float64) *ReplaySnapshot {
	t.Scale = v
	return t
}

// ReplaySnapshotResponse contains the browser's response
// to calling the ReplaySnapshot CDP command with Do().
type ReplaySnapshotResponse struct {
	// A data: URL for resulting image.
	DataURL string
}

// Do sends the ReplaySnapshot CDP command to a browser,
// and returns the browser's response.
func (t *ReplaySnapshot) Do(ctx context.Context) (*ReplaySnapshotResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "LayerTree.replaySnapshot", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &ReplaySnapshotResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SnapshotCommandLog contains the parameters, and acts as
// a Go receiver, for the CDP command `snapshotCommandLog`.
//
// Replays the layer snapshot and returns canvas log.
//
// https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-snapshotCommandLog
type SnapshotCommandLog struct {
	// The id of the layer snapshot.
	SnapshotID string
}

// NewSnapshotCommandLog constructs a new SnapshotCommandLog struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-snapshotCommandLog
func NewSnapshotCommandLog(snapshotID string) *SnapshotCommandLog {
	return &SnapshotCommandLog{
		SnapshotID: snapshotID,
	}
}

// SnapshotCommandLogResponse contains the browser's response
// to calling the SnapshotCommandLog CDP command with Do().
type SnapshotCommandLogResponse struct {
	// The array of canvas function calls.
	CommandLog []json.RawMessage
}

// Do sends the SnapshotCommandLog CDP command to a browser,
// and returns the browser's response.
func (t *SnapshotCommandLog) Do(ctx context.Context) (*SnapshotCommandLogResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "LayerTree.snapshotCommandLog", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &SnapshotCommandLogResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}
