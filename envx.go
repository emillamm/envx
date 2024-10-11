// Package envx is the only package you will need to import. See [EnvX] for examples that will get you started.
package envx

import "fmt"

// EnvX is a type definition that has the same signature as [os.Getenv]. It exposes several several methods that allows you to read and parse environment variables of different types.
// Each method returns a [Value] which in turn provides various methods to retrieve the value.
type EnvX func(string)string

type Value[T comparable] struct {
	// Name is the name (or key) of the environment variable
	Name string
	read func()(T,error)
}

func (v *Value[T]) Value() (t T, err error) {
	t, err = v.read()
	if err != nil {
		err = wrapError(v.Name, t, err)
	}
	return
}

func wrapError[T comparable](name string, value T, err error) error {
	return fmt.Errorf("Error reading environment variable '%s' with type '%T': %w", name, value, err)
}

func (v *Value[T]) Default(value T) (t T, err error) {
	t, err = v.read()
	if err == ErrEmptyValue {
		t = value
		err = nil
	}
	return
}

func getValue[T comparable](name string, env EnvX, conv func(string)(T,error)) *Value[T] {
	read := func() (t T, err error) {
		rawValue := env(name)
		if rawValue != "" {
			t, err = conv(rawValue)
			if err != nil {
				err = ErrInvalidType
			}
		} else {
			err = ErrEmptyValue
		}
		return
	}

	return &Value[T]{
		Name: name,
		read: read,
	}
}

