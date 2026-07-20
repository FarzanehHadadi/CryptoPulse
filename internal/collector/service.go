package collector

import (
	"context"
	"cryptoPulse/internal/exchange"
	"cryptoPulse/internal/logger"
)

func (s *Service) CollectCandles(ctx context.Context, req exchange.CandleRequest) error {
	candles, err := s.exchange.GetCandles(ctx, req)
	if err != nil {
		return err
	}
	err = s.repo.Candles.CreateCandles(ctx, candles)
	if err != nil {
		logger.Error("Failed to save candles to database", "error", err)
		return err
	}
	logger.Info("Candles saved successfully")
	return nil
}
