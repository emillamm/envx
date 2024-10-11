package envx

import "strconv"

// [EnvX.Int] returns a [Value] of type int that references the environment variable given by the name.
// Calling [Value.Value] will read, parse and return the environment variable.
func (env EnvX) Int(name string) *Value[int] {
	return getValue[int](name, env, strconv.Atoi)
}

