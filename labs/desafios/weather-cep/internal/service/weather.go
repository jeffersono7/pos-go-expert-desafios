package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"regexp"

	"github.com/jeffersono7/pos-go-expert-desafios/labs/desafios/weather-cep/internal/domain"
)

var (
	ErrInvalidCEP     = errors.New("invalid zipcode")
	ErrNotFoundCEP    = errors.New("can not find zipcode")
	ErrFailGetWeather = errors.New("fail get weather")
	cepRegex          *regexp.Regexp
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

func (ws *WeatherService) GetWeatherFromCEP(ctx context.Context, cep string) (domain.Weather, error) {
	if !cepRegex.MatchString(cep) {
		return domain.Weather{}, ErrInvalidCEP
	}

	cepResp, err := ws.cepClient.GetCEP(ctx, cep)
	if err != nil {
		log.Println(fmt.Errorf("fail get cep: %w", err))
		return domain.Weather{}, ErrNotFoundCEP
	}

	weatherResp, err := ws.weatherClient.GetTemp(ctx, fmt.Sprintf("%s %s", cepResp.Localidade, cepResp.Estado))
	if err != nil {
		log.Println(fmt.Errorf("fail get weather: %w", err))
		return domain.Weather{}, ErrFailGetWeather
	}

	return domain.NewWeather(weatherResp.Current.TempC), nil
}
