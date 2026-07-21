package scheduler

import (
	"context"
	"sync"
)

type WorkerPool struct {
	jobs    chan Job
	runner  *Runner
	workers int
	wg      sync.WaitGroup
}

func newWorkerPool(worker int, r *Runner) *WorkerPool {
	if worker < 1 {
		worker = 1
	}
	return &WorkerPool{
		workers: worker,
		runner:  r,
		jobs:    make(chan Job, worker*2),
	}

}

func (p *WorkerPool) start(ctx context.Context) {
	for i := 0; i < p.workers; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case job, ok := <-p.jobs:
					if !ok {
						return
					}
					p.runner.run(ctx, job)
				}
			}
		}()
	}
}
func (p *WorkerPool) submit(ctx context.Context, job Job) {
	select {
	case <-ctx.Done():
		return
	case p.jobs <- job:
	}
}
func (p *WorkerPool) stop() {
	close(p.jobs)
	p.wg.Wait()
}
