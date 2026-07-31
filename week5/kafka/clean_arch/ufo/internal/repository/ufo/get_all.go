package ufo

import (
	"context"

	"github.com/mbakhodurov/examples2/week_5/kafka/clean_arch/ufo/internal/model"
)

func (r *repository) GetAll(ctx context.Context) ([]model.Sighting, error) {
	query := `SELECT uuid, observed_at, location, description, color, sound, duration_seconds, created_at, updated_at, deleted_at
		FROM sightings WHERE deleted_at IS NULL ORDER BY created_at DESC`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sightings []model.Sighting
	for rows.Next() {
		var s model.Sighting
		if err := rows.Scan(
			&s.Uuid, &s.ObservedAt, &s.Location, &s.Description,
			&s.Color, &s.Sound, &s.DurationSeconds,
			&s.CreatedAt, &s.UpdatedAt, &s.DeletedAt,
		); err != nil {
			return nil, err
		}
		sightings = append(sightings, s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return sightings, nil
}
