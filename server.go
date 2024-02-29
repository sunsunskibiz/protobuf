package main

import (
	"context"

	gatewayv1 "github.com/sunsunskibiz/protobuf/gen/gateway/v1"
)

type server struct {
	gatewayv1.UnimplementedGeeterServiceServer
}

func (s *server) SayHello(ctx context.Context, request *gatewayv1.SayHelloRequest) (*gatewayv1.SayHelloResponse, error) {
	return &gatewayv1.SayHelloResponse{
		Message: "Hello " + request.GetName(),
	}, nil
}
