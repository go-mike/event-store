package abstractions

// EventStore is an interface for an event store, with all features implemented.
type EventStore interface {
	EventAdder

	InstanceEventReader
}