package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWeather(t *testing.T) {
	var (
		expectedC float32 = 29.8
		expectedF float32 = 85.64
		expectedK float32 = 302.8
	)

	actual := NewWeather(expectedC)

	assert.Equal(t, expectedC, actual.TempC)
	assert.Equal(t, expectedF, actual.TempF)
	assert.Equal(t, expectedK, actual.TempK)
}
