package link

import (
	"github.com/dockloud/link/internal/domain/link"
)

func ToCacheableDTO(link *link.Link) *CacheableLink {
	return &CacheableLink{
		Token:         link.Token,
		DocumentID:    link.DocumentID,
		CreatorID:     link.CreatorID,
		ReadAllowed:   link.ReadAllowed,
		WriteAllowed:  link.WriteAllowed,
		DeleteAllowed: link.DeleteAllowed,
		CreatedAt:     link.CreatedAt,
		ExpiresAt:     link.ExpiresAt,
		RevokedAt:     link.RevokedAt,
	}
}
