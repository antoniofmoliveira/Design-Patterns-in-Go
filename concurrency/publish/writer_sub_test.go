package publish

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"
)

type mockWriter struct {
	testingFunc func(string)
}

func (m *mockWriter) Write(p []byte) (n int, err error) {
	m.testingFunc(string(p))
	return len(p), nil
}

// func TestStdoutPrinter(t *testing.T) {
// }
// func TestWriter(t *testing.T) {
// 	sub := NewWriterSubscriber(0, nil)
// }

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

// publisher_test.go
type mockSubscriber struct {
	notifyTestingFunc func(msg interface{})
	closeTestingFunc  func()
}

func (m *mockSubscriber) Close() {
	m.closeTestingFunc()
}
func (m *mockSubscriber) Notify(msg interface{}) error {
	m.notifyTestingFunc(msg)
	return nil
}

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
