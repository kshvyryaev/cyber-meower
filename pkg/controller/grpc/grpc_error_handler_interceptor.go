package controller

import (
	"context"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GrpcErrorHandlerInterceptor struct {
	logger *zap.Logger
}

func ProvideGrpcErrorHandlerInterceptor(logger *zap.Logger) *GrpcErrorHandlerInterceptor {
	return &GrpcErrorHandlerInterceptor{
		logger: logger,
	}
}

func (interceptor *GrpcErrorHandlerInterceptor) Handle(context context.Context, request interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	response, err := handler(context, request)
	if err != nil {
		interceptor.logger.Error("error happend", zap.Error(err))
		return nil, errors.Wrap(err, "grpc error handler interceptor")
	}

	return response, err
}
