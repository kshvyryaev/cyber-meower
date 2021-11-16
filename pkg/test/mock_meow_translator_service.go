package test

import "github.com/stretchr/testify/mock"

type MockMeowTranslatorService struct {
	mock.Mock
}

func (service *MockMeowTranslatorService) Translate(body string) string {
	args := service.Called(body)
	return args.String(0)
}
