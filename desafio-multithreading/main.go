package main

import (
	"fmt"
	"os"
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
	for cep := range cepToConsult {
		cepResponse <- cep
	}
}
