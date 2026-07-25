package backfill

import "cryptoPulse/internal/exchange"

type Strategy interface {
	Split(r Range) []exchange.CandleRequest
}
