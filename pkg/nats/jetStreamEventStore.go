package nats

import (
	"context"

	natsGo "github.com/nats-io/nats.go"

	"github.com/go-mike/event-store/pkg/abstractions"
)

type jetStreamEventStore struct {
	prefix string
	jsConn natsGo.JetStream
}

// AddEvents implements abstractions.EventStore for jetStreamEventStore
func (store *jetStreamEventStore) AddEvents(
	ctx context.Context,
	streamKey abstractions.EventStreamKey,
	transactionId string,
	events []abstractions.EventEnvelope,
	options *abstractions.AddEventsOptions,
) (abstractions.AddEventsResult, error) {
	subject := createEntityIdSubject(store.prefix, streamKey.PartitionKey, streamKey.EntityType, streamKey.EntityId)

	store.jsConn.PublishMsg()
}

// ReadInstanceEvents implements abstractions.EventStore for jetStreamEventStore
func (store *jetStreamEventStore) ReadInstanceEvents(
	ctx context.Context,
	key abstractions.ReadInstanceEventsKey,
	options abstractions.ReadInstanceEventsOptions,
) (<-chan abstractions.ReadEventsResult, error) {
	panic("unimplemented")
}

// NewJetStreamEventStore creates a new JetStream event store.
func NewJetStreamEventStore(
	jsConn natsGo.JetStream,
	prefix string,
) abstractions.EventStore {
	return &jetStreamEventStore{
		jsConn: jsConn,
		prefix: cleanSubjectStep(prefix),
	}
}
