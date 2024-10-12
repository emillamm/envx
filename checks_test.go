package envx

import (
	"fmt"
	"os"
	"testing"
)

func TestChecks(t *testing.T) {

	var env EnvX = func(name string) string {
		switch name {
		case "FOO":
			return "917"
		case "BAR":
			return "1INVALID987"
		case "BAZ":
			return "ABC"
		default:
			return ""
		}
	}

	t.Run("checks.Err() should initially be nil" , func(t *testing.T) {
		checks := NewChecks()
		if err := checks.Err(); err != nil || len(*checks) > 0 {
			t.Errorf("got (%v, %v), want (nil, 0)", err, len(*checks))
		}
	})

	t.Run("checks.Err() be nil if no errors are returned in the checks" , func(t *testing.T) {
		checks := NewChecks()

		if v := Check(env.Int("FOO").Value())(checks); v != 917 {
			t.Errorf("got %v, want 917", v)
		}

		if v := Check(env.String("BAZ").Value())(checks); v != "ABC" {
			t.Errorf("got %v, want ABC", v)
		}

		if err := checks.Err(); err != nil || len(*checks) > 0 {
			t.Errorf("got (%v, %v), want (nil, 0)", err, len(*checks))
		}
	})

	t.Run("checks.Err() be not be nil if errors are returned in the checks" , func(t *testing.T) {
		checks := NewChecks()

		// not resulting in errors
		Check(env.Int("FOO").Value())(checks)
		Check(env.String("BAZ").Value())(checks)

		// resulting in errors
		if v := Check(env.Int("BAR").Value())(checks); v != 0 {
			t.Errorf("got %v, want 0", v)
		}

		if v := Check(env.String("XYZ").Value())(checks); v != "" {
			t.Errorf("got %v, want ABC", v)
		}

		if err := checks.Err(); err == nil || len(*checks) != 2 {
			t.Errorf("got (%v, %v), want (AggregateError, 2)", err, len(*checks))
		}
	})
}

func ExampleCheck() {

	// Initialize env and checks
	var env EnvX = os.Getenv
	checks := NewChecks()

	// Intercepts ErrEmptyValue
	str := Check(env.String("NON_EXISTING_STRING").Value())(checks)

	// Intercepts ErrInvalidType
	os.Setenv("INVALID_INT", "A1Q")
	i := Check(env.Int("INVALID_INT").Value())(checks)

	fmt.Printf("Returned values str: %#v, i: %v\n", str, i)
	fmt.Print(checks.Err())

	// Output:
	// Returned values str: "", i: 0
	// Unable to read 2 environment variable(s):
	// Error reading environment variable 'NON_EXISTING_STRING' with type 'string': environment variable does not exist
	// Error reading environment variable 'INVALID_INT' with type 'int': environment variable could not be converted to expected type
}

