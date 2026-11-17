package document

import "time"

type Document struct {
	ID        int
	OwnerID   string
	Name      string
	Content   string
	IsPrivate bool
	IsDeleted bool
	CreatedAt time.Time
	EditedAt  time.Time
}
