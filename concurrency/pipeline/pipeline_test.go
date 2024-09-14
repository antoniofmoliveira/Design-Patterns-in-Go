package pipeline

import "testing"

// This code snippet is a test function written in Go. It tests the `LaunchPipeline` function by providing a table of test cases. Each test case is a pair of integers stored in a two-dimensional slice called `tableTest`. The function iterates over each test case, calls the `LaunchPipeline` function with the first integer of the test case, and compares the result with the second integer of the test case. If the result does not match the expected value, the test fails and the test function terminates. If the result matches the expected value, a log message is printed.
func TestLaunchPipeline(t *testing.T) {
	tableTest := [][]int{
		{3, 14},
		{5, 55},
	}
	// ...
	var res int
	for _, test := range tableTest {
		res = LaunchPipeline(test[0])
		if res != test[1] {
			t.Fatal()
		}
		t.Logf("%d == %d\n", res, test[1])
	}
}
