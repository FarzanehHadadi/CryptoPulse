package scheduler

import (
	"context"
	"cryptoPulse/internal/collector"
	"cryptoPulse/internal/logger"
	"time"
)

type Scheduler struct {
	jobs      []Job
	collector *collector.Service
	workers   int
	interval  time.Duration
	timeout   time.Duration

	cancel context.CancelFunc
	pool   *WorkerPool
}

func NewScheduler(c *collector.Service, jobs []Job, workers int) *Scheduler {
	return &Scheduler{
		collector: c,
		jobs:      jobs,
		workers:   workers,
		interval:  time.Minute, // match 1m candles for V1
		timeout:   20 * time.Second,
	}
}

func (s *Scheduler) Start(ctx context.Context) error {
	ctx, s.cancel = context.WithCancel(ctx)

	r := newRunner(s.collector, s.timeout)
	s.pool = newWorkerPool(s.workers, r)
	s.pool.start(ctx)

	// run once immediately, then on ticker
	s.dispatch(ctx)

	ticker := time.NewTicker(s.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			s.pool.stop()
			logger.Info("scheduler stopped")
			return ctx.Err()
		case <-ticker.C:
			s.dispatch(ctx)
		}
	}
}

func (s *Scheduler) dispatch(ctx context.Context) {
	for _, job := range s.jobs {
		s.pool.submit(ctx, job)
	}
}

func (s *Scheduler) Stop() {
	if s.cancel != nil {
		s.cancel()
	}
}
