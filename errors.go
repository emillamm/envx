package envx

type EmptyValueError struct {}

func (e *EmptyValueError) Error() string {
	return "environment variable does not exist"
}

type InvalidValueTypeError struct {}

func (e *InvalidValueTypeError) Error() string {
	return "environment variable could not be converted to expected type"
}

var ErrInvalidType = &InvalidValueTypeError{}

var ErrEmptyValue = &EmptyValueError{}

