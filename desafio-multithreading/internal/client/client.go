package client

import "context"

type CepResponse struct {
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Api          string `json:"api"`
}

type CepAPI interface {
	GetCEP(ctx context.Context, cep string, channelResponse chan<- CepResponse)
}
