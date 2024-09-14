package publish

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"
)



// This is not a class definition, but a struct definition in Go. Here's a succinct explanation of what the struct and its method do:
// * `mockWriter`: a struct that holds a single field `testingFunc`, which is a function that takes a string as an argument.
// * The `Write` method (not shown in this snippet, but defined elsewhere in the code):
// 	+ Calls the `testingFunc` function with the string representation of the input byte slice.
// 	+ Returns the length of the input byte slice and a nil error.
type mockWriter struct {
	testingFunc func(string)
}

// Write implements the io.Writer interface by calling the testingFunc field
// with the string representation of the input byte slice.
//
// It returns the length of the input byte slice and a nil error.
func (m *mockWriter) Write(p []byte) (n int, err error) {
	m.testingFunc(string(p))
	return len(p), nil
}

// TestPublisher tests the Publisher type.
//
// It creates a WriterSubscriber connected to stdout, and a mockWriter that
// checks whether the message sent to the Publisher is the same as the one
// received by the subscriber.
//
// The Notify method is called on the Publisher with the message "Hello".
//
// The test checks that the message is received by the subscriber and that
// the Notify method does not return an error.
func TestPublisher(t *testing.T) {
	sub := NewWriterSubscriber(0, os.Stdout)
	msg := "Hello"
	var wg sync.WaitGroup
	wg.Add(1)
	stdoutPrinter := sub.(*writerSubscriber)
	stdoutPrinter.Writer = &mockWriter{
		testingFunc: func(res string) {
			if !strings.Contains(res, msg) {
				t.Fatal(fmt.Errorf("Incorrect string: %s", res))
			}
			wg.Done()
		},
	}
	err := sub.Notify(msg)
	if err != nil {
		wg.Done()
		t.Error(err)

	}
	wg.Wait()
	sub.Close()
}



// This class definition defines a `mockSubscriber` struct that implements a mock subscriber for testing purposes. Here's what each method does:
// * `Close()`: Closes the subscriber by calling the `closeTestingFunc` function, which is typically set up to perform some test-specific cleanup or verification.
// * `Notify(msg interface{}) error`: Notifies the subscriber with a message of any type (`interface{}`) by calling the `notifyTestingFunc` function, which is typically set up to perform some test-specific verification or assertion. The method returns `nil` as an error.
type mockSubscriber struct {
	notifyTestingFunc func(msg interface{})
	closeTestingFunc  func()
}

// Close calls the closeTestingFunc, which is typically set up to perform
// some test-specific cleanup or verification.
func (m *mockSubscriber) Close() {
	m.closeTestingFunc()
}

// Notify notifies the subscriber with a message of any type (`interface{}`)
// by calling the `notifyTestingFunc` function, which is typically set up to
// perform some test-specific verification or assertion. The method returns `nil`
// as an error.
func (m *mockSubscriber) Notify(msg interface{}) error {
	m.notifyTestingFunc(msg)
	return nil
}

// TestPublisher2 tests the Publisher type by verifying that a subscriber can be
// added, notified, removed, and then the publisher can be stopped. It also
// verifies that the number of subscribers is correctly updated at each step.
func TestPublisher2(t *testing.T) {
	msg := "Hello"
	p := NewPublisher()
	go p.start()
	var wg sync.WaitGroup
	sub := &mockSubscriber{
		notifyTestingFunc: func(msg interface{}) {
			defer wg.Done()
			s, ok := msg.(string)
			if !ok {
				t.Fatal(errors.New("Could not assert result"))
			}
			if s != msg {
				t.Fail()
			}
		},
		closeTestingFunc: func() {
			wg.Done()
		},
	}
	wg.Add(1)
	t.Log("p len", len(p.(*publisher).subscribers))
	p.AddSubscriberCh() <- sub
	t.Log("p len", len(p.(*publisher).subscribers))

	p.PublishingCh() <- msg
	wg.Wait()

	pubCon := p.(*publisher)
	if len(pubCon.subscribers) != 1 {
		t.Error("Unexpected number of subscribers")
	}
	wg.Add(1)
	p.RemoveSubscriberCh() <- sub
	wg.Wait()
	// Number of subscribers is restored to zero
	if len(pubCon.subscribers) != 0 {
		t.Error("Expected no subscribers")
	}
	p.Stop()
}
