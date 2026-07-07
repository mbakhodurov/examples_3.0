package v1

import (
	"context"
	"fmt"

	"github.com/mbakhodurov/examples2/week_2/layers/internal/client/grpc/weather/v1/converter"
	errs "github.com/mbakhodurov/examples2/week_2/layers/internal/errors"
	"github.com/mbakhodurov/examples2/week_2/layers/internal/model"
	weather_v1 "github.com/mbakhodurov/examples2/week_2/layers/pkg/proto/weather/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type client struct {
	grpcClient weather_v1.WeatherServiceClient
}

// New создаёт обёртку над gRPC клиентом WeatherService
func New(grpcClient weather_v1.WeatherServiceClient) *client {
	return &client{
		grpcClient: grpcClient,
	}
}

// GetCurrentWeather получает текущие погодные условия по местоположению
func (c *client) GetCurrentWeather(ctx context.Context, location string) (model.WeatherConditions, error) {
	resp, err := c.grpcClient.GetCurrentWeather(ctx, &weather_v1.GetCurrentWeatherRequest{
		Location: location,
	})
	if err != nil {
		// Маппинг gRPC ошибок в доменные ошибки
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.Unavailable {
				return model.WeatherConditions{}, errs.ErrWeatherUnavailable
			}
		}

		return model.WeatherConditions{}, fmt.Errorf("получить текущую погоду: %w", err)
	}

	return converter.WeatherConditionsFromProto(resp.GetConditions()), nil
}
