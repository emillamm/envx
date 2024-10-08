package envx

func (env EnvX) String(name string) *Value[string] {
	return getValue[string](name, env, func(s string) (string, error) {
		return s, nil
	})
}

