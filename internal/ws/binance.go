package ws

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/werastine/OrderBookSpectator.git/internal/analysis"
	"github.com/werastine/OrderBookSpectator.git/models"
)

// wss://stream.binance.com:9443/ws/linkusdt@bookTicker

func BnGetOrderBook(bnStream chan *models.Payload) error {
	url := "wss://stream.binance.com:9443/ws/linkusdt@bookTicker"

	conn, resp, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return fmt.Errorf("connecting to Binance default dialer %w", err)
	}

	if resp.StatusCode < 101 || resp.StatusCode > 299 {
		return fmt.Errorf("recieved status code while connecting to Binance default dialer %d", resp.StatusCode)
	}

	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("[ERROR] Binance: closing connection")
		}
	}()

	log.Println("Websocket is created!")

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			return fmt.Errorf("reading message %w", err)
		}

		var pl models.Payload
		if err = json.Unmarshal(p, &pl); err != nil {
			return fmt.Errorf("unmarshaling bytes into json: %w", err)
		}

		pl.Spread, err = analysis.CountSpread(pl.BidPrice, pl.AskPrice)
		if err != nil {
			return fmt.Errorf("counting spread %w", err)
		}
		bnStream <- &pl
	}
}
