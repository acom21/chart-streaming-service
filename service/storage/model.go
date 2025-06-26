package storage

import (
	"time"

	"github.com/shopspring/decimal"
)

// OHLC represents a candlestick data point.
type OHLC struct {
	Symbol    string
	StartTime time.Time
	Open      decimal.Decimal
	Close     decimal.Decimal
	High      decimal.Decimal
	Low       decimal.Decimal
	Volume    float64
}
