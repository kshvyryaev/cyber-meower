package command

import (
	"time"

	contractEvent "github.com/kshvyryaev/cyber-meower/internal/contract/event"
	"github.com/kshvyryaev/cyber-meower/internal/meow-service/domain"
	"github.com/kshvyryaev/cyber-meower/internal/meow-service/event"
	"github.com/kshvyryaev/cyber-meower/internal/meow-service/repository"
	"github.com/kshvyryaev/cyber-meower/internal/meow-service/service"
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

	event := &contractEvent.MeowCreatedEvent{
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
