package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/bufbuild/protovalidate-go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	gatewayv1 "github.com/sunsunskibiz/protobuf/gen/gateway/v1"
	grpcv1 "github.com/sunsunskibiz/protobuf/gen/grpc/v1"
	mapv1 "github.com/sunsunskibiz/protobuf/gen/map/v1"
	mediav1 "github.com/sunsunskibiz/protobuf/gen/media/v1"
	simplev1 "github.com/sunsunskibiz/protobuf/gen/simple/v1"
	validatesimplev1 "github.com/sunsunskibiz/protobuf/gen/validatesimple/v1"
	"github.com/sunsunskibiz/protobuf/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func doSimple() *simplev1.Simple {
	return &simplev1.Simple{
		Id:         1,
		IsSimple:   true,
		Name:       "Sun",
		SimpleList: []int32{1, 2, 3, 4},
	}
}

func doMap() *mapv1.MapExample {
	return &mapv1.MapExample{
		Ids: map[string]*mapv1.IdWrapper{
			"1": {Id: 1},
			"2": {Id: 2},
			"3": {Id: 3},
		},
	}
}

func doValidate() {
	v, err := protovalidate.New()
	if err != nil {
		fmt.Println("failed to initialize validator:", err)
	}

	msg := validatesimplev1.Simple{
		Id: 123, // Modify here to get validate failed
	}

	if err = v.Validate(&msg); err != nil {
		fmt.Println("validation failed:", err)
	} else {
		fmt.Println("validation succeeded")
	}
}

func doGRPC(registerServer func(s *grpc.Server)) {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()

	// Attach the Greeter service to the server
	registerServer(s)

	// Register reflaction
	reflection.Register(s)

	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:8090")
	go func() {
		log.Fatal(s.Serve(lis))
	}()
}

func main() {
	fmt.Println(doSimple())
	fmt.Println(doMap())

	doValidate()

	doGRPC(func(s *grpc.Server) {
		grpc.ServiceRegistrar(s).RegisterService(&gatewayv1.GeeterService_ServiceDesc, gatewayv1.GeeterServiceServer(&server.Server{}))
		grpcv1.RegisterEchoServiceServer(s, &server.GRPCServer{})
	})

	doGateway(func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
		err := gatewayv1.RegisterGeeterServiceHandler(context.Background(), mux, conn)
		if err != nil {
			return err
		}

		return nil
	})
}
