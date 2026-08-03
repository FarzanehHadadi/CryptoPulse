package repository

import (
	"context"
	"cryptoPulse/internal/db"
	"cryptoPulse/internal/domain"
)

type Repository struct {
	Symbols    Symbols
	Candles    Candles
	Indicators Indicators
}
type Symbols interface {
	GetSymbols() ([]db.Symbol, error)
}
type Candles interface {
	CreateCandles(ctx context.Context, candles []domain.Candle) error
	GetLastCandleBySymbol(ctx context.Context, symbol string, interval string) (*domain.Candle, error)
	GetCandlesBySymbol(ctx context.Context, symbol string, interval string, count int) ([]domain.Candle, error)
}
type Indicators interface {
	CreateIndicators(ctx context.Context, indicators []domain.IndicatorResult) error
}
