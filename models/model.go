package models

type Payload struct {
	Symbol      string `json:"s"`
	BidPrice    string `json:"b"`
	BidQuantity string `json:"B"`
	AskPrice    string `json:"a"`
	AskQuantity string `json:"A"`
	Spread      float64
}
