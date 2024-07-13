# envx
Envx is a tiny go library that aims to make your life a little easier when consuming environment variables in your go programs. 

It extends the `func(string)string` signature with helper methods so you can continue using `os.Getenv` the way you are used to in your workflow. 

It helps you reduce common boilerplate code for handling errors, defaults and converting to other primitive types. 

It only uses pure functions and generics (no reflection) and requires no initialization. 
