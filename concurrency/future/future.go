package future

import "fmt"

// SuccessFunc is a type alias for a function that takes a string and returns nothing.
type SuccessFunc func(string)

// FailFunc is a type alias for a function that takes an error and returns nothing.
type FailFunc func(error)

// ExecuteStringFunc is a type alias for a function that returns a string and an error.
type ExecuteStringFunc func() (string, error)

// The `MaybeString` class is a struct that represents a future string value. It has two methods:
// * `Success(f SuccessFunc) *MaybeString`: This method sets the success function that will be called when the future string value is successfully retrieved. It returns a pointer to the `MaybeString` object, allowing method chaining.
// * `Fail(f FailFunc) *MaybeString`: This method sets the fail function that will be called when there is an error retrieving the future string value. It returns a pointer to the `MaybeString` object, allowing method chaining.
// The `MaybeString` struct also has two fields:
// * `successFunc SuccessFunc`: This field holds the success function that will be called when the future string value is successfully retrieved.
// * `failFunc FailFunc`: This field holds the fail function that will be called when there is an error retrieving the future string value.
// In summary, the `MaybeString` class is a way to represent a future string value that can be retrieved asynchronously. It allows you to set success and fail functions that will be called when the future value is either successfully retrieved or an error occurs.
type MaybeString struct {
	successFunc SuccessFunc
	failFunc    FailFunc
}

// Success sets the success function that will be called when the future string
// value is successfully retrieved. It returns a pointer to the `MaybeString`
// object, allowing method chaining.
func (s *MaybeString) Success(f SuccessFunc) *MaybeString {
	s.successFunc = f
	return s
}

// Fail sets the fail function that will be called when there is an error
// retrieving the future string value. It returns a pointer to the
// `MaybeString` object, allowing method chaining.
func (s *MaybeString) Fail(f FailFunc) *MaybeString {
	s.failFunc = f
	return s
}

// Execute calls the given function and calls either the success or fail function
// that was previously set, depending on whether the given function returns an
// error or not. It does this in a goroutine, so it will not block the calling
// goroutine.
func (s *MaybeString) Execute(f ExecuteStringFunc) {
	go func(s *MaybeString) {
		str, err := f()
		if err != nil {
			s.failFunc(err)
		} else {
			s.successFunc(str)
		}
	}(s)
}

// setContext returns a function that satisfies the ExecuteStringFunc interface by
// simply returning the given message string and a nil error. The returned
// function is a closure that captures the message string and returns it when
// called.
func setContext(msg string) ExecuteStringFunc {
	msg = fmt.Sprintf("%s Closure!\n", msg)
	return func() (string, error) {
		return msg, nil
	}
}
