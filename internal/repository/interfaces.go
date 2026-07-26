package repository

import (
	"context"
	"cryptoPulse/internal/db"
	"cryptoPulse/internal/domain"
)

type Repository struct {
	Symbols Symbols
	Candles Candles
}
type Symbols interface {
	GetSymbols() ([]db.Symbol, error)
}
type Candles interface {
	CreateCandles(ctx context.Context, candles []domain.Candle) error
	GetLastCandleBySymbol(ctx context.Context, symbol string, interval string) (*domain.Candle, error)
}
