package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/jeffersono7/pos-go-expert-desafios/labs/desafios/otel/service_b/config"
	"github.com/jeffersono7/pos-go-expert-desafios/labs/desafios/otel/service_b/internal/handler"
	"github.com/jeffersono7/pos-go-expert-desafios/labs/desafios/otel/service_b/internal/service"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

var port = "8080"

func init() {
	portEnv := os.Getenv("PORT")
	if portEnv == "" {
		return
	}
	port = portEnv
}

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// otel
	otelShutdown, err := config.SetupOTelSDK(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		err = otelShutdown(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
	}()
	// end

	httpHandler, handleFunc := newOtelHTTPHandler()
	server := http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      httpHandler,
	}

	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	// boot
	httpClient := http.Client{}

	weatherService := service.WeatherService{Client: httpClient}

	weatherController := handler.WeatherHandler{Service: weatherService}

	handleFunc("/weather", weatherController.GetWeather, "POST")
	// end

	log.Printf("server listen on :%s", port)
	<-sig
	ctx, shutdownRelease := context.WithTimeout(context.Background(), time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("http shutdown error: %v", err)
		return
	}

	defer log.Println("graceful shutdown complete.")
}

func newOtelHTTPHandler() (http.Handler, func(string, func(http.ResponseWriter, *http.Request), ...string)) {
	r := mux.NewRouter()

	handleFunc := func(pattern string, handlerFunc func(http.ResponseWriter, *http.Request), methods ...string) {
		handler := otelhttp.WithRouteTag(pattern, http.HandlerFunc(handlerFunc))
		r.Handle(pattern, handler).Methods(methods...)
	}

	handler := otelhttp.NewHandler(r, "/")

	return handler, handleFunc
}
