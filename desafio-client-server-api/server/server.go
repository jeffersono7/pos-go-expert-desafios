package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/jeffersono7/pos-go-expert-desafios/desafio-client-server-api/server/domain"
	"github.com/jeffersono7/pos-go-expert-desafios/desafio-client-server-api/server/repository"
)

var (
	quoteRepository repository.QuoteRepository
)

func main() {
	now := time.Now()
	log.Println("Booting server...")

	// db
	DB := initDB()
	defer DB.Close()
	quoteRepository = repository.QuoteRepository{DB: DB}

	// handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", logMiddleware(handleCotacao))

	log.Printf("\nListen on :%s [time_to_start: %0.2f]\n", "8080", time.Since(now).Seconds())

	http.ListenAndServe(":8080", mux)
}

func initDB() *sql.DB {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	DB, err := sql.Open("sqlite3", "./quote.db")
	if err != nil {
		log.Fatalln(err)
	}

	// create tables
	stmt := `
		CREATE TABLE IF NOT EXISTS quotes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			bid TEXT NOT NULL
		)
	`
	_, err = DB.ExecContext(ctx, stmt)
	if err != nil {
		log.Fatalln(err)
	}

	return DB
}

func handleCotacao(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
	defer cancel()

	quote, err := getDollarQuote(ctx)
	if err != nil {
		log.Printf("timeout reached on request to quote: %v", err)
		w.WriteHeader(http.StatusRequestTimeout)
		w.Write([]byte("{\"error\": \"timeout\"}"))
		return
	}

	ctxDb, cancel := context.WithTimeout(ctx, time.Millisecond*10)
	defer cancel()
	err = quoteRepository.Insert(ctxDb, quote)
	if err != nil {
		log.Printf("timeout reached on request to quote: %v", err)
		w.WriteHeader(http.StatusRequestTimeout)
		w.Write([]byte("{\"error\": \"timeout\"}"))
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(quote); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"error\": \"internal server error\"}"))
		return
	}
}

func getDollarQuote(ctx context.Context) (domain.Quote, error) {
	client := http.DefaultClient

	request, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return domain.Quote{}, fmt.Errorf("failed create request: %w", err)
	}
	resp, err := client.Do(request)
	if err != nil {
		return domain.Quote{}, fmt.Errorf("failed do request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return domain.Quote{}, fmt.Errorf("failed read response body: %w", err)
	}
	var responseQuote map[string]map[string]string
	if err := json.Unmarshal(body, &responseQuote); err != nil {
		return domain.Quote{}, fmt.Errorf("failed unmarshal response: %w", err)
	}

	usdBrl, exists := responseQuote["USDBRL"]
	if !exists {
		return domain.Quote{}, fmt.Errorf("unexpected response: %w", err)
	}
	bid, exists := usdBrl["bid"]
	if !exists {
		return domain.Quote{}, fmt.Errorf("unexpected response: %w", err)
	}

	return domain.Quote{Bid: bid}, nil
}

func logMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	}
}
