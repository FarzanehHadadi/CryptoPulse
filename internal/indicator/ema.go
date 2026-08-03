package indicator

import "cryptoPulse/internal/domain"

type ema struct {
	name string
}

func (e *ema) Name() string {
	return e.name
}
func (e *ema) Calculate(candles []domain.Candle) (any, error) {
	return nil, nil
}
