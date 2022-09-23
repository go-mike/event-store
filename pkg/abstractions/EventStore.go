package abstractions

type EventStore interface {
	EventAdder

	InstanceEventReader
}