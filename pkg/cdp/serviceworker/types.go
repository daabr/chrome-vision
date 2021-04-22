package serviceworker

// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#type-RegistrationID
type RegistrationID string

// ServiceWorker registration.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#type-ServiceWorkerRegistration
type ServiceWorkerRegistration struct {
	RegistrationID string
	ScopeURL       string
	IsDeleted      bool
}

// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#type-ServiceWorkerVersionRunningStatus
type ServiceWorkerVersionRunningStatus string

// ServiceWorkerVersionRunningStatus valid values.
const (
	ServiceWorkerVersionRunningStatusStopped  ServiceWorkerVersionRunningStatus = "stopped"
	ServiceWorkerVersionRunningStatusStarting ServiceWorkerVersionRunningStatus = "starting"
	ServiceWorkerVersionRunningStatusRunning  ServiceWorkerVersionRunningStatus = "running"
	ServiceWorkerVersionRunningStatusStopping ServiceWorkerVersionRunningStatus = "stopping"
)

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

// ServiceWorker version.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#type-ServiceWorkerVersion
type ServiceWorkerVersion struct {
	VersionID      string
	RegistrationID string
	ScriptURL      string
	RunningStatus  string
	Status         string
	// The Last-Modified header value of the main script.
	ScriptLastModified float64 `json:"scriptLastModified,omitempty"`
	// The time at which the response headers of the main script were received from the server.
	// For cached script it is the last time the cache entry was validated.
	ScriptResponseTime float64  `json:"scriptResponseTime,omitempty"`
	ControlledClients  []string `json:"controlledClients,omitempty"`
	TargetID           string   `json:"targetId,omitempty"`
}

// ServiceWorker error message.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#type-ServiceWorkerErrorMessage
type ServiceWorkerErrorMessage struct {
	ErrorMessage   string
	RegistrationID string
	VersionID      string
	SourceURL      string
	LineNumber     int64
	ColumnNumber   int64
}
