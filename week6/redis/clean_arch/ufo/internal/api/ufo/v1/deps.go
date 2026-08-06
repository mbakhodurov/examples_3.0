package v1

import (
	"context"

	"github.com/mbakhodurov/examples2/week_6/redis/clean_arch/ufo/internal/model"
	"github.com/mbakhodurov/examples2/week_6/redis/clean_arch/ufo/internal/service/input"
)

// UFOService определяет контракт бизнес-логики для работы с наблюдениями НЛО
type UFOService interface {
	Create(ctx context.Context, in input.CreateSightingInput) (string, error)
	Get(ctx context.Context, uuid string) (model.Sighting, error)
	GetAll(ctx context.Context) ([]model.Sighting, error)
	Update(ctx context.Context, uuid string, updateInfo input.UpdateSightingInput) error
	Delete(ctx context.Context, uuid string) error
}
