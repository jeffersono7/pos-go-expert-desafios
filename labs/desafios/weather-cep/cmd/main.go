package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jeffersono7/pos-go-expert-desafios/labs/desafios/weather-cep/internal/controller"
	"github.com/jeffersono7/pos-go-expert-desafios/labs/desafios/weather-cep/internal/service"
	"github.com/jeffersono7/pos-go-expert-desafios/labs/desafios/weather-cep/internal/service/impl"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(time.Second * 5))

	// clients
	cepClient := impl.NewCepServiceImpl()
	weatherClient := impl.NewWeatherClientImpl()
	// services
	weatherService := service.NewWeatherService(cepClient, weatherClient)
	// controllers
	wc := controller.NewWeatherController(weatherService)

	r.Get("/cep/{cep}/weather", wc.GetWeather)

	log.Println("Listen on :8080")
	http.ListenAndServe(net.JoinHostPort("", "8080"), r)
}
