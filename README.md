# EnvX
EnvX is a zero-dependency, lightweight, Go module that simplifies using environment variables in your go programs. It decorates `os.Getenv` with helper methods that reduces the need for boiler plate code when reading and parsing environment variables.

## Usage

### Reading and parsing environment variables
The following example demonstrates you can use EnvX to read environment variables of various types in Go code.

```
import {
    "github.com/emillamm/envx"
    "os"
}

func PrintEnv(env envx.EnvX) {

    // Returns value of FOO as a string
    i, err := env.String("FOO").Value()
    println(i)      // apple
    println(err)    // <nil>

    // Returns value of BAR as an int
    j, err := env.Int("BAR").Value()
    println(j)      // 97
    println(err)    // <nil>

    // BAZ doesn't exist, returns ErrEmptyValue wrapped in a descriptive error
    k, err := env.Int("BAZ").Value()
    println(k)      // 0
    println(err)    // Error reading environment variable 'FOO' with type 'string': environment variable does not exist

    // BAZ doesn't exist, returns provided default value
    l, err := env.Int("BAZ").Default(88)
    println(l)      // 88
    println(err)    // <nil>
}

func main() {
    os.Setenv("FOO", "apple")
    os.Setenv("BAR", "97")
    PrintEnv(os.Getenv)
}
```

### Reading multiple variables and returning a single error
When reading multiple environment variables, it can be useful to delay error handling until the very end, instead of having to check each error along the way. The following example demonstrates how to do this.

```
func GetConnection(env envx.EnvX) (*Connection, error) {

    // This variable will accumulate errors that arise from reading the following environment variables
    checks := envx.NewChecks()

    // By wrapping each read statement inside a "envx.Check(...)(checks)" expression
    // we are able to capture errors for later processing.
	host := envx.Check(env.String("HOST").Default("localhost"))(checks)
	port := envx.Check(env.Int("PORT").Default(8080))(checks)
	user := envx.Check(env.String("USER").Default("postgres"))(checks)
	pass := envx.Check(env.String("PASSWORD").Default("postgres"))(checks)

    // If any errors were encountered, "checks.Err()" will return a descriptive error
    // that summarizes all the underlying errors. Otherwise it returns nil.
    if err := checks.Err(); err != nil {
        return nil, err
    }

    return &Connection{
        Host: host,
        Port: port,
        User: user,
        Pass: pass,
    }
}
```

