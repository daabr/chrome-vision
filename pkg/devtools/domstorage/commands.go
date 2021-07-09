package domstorage

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
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
func NewClear(storageID StorageID) *Clear {
	return &Clear{
		StorageID: storageID,
	}
}

// Do sends the Clear CDP command to a browser,
// and returns the browser's response.
func (t *Clear) Do(ctx context.Context) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	m, err := devtools.SendAndWait(ctx, "DOMStorage.clear", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the Clear CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *Clear) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "DOMStorage.clear", b)
}

// ParseResponse parses the browser's response
// to the Clear CDP command.
func (t *Clear) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
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
	m, err := devtools.SendAndWait(ctx, "DOMStorage.disable", nil)
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
	return devtools.Send(ctx, "DOMStorage.disable", nil)
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
	m, err := devtools.SendAndWait(ctx, "DOMStorage.enable", nil)
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
	return devtools.Send(ctx, "DOMStorage.enable", nil)
}

// ParseResponse parses the browser's response
// to the Enable CDP command.
func (t *Enable) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
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
func NewGetDOMStorageItems(storageID StorageID) *GetDOMStorageItems {
	return &GetDOMStorageItems{
		StorageID: storageID,
	}
}

// GetDOMStorageItemsResult contains the browser's response
// to calling the GetDOMStorageItems CDP command with Do().
type GetDOMStorageItemsResult struct {
	Entries []Item `json:"entries"`
}

// Do sends the GetDOMStorageItems CDP command to a browser,
// and returns the browser's response.
func (t *GetDOMStorageItems) Do(ctx context.Context) (*GetDOMStorageItemsResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "DOMStorage.getDOMStorageItems", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetDOMStorageItems CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetDOMStorageItems) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "DOMStorage.getDOMStorageItems", b)
}

// ParseResponse parses the browser's response
// to the GetDOMStorageItems CDP command.
func (t *GetDOMStorageItems) ParseResponse(m *devtools.Message) (*GetDOMStorageItemsResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetDOMStorageItemsResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
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
func NewRemoveDOMStorageItem(storageID StorageID, key string) *RemoveDOMStorageItem {
	return &RemoveDOMStorageItem{
		StorageID: storageID,
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
	m, err := devtools.SendAndWait(ctx, "DOMStorage.removeDOMStorageItem", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the RemoveDOMStorageItem CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *RemoveDOMStorageItem) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "DOMStorage.removeDOMStorageItem", b)
}

// ParseResponse parses the browser's response
// to the RemoveDOMStorageItem CDP command.
func (t *RemoveDOMStorageItem) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
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
func NewSetDOMStorageItem(storageID StorageID, key string, value string) *SetDOMStorageItem {
	return &SetDOMStorageItem{
		StorageID: storageID,
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
	m, err := devtools.SendAndWait(ctx, "DOMStorage.setDOMStorageItem", b)
	if err != nil {
		return err
	}
	return t.ParseResponse(m)
}

// Start sends the SetDOMStorageItem CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *SetDOMStorageItem) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "DOMStorage.setDOMStorageItem", b)
}

// ParseResponse parses the browser's response
// to the SetDOMStorageItem CDP command.
func (t *SetDOMStorageItem) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}
