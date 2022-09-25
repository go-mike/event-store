package abstractions

import (
	"context"
	"time"
)

// InstanceEventReader is an interface for reading events from the event store.
type InstanceEventReader interface {
	// ReadInstanceEvents reads events from the event store.
	ReadInstanceEvents(
		ctx context.Context,
		key ReadInstanceEventsKey,
		options ReadInstanceEventsOptions,
	) (<-chan ReadEventsResult, error)
}

// ReadInstanceEventsKey contains the key for reading events from the event store.
type ReadInstanceEventsKey struct {
	// PartitionKey is the partition key.
	PartitionKey string
	// EntityType is the entity type.
	EntityType string
	// EntityId is the entity id.
	EntityId string
}

// ReadInstanceEventsOptions contains the options for reading events from the event store.
type ReadInstanceEventsOptions struct {
	// FromVersion is the entity version to start reading from.
	FromVersion *int
	// FromTime is the event creation time to start reading from.
	FromTime *time.Time
	// FromTransactionId is the transaction id to read from.
	FromTransactionId *string
	// PageSize is the number of events to read on each page.
	PageSize *int
	// ChannelBufferSize is the size of the channel buffer.
	ChannelBufferSize *int
}

// ReadEventsResult contains the result of reading events from the event store.
type ReadEventsResult struct {
	// Events are the events.
	Events     []PersistedEventEnvelope
	// IsLastPage is true if this is the last page.
	IsLastPage bool
}
