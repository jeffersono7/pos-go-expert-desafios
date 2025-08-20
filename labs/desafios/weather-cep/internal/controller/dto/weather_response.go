package dto

import "github.com/jeffersono7/pos-go-expert-desafios/labs/desafios/weather-cep/internal/domain"

type WeatherResponseDTO struct {
	TempC float32 `json:"temp_C"`
	TempF float32 `json:"temp_F"`
	TempK float32 `json:"temp_K"`
}

func ToDTO(weather domain.Weather) WeatherResponseDTO {
	return WeatherResponseDTO{
		TempC: weather.TempC,
		TempF: weather.TempF,
		TempK: weather.TempK,
	}
}
