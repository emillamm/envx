package envx

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
		err = NewError[T](v.err, v.name, value)
	}
	return value, err
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

