package ufo

import (
	"context"

	"github.com/mbakhodurov/examples2/week_4/config/ufo/internal/service/input"
)

func (s *service) Update(ctx context.Context, uuid string, in input.UpdateSightingInput) error {
	return s.ufoRepository.Update(ctx, uuid, in)
}
