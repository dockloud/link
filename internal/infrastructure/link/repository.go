package link

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/dockloud/link/internal/domain/link"
	"github.com/redis/go-redis/v9"
)

var ErrNotFound = errors.New("link not found")
var logger = log.Default()

type Repository struct {
	db    *sql.DB
	cache *redis.Client
}

func NewRepository(db *sql.DB, cache *redis.Client) *Repository {
	return &Repository{
		db:    db,
		cache: cache,
	}
}

func scanLink(scanner interface {
	Scan(dest ...any) error
}) (*link.Link, error) {
	var l link.Link

	err := scanner.Scan(
		&l.Token,
		&l.DocumentID,
		&l.CreatorID,
		&l.ReadAllowed,
		&l.WriteAllowed,
		&l.DeleteAllowed,
		&l.CreatedAt,
		&l.ExpiresAt,
		&l.RevokedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &l, nil
}

func (r *Repository) GetByToken(ctx context.Context, token string) (*link.Link, error) {
	var l *link.Link
	linkKey := fmt.Sprintf("link:%s", token)
	jsonLink, err := r.cache.Get(ctx, linkKey).Result()

	switch {
	case err == nil:
		err := json.Unmarshal([]byte(jsonLink), &l)
		if err != nil {
			logger.Printf("Error unmarshaling cached link: %v", err)
			return nil, err
		}
		return l, nil
	case errors.Is(err, redis.Nil):
		// cache miss, proceed to fetch from the database
	default:
		logger.Printf("Error fetching link from cache: %v", err)
		return nil, err
	}

	query := `
		SELECT token, document_id, creator_id, 
		       read_allowed, write_allowed, delete_allowed,
		       created_at, expires_at, revoked_at
		FROM links
		WHERE token = $1
	`
	row := r.db.QueryRowContext(ctx, query, token)
	l, err = scanLink(row)
	if err != nil {
		return nil, err
	}

	linkJson, err := json.Marshal(l)
	if err != nil {
		logger.Printf("Error marshaling link: %v", err)
	}
	err = r.cache.Set(ctx, linkKey, linkJson, 0).Err()
	if err != nil {
		logger.Printf("Error setting link in cache: %v", err)
	}

	return l, nil
}

func (r *Repository) GetAllByUser(ctx context.Context, uid string) ([]*link.Link, error) {
	var links []*link.Link

	query := `
		SELECT token, document_id, creator_id, 
		       read_allowed, write_allowed, delete_allowed,
		       created_at, expires_at, revoked_at
		FROM links
		WHERE creator_id = $1
	`
	rows, err := r.db.QueryContext(ctx, query, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		l, err := scanLink(rows)
		if err != nil {
			return nil, err
		}
		links = append(links, l)
	}

	return links, nil
}

func (r *Repository) Create(
	ctx context.Context, token, creatorID string, documentID int,
	readAllowed, writeAllowed, deleteAllowed bool, expiresAt *string,
) (*link.Link, error) {
	query := `
		INSERT INTO links (token, document_id, creator_id, read_allowed, write_allowed, delete_allowed, expires_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING token, document_id, creator_id, read_allowed, write_allowed, delete_allowed, created_at, expires_at, revoked_at
	`
	row := r.db.QueryRowContext(ctx, query,
		token,
		documentID,
		creatorID,
		readAllowed,
		writeAllowed,
		deleteAllowed,
		expiresAt,
	)

	l, err := scanLink(row)
	if err != nil {
		return nil, err
	}

	linkKey := fmt.Sprintf("link:%s", token)
	linkJson, err := json.Marshal(l)
	if err != nil {
		logger.Printf("Error marshaling link: %v", err)
	}
	err = r.cache.Set(ctx, linkKey, linkJson, 0).Err()
	if err != nil {
		logger.Printf("Error setting link in cache: %v", err)
	}

	return l, nil
}

func (r *Repository) Revoke(ctx context.Context, token string) error {
	query := `
		UPDATE links
		SET revoked_at = NOW()
		WHERE token = $1
	`
	_, err := r.db.ExecContext(ctx, query, token)
	if err != nil {
		return err
	}

	linkKey := fmt.Sprintf("link:%s", token)
	err = r.cache.Del(ctx, linkKey).Err()
	if err != nil {
		logger.Printf("Error deleting link from cache: %v", err)
	}

	return nil
}
