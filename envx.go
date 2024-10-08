package envx

//import "sync"

type EnvX func(string)string

func getValue[T comparable](name string, env EnvX, conv func(string)(T,error)) *Value[T] {
	//e := Value[T]{}
	//var initOnce sync.Once
	//e.init = func() {
	//	initOnce.Do(func () {
	//		e.raw = env(name)
	//		if e.raw == "" {
	//			e.err = NewError[T](ErrEmptyValue, name, e.value)
	//			return
	//		}
	//		value, err := conv(e.raw)
	//		e.value = value
	//		// for security regetValueons, don't include regetValueon behind the invalid type
	//		if err != nil {
	//			e.err = NewError[T](ErrInvalidType, name, value)
	//		}
	//	})
	//}
	//return &e


	//v := Value[T]{}
	//v.raw = env(name)
	//if v.raw == "" {
	//	v.err = NewError[T](ErrEmptyValue, name, v.value)
	//	return &v
	//}
	//value, err := conv(v.raw)
	//v.value = value
	//// for security reasons, don't include underlying invalid type error
	//if err != nil {
	//	v.err = NewError[T](ErrInvalidType, name, value)
	//}
	//return &v

	//value := Value[T]{}
	rawValue := env(name)	
	//var value *T
	//var err error

	v := Value[T]{
		name: name,
	}

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

	//return &Value[T]{
	//	name: name,
	//	err: err,
	//	value: value,
	//}
	return &v
}

type Value[T comparable] struct {
	name string
	//raw string
	//conv func(string)
	err error
	value *T
	//init func()
}

func (v *Value[T]) Value() (T, error) {
	//e.init()
	var err error
	var value T
	if v.value != nil {
		value = *v.value
	}
	if v.err != nil {
		err = NewError[T](v.err, v.name, value)
	}
	return value, err
}

// TODO
func (v *Value[T]) Default(value T) (T, error) {
	//e.init()
	if v.err == ErrEmptyValue {
		return value, nil
	}
	return v.Value()
}

//func (e *Value[T]) Err() error {
//	e.init()
//	return e.err
//}

//func (e *Value[T]) Raw() string {
//	//e.init()
//	return e.raw
//}

//func (e EnvX) Getenv(name string) string {
//	return e(name)
//}
//
//type GetenvFunc interface {
//	Getenv(name string) string
//}
//
////func (a AA) Int(name string) *Value[int] {
////	return nil
////}
//
//func (env EnvX) Int(name string) *Value[int] {
//	return getValue[int](name, env, strconv.Atoi)
//}
//
//type EnvXAny struct {
//	env EnvX
//	errs chan error
//}
//
//func (a EnvXAny) Int(name string) *Value[int] {
//	return a.env.Int(name)
//}
//
//type Conf interface {
//	Int(string) *Value[int]
//}
//
//func Wrap(env EnvX) Conf {
//	return EnvXAny{
//		env: env,
//		errs: make(chan error),
//	}
//}




// var checks Checks
//
//host := checks.Check(env.String("HOST").Default("localhost"))
//
//i.Error() // AggregateError
//
//
//errs, catch := ErrorCatcher()
//
//
//
//env, err := 
//
//obs := envx.Observe(env)
//
//host, err := env.String("HOST").Default("abc").AppendErr(err).Value()
//
//host := errs.intercept(env.String("HOST").Default("abc"))
//
//host := env.String("HOST").Default("abc"))
//
//env env.Start()
//
//host := env.String("HOST").Default().Err(errs)

//type EnvXAny interface {
//	getValue[T comparable](name string, env EnvX, conv func(string)(T,error)) *Value[T] {
//}


//func Batch(env EnvX) EnvX {
//}

//type EnvXAny[T comparable] func(string)Value[T]


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
//func getValue[T comparable](env EnvX, conv func(string)(T,error)) EnvXAny[T] {
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
