package engine

import (
	"fmt"
	"simple_work_scheduler/worker"
)

type Scheduler interface {
	Submit(worker.Request)
	ConfigWorkChan(chan worker.Request)
}

type Engine struct {
	Scheduler
	WorkerCount int
}

func (e Engine) Run(seeds ...worker.Request) {
	for _, request := range seeds {
		e.Scheduler.Submit(request)
	}

	in := make(chan worker.Request)
	e.Scheduler.ConfigWorkChan(in)
	out := make(chan worker.ParseResult)

	for i := 0; i < e.WorkerCount; i++ {
		worker.CreateWorker(in, out)
	}

	i := 0
	for {
		result := <-out
		fmt.Println(i, result.Items)
		i++
		for _, req := range result.Requests {
			e.Scheduler.Submit(req)
		}
	}
}
