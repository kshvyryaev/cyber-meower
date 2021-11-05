package controller

import (
	"net"

	"github.com/kshvyryaev/cyber-meower-meower-service/pkg"
	"github.com/kshvyryaev/cyber-meower-proto/pkg/proto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	config         *pkg.Config
	meowController *GrpcMeowController
}

func ProvideGrpcServer(config *pkg.Config, meowController *GrpcMeowController) *GrpcServer {
	return &GrpcServer{
		config:         config,
		meowController: meowController,
	}
}

func (server *GrpcServer) Run() error {
	listener, err := net.Listen("tcp", ":"+server.config.Port)
	if err != nil {
		return errors.Wrap(err, "grpc server")
	}

	// TODO: Add error and panic interceptors
	grpcServer := grpc.NewServer()
	proto.RegisterMeowServiceServer(grpcServer, server.meowController)

	grpcServer.Serve(listener)
	return nil
}
