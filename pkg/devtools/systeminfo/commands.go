package systeminfo

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
)

// GetInfo contains the parameters, and acts as
// a Go receiver, for the CDP command `getInfo`.
//
// Returns information about the system.
//
// https://chromedevtools.github.io/devtools-protocol/tot/SystemInfo/#method-getInfo
type GetInfo struct{}

// NewGetInfo constructs a new GetInfo struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/SystemInfo/#method-getInfo
func NewGetInfo() *GetInfo {
	return &GetInfo{}
}

// GetInfoResult contains the browser's response
// to calling the GetInfo CDP command with Do().
type GetInfoResult struct {
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
func (t *GetInfo) Do(ctx context.Context) (*GetInfoResult, error) {
	m, err := devtools.SendAndWait(ctx, "SystemInfo.getInfo", nil)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetInfo CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetInfo) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "SystemInfo.getInfo", nil)
}

// ParseResponse parses the browser's response
// to the GetInfo CDP command.
func (t *GetInfo) ParseResponse(m *devtools.Message) (*GetInfoResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetInfoResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetProcessInfo contains the parameters, and acts as
// a Go receiver, for the CDP command `getProcessInfo`.
//
// Returns information about all running processes.
//
// https://chromedevtools.github.io/devtools-protocol/tot/SystemInfo/#method-getProcessInfo
type GetProcessInfo struct{}

// NewGetProcessInfo constructs a new GetProcessInfo struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/SystemInfo/#method-getProcessInfo
func NewGetProcessInfo() *GetProcessInfo {
	return &GetProcessInfo{}
}

// GetProcessInfoResult contains the browser's response
// to calling the GetProcessInfo CDP command with Do().
type GetProcessInfoResult struct {
	// An array of process info blocks.
	ProcessInfo []ProcessInfo `json:"processInfo"`
}

// Do sends the GetProcessInfo CDP command to a browser,
// and returns the browser's response.
func (t *GetProcessInfo) Do(ctx context.Context) (*GetProcessInfoResult, error) {
	m, err := devtools.SendAndWait(ctx, "SystemInfo.getProcessInfo", nil)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetProcessInfo CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetProcessInfo) Start(ctx context.Context) (chan *devtools.Message, error) {
	return devtools.Send(ctx, "SystemInfo.getProcessInfo", nil)
}

// ParseResponse parses the browser's response
// to the GetProcessInfo CDP command.
func (t *GetProcessInfo) ParseResponse(m *devtools.Message) (*GetProcessInfoResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetProcessInfoResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}
