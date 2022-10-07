package abstractions

import "time"

// EventStreamKey is the key for an event stream.
type EventStreamKey struct {
	// PartitionKey is the partition key for the event stream.
	PartitionKey string
	// EntityType is the type of the entity.
	EntityType string
	// EntityId is the ID of the entity.
	EntityId string
}

// EventMetadata is the metadata for an event.
type EventMetadata struct {
	// EventType is the type of the event.
	EventType string
	// ContentType is the content type of the event.
	ContentType string

	// Metadata is the metadata for the event.
	Metadata map[string]string
}

// PersistedEventMetadata is the persisted metadata for an event.
type PersistedEventMetadata struct {
	// EntityVersion is the version of the entity after the event was added.
	EntityVersion int64
	// CreatedAt is the time the event was created.
	CreatedAt time.Time
	// TransactionId is the ID of the transaction that added the event.
	TransactionId string

	// EventMetadata is the metadata for the event.
	EventMetadata
}

// EventEnvelope is an envelope for an event data and metadata.
type EventEnvelope struct {
	// Data is the event data.
	Data []byte
	// Metadata is the metadata for the event.
	Metadata EventMetadata
}

// PersistedEventEnvelope is an envelope for an event data and persisted metadata.
type PersistedEventEnvelope struct {
	// Data is the event data.
	Data []byte
	// Metadata is the persisted metadata for the event.
	Metadata PersistedEventMetadata
}

func AsPersistedEvents(
	now time.Time,
	transactionId string,
	previousVersion int64,
	events []EventEnvelope,
) []PersistedEventEnvelope {
	persistedEvents := make([]PersistedEventEnvelope, len(events))

	for i, event := range events {
		persistedEvents[i] = PersistedEventEnvelope {
			Data: event.Data,
			Metadata: PersistedEventMetadata {
				EntityVersion: previousVersion + int64(i) + 1,
				CreatedAt: now,
				TransactionId: transactionId,
				EventMetadata: event.Metadata,
			},
		}
	}

	return persistedEvents
}
