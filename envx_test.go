package envx

import "testing"

func TestEnvX(t *testing.T) {

	var env EnvX = func(name string) string {
		switch name {
		case "FOO":
			return "bar"
		default:
			return ""
		}
	}

	t.Run("Value() should return a value if a variable exists" , func(t *testing.T) {
		if v, err := env.String("FOO").Value(); v != "bar" || err != nil {
			t.Errorf("got (%v, %#v), want (bar, nil)", v, err)
		}
	})

	t.Run("Value() should return ErrEmptyValue if the variable doesn't exists" , func(t *testing.T) {
		if v, err := env.String("BAR").Value(); v != "" || err == nil || err.(Error).Err != ErrEmptyValue {
			t.Errorf("got (%v, %#v), want ('', ErrEmptyValue)", v, err)
		}
	})

	t.Run("Value() should return ErrInvalidType if the variable can't be converted to the correct type" , func(t *testing.T) {
		if v, err := env.Int("FOO").Value(); v != 0 || err == nil || err.(Error).Err != ErrInvalidType {
			t.Errorf("got (%v, %#v), want (0, ErrInvalidType)", v, err)
		}
	})

	t.Run("Default() should return a default value if the variable doesn't exists" , func(t *testing.T) {
		if v, err := env.String("BAR").Default("baz"); v != "baz" || err != nil {
			t.Errorf("got (%v, %#v), want (baz, nil)", v, err)
		}
	})
}

