package log

import (
	"context"
	"fmt"
)

// Clear contains the parameters, and acts as
// a Go receiver, for the CDP command `clear`.
//
// Clears the log.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-clear
type Clear struct{}

// NewClear constructs a new Clear struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-clear
func NewClear() *Clear {
	return &Clear{}
}

// Do sends the Clear CDP command to a browser,
// and returns the browser's response.
func (t *Clear) Do(ctx context.Context) error {
	fmt.Println(ctx)
	return nil
}

// Disable contains the parameters, and acts as
// a Go receiver, for the CDP command `disable`.
//
// Disables log domain, prevents further log entries from being reported to the client.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-disable
type Disable struct{}

// NewDisable constructs a new Disable struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-disable
func NewDisable() *Disable {
	return &Disable{}
}

// Do sends the Disable CDP command to a browser,
// and returns the browser's response.
func (t *Disable) Do(ctx context.Context) error {
	fmt.Println(ctx)
	return nil
}

// Enable contains the parameters, and acts as
// a Go receiver, for the CDP command `enable`.
//
// Enables log domain, sends the entries collected so far to the client by means of the
// `entryAdded` notification.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-enable
type Enable struct{}

// NewEnable constructs a new Enable struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-enable
func NewEnable() *Enable {
	return &Enable{}
}

// Do sends the Enable CDP command to a browser,
// and returns the browser's response.
func (t *Enable) Do(ctx context.Context) error {
	fmt.Println(ctx)
	return nil
}

// StartViolationsReport contains the parameters, and acts as
// a Go receiver, for the CDP command `startViolationsReport`.
//
// start violation reporting.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-startViolationsReport
type StartViolationsReport struct {
	// Configuration for violations.
	Config []ViolationSetting `json:"config"`
}

// NewStartViolationsReport constructs a new StartViolationsReport struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-startViolationsReport
func NewStartViolationsReport(config []ViolationSetting) *StartViolationsReport {
	return &StartViolationsReport{
		Config: config,
	}
}

// Do sends the StartViolationsReport CDP command to a browser,
// and returns the browser's response.
func (t *StartViolationsReport) Do(ctx context.Context) error {
	fmt.Println(ctx)
	return nil
}

// StopViolationsReport contains the parameters, and acts as
// a Go receiver, for the CDP command `stopViolationsReport`.
//
// Stop violation reporting.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-stopViolationsReport
type StopViolationsReport struct{}

// NewStopViolationsReport constructs a new StopViolationsReport struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Log/#method-stopViolationsReport
func NewStopViolationsReport() *StopViolationsReport {
	return &StopViolationsReport{}
}

// Do sends the StopViolationsReport CDP command to a browser,
// and returns the browser's response.
func (t *StopViolationsReport) Do(ctx context.Context) error {
	fmt.Println(ctx)
	return nil
}
