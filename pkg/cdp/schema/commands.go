package schema

import (
	"context"
	"fmt"
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

// GetDomainsResponse contains the browser's response
// to calling the GetDomains CDP command with Do().
type GetDomainsResponse struct {
	// List of supported domains.
	Domains []Domain `json:"domains"`
}

// Do sends the GetDomains CDP command to a browser,
// and returns the browser's response.
func (t *GetDomains) Do(ctx context.Context) (*GetDomainsResponse, error) {
	fmt.Println(ctx)
	return new(GetDomainsResponse), nil
}
