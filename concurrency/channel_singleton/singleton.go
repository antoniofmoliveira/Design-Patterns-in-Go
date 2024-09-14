package channel_singleton

// The channel that will be used to increment the counter
var addCh chan bool = make(chan bool)

// The channel that will receive a channel that will be used to receive the count
var getCountCh chan chan int = make(chan chan int)

// The channel that will be used to stop the goroutine
var quitCh chan bool = make(chan bool)

// The init function starts a goroutine that increments a counter every time it
// receives on the addCh channel, and returns the current value of the counter
// every time it receives on the getCountCh channel. It stops when it receives on
// the quitCh channel.
func init() {
	var count int
	go func(addCh <-chan bool, getCountCh <-chan chan int, quitCh <-chan bool) {
		for {
			select {
			case <-addCh:
				count++
			case ch := <-getCountCh:
				ch <- count
			case <-quitCh:
				return
			}
		}
	}(addCh, getCountCh, quitCh)
}

// The singleton struct
type singleton struct{}

// The only instance of the singleton
var instance singleton

// GetInstance returns the only instance of singleton
func GetInstance() *singleton {
	return &instance
}

// AddOne sends a signal on the add channel, which causes the goroutine in init
// to increment its counter.
func (s *singleton) AddOne() {
	addCh <- true
}

// GetCount returns the current count of the counter.
// This code snippet is a method of the `singleton` struct that returns the current count of a counter. Here's a succinct explanation:
// 1. It creates a new channel `resCh` to receive the count.
// 2. It sends `resCh` on the `getCountCh` channel, which is likely handled by a goroutine that maintains the counter.
// 3. The goroutine will send the current count on `resCh`.
// 4. The method receives the count from `resCh` and returns it.
// In essence, this method uses a channel to request the current count from a separate goroutine that manages the counter.
func (s *singleton) GetCount() int {
	resCh := make(chan int)
	defer close(resCh)
	getCountCh <- resCh
	return <-resCh
}

// Stop the goroutine in the singleton by sending a signal on the quit
// channel, and then closing all three channels. This is a no-op if the
// goroutine has already been stopped.
func (s *singleton) Stop() {
	quitCh <- true
	close(addCh)
	close(getCountCh)
	close(quitCh)
}
