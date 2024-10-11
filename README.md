# EnvX
[![Go Reference](https://pkg.go.dev/badge/github.com/emillamm/envx.svg)](https://pkg.go.dev/github.com/emillamm/envx)

EnvX is a zero-dependency, lightweight, Go module that simplifies reading environment variables in your go programs. It reduces the need for boiler plate code when parsing environment variables into desired types.

## Usage
### Input source
`EnvX` is just a type definition of `func(string)string`. It has the same signature as `os.Getenv`.
```
var env EnvX = os.Getenv
```

It can use any function as input, it doesn't have to be `os.Getenv`. 
```
var env EnvX = func(name string) string {
        switch name {
        case "MY_TIMESTAMP":
            return "2006-01-02T15:04:05+07:00"
        case "MY_INT":
            return "844"
        default:
            return ""
        }
    }
}
```

### Reading variables
`EnvX` is decorated with helper method for reading `string` data and parsing it to various types
```
// Return string
str, err := env.String("MY_STRING").Value()

// Parse and return int
i, err := env.Int("MY_INT").Value()

// Parse and return time.Time using `time.RFC3339` as layout
timestamp, err := env.Time("MY_TIMESTAMP", time.RFC3339).Value()
```

### Errors
A descriptive error is returned if the environment variable doesn't exist

```
i, err := env.Int("NON_EXISTING_VAR").Value()
fmt.Println(i)      // 0
fmt.Println(err)    // Error reading environment variable 'NON_EXISTING_VAR' with type 'int': environment variable does not exist
```

Use `errors.Is` to check the type of error. For non-existing environment variables, the returned error wraps `ErrEmptyValue`.
```
i, err := env.Int("NON_EXISTING_VAR").Value()
if errors.Is(err, ErrEmptyValue) {
    fmt.Println("Doesn't exist")
}
```

If the underlying string can't be parsed to the desired type, the returned error wraps `ErrInvalidType`.
```
i, err := env.Int("NOT_AN_INT").Value()
if errors.Is(err, ErrInvalidType) {
    fmt.Printf("%v could not be parsed as an int", v)
}

fmt.Println(i)      // 0
fmt.Println(err)    // Error reading environment variable 'NOT_AN_INT' with type 'int': environment variable could not be converted to expected type
```

### Defaults
`.Value()` returns the value and an optional error as seen in previous examples.


`.Default(...)` behaves the same way, except that it falls back to default value instead of returning `ErrEmptyValue`.
```
i, err := env.Int("NON_EXISTING_VAR").Default(844)
fmt.Println(i)      // 844
fmt.Println(err)    // <nil>
```

### Deferred error handling
When reading multiple environment variables, it can be cumbersome to check each error along the way. `envx.Check`is a helper function that simplifies error handling when reading a batch of environment variables by summarizing all errors into a single error.

```
func GetConnection(env envx.EnvX) (*Connection, error) {

    // This variable will accumulate errors that arise from reading the following environment variables
    checks := envx.NewChecks()

    // By wrapping each read statement inside a "envx.Check(...)(checks)" expression
    // we are able to capture errors for later processing.
    host := Check(env.String("HOST").Default("localhost"))(checks)
    port := Check(env.Int("PORT").Default(8080))(checks)
    user := Check(env.String("USER").Value())(checks)
    pass := Check(env.String("PASSWORD").Value())(checks)

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
    }, nil
}
```

## Supported types
Review the [docs](https://pkg.go.dev/github.com/emillamm/envx#EnvX) to learn which types EnvX natively supports.

