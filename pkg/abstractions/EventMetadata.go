package abstractions

import "time"

type EventStreamKey struct {
	PartitionKey string
	EntityType   string
	EntityId     string
}

type EventMetadata struct {
	EventType   string
	ContentType string

	Metadata map[string]string
}

type PersistedEventMetadata struct {
	EntityVersion int
	CreatedAt     time.Time
	TransactionId string

	EventMetadata
}

type EventEnvelope struct {
	Data     []byte
	Metadata EventMetadata
}

type PersistedEventEnvelope struct {
	Data     []byte
	Metadata PersistedEventMetadata
}
