package envx

import "testing"

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
		var checks Checks
		if err := checks.Err(); err != nil || len(checks.Errs) > 0 {
			t.Errorf("got (%v, %v), want (nil, 0)", err, len(checks.Errs))
		}
	})

	t.Run("checks.Err() be nil if no errors are returned in the checks" , func(t *testing.T) {
		var checks Checks

		if v := Check(env.Int("FOO"), &checks); v != 917 {
			t.Errorf("got %v, want 917", v)
		}

		if v := Check(env.String("BAZ"), &checks); v != "ABC" {
			t.Errorf("got %v, want ABC", v)
		}

		if err := checks.Err(); err != nil || len(checks.Errs) > 0 {
			t.Errorf("got (%v, %v), want (nil, 0)", err, len(checks.Errs))
		}
	})

	t.Run("checks.Err() be not be nil if errors are returned in the checks" , func(t *testing.T) {
		var checks Checks

		// not resulting in errors
		Check(env.Int("FOO"), &checks)
		Check(env.String("BAZ"), &checks)

		//var checks2 Checks2

		//s := Check2(env.String("FOO1").Value())(checks2)


		// resulting in errors
		if v := Check(env.Int("BAR"), &checks); v != 0 {
			t.Errorf("got %v, want 0", v)
		}

		if v := Check(env.String("XYZ"), &checks); v != "" {
			t.Errorf("got %v, want ABC", v)
		}

		if err := checks.Err(); err == nil || len(checks.Errs) != 2 {
			t.Errorf("got (%v, %v), want (AggregateError, 2)", err, len(checks.Errs))
		}
	})
}

