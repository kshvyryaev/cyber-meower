package domain

import (
	"time"

	eventContract "github.com/kshvyryaev/cyber-meower-event/pkg/event"
)

type MeowUsecase interface {
	Create(body string) (int, error)
}

type MeowTranslatorService interface {
	Translate(body string) string
}

type MeowRepository interface {
	Create(meow *Meow) (int, error)
}

type MeowEventPublisher interface {
	Publish(event *eventContract.MeowCreatedEvent)
}

type Meow struct {
	ID        int
	Body      string
	CreatedOn time.Time
}
