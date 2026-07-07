package main

import (
	"log"

	"github.com/werastine/OrderBookSpectator.git/database"
	"github.com/werastine/OrderBookSpectator.git/internal/ws"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Println("[ERROR] Postgres:", err)
		return
	}
	defer func() {
		log.Println("Connection is closed!")
		if err := db.Close(); err != nil {
			log.Println("[ERROR] Postgres: closing connection", err)
		}
	}()

	_, err = ws.BnGetOrderBook(db)
	if err != nil {
		log.Println("[ERROR] Binance:", err)
	}
}
