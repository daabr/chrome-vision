package performance

// Metrics asynchronous event.
//
// Current values of the metrics.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Performance/#event-metrics
type Metrics struct {
	// Current values of the metrics.
	Metrics []Metric `json:"metrics"`
	// Timestamp title.
	Title string `json:"title"`
}
