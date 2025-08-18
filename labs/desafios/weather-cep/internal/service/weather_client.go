package service

import "context"

type WeatherResp struct {
	Current struct {
		TempC float32 `json:"temp_c"`
	} `json:"current"`
}

type WeatherClient interface {
	GetTemp(ctx context.Context, neighborhood string) (WeatherResp, error)
}
