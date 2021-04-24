package serviceworker

// WorkerErrorReported asynchronous event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerErrorReported
type WorkerErrorReported struct {
	ErrorMessage ServiceWorkerErrorMessage `json:"errorMessage"`
}

// WorkerRegistrationUpdated asynchronous event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerRegistrationUpdated
type WorkerRegistrationUpdated struct {
	Registrations []ServiceWorkerRegistration `json:"registrations"`
}

// WorkerVersionUpdated asynchronous event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerVersionUpdated
type WorkerVersionUpdated struct {
	Versions []ServiceWorkerVersion `json:"versions"`
}
