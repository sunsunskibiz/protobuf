package server

import (
	"context"
	"encoding/base64"
	"fmt"

	mediav1 "github.com/sunsunskibiz/protobuf/gen/media/v1"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type MediaServer struct {
	mediav1.UnimplementedMediaServiceServer
}

func (s *MediaServer) Image(ctx context.Context, request *mediav1.ImageRequest) (*httpbody.HttpBody, error) {
	filename := "Kirby.png"
	err := grpc.SendHeader(ctx, metadata.Pairs("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename)))
	if err != nil {
		return nil, err
	}

	img, err := base64.StdEncoding.DecodeString(string(request.File))
	if err != nil {
		return nil, err
	}

	return &httpbody.HttpBody{
		ContentType: "image/png",
		Data:        img,
	}, nil
}
