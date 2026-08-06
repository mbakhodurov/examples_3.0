package ufo

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/mbakhodurov/examples2/week_6/redis/clean_arch/ufo/internal/model"
	repoConverter "github.com/mbakhodurov/examples2/week_6/redis/clean_arch/ufo/internal/repository/converter"
	repoModel "github.com/mbakhodurov/examples2/week_6/redis/clean_arch/ufo/internal/repository/model"
)

func (r *repository) GetAll(ctx context.Context) ([]model.Sighting, error) {
	query := `SELECT uuid, observed_at, location, description, color, sound, duration_seconds, created_at, updated_at, deleted_at
		FROM sightings WHERE deleted_at IS NULL ORDER BY created_at DESC`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	pgViews, err := pgx.CollectRows(rows, pgx.RowToStructByName[repoModel.SightingPGView])
	if err != nil {
		return nil, err
	}

	sightings := make([]model.Sighting, 0, len(pgViews))
	for _, pgView := range pgViews {
		sightings = append(sightings, repoConverter.SightingFromPGView(pgView))
	}

	return sightings, nil
}
