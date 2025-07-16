package brasilapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jeffersono7/pos-go-expert-desafios/desafio-multithreading/internal/client"
)

const url = "https://brasilapi.com.br/api/cep/v1"

type BrasilAPIClient struct {
}

func (b BrasilAPIClient) GetCEP(ctx context.Context, cep string) (*client.CepResponse, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/%s", url, cep), nil)
	if err != nil {
		return nil, fmt.Errorf("fail new req: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fail do request: %w", err)
	}
	defer resp.Body.Close()

	var cepResp client.CepResponse
	if err = json.NewDecoder(resp.Body).Decode(&cepResp); err != nil {
		return nil, fmt.Errorf("fail unmarshal resp: %w", err)
	}

	cepResp.Api = "brasilapi"

	return &cepResp, nil
}
