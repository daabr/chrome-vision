package backgroundservice

// The Background Service that will be associated with the commands/events.
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

// A key-value pair for additional event information to pass along.
//
// https://chromedevtools.github.io/devtools-protocol/tot/BackgroundService/#type-EventMetadata
type EventMetadata struct {
	Key   string
	Value string
}

// https://chromedevtools.github.io/devtools-protocol/tot/BackgroundService/#type-BackgroundServiceEvent
type BackgroundServiceEvent struct {
	// Timestamp of the event (in seconds).
	Timestamp float64
	// The origin this event belongs to.
	Origin string
	// The Service Worker ID that initiated the event.
	ServiceWorkerRegistrationID string
	// The Background Service this event belongs to.
	Service string
	// A description of the event.
	EventName string
	// An identifier that groups related events together.
	InstanceID string
	// A list of event-specific information.
	EventMetadata []EventMetadata
}
