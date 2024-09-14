package main

import (
	"fmt"
	"regexp"
	"sync"
	"testing"
)

// This is a test function in Go, `Test_Dispatcher`, which tests the functionality of a `Dispatcher` by launching 3 workers and making 10 requests to the dispatcher. The test checks if the requests are processed correctly by the workers and if the responses match a specific pattern.
func Test_Dispatcher(t *testing.T) {
	// pasted code from main function
	bufferSize := 100
	var dispatcher Dispatcher = NewDispatcher(bufferSize)
	workers := 3
	for i := 0; i < workers; i++ {
		var w WorkerLauncher = &PreffixSuffixWorker{prefixS: fmt.Sprintf("WorkerID: %d -> ", i), suffixS: " World", id: i}
		dispatcher.LaunchWorker(w)
	}
	// Simulate Requests
	requests := 10
	var wg sync.WaitGroup
	wg.Add(requests)
	for i := 0; i < requests; i++ {
		req := Request{
			Data: fmt.Sprintf("(Msg_id: %d) -> Hello", i),
			Handler: func(i interface{}) {
				s, ok := i.(string)
				defer wg.Done()
				if !ok {
					t.Fail()
				}
				ok, err := regexp.Match(`WorkerID\: \d* -\> \(MSG_ID: \d*\) -> [A-Z]*\sWorld`, []byte(s))
				if !ok || err != nil {
					t.Fail()
				}
			},
		}
		dispatcher.MakeRequest(req)
	}
}
