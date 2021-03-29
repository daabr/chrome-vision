package inspector

import (
	"context"
	"fmt"
)

// Disable contains the parameters, and acts as
// a Go receiver, for the CDP command `disable`.
//
// Disables inspector domain notifications.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Inspector/#method-disable
type Disable struct{}

// NewDisable constructs a new Disable struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Inspector/#method-disable
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
// Enables inspector domain notifications.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Inspector/#method-enable
type Enable struct{}

// NewEnable constructs a new Enable struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Inspector/#method-enable
func NewEnable() *Enable {
	return &Enable{}
}

// Do sends the Enable CDP command to a browser,
// and returns the browser's response.
func (t *Enable) Do(ctx context.Context) error {
	fmt.Println(ctx)
	return nil
}
