package link

import "time"

type Link struct {
	Token         string
	DocumentID    int
	CreatorID     string
	ReadAllowed   bool
	WriteAllowed  bool
	DeleteAllowed bool
	CreatedAt     time.Time
	ExpiresAt     *time.Time
	RevokedAt     *time.Time
}
