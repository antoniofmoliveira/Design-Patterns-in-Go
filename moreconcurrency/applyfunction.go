package lib

// ApplyFunction takes a channel of done signal, a channel of values,
// a value of the same type as the values in the channel, and a function that
// takes two values of the same type and returns a value of the same type.
// It returns a new channel that will yield values returned by the given
// function in order, and will close immediately after the done signal is
// received.
func ApplyFunction[T any](
	done <-chan interface{},
	stream <-chan T,
	value T,
	function func(T, T) T, // need more generics here
) <-chan T {
	operStream := make(chan T)
	go func() {
		defer close(operStream)
		for i := range stream {
			select {
			case <-done:
				return
			case operStream <- function(i, value):
			}
		}
	}()
	return operStream
}
