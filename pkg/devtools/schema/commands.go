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
	response, err := devtools.Send(ctx, "Schema.getDomains", nil)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetDomainsResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}
