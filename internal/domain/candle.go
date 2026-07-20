package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

type Candle struct {
	ID       int32
	Symbol   string
	Interval string

	Open   decimal.Decimal
	High   decimal.Decimal
	Low    decimal.Decimal
	Close  decimal.Decimal
	Volume decimal.Decimal

	OpenTime  time.Time
	CloseTime time.Time
}
