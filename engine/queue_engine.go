package engine

import (
	"fmt"
	"simple_work_scheduler/worker"
)

type QueueScheduler interface {
	Submit(worker.Request)
	WorkerReady(chan worker.Request)
	Run()
}

type QueueEngine struct {
	QueueScheduler
	WorkerCount int
}

func (e QueueEngine) Run(seeds ...worker.Request) {
	e.QueueScheduler.Run()
	out := make(chan worker.ParseResult)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(out, e.QueueScheduler)
	}

	for _, request := range seeds {
		e.QueueScheduler.Submit(request)
	}

	i := 0
	for {
		result := <-out
		fmt.Println(i, result.Items)
		i++
		for _, req := range result.Requests {
			e.QueueScheduler.Submit(req)
		}
	}
}

func createWorker(out chan worker.ParseResult, scheduler QueueScheduler) {
	in := make(chan worker.Request)

	go func() {
		for {
			scheduler.WorkerReady(in)
			request := <-in
			result := worker.Work(request)
			out <- result
		}
	}()
}
