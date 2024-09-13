package lib

// OrDone takes a channel of done signal and a channel of values.
// It returns a new channel that will yield values from the given channel of values,
// and will close immediately after the done signal is received.
func OrDone(done, c <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if !ok {
					return
				}
				select {
				case valStream <- v:
				case <-done:
				}
			}
		}
	}()
	return valStream
}

// for val := range orDone(done, myChan) {
// 	// Do something with val
// }
