package service

type WeatherResp struct {
	Current struct {
		TempC float32 `json:"temp_c"`
	} `json:"current"`
}

type WeatherClient interface {
	GetTemp(neighborhood string) (WeatherResp, error)
}
