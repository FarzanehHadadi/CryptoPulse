package backfill

import (
	"cryptoPulse/internal/exchange"
	"time"
)

type Planner interface {
	Plan(
		req exchange.CandleRequest,
		r Range,
		interval time.Duration,
	) []exchange.CandleRequest
}
type ChunkPlanner struct {
	MaxCandlesPerRequest int
}

func (cp *ChunkPlanner) Plan(req exchange.CandleRequest,
	r Range, interval time.Duration,
) []exchange.CandleRequest {
	if r.isEmpty() {
		return nil
	}
	limit := cp.MaxCandlesPerRequest
	if limit <= 0 {
		limit = req.Limit
	}
	if limit <= 0 {
		limit = 1000
	}
	var requests []exchange.CandleRequest

	cursor := r.From

	for !cursor.After(r.To) {

		end := cursor.Add(time.Duration(limit-1) * interval)
		if end.After(r.To) {
			end = r.To
		}
		start := cursor
		finish := end
		requests = append(requests, exchange.CandleRequest{
			Symbol:    req.Symbol,
			Interval:  req.Interval,
			Limit:     limit,
			StartTime: &start,
			EndTime:   &finish,
		})
		cursor = end.Add(interval)

	}
	return requests
}
