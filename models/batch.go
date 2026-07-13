package models

type OrderBookBatch struct {
	Symbols  []string
	BidPrice []string
	AskPrice []string
}

func (b *OrderBookBatch) BathcAppend(symb string, bidP string, askP string) {
	b.Symbols = append(b.Symbols, symb)
	b.BidPrice = append(b.BidPrice, bidP)
	b.AskPrice = append(b.AskPrice, askP)
}

func (b *OrderBookBatch) BatchClear() {
	b.Symbols = b.Symbols[:0]
	b.BidPrice = b.BidPrice[:0]
	b.AskPrice = b.AskPrice[:0]
}

func (b *OrderBookBatch) BatchLen() int {
	return len(b.Symbols)
}
