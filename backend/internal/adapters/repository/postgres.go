package repository

import (
	"collector-library/internal/core/domain"
	"collector-library/internal/core/ports"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) ports.AmiiboRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) List(ctx context.Context) ([]domain.Amiibo, error) {
	query := `SELECT id, name, image_url, series, release_date FROM amiibos`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query amiibos: %w", err)
	}
	defer rows.Close()

	var amiibos []domain.Amiibo
	for rows.Next() {
		var a domain.Amiibo
		err := rows.Scan(&a.ID, &a.Name, &a.ImageURL, &a.Series, &a.ReleaseDate)
		if err != nil {
			return nil, fmt.Errorf("failed to scan amiibo: %w", err)
		}
		amiibos = append(amiibos, a)
	}
	return amiibos, nil
}

func (r *PostgresRepository) CreateBatch(ctx context.Context, amiibos []domain.Amiibo) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `INSERT INTO amiibos (name, image_url, series, release_date) VALUES ($1, $2, $3, $4)`

	for _, a := range amiibos {
		// Series defaults to 'Super Smash Bros' if empty, but we can set it explicitly
		series := a.Series
		if series == "" {
			series = "Super Smash Bros"
		}

		_, err := tx.Exec(ctx, query, a.Name, a.ImageURL, series, a.ReleaseDate)
		if err != nil {
			return fmt.Errorf("failed to insert amiibo %s: %w", a.Name, err)
		}
	}

	return tx.Commit(ctx)
}
