package serviceworker

// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerErrorReported
type WorkerErrorReported struct {
	ErrorMessage ServiceWorkerErrorMessage
}

// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerRegistrationUpdated
type WorkerRegistrationUpdated struct {
	Registrations []ServiceWorkerRegistration
}

// https://chromedevtools.github.io/devtools-protocol/tot/ServiceWorker/#event-workerVersionUpdated
type WorkerVersionUpdated struct {
	Versions []ServiceWorkerVersion
}
