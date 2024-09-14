package mutexsingleton

import "sync"

// This class definition defines a singleton struct in Go, which is a design pattern that restricts the instantiation of a class to a single instance. Here's a brief explanation of each class method:
// * `GetInstance() *singleton`: Returns a pointer to the single instance of the singleton struct.
// * `(s *singleton) AddOne()`: Increments the `count` field of the singleton instance while acquiring a lock to ensure thread safety.
// * `(s *singleton) GetCount() int`: Returns the current value of the `count` field of the singleton instance while acquiring a read lock to ensure thread safety.
// Note that the `sync.RWMutex` field is used to provide thread-safe access to the `count` field, allowing multiple readers to access the field simultaneously while preventing writers from accessing it until all readers have finished.
type singleton struct {
	count int
	sync.RWMutex
}

// The single instance of the singleton struct.
var instance singleton

// GetInstance returns the single instance of the singleton struct, which is
// initialized the first time this function is called. Subsequent calls return
// the same instance. This ensures thread safety by preventing multiple goroutines
// from creating multiple instances of the singleton struct.
func GetInstance() *singleton {
	return &instance
}

// AddOne increments the `count` field of the singleton instance while acquiring
// a lock to ensure thread safety.
func (s *singleton) AddOne() {
	s.Lock()
	defer s.Unlock()
	s.count++
}

// GetCount returns the current value of the `count` field of the singleton instance
// while acquiring a read lock to ensure thread safety.
func (s *singleton) GetCount() int {
	s.RLock()
	defer s.RUnlock()
	return s.count
}
