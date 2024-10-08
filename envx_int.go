package envx

import "strconv"

func (env EnvX) Int(name string) *Value[int] {
	return getValue[int](name, env, strconv.Atoi)
}

