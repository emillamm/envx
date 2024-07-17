# EnvX
Envx is a lightweight Go module that simplifies consuming environment variables in your go programs.

It extends the `func(string)string` signature with helper methods and allows you to work directly with `os.Getenv` in your code while reducing common boilerplate code for handling errors, defaults and outputting non-string types.

## Usage

