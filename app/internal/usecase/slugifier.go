package usecase

import (
	"context"
	"fmt"
)

type retrieverSender interface {
	RetrieveURL(ctx context.Context, slug string) (string, error)
	CreateSlug(ctx context.Context, url, slug string) error
}

type hasher interface {
	GenerateShortLink(initialLink string) string
}

type Slugifier struct {
	retrieverSender retrieverSender
	hasher          hasher
}

func NewSlugifier(retrieverSender retrieverSender, hasher hasher) *Slugifier {
	return &Slugifier{
		retrieverSender: retrieverSender,
		hasher:          hasher,
	}
}

func (s *Slugifier) RetrieveURL(ctx context.Context, slug string) (string, error) {
	url, err := s.retrieverSender.RetrieveURL(ctx, slug)
	if err != nil {
		return "", fmt.Errorf("when s.retriever.RetrieveURL: %w", err)
	}
	return url, nil
}

func (s *Slugifier) CreateSlug(ctx context.Context, url string) (string, error) {
	slug := s.hasher.GenerateShortLink(url)
	err := s.retrieverSender.CreateSlug(ctx, url, slug)
	if err != nil {
		return "", fmt.Errorf("not able to create slug: %w", err)
	}
	return slug, nil
}
