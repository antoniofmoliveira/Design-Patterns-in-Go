package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
)

// workers_pipeline.go file
func main() {
	bufferSize := 100
	var dispatcher Dispatcher = NewDispatcher(bufferSize)
	workers := 3
	for i := 0; i < workers; i++ {
		var w WorkerLauncher = &PreffixSuffixWorker{
			prefixS: fmt.Sprintf("WorkerID: %d -> ", i),
			suffixS: " World",
			id:      i,
		}
		dispatcher.LaunchWorker(w)
	}
	requests := 10
	var wg sync.WaitGroup
	wg.Add(requests)
	for i := 0; i < requests; i++ {
		req := NewStringRequest("(Msg_id: %d) -> Hello", i, &wg)
		dispatcher.MakeRequest(req)
	}
	dispatcher.Stop()
	wg.Wait()
}

type Request struct {
	Data    interface{}
	Handler RequestHandler
}

type RequestHandler func(interface{})

func NewStringRequest(s string, id int, wg *sync.WaitGroup) Request {
	return Request{
		Data: "Hello", Handler: func(i interface{}) {
			defer wg.Done()
			s, ok := i.(string)
			if !ok {
				log.Fatal("Invalid casting to string")
			}
			fmt.Println(s)
		},
	}
}

// worker.go file
type WorkerLauncher interface {
	LaunchWorker(in chan Request)
}

type PreffixSuffixWorker struct {
	id      int
	prefixS string
	suffixS string
}

func (w *PreffixSuffixWorker) LaunchWorker(in chan Request) {
	w.prefix(w.append(w.uppercase(in)))
}

func (w *PreffixSuffixWorker) uppercase(in <-chan Request) <-chan Request {
	out := make(chan Request)
	go func() {
		for msg := range in {
			s, ok := msg.Data.(string)
			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Data = strings.ToUpper(s)
			out <- msg
		}
		close(out)
	}()
	return out
}

func (w *PreffixSuffixWorker) append(in <-chan Request) <-chan Request {
	out := make(chan Request)
	go func() {
		for msg := range in {
			uppercaseString, ok := msg.Data.(string)
			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Data = fmt.Sprintf("%s%s", uppercaseString, w.suffixS)
			out <- msg
		}
		close(out)
	}()
	return out
}

func (w *PreffixSuffixWorker) prefix(in <-chan Request) {
	go func() {
		for msg := range in {
			uppercasedStringWithSuffix, ok := msg.Data.(string)
			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Handler(fmt.Sprintf("%s%s", w.prefixS,
				uppercasedStringWithSuffix))
		}
	}()
}

// dispatcher.go file
type Dispatcher interface {
	LaunchWorker(w WorkerLauncher)
	MakeRequest(Request)
	Stop()
}

type dispatcher struct {
	inCh chan Request
}

func (d *dispatcher) LaunchWorker(w WorkerLauncher) {
	w.LaunchWorker(d.inCh)
}
func (d *dispatcher) Stop() {
	close(d.inCh)
}
func (d *dispatcher) MakeRequest(r Request) {
	// d.inCh <- r
	select {
	case d.inCh <- r:
	case <-time.After(time.Second * 5):
		return
	}

}
func NewDispatcher(b int) Dispatcher {
	return &dispatcher{
		inCh: make(chan Request, b),
	}
}
