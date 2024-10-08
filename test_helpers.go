package envx

import (
	"testing"
	"errors"
)

func testGetenv[T comparable](
	t testing.TB,
	getenvFunc func(string)*Value[T],
	name string,
	expectedValue T,
	expectedError error,
) {
	t.Helper()

	equalityCheck := func(a T, b T) bool { return a == b }
	testGetenvWithEqualityCheck(t, getenvFunc, name, expectedValue, expectedError, equalityCheck)
}

func testGetenvWithEqualityCheck[T comparable](
	t testing.TB,
	getenvFunc func(string)*Value[T],
	name string,
	expectedValue T,
	expectedError error,
	equalityCheck func(T,T)bool,
) {
	t.Helper()

	value, err := getenvFunc(name).Value()

	var underlyingErr error
	if err != nil {
		underlyingErr = err.(Error).Err // cast to Error - we should always expect the underlying error of being this type
	}

	if !errors.Is(expectedError, underlyingErr) {
		t.Errorf("unenexpected error for %s: got %#v, want %#v", name, err, expectedError)
	}

	if !equalityCheck(value, expectedValue) {
		t.Errorf("unenexpected value for %s: got ('%v'), want ('%v')", name, value, expectedValue)
	}
}

