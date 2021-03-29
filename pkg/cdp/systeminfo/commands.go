package systeminfo

import (
	"context"
	"fmt"
)

// GetInfo contains the parameters, and acts as
// a Go receiver, for the CDP command `getInfo`.
//
// Returns information about the system.
//
// https://chromedevtools.github.io/devtools-protocol/tot/SystemInfo/#method-getInfo
type GetInfo struct{}

// NewGetInfo constructs a new GetInfo struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/SystemInfo/#method-getInfo
func NewGetInfo() *GetInfo {
	return &GetInfo{}
}

// GetInfoResponse contains the browser's response
// to calling the GetInfo CDP command with Do().
type GetInfoResponse struct {
	// Information about the GPUs on the system.
	Gpu GPUInfo `json:"gpu"`
	// A platform-dependent description of the model of the machine. On Mac OS, this is, for
	// example, 'MacBookPro'. Will be the empty string if not supported.
	ModelName string `json:"modelName"`
	// A platform-dependent description of the version of the machine. On Mac OS, this is, for
	// example, '10.1'. Will be the empty string if not supported.
	ModelVersion string `json:"modelVersion"`
	// The command line string used to launch the browser. Will be the empty string if not
	// supported.
	CommandLine string `json:"commandLine"`
}

// Do sends the GetInfo CDP command to a browser,
// and returns the browser's response.
func (t *GetInfo) Do(ctx context.Context) (*GetInfoResponse, error) {
	fmt.Println(ctx)
	return new(GetInfoResponse), nil
}

// GetProcessInfo contains the parameters, and acts as
// a Go receiver, for the CDP command `getProcessInfo`.
//
// Returns information about all running processes.
//
// https://chromedevtools.github.io/devtools-protocol/tot/SystemInfo/#method-getProcessInfo
type GetProcessInfo struct{}

// NewGetProcessInfo constructs a new GetProcessInfo struct instance, with
// all the required parameters, and only them. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/SystemInfo/#method-getProcessInfo
func NewGetProcessInfo() *GetProcessInfo {
	return &GetProcessInfo{}
}

// GetProcessInfoResponse contains the browser's response
// to calling the GetProcessInfo CDP command with Do().
type GetProcessInfoResponse struct {
	// An array of process info blocks.
	ProcessInfo []ProcessInfo `json:"processInfo"`
}

// Do sends the GetProcessInfo CDP command to a browser,
// and returns the browser's response.
func (t *GetProcessInfo) Do(ctx context.Context) (*GetProcessInfoResponse, error) {
	fmt.Println(ctx)
	return new(GetProcessInfoResponse), nil
}
