package serviceworker

// RegistrationID data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#type-RegistrationID
type RegistrationID string

// ServiceWorkerRegistration data type. ServiceWorker registration.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#type-ServiceWorkerRegistration
type Registration struct {
	RegistrationID string `json:"registrationId"`
	ScopeURL       string `json:"scopeURL"`
	IsDeleted      bool   `json:"isDeleted"`
}

// ServiceWorkerVersionRunningStatus data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#type-ServiceWorkerVersionRunningStatus
type VersionRunningStatus string

// VersionRunningStatus valid values.
const (
	VersionRunningStatusStopped  VersionRunningStatus = "stopped"
	VersionRunningStatusStarting VersionRunningStatus = "starting"
	VersionRunningStatusRunning  VersionRunningStatus = "running"
	VersionRunningStatusStopping VersionRunningStatus = "stopping"
)

// String returns the VersionRunningStatus value as a built-in string.
func (t VersionRunningStatus) String() string {
	return string(t)
}

// ServiceWorkerVersionStatus data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#type-ServiceWorkerVersionStatus
type VersionStatus string

// VersionStatus valid values.
const (
	VersionStatusNew        VersionStatus = "new"
	VersionStatusInstalling VersionStatus = "installing"
	VersionStatusInstalled  VersionStatus = "installed"
	VersionStatusActivating VersionStatus = "activating"
	VersionStatusActivated  VersionStatus = "activated"
	VersionStatusRedundant  VersionStatus = "redundant"
)

// String returns the VersionStatus value as a built-in string.
func (t VersionStatus) String() string {
	return string(t)
}

// ServiceWorkerVersion data type. ServiceWorker version.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#type-ServiceWorkerVersion
type Version struct {
	VersionID      string               `json:"versionId"`
	RegistrationID string               `json:"registrationId"`
	ScriptURL      string               `json:"scriptURL"`
	RunningStatus  VersionRunningStatus `json:"runningStatus"`
	Status         VersionStatus        `json:"status"`
	// The Last-Modified header value of the main script.
	ScriptLastModified float64 `json:"scriptLastModified,omitempty"`
	// The time at which the response headers of the main script were received from the server.
	// For cached script it is the last time the cache entry was validated.
	ScriptResponseTime float64  `json:"scriptResponseTime,omitempty"`
	ControlledClients  []string `json:"controlledClients,omitempty"`
	TargetID           string   `json:"targetId,omitempty"`
}

// ServiceWorkerErrorMessage data type. ServiceWorker error message.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#type-ServiceWorkerErrorMessage
type ErrorMessage struct {
	ErrorMessage   string `json:"errorMessage"`
	RegistrationID string `json:"registrationId"`
	VersionID      string `json:"versionId"`
	SourceURL      string `json:"sourceURL"`
	LineNumber     int64  `json:"lineNumber"`
	ColumnNumber   int64  `json:"columnNumber"`
}
