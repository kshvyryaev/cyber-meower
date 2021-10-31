package command

import (
	"time"

	"github.com/google/wire"
	"github.com/kshvyryaev/cyber-meower/internal/meow-service/domain"
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

	return &CreateMeowCommandResponse{ID: id}, nil
}

func ProvideСreateMeowCommandHandler(
	meowTranslator *service.MeowTranslatorService,
	repository repository.MeowRepository) *СreateMeowCommandHandler {
	return &СreateMeowCommandHandler{
		meowTranslator: meowTranslator,
		repository:     repository,
	}
}

var СreateMeowCommandHandlerSet = wire.NewSet(
	ProvideСreateMeowCommandHandler,
)
