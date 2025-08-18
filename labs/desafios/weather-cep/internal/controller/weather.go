package controller

import "net/http"

type WeatherController struct {
}

func NewWeatherController() *WeatherController {
	return &WeatherController{}
}

func (wc *WeatherController) GetWeather(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("i'm working Jeff"))
}
