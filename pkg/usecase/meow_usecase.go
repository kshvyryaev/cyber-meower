package usecase

import (
	"time"

	eventContract "github.com/kshvyryaev/cyber-meower-event/pkg/event"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/domain"
	"github.com/pkg/errors"
)

type MeowUsecase struct {
	translator     domain.MeowTranslatorService
	repository     domain.MeowRepository
	eventPublisher domain.MeowEventPublisher
}

func ProvideMeowUsecase(
	translator domain.MeowTranslatorService,
	repository domain.MeowRepository,
	eventPublisher domain.MeowEventPublisher) *MeowUsecase {
	return &MeowUsecase{
		translator:     translator,
		repository:     repository,
		eventPublisher: eventPublisher,
	}
}

func (usecase *MeowUsecase) Create(body string) (int, error) {
	translatedBody := usecase.translator.Translate(body)
	meow := &domain.Meow{
		Body:      translatedBody,
		CreatedOn: time.Now().UTC(),
	}

	id, err := usecase.repository.Create(meow)
	if err != nil {
		return 0, errors.Wrap(err, "create meow command handler")
	}

	event := &eventContract.MeowCreatedEvent{
		ID:        id,
		Body:      meow.Body,
		CreatedOn: meow.CreatedOn,
	}

	usecase.eventPublisher.Publish(event)

	return id, nil
}
