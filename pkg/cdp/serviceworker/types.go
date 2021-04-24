package serviceworker

// RegistrationID data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#type-RegistrationID
type RegistrationID string

// ServiceWorkerRegistration data type. ServiceWorker registration.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#type-ServiceWorkerRegistration
type ServiceWorkerRegistration struct {
	RegistrationID string `json:"registrationId"`
	ScopeURL       string `json:"scopeURL"`
	IsDeleted      bool   `json:"isDeleted"`
}

// ServiceWorkerVersionRunningStatus data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#type-ServiceWorkerVersionRunningStatus
type ServiceWorkerVersionRunningStatus string

// ServiceWorkerVersionRunningStatus valid values.
const (
	ServiceWorkerVersionRunningStatusStopped  ServiceWorkerVersionRunningStatus = "stopped"
	ServiceWorkerVersionRunningStatusStarting ServiceWorkerVersionRunningStatus = "starting"
	ServiceWorkerVersionRunningStatusRunning  ServiceWorkerVersionRunningStatus = "running"
	ServiceWorkerVersionRunningStatusStopping ServiceWorkerVersionRunningStatus = "stopping"
)

// ServiceWorkerVersionStatus data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#type-ServiceWorkerVersionStatus
type ServiceWorkerVersionStatus string

// ServiceWorkerVersionStatus valid values.
const (
	ServiceWorkerVersionStatusNew        ServiceWorkerVersionStatus = "new"
	ServiceWorkerVersionStatusInstalling ServiceWorkerVersionStatus = "installing"
	ServiceWorkerVersionStatusInstalled  ServiceWorkerVersionStatus = "installed"
	ServiceWorkerVersionStatusActivating ServiceWorkerVersionStatus = "activating"
	ServiceWorkerVersionStatusActivated  ServiceWorkerVersionStatus = "activated"
	ServiceWorkerVersionStatusRedundant  ServiceWorkerVersionStatus = "redundant"
)

// ServiceWorkerVersion data type. ServiceWorker version.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#type-ServiceWorkerVersion
type ServiceWorkerVersion struct {
	VersionID      string `json:"versionId"`
	RegistrationID string `json:"registrationId"`
	ScriptURL      string `json:"scriptURL"`
	RunningStatus  string `json:"runningStatus"`
	Status         string `json:"status"`
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
type ServiceWorkerErrorMessage struct {
	ErrorMessage   string `json:"errorMessage"`
	RegistrationID string `json:"registrationId"`
	VersionID      string `json:"versionId"`
	SourceURL      string `json:"sourceURL"`
	LineNumber     int64  `json:"lineNumber"`
	ColumnNumber   int64  `json:"columnNumber"`
}
