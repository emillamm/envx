package envx

//import (
//	"testing"
//	"time"
//)
//
//func TestEnvXDuration(t *testing.T) {
//
//	var env EnvX = func(name string) string {
//		switch name {
//		case "FOO":
//			return "10h"
//		case "BAR":
//			return "1INVALID987"
//		default:
//			return ""
//		}
//	}
//
//	t.Run("Getenv should return the value and provide no errors if the variable exists" , func(t *testing.T) {
//		testGetenv[time.Duration](
//			t,
//			env.AsDuration().Getenv,	// target func
//			"FOO",				// variable name
//			10 * time.Hour,			// expected value
//			nil,				// expected error
//		)
//	})
//
//	t.Run("Getenv should return the zero-value and provide ErrEmptyValue if the variable doesn't exists" , func(t *testing.T) {
//		testGetenv[time.Duration](
//			t,
//			env.AsDuration().Getenv,	// target func
//			"BAZ",				// variable name
//			0,				// expected value
//			ErrEmptyValue,			// expected error
//		)
//	})
//
//	t.Run("Getenv should return the zero-value and provide ErrInvalidType if the string value doesn't match the expected type" , func(t *testing.T) {
//		testGetenv[time.Duration](
//			t,
//			env.AsDuration().Getenv,	// target func
//			"BAR",				// variable name
//			0,				// expected value
//			ErrInvalidType,			// expected error
//		)
//	})
//}
//
