package ws

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/werastine/OrderBookSpectator.git/database"
	"github.com/werastine/OrderBookSpectator.git/internal/analysis"
)

// wss://stream.binance.com:9443/ws/linkusdt@bookTicker

type Payload struct {
	Symbol      string `json:"s"`
	BidPrice    string `json:"b"`
	BidQuantity string `json:"B"`
	AskPrice    string `json:"a"`
	AskQuantity string `json:"A"`
	Spread      float64
}

func BnGetOrderBook(db *sql.DB) (*Payload, error) {
	url := "wss://stream.binance.com:9443/ws/gramusdt@bookTicker"
	var payload Payload

	conn, resp, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, fmt.Errorf("connecting to Binance default dialer %w", err)
	}

	if resp.StatusCode < 101 && resp.StatusCode > 299 {
		return nil, fmt.Errorf("recieved status code while connecting to Binance default dialer %d", resp.StatusCode)
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
			return nil, fmt.Errorf("reading message %w", err)
		}

		if err = json.Unmarshal(p, &payload); err != nil {
			return nil, fmt.Errorf("unmarshaling bytes into json: %w", err)
		}

		payload.Spread, err = analysis.CountSpread(payload.BidPrice, payload.AskPrice)
		if err != nil {
			return nil, fmt.Errorf("counting spread %w", err)
		}

		if err = database.SaveOrderBook(db, payload.Symbol, payload.BidPrice, payload.AskPrice); err != nil {
			return &payload, fmt.Errorf("saving info into order book %w", err)
		}
	}
}
