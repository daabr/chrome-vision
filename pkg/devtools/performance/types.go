package performance

// Metric data type. Run-time execution metric.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Performance/#type-Metric
type Metric struct {
	// Metric name.
	Name string `json:"name"`
	// Metric value.
	Value float64 `json:"value"`
}
