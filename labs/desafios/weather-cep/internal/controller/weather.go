package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jeffersono7/pos-go-expert-desafios/labs/desafios/weather-cep/internal/controller/dto"
	"github.com/jeffersono7/pos-go-expert-desafios/labs/desafios/weather-cep/internal/service"
)

type WeatherController struct {
	weatherService *service.WeatherService
}

func NewWeatherController(weatherService *service.WeatherService) *WeatherController {
	return &WeatherController{weatherService: weatherService}
}

func (wc *WeatherController) GetWeather(w http.ResponseWriter, r *http.Request) {
	cepParam := chi.URLParam(r, "cep")
	ctx := context.Background()

	weather, err := wc.weatherService.GetWeatherFromCEP(ctx, cepParam)
	if err != nil {
		wc.handleError(w, err)
		return
	}

	response := dto.ToDTO(weather)

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
	}
}

func (wc *WeatherController) handleError(w http.ResponseWriter, err error) {
	message := "internal server error"
	status := http.StatusInternalServerError

	switch err {
	case service.ErrFailGetWeather:
		message = err.Error()
	case service.ErrInvalidCEP:
		message = err.Error()
		status = http.StatusUnprocessableEntity
	case service.ErrNotFoundCEP:
		message = err.Error()
		status = http.StatusNotFound
	}

	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(struct {
		Message string `json:"message"`
	}{Message: message}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
	}
}
