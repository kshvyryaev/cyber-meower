package test

import (
	eventContract "github.com/kshvyryaev/cyber-meower-event/pkg/event"
	"github.com/stretchr/testify/mock"
)

type MockMeowEventPublisher struct {
	mock.Mock
}

func (publisher *MockMeowEventPublisher) Publish(event *eventContract.MeowCreatedEvent) {
	publisher.Called(event)
}
