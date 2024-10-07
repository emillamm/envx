package envx

func (env EnvX) String(name string) *Value[string] {
	return getValue[string](name, env, func(s string) (string, error) {
		return s, nil
	})
}

//
//func (env EnvX) AsString() EnvXAny[string] {
//	return getValue[string](env, func(value string) (string, error) {
//		return value, nil
//	})
//}
//
//// Proxy method so you can write
//// `env.Getenv("something")` instead of 
//// `env.AsString().Getenv("something")`
//func (env EnvX) Getenv(name string, options ...EnvXHandler[string]) string {
//	return env.AsString().Getenv(name, options...)
//}
//
