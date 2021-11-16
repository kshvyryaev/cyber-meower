package test

import (
	"testing"

	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/contract"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/usecase"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Create(t *testing.T) {
	databaseErr := errors.New("database error")

	type dependencies struct {
		translator     contract.MeowTranslatorService
		repository     contract.MeowRepository
		eventPublisher contract.MeowEventPublisher
	}

	type args struct {
		body string
	}

	tests := []struct {
		name         string
		dependencies dependencies
		args         args
		want         int
		wantErr      error
	}{
		{
			name: "Returns id when everything is ok",
			dependencies: dependencies{
				translator:     getTranslatorMock(""),
				repository:     getRepositoryMock(1, nil),
				eventPublisher: getEventPublisherMock(),
			},
			args: args{
				body: "A",
			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "Returns wrapped error when repository returns error",
			dependencies: dependencies{
				translator:     getTranslatorMock(""),
				repository:     getRepositoryMock(0, databaseErr),
				eventPublisher: getEventPublisherMock(),
			},
			args: args{
				body: "A",
			},
			want:    0,
			wantErr: errors.Wrap(databaseErr, "create meow command handler"),
		},
	}

	for _, tt := range tests {
		usecase := usecase.ProvideMeowUsecase(
			tt.dependencies.translator,
			tt.dependencies.repository,
			tt.dependencies.eventPublisher)

		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := usecase.Create(tt.args.body)
			assert.Equal(t, tt.want, got)

			if tt.wantErr != nil {
				assert.Error(t, tt.wantErr, gotErr)
				assert.EqualError(t, tt.wantErr, gotErr.Error())
			} else {
				assert.Nil(t, gotErr)
			}
		})
	}
}

func getTranslatorMock(output string) *MockMeowTranslatorService {
	translator := new(MockMeowTranslatorService)
	translator.On("Translate", mock.Anything).Return(output)
	return translator
}

func getRepositoryMock(output int, outputErr error) *MockMeowRepository {
	repository := new(MockMeowRepository)
	repository.On("Create", mock.Anything).Return(output, outputErr)
	return repository
}

func getEventPublisherMock() *MockMeowEventPublisher {
	eventPublisher := new(MockMeowEventPublisher)
	eventPublisher.On("Create", mock.Anything)
	return eventPublisher
}
