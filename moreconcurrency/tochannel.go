package lib

// ToChannel takes a channel of done signal and a channel of values.
// It returns a new channel that will yield values from the given value stream,
// but with the type of the values asserted to T, and will close immediately
// after the done signal is received.
func ToChannel[T any](done <-chan interface{}, valueStream <-chan interface{}) <-chan T {
	stream := make(chan T)
	go func() {
		defer close(stream)
		for v := range valueStream {
			select {
			case <-done:
				return
			case stream <- v.(T):
			}
		}
	}()
	return stream
}
