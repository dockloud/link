package document

import (
	"context"

	pb "github.com/dockloud/contract/gen/document/v1"
)

type Repository struct {
	client pb.DocumentServiceClient
}

func NewRepository(client pb.DocumentServiceClient) *Repository {
	return &Repository{
		client: client,
	}
}

func (r *Repository) Get(ctx context.Context, documentID int) (*pb.GetUnsafeDocumentResponse, error) {
	resp, err := r.client.GetUnsafeDocument(
		ctx,
		&pb.GetUnsafeDocumentRequest{
			DocumentId: int64(documentID),
		})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *Repository) Update(
	ctx context.Context, documentID int,
	content, name *string, isPrivate *bool,
) (*pb.UpdateUnsafeDocumentResponse, error) {
	resp, err := r.client.UpdateUnsafeDocument(
		ctx,
		&pb.UpdateUnsafeDocumentRequest{
			DocumentId: int64(documentID),
			Content:    content,
			Name:       name,
			IsPrivate:  isPrivate,
		})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *Repository) Delete(ctx context.Context, documentID int) (*pb.DeleteUnsafeDocumentResponse, error) {
	resp, err := r.client.DeleteUnsafeDocument(
		ctx,
		&pb.DeleteUnsafeDocumentRequest{
			DocumentId: int64(documentID),
		})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
