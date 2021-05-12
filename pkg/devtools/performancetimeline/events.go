package performancetimeline

// TimelineEventAdded asynchronous event. Sent when a performance timeline event is added. See reportPerformanceTimeline method.
//
// https://chromedevtools.github.io/devtools-protocol/tot/PerformanceTimeline/#event-timelineEventAdded
type TimelineEventAdded struct {
	Event TimelineEvent `json:"event"`
}
