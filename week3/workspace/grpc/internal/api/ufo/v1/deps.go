package v1

import (
	"context"

	"github.com/mbakhodurov/examples2/week_3/workspace/grpc/internal/model"
	"github.com/mbakhodurov/examples2/week_3/workspace/grpc/internal/service/input"
)

// UFORepository — интерфейс репозитория для работы с наблюдениями НЛО
type UFORepository interface {
	Create(ctx context.Context, in input.CreateSightingInput) (string, error)
	Get(ctx context.Context, uuid string) (model.Sighting, error)
	Update(ctx context.Context, uuid string, in input.UpdateSightingInput) error
	Delete(ctx context.Context, uuid string) error
}
