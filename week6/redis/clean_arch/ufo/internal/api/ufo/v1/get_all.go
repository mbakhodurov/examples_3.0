package v1

import (
	"context"

	ufo_v1 "github.com/mbakhodurov/examples2/week_6/redis/clean_arch/shared/pkg/proto/ufo/v1"
	"github.com/mbakhodurov/examples2/week_6/redis/clean_arch/ufo/internal/api/converter"
)

func (a *api) GetAll(ctx context.Context, _ *ufo_v1.GetAllRequest) (*ufo_v1.GetAllResponse, error) {
	sightings, err := a.ufoService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return &ufo_v1.GetAllResponse{
		Sightings: converter.SightingsToDTO(sightings),
	}, nil
}
