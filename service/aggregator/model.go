package aggregator

import (
	"time"

	"github.com/shopspring/decimal"
)

type TradeData struct {
	Symbol    string
	Price     decimal.Decimal
	Quantity  float64
	Timestamp time.Time
	TradeID   int64
}

type Event struct {
	Trade *TradeData
	Err   error
}
