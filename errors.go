package envx

import "fmt"

type EmptyValueError struct {
	Name string
}

func (e *EmptyValueError) Error() string {
	return fmt.Sprintf("environment variable %s does not exist", e.Name)
}

type InvalidValueTypeError struct {
	Name string
	Type string
}

func (e *InvalidValueTypeError) Error() string {
	return fmt.Sprintf("environment variable %s could not be converted to type %s", e.Name, e.Type)
}

