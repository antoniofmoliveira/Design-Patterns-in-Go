package lib

import (
	"crypto/rand"
	"math/big"
)

// RandomBigIntGenerator takes a quantity to generate,
// and a limit to generate random bigints within.
// It returns a new channel that will yield bigints in order, and will close
// immediately after the done signal is received.
// The returned channel will yield at most qt values before closing.
func RandomBigIntGenerator(qt int, limit int64) <-chan interface{} {
	done := make(chan interface{})
	defer close(done)
	randInt := func(limit int64) interface{} {
		n, _ := rand.Int(rand.Reader, big.NewInt(limit)) //1_000_000_000_000_000_000))
		return n
	}
	randIntWrapper := func(args ...int64) interface{} {
		return randInt(args[0])
	}
	return Take(done, RepeatFuncWithArgs(done, randIntWrapper, limit), qt)
}
