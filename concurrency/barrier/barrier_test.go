package barrier

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// captureBarrierOutput captures the output of the barrier function by
// redirecting os.Stdout to a pipe and copying from it. It's meant to be used
// in tests.
func captureBarrierOutput(endpoints ...string) string {
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	out := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	barrier(endpoints...)
	writer.Close()
	temp := <-out
	return temp
}

// This is a Go test function named `TestBarrier`. It contains three sub-tests:
// 1. "Correct endpoints": Tests that the `barrier` function correctly handles valid endpoints.
// 2. "One endpoint incorrect": Tests that the `barrier` function correctly handles a malformed URL.
// 3. "Very short timeout": Tests that the `barrier` function correctly handles a very short timeout.
// Each sub-test calls the `captureBarrierOutput` function with different endpoints and checks the output for expected strings. If the output does not contain the expected strings, the test fails.
func TestBarrier(t *testing.T) {
	t.Run("Correct endpoints", func(t *testing.T) {
		endpoints := []string{"http://httpbin.org/headers",
			"http://httpbin.org/user-agent"}
		result := captureBarrierOutput(endpoints...)
		if !(strings.Contains(result, "Accept-Encoding") || strings.Contains(result, "user-agent")) {
			t.Fail()
		}
		t.Log(result)
	})

	t.Run("One endpoint incorrect", func(t *testing.T) {
		endpoints := []string{"http://malformed-url",
			"http://httpbin.org/user-agent"}
		result := captureBarrierOutput(endpoints...)
		if !strings.Contains(result, "ERROR") {
			t.Fail()
		}
		t.Log(result)
	})

	t.Run("Very short timeout", func(t *testing.T) {
		endpoints := []string{"http://httpbin.org/headers",
			"http://httpbin.org/user-agent"}
		timeoutMilliseconds = 1
		result := captureBarrierOutput(endpoints...)
		if !strings.Contains(result, "Timeout") {
			t.Fail()
		}
		t.Log(result)
	})
}
