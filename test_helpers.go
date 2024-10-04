package envx

//import "testing"

//func testGetenv[T comparable](
//	t testing.TB,
//	getenvFunc func(string, ...EnvXHandler[T])T,
//	name string,
//	expectedValue T,
//	expectedError error,
//) {
//	t.Helper()
//
//	equalityCheck := func(a T, b T) bool { return a == b }
//	testGetenvWithEqualityCheck(t, getenvFunc, name, expectedValue, expectedError, equalityCheck)
//}
//
//func testGetenvWithEqualityCheck[T comparable](
//	t testing.TB,
//	getenvFunc func(string, ...EnvXHandler[T])T,
//	name string,
//	expectedValue T,
//	expectedError error,
//	equalityCheck func(T,T)bool,
//) {
//	t.Helper()
//
//	var err error
//	var value T = getenvFunc(name, Intercept[T](&err))
//
//	if err != expectedError {
//		t.Errorf("unenexpected error for %s: got %#v, want %#v", name, err, expectedError)
//	}
//
//	if !equalityCheck(value, expectedValue) {
//		t.Errorf("unenexpected value for %s: got ('%v'), want ('%v')", name, value, expectedValue)
//	}
//}
//
