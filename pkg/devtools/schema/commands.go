package schema

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
)

// GetDomains contains the parameters, and acts as
// a Go receiver, for the CDP command `getDomains`.
//
// Returns supported domains.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Schema/#method-getDomains
type GetDomains struct{}

// NewGetDomains constructs a new GetDomains struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Schema/#method-getDomains
func NewGetDomains() *GetDomains {
	return &GetDomains{}
}

// GetDomainsResult contains the browser's response
// to calling the GetDomains CDP command with Do().
type GetDomainsResult struct {
	// List of supported domains.
	Domains []Domain `json:"domains"`
}

// Do sends the GetDomains CDP command to a browser,
// and returns the browser's response.
func (t *GetDomains) Do(ctx context.Context) (*GetDomainsResult, error) {
	m, err := devtools.SendAndWait(ctx, "Schema.getDomains", nil)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetDomains CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetDomains) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "Schema.getDomains", nil)
}

// ParseResponse parses the browser's response
// to the GetDomains CDP command.
func (t *GetDomains) ParseResponse(m *devtools.Message) (*GetDomainsResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetDomainsResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}
