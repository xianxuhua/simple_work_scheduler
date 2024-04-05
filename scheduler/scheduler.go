package scheduler

import "queue_scheduler/worker"

type Scheduler struct {
	workerChan chan worker.Request
}

func (s *Scheduler) Submit(request worker.Request) {
	go func() {
		s.workerChan <- request
	}()
}

func (s *Scheduler) ConfigWorkChan(in chan worker.Request) {
	s.workerChan = in
}
