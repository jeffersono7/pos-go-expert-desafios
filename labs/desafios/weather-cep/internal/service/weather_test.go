package service

import (
	"testing"

	"github.com/jeffersono7/pos-go-expert-desafios/labs/desafios/weather-cep/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewWeather(t *testing.T) {
	subject := NewWeatherService(nil, nil)
	assert.NotNil(t, subject)
}

func TestGetWeatherFromCEP(t *testing.T) {
	suite := []struct {
		description    string
		input          string
		expectedErr    error
		expectedResult domain.Weather
	}{
		{
			description:    "when is invalid cep",
			input:          "11122",
			expectedErr:    ErrInvalidCEP,
			expectedResult: domain.Weather{},
		},
	}

	for _, cc := range suite {
		t.Run(cc.description, func(t *testing.T) {
			subject := NewWeatherService(nil, nil)

			actual, err := subject.GetWeatherFromCEP(cc.input)

			if cc.expectedErr != nil {
				assert.Equal(t, cc.expectedErr, err)
			}
			assert.Equal(t, cc.expectedResult, actual)
		})
	}
}
