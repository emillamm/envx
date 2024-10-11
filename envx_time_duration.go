package envx

import "time"

// [EnvX.Duration] returns a [Value] of type [time.Duration] that references the environment variable given by the name.
// Calling [Value.Value] will read, parse and return the environment variable.
func (env EnvX) Duration(name string) *Value[time.Duration] {
	return getValue[time.Duration](name, env, time.ParseDuration)
}

