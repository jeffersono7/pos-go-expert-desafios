package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWeather(t *testing.T) {
	subject := NewWeatherService(nil, nil)
	assert.NotNil(t, subject)
}

func TestGetWeatherFromCEP(t *testing.T) {

}
