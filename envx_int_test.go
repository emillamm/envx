package envx

import (
	"testing"
	"os"
	"fmt"
)

func TestEnvXInt(t *testing.T) {

	var env EnvX = func(name string) string {
		switch name {
		case "FOO":
			return "917"
		case "BAR":
			return "1INVALID987"
		default:
			return ""
		}
	}

	t.Run("Getenv should return the value and provide no errors if the variable exists" , func(t *testing.T) {
		testGetenv[int](
			t,
			env.Int,		// target func
			"FOO",			// variable name
			917,			// expected value
			nil,			// expected error
		)
	})

	t.Run("Getenv should return the zero-value and provide ErrEmptyValue if the variable doesn't exists" , func(t *testing.T) {
		testGetenv[int](
			t,
			env.Int,		// target func
			"BAZ",			// variable name
			0,			// expected value
			ErrEmptyValue,		// expected error
		)
	})

	t.Run("Getenv should return the zero-value and provide ErrInvalidType if the string value doesn't match the expected type" , func(t *testing.T) {
		testGetenv[int](
			t,
			env.Int,		// target func
			"BAR",			// variable name
			0,			// expected value
			ErrInvalidType,		// expected error
		)
	})
}

func ExampleEnvX_Int() {
	os.Setenv("FOO_INT", "87")

	var env EnvX = os.Getenv
	var value int
	var err error

	value, err = env.Int("FOO_INT").Value()

	fmt.Printf("value: %v, error: %v", value, err)
	// Output:
	// value: 87, error: <nil>
}

