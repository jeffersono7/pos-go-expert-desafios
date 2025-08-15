package service_test

import (
	"context"
	"errors"
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
	suite := []struct {
		description    string
		input          string
		expectedErr    error
		expectedResult domain.Weather
		setup          func(cepClient *mocks.CepClientMock)
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
			setup: func(cepClient *mocks.CepClientMock) {
				cepClient.On("GetCEP", mock.Anything, "11100077").Return(service.CepResp{}, errors.New("an error"))
			},
		},
	}

	for _, cc := range suite {
		t.Run(cc.description, func(t *testing.T) {
			ctx := context.Background()

			cepClientMock := new(mocks.CepClientMock)
			subject := service.NewWeatherService(cepClientMock, nil)

			if cc.setup != nil {
				cc.setup(cepClientMock)
			}
			actual, err := subject.GetWeatherFromCEP(ctx, cc.input)

			if cc.expectedErr != nil {
				assert.Equal(t, cc.expectedErr, err)
			}
			assert.Equal(t, cc.expectedResult, actual)
		})
	}
}
