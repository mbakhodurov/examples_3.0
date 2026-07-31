package ufo

import (
	"context"

	"github.com/mbakhodurov/examples2/week_5/kafka/clean_arch/ufo/internal/model"
	"github.com/mbakhodurov/examples2/week_5/kafka/clean_arch/ufo/internal/service/input"
)

// UFORepository определяет контракт для работы с хранилищем наблюдений НЛО
type UFORepository interface {
	Create(ctx context.Context, sighting model.Sighting) error
	Get(ctx context.Context, uuid string) (model.Sighting, error)
	GetAll(ctx context.Context) ([]model.Sighting, error)
	Update(ctx context.Context, uuid string, updateInfo input.UpdateSightingInput) error
	Delete(ctx context.Context, uuid string) error
}

// UFOProducerService определяет контракт для отправки событий о наблюдениях НЛО в Kafka
type UFOProducerService interface {
	ProduceUFORecorded(ctx context.Context, event model.UFORecordedEvent) error
}
