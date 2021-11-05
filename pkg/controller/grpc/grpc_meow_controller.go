package controller

import (
	"context"

	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/command"
	"github.com/kshvyryaev/cyber-meower-proto/pkg/proto"
	"github.com/pkg/errors"
)

type GrpcMeowController struct {
	proto.UnimplementedMeowServiceServer
	createMeowCommandHandler *command.СreateMeowCommandHandler
}

func ProvideGrpcMeowController(createMeowCommandHandler *command.СreateMeowCommandHandler) *GrpcMeowController {
	return &GrpcMeowController{
		createMeowCommandHandler: createMeowCommandHandler,
	}
}

func (controller *GrpcMeowController) Create(context context.Context, request *proto.CreateMeowRequest) (*proto.CreateMeowResponse, error) {
	command := &command.CreateMeowCommand{Body: request.Body}

	response, err := controller.createMeowCommandHandler.Handle(command)
	if err != nil {
		return nil, errors.Wrap(err, "grpc meow controller")
	}

	return &proto.CreateMeowResponse{ID: int64(response.ID)}, nil
}
