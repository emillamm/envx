package envx

import (
	"fmt"
	"strings"
)

// Error - generic error that wraps other error types
type Error struct {
	Err error
	Name string
	Type string
}

func WrapError[T comparable](err error, name string, value T) Error {
	return Error{
		Err: err,
		Name: name,
		Type: fmt.Sprintf("%T", value),
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error reading environment variable '%s' with type '%s': %s", e.Name, e.Type, e.Err)
}

func (e *Error) Unwrap() error { return e.Err }


// Errors - represents a temporary aggregation of multiple Error. It can be converted to the AggregateError type.
// Having two separate structs helps simplifying nil checks as the `Error() error` method explicitly returns nil.
type Errors struct {
	Observed []Error
}

func (e Errors) Error() error {
	if len(e.Observed) == 0 {
		return nil
	}
	return &AggregateError{e.Observed}
}


// AggregateError - the error type that represents an aggregation of multiple Error
type AggregateError struct {
	Observed []Error
}

func (e *AggregateError) Error() string {
	var b strings.Builder
	if len(e.Observed) > 0 {
		b.WriteString(fmt.Sprintf("Unable to read %d environment variable(s):", len(e.Observed)))
		for _, err := range e.Observed {
			b.WriteString("\n")
			b.WriteString(err.Error())
		}
	}
	return b.String()
}


// EmptyValueError
type EmptyValueError struct {}

func (e *EmptyValueError) Error() string {
	return "environment variable does not exist"
}

var ErrEmptyValue = &EmptyValueError{}


// InvalidTypeError
type InvalidValueTypeError struct {}

func (e *InvalidValueTypeError) Error() string { return "environment variable could not be converted to expected type" }

var ErrInvalidType = &InvalidValueTypeError{}

