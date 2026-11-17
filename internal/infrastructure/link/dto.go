package link

import "time"

type CacheableLink struct {
	Token         string     `redis:"token"`
	DocumentID    int        `redis:"document_id"`
	CreatorID     string     `redis:"creator_id"`
	ReadAllowed   bool       `redis:"read_allowed"`
	WriteAllowed  bool       `redis:"write_allowed"`
	DeleteAllowed bool       `redis:"delete_allowed"`
	CreatedAt     time.Time  `redis:"created_at"`
	ExpiresAt     *time.Time `redis:"expires_at"`
	RevokedAt     *time.Time `redis:"revoked_at"`
}
