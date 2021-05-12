package io

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
	"github.com/daabr/chrome-vision/pkg/devtools/runtime"
)

// Close contains the parameters, and acts as
// a Go receiver, for the CDP command `close`.
//
// Close the stream, discard any temporary backing storage.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IO/#method-close
type Close struct {
	// Handle of the stream to close.
	Handle string `json:"handle"`
}

// NewClose constructs a new Close struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IO/#method-close
func NewClose(handle string) *Close {
	return &Close{
		Handle: handle,
	}
}

// Do sends the Close CDP command to a browser,
// and returns the browser's response.
func (t *Close) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := devtools.Send(ctx, "IO.close", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// Read contains the parameters, and acts as
// a Go receiver, for the CDP command `read`.
//
// Read a chunk of the stream
//
// https://chromedevtools.github.io/devtools-protocol/tot/IO/#method-read
type Read struct {
	// Handle of the stream to read.
	Handle string `json:"handle"`
	// Seek to the specified offset before reading (if not specificed, proceed with offset
	// following the last read). Some types of streams may only support sequential reads.
	Offset int64 `json:"offset,omitempty"`
	// Maximum number of bytes to read (left upon the agent discretion if not specified).
	Size int64 `json:"size,omitempty"`
}

// NewRead constructs a new Read struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IO/#method-read
func NewRead(handle string) *Read {
	return &Read{
		Handle: handle,
	}
}

// SetOffset adds or modifies the value of the optional
// parameter `offset` in the Read CDP command.
//
// Seek to the specified offset before reading (if not specificed, proceed with offset
// following the last read). Some types of streams may only support sequential reads.
func (t *Read) SetOffset(v int64) *Read {
	t.Offset = v
	return t
}

// SetSize adds or modifies the value of the optional
// parameter `size` in the Read CDP command.
//
// Maximum number of bytes to read (left upon the agent discretion if not specified).
func (t *Read) SetSize(v int64) *Read {
	t.Size = v
	return t
}

// ReadResult contains the browser's response
// to calling the Read CDP command with Do().
type ReadResult struct {
	// Set if the data is base64-encoded
	Base64Encoded bool `json:"base64Encoded,omitempty"`
	// Data that were read.
	Data string `json:"data"`
	// Set if the end-of-file condition occurred while reading.
	Eof bool `json:"eof"`
}

// Do sends the Read CDP command to a browser,
// and returns the browser's response.
func (t *Read) Do(ctx context.Context) (*ReadResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.Send(ctx, "IO.read", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &ReadResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ResolveBlob contains the parameters, and acts as
// a Go receiver, for the CDP command `resolveBlob`.
//
// Return UUID of Blob object specified by a remote object id.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IO/#method-resolveBlob
type ResolveBlob struct {
	// Object id of a Blob object wrapper.
	ObjectID runtime.RemoteObjectID `json:"objectId"`
}

// NewResolveBlob constructs a new ResolveBlob struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/IO/#method-resolveBlob
func NewResolveBlob(objectID runtime.RemoteObjectID) *ResolveBlob {
	return &ResolveBlob{
		ObjectID: objectID,
	}
}

// ResolveBlobResult contains the browser's response
// to calling the ResolveBlob CDP command with Do().
type ResolveBlobResult struct {
	// UUID of the specified Blob.
	Uuid string `json:"uuid"`
}

// Do sends the ResolveBlob CDP command to a browser,
// and returns the browser's response.
func (t *ResolveBlob) Do(ctx context.Context) (*ResolveBlobResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := devtools.Send(ctx, "IO.resolveBlob", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &ResolveBlobResult{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}
