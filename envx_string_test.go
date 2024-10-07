package envx

import "testing"

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

