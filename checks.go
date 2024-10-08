package envx

import "fmt"

type Checks struct {
	Errs []Error
}

func (c Checks) Err() error {
	if len(c.Errs) == 0 {
		return nil
	}
	return &AggregateError{
		Errs: c.Errs,
	}
}

func Check[T comparable](t T, err error) func(Checks)T {
	return func(checks Checks) T {
		if err != nil {
			if checks.Errs == nil {
				s := []Error{}
				checks.Errs = &s
			}
			*checks.Errs = append(*checks.Errs, err.(Error))
		}
		return t
	}
}

