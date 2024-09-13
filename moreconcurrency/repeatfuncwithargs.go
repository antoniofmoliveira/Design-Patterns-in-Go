package lib

// RepeatFuncWithArgs takes a channel of done signal, a function that
// takes a variable number of arguments and returns interface, and a variable
// number of arguments.
// It returns a new channel that will yield values returned by the given
// function in order, and will repeat them indefinitely until the done channel
// is closed.
func RepeatFuncWithArgs[T any](
	done <-chan interface{},
	fn func(...T) interface{},
	args ...T,
) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				return
			case valueStream <- fn(args...):
			}
		}
	}()
	return valueStream
}