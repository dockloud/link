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

type CreateLinkOutput struct {
	Link link.Link
}

type RevokeLinkInput struct {
	UID   string
	Token string
}

type RevokeLinkOutput struct {
	// empty
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
