package performancetimeline

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
)

// Enable contains the parameters, and acts as
// a Go receiver, for the CDP command `enable`.
//
// Previously buffered events would be reported before method returns.
// See also: timelineEventAdded
//
// https://chromedevtools.github.io/devtools-protocol/tot/PerformanceTimeline/#method-enable
type Enable struct {
	// The types of event to report, as specified in
	// https://w3c.github.io/performance-timeline/#dom-performanceentry-entrytype
	// The specified filter overrides any previous filters, passing empty
	// filter disables recording.
	// Note that not all types exposed to the web platform are currently supported.
	EventTypes []string `json:"eventTypes"`
}

// NewEnable constructs a new Enable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/PerformanceTimeline/#method-enable
func NewEnable(eventTypes []string) *Enable {
	return &Enable{
		EventTypes: eventTypes,
	}
}

// Do sends the Enable CDP command to a browser,
// and returns the browser's response.
func (t *Enable) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "PerformanceTimeline.enable", b)
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
	return devtools.Send(ctx, "PerformanceTimeline.enable", b)
}

// ParseResponse parses the browser's response
// to the Enable CDP command.
func (t *Enable) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}
