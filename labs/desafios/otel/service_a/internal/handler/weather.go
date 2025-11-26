package handler

import (
	"encoding/json"
	"io"
	"net/http"
)

type WeatherHandler struct {
}

type GetWeatherReqBody struct {
	Cep string `json:"cep"`
}

func (h WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
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

	// validate body

	w.WriteHeader(http.StatusOK)
}
