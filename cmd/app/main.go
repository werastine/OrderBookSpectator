package main

import (
	"log"
	"sync"

	"github.com/werastine/OrderBookSpectator.git/database"
	"github.com/werastine/OrderBookSpectator.git/domain"
	"github.com/werastine/OrderBookSpectator.git/internal/ws"
	"github.com/werastine/OrderBookSpectator.git/models"
)

func main() {
	var wg = sync.WaitGroup{}
	chn := make(chan *models.Payload, 15)

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

	wg.Add(1)
	go func() {
		defer close(chn)
		defer wg.Done()
		err = ws.BnGetOrderBook(chn)
		if err != nil {
			log.Println("[ERROR] Binance:", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		domain.DataStreamReciever(db, chn)
	}()

	wg.Wait()
}

// проверить правки клода
// доработать обработчики ошибок в логике, скюл запросе
// мб улучшать архитектуру!
