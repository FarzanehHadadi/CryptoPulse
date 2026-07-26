package backfill

import (
	"time"
)

type Resolver struct{}

type Range struct {
	From time.Time
	To   time.Time
}

func (r *Range) isEmpty() bool {
	return !r.From.Before(r.To)
}

// missing: 12:01 → 12:14
func (r *Resolver) Resolve(fromTime time.Time, now time.Time, interval time.Duration) Range {
	//get first time after lastCandle + interval and last time before now - interval
	toTime := now.Add(-interval)
	if fromTime.After(toTime) {
		return Range{From: fromTime, To: fromTime} // empty
	}
	return Range{
		From: fromTime,
		To:   toTime,
	}
}
