package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jeffersono7/pos-go-expert-desafios/labs/desafios/otel/service_b/internal/domain"
)

type WeatherService struct {
	Client http.Client
}

func (s WeatherService) GetWeather(ctx context.Context, cep string) (*domain.Weather, error) {
	url := fmt.Sprintf("http://localhost:8081/cep/%s/weather", cep)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("fail new req: %w", err)
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fail do req: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response with status not ok: status=%s", resp.Status)
	}

	defer resp.Body.Close()

	var weather domain.Weather
	if err = json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return nil, fmt.Errorf("fail decode resp: %w", err)
	}

	return &weather, nil
}
