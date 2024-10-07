package envx

import (
	//"errors"
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
		var checks Checks
		if err := checks.Err(); err != nil {
			t.Errorf("got %v, want nil", err)
		}
	})

	t.Run("checks.Err() be nil if no errors are returned in the checks" , func(t *testing.T) {
		var checks Checks
		if v := Check(checks, env.Int("FOO")); v != 917 {
			t.Errorf("got %v, want 917", v)
		}
		if v := Check(checks, env.String("BAZ")); v != "ABC" {
			t.Errorf("got %v, want ABC", v)
		}
	})

	//t.Run("checks.Err() be not be nil if errors are returned in the checks" , func(t *testing.T) {
	//	var checks Checks
	//	checks.Check(env
	//}

	//t.Run("Default should return a default value if a variable doesn't exist and unset the error" , func(t *testing.T) {
	//	if v, err := Default[string]("bar")("FOO", "", ErrEmptyValue); v != "bar" || err != nil {
	//		t.Errorf("unenexpected value/error: got ('%s','%#v'), want ('bar','nil')", v, err)
	//	}
	//})
	//t.Run("Default should return existing value if a variable already exist leave the error unchanged" , func(t *testing.T) {
	//	if v, err := Default[string]("bar")("FOO", "foo", nil); v != "foo" || err != nil {
	//		t.Errorf("unenexpected value/error: got ('%s','%#v'), want ('bar','nil')", v, err)
	//	}
	//	someError := errors.New("some error")
	//	if v, err := Default[string]("bar")("FOO", "foo", someError); v != "foo" || err != someError {
	//		t.Errorf("unenexpected value/error: got ('%s','%#v'), want ('foo','nil')", v, err)
	//	}
	//})

	//t.Run("Intercept should point to the provided error if the pointer is empty" , func(t *testing.T) {
}
