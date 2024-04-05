package main

import (
	"queue_scheduler/engine"
	"queue_scheduler/scheduler"
	"queue_scheduler/worker"
)

func main() {
	engine.Engine{
		Scheduler:   &scheduler.Scheduler{},
		WorkerCount: 100}.Run(worker.Request{})
}
