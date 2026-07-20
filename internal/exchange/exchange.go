package exchange

import (
	"context"
	"cryptoPulse/internal/domain"
)

type Exchange interface {
	GetSymbols() ([]string, error)
	GetCandles(ctx context.Context, request CandleRequest) ([]domain.Candle, error)
}

type CandleRequest struct {
	Symbol   string
	Interval string
	Limit    int
}
