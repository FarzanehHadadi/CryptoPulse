package indicator

import (
	"context"
	"cryptoPulse/internal/domain"
	"cryptoPulse/internal/logger"
	"cryptoPulse/internal/repository"
	"errors"
)

type Service struct {
	candleRepo    repository.Candles
	indicatorRepo repository.Indicators
	indicators    []Indicator
}

func NewService(candleRepo repository.Candles, indicatorRepo repository.Indicators) *Service {
	return &Service{
		candleRepo:    candleRepo,
		indicatorRepo: indicatorRepo,
	}
}
func (s *Service) Register(i Indicator) {
	if i == nil {
		return
	}
	s.indicators = append(s.indicators, i)
}
func (s *Service) Process(
	ctx context.Context,
	symbol string,
	interval string,
) error {
	if len(s.indicators) == 0 {
		return errors.New("no indicators registered")
	}
	candles, err := s.candleRepo.GetCandlesBySymbol(ctx, symbol, interval, s.requiredCandles())
	if err != nil {
		return err
	}
	var results []domain.IndicatorResult
	for _, indicator := range s.indicators {
		res, err := indicator.Calculate(candles)
		if err != nil {
			logger.Warn(
				"indicator calculation failed",
				"indicator", indicator.Name(),
				"symbol", symbol,
				"interval", interval,
				"error", err,
			)
			continue
		}
		results = append(results, res...)
	}
	if len(results) == 0 {
		logger.Info(
			"no indicator results generated",
			"symbol", symbol,
			"interval", interval,
		)
		return nil
	}
	if err := s.indicatorRepo.CreateIndicators(ctx, results); err != nil {
		return err
	}
	return nil
}
func (s *Service) requiredCandles() int {
	max := 0
	for _, i := range s.indicators {
		if i.RequiredCandles() > max {
			max = i.RequiredCandles()
		}
	}
	return max
}
