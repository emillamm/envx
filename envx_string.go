package envx

// [EnvX.String] returns a [Value] of type string that references the environment variable given by the name.
// Calling [Value.Value] will read and return the environment variable.
// Note that if an empty string is read, it will result in an [ErrEmptyValue] error.
func (env EnvX) String(name string) *Value[string] {
	return getValue[string](name, env, func(s string) (string, error) {
		return s, nil
	})
}

