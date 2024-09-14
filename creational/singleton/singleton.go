package singleton

// * `AddOne() int`: This method is declared but not implemented. It is intended to be implemented by any type that satisfies the `Singleton` interface. When implemented, it should increment a counter and return the new count.
type Singleton interface {
	AddOne() int
}

// * The `singleton` struct is a custom data type that holds a single field called `count` of type `int`.
// * It does not have any methods defined in this snippet, but it is likely used as a building block for the `Singleton` interface and other related functions in the codebase.
// Note that in Go, classes are not explicitly defined like in other languages. Instead, structs are used to define custom data types, and methods can be attached to these structs using a separate syntax.
// This is not a thread safe implementation.
type singleton struct {
	count int
}

// The only instance of the singleton
var instance *singleton

// GetInstance returns the single instance of the singleton struct, which is
// initialized the first time this function is called. Subsequent calls return
// the same instance. This ensures thread safety by preventing multiple goroutines
// from creating multiple instances of the singleton struct.
func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

// AddOne increments the `count` field of the singleton instance and returns the
// new count.
func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
