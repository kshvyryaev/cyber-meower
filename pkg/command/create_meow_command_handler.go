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

type CreateMeowCommand struct {
	Body string `json:"body"`
}

type CreateMeowCommandResponse struct {
	ID int `json:"id"`
}

type СreateMeowCommandHandler struct {
	meowTranslator *service.MeowTranslatorService
	repository     repository.MeowRepository
	eventPublisher event.EventPublisher
}

func ProvideСreateMeowCommandHandler(
	meowTranslator *service.MeowTranslatorService,
	repository repository.MeowRepository,
	eventPublisher event.EventPublisher) *СreateMeowCommandHandler {
	return &СreateMeowCommandHandler{
		meowTranslator: meowTranslator,
		repository:     repository,
		eventPublisher: eventPublisher,
	}
}

func (handler *СreateMeowCommandHandler) Handle(command *CreateMeowCommand) (*CreateMeowCommandResponse, error) {
	meowBody := handler.meowTranslator.Translate(command.Body)
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
	if err != nil {
		return nil, errors.Wrap(err, "create meow command handler")
	}

	return &CreateMeowCommandResponse{ID: id}, nil
}
