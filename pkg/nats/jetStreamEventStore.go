package nats

import (
	"context"

	natsGo "github.com/nats-io/nats.go"

	"github.com/go-mike/event-store/pkg/abstractions"
	"github.com/go-mike/event-store/pkg/maps"
)

type jetStreamEventStore struct {
	prefix string
	jsConn natsGo.JetStream
	clock  abstractions.Clock
}

// AddEvents implements abstractions.EventStore for jetStreamEventStore
func (store *jetStreamEventStore) AddEvents(
	ctx context.Context,
	streamKey abstractions.EventStreamKey,
	transactionId string,
	previousVersion int64,
	events []abstractions.EventEnvelope,
) (abstractions.AddEventsResult, error) {
	persistedEvents := abstractions.AsPersistedEvents(
		store.clock.Now(),
		transactionId,
		previousVersion,
		events)

	subject := createEntityIdSubject(
		store.prefix,
		streamKey.PartitionKey,
		streamKey.EntityType,
		streamKey.EntityId)

	for index, event := range persistedEvents {
		msg := natsGo.NewMsg(subject)
		msg.Data = event.Data
		msg.Header = maps.MapValue(event.Metadata.Metadata, func(key string, value string) []string {
			return []string{value}
		})

		_, err := store.jsConn.PublishMsg(msg)
		if err != nil {
			return abstractions.AddEventsResult{
				PublishedEventsCount: index,
				EntityVersion:        int(previousVersion) + index,
			}, err
		}
	}

	return abstractions.AddEventsResult{
		PublishedEventsCount: len(persistedEvents),
		EntityVersion:        int(previousVersion) + len(persistedEvents),
	}, nil
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
	clock abstractions.Clock,
	prefix string,
) abstractions.EventStore {
	return &jetStreamEventStore{
		jsConn: jsConn,
		prefix: cleanSubjectStep(prefix),
	}
}
