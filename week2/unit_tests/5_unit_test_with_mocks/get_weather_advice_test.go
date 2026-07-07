package main

import (
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/mbakhodurov/examples2/week_2/unit_tests/5_unit_test_with_mocks/mocks"
	weatherCenter "github.com/mbakhodurov/examples2/week_2/unit_tests/5_unit_test_with_mocks/weather_center"
	"github.com/stretchr/testify/require"
)

func TestGetWeatherAdvice(t *testing.T) {
	type weatherCenterClientMockFunc func(t *testing.T) WeatherCenterClient

	city := gofakeit.City()

	tests := []struct {
		name                    string
		city                    string
		expected                string
		errCheck                func(t *testing.T, err error)
		weatherCenterClientMock weatherCenterClientMockFunc
	}{
		{
			name:     "Температура +25 градусов",
			city:     city,
			expected: "Отличная погода для прогулок",
			errCheck: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
			weatherCenterClientMock: func(t *testing.T) WeatherCenterClient {
				mockClient := mocks.NewWeatherCenterClient(t)
				mockClient.EXPECT().GetTemperature(city).Return(float32(25), nil)

				return mockClient
			},
		},
		{
			name:     "Температура -15 градусов",
			city:     city,
			expected: "Прохладно, но можно выйти на улицу",
			errCheck: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
			weatherCenterClientMock: func(t *testing.T) WeatherCenterClient {
				mockClient := mocks.NewWeatherCenterClient(t)
				mockClient.EXPECT().GetTemperature(city).Return(float32(-15), nil)

				return mockClient
			},
		},
		{
			name:     "Город не найден",
			city:     city,
			expected: "",
			errCheck: func(t *testing.T, err error) {
				require.Error(t, err)
				require.ErrorIs(t, err, weatherCenter.ErrCityNotFound)
			},
			weatherCenterClientMock: func(t *testing.T) WeatherCenterClient {
				mockClient := mocks.NewWeatherCenterClient(t)
				mockClient.EXPECT().GetTemperature(city).Return(float32(0), weatherCenter.ErrCityNotFound)

				return mockClient
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res, err := getWeatherAdvice(tc.weatherCenterClientMock(t), tc.city)
			tc.errCheck(t, err)
			require.Equal(t, tc.expected, res)
		})
	}
}
