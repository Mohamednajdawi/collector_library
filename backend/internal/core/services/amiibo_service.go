package services

import (
	"collector-library/internal/core/domain"
	"collector-library/internal/core/ports"
	"context"
)

type AmiiboService struct {
	repo ports.AmiiboRepository
}

func NewAmiiboService(repo ports.AmiiboRepository) *AmiiboService {
	return &AmiiboService{
		repo: repo,
	}
}

func (s *AmiiboService) ListAmiibos(ctx context.Context) ([]domain.Amiibo, error) {
	return s.repo.List(ctx)
}

func (s *AmiiboService) IngestAmiibos(ctx context.Context, amiibos []domain.Amiibo) error {
	return s.repo.CreateBatch(ctx, amiibos)
}
