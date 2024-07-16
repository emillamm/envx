package envx

type EnvX func(string)string

type EnvXAny[T comparable] func(string)(T, error)

type EnvXHandler[T comparable] func(T, error) (T, error)

func (env EnvXAny[T]) Getenv(name string, handlers ...EnvXHandler[T]) T {
	v, err := env(name)
	for _, handler := range handlers {
		v, err = handler(v, err)
	}
	return v
}

func as[T comparable](env EnvX, conv func(string)(T,error)) EnvXAny[T] {
	return func(name string) (T, error) {
		v := env(name)
		if v == "" {
			return *new(T), ErrEmptyValue
		}
		converted, err := conv(v)
		if err != nil {
			return converted, ErrInvalidType
		}
		return converted, nil
	}
}

