package lib

// Repeat takes a channel of values and a done channel.
// It returns a new channel that will yield all the values in the given channel
// in order, and will repeat them indefinitely until the done channel is closed.
func Repeat(
	done <-chan interface{},
	values ...interface{},
) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case valueStream <- v:
				}
			}
		}
	}()
	return valueStream
}
