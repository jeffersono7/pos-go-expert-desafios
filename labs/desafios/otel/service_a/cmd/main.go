package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/jeffersono7/pos-go-expert-desafios/labs/desafios/otel/service_a/internal/handler"
)

func main() {
	r := mux.NewRouter()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	server := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	// boot
	weatherController := handler.WeatherHandler{}

	r.HandleFunc("/weather", weatherController.GetWeather).Methods("POST")
	// end

	log.Println("server listen...")

	<-sig
	ctx, shutdownRelease := context.WithTimeout(context.Background(), time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("http shutdown error: %v", err)
		return
	}

	log.Println("graceful shutdown complete.")
}
