// This example demonstrates how to use `Checks` to simplify error handling
// when reading multiple environment variables.
package envx

import (
	"fmt"
	"os"
)

func Example_checks() {

	connection, err := GetConnection(os.Getenv)
	fmt.Printf("connection: %+v\n", connection)
	fmt.Printf("error: %v\n", err)

	fmt.Println("setting environment variables USER_TEST and PASSWORD_TEST")
	os.Setenv("USER_TEST", "bob")
	os.Setenv("PASSWORD_TEST", "abc123")

	connection, err = GetConnection(os.Getenv)
	fmt.Printf("connection: %+v\n", connection)
	fmt.Printf("error: %v\n", err)

	// Output:
	// connection: <nil>
	// error: Unable to read 2 environment variable(s):
	// Error reading environment variable 'USER_TEST' with type 'string': environment variable does not exist
	// Error reading environment variable 'PASSWORD_TEST' with type 'string': environment variable does not exist
	// setting environment variables USER_TEST and PASSWORD_TEST
	// connection: &{Host:localhost Port:8080 User:bob Pass:abc123}
	// error: <nil>
}

type Connection struct {
	Host string
	Port int
	User string
	Pass string
}

func GetConnection(env EnvX) (*Connection, error) {

	// This variable will accumulate errors that arise from reading the following environment variables
	checks := NewChecks()

	// By wrapping each read statement inside a "envx.Check(...)(checks)" expression
	// we are able to capture errors for later processing.
	host := Check(env.String("HOST_TEST").Default("localhost"))(checks)
	port := Check(env.Int("PORT_TEST").Default(8080))(checks)
	user := Check(env.String("USER_TEST").Value())(checks)
	pass := Check(env.String("PASSWORD_TEST").Value())(checks)

	// If any errors were encountered, "checks.Err()" will return a descriptive error
	// that summarizes all the underlying errors. Otherwise it returns nil.
	if err := checks.Err(); err != nil {
		return nil, err
	}

	con := Connection{
		Host: host,
		Port: port,
		User: user,
		Pass: pass,
	}

	return &con, nil
}

