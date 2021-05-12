package security

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
)

// Disable contains the parameters, and acts as
// a Go receiver, for the CDP command `disable`.
//
// Disables tracking security state changes.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-disable
type Disable struct{}

// NewDisable constructs a new Disable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-disable
func NewDisable() *Disable {
	return &Disable{}
}

// Do sends the Disable CDP command to a browser,
// and returns the browser's response.
func (t *Disable) Do(ctx context.Context) error {
	response, err := devtools.Send(ctx, "Security.disable", nil)
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
// Enables tracking security state changes.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-enable
type Enable struct{}

// NewEnable constructs a new Enable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-enable
func NewEnable() *Enable {
	return &Enable{}
}

// Do sends the Enable CDP command to a browser,
// and returns the browser's response.
func (t *Enable) Do(ctx context.Context) error {
	response, err := devtools.Send(ctx, "Security.enable", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetIgnoreCertificateErrors contains the parameters, and acts as
// a Go receiver, for the CDP command `setIgnoreCertificateErrors`.
//
// Enable/disable whether all certificate errors should be ignored.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-setIgnoreCertificateErrors
//
// This CDP method is experimental.
type SetIgnoreCertificateErrors struct {
	// If true, all certificate errors will be ignored.
	Ignore bool `json:"ignore"`
}

// NewSetIgnoreCertificateErrors constructs a new SetIgnoreCertificateErrors struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-setIgnoreCertificateErrors
//
// This CDP method is experimental.
func NewSetIgnoreCertificateErrors(ignore bool) *SetIgnoreCertificateErrors {
	return &SetIgnoreCertificateErrors{
		Ignore: ignore,
	}
}

// Do sends the SetIgnoreCertificateErrors CDP command to a browser,
// and returns the browser's response.
func (t *SetIgnoreCertificateErrors) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := devtools.Send(ctx, "Security.setIgnoreCertificateErrors", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// HandleCertificateError contains the parameters, and acts as
// a Go receiver, for the CDP command `handleCertificateError`.
//
// Handles a certificate error that fired a certificateError event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-handleCertificateError
//
// This CDP method is deprecated.
type HandleCertificateError struct {
	// The ID of the event.
	EventID int64 `json:"eventId"`
	// The action to take on the certificate error.
	Action CertificateErrorAction `json:"action"`
}

// NewHandleCertificateError constructs a new HandleCertificateError struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-handleCertificateError
//
// This CDP method is deprecated.
func NewHandleCertificateError(eventID int64, action CertificateErrorAction) *HandleCertificateError {
	return &HandleCertificateError{
		EventID: eventID,
		Action:  action,
	}
}

// Do sends the HandleCertificateError CDP command to a browser,
// and returns the browser's response.
func (t *HandleCertificateError) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := devtools.Send(ctx, "Security.handleCertificateError", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetOverrideCertificateErrors contains the parameters, and acts as
// a Go receiver, for the CDP command `setOverrideCertificateErrors`.
//
// Enable/disable overriding certificate errors. If enabled, all certificate error events need to
// be handled by the DevTools client and should be answered with `handleCertificateError` commands.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-setOverrideCertificateErrors
//
// This CDP method is deprecated.
type SetOverrideCertificateErrors struct {
	// If true, certificate errors will be overridden.
	Override bool `json:"override"`
}

// NewSetOverrideCertificateErrors constructs a new SetOverrideCertificateErrors struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-setOverrideCertificateErrors
//
// This CDP method is deprecated.
func NewSetOverrideCertificateErrors(override bool) *SetOverrideCertificateErrors {
	return &SetOverrideCertificateErrors{
		Override: override,
	}
}

// Do sends the SetOverrideCertificateErrors CDP command to a browser,
// and returns the browser's response.
func (t *SetOverrideCertificateErrors) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := devtools.Send(ctx, "Security.setOverrideCertificateErrors", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}
