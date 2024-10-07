package envx

import "cmp"

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

func Check[T comparable](checks Checks, entry *Value[T]) T {
	value, err := entry.Value()
	if err != nil {
		checks.Errs = append(checks.Errs, err.(Error))
	}
	return value
}

func (c Checks) Ch(entry *Value[cmp.Ordered]) {
	value, err := entry.Value()
	if err != nil {
		checks.Errs = append(checks.Errs, err.(Error))
	}
	return value
}

