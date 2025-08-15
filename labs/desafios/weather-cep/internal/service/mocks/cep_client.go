package mocks

import (
	"context"

	"github.com/jeffersono7/pos-go-expert-desafios/labs/desafios/weather-cep/internal/service"
	"github.com/stretchr/testify/mock"
)

type CepClientMock struct {
	mock.Mock
}

func (c *CepClientMock) GetCEP(ctx context.Context, cep string) (service.CepResp, error) {
	args := c.Called(ctx, cep)
	return args.Get(0).(service.CepResp), args.Error(1)
}
