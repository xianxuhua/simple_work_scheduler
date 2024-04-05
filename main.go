package main

import (
	"simple_work_scheduler/engine"
	"simple_work_scheduler/scheduler"
	"simple_work_scheduler/worker"
)

func main() {
	//engine.Engine{
	//	Scheduler:   &scheduler.Scheduler{},
	//	WorkerCount: 100}.Run(worker.Request{})
	engine.QueueEngine{
		QueueScheduler: &scheduler.QueueScheduler{},
		WorkerCount:    100,
	}.Run(worker.Request{})
}
