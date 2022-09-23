package inmemory

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-mike/event-store/pkg/abstractions"
)

type inMemoryEventStore struct {
	mutex      sync.Mutex
	clock      abstractions.Clock
	partitions map[string]*inMemoryPartition
}

type inMemoryPartition struct {
	mutex    sync.Mutex
	events   []abstractions.PersistedEventEnvelope
	versions map[string]int
}

// New creates a new in-memory event store.
func New(
	clock abstractions.Clock,
) (abstractions.EventStore, error) {
	return &inMemoryEventStore{
		clock: clock,
	}, nil
}

// AddEvents implements abstractions.EventStore
func (store *inMemoryEventStore) AddEvents(
	ctx context.Context,
	streamKey abstractions.EventStreamKey,
	transactionId string,
	events []abstractions.EventEnvelope,
	options *abstractions.AddEventsOptions,
) (abstractions.AddEventsResult, error) {
	partition := store.getOrCreatePartition(streamKey.PartitionKey)

	partition.mutex.Lock()
	defer partition.mutex.Unlock()

	instanceKey := fmt.Sprintf("%s/%s", streamKey.EntityType, streamKey.EntityId)
	nextVersion, ok := partition.versions[instanceKey]
	if !ok {
		nextVersion = 1
	}

	if options != nil && options.ExpectedVersion != nil {
		if *options.ExpectedVersion != nextVersion-1 {
			return abstractions.AddEventsResult{},
				fmt.Errorf(
					"expected version %d, but was %d",
					*options.ExpectedVersion,
					nextVersion-1)
		}
	}
}

// ReadInstanceEvents implements abstractions.EventStore
func (store *inMemoryEventStore) ReadInstanceEvents(
	ctx context.Context,
	key abstractions.ReadInstanceEventsKey,
	options abstractions.ReadInstanceEventsOptions,
) (<-chan abstractions.ReadInstanceEventsResult, error) {
}

// getOrCreatePartition returns the partition for the given key, creating it if it doesn't exist.
func (store *inMemoryEventStore) getOrCreatePartition(
	partitionKey string,
) *inMemoryPartition {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	partition, ok := store.partitions[partitionKey]
	if !ok {
		partition = &inMemoryPartition{}
		store.partitions[partitionKey] = partition
	}

	return partition
}

var _ abstractions.EventStore = (*inMemoryEventStore)(nil)
