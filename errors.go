package envx

import "fmt"

type EmptyEnvVarError struct {
	Name string
}

func (e *EmptyEnvVarError) Error() string {
	return fmt.Sprintf("environment variable %s does not exist", e.Name)
}

type InvalidEnvVarTypeError struct {
	Name string
	Type string
}

func (e *InvalidEnvVarTypeError) Error() string {
	return fmt.Sprintf("environment variable %s could not be converted to type %s", e.Name, e.Type)
}

