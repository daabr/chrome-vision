package serviceworker

import "github.com/daabr/chrome-vision/pkg/cdp"

// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#type-RegistrationID
type RegistrationID string

// ServiceWorker registration.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#type-ServiceWorkerRegistration
type ServiceWorkerRegistration struct {
	RegistrationID RegistrationID `json:"registrationId"`
	ScopeURL       string         `json:"scopeURL"`
	IsDeleted      bool           `json:"isDeleted"`
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
	VersionID      string                            `json:"versionId"`
	RegistrationID RegistrationID                    `json:"registrationId"`
	ScriptURL      string                            `json:"scriptURL"`
	RunningStatus  ServiceWorkerVersionRunningStatus `json:"runningStatus"`
	Status         ServiceWorkerVersionStatus        `json:"status"`
	// The Last-Modified header value of the main script.
	ScriptLastModified float64 `json:"scriptLastModified,omitempty"`
	// The time at which the response headers of the main script were received from the server.
	// For cached script it is the last time the cache entry was validated.
	ScriptResponseTime float64        `json:"scriptResponseTime,omitempty"`
	ControlledClients  []cdp.TargetID `json:"controlledClients,omitempty"`
	TargetID           *cdp.TargetID  `json:"targetId,omitempty"`
}

// ServiceWorker error message.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#type-ServiceWorkerErrorMessage
type ServiceWorkerErrorMessage struct {
	ErrorMessage   string         `json:"errorMessage"`
	RegistrationID RegistrationID `json:"registrationId"`
	VersionID      string         `json:"versionId"`
	SourceURL      string         `json:"sourceURL"`
	LineNumber     int64          `json:"lineNumber"`
	ColumnNumber   int64          `json:"columnNumber"`
}
