package grpc

import (
	"log"

	linkv1 "github.com/dockloud/contract/gen/link/v1"
	"github.com/dockloud/link/internal/usecase/link"
)

type linkServer struct {
	linkv1.UnimplementedLinkServiceServer
	service link.Service
	logger  log.Logger
}

func NewLinkServer(service link.Service) linkv1.LinkServiceServer {
	return &linkServer{
		service: service,
		logger:  *log.Default(),
	}
}
