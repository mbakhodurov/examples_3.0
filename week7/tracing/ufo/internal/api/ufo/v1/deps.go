package v1

import (
	"context"

	"github.com/mbakhodurov/homeworks2/week7/tracing/ufo/internal/model"
	"github.com/mbakhodurov/homeworks2/week7/tracing/ufo/internal/service/input"
)

// UFOService интерфейс бизнес-логики UFO сервиса
type UFOService interface {
	Create(ctx context.Context, in input.CreateSightingInput) (string, error)
	Get(ctx context.Context, uuid string) (model.Sighting, error)
	Update(ctx context.Context, uuid string, updateInfo input.UpdateSightingInput) error
	Delete(ctx context.Context, uuid string) error
	AnalyzeSighting(ctx context.Context, uuid string) (model.AnalysisResult, error)
}
