package abstractions

import (
	"context"
	"time"
)

type InstanceEventReader interface {
	ReadInstanceEvents(
		ctx context.Context,
		key ReadInstanceEventsKey,
		options ReadInstanceEventsOptions,
	) (<-chan ReadInstanceEventsResult, error)
}

type ReadInstanceEventsKey struct {
	PartitionKey string
	EntityType   string
	EntityId     string
}

type ReadInstanceEventsOptions struct {
	FromVersion       *int
	FromTime          *time.Time
	FromTransactionId *string
}

type ReadInstanceEventsResult struct {
	Events []PersistedEventEnvelope
}
