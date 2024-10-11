package envx

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestEnvXDuration(t *testing.T) {

	var env EnvX = func(name string) string {
		switch name {
		case "FOO":
			return "5h31m"
		case "BAR":
			return "1INVALID987"
		default:
			return ""
		}
	}

	t.Run("Getenv should return the value and provide no errors if the variable exists" , func(t *testing.T) {
		testGetenv[time.Duration](
			t,
			env.Duration,				// target func
			"FOO",					// variable name
			5 * time.Hour + 31 * time.Minute,	// expected value
			nil,					// expected error
		)
	})

	t.Run("Getenv should return the zero-value and provide ErrEmptyValue if the variable doesn't exists" , func(t *testing.T) {
		testGetenv[time.Duration](
			t,
			env.Duration,	// target func
			"BAZ",				// variable name
			0,				// expected value
			ErrEmptyValue,			// expected error
		)
	})

	t.Run("Getenv should return the zero-value and provide ErrInvalidType if the string value doesn't match the expected type" , func(t *testing.T) {
		testGetenv[time.Duration](
			t,
			env.Duration,	// target func
			"BAR",				// variable name
			0,				// expected value
			ErrInvalidType,			// expected error
		)
	})
}

func ExampleEnvX_Duration() {
	os.Setenv("FOO_DURATION", "5h31m")

	var env EnvX = os.Getenv
	var value time.Duration
	var err error

	value, err = env.Duration("FOO_DURATION").Value()

	fmt.Printf("value: %v, error: %v", value, err)
	// Output:
	// value: 5h31m0s, error: <nil>
}

