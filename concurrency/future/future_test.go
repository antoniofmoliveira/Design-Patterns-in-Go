package future

import (
	"errors"
	"sync"
	"testing"
	"time"
)

// timeout logs "Timeout!" after 1 second and calls t.Fail().
// It must be called as a goroutine.
func timeout(t *testing.T, wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	t.Log("Timeout!")
	t.Fail()
	wg.Done()
}



// This is a Go test function that exercises the `MaybeString` type's `Execute` method in three different scenarios:
// 1. **Success result**: The `Execute` method is called with a function that returns a successful string result ("Hello World!"). The `Success` callback is expected to be called with this result.
// 2. **Failed result**: The `Execute` method is called with a function that returns an error. The `Fail` callback is expected to be called with this error.
// 3. **Closure Success result**: The `Execute` method is called with a function that returns a successful string result, but this function is created using a closure (`setContext("Hello")`). The `Success` callback is expected to be called with this result.
// In each scenario, a timeout goroutine is started to ensure the test fails if the `Execute` method doesn't complete within a certain time. The `WaitGroup` is used to synchronize the test and ensure that the `Execute` method has completed before the test finishes.
func TestStringOrError_Execute(t *testing.T) {
	future := &MaybeString{}
	t.Run("Success result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)
		go timeout(t, &wg)
		future.Success(func(s string) {
			t.Log(s)
			wg.Done()
		}).Fail(func(e error) {
			t.Fail()
			wg.Done()
		})
		future.Execute(func() (string, error) {
			return "Hello World!", nil
		})
		wg.Wait()
	})

	t.Run("Failed result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)
		go timeout(t, &wg)
		future.Success(func(s string) {
			t.Fail()
			wg.Done()
		}).Fail(func(e error) {
			t.Log(e.Error())
			wg.Done()
		})
		future.Execute(func() (string, error) {
			return "", errors.New("Error ocurred")
		})
		wg.Wait()
	})

	t.Run("Closure Success result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)
		// Timeout!
		go timeout(t, &wg)
		future.Success(func(s string) {
			t.Log(s)
			wg.Done()
		}).Fail(func(e error) {
			t.Fail()
			wg.Done()
		})
		future.Execute(setContext("Hello"))
		wg.Wait()
	})
}
