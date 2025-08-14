package service

import "github.com/jeffersono7/pos-go-expert-desafios/labs/desafios/weather-cep/internal/domain"

type WeatherService struct {
	cepService CepService
}

func NewWeatherService(cepService CepService) *WeatherService {
	return &WeatherService{cepService: cepService}
}

func GetWeatherFromCEP(cep string) (domain.Weather, error) {
	return domain.Weather{}, nil
}
