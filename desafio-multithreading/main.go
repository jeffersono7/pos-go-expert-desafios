package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jeffersono7/pos-go-expert-desafios/desafio-multithreading/internal/brasilapi"
	"github.com/jeffersono7/pos-go-expert-desafios/desafio-multithreading/internal/client"
	"github.com/jeffersono7/pos-go-expert-desafios/desafio-multithreading/internal/viacep"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage go run main.go <cep> <cep>...\n")
		return
	}
	ceps := os.Args[1:]

	cepToConsult := make(chan string, 3)
	cepResponse := make(chan string, 3)

	go worker(cepToConsult, cepResponse)
	go worker(cepToConsult, cepResponse)
	go worker(cepToConsult, cepResponse)

	for _, cep := range ceps {
		cepToConsult <- cep
	}
	close(cepToConsult)

	for range ceps {
		resp := <-cepResponse

		fmt.Println(resp)
	}
}

func worker(cepToConsult <-chan string, cepResponse chan<- string) {
	var client1 client.CepAPI = brasilapi.BrasilAPIClient{}
	var client2 client.CepAPI = viacep.ViaCepAPIClient{}

	for cep := range cepToConsult {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		channel1 := make(chan client.CepResponse)
		channel2 := make(chan client.CepResponse)

		go client1.GetCEP(ctx, cep, channel1)
		go client2.GetCEP(ctx, cep, channel2)

		select {
		case cepData := <-channel1:
			cancel()
			cepResponse <- fmt.Sprintf("CEP: %s -> %v", cep, cepData)
		case cepData := <-channel2:
			cancel()
			cepResponse <- fmt.Sprintf("CEP: %s -> %v", cep, cepData)
		case <-time.After(time.Second):
			cancel()
			cepResponse <- fmt.Sprintf("CEP: %s -> timeout", cep)
		}
	}
}
