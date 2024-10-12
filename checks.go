package envx

// Checks represents a temporary aggregation of multiple errors that can be converted to an AggregateError type.
type Checks []error

// NewChecks allocates a new [Checks] and returns the pointer.
func NewChecks() *Checks {
	return &Checks{}
}

// Err returns an [AggregateError] conatining all the errors that were collected while using [Check] for error handling.
// If no errors were encountered, this method returns nil.
func (c *Checks) Err() error {
	if c == nil || len(*c) == 0 {
		return nil
	}
	return &AggregateError{
		Errs: *c,
	}
}

// Check is used for collecting errors when reading multiple environment variables.
// It should be used as a wrapper around calls to [Value.Value] or [Value.Default] after which it will
// the value will pass through and the error will be intercepted and stored in a [Checks].
// The return type of this function is another function that accepts the [Checks] that will collect the error
// and returns the value. This functionality is best described via an example as seen below.
func Check[T comparable](t T, err error) func(*Checks)T {
	return func(checks *Checks) T {
		if err != nil {
			if checks == nil {
				s := []error{}
				*checks = s
			}
			*checks = append(*checks, err)
		}
		return t
	}
}

