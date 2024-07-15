package envx

import "testing"

func TestEnv(t *testing.T) {

	var env EnvX = func(name string) string {
		switch name {
		case "FOO":
			return "bar"
		default:
			return ""
		}
	}

	t.Run("Getenv should return the value of a variable if it exists and otherwise return the zero-value", func(t *testing.T) {
		if v := env.Getenv("FOO"); v != "bar" {
			t.Errorf("FOO: got '%s', want 'bar'", v)
		}

		if v := env.Getenv("BAZ"); v != "" {
			t.Errorf("BAZ: got '%s', want ''", v)
		}
	})

	t.Run("GetenvOrError should return an error if the variable doesn't exists", func(t *testing.T) {
		if v, err := env.GetenvOrError("FOO"); v != "bar" || err != nil {
			t.Errorf("unenexpected value/error for FOO: got ('%s','%#v'), want ('bar','nil')", v, err)
		}

		v, err := env.GetenvOrError("BAZ")
		emptyValueError, isExpectedErr := err.(*EmptyValueError)
		if v != "" || !isExpectedErr || emptyValueError.Name != "BAZ" {
			t.Errorf("unenexpected value/error for BAZ: got ('%s','%#v'), want ('','EmptyValueError{BAZ}')", v, err)
		}
	})

	t.Run("GetenvOrDefault should return a default value if the variable doesn't exists", func(t *testing.T) {
		if v, err := env.GetenvOrDefault("FOO", "foo_default"); v != "bar" || err != nil {
			t.Errorf("unenexpected value/error for FOO: got ('%s','%#v'), want ('bar','nil')", v, err)
		}

		if v, err := env.GetenvOrDefault("BAZ", "baz_default"); v != "baz_default" || err != nil {
			t.Errorf("unenexpected value/error for BAZ: got ('%s','%#v'), want ('baz_value','nil')", v, err)
		}
	})
}

