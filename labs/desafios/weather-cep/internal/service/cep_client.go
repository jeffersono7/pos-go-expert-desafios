package service

import "context"

type CepResp struct {
	Localidade string `json:"localidade"`
	UF         string `json:"uf"`
	Estado     string `json:"estado"`
}

type CepClient interface {
	GetCEP(ctx context.Context, cep string) (CepResp, error)
}
