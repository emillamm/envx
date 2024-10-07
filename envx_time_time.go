package envx

import "time"

func (env EnvX) Time(name string, layout string) *Value[time.Time] {
	return getValue[time.Time](name, env, func(s string) (time.Time, error) {
		return time.Parse(layout, s)
	})
}

