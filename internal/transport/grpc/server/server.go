package server

import (
	"context"
	"errors"
	"log"

	linkv1 "github.com/dockloud/contract/gen/link/v1"
	clientmapper "github.com/dockloud/link/internal/transport/grpc/client/mapper"
	servermapper "github.com/dockloud/link/internal/transport/grpc/server/mapper"
	"github.com/dockloud/link/internal/usecase/link"
	"github.com/dockloud/user-auth/pkg/transport/grpc/metadata"
)

var ErrNoUIDPassed = errors.New("no uid passed")

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

func (s *linkServer) GetLinks(ctx context.Context, req *linkv1.GetLinksRequest) (*linkv1.GetLinksResponse, error) {
	s.logger.Printf("Received GetLinksRequest %+v", req)

	uid, ok := metadata.GetUIDFromMetadata(ctx)
	if !ok {
		s.logger.Printf("Failed to process request %+v", req)
		return nil, ErrNoUIDPassed
	}

	input := servermapper.ToGetUserLinksInput(req, uid)
	out, err := s.service.GetUserLinks(ctx, input)
	if err != nil {
		s.logger.Printf("Failed to process request %+v", req)
		return nil, err
	}

	return servermapper.ToGetUserLinksResponse(out), nil
}

func (s *linkServer) CreateLink(ctx context.Context, req *linkv1.CreateLinkRequest) (*linkv1.CreateLinkResponse, error) {
	s.logger.Printf("Received CreateLinkRequest %+v", req)

	uid, ok := metadata.GetUIDFromMetadata(ctx)
	if !ok {
		s.logger.Printf("Failed to process request %+v", req)
		return nil, ErrNoUIDPassed
	}

	input := servermapper.ToCreateLinkInput(req, uid)
	out, err := s.service.CreateLink(ctx, input)
	if err != nil {
		s.logger.Printf("Failed to process request %+v", req)
		return nil, err
	}

	return servermapper.ToCreateLinkResponse(out), nil
}

func (s *linkServer) RevokeLink(ctx context.Context, req *linkv1.RevokeLinkRequest) (*linkv1.RevokeLinkResponse, error) {
	s.logger.Printf("Received RevokeLinkRequest %+v", req)

	uid, ok := metadata.GetUIDFromMetadata(ctx)
	if !ok {
		s.logger.Printf("Failed to process request %+v", req)
		return nil, ErrNoUIDPassed
	}

	input := servermapper.ToRevokeLinkInput(req, uid)
	out, err := s.service.RevokeLink(ctx, input)
	if err != nil {
		s.logger.Printf("Failed to process request %+v", req)
		return nil, err
	}

	return servermapper.ToRevokeLinkResponse(out), nil
}

func (s *linkServer) GetDocumentByLink(ctx context.Context, req *linkv1.GetDocumentByLinkRequest) (*linkv1.GetDocumentByLinkResponse, error) {
	s.logger.Printf("Received GetDocumentByLinkRequest %+v", req)

	input := servermapper.ToGetDocumentInput(req)
	out, err := s.service.GetDocument(ctx, input)
	if err != nil {
		s.logger.Printf("Failed to process request %+v", req)
		return nil, err
	}

	return clientmapper.ToGetDocumentByLinkResponse(out), nil
}

func (s *linkServer) UpdateDocumentByLink(ctx context.Context, req *linkv1.UpdateDocumentByLinkRequest) (*linkv1.UpdateDocumentByLinkResponse, error) {
	s.logger.Printf("Received UpdateDocumentByLinkRequest %+v", req)

	input := servermapper.ToUpdateDocumentInput(req)
	out, err := s.service.UpdateDocument(ctx, input)
	if err != nil {
		s.logger.Printf("Failed to process request %+v", req)
		return nil, err
	}

	return clientmapper.ToUpdateDocumentByLinkResponse(out), nil
}

func (s *linkServer) DeleteDocumentByLink(ctx context.Context, req *linkv1.DeleteDocumentByLinkRequest) (*linkv1.DeleteDocumentByLinkResponse, error) {
	s.logger.Printf("Received DeleteDocumentByLinkRequest %+v", req)

	input := servermapper.ToDeleteDocumentInput(req)
	out, err := s.service.DeleteDocument(ctx, input)
	if err != nil {
		s.logger.Printf("Failed to process request %+v", req)
		return nil, err
	}

	return clientmapper.ToDeleteDocumentByLinkResponse(out), nil
}
