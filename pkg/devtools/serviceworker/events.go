package serviceworker

// WorkerErrorReported asynchronous event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerErrorReported
type WorkerErrorReported struct {
	ErrorMessage ErrorMessage `json:"errorMessage"`
}

// WorkerRegistrationUpdated asynchronous event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerRegistrationUpdated
type WorkerRegistrationUpdated struct {
	Registrations []Registration `json:"registrations"`
}

// WorkerVersionUpdated asynchronous event.
//
// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerVersionUpdated
type WorkerVersionUpdated struct {
	Versions []Version `json:"versions"`
}
