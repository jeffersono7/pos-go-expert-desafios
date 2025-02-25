package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jeffersono7/pos-go-expert-desafios/desafio-client-server-api/server/domain"
)

type QuoteRepository struct {
	DB *sql.DB
}

func (q QuoteRepository) Insert(ctx context.Context, quote domain.Quote) error {
	tx, err := q.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault, ReadOnly: false})
	if err != nil {
		return fmt.Errorf("failed open tx: %w", err)
	}

	query := `
		INSERT INTO quotes (bid) VALUES (?)
	`

	_, err = tx.ExecContext(ctx, query, quote.Bid)
	if err != nil {
		return fmt.Errorf("failed insert quote: %w", err)
	}

	return nil
}
