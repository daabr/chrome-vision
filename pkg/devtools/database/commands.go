package database

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/devtools"
)

// Disable contains the parameters, and acts as
// a Go receiver, for the CDP command `disable`.
//
// Disables database tracking, prevents database events from being sent to the client.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-disable
type Disable struct{}

// NewDisable constructs a new Disable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-disable
func NewDisable() *Disable {
	return &Disable{}
}

// Do sends the Disable CDP command to a browser,
// and returns the browser's response.
func (t *Disable) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Database.disable", nil)
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
	return devtools.Send(ctx, "Database.disable", nil)
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
// Enables database tracking, database events will now be delivered to the client.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-enable
type Enable struct{}

// NewEnable constructs a new Enable struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-enable
func NewEnable() *Enable {
	return &Enable{}
}

// Do sends the Enable CDP command to a browser,
// and returns the browser's response.
func (t *Enable) Do(ctx context.Context) error {
	m, err := devtools.SendAndWait(ctx, "Database.enable", nil)
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
	return devtools.Send(ctx, "Database.enable", nil)
}

// ParseResponse parses the browser's response
// to the Enable CDP command.
func (t *Enable) ParseResponse(m *devtools.Message) error {
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}
	return nil
}

// ExecuteSQL contains the parameters, and acts as
// a Go receiver, for the CDP command `executeSQL`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-executeSQL
type ExecuteSQL struct {
	DatabaseID string `json:"databaseId"`
	Query      string `json:"query"`
}

// NewExecuteSQL constructs a new ExecuteSQL struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-executeSQL
func NewExecuteSQL(databaseID string, query string) *ExecuteSQL {
	return &ExecuteSQL{
		DatabaseID: databaseID,
		Query:      query,
	}
}

// ExecuteSQLResult contains the browser's response
// to calling the ExecuteSQL CDP command with Do().
type ExecuteSQLResult struct {
	ColumnNames []string          `json:"columnNames,omitempty"`
	Values      []json.RawMessage `json:"values,omitempty"`
	SQLError    *Error            `json:"sqlError,omitempty"`
}

// Do sends the ExecuteSQL CDP command to a browser,
// and returns the browser's response.
func (t *ExecuteSQL) Do(ctx context.Context) (*ExecuteSQLResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Database.executeSQL", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the ExecuteSQL CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *ExecuteSQL) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Database.executeSQL", b)
}

// ParseResponse parses the browser's response
// to the ExecuteSQL CDP command.
func (t *ExecuteSQL) ParseResponse(m *devtools.Message) (*ExecuteSQLResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &ExecuteSQLResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetDatabaseTableNames contains the parameters, and acts as
// a Go receiver, for the CDP command `getDatabaseTableNames`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-getDatabaseTableNames
type GetDatabaseTableNames struct {
	DatabaseID string `json:"databaseId"`
}

// NewGetDatabaseTableNames constructs a new GetDatabaseTableNames struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-getDatabaseTableNames
func NewGetDatabaseTableNames(databaseID string) *GetDatabaseTableNames {
	return &GetDatabaseTableNames{
		DatabaseID: databaseID,
	}
}

// GetDatabaseTableNamesResult contains the browser's response
// to calling the GetDatabaseTableNames CDP command with Do().
type GetDatabaseTableNamesResult struct {
	TableNames []string `json:"tableNames"`
}

// Do sends the GetDatabaseTableNames CDP command to a browser,
// and returns the browser's response.
func (t *GetDatabaseTableNames) Do(ctx context.Context) (*GetDatabaseTableNamesResult, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	m, err := devtools.SendAndWait(ctx, "Database.getDatabaseTableNames", b)
	if err != nil {
		return nil, err
	}
	return t.ParseResponse(m)
}

// Start sends the GetDatabaseTableNames CDP command to a browser,
// and returns a channel to receive the browser's response.
// Callers should close the returned channel on their own,
// although closing unused channels isn't strictly required.
func (t *GetDatabaseTableNames) Start(ctx context.Context) (chan *devtools.Message, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return devtools.Send(ctx, "Database.getDatabaseTableNames", b)
}

// ParseResponse parses the browser's response
// to the GetDatabaseTableNames CDP command.
func (t *GetDatabaseTableNames) ParseResponse(m *devtools.Message) (*GetDatabaseTableNamesResult, error) {
	if m.Error != nil {
		return nil, errors.New(m.Error.Error())
	}
	result := &GetDatabaseTableNamesResult{}
	if err := json.Unmarshal(m.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}
