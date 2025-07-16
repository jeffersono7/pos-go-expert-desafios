package viacep

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jeffersono7/pos-go-expert-desafios/desafio-multithreading/internal/client"
)

const url = "http://viacep.com.br/ws"

type ViaCepAPIClient struct {
}

func (b ViaCepAPIClient) GetCEP(ctx context.Context, cep string, channelResp chan<- client.CepResponse) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/%s/json/", url, cep), nil)
	if err != nil {
		// fmt.Printf("fail new req: %v\n", err)
		close(channelResp)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// fmt.Printf("fail do request: %v\n", err)
		close(channelResp)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		// fmt.Printf("api returned error: %v\n", resp.Status)
		close(channelResp)
		return
	}

	var cepResp client.CepResponse
	if err = json.NewDecoder(resp.Body).Decode(&cepResp); err != nil {
		// fmt.Printf("fail unmarshal resp: %v\n", err)
		close(channelResp)
		return
	}

	cepResp.Api = "viacep"

	channelResp <- cepResp
	close(channelResp)
}
