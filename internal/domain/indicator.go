package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

type IndicatorResult struct {
	Symbol   string
	Interval string

	Name   string
	Period int

	CandleTime time.Time

	Value decimal.Decimal
}
