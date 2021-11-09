package event

import (
	eventContract "github.com/kshvyryaev/cyber-meower-event/pkg/event"
	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
)

type NatsMeowEventPublisher struct {
	connection *nats.EncodedConn
	channel    chan<- *eventContract.MeowCreatedEvent
}

func ProvideNatsMeowEventPublisher(connection *nats.EncodedConn) (*NatsMeowEventPublisher, func(), error) {
	channel := make(chan *eventContract.MeowCreatedEvent)

	err := connection.BindSendChan(eventContract.MeowCreatedEventSubject, channel)
	if err != nil {
		return nil, nil, errors.Wrap(err, "nats meow event publisher")
	}

	cleanup := func() {
		close(channel)
	}

	return &NatsMeowEventPublisher{connection: connection, channel: channel}, cleanup, nil
}

func (publisher *NatsMeowEventPublisher) Publish(event *eventContract.MeowCreatedEvent) {
	publisher.channel <- event
}
