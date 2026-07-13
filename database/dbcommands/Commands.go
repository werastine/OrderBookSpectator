package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/werastine/OrderBookSpectator.git/models"
)

func SaveOrderBook(db *sql.DB, batch *models.OrderBookBatch) error {
	if batch.BatchLen() == 0 {
		return fmt.Errorf("batch is 0")
	}

	query := `INSERT INTO order_books (symbol, best_bid, best_ask)
			  SELECT * FROM unnest($1::text[], $2::numeric[], $3::numeric[])`

	if _, err := db.Exec(query, batch.Symbols, batch.BidPrice, batch.AskPrice); err != nil {
		return fmt.Errorf("sending sql query: %w", err)
	}
	log.Println("Sent batch of data into Database")
	return nil
}
