package client

import (
	pb "github.com/dockloud/contract/gen/document/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewDocumentClient(addr string) (*pb.DocumentServiceClient, error) {
	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	client := pb.NewDocumentServiceClient(conn)

	return &client, nil
}
