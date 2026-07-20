package pc_builder

import (
	"context"

	"github.com/mbakhodurov/examples2/week_4/ddd/app/internal/model/valueobject"
)

// PCBuilderService определяет контракт сервиса сборки ПК для API-слоя
type PCBuilderService interface {
	CreateBuild(ctx context.Context, componentUUIDs []string) (string, valueobject.BuildStatus, error)
	CancelBuild(ctx context.Context, buildUUID string) (valueobject.BuildStatus, error)
}
