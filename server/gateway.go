package server

import (
	"context"

	gatewayv1 "github.com/sunsunskibiz/protobuf/gen/gateway/v1"
)

type Server struct {
	gatewayv1.UnimplementedGeeterServiceServer
}

func (s *Server) SayHello(ctx context.Context, request *gatewayv1.SayHelloRequest) (*gatewayv1.SayHelloResponse, error) {
	return &gatewayv1.SayHelloResponse{
		Message: "Hello " + request.GetName(),
	}, nil
}
