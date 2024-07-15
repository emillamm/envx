package envx

func (env EnvX) AsString() EnvXAny[string] {
	return func(name string) (string, bool, error) {
		v := env(name)
		if v == "" {
			return "", false, nil
		}
		return v, true, nil
	}
}

// Proxy methods so you can write
// `env.Getenv("something")` instead of 
// `env.asString().Getenv("something")`
func (env EnvX) Getenv(name string) string {
	return env.AsString().Getenv(name)
}

func (env EnvX) GetenvOrError(name string) (string, error) {
	return env.AsString().GetenvOrError(name)
}

func (env EnvX) GetenvOrFatal(name string) string {
	return env.AsString().GetenvOrFatal(name)
}

func (env EnvX) GetenvOrDefault(name string, defaultValue string) (string, error) {
	return env.AsString().GetenvOrDefault(name, defaultValue)
}

