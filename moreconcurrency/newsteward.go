package lib

import (
	"log"
	"time"
)

type StartGoroutineFn func(
	done <-chan interface{},
	pulseInterval time.Duration,
) (heartbeat <-chan interface{})

// NewSteward takes a timeout duration, and a StartGoroutineFn. It
// returns a new StartGoroutineFn that will start a goroutine that will
// start the given function, and will restart the given function if it
// fails to produce a value within the given timeout duration. The new
// function will also close the channel it returns if the done channel is
// closed.
func NewSteward(
	timeout time.Duration,
	startGoroutine StartGoroutineFn,
) StartGoroutineFn {
	return func(
		done <-chan interface{},
		pulseInterval time.Duration,
	) <-chan interface{} {
		heartbeat := make(chan interface{})
		go func() {
			defer close(heartbeat)
			var wardDone chan interface{}
			var wardHeartbeat <-chan interface{}
			startWard := func() {
				wardDone = make(chan interface{})
				wardHeartbeat = startGoroutine(OrDone(wardDone, done), timeout/2)
			}
			startWard()
			pulse := time.Tick(pulseInterval)
		monitorLoop:
			for {
				timeoutSignal := time.After(timeout)
				for {
					select {
					case <-pulse:
						select {
						case heartbeat <- struct{}{}:
						default:
						}
					case <-wardHeartbeat:
						continue monitorLoop
					case <-timeoutSignal:
						log.Println("steward: ward unhealthy; restarting")
						close(wardDone)
						startWard()
						continue monitorLoop
					case <-done:
						return
					}
				}
			}
		}()
		return heartbeat
	}
}
