package contract

import eventContract "github.com/kshvyryaev/cyber-meower-event/pkg/event"

type MeowEventPublisher interface {
	Publish(event *eventContract.MeowCreatedEvent)
}
