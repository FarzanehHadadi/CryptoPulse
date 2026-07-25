package backfill

import (
	"context"
	"cryptoPulse/internal/exchange"
	"cryptoPulse/internal/repository"
)

type Service struct {
	exchange exchange.Exchange
	repo     repository.Candles
	resolver *Resolver
}

func (s *Service) Run(ctx context.Context, req exchange.CandleRequest) error {
	// get last candle in repo for req.Name
	// send it to resolver
	//get missing time and send it to strategy
	return nil
}
