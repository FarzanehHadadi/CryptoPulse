package scheduler

import (
	"context"
	"cryptoPulse/internal/collector"
)

type scheduler struct {
	Jobs      []Job
	collector *collector.Service
	workers   int
}

func NewScheduler(collector *collector.Service, workers int) *scheduler {
	return &scheduler{
		collector: collector,
		workers:   workers,
	}
}

func (s *scheduler) Start(ctx context.Context) error {
	return nil
}
func (s *scheduler) Stop() {}

func (s *scheduler) dispatch() {}
