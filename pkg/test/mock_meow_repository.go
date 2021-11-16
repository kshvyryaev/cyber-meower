package test

import (
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/domain"
	"github.com/stretchr/testify/mock"
)

type MockMeowRepository struct {
	mock.Mock
}

func (repository *MockMeowRepository) Create(meow *domain.Meow) (int, error) {
	args := repository.Called(meow)
	return args.Int(0), args.Error(1)
}
