package envx

import "testing"

func testGetenv[T comparable](
	t testing.TB,
	getenvFunc func(string, ...EnvXHandler[T])T,
	name string,
	expectedValue T,
	expectedError error,
) {
	t.Helper()

	var err error
	var value T = getenvFunc(name, Intercept[T](&err))

	if err != expectedError {
		t.Errorf("unenexpected error for %s: got %#v, want %#v", name, err, expectedError)
	}

	if value != expectedValue {
		t.Errorf("unenexpected value for %s: got ('%v'), want ('%v')", name, value, expectedValue)
	}
}

