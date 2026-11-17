package link

import (
	"errors"
	"time"
)

var ErrLinkExpired = errors.New("link has expired")
var ErrLinkRevoked = errors.New("link has been revoked")
var ErrOperationDeclined = errors.New("operation declined")

func (l *Link) IsExpired() error {
	if l.ExpiresAt == nil {
		return nil
	}
	if l.ExpiresAt.Before(time.Now()) {
		return ErrLinkExpired
	}
	return nil
}

func (l *Link) IsRevoked() error {
	if l.RevokedAt == nil {
		return nil
	}
	return ErrLinkRevoked
}

func (l *Link) CanRead() error {
	if !l.ReadAllowed {
		return ErrOperationDeclined
	}
	if err := l.IsExpired(); err != nil {
		return err
	}
	if err := l.IsRevoked(); err != nil {
		return err
	}
	return nil
}

func (l *Link) CanWrite() error {
	if !l.WriteAllowed {
		return ErrOperationDeclined
	}
	if err := l.IsExpired(); err != nil {
		return err
	}
	if err := l.IsRevoked(); err != nil {
		return err
	}
	return nil
}

func (l *Link) CanDelete() error {
	if !l.DeleteAllowed {
		return ErrOperationDeclined
	}
	if err := l.IsExpired(); err != nil {
		return err
	}
	if err := l.IsRevoked(); err != nil {
		return err
	}
	return nil
}
