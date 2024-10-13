package envx

import "strconv"

// [EnvX.Bool] returns a [Value] of type bool that references the environment variable given by the name.
// Calling [Value.Value] will read, parse and return the environment variable.
func (env EnvX) Bool(name string) *Value[bool] {
	return getValue[bool](name, env, strconv.ParseBool)
}

