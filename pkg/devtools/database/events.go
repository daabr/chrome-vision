package database

// AddDatabase asynchronous event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Database/#event-addDatabase
type AddDatabase struct {
	Database Database `json:"database"`
}
