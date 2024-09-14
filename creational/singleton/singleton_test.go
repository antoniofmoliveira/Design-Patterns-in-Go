package singleton

import "testing"

// This is a test function in Go that verifies the behavior of a singleton instance. Here's a brief explanation:
// 1. It gets an instance of the singleton (`counter1`) and checks if it's not nil.
// 2. It increments the counter using `AddOne()` and checks if the count is 1.
// 3. It gets another instance of the singleton (`counter2`) and checks if it's the same instance as `counter1`.
// 4. It increments the counter again using `AddOne()` and checks if the count is 2.
// The test ensures that the singleton instance is correctly initialized, incremented, and that multiple calls to `GetInstance()` return the same instance.
func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()
	if counter1 == nil {
		// Test of acceptance criteria 1 failed
		t.Error("expected pointer to Singleton after callingGetInstance(), not nil")
	}
	expectedCounter := counter1
	currentCount := counter1.AddOne()
	if currentCount != 1 {
		t.Errorf("After calling for the first time to count, the count must be 1 but it is %d\n", currentCount)
	}
	counter2 := GetInstance()
	if counter2 != expectedCounter {
		// Test 2 failed
		t.Error("Expected same instance in counter2 but it got a differentinstance")
	}
	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Errorf("After calling 'AddOne' using the second counter, the currentcount must be 2 but was %d\n", currentCount)
	}
}
