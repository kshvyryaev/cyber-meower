package event

import (
	"bytes"
	"encoding/gob"

	"github.com/google/wire"
	"github.com/kshvyryaev/cyber-meower/internal/meow-service/config"
	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
)

type NatsEventPublisher struct {
	connection *nats.Conn
}

func ProvideNatsEventPublisher(config *config.Config) (*NatsEventPublisher, func(), error) {
	connection, err := nats.Connect(config.EventStoreAddress)
	if err != nil {
		return nil, nil, errors.Wrap(err, "nats event publisher")
	}

	cleanup := func() {
		connection.Close()
	}

	return &NatsEventPublisher{connection: connection}, cleanup, nil
}

func (publisher *NatsEventPublisher) Publish(event Event) error {
	bytes, err := publisher.parseEventToBytes(event)
	if err != nil {
		return errors.Wrap(err, "nats event publisher")
	}

	err = publisher.connection.Publish(event.GetKey(), bytes)
	if err != nil {
		return errors.Wrap(err, "nats event publisher")
	}

	return nil
}

func (publisher *NatsEventPublisher) parseEventToBytes(event Event) ([]byte, error) {
	buffer := bytes.Buffer{}

	err := gob.NewEncoder(&buffer).Encode(event)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

var NatsEventPublisherSet = wire.NewSet(
	ProvideNatsEventPublisher,
	wire.Bind(new(EventPublisher), new(*NatsEventPublisher)),
)
