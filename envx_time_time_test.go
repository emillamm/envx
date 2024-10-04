package envx

//import (
//	"testing"
//	"time"
//)
//
//func TestEnvXTime(t *testing.T) {
//
//	var env EnvX = func(name string) string {
//		switch name {
//		case "FOO":
//			return "2006-01-02T15:04:05+07:00"
//		case "BAR":
//			return "1INVALID987"
//		default:
//			return ""
//		}
//	}
//
//	t.Run("Getenv should return the value and provide no errors if the variable exists" , func(t *testing.T) {
//		expectedTime, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
//		testGetenvWithEqualityCheck[time.Time](
//			t,
//			env.AsTime(time.RFC3339).Getenv,	// target func
//			"FOO",					// variable name
//			expectedTime,				// expected value
//			nil,					// expected error
//			func (a time.Time, b time.Time) bool { return a.Equal(b) },
//		)
//		if err != nil {
//			t.Fatal(err)
//		}
//	})
//
//	t.Run("Getenv should return the zero-value and provide ErrEmptyValue if the variable doesn't exists" , func(t *testing.T) {
//		testGetenv[time.Time](
//			t,
//			env.AsTime(time.RFC3339).Getenv,	// target func
//			"BAZ",					// variable name
//			time.Time{},				// expected value
//			ErrEmptyValue,				// expected error
//		)
//	})
//
//	t.Run("Getenv should return the zero-value and provide ErrInvalidType if the string value doesn't match the expected type" , func(t *testing.T) {
//		testGetenv[time.Time](
//			t,
//			env.AsTime(time.RFC3339).Getenv,	// target func
//			"BAR",					// variable name
//			time.Time{},				// expected value
//			ErrInvalidType,				// expected error
//		)
//	})
//
//	t.Run("Getenv should return the zero-value and provide ErrInvalidType if the time layout is invalid" , func(t *testing.T) {
//		testGetenv[time.Time](
//			t,
//			env.AsTime("invalid").Getenv,	// target func
//			"FOO",				// variable name
//			time.Time{},			// expected value
//			ErrInvalidType,			// expected error
//		)
//	})
//}
//
