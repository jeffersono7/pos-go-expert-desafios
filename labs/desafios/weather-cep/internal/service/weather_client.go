package service

type WeatherResp struct {
}

type WeatherClient interface {
	GetTemp(neighborhood string) (WeatherResp, error)
}
