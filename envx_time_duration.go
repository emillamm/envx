package envx

import "time"

func (env EnvX) Duration(name string) *Value[time.Duration] {
	return getValue[time.Duration](name, env, time.ParseDuration)
}

