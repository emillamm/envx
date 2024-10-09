// Package envx provides a set of helper methods that simplifies comsuming environment variables in your go programs.
//
// It introduces a type definition, EnvX, which has the same signature as os.Getenv. EnvX has several several methods that allows you to read and parse environment variables of different types.
//
// Example:
//   // Create an `EnvX` type reference to `os.Getenv`
//   var env envx.EnvX = os.Getenv
//
//   // Returns the value of FOO parsed as an int.
//   // If FOO can't be parsed as an int, it returns ErrInvalidType
//   // If FOO doesn't exist, it returns ErrEmptyValue
//   i, err := env.Int("FOO").Value()
//
//   // Returns the value of BAR as a string.
//   // If BAR doesn't exist, it returns the default value "apple" instead.
//   j, err := env.String("BAR").Default("apple")
//
package envx

import "fmt"

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

