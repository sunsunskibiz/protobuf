package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getFormFile(r *http.Request, name string) ([]byte, error) {
	file, _, err := r.FormFile(name)
	if err != nil {
		return nil, fmt.Errorf("not found")
	}
	defer file.Close()

	buf := bytes.Buffer{}
	io.Copy(&buf, file)
	return buf.Bytes(), nil
}

func createRequestFromMultiPart(r *http.Request) (*http.Request, error) {
	dataBytes, err := getFormFile(r, "file")
	if err != nil {
		fmt.Printf("%v", err)
		return nil, err
	}

	dataBase64 := base64.StdEncoding.EncodeToString(dataBytes)
	req := map[string]any{
		"file": []byte(dataBase64),
	}
	str, err := json.Marshal(&req)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(str)
	newR, err := http.NewRequest(http.MethodPost, r.URL.String(), reader)
	if err != nil {
		return nil, err
	}
	return newR, nil
}

func customizeMiddleware(otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data") {
			newR, err := createRequestFromMultiPart(r)
			if err != nil {
				w.WriteHeader(400)
				return
			}
			otherHandler.ServeHTTP(w, newR)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}

func doGateway(registerGWHandler func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error) {
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8090",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()

	// Register Greeter
	err = registerGWHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := http.Server{
		Addr:    ":8080",
		Handler: customizeMiddleware(gwmux),
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8080")
	log.Fatalln(gwServer.ListenAndServe())
}
