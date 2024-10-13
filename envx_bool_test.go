package envx

import (
	"testing"
	"os"
	"fmt"
)

func TestEnvXBool(t *testing.T) {

	var env EnvX = func(name string) string {
		switch name {
		case "FOO":
			return "true"
		case "BAR":
			return "apple"
		default:
			return ""
		}
	}

	t.Run("Getenv should return the value and provide no errors if the variable exists" , func(t *testing.T) {
		testGetenv[bool](
			t,
			env.Bool,		// target func
			"FOO",			// variable name
			true,			// expected value
			nil,			// expected error
		)
	})

	t.Run("Getenv should return the zero-value and provide ErrEmptyValue if the variable doesn't exists" , func(t *testing.T) {
		testGetenv[bool](
			t,
			env.Bool,		// target func
			"BAZ",			// variable name
			false,			// expected value
			ErrEmptyValue,		// expected error
		)
	})

	t.Run("Getenv should return the zero-value and provide ErrInvalidType if the string value doesn't match the expected type" , func(t *testing.T) {
		testGetenv[bool](
			t,
			env.Bool,		// target func
			"BAR",			// variable name
			false,			// expected value
			ErrInvalidType,		// expected error
		)
	})
}

func ExampleEnvX_Bool() {
	os.Setenv("FOO_BOOL", "true")

	var env EnvX = os.Getenv
	var value bool
	var err error

	value, err = env.Bool("FOO_BOOL").Value()

	fmt.Printf("value: %v, error: %v", value, err)
	// Output:
	// value: true, error: <nil>
}


