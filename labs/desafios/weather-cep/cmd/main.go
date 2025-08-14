package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jeffersono7/pos-go-expert-desafios/labs/desafios/weather-cep/internal/controller"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(time.Second * 5))

	// controllers
	wc := controller.NewWeatherController()

	r.Get("/cep/{cep}/weather", wc.GetWeather)

	log.Println("Listen on :8080")
	http.ListenAndServe(net.JoinHostPort("", "8080"), r)
}
