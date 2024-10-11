package envx

import "time"

// [EnvX.Time] returns a [Value] of type [time.Time] that references the environment variable given by the name.
// It also accepts a layout string (see [time.Layout]) to control the parsing format.
// Calling [Value.Value] will read, parse and return the environment variable.
func (env EnvX) Time(name string, layout string) *Value[time.Time] {
	return getValue[time.Time](name, env, func(s string) (time.Time, error) {
		return time.Parse(layout, s)
	})
}

