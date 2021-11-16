package test

import (
	"testing"

	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/usecase"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Create(t *testing.T) {
	databaseErr := errors.New("database error")

	type dependencies struct {
		translator     *MockMeowTranslatorService
		repository     *MockMeowRepository
		eventPublisher *MockMeowEventPublisher
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
				translator:     getTranslatorMock(mock.Anything, ""),
				repository:     getRepositoryMock(mock.Anything, 1, nil),
				eventPublisher: getEventPublisherMock(mock.Anything),
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
				translator:     getTranslatorMock(mock.Anything, ""),
				repository:     getRepositoryMock(mock.Anything, 0, databaseErr),
				eventPublisher: getEventPublisherMock(mock.Anything),
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
				tt.dependencies.translator.AssertExpectations(t)
				tt.dependencies.repository.AssertExpectations(t)
			} else {
				assert.Nil(t, gotErr)
				tt.dependencies.translator.AssertExpectations(t)
				tt.dependencies.repository.AssertExpectations(t)
				tt.dependencies.eventPublisher.AssertExpectations(t)
			}
		})
	}
}

func getTranslatorMock(input, output interface{}) *MockMeowTranslatorService {
	translator := &MockMeowTranslatorService{}
	translator.On("Translate", input).Return(output)
	return translator
}

func getRepositoryMock(input, output, outputErr interface{}) *MockMeowRepository {
	repository := &MockMeowRepository{}
	repository.On("Create", input).Return(output, outputErr)
	return repository
}

func getEventPublisherMock(input interface{}) *MockMeowEventPublisher {
	eventPublisher := &MockMeowEventPublisher{}
	eventPublisher.On("Publish", input)
	return eventPublisher
}
