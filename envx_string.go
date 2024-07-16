package envx

func (env EnvX) AsString() EnvXAny[string] {
	return as[string](env, func(value string) (string, error) {
		return value, nil
	})
}

// Proxy method so you can write
// `env.Getenv("something")` instead of 
// `env.AsString().Getenv("something")`
func (env EnvX) Getenv(name string, options ...EnvXHandler[string]) string {
	return env.AsString().Getenv(name, options...)
}

