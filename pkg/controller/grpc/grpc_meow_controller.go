package grpc

import (
	"context"

	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/domain"
	"github.com/kshvyryaev/cyber-meower-proto/pkg/proto"
	"github.com/pkg/errors"
)

type GrpcMeowController struct {
	proto.UnimplementedMeowServiceServer
	usecase domain.MeowUsecase
}

func ProvideGrpcMeowController(usecase domain.MeowUsecase) *GrpcMeowController {
	return &GrpcMeowController{
		usecase: usecase,
	}
}

func (controller *GrpcMeowController) Create(context context.Context, request *proto.CreateMeowRequest) (*proto.CreateMeowResponse, error) {
	id, err := controller.usecase.Create(request.Body)
	if err != nil {
		return nil, errors.Wrap(err, "grpc meow controller")
	}

	return &proto.CreateMeowResponse{ID: int64(id)}, nil
}
