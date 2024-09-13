package lib

// Take takes a channel of done signal, a channel of values, and a
// number of times to take from the value stream. It returns a new channel
// that will yield at most num values from the given value stream, and will
// close immediately after the done signal is received.
func Take(
	done <-chan interface{},
	valueStream <-chan interface{},
	num int,
) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}
