package link

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"time"

	docv1 "github.com/dockloud/contract/gen/document/v1"
	"github.com/dockloud/link/internal/domain/link"
)

type LinkRepository interface {
	GetByToken(ctx context.Context, token string) (*link.Link, error)
	GetAllByUser(ctx context.Context, uid string) ([]link.Link, error)
	Create(ctx context.Context, token, creatorID string, documentID int, readAllowed, writeAllowed, deleteAllowed bool, expiresAt *string) (*link.Link, error)
	Revoke(ctx context.Context, token string) error
}
type DocumentRepository interface {
	Get(ctx context.Context, documentID int) (*docv1.GetUnsafeDocumentResponse, error)
	Update(ctx context.Context, documentID int, content, name *string, isPrivate *bool) (*docv1.UpdateUnsafeDocumentResponse, error)
	Delete(ctx context.Context, documentID int) (*docv1.DeleteUnsafeDocumentResponse, error)
}

type Service struct {
	linkRepo     LinkRepository
	documentRepo DocumentRepository
}

func NewService(linkRepo LinkRepository, documentRepo DocumentRepository) *Service {
	return &Service{
		linkRepo:     linkRepo,
		documentRepo: documentRepo,
	}
}

func generateToken(id uint64) string {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, id)
	digest := sha256.Sum256(buf)
	return base64.RawURLEncoding.EncodeToString(digest[:6])
}

func (s *Service) GetUserLinks(ctx context.Context, in *GetUserLinksInput) (*GetUserLinksOutput, error) {
	links, err := s.linkRepo.GetAllByUser(ctx, in.UID)
	if err != nil {
		return nil, err
	}

	var filteredLinks []link.Link
	for _, l := range links {
		if l.IsExpired() != nil || l.IsRevoked() != nil {
			continue
		}
		filteredLinks = append(filteredLinks, l)
	}

	return &GetUserLinksOutput{
		Links: filteredLinks,
	}, nil
}

func (s *Service) CreateLink(ctx context.Context, in *CreateLinkInput) (*CreateLinkOutput, error) {
	var timeStr *string
	if in.HoursToExpire != nil {
		expiration := time.Now().Add(time.Duration(*in.HoursToExpire) * time.Hour)
		formattedExpiration := expiration.Format(time.RFC3339)
		timeStr = &formattedExpiration
	}

	token := generateToken(uint64(time.Now().UnixNano() + int64(in.DocumentID)))
	l, err := s.linkRepo.Create(ctx, token, in.UID, in.DocumentID, in.ReadAllowed, in.WriteAllowed, in.DeleteAllowed, timeStr)
	if err != nil {
		return nil, err
	}

	return &CreateLinkOutput{
		Link: *l,
	}, nil
}

func (s *Service) RevokeLink(ctx context.Context, in *RevokeLinkInput) (*RevokeLinkOutput, error) {
	l, err := s.linkRepo.GetByToken(ctx, in.Token)
	if err != nil {
		return nil, err
	}

	if err := l.CanRevoke(in.UID); err != nil {
		return nil, err
	}

	err = s.linkRepo.Revoke(ctx, in.Token)
	if err != nil {
		return nil, err
	}

	return &RevokeLinkOutput{}, nil
}

func (s *Service) GetDocument(ctx context.Context, in *GetDocumentInput) (*docv1.GetUnsafeDocumentResponse, error) {
	l, err := s.linkRepo.GetByToken(ctx, in.Token)
	if err != nil {
		return nil, err
	}

	if err := l.CanRead(); err != nil {
		return nil, err
	}

	doc, err := s.documentRepo.Get(ctx, l.DocumentID)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func (s *Service) UpdateDocument(ctx context.Context, in *UpdateDocumentInput) (*docv1.UpdateUnsafeDocumentResponse, error) {
	l, err := s.linkRepo.GetByToken(ctx, in.Token)
	if err != nil {
		return nil, err
	}

	if err := l.CanWrite(); err != nil {
		return nil, err
	}

	doc, err := s.documentRepo.Update(ctx, l.DocumentID, in.Content, in.Name, in.IsPrivate)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func (s *Service) DeleteDocument(ctx context.Context, in *DeleteDocumentInput) (*docv1.DeleteUnsafeDocumentResponse, error) {
	l, err := s.linkRepo.GetByToken(ctx, in.Token)
	if err != nil {
		return nil, err
	}

	if err := l.CanDelete(); err != nil {
		return nil, err
	}

	if err := s.linkRepo.Revoke(ctx, in.Token); err != nil {
		return nil, err
	}

	doc, err := s.documentRepo.Delete(ctx, l.DocumentID)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
