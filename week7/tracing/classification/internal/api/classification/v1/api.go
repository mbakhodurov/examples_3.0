package classification_v1

import (
	classificationv1 "github.com/mbakhodurov/homeworks2/week7/tracing/shared/pkg/proto/classification/v1"
)

// api — gRPC хендлер для сервиса классификации
type api struct {
	classificationv1.UnimplementedClassificationServiceServer
	classificationService ClassificationService
}

// NewAPI создаёт новый API классификации
func NewAPI(classificationService ClassificationService) *api {
	return &api{
		classificationService: classificationService,
	}
}
