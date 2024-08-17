package envx

import "time"

func (env EnvX) AsTime(layout string) EnvXAny[time.Time] {
	return as[time.Time](env, func(s string) (time.Time, error) {
		return time.Parse(layout, s)
	})
}

