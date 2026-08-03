package indicator

import (
	"fmt"

	"cryptoPulse/internal/domain"

	"github.com/shopspring/decimal"
)

type SMA struct {
	period int
}

func NewSMA(period int) *SMA {
	return &SMA{period: period}
}

func (s *SMA) Name() string {
	return "SMA"
}

func (s *SMA) Calculate(candles []domain.Candle) ([]domain.IndicatorResult, error) {
	if s.period <= 0 {
		return []domain.IndicatorResult{}, fmt.Errorf("sma: period must be > 0")
	}
	if len(candles) < s.period {
		return []domain.IndicatorResult{}, fmt.Errorf("sma: need at least %d candles, got %d", s.period, len(candles))
	}

	// first window
	sum := decimal.Zero
	for i := 0; i < s.period; i++ {
		sum = sum.Add(candles[i].Close)
	}
	period := decimal.NewFromInt(int64(s.period))
	results := make([]domain.IndicatorResult, 0, len(candles)-s.period+1)
	results = append(results, domain.IndicatorResult{
		Symbol:     candles[s.period-1].Symbol,
		Interval:   candles[s.period-1].Interval,
		Name:       s.Name(),
		Period:     s.period,
		CandleTime: candles[s.period-1].OpenTime,
		Value:      sum.Div(period),
	})
	// slide window
	for i := s.period; i < len(candles); i++ {
		sum = sum.Add(candles[i].Close).Sub(candles[i-s.period].Close)
		sma := sum.Div(period)
		results = append(results, domain.IndicatorResult{Symbol: candles[i].Symbol,
			Interval: candles[i].Interval, Name: s.Name(), Period: s.period,
			CandleTime: candles[i].OpenTime, Value: sma})
	}

	return results, nil
}
func (s *SMA) RequiredCandles() int {
	return s.period
}
