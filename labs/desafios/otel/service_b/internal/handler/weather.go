package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jeffersono7/pos-go-expert-desafios/labs/desafios/otel/service_b/internal/service"
)

type WeatherHandler struct {
	Service service.WeatherService
}

type GetWeatherReqBody struct {
	Cep string `json:"cep" validate:"required,min=8"`
}

func (g GetWeatherReqBody) validate() error {
	validate := validator.New()

	if err := validate.Struct(g); err != nil {
		return errors.New("invalid zipcode")
	}

	return nil
}

func (h WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("invalid zipcode"))
		return
	}

	var getWeatherReqBody GetWeatherReqBody
	if err = json.Unmarshal(body, &getWeatherReqBody); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("invalid zipcode"))
		return
	}

	if err = getWeatherReqBody.validate(); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("invalid zipcode"))
	}

	weather, err := h.Service.GetWeather(ctx, getWeatherReqBody.Cep)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(weather); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
	}
}
