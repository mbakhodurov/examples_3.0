package converter

import (
	"github.com/mbakhodurov/examples2/week_2/layers/internal/model"
	weather_v1 "github.com/mbakhodurov/examples2/week_2/layers/pkg/proto/weather/v1"
)

// WeatherConditionsFromProto конвертирует транспортный DTO погодных условий в доменную модель
func WeatherConditionsFromProto(conditions *weather_v1.WeatherConditions) model.WeatherConditions {
	return model.WeatherConditions{
		Description:        conditions.GetDescription(),
		TemperatureCelsius: conditions.GetTemperatureCelsius(),
		VisibilityKm:       conditions.GetVisibilityKm(),
	}
}
