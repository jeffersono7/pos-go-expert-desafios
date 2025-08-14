package service

import "github.com/jeffersono7/pos-go-expert-desafios/labs/desafios/weather-cep/internal/domain"

type WeatherService struct {
	cepClient     CepClient
	weatherClient WeatherClient
}

func NewWeatherService(cepClient CepClient, weatherClient WeatherClient) *WeatherService {
	return &WeatherService{cepClient: cepClient, weatherClient: weatherClient}
}

func GetWeatherFromCEP(cep string) (domain.Weather, error) {
	return domain.Weather{}, nil
}
