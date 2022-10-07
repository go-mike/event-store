package abstractions

import "context"

// EventAdder is an interface for adding events to the event store.
type EventAdder interface {
	// AddEvents adds events to the event store.
	AddEvents(
		ctx context.Context,
		streamKey EventStreamKey,
		transactionId string,
		previousVersion int64,
		events []EventEnvelope,
	) (AddEventsResult, error)
}

// AddEventsResult contains the result of adding events to the event store.
type AddEventsResult struct {
	// EntityVersion is the version of the entity after the events were added.
	EntityVersion int
	// PublishedEventsCount is the number of events that were published.
	PublishedEventsCount int
}
