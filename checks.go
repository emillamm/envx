package envx

// Checks represents a temporary aggregation of multiple Error that can be converted to an AggregateError type.
type Checks []Error

func NewChecks() *Checks {
	return &Checks{}
}

func (c *Checks) Err() error {
	if c == nil || len(*c) == 0 {
		return nil
	}
	return &AggregateError{
		Errs: *c,
	}
}

func Check[T comparable](t T, err error) func(*Checks)T {
	return func(checks *Checks) T {
		if err != nil {
			if checks == nil {
				s := []Error{}
				*checks = s
			}
			*checks = append(*checks, err.(Error))
		}
		return t
	}
}

