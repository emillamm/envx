package envx

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestEnvXTime(t *testing.T) {

	var env EnvX = func(name string) string {
		switch name {
		case "FOO":
			return "2006-01-02T15:04:05+07:00"
		case "BAR":
			return "1INVALID987"
		default:
			return ""
		}
	}

	getenvFunc := func(layout string) func(string)*Value[time.Time] {
		return func(name string) *Value[time.Time] {
			return env.Time(name, layout)
		}
	}

	t.Run("Getenv should return the value and provide no errors if the variable exists" , func(t *testing.T) {
		expectedTime, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
		testGetenvWithEqualityCheck[time.Time](
			t,
			getenvFunc(time.RFC3339),	// target func
			"FOO",				// variable name
			expectedTime,			// expected value
			nil,				// expected error
			func (a time.Time, b time.Time) bool { return a.Equal(b) },
		)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Getenv should return the zero-value and provide ErrEmptyValue if the variable doesn't exists" , func(t *testing.T) {
		testGetenv[time.Time](
			t,
			getenvFunc(time.RFC3339),	// target func
			"BAZ",				// variable name
			time.Time{},			// expected value
			ErrEmptyValue,			// expected error
		)
	})

	t.Run("Getenv should return the zero-value and provide ErrInvalidType if the string value doesn't match the expected type" , func(t *testing.T) {
		testGetenv[time.Time](
			t,
			getenvFunc(time.RFC3339),	// target func
			"BAR",				// variable name
			time.Time{},			// expected value
			ErrInvalidType,			// expected error
		)
	})

	t.Run("Getenv should return the zero-value and provide ErrInvalidType if the time layout is invalid" , func(t *testing.T) {
		testGetenv[time.Time](
			t,
			getenvFunc("invalid"),		// target func
			"FOO",				// variable name
			time.Time{},			// expected value
			ErrInvalidType,			// expected error
		)
	})
}

func ExampleEnvX_Time() {
	os.Setenv("FOO_TIME", "2006-01-02T15:04:05Z")

	var env EnvX = os.Getenv
	var value time.Time
	var err error

	value, err = env.Time("FOO_TIME", time.RFC3339).Value()

	fmt.Printf("value: %v, error: %v", value, err)
	// Output:
	// value: 2006-01-02 15:04:05 +0000 UTC, error: <nil>
}

