package performance

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
)

// Disable contains the parameters, and acts as
// a Go receiver, for the CDP command `disable`.
//
// Disable collecting and reporting metrics.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Performance/#method-disable
type Disable struct{}

// NewDisable constructs a new Disable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Performance/#method-disable
func NewDisable() *Disable {
	return &Disable{}
}

// Do sends the Disable CDP command to a browser,
// and returns the browser's response.
func (t *Disable) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Performance.disable", nil)
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
	return devtools.Send(ctx, "Performance.disable", nil)
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
// Enable collecting and reporting metrics.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Performance/#method-enable
type Enable struct {
	// Time domain to use for collecting and reporting duration metrics.
	TimeDomain string `json:"timeDomain,omitempty"`
}

// NewEnable constructs a new Enable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Performance/#method-enable
func NewEnable() *Enable {
	return &Enable{}
}

// SetTimeDomain adds or modifies the value of the optional
// parameter `timeDomain` in the Enable CDP command.
//
// Time domain to use for collecting and reporting duration metrics.
func (t *Enable) SetTimeDomain(v string) *Enable {
	t.TimeDomain = v
	return t
}

// Do sends the Enable CDP command to a browser,
// and returns the browser's response.
func (t *Enable) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Performance.enable", b)
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
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Performance.enable", b)
}

// ParseResponse parses the browser's response
// to the Enable CDP command.
func (t *Enable) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// SetTimeDomain contains the parameters, and acts as
// a Go receiver, for the CDP command `setTimeDomain`.
//
// Sets time domain to use for collecting and reporting duration metrics.
// Note that this must be called before enabling metrics collection. Calling
// this method while metrics collection is enabled returns an error.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Performance/#method-setTimeDomain
//
// This CDP method is deprecated.
// This CDP method is experimental.
type SetTimeDomain struct {
	// Time domain
	TimeDomain string `json:"timeDomain"`
}

// NewSetTimeDomain constructs a new SetTimeDomain struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Performance/#method-setTimeDomain
//
// This CDP method is deprecated.
// This CDP method is experimental.
func NewSetTimeDomain(timeDomain string) *SetTimeDomain {
	return &SetTimeDomain{
		TimeDomain: timeDomain,
	}
}

// Do sends the SetTimeDomain CDP command to a browser,
// and returns the browser's response.
func (t *SetTimeDomain) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "Performance.setTimeDomain", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetTimeDomain CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetTimeDomain) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Performance.setTimeDomain", b)
}

// ParseResponse parses the browser's response
// to the SetTimeDomain CDP command.
func (t *SetTimeDomain) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// GetMetrics contains the parameters, and acts as
// a Go receiver, for the CDP command `getMetrics`.
//
// Retrieve current values of run-time metrics.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Performance/#method-getMetrics
type GetMetrics struct{}

// NewGetMetrics constructs a new GetMetrics struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Performance/#method-getMetrics
func NewGetMetrics() *GetMetrics {
	return &GetMetrics{}
}

// GetMetricsResult contains the browser's response
// to calling the GetMetrics CDP command with Do().
type GetMetricsResult struct {
	// Current values for run-time metrics.
	Metrics []Metric `json:"metrics"`
}

// Do sends the GetMetrics CDP command to a browser,
// and returns the browser's response.
func (t *GetMetrics) Do(ctx context.Context) (*GetMetricsResult, error) {
	m, err := devtools.SendAndWait(ctx, "Performance.getMetrics", nil)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetMetrics CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetMetrics) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Performance.getMetrics", nil)
}

// ParseResponse parses the browser's response
// to the GetMetrics CDP command.
func (t *GetMetrics) ParseResponse(m *devtools.Message) (*GetMetricsResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetMetricsResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}
