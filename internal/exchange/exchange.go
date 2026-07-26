package exchange

import (
	"context"
	"cryptoPulse/internal/domain"
	"time"
)

type Capabilities struct {
	MaxCandleLimit int
}
type Exchange interface {
	// GetSymbols() ([]string, error)
	GetCandles(ctx context.Context, request CandleRequest) ([]domain.Candle, error)
	Capabilities() Capabilities
}

type CandleRequest struct {
	Symbol    string
	Interval  string
	Limit     int
	StartTime *time.Time
	EndTime   *time.Time
}
