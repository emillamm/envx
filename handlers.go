package envx

// return default value on error ErrEmptyValue, i.e. if the given variable doesn't exists
func Default[T comparable](defaultValue T) EnvXHandler[T] {
	return func(v T, err error) (T, error) {
		if err == ErrEmptyValue {
			return defaultValue, nil
		}
		return v, err
	}
}

// set error only if an error wasn't previously set (i.e. pointer is nil), otherwise no-op
func Error[T comparable](ptr *error) EnvXHandler[T] {
	return func(v T, err error) (T, error) {
		if err != nil && *ptr == nil {
			*ptr = err
		}
		return v, err
	}
}

// panic on any error (including if the variable doesn't exist)
func Panic[T comparable]() EnvXHandler[T] {
	return func(v T, err error) (T, error) {
		if err != nil {
			panic(err)
		}
		return v, err
	}
}

