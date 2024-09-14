package publish

import (
	"fmt"
	"io"
	"os"
	"time"
)

// This class definition defines an interface for a `Subscriber` in a publish-subscribe pattern. Here's what each method does:
// * `Notify(interface{}) error`: Notifies the subscriber with a message of any type (`interface{}`) and returns an error if the notification fails.
// * `Close()`: Closes the subscriber, likely releasing any resources it holds.
// Note that this is an interface, not a class, which means it defines a contract that must be implemented by any type that wants to be a `Subscriber`.
type Subscriber interface {
	Notify(interface{}) error
	Close()
}

// This class definition defines an interface for a `Publisher` in a publish-subscribe pattern. Here's what each method does:
// * `start()`: Starts the publisher, likely initializing its internal state and beginning the publishing process.
// * `AddSubscriberCh() chan<- Subscriber`: Returns a channel that allows subscribers to be added to the publisher.
// * `RemoveSubscriberCh() chan<- Subscriber`: Returns a channel that allows subscribers to be removed from the publisher.
// * `PublishingCh() chan<- interface{}`: Returns a channel that allows messages to be published to the subscribers.
// * `Stop()`: Stops the publisher, likely releasing any resources it holds and halting the publishing process.
type Publisher interface {
	start()
	AddSubscriberCh() chan<- Subscriber
	RemoveSubscriberCh() chan<- Subscriber
	PublishingCh() chan<- interface{}
	Stop()
}

// This is not a class definition, but a struct definition in Go. Here's a succinct explanation of what each field does:

// * `in`: a channel that receives messages of any type (`interface{}`)
// * `id`: an integer identifier for the subscriber
// * `Writer`: an output stream where messages will be written
type writerSubscriber struct {
	in     chan interface{}
	id     int
	Writer io.Writer
}

// Notify notifies the subscriber with the given message, and returns an error if
// the notification times out or fails in some other way.
func (s *writerSubscriber) Notify(msg interface{}) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%#v", rec)
		}
	}()
	select {
	case s.in <- msg:
	case <-time.After(time.Second):
		err = fmt.Errorf("Timeout")
	}
	return
}

// Close closes the subscriber, likely releasing any resources it holds.
func (s *writerSubscriber) Close() {
	close(s.in)
}

// NewWriterSubscriber returns a new Subscriber that writes the given message
// to the given output stream. If the output stream is nil, it defaults to
// os.Stdout. The subscriber is run in a separate goroutine, so the caller
// should call Close() on the subscriber to clean up resources when finished.
func NewWriterSubscriber(id int, out io.Writer) Subscriber {
	if out == nil {
		out = os.Stdout
	}
	s := &writerSubscriber{
		id:     id,
		in:     make(chan interface{}),
		Writer: out,
	}
	go func() {
		for msg := range s.in {
			fmt.Fprintf(s.Writer, "(W%d): %v\n", s.id, msg)
		}
	}()
	return s
}

// This class definition defines a `publisher` struct that represents a publisher in a publish-subscribe pattern. Here's what each field does:
// * `subscribers`: a list of subscribers that are currently subscribed to the publisher.
// * `addSubCh`: a channel that allows new subscribers to be added to the publisher.
// * `removeSubCh`: a channel that allows subscribers to be removed from the publisher.
// * `in`: a channel that allows messages to be published to the subscribers.
// * `stop`: a channel that allows the publisher to be stopped.
type publisher struct {
	subscribers []Subscriber
	addSubCh    chan Subscriber
	removeSubCh chan Subscriber
	in          chan interface{}
	stop        chan struct{}
}

// AddSubscriberCh returns a channel that can be used to add a new subscriber to the publisher.
func (p *publisher) AddSubscriberCh() chan<- Subscriber {
	return p.addSubCh
}

// RemoveSubscriberCh returns a channel that can be used to remove a subscriber from the publisher.
func (p *publisher) RemoveSubscriberCh() chan<- Subscriber {
	return p.removeSubCh
}

// PublishingCh returns a channel that can be used to publish messages to the subscribers.
func (p *publisher) PublishingCh() chan<- interface{} {
	return p.in
}

// Stop stops the publisher, likely releasing any resources it holds and halting the publishing process.
func (p *publisher) Stop() {
	close(p.stop)
}

// NewPublisher returns a new Publisher that can be used to publish messages to a list of subscribers.
func NewPublisher() Publisher {
	return &publisher{
		subscribers: make([]Subscriber, 0),
		addSubCh:    make(chan Subscriber),
		removeSubCh: make(chan Subscriber),
		in:          make(chan interface{}),
		stop:        make(chan struct{}),
	}
}

// start runs the main loop of the publisher. It continually listens for messages to publish, new subscribers to add, or subscribers to remove.
// If the publisher is stopped, it will close all the subscribers and channels and return.
func (p *publisher) start() {
	for {
		select {
		case msg := <-p.in:
			for _, sub := range p.subscribers {
				sub.Notify(msg)
			}
		case sub := <-p.addSubCh:
			p.subscribers = append(p.subscribers, sub)
		case sub := <-p.removeSubCh:
			for i, candidate := range p.subscribers {
				if candidate == sub {
					p.subscribers = append(p.subscribers[:i],
						p.subscribers[i+1:]...)
					candidate.Close()
					break
				}
			}
		case <-p.stop:
			for _, sub := range p.subscribers {
				sub.Close()
			}
			close(p.addSubCh)
			close(p.in)
			close(p.removeSubCh)
			return
		}
	}
}
