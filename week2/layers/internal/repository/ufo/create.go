package ufo

import (
	"context"

	"github.com/mbakhodurov/examples2/week_2/layers/internal/model"
	"github.com/mbakhodurov/examples2/week_2/layers/internal/repository/converter"
)

func (r *repository) Create(_ context.Context, sighting model.Sighting) error {
	r.mu.Lock()
	r.mu.Unlock()

	r.data[sighting.Uuid] = converter.SightingToRepoModel(sighting)
	return nil
}
