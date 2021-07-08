package database

// DatabaseID data type. Unique identifier of Database object.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Database/#type-DatabaseId
type ID string

// Database object.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Database/#type-Database
type Database struct {
	// Database ID.
	ID string `json:"id"`
	// Database domain.
	Domain string `json:"domain"`
	// Database name.
	Name string `json:"name"`
	// Database version.
	Version string `json:"version"`
}

// Error data type. Database error.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Database/#type-Error
type Error struct {
	// Error message.
	Message string `json:"message"`
	// Error code.
	Code int64 `json:"code"`
}
