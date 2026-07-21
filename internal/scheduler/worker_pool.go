package scheduler

import "sync"

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

func start()  {}
func stop()   {}
func submit() {}
