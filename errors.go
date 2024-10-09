package envx

import (
	"fmt"
	"strings"
	"errors"
)

var ErrEmptyValue = errors.New("environment variable does not exist")
var ErrInvalidType = errors.New("environment variable could not be converted to expected typek")

// AggregateError - the error type that represents an aggregation of multiple errors
type AggregateError struct {
	Errs []error
}

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

