package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/jeffersono7/pos-go-expert-desafios/labs/desafios/weather-cep/internal/config"
	"github.com/jeffersono7/pos-go-expert-desafios/labs/desafios/weather-cep/internal/service"
)

type WeatherClientImpl struct {
}

func NewWeatherClientImpl() *WeatherClientImpl {
	return &WeatherClientImpl{}
}

func (wc *WeatherClientImpl) GetTemp(ctx context.Context, neighborhood string) (service.WeatherResp, error) {
	neighborhood = url.QueryEscape(neighborhood)
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?q=%s&key=%s", neighborhood, config.ApiKey)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return service.WeatherResp{}, fmt.Errorf("fail make req: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return service.WeatherResp{}, fmt.Errorf("fail do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println(req)
		return service.WeatherResp{}, fmt.Errorf("request returns an error with status: %s", resp.Status)
	}

	var weatherResp service.WeatherResp
	if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
		return service.WeatherResp{}, fmt.Errorf("fail unmarshal response: %w", err)
	}

	return weatherResp, nil
}
