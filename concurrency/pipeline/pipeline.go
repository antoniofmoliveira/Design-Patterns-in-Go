package pipeline

// LaunchPipeline takes an amount of integers to generate, and returns the sum of their squares.
func LaunchPipeline(amount int) int {
	// 	firstCh := generator(amount)
	// 	secondCh := power(firstCh)
	// 	thirdCh := sum(secondCh)
	// 	result := <-thirdCh
	// 	return result
	return <-sum(power(generator(amount)))
}

// generator takes an amount of integers to generate, and returns a channel that
// will yield them in order, and will close immediately after the done signal is
// received.
func generator(max int) <-chan int {
	outChInt := make(chan int, 100)
	go func() {
		for i := 1; i <= max; i++ {
			outChInt <- i
		}
		close(outChInt)
	}()
	return outChInt
}

// power takes a channel of integers, squares each of them and returns a new channel
// that will yield the squares in order, and will close immediately after the done
// signal is received from the input channel.
func power(in <-chan int) <-chan int {
	out := make(chan int, 100)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

// sum takes a channel of integers and returns a new channel that will yield
// the sum of the integers in order, and will close immediately after the done
// signal is received from the input channel.
func sum(in <-chan int) <-chan int {
	out := make(chan int, 100)
	go func() {
		var sum int
		for v := range in {
			sum += v
		}
		out <- sum
		close(out)
	}()
	return out
}
