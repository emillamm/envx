// Package envx is the only package you will need to import. Refer to [EnvX] for the main functionality for reading enviroment variables. Use examples below to get started.

package envx

import "fmt"

// EnvX is a type definition that has the same signature as [os.Getenv]. It exposes several several methods that allows you to read and parse environment variables of different types.
// Each method returns a [Value] which in turn provides various methods to retrieve the value.
type EnvX func(string)string

// Value represents the name and content of an environment variable after it has been parsed into a desired type.
type Value[T comparable] struct {
	// Name is the name (or key) of the environment variable
	Name string
	read func()(T,error)
}

// [Value.Value] reads and parses the environment variable given by Value.Name.
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

// [Value.Default] reads and parses the environment variable given by Value.Name. If it doesn't exist,
// it falls back to returning the the provided default value.
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

