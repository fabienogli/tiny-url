package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/fabienogli/stoik/internal/domain"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type sqlConnector interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}

type SlugRepository struct {
	conn sqlConnector
}

func NewSlugRepository(conn *pgx.Conn) *SlugRepository {
	return &SlugRepository{
		conn: conn,
	}
}

func (s *SlugRepository) RetrieveURL(ctx context.Context, slug string) (string, error) {
	var url string
	row := s.conn.QueryRow(ctx, `SELECT url From tiny_url where slug=$1`, slug)
	err := row.Scan(&url)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", fmt.Errorf("when no Result: %w", domain.ErrNoResult)
	}
	if err != nil {
		return "", fmt.Errorf("when QueryRow: %w", err)
	}
	return url, nil
}

func (s *SlugRepository) CreateSlug(ctx context.Context, url, slug string) error {
	_, err := s.conn.Exec(ctx, `INSERT INTO tiny_url(url,slug) VALUES ($1,$2)`, url, slug)
	if err != nil {
		return fmt.Errorf("when creating: %w", err)
	}
	return nil
}
