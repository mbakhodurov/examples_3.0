package v1

import (
	"context"

	ufov1 "github.com/mbakhodurov/examples2/week_5/kafka/clean_arch/shared/pkg/proto/ufo/v1"
	"github.com/mbakhodurov/examples2/week_5/kafka/clean_arch/ufo/internal/api/converter"
)

func (a *api) GetAll(ctx context.Context, _ *ufov1.GetAllRequest) (*ufov1.GetAllResponse, error) {
	sightings, err := a.ufoService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	dtos := make([]*ufov1.Sighting, 0, len(sightings))
	for _, s := range sightings {
		dtos = append(dtos, converter.SightingToDTO(s))
	}

	return &ufov1.GetAllResponse{
		Sightings: dtos,
	}, nil
}
