package envx

//import (
//	"errors"
//	"testing"
//)
//
//func TestHandlers(t *testing.T) {
//
//	t.Run("Default should return a default value if a variable doesn't exist and unset the error" , func(t *testing.T) {
//		if v, err := Default[string]("bar")("FOO", "", ErrEmptyValue); v != "bar" || err != nil {
//			t.Errorf("unenexpected value/error: got ('%s','%#v'), want ('bar','nil')", v, err)
//		}
//	})
//	t.Run("Default should return existing value if a variable already exist leave the error unchanged" , func(t *testing.T) {
//		if v, err := Default[string]("bar")("FOO", "foo", nil); v != "foo" || err != nil {
//			t.Errorf("unenexpected value/error: got ('%s','%#v'), want ('bar','nil')", v, err)
//		}
//		someError := errors.New("some error")
//		if v, err := Default[string]("bar")("FOO", "foo", someError); v != "foo" || err != someError {
//			t.Errorf("unenexpected value/error: got ('%s','%#v'), want ('foo','nil')", v, err)
//		}
//	})
//
//	t.Run("Intercept should point to the provided error if the pointer is empty" , func(t *testing.T) {
//		var errPtr error
//		v, err := Intercept[string](&errPtr)("FOO", "foo", ErrEmptyValue)
//
//		if v != "foo" || err != ErrEmptyValue || errPtr != ErrEmptyValue {
//			t.Errorf("unenexpected value/error: got ('%s','%#v','%#v'), want ('foo','ErrEmptyValue','ErrEmptyValue')", v, err, errPtr)
//		}
//	})
//	t.Run("Intercept should not point to the provided error if the pointer is already pointing to another error" , func(t *testing.T) {
//		var someError error = errors.New("some error")
//		var errPtr error = someError
//		v, err := Intercept[string](&errPtr)("FOO", "foo", ErrEmptyValue)
//
//		if v != "foo" || err != ErrEmptyValue || errPtr != someError {
//			t.Errorf("unenexpected value/error: got ('%s','%#v','%#v'), want ('foo','ErrEmptyValue','ErrEmptyValue')", v, err, errPtr)
//		}
//	})
//
//	t.Run("Observe should append errors to the provided array pointer" , func(t *testing.T) {
//		var errs Errors
//		Observe[string](&errs)("FOO", "foo", ErrEmptyValue)
//		Observe[int](&errs)("BAR", 0, ErrInvalidType)
//
//		if len(errs.Observed) != 2 {
//			t.Errorf("expected 3 errors, got %d", len(errs.Observed))
//			return
//		}
//
//		var checkError = func(index int, name string, valueType string, err error) {
//			e := errs.Observed[index]
//			if e.Name != name {
//				t.Errorf("invalid name at index %d: got %s, want %s", index, e.Name, name)
//			}
//			if e.Type != valueType {
//				t.Errorf("invalid type at index %d: got %s, want %s", index, e.Type, valueType)
//			}
//			if e.Err != err {
//				t.Errorf("invalid error at index %d: got %#v, want %#v", index, e.Err, err)
//			}
//		}
//
//		checkError(0, "FOO", "string", ErrEmptyValue)
//		checkError(1, "BAR", "int", ErrInvalidType)
//
//		var err error = errs.Error()
//		if err == nil {
//			t.Errorf("err should not be nil")
//		}
//	})
//
//	t.Run("Panic should panic if any error is provided" , func(t *testing.T) {
//		defer func() {
//			if r := recover(); r == nil {
//				t.Errorf("expected a panic")
//			}
//		}()
//		Panic[string]()("FOO", "foo", ErrEmptyValue)
//	})
//	t.Run("Panic should not panic if no error is provided" , func(t *testing.T) {
//		defer func() {
//			if r := recover(); r != nil {
//				t.Errorf("expected no panic")
//			}
//		}()
//		v, err := Panic[string]()("FOO", "foo", nil)
//		if v != "foo" || err != nil {
//			t.Errorf("unenexpected value/error: got ('%s','%#v'), want ('foo','nil')", v, err)
//		}
//	})
//}
//
