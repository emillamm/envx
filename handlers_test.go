package envx

import (
	"errors"
	"testing"
)

func TestHandlers(t *testing.T) {

	t.Run("Default should return a default value if a variable doesn't exist and unset the error" , func(t *testing.T) {
		if v, err := Default[string]("bar")("", ErrEmptyValue); v != "bar" || err != nil {
			t.Errorf("unenexpected value/error: got ('%s','%#v'), want ('bar','nil')", v, err)
		}
	})
	t.Run("Default should return existing value if a variable already exist leave the error unchanged" , func(t *testing.T) {
		if v, err := Default[string]("bar")("foo", nil); v != "foo" || err != nil {
			t.Errorf("unenexpected value/error: got ('%s','%#v'), want ('bar','nil')", v, err)
		}
		someError := errors.New("some error")
		if v, err := Default[string]("bar")("foo", someError); v != "foo" || err != someError {
			t.Errorf("unenexpected value/error: got ('%s','%#v'), want ('foo','nil')", v, err)
		}
	})

	t.Run("Error should point to the provided error if the pointer is empty" , func(t *testing.T) {
		var errPtr error
		v, err := Error[string](&errPtr)("foo", ErrEmptyValue)

		if v != "foo" || err != ErrEmptyValue || errPtr != ErrEmptyValue {
			t.Errorf("unenexpected value/error: got ('%s','%#v','%#v'), want ('foo','ErrEmptyValue','ErrEmptyValue')", v, err, errPtr)
		}
	})
	t.Run("Error should not point to the provided error if the pointer is already pointing to another error" , func(t *testing.T) {
		var someError error = errors.New("some error")
		var errPtr error = someError
		v, err := Error[string](&errPtr)("foo", ErrEmptyValue)

		if v != "foo" || err != ErrEmptyValue || errPtr != someError {
			t.Errorf("unenexpected value/error: got ('%s','%#v','%#v'), want ('foo','ErrEmptyValue','ErrEmptyValue')", v, err, errPtr)
		}
	})

	t.Run("Panic should panic if any error is provided" , func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected a panic")
			}
		}()
		Panic[string]()("foo", ErrEmptyValue)
	})
	t.Run("Panic should not panic if no error is provided" , func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("expected no panic")
			}
		}()
		v, err := Panic[string]()("foo", nil)
		if v != "foo" || err != nil {
			t.Errorf("unenexpected value/error: got ('%s','%#v'), want ('foo','nil')", v, err)
		}
	})
}

