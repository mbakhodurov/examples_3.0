package classification_v1

import (
	"context"

	"github.com/mbakhodurov/homeworks2/week7/tracing/classification/internal/model"
)

// ClassificationService — интерфейс сервиса классификации
type ClassificationService interface {
	ClassifyObject(ctx context.Context, description, color string, durationSeconds int32) (model.ClassificationResult, error)
}
