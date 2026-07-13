package domain

import (
	"database/sql"
	"log"
	"time"

	dbComms "github.com/werastine/OrderBookSpectator.git/database/dbcommands"
	"github.com/werastine/OrderBookSpectator.git/models"
)

func DataStreamReciever(db *sql.DB, bnPayload chan *models.Payload) {
	batch := models.OrderBookBatch{
		Symbols:  make([]string, 0, 15),
		BidPrice: make([]string, 0, 15),
		AskPrice: make([]string, 0, 15),
	}

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if batch.BatchLen() > 0 {
				if err := dbComms.SaveOrderBook(db, &batch); err != nil {
					log.Println("[ERROR] saving data intxo order book:", err)
					return
				}
				batch.BatchClear()

			}
		case pl, ok := <-bnPayload:
			if !ok {
				if batch.BatchLen() > 0 {
					if err := dbComms.SaveOrderBook(db, &batch); err != nil {
						log.Println("[ERROR] saving data into order book:", err)
						return
					}
				}
				return
			}
			batch.BathcAppend(pl.Symbol, pl.BidPrice, pl.AskPrice)

			if batch.BatchLen() >= 15 {
				if err := dbComms.SaveOrderBook(db, &batch); err != nil {
					log.Println("[ERROR] saving data into order book:", err)
					return
				}
				batch.BatchClear()
				ticker.Reset(5 * time.Second)
			}

		}
	}
}
