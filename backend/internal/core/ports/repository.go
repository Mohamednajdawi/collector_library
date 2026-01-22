package ports

import (
	"context"
	"collector-library/internal/core/domain"
)

// AmiiboRepository defines the interface for storing and retrieving Amiibos.
type AmiiboRepository interface {
	// List returns a list of all Amiibos.
	List(ctx context.Context) ([]domain.Amiibo, error)
	// CreateBatch inserts multiple Amiibos into the store.
	CreateBatch(ctx context.Context, amiibos []domain.Amiibo) error
}
