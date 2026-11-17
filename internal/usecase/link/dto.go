package link

import (
	"github.com/dockloud/link/internal/domain/link"
)

type GetUserLinksInput struct {
	UID string
}

type GetUserLinksOutput struct {
	Links []link.Link
}

type CreateLinkInput struct {
	UID           string
	DocumentID    int
	ReadAllowed   bool
	WriteAllowed  bool
	DeleteAllowed bool
	HoursToExpire *int
}

type RevokeLinkInput struct {
	Token string
}

type GetDocumentInput struct {
	Token string
}

type UpdateDocumentInput struct {
	Token     string
	Content   *string
	Name      *string
	IsPrivate *bool
}

type DeleteDocumentInput struct {
	Token string
}
