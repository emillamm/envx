package envx

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestEnvXString(t *testing.T) {

	var env EnvX = func(name string) string {
		switch name {
		case "FOO":
			return "bar"
		default:
			return ""
		}
	}

	t.Run("Getenv should return the value and provide no errors if the variable exists" , func(t *testing.T) {
		testGetenv[string](
			t,
			env.String,	// target func
			"FOO",		// variable name
			"bar",		// expected value
			nil,		// expected error
		)
	})

	t.Run("Getenv should return the zero-value and provide ErrEmptyValue if the variable doesn't exists" , func(t *testing.T) {
		testGetenv[string](
			t,
			env.String,		// target func
			"BAZ",			// variable name
			"",			// expected value
			ErrEmptyValue,		// expected error
		)
	})
}

func ExampleEnvX_String() {
	os.Setenv("FOO_STRING", "apple")

	var env EnvX = os.Getenv
	var value string
	var err error

	value, err = env.String("FOO_STRING").Value() // value is apple, err is nil
	fmt.Printf("value: %v, error: %v\n", value, err)

	value, err = env.String("BAR_STRING").Value() // BAR does not exist: value is "", err is ErrEmptyValue
	fmt.Printf("value: %#v, error: %v\n", value, err)
	fmt.Printf("error is ErrEmptyValue: %v\n", errors.Is(err, ErrEmptyValue))
	// Output:
	// value: apple, error: <nil>
	// value: "", error: Error reading environment variable 'BAR_STRING' with type 'string': environment variable does not exist
	// error is ErrEmptyValue: true
}

