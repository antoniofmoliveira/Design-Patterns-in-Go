package channel_singleton

import (
	"fmt"
	"testing"
	"time"
)

// This is a test function in Go that tests the behavior of a singleton instance in a concurrent environment. Here's a succinct explanation:
// 1. It gets two instances of a singleton (`singleton` and `singleton2`) and verifies that they are the same instance (not shown in this code snippet, but likely done elsewhere).
// 2. It starts 10,000 concurrent goroutines (5000 times 2) that increment the singleton's counter using `AddOne()`.
// 3. It prints the current count before the loop finishes.
// 4. It waits until the count reaches 10,000 (i.e., all goroutines have finished incrementing the counter) by polling the count every 5 milliseconds.
// 5. Finally, it calls `Stop()` on the singleton instance.
// This test likely aims to verify that the singleton instance is thread-safe and can handle concurrent access correctly.
func TestStartInstance(t *testing.T) {
	singleton := GetInstance()
	singleton2 := GetInstance()
	n := 5000
	for i := 0; i < n; i++ {
		go singleton.AddOne()
		go singleton2.AddOne()
	}
	fmt.Printf("Before loop, current count is %d\n", singleton.GetCount())
	var val int
	for val != n*2 {
		val = singleton.GetCount()
		time.Sleep(5 * time.Millisecond)
	}
	singleton.Stop()
}
