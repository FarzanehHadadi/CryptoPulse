package scheduler

import (
	"cryptoPulse/internal/config"
	"cryptoPulse/internal/exchange"
)

type Job struct {
	Name          string
	CandleRequest exchange.CandleRequest
}

func JobsFromConfig(cfg config.SchedulerConfig) []Job {
	jobs := make([]Job, len(cfg.Jobs))
	for i, job := range cfg.Jobs {
		jobs[i] = Job{
			Name: job.Name,
			CandleRequest: exchange.CandleRequest{
				Symbol:   job.CandleRequest.Symbol,
				Limit:    job.CandleRequest.Limit,
				Interval: job.CandleRequest.Interval,
			},
		}
	}
	return jobs
}
