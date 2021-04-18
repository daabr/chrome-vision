package domstorage

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/cdp"
)

// Clear contains the parameters, and acts as
// a Go receiver, for the CDP command `clear`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-clear
type Clear struct {
	StorageID StorageID `json:"storageId"`
}

// NewClear constructs a new Clear struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-clear
func NewClear(storageId StorageID) *Clear {
	return &Clear{
		StorageID: storageId,
	}
}

// Do sends the Clear CDP command to a browser,
// and returns the browser's response.
func (t *Clear) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "Clear", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// Disable contains the parameters, and acts as
// a Go receiver, for the CDP command `disable`.
//
// Disables storage tracking, prevents storage events from being sent to the client.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-disable
type Disable struct{}

// NewDisable constructs a new Disable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-disable
func NewDisable() *Disable {
	return &Disable{}
}

// Do sends the Disable CDP command to a browser,
// and returns the browser's response.
func (t *Disable) Do(ctx context.Context) error {
	response, err := cdp.Send(ctx, "Disable", nil)
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
// Enables storage tracking, storage events will now be delivered to the client.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-enable
type Enable struct{}

// NewEnable constructs a new Enable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-enable
func NewEnable() *Enable {
	return &Enable{}
}

// Do sends the Enable CDP command to a browser,
// and returns the browser's response.
func (t *Enable) Do(ctx context.Context) error {
	response, err := cdp.Send(ctx, "Enable", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// GetDOMStorageItems contains the parameters, and acts as
// a Go receiver, for the CDP command `getDOMStorageItems`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-getDOMStorageItems
type GetDOMStorageItems struct {
	StorageID StorageID `json:"storageId"`
}

// NewGetDOMStorageItems constructs a new GetDOMStorageItems struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-getDOMStorageItems
func NewGetDOMStorageItems(storageId StorageID) *GetDOMStorageItems {
	return &GetDOMStorageItems{
		StorageID: storageId,
	}
}

// GetDOMStorageItemsResponse contains the browser's response
// to calling the GetDOMStorageItems CDP command with Do().
type GetDOMStorageItemsResponse struct {
	Entries []Item `json:"entries"`
}

// Do sends the GetDOMStorageItems CDP command to a browser,
// and returns the browser's response.
func (t *GetDOMStorageItems) Do(ctx context.Context) (*GetDOMStorageItemsResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "GetDOMStorageItems", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetDOMStorageItemsResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// RemoveDOMStorageItem contains the parameters, and acts as
// a Go receiver, for the CDP command `removeDOMStorageItem`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-removeDOMStorageItem
type RemoveDOMStorageItem struct {
	StorageID StorageID `json:"storageId"`
	Key       string    `json:"key"`
}

// NewRemoveDOMStorageItem constructs a new RemoveDOMStorageItem struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-removeDOMStorageItem
func NewRemoveDOMStorageItem(storageId StorageID, key string) *RemoveDOMStorageItem {
	return &RemoveDOMStorageItem{
		StorageID: storageId,
		Key:       key,
	}
}

// Do sends the RemoveDOMStorageItem CDP command to a browser,
// and returns the browser's response.
func (t *RemoveDOMStorageItem) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "RemoveDOMStorageItem", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// SetDOMStorageItem contains the parameters, and acts as
// a Go receiver, for the CDP command `setDOMStorageItem`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-setDOMStorageItem
type SetDOMStorageItem struct {
	StorageID StorageID `json:"storageId"`
	Key       string    `json:"key"`
	Value     string    `json:"value"`
}

// NewSetDOMStorageItem constructs a new SetDOMStorageItem struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/DOMStorage/#method-setDOMStorageItem
func NewSetDOMStorageItem(storageId StorageID, key string, value string) *SetDOMStorageItem {
	return &SetDOMStorageItem{
		StorageID: storageId,
		Key:       key,
		Value:     value,
	}
}

// Do sends the SetDOMStorageItem CDP command to a browser,
// and returns the browser's response.
func (t *SetDOMStorageItem) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	response, err := cdp.Send(ctx, "SetDOMStorageItem", b)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}
