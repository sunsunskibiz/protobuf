package server

import (
	"context"

	grpcv1 "github.com/sunsunskibiz/protobuf/gen/grpc/v1"
)

type GRPCServer struct {
	grpcv1.UnimplementedEchoServiceServer
}

func (s *GRPCServer)  Echo(ctx context.Context, request *grpcv1.EchoRequest) (*grpcv1.EchoResponse, error) {
	return &grpcv1.EchoResponse{
		Message: "Hello " + request.GetName(),
	}, nil
}