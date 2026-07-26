package backfill

import (
	"context"
	"cryptoPulse/internal/domain"
	"cryptoPulse/internal/exchange"
	"cryptoPulse/internal/logger"
	"cryptoPulse/internal/repository"
	"time"
)

type Service struct {
	exchange exchange.Exchange
	repo     repository.Candles
	resolver *Resolver
}

func (s *Service) Run(ctx context.Context, req exchange.CandleRequest) error {
	// get last candle in repo for req.Name
	candle, err := s.repo.GetLastCandleBySymbol(ctx, req.Symbol, req.Interval)
	if err != nil {
		return err
	}
	parsedInterval, err := domain.ParseInterval(req.Interval)
	if err != nil {
		return err
	}
	lastCandleTime := candle.OpenTime.Add(parsedInterval)
	// send it to resolver
	missingRange := s.resolver.Resolve(lastCandleTime, time.Now(), parsedInterval)
	//get missing time and send it to strategy
	logger.Info("missing range:", missingRange)
	return nil
}
