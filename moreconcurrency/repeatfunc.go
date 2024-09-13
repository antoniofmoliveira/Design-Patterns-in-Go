package lib

// RepeatFunc takes a channel of done signal and a function that
// returns interface.
// It returns a new channel that will yield values returned by the given
// function in order, and will repeat them indefinitely until the done channel
// is closed.
func RepeatFunc(
	done <-chan interface{},
	fn func() interface{},
) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				return
			case valueStream <- fn():
			}
		}
	}()
	return valueStream
}
