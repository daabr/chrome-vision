package animation

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/cdp"
	"github.com/daabr/chrome-vision/pkg/cdp/runtime"
)

// Disable contains the parameters, and acts as
// a Go receiver, for the CDP command `disable`.
//
// Disables animation domain notifications.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-disable
type Disable struct{}

// NewDisable constructs a new Disable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-disable
func NewDisable() *Disable {
	return &Disable{}
}

// Do sends the Disable CDP command to a browser,
// and returns the browser's response.
func (t *Disable) Do(ctx context.Context) error {
	response, err := cdp.Send(ctx, "Animation.disable", nil)
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
// Enables animation domain notifications.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-enable
type Enable struct{}

// NewEnable constructs a new Enable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-enable
func NewEnable() *Enable {
	return &Enable{}
}

// Do sends the Enable CDP command to a browser,
// and returns the browser's response.
func (t *Enable) Do(ctx context.Context) error {
	response, err := cdp.Send(ctx, "Animation.enable", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// GetCurrentTime contains the parameters, and acts as
// a Go receiver, for the CDP command `getCurrentTime`.
//
// Returns the current time of the an animation.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-getCurrentTime
type GetCurrentTime struct {
	// Id of animation.
	ID string `json:"id"`
}

// NewGetCurrentTime constructs a new GetCurrentTime struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-getCurrentTime
func NewGetCurrentTime(id string) *GetCurrentTime {
	return &GetCurrentTime{
		ID: id,
	}
}

// GetCurrentTimeResponse contains the browser's response
// to calling the GetCurrentTime CDP command with Do().
type GetCurrentTimeResponse struct {
	// Current time of the page.
	CurrentTime float64 `json:"currentTime"`
}

// Do sends the GetCurrentTime CDP command to a browser,
// and returns the browser's response.
func (t *GetCurrentTime) Do(ctx context.Context) (*GetCurrentTimeResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "Animation.getCurrentTime", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetCurrentTimeResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetPlaybackRate contains the parameters, and acts as
// a Go receiver, for the CDP command `getPlaybackRate`.
//
// Gets the playback rate of the document timeline.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-getPlaybackRate
type GetPlaybackRate struct{}

// NewGetPlaybackRate constructs a new GetPlaybackRate struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-getPlaybackRate
func NewGetPlaybackRate() *GetPlaybackRate {
	return &GetPlaybackRate{}
}

// GetPlaybackRateResponse contains the browser's response
// to calling the GetPlaybackRate CDP command with Do().
type GetPlaybackRateResponse struct {
	// Playback rate for animations on page.
	PlaybackRate float64 `json:"playbackRate"`
}

// Do sends the GetPlaybackRate CDP command to a browser,
// and returns the browser's response.
func (t *GetPlaybackRate) Do(ctx context.Context) (*GetPlaybackRateResponse, error) {
	response, err := cdp.Send(ctx, "Animation.getPlaybackRate", nil)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetPlaybackRateResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ReleaseAnimations contains the parameters, and acts as
// a Go receiver, for the CDP command `releaseAnimations`.
//
// Releases a set of animations to no longer be manipulated.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-releaseAnimations
type ReleaseAnimations struct {
	// List of animation ids to seek.
	Animations []string `json:"animations"`
}

// NewReleaseAnimations constructs a new ReleaseAnimations struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-releaseAnimations
func NewReleaseAnimations(animations []string) *ReleaseAnimations {
	return &ReleaseAnimations{
		Animations: animations,
	}
}

// Do sends the ReleaseAnimations CDP command to a browser,
// and returns the browser's response.
func (t *ReleaseAnimations) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Animation.releaseAnimations", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// ResolveAnimation contains the parameters, and acts as
// a Go receiver, for the CDP command `resolveAnimation`.
//
// Gets the remote object of the Animation.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-resolveAnimation
type ResolveAnimation struct {
	// Animation id.
	AnimationID string `json:"animationId"`
}

// NewResolveAnimation constructs a new ResolveAnimation struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-resolveAnimation
func NewResolveAnimation(animationId string) *ResolveAnimation {
	return &ResolveAnimation{
		AnimationID: animationId,
	}
}

// ResolveAnimationResponse contains the browser's response
// to calling the ResolveAnimation CDP command with Do().
type ResolveAnimationResponse struct {
	// Corresponding remote object.
	RemoteObject runtime.RemoteObject `json:"remoteObject"`
}

// Do sends the ResolveAnimation CDP command to a browser,
// and returns the browser's response.
func (t *ResolveAnimation) Do(ctx context.Context) (*ResolveAnimationResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "Animation.resolveAnimation", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &ResolveAnimationResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// SeekAnimations contains the parameters, and acts as
// a Go receiver, for the CDP command `seekAnimations`.
//
// Seek a set of animations to a particular time within each animation.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-seekAnimations
type SeekAnimations struct {
	// List of animation ids to seek.
	Animations []string `json:"animations"`
	// Set the current time of each animation.
	CurrentTime float64 `json:"currentTime"`
}

// NewSeekAnimations constructs a new SeekAnimations struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-seekAnimations
func NewSeekAnimations(animations []string, currentTime float64) *SeekAnimations {
	return &SeekAnimations{
		Animations:  animations,
		CurrentTime: currentTime,
	}
}

// Do sends the SeekAnimations CDP command to a browser,
// and returns the browser's response.
func (t *SeekAnimations) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Animation.seekAnimations", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetPaused contains the parameters, and acts as
// a Go receiver, for the CDP command `setPaused`.
//
// Sets the paused state of a set of animations.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-setPaused
type SetPaused struct {
	// Animations to set the pause state of.
	Animations []string `json:"animations"`
	// Paused state to set to.
	Paused bool `json:"paused"`
}

// NewSetPaused constructs a new SetPaused struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-setPaused
func NewSetPaused(animations []string, paused bool) *SetPaused {
	return &SetPaused{
		Animations: animations,
		Paused:     paused,
	}
}

// Do sends the SetPaused CDP command to a browser,
// and returns the browser's response.
func (t *SetPaused) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Animation.setPaused", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetPlaybackRate contains the parameters, and acts as
// a Go receiver, for the CDP command `setPlaybackRate`.
//
// Sets the playback rate of the document timeline.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-setPlaybackRate
type SetPlaybackRate struct {
	// Playback rate for animations on page
	PlaybackRate float64 `json:"playbackRate"`
}

// NewSetPlaybackRate constructs a new SetPlaybackRate struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-setPlaybackRate
func NewSetPlaybackRate(playbackRate float64) *SetPlaybackRate {
	return &SetPlaybackRate{
		PlaybackRate: playbackRate,
	}
}

// Do sends the SetPlaybackRate CDP command to a browser,
// and returns the browser's response.
func (t *SetPlaybackRate) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Animation.setPlaybackRate", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetTiming contains the parameters, and acts as
// a Go receiver, for the CDP command `setTiming`.
//
// Sets the timing of an animation node.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-setTiming
type SetTiming struct {
	// Animation id.
	AnimationID string `json:"animationId"`
	// Duration of the animation.
	Duration float64 `json:"duration"`
	// Delay of the animation.
	Delay float64 `json:"delay"`
}

// NewSetTiming constructs a new SetTiming struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Animation/#method-setTiming
func NewSetTiming(animationId string, duration float64, delay float64) *SetTiming {
	return &SetTiming{
		AnimationID: animationId,
		Duration:    duration,
		Delay:       delay,
	}
}

// Do sends the SetTiming CDP command to a browser,
// and returns the browser's response.
func (t *SetTiming) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Animation.setTiming", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}
