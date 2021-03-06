package backgroundservice

// ServiceName data type. The Background Service that will be associated with the commands/events.
// Every Background Service operates independently, but they share the same
// API.
//
// https://chromedevtools.github.io/devtools-protocol/tot/BackgroundService/#type-ServiceName
type ServiceName string

// ServiceName valid values.
const (
	ServiceNameBackgroundFetch        ServiceName = "backgroundFetch"
	ServiceNameBackgroundSync         ServiceName = "backgroundSync"
	ServiceNamePushMessaging          ServiceName = "pushMessaging"
	ServiceNameNotifications          ServiceName = "notifications"
	ServiceNamePaymentHandler         ServiceName = "paymentHandler"
	ServiceNamePeriodicBackgroundSync ServiceName = "periodicBackgroundSync"
)

// String returns the ServiceName value as a built-in string.
func (t ServiceName) String() string {
	return string(t)
}

// EventMetadata data type. A key-value pair for additional event information to pass along.
//
// https://chromedevtools.github.io/devtools-protocol/tot/BackgroundService/#type-EventMetadata
type EventMetadata struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Event data type.
//
// https://chromedevtools.github.io/devtools-protocol/tot/BackgroundService/#type-BackgroundServiceEvent
type Event struct {
	// Timestamp of the event (in seconds).
	Timestamp float64 `json:"timestamp"`
	// The origin this event belongs to.
	Origin string `json:"origin"`
	// The Service Worker ID that initiated the event.
	ServiceWorkerRegistrationID string `json:"serviceWorkerRegistrationId"`
	// The Background Service this event belongs to.
	Service ServiceName `json:"service"`
	// A description of the event.
	EventName string `json:"eventName"`
	// An identifier that groups related events together.
	InstanceID string `json:"instanceId"`
	// A list of event-specific information.
	EventMetadata []EventMetadata `json:"eventMetadata"`
}
