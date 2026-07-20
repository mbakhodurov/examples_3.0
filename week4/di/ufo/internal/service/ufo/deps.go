package ufo

import (
	"context"

	"github.com/mbakhodurov/examples2/week_4/di/ufo/internal/model"
	"github.com/mbakhodurov/examples2/week_4/di/ufo/internal/service/input"
)

// UFORepository определяет контракт для работы с хранилищем наблюдений НЛО
type UFORepository interface {
	Create(ctx context.Context, sighting model.Sighting) error
	Get(ctx context.Context, uuid string) (model.Sighting, error)
	Update(ctx context.Context, uuid string, in input.UpdateSightingInput) error
	Delete(ctx context.Context, uuid string) error
}
