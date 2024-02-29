package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/bufbuild/protovalidate-go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	gatewayv1 "github.com/sunsunskibiz/protobuf/gen/gateway/v1"
	grpcv1 "github.com/sunsunskibiz/protobuf/gen/grpc/v1"
	mapv1 "github.com/sunsunskibiz/protobuf/gen/map/v1"
	simplev1 "github.com/sunsunskibiz/protobuf/gen/simple/v1"
	validatesimplev1 "github.com/sunsunskibiz/protobuf/gen/validatesimple/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

func doGateway() {
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()

	// Register Greeter
	err = gatewayv1.RegisterGeeterServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}

func doGRPC(registerServer func(s *grpc.Server)) {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
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
	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatal(s.Serve(lis))
	}()
}

func main() {
	fmt.Println(doSimple())
	fmt.Println(doMap())

	doValidate()

	doGRPC(func(s *grpc.Server) {
		gatewayv1.RegisterGeeterServiceServer(s, &server{})
		grpcv1.RegisterEchoServiceServer(s, &grpcServer{})
	})

	doGateway()
}
