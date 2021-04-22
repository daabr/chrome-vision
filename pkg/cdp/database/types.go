package database

// Unique identifier of Database object.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Database/#type-DatabaseId
type DatabaseID string

// Database object.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Database/#type-Database
type Database struct {
	// Database ID.
	ID string
	// Database domain.
	Domain string
	// Database name.
	Name string
	// Database version.
	Version string
}

// Database error.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Database/#type-Error
type Error struct {
	// Error message.
	Message string
	// Error code.
	Code int64
}
