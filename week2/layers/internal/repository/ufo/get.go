package ufo

import (
	"context"

	errs "github.com/mbakhodurov/examples2/week_2/layers/internal/errors"
	"github.com/mbakhodurov/examples2/week_2/layers/internal/model"
	repoConverter "github.com/mbakhodurov/examples2/week_2/layers/internal/repository/converter"
)

func (r *repository) Get(_ context.Context, uuid string) (model.Sighting, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	repoSighting, ok := r.data[uuid]

	if !ok {
		return model.Sighting{}, errs.ErrSightingNotFound
	}

	if repoSighting.DeletedAt != nil {
		return model.Sighting{}, errs.ErrSightingNotFound
	}

	return repoConverter.SightingToModel(repoSighting), nil
}
