package ufo

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/mbakhodurov/examples2/week_6/redis/clean_arch/ufo/internal/model"
)

func (s *service) GetAll(ctx context.Context) ([]model.Sighting, error) {
	sightings, err := s.ufoRepo.GetAll(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "не удалось получить список наблюдений", "error", err)
		return nil, fmt.Errorf("получить список наблюдений: %w", err)
	}

	return sightings, nil
}
