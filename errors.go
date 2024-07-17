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


// Errors - special error that represents an aggregation of multiple Error
type Errors []Error

func (e *Errors) Error() string {
	var b strings.Builder
	if len(*e) > 0 {
		b.WriteString(fmt.Sprintf("Unable to read %d environment variable(s):", len(*e)))
		for _, err := range *e {
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


