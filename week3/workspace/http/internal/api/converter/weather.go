// Package converter содержит функции преобразования между транспортными DTO
// и доменными моделями сервиса погоды
package converter

import (
	"github.com/mbakhodurov/examples2/week_3/workspace/http/internal/model"
	weatherv1 "github.com/mbakhodurov/examples2/week_3/workspace/shared/pkg/openapi/weather/v1"
)

// WeatherToDTO конвертирует доменную модель погоды в транспортный DTO
func WeatherToDTO(w model.Weather) *weatherv1.Weather {
	return &weatherv1.Weather{
		City:        w.City,
		Temperature: w.Temperature,
		UpdatedAt:   w.UpdatedAt,
	}
}
