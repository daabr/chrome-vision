package performance

// Current values of the metrics.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Performance/#event-metrics
type Metrics struct {
	// Current values of the metrics.
	Metrics []Metric
	// Timestamp title.
	Title string
}
