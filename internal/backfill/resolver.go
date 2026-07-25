package backfill

import (
	"time"
)

type Resolver struct{}

type Range struct {
	From time.Time
	To   time.Time
}

// missing: 12:01 → 12:14
func (r *Resolver) Resolve(lastCandle time.Time, now time.Time, interval time.Duration) Range {
	return Range{}
	//get first time after lastCandle + interval and last time before now - interval
}
