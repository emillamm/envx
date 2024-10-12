package envx

import (
	"fmt"
	"strings"
	"errors"
)

// ErrEmptyValue is returned as a wrapped error when a non-existing environment variable is read without a default value.
var ErrEmptyValue = errors.New("environment variable does not exist")

// ErrInvalidType is returned as a wrapped error when an environment variable cannot be parsed as the desired type.
var ErrInvalidType = errors.New("environment variable could not be converted to expected type")

// AggregateError is returned when calling [Checks.Err]. It represents an aggregation of multiple errors.
type AggregateError struct {
	Errs []error
}

// Error returns the error string of the aggregate error with one line per error.
func (e *AggregateError) Error() string {
	var b strings.Builder
	if len(e.Errs) > 0 {
		b.WriteString(fmt.Sprintf("Unable to read %d environment variable(s):", len(e.Errs)))
		for _, err := range e.Errs {
			b.WriteString("\n")
			b.WriteString(err.Error())
		}
	}
	return b.String()
}

