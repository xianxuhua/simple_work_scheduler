package worker

import (
	"math/rand"
	"strconv"
	"time"
)

type Request struct{}
type ParseResult struct {
	Items    []string
	Requests []Request
}

func CreateWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result := Work(request)
			out <- result
		}
	}()
}

func Work(request Request) ParseResult {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	return ParseResult{Requests: []Request{request}, Items: []string{"golang" + strconv.Itoa(rand.Intn(100))}}
}
