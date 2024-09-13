package publish

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Subscriber interface {
	Notify(interface{}) error
	Close()
}

type Publisher interface {
	start()
	AddSubscriberCh() chan<- Subscriber
	RemoveSubscriberCh() chan<- Subscriber
	PublishingCh() chan<- interface{}
	Stop()
}

type writerSubscriber struct {
	in     chan interface{}
	id     int
	Writer io.Writer
}

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

func (s *writerSubscriber) Close() {
	close(s.in)
}

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

type publisher struct {
	subscribers []Subscriber
	addSubCh    chan Subscriber
	removeSubCh chan Subscriber
	in          chan interface{}
	stop        chan struct{}
}

func (p *publisher) AddSubscriberCh() chan<- Subscriber {
	return p.addSubCh
}

func (p *publisher) RemoveSubscriberCh() chan<- Subscriber {
	return p.removeSubCh
}
func (p *publisher) PublishingCh() chan<- interface{} {
	return p.in
}

//	func (p *publisher) PublishingCh() chan<- interface{} {
//		return nil
//	}
func (p *publisher) Stop() {
	close(p.stop)
}

// publisher.go file
func NewPublisher() Publisher {
	return &publisher{
		subscribers: make([]Subscriber, 0),
		addSubCh:    make(chan Subscriber),
		removeSubCh: make(chan Subscriber),
		in:          make(chan interface{}),
		stop:        make(chan struct{}),
	}
}

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
