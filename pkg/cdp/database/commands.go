package database

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/daabr/chrome-vision/pkg/cdp"
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
	response, err := cdp.Send(ctx, "Database.disable", nil)
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
	response, err := cdp.Send(ctx, "Database.enable", nil)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return errors.New(response.Error.Error())
	}
	return nil
}

// ExecuteSQL contains the parameters, and acts as
// a Go receiver, for the CDP command `executeSQL`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-executeSQL
type ExecuteSQL struct {
	DatabaseID DatabaseID `json:"databaseId"`
	Query      string     `json:"query"`
}

// NewExecuteSQL constructs a new ExecuteSQL struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-executeSQL
func NewExecuteSQL(databaseID DatabaseID, query string) *ExecuteSQL {
	return &ExecuteSQL{
		DatabaseID: databaseID,
		Query:      query,
	}
}

// ExecuteSQLResponse contains the browser's response
// to calling the ExecuteSQL CDP command with Do().
type ExecuteSQLResponse struct {
	ColumnNames []string          `json:"columnNames,omitempty"`
	Values      []json.RawMessage `json:"values,omitempty"`
	SqlError    *Error            `json:"sqlError,omitempty"`
}

// Do sends the ExecuteSQL CDP command to a browser,
// and returns the browser's response.
func (t *ExecuteSQL) Do(ctx context.Context) (*ExecuteSQLResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "Database.executeSQL", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &ExecuteSQLResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetDatabaseTableNames contains the parameters, and acts as
// a Go receiver, for the CDP command `getDatabaseTableNames`.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-getDatabaseTableNames
type GetDatabaseTableNames struct {
	DatabaseID DatabaseID `json:"databaseId"`
}

// NewGetDatabaseTableNames constructs a new GetDatabaseTableNames struct instance, with
// all (but only) the required parameters. Optional parameters
// may be added using the builder-like methods below.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Database/#method-getDatabaseTableNames
func NewGetDatabaseTableNames(databaseID DatabaseID) *GetDatabaseTableNames {
	return &GetDatabaseTableNames{
		DatabaseID: databaseID,
	}
}

// GetDatabaseTableNamesResponse contains the browser's response
// to calling the GetDatabaseTableNames CDP command with Do().
type GetDatabaseTableNamesResponse struct {
	TableNames []string `json:"tableNames"`
}

// Do sends the GetDatabaseTableNames CDP command to a browser,
// and returns the browser's response.
func (t *GetDatabaseTableNames) Do(ctx context.Context) (*GetDatabaseTableNamesResponse, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	response, err := cdp.Send(ctx, "Database.getDatabaseTableNames", b)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Error())
	}
	result := &GetDatabaseTableNamesResponse{}
	if err := json.Unmarshal(response.Result, result); err != nil {
		return nil, err
	}
	return result, nil
}
