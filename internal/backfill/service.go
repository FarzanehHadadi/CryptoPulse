package backfill

import (
	"context"
	"cryptoPulse/internal/domain"
	"cryptoPulse/internal/exchange"
	"cryptoPulse/internal/logger"
	"cryptoPulse/internal/repository"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
)

type Service struct {
	exchange exchange.Exchange
	repo     repository.Candles
	resolver *Resolver
	planner  Planner
}

func NewService(ex exchange.Exchange, repo repository.Candles, maxPerReq int) *Service {
	return &Service{
		exchange: ex,
		repo:     repo,
		resolver: &Resolver{},
		planner:  &ChunkPlanner{MaxCandlesPerRequest: maxPerReq},
	}
}
func (s *Service) Run(ctx context.Context, req exchange.CandleRequest) error {
	// get last candle in repo for req.Name
	candle, err := s.repo.GetLastCandleBySymbol(ctx, req.Symbol, req.Interval)

	parsedInterval, _ := domain.ParseInterval(req.Interval)
	var lastCandleTime time.Time
	switch {
	case err == nil:
		lastCandleTime = candle.OpenTime.Add(parsedInterval)
	case errors.Is(err, pgx.ErrNoRows):
		lastCandleTime = time.Now().Add(-7 * 24 * time.Hour) // configurable later
	default:
		return err
	}
	// send it to resolver
	missingRange := s.resolver.Resolve(lastCandleTime, time.Now(), parsedInterval)
	//get missing time and send it to strategy
	requests := s.planner.Plan(req, missingRange, parsedInterval)

	logger.Info("adding candles ... ")
	for _, req := range requests {
		candles, err := s.exchange.GetCandles(ctx, req)

		if err != nil {
			return err
		}
		if err := s.repo.CreateCandles(ctx, candles); err != nil {
			return err
		}
	}
	logger.Info("candles added successfully ")
	return nil
}
