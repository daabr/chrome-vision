package audits

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
)

// GetEncodedResponse contains the parameters, and acts as
// a Go receiver, for the CDP command `getEncodedResponse`.
//
// Returns the response body and size if it were re-encoded with the specified settings. Only
// applies to images.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#method-getEncodedResponse
type GetEncodedResponse struct {
	// Identifier of the network request to get content for.
	RequestID string `json:"requestId"`
	// The encoding to use.
	Encoding string `json:"encoding"`
	// The quality of the encoding (0-1). (defaults to 1)
	Quality float64 `json:"quality,omitempty"`
	// Whether to only return the size information (defaults to false).
	SizeOnly bool `json:"sizeOnly,omitempty"`
}

// NewGetEncodedResponse constructs a new GetEncodedResponse struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#method-getEncodedResponse
func NewGetEncodedResponse(requestID string, encoding string) *GetEncodedResponse {
	return &GetEncodedResponse{
		RequestID: requestID,
		Encoding:  encoding,
	}
}

// SetQuality adds or modifies the value of the optional
// parameter `quality` in the GetEncodedResponse CDP command.
//
// The quality of the encoding (0-1). (defaults to 1)
func (t *GetEncodedResponse) SetQuality(v float64) *GetEncodedResponse {
	t.Quality = v
	return t
}

// SetSizeOnly adds or modifies the value of the optional
// parameter `sizeOnly` in the GetEncodedResponse CDP command.
//
// Whether to only return the size information (defaults to false).
func (t *GetEncodedResponse) SetSizeOnly(v bool) *GetEncodedResponse {
	t.SizeOnly = v
	return t
}

// GetEncodedResponseResult contains the browser's response
// to calling the GetEncodedResponse CDP command with Do().
type GetEncodedResponseResult struct {
	// The encoded body as a base64 string. Omitted if sizeOnly is true. (Encoded as a base64 string when passed over JSON)
	Body string `json:"body,omitempty"`
	// Size before re-encoding.
	OriginalSize int64 `json:"originalSize"`
	// Size after re-encoding.
	EncodedSize int64 `json:"encodedSize"`
}

// Do sends the GetEncodedResponse CDP command to a browser,
// and returns the browser's response.
func (t *GetEncodedResponse) Do(ctx context.Context) (*GetEncodedResponseResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.SendAndWait(ctx, "Audits.getEncodedResponse", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetEncodedResponseResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Disable contains the parameters, and acts as
// a Go receiver, for the CDP command `disable`.
//
// Disables issues domain, prevents further issues from being reported to the client.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#method-disable
type Disable struct{}

// NewDisable constructs a new Disable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#method-disable
func NewDisable() *Disable {
	return &Disable{}
}

// Do sends the Disable CDP command to a browser,
// and returns the browser's response.
func (t *Disable) Do(ctx context.Context) error {
	response, err := devtools.SendAndWait(ctx, "Audits.disable", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// Enable contains the parameters, and acts as
// a Go receiver, for the CDP command `enable`.
//
// Enables issues domain, sends the issues collected so far to the client by means of the
// `issueAdded` event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#method-enable
type Enable struct{}

// NewEnable constructs a new Enable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#method-enable
func NewEnable() *Enable {
	return &Enable{}
}

// Do sends the Enable CDP command to a browser,
// and returns the browser's response.
func (t *Enable) Do(ctx context.Context) error {
	response, err := devtools.SendAndWait(ctx, "Audits.enable", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// CheckContrast contains the parameters, and acts as
// a Go receiver, for the CDP command `checkContrast`.
//
// Runs the contrast check for the target page. Found issues are reported
// using Audits.issueAdded event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#method-checkContrast
type CheckContrast struct {
	// Whether to report WCAG AAA level issues. Default is false.
	ReportAAA bool `json:"reportAAA,omitempty"`
}

// NewCheckContrast constructs a new CheckContrast struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Audits/#method-checkContrast
func NewCheckContrast() *CheckContrast {
	return &CheckContrast{}
}

// SetReportAAA adds or modifies the value of the optional
// parameter `reportAAA` in the CheckContrast CDP command.
//
// Whether to report WCAG AAA level issues. Default is false.
func (t *CheckContrast) SetReportAAA(v bool) *CheckContrast {
	t.ReportAAA = v
	return t
}

// Do sends the CheckContrast CDP command to a browser,
// and returns the browser's response.
func (t *CheckContrast) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := devtools.SendAndWait(ctx, "Audits.checkContrast", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}
