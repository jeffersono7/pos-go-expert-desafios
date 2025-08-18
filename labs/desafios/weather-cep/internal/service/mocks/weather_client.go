package mocks

import (
	"context"

	"github.com/jeffersono7/pos-go-expert-desafios/labs/desafios/weather-cep/internal/service"
	"github.com/stretchr/testify/mock"
)

type WeatherClientMock struct {
	mock.Mock
}

func (wc *WeatherClientMock) GetTemp(ctx context.Context, neighborhook string) (service.WeatherResp, error) {
	args := wc.Called(ctx, neighborhook)
	return args.Get(0).(service.WeatherResp), args.Error(1)
}
