package envx


import (
	"fmt"
	"log"
)

type EnvX func(string)string

type EnvXAny[T any] func(string)(T, bool, error)

func (env EnvXAny[T]) Getenv(name string) T {
	v, _, _ := env.getenv(name)
	return v
}

func (env EnvXAny[T]) GetenvOrError(name string) (v T, err error) {
	v, exists, err := env.getenv(name)
	if !exists {
		err = &EmptyValueError{Name: name}
	}
	return
}

func (env EnvXAny[T]) GetenvOrFatal(name string) T {
	v, err := env.GetenvOrError(name)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func (env EnvXAny[T]) GetenvOrDefault(name string, defaultValue T) (T, error) {
	v, exists, err := env.getenv(name)
	// err != nil only if a value already exists and it cannot be parsed as T
	// i.e. the error will always be nil if T is of type string.
	if !exists || err != nil {
		v = defaultValue
	}
	return v, err
}


func as[T any](env EnvX, conv func(string)(T,error)) EnvXAny[T] {
	return func(name string) (T, bool, error) {
		v, exists := env.getenv(name)
		if !exists {
			return *new(T), false, nil
		}
		converted, err := conv(v)
		if err != nil {
			err := &InvalidValueTypeError{
				Name: name,
				Type: fmt.Sprintf("%T", converted),
			}
			return converted, true, err
		}
		return converted, true, nil
	}
}


func (env EnvX) getenv(name string) (string, bool) {
	v := env(name)
	if v == "" {
		return "", false
	}
	return v, true
}

func (env EnvXAny[T]) getenv(name string) (T, bool, error) {
	return env(name)
}

