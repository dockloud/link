package client

import (
	pb "github.com/dockloud/contract/gen/document/v1"
	"github.com/dockloud/link/internal/transport/grpc/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewDocumentClient(addr string, secret []byte) (*pb.DocumentServiceClient, error) {
	interceptor, err := auth.NewSignInterceptor(secret)
	if err != nil {
		return nil, err
	}

	conn, err := grpc.NewClient(
		addr,
		grpc.WithUnaryInterceptor(interceptor),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	client := pb.NewDocumentServiceClient(conn)

	return &client, nil
}
