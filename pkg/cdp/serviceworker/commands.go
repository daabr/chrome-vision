package serviceworker

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/cdp"
)

// DeliverPushMessage contains the parameters, and acts as
// a Go receiver, for the CDP command `deliverPushMessage`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-deliverPushMessage
type DeliverPushMessage struct {
	Origin         string
	RegistrationID string
	Data           string
}

// NewDeliverPushMessage constructs a new DeliverPushMessage struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-deliverPushMessage
func NewDeliverPushMessage(origin string, registrationID string, data string) *DeliverPushMessage {
	return &DeliverPushMessage{
		Origin:         origin,
		RegistrationID: registrationID,
		Data:           data,
	}
}

// Do sends the DeliverPushMessage CDP command to a browser,
// and returns the browser's response.
func (t *DeliverPushMessage) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "ServiceWorker.deliverPushMessage", b)
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
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-disable
type Disable struct{}

// NewDisable constructs a new Disable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-disable
func NewDisable() *Disable {
	return &Disable{}
}

// Do sends the Disable CDP command to a browser,
// and returns the browser's response.
func (t *Disable) Do(ctx context.Context) error {
	response, err := cdp.Send(ctx, "ServiceWorker.disable", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// DispatchSyncEvent contains the parameters, and acts as
// a Go receiver, for the CDP command `dispatchSyncEvent`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-dispatchSyncEvent
type DispatchSyncEvent struct {
	Origin         string
	RegistrationID string
	Tag            string
	LastChance     bool
}

// NewDispatchSyncEvent constructs a new DispatchSyncEvent struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-dispatchSyncEvent
func NewDispatchSyncEvent(origin string, registrationID string, tag string, lastChance bool) *DispatchSyncEvent {
	return &DispatchSyncEvent{
		Origin:         origin,
		RegistrationID: registrationID,
		Tag:            tag,
		LastChance:     lastChance,
	}
}

// Do sends the DispatchSyncEvent CDP command to a browser,
// and returns the browser's response.
func (t *DispatchSyncEvent) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "ServiceWorker.dispatchSyncEvent", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// DispatchPeriodicSyncEvent contains the parameters, and acts as
// a Go receiver, for the CDP command `dispatchPeriodicSyncEvent`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-dispatchPeriodicSyncEvent
type DispatchPeriodicSyncEvent struct {
	Origin         string
	RegistrationID string
	Tag            string
}

// NewDispatchPeriodicSyncEvent constructs a new DispatchPeriodicSyncEvent struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-dispatchPeriodicSyncEvent
func NewDispatchPeriodicSyncEvent(origin string, registrationID string, tag string) *DispatchPeriodicSyncEvent {
	return &DispatchPeriodicSyncEvent{
		Origin:         origin,
		RegistrationID: registrationID,
		Tag:            tag,
	}
}

// Do sends the DispatchPeriodicSyncEvent CDP command to a browser,
// and returns the browser's response.
func (t *DispatchPeriodicSyncEvent) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "ServiceWorker.dispatchPeriodicSyncEvent", b)
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
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-enable
type Enable struct{}

// NewEnable constructs a new Enable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-enable
func NewEnable() *Enable {
	return &Enable{}
}

// Do sends the Enable CDP command to a browser,
// and returns the browser's response.
func (t *Enable) Do(ctx context.Context) error {
	response, err := cdp.Send(ctx, "ServiceWorker.enable", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// InspectWorker contains the parameters, and acts as
// a Go receiver, for the CDP command `inspectWorker`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-inspectWorker
type InspectWorker struct {
	VersionID string
}

// NewInspectWorker constructs a new InspectWorker struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-inspectWorker
func NewInspectWorker(versionID string) *InspectWorker {
	return &InspectWorker{
		VersionID: versionID,
	}
}

// Do sends the InspectWorker CDP command to a browser,
// and returns the browser's response.
func (t *InspectWorker) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "ServiceWorker.inspectWorker", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetForceUpdateOnPageLoad contains the parameters, and acts as
// a Go receiver, for the CDP command `setForceUpdateOnPageLoad`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-setForceUpdateOnPageLoad
type SetForceUpdateOnPageLoad struct {
	ForceUpdateOnPageLoad bool
}

// NewSetForceUpdateOnPageLoad constructs a new SetForceUpdateOnPageLoad struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-setForceUpdateOnPageLoad
func NewSetForceUpdateOnPageLoad(forceUpdateOnPageLoad bool) *SetForceUpdateOnPageLoad {
	return &SetForceUpdateOnPageLoad{
		ForceUpdateOnPageLoad: forceUpdateOnPageLoad,
	}
}

// Do sends the SetForceUpdateOnPageLoad CDP command to a browser,
// and returns the browser's response.
func (t *SetForceUpdateOnPageLoad) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "ServiceWorker.setForceUpdateOnPageLoad", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SkipWaiting contains the parameters, and acts as
// a Go receiver, for the CDP command `skipWaiting`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-skipWaiting
type SkipWaiting struct {
	ScopeURL string
}

// NewSkipWaiting constructs a new SkipWaiting struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-skipWaiting
func NewSkipWaiting(scopeURL string) *SkipWaiting {
	return &SkipWaiting{
		ScopeURL: scopeURL,
	}
}

// Do sends the SkipWaiting CDP command to a browser,
// and returns the browser's response.
func (t *SkipWaiting) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "ServiceWorker.skipWaiting", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// StartWorker contains the parameters, and acts as
// a Go receiver, for the CDP command `startWorker`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-startWorker
type StartWorker struct {
	ScopeURL string
}

// NewStartWorker constructs a new StartWorker struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-startWorker
func NewStartWorker(scopeURL string) *StartWorker {
	return &StartWorker{
		ScopeURL: scopeURL,
	}
}

// Do sends the StartWorker CDP command to a browser,
// and returns the browser's response.
func (t *StartWorker) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "ServiceWorker.startWorker", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// StopAllWorkers contains the parameters, and acts as
// a Go receiver, for the CDP command `stopAllWorkers`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-stopAllWorkers
type StopAllWorkers struct{}

// NewStopAllWorkers constructs a new StopAllWorkers struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-stopAllWorkers
func NewStopAllWorkers() *StopAllWorkers {
	return &StopAllWorkers{}
}

// Do sends the StopAllWorkers CDP command to a browser,
// and returns the browser's response.
func (t *StopAllWorkers) Do(ctx context.Context) error {
	response, err := cdp.Send(ctx, "ServiceWorker.stopAllWorkers", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// StopWorker contains the parameters, and acts as
// a Go receiver, for the CDP command `stopWorker`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-stopWorker
type StopWorker struct {
	VersionID string
}

// NewStopWorker constructs a new StopWorker struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-stopWorker
func NewStopWorker(versionID string) *StopWorker {
	return &StopWorker{
		VersionID: versionID,
	}
}

// Do sends the StopWorker CDP command to a browser,
// and returns the browser's response.
func (t *StopWorker) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "ServiceWorker.stopWorker", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// Unregister contains the parameters, and acts as
// a Go receiver, for the CDP command `unregister`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-unregister
type Unregister struct {
	ScopeURL string
}

// NewUnregister constructs a new Unregister struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-unregister
func NewUnregister(scopeURL string) *Unregister {
	return &Unregister{
		ScopeURL: scopeURL,
	}
}

// Do sends the Unregister CDP command to a browser,
// and returns the browser's response.
func (t *Unregister) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "ServiceWorker.unregister", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// UpdateRegistration contains the parameters, and acts as
// a Go receiver, for the CDP command `updateRegistration`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-updateRegistration
type UpdateRegistration struct {
	ScopeURL string
}

// NewUpdateRegistration constructs a new UpdateRegistration struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#method-updateRegistration
func NewUpdateRegistration(scopeURL string) *UpdateRegistration {
	return &UpdateRegistration{
		ScopeURL: scopeURL,
	}
}

// Do sends the UpdateRegistration CDP command to a browser,
// and returns the browser's response.
func (t *UpdateRegistration) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "ServiceWorker.updateRegistration", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}
