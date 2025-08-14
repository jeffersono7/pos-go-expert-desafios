package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jeffersono7/pos-go-expert-desafios/labs/desafios/weather-cep/service"
)

type CepServiceImpl struct {
}

func NewCepServiceImpl() *CepServiceImpl {
	return &CepServiceImpl{}
}

func (cs CepServiceImpl) GetCEP(ctx context.Context, cep string) (service.Cep, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep), nil)
	if err != nil {
		return service.Cep{}, fmt.Errorf("fail make req: %w", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return service.Cep{}, fmt.Errorf("fail do req: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return service.Cep{}, fmt.Errorf("request returns an error with status: %s", res.Status)
	}

	var cepRes service.Cep
	if err := json.NewDecoder(res.Body).Decode(&cepRes); err != nil {
		return service.Cep{}, fmt.Errorf("fail unmarshal response: %w", err)
	}

	return cepRes, nil
}
