package envx

import (
	"sync"
	"strconv"
)

type EnvX func(string)string

func (env EnvX) Int(name string) *Entry[int] {
	return as[int](name, env, strconv.Atoi)
}

type EnvXAny struct {
	env EnvX
	errs chan error
}

func (a EnvXAny) Int(name string) *Entry[int] {
	return a.env.Int(name)
}

type Conf interface {
	Int(string) *Entry[int]
}

func Wrap(env EnvX) Conf {
	return EnvXAny{
		env: env,
		errs: make(chan error),
	}
}

//type EnvXAny interface {
//	as[T comparable](name string, env EnvX, conv func(string)(T,error)) *Entry[T] {
//}

func as[T comparable](name string, env EnvX, conv func(string)(T,error)) *Entry[T] {
	e := Entry[T]{}
	var initOnce sync.Once
	e.init = func() {
		initOnce.Do(func () {
			e.raw = env(name)
			e.value, e.err = conv(e.raw)
		})
	}
	return &e
}

type Entry[T comparable] struct {
	raw string
	err error
	value T
	init func()
}

func (e *Entry[T]) Value() T {
	e.init()
	return e.value
}

func (e *Entry[T]) Err() error {
	e.init()
	return e.err
}

func (e *Entry[T]) Raw() string {
	e.init()
	return e.raw
}

//func Batch(env EnvX) EnvX {
//}

//type EnvXAny[T comparable] func(string)Entry[T]


//type EnvXAny[T comparable] func(string)(T, error)
//
//type EnvXHandler[T comparable] func(string, T, error) (T, error)
//
//func (env EnvXAny[T]) Getenv(name string, handlers ...EnvXHandler[T]) T {
//	v, err := env(name)
//	for _, handler := range handlers {
//		v, err = handler(name, v, err)
//	}
//	return v
//}
//
//func as[T comparable](env EnvX, conv func(string)(T,error)) EnvXAny[T] {
//	return func(name string) (T, error) {
//		v := env(name)
//		if v == "" {
//			return *new(T), ErrEmptyValue
//		}
//		converted, err := conv(v)
//		if err != nil {
//			return converted, ErrInvalidType
//		}
//		return converted, nil
//	}
//}
//
