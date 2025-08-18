package service_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/jeffersono7/pos-go-expert-desafios/labs/desafios/weather-cep/internal/domain"
	"github.com/jeffersono7/pos-go-expert-desafios/labs/desafios/weather-cep/internal/service"
	"github.com/jeffersono7/pos-go-expert-desafios/labs/desafios/weather-cep/internal/service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewWeather(t *testing.T) {
	subject := service.NewWeatherService(nil, nil)
	assert.NotNil(t, subject)
}

func TestGetWeatherFromCEP(t *testing.T) {
	validCep := "74343135"

	suite := []struct {
		description    string
		input          string
		expectedErr    error
		expectedResult domain.Weather
		setup          func(cepClient *mocks.CepClientMock, weatherClient *mocks.WeatherClientMock)
	}{
		{
			description:    "when is invalid cep",
			input:          "11122",
			expectedErr:    service.ErrInvalidCEP,
			expectedResult: domain.Weather{},
		},
		{
			description:    "when cep is not found",
			input:          "11100077",
			expectedErr:    service.ErrNotFoundCEP,
			expectedResult: domain.Weather{},
			setup: func(cepClient *mocks.CepClientMock, _ *mocks.WeatherClientMock) {
				cepClient.On("GetCEP", mock.Anything, "11100077").Return(service.CepResp{}, errors.New("an error"))
			},
		},
		{
			description:    "when fail get weather",
			input:          validCep,
			expectedErr:    service.ErrFailGetWeather,
			expectedResult: domain.Weather{},
			setup: func(cepClient *mocks.CepClientMock, weatherClient *mocks.WeatherClientMock) {
				cepClient.On("GetCEP", mock.Anything, validCep).Return(service.CepResp{Localidade: "sobradinho", UF: "DF", Estado: "Distrito Federal"}, nil)
				weatherClient.On("GetTemp", mock.Anything, "sobradinho Distrito Federal").Return(service.WeatherResp{}, fmt.Errorf("an error"))
			},
		},
		{
			description:    "when successfully get weather",
			input:          validCep,
			expectedErr:    nil,
			expectedResult: domain.Weather{TempC: 29, TempF: 84.2, TempK: 302},
			setup: func(cepClient *mocks.CepClientMock, weatherClient *mocks.WeatherClientMock) {
				cepClient.On("GetCEP", mock.Anything, validCep).Return(service.CepResp{Localidade: "sobradinho", UF: "DF", Estado: "Distrito Federal"}, nil)
				weatherClient.On("GetTemp", mock.Anything, "sobradinho Distrito Federal").
					Return(service.WeatherResp{
						Current: struct {
							TempC float32 "json:\"temp_c\""
						}{
							TempC: 29.0,
						}},
						nil)
			},
		},
	}

	for _, cc := range suite {
		t.Run(cc.description, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			cepClientMock := new(mocks.CepClientMock)
			weatherClientMock := new(mocks.WeatherClientMock)
			subject := service.NewWeatherService(cepClientMock, weatherClientMock)

			if cc.setup != nil {
				cc.setup(cepClientMock, weatherClientMock)
			}
			actual, err := subject.GetWeatherFromCEP(ctx, cc.input)

			if cc.expectedErr != nil {
				assert.Equal(t, cc.expectedErr, err)
			}
			assert.Equal(t, cc.expectedResult, actual)
			cepClientMock.AssertExpectations(t)
		})
	}
}
