package console

import (
	"context"
	"fmt"
)

// ClearMessages contains the parameters, and acts as
// a Go receiver, for the CDP command `clearMessages`.
//
// Does nothing.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Console/#method-clearMessages
type ClearMessages struct{}

// NewClearMessages constructs a new ClearMessages struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Console/#method-clearMessages
func NewClearMessages() *ClearMessages {
	return &ClearMessages{}
}

// Do sends the ClearMessages CDP command to a browser,
// and returns the browser's response.
func (t *ClearMessages) Do(ctx context.Context) error {
	fmt.Println(ctx)
	return nil
}

// Disable contains the parameters, and acts as
// a Go receiver, for the CDP command `disable`.
//
// Disables console domain, prevents further console messages from being reported to the client.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Console/#method-disable
type Disable struct{}

// NewDisable constructs a new Disable struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Console/#method-disable
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
// Enables console domain, sends the messages collected so far to the client by means of the
// `messageAdded` notification.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Console/#method-enable
type Enable struct{}

// NewEnable constructs a new Enable struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Console/#method-enable
func NewEnable() *Enable {
	return &Enable{}
}

// Do sends the Enable CDP command to a browser,
// and returns the browser's response.
func (t *Enable) Do(ctx context.Context) error {
	fmt.Println(ctx)
	return nil
}
