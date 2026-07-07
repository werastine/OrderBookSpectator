package database

import (
	"database/sql"
	"fmt"
)

func SaveOrderBook(db *sql.DB, symbol string, bidPrice string, askPrice string) error {
	query := `INSERT INTO order_books (symbol, best_bid, best_ask)
			  VALUES ($1, $2, $3);
		`

	_, err := db.Exec(query, symbol, bidPrice, askPrice)
	if err != nil {
		return fmt.Errorf(" inserting into order book %w", err)
	}
	return nil
}
