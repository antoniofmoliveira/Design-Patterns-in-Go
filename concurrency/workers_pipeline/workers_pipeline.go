package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
)

// main is an example of a dispatcher with multiple workers and multiple requests.
// The dispatcher is created with a buffer size of 100, and 3 workers are launched.
// Each worker is a PreffixSuffixWorker, which appends a prefix to the incoming
// message and appends a suffix to the result before sending it back to the
// dispatcher.
// The dispatcher is then asked to make a total of 10 requests, each of which is
// a NewStringRequest with the format "(Msg_id: %d) -> Hello". The dispatcher
// will execute each request in order, and the main goroutine will wait for all
// the requests to complete before exiting.
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

// This class definition defines a `Request` struct with two fields: `Data` and `Handler`.
// Here's what each field does:
// * `Data`: holds the data associated with the request, which can be of any type (`interface{}`)
// * `Handler`: a function that will be executed to handle the request, of type `RequestHandler` (which is a `func(interface{})`)
// Note that this is not a class definition in the classical sense, but rather a struct definition in Go.
type Request struct {
	Data    interface{}
	Handler RequestHandler
}

// This is a type alias for a function that takes an interface{} and returns nothing.
type RequestHandler func(interface{})

// NewStringRequest returns a new Request with Data set to "Hello" and Handler set to a function
// that prints the Data as a string and calls Done on the given WaitGroup when done.
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

// The `WorkerLauncher` interface defines a single method `LaunchWorker` that takes a channel of `Request` objects as input. This interface is used to define the behavior of a worker that can be launched by a dispatcher. The `LaunchWorker` method is responsible for starting the worker and passing in the input channel that the worker will read from.
type WorkerLauncher interface {
	LaunchWorker(in chan Request)
}

// This is not a class definition in the classical sense, but rather a struct definition in Go. Here's what each field does:
// * `id`: an integer identifier for the worker
// * `prefixS`: a string that will be prefixed to the data processed by the worker
// * `suffixS`: a string that will be suffixed to the data processed by the worker
// Note that this struct has methods associated with it, which are defined elsewhere in the code. These methods are:
// * `LaunchWorker(in chan Request)`: launches the worker and passes in the input channel that the worker will read from.
// * `uppercase(in <-chan Request) <-chan Request`: takes a channel of requests, uppercases the data in each request, and returns a new channel with the uppercased requests.
// * `append(in <-chan Request) <-chan Request`: takes a channel of requests, appends the suffix to the data in each request, and returns a new channel with the appended requests.
// * `prefix(in <-chan Request)`: takes a channel of requests, prefixes the prefix to the data in each request, and handles the requests.
type PreffixSuffixWorker struct {
	id      int
	prefixS string
	suffixS string
}

// LaunchWorker calls the LaunchWorker function of the given WorkerLauncher,
// passing in the input channel that the worker will read from.
func (w *PreffixSuffixWorker) LaunchWorker(in chan Request) {
	w.prefix(w.append(w.uppercase(in)))
}

// uppercase takes a channel of requests, uppercases the data in each request, and returns a new channel with the uppercased requests.
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

// append takes a channel of requests, appends the suffix to the data in each request,
// and returns a new channel with the appended requests.
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

// prefix takes a channel of requests, prefixes the prefix to the data in each request,
// and handles the requests.
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

// The `Dispatcher` interface defines a contract for a dispatcher that can manage workers. Here's a brief explanation of each method:
// * `LaunchWorker(w WorkerLauncher)`: Launches a worker that implements the `WorkerLauncher` interface, passing control to the worker.
// * `MakeRequest(Request)`: Sends a request to the dispatcher, which will be processed by the launched workers.
// * `Stop()`: Stops the dispatcher, likely by closing its input channel and causing launched workers to exit.
type Dispatcher interface {
	LaunchWorker(w WorkerLauncher)
	MakeRequest(Request)
	Stop()
}

// This is not a class definition, but rather a struct definition in Go. Here's a succinct explanation of what it does:
// The `dispatcher` struct represents a dispatcher that manages requests. It has a single field `inCh`, which is a channel of type `Request`.
type dispatcher struct {
	inCh chan Request
}

// LaunchWorker calls the LaunchWorker function of the given WorkerLauncher,
// passing in the dispatcher's input channel.
func (d *dispatcher) LaunchWorker(w WorkerLauncher) {
	w.LaunchWorker(d.inCh)
}

// Stop the dispatcher by closing its input channel. This will cause all
// launched workers to eventually exit.
func (d *dispatcher) Stop() {
	close(d.inCh)
}

// MakeRequest sends a Request to the dispatcher's input channel. If the
// channel is full, it will wait for 5 seconds and then return.
func (d *dispatcher) MakeRequest(r Request) {
	// d.inCh <- r
	select {
	case d.inCh <- r:
	case <-time.After(time.Second * 5):
		return
	}

}

// NewDispatcher returns a new Dispatcher with a channel of the given buffer size.
func NewDispatcher(b int) Dispatcher {
	return &dispatcher{
		inCh: make(chan Request, b),
	}
}
