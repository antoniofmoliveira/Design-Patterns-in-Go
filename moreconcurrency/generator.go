package lib

// Generator takes a channel of done signal and a variable number of values.
// It returns a new channel that will yield all the values in the given values
// in order, and will close immediately after the done signal is received.
func Generator[T any](done <-chan interface{}, values ...T) <-chan T {
	intStream := make(chan T)
	go func() {
		defer close(intStream)
		for _, i := range values {
			select {
			case <-done:
				return
			case intStream <- i:
			}
		}
	}()
	return intStream
}
