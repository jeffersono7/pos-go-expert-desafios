package service

import (
	"errors"
	"regexp"

	"github.com/jeffersono7/pos-go-expert-desafios/labs/desafios/weather-cep/internal/domain"
)

var (
	ErrInvalidCEP  = errors.New("invalid zipcode")
	ErrNotFoundCEP = errors.New("can not find zipcode")
	cepRegex       *regexp.Regexp
)

type WeatherService struct {
	cepClient     CepClient
	weatherClient WeatherClient
}

func init() {
	reg, err := regexp.Compile("^[0-9]{8}$")
	if err != nil {
		panic(err)
	}
	cepRegex = reg
}

func NewWeatherService(cepClient CepClient, weatherClient WeatherClient) *WeatherService {
	return &WeatherService{cepClient: cepClient, weatherClient: weatherClient}
}

func (ws *WeatherService) GetWeatherFromCEP(cep string) (domain.Weather, error) {
	if !cepRegex.MatchString(cep) {
		return domain.Weather{}, ErrInvalidCEP
	}

	return domain.Weather{}, nil
}
