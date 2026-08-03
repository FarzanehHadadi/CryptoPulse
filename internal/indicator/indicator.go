package indicator

import "cryptoPulse/internal/domain"

type Indicator interface {
	Name() string
	Calculate(candles []domain.Candle) ([]domain.IndicatorResult, error)
	RequiredCandles() int
}
