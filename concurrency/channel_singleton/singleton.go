package channel_singleton

var addCh chan bool = make(chan bool)
var getCountCh chan chan int = make(chan chan int)
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

type singleton struct{}

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
