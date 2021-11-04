package command

import (
	"time"

	eventContract "github.com/kshvyryaev/cyber-meower-event/pkg/event"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/domain"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/event"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/repository"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/service"
	"github.com/pkg/errors"
)

type СreateMeowCommandHandler struct {
	translator     *service.MeowTranslatorService
	repository     repository.MeowRepository
	eventPublisher event.MeowEventPublisher
}

func ProvideСreateMeowCommandHandler(
	translator *service.MeowTranslatorService,
	repository repository.MeowRepository,
	eventPublisher event.MeowEventPublisher) *СreateMeowCommandHandler {
	return &СreateMeowCommandHandler{
		translator:     translator,
		repository:     repository,
		eventPublisher: eventPublisher,
	}
}

func (handler *СreateMeowCommandHandler) Handle(command *CreateMeowCommand) (*CreateMeowResponse, error) {
	meowBody := handler.translator.Translate(command.Body)
	meow := &domain.Meow{
		Body:      meowBody,
		CreatedOn: time.Now().UTC(),
	}

	id, err := handler.repository.Create(meow)
	if err != nil {
		return nil, errors.Wrap(err, "create meow command handler")
	}

	event := &eventContract.MeowCreatedEvent{
		ID:        id,
		Body:      meow.Body,
		CreatedOn: meow.CreatedOn,
	}

	handler.eventPublisher.Publish(event)

	return &CreateMeowResponse{ID: id}, nil
}
