package envx

import "testing"

func TestHandlersIntegration(t *testing.T) {

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

	t.Run("Getenv should return a default value and provide no errors" , func(t *testing.T) {
		var err error
		v := env.AsInt().Getenv("BAZ", Default(101), Intercept[int](&err))
		if v != 101 || err != nil {
			t.Errorf("unenexpected value/error for BAZ: got ('%d','%#v'), want ('101',nil)", v, err)
		}
	})

	t.Run("Getenv should return a zero-value and provide ErrInvalidType" , func(t *testing.T) {
		var err error
		v := env.AsInt().Getenv("BAR", Default(101), Intercept[int](&err))
		if v != 0 || err != ErrInvalidType {
			t.Errorf("unenexpected value/error for BAZ: got ('%d','%#v'), want ('0','ErrInvalidType')", v, err)
		}
	})

	t.Run("Getenv should observe multiple errors" , func(t *testing.T) {
		var errs Errors
		if v := env.Getenv("FOO", Default("foo"), Observe[string](&errs)); v != "917" {
			t.Errorf("got %s, want '917'", v)
		}
		if v := env.AsInt().Getenv("BAR", Default(101), Observe[int](&errs)); v != 0 {
			t.Errorf("got %d, want 0", v)
		}
		if v := env.AsInt().Getenv("BAZ", Default(101), Observe[int](&errs)); v != 101 {
			t.Errorf("got %d, want 101", v)
		}
		if v := env.AsInt().Getenv("QUX", Observe[int](&errs)); v != 0 {
			t.Errorf("got %d, want 0", v)
		}
		if len(errs.Observed) != 2 {
			t.Errorf("wrong number of errors: got %d, want 2", len(errs.Observed))
		}
	})
}

