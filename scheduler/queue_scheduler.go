package scheduler

import "simple_work_scheduler/worker"

type QueueScheduler struct {
	requestChan chan worker.Request
	workerChan  chan chan worker.Request
}

func (s *QueueScheduler) Submit(request worker.Request) {
	s.requestChan <- request
}

func (s *QueueScheduler) WorkerReady(w chan worker.Request) {
	s.workerChan <- w
}

func (s *QueueScheduler) Run() {
	s.workerChan = make(chan chan worker.Request)
	s.requestChan = make(chan worker.Request)

	go func() {
		var workerQueue []chan worker.Request
		var requestQueue []worker.Request

		for {
			var activeWorker chan worker.Request
			var activeRequest worker.Request

			if len(workerQueue) != 0 && len(requestQueue) != 0 {
				activeWorker = workerQueue[0]
				activeRequest = requestQueue[0]
			}

			select {
			case w := <-s.workerChan:
				workerQueue = append(workerQueue, w)
			case r := <-s.requestChan:
				requestQueue = append(requestQueue, r)
			case activeWorker <- activeRequest:
				workerQueue = workerQueue[1:]
				requestQueue = requestQueue[1:]
			}
		}
	}()
}
