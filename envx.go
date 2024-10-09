// Package envx is the only package you will need to import. See [EnvX] for examples that will get you started.
package envx

import "fmt"

// EnvX is a type definition that has the same signature as os.Getenv. It exposes several several methods that allows you to read and parse environment variables of different types.
type EnvX func(string)string

type Value[T comparable] struct {
	name string
	err error
	value *T
}

func (v *Value[T]) Value() (T, error) {
	var err error
	var value T
	if v.value != nil {
		value = *v.value
	}
	if v.err != nil {
		err = wrapError(v.name, value, v.err)
	}
	return value, err
}

func wrapError[T comparable](name string, value T, err error) error {
	return fmt.Errorf("Error reading environment variable '%s' with type '%T': %w", name, value, err)
}

func (v *Value[T]) Default(value T) (T, error) {
	if v.err == ErrEmptyValue {
		return value, nil
	}
	return v.Value()
}

func getValue[T comparable](name string, env EnvX, conv func(string)(T,error)) *Value[T] {

	v := Value[T]{
		name: name,
	}

	rawValue := env(name)

	if rawValue != "" {
		value, err := conv(rawValue)
		if err != nil {
			v.err = ErrInvalidType
		} else {
			v.value = &value
		}
	} else {
		v.err = ErrEmptyValue
	}

	return &v
}

