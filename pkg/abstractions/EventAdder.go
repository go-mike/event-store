package abstractions

import "context"

// EventAdder is an interface for adding events to the event store.
type EventAdder interface {
	// AddEvents adds events to the event store.
	AddEvents(
		ctx context.Context,
		streamKey EventStreamKey,
		transactionId string,
		events []EventEnvelope,
		options *AddEventsOptions,
	) (AddEventsResult, error)
}

// AddEventsOptions contains the options for adding events to the event store.
type AddEventsOptions struct {
	// ExpectedVersion is the expected version of the stream.
	ExpectedVersion *int
}

// AddEventsResult contains the result of adding events to the event store.
type AddEventsResult struct {
	// EntityVersion is the version of the entity after the events were added.
	EntityVersion int
}
