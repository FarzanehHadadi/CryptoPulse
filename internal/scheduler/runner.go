package scheduler

import (
	"context"
	"cryptoPulse/internal/collector"
	"cryptoPulse/internal/logger"
	"time"
)

type Runner struct {
	Collector *collector.Service
	timeout   time.Duration
}

func newRunner(s *collector.Service, t time.Duration) *Runner {
	return &Runner{
		Collector: s,
		timeout:   t,
	}

}
func (r *Runner) run(ctx context.Context, job Job) {
	jobCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()
	if err := r.Collector.CollectCandles(jobCtx, job.CandleRequest); err != nil {
		logger.Error("job failed", "job", job.Name, "error", err)
		return
	}
	logger.Info("job finished", "job", job.Name)

}
