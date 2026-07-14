package ufo

import (
	"context"

	"github.com/mbakhodurov/examples2/week_4/config/ufo/internal/model"
)

func (s *service) Get(ctx context.Context, uuid string) (model.Sighting, error) {
	return s.ufoRepository.Get(ctx, uuid)
}
