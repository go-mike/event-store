package abstractions

import "context"

// EventAdder is an interface for adding events to the event store.
type EventAdder interface {
	AddEvents(
		ctx context.Context,
		streamKey EventStreamKey,
		transactionId string,
		events []EventEnvelope,
		options *AddEventsOptions,
	) (AddEventsResult, error)
}

type AddEventsOptions struct {
	ExpectedVersion *int
}

type AddEventsResult struct {
	EntityVersion int
}
