package barrier

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

var timeoutMilliseconds int = 5000

// This is not a class definition, but a struct definition in Go.
// The `barrierResp` struct has two fields:
// * `Err`: an error type that holds any error that occurred
// * `Resp`: a string type that holds the response
// There are no methods defined in this struct.
type barrierResp struct {
	Err  error
	Resp string
}

// barrier makes a GET request to each of the given endpoints and
// prints their responses. If any of the requests fail, it prints an
// error message instead of the response. If all requests succeed, it
// prints their responses in order.
func barrier(endpoints ...string) {
	requestNumber := len(endpoints)
	in := make(chan barrierResp, requestNumber)
	defer close(in)
	responses := make([]barrierResp, requestNumber)
	for _, endpoint := range endpoints {
		go makeRequest(in, endpoint)
	}
	var hasError bool
	for i := 0; i < requestNumber; i++ {
		resp := <-in
		if resp.Err != nil {
			fmt.Println("ERROR: ", resp.Err)
			hasError = true
		}
		responses[i] = resp
	}
	if !hasError {
		for _, resp := range responses {
			fmt.Println(resp.Resp)
		}
	}
}

// makeRequest makes a GET request to the given URL and writes the response
// to the output channel. If the request fails, it writes an error to the
// output channel.
func makeRequest(out chan<- barrierResp, url string) {
	res := barrierResp{}
	client := http.Client{
		Timeout: time.Duration(time.Duration(timeoutMilliseconds) *
			time.Millisecond),
	}
	resp, err := client.Get(url)
	if err != nil {
		res.Err = err
		out <- res
		return
	}
	byt, err := io.ReadAll(resp.Body)
	if err != nil {
		res.Err = err
		out <- res
		return
	}
	res.Resp = string(byt)
	out <- res
}
