package service

import "context"

type Cep struct {
	Localidade string `json:"localidade"`
	UF         string `json:"uf"`
}

type CepService interface {
	GetCEP(ctx context.Context, cep string) (Cep, error)
}
