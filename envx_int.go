package envx

import "strconv"

func (env EnvX) AsInt() EnvXAny[int] {
	return as[int](env, strconv.Atoi)
}

