package envx

import (
	"fmt"
	"os"
)

func Example_basic() {

	// Initialize values
	os.Setenv("FOO", "apple")
	os.Setenv("BAR", "97")
	var env EnvX = os.Getenv

	// Returns value of FOO as a string
	i, err := env.String("FOO").Value()
	fmt.Println(i)		// apple
	fmt.Println(err)	// <nil>

	// Returns value of BAR as an int
	j, err := env.Int("BAR").Value()
	fmt.Println(j)		// 97
	fmt.Println(err)	// <nil>

	// BAZ doesn't exist, returns ErrEmptyValue wrapped in a descriptive error
	k, err := env.Int("BAZ").Value()
	fmt.Println(k)		// 0
	fmt.Println(err)	// Error reading environment variable 'BAZ' with type 'int': environment variable does not exist

	// BAZ doesn't exist, returns provided default value
	l, err := env.Int("BAZ").Default(88)
	fmt.Println(l)		// 88
	fmt.Println(err)	// <nil>

	// Output:
	// apple
	// <nil>
	// 97
	// <nil>
	// 0
	// Error reading environment variable 'BAZ' with type 'int': environment variable does not exist
	// 88
	// <nil>
}

