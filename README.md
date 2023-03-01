# Excercises in Go

This repository contains a collection of exercises and reference material for [Go](https://go.dev/).

See [Syntax Cheatsheet](syntaxCheatsheet.go) for all Go syntax on 1 page.

Excercises are seperated by directories named using the following naming convention `<number>_<ExcerciseName>`

* Excercises are adapted from [Gophercises](https://gophercises.com/)
* Solid basics introduction can be found at [Go By Example](https://gobyexample.com/)

## Biased opinion

> Go is simply amazing! Think Python simplicity with C++ performance, Rust memory safety, no Java BS and no JavaScript esotericism. Testing, documentation and package management ships with the language. Also compiles to binary. What more can you possibly want? Composition over Inheritance every time baby! I love it! Let's Go!

## Quick Go types Cheatsheet 

<table><thead><tr><th>Type</th><th>Description</th></tr></thead><tbody><tr><td><code>bool</code></td><td>Boolean, either <code>true</code> or <code>false</code></td></tr><tr><td><code>int</code></td><td>Signed integer, either 32 or 64 bits</td></tr><tr><td><code>uint</code></td><td>Unsigned integer, either 32 or 64 bits</td></tr><tr><td><code>float32</code></td><td>IEEE-754 32-bit floating-point number</td></tr><tr><td><code>float64</code></td><td>IEEE-754 64-bit floating-point number</td></tr><tr><td><code>complex64</code></td><td>Complex number with float32 real and imaginary parts</td></tr><tr><td><code>complex128</code></td><td>Complex number with float64 real and imaginary parts</td></tr><tr><td><code>byte</code></td><td>Alias for <code>uint8</code>, represents a byte</td></tr><tr><td><code>rune</code></td><td>Alias for <code>int32</code>, represents a Unicode code point</td></tr><tr><td><code>string</code></td><td>A string of characters</td></tr><tr><td><code>error</code></td><td>An interface type that represents an error condition</td></tr><tr><td><code>uintptr</code></td><td>An unsigned integer that is large enough to hold the bit pattern of any pointer</td></tr></tbody></table>

## OO in Go Gotchas

- there are no classes in Go
- there is only `struct` type that have the same function
- there is no concept of inheritance in Go
- structs can be composed (composition over inheritance)

## Interfaces Gotchas

> instead of designing our abstractions in terms of what kind of data our types can hold, we design our abstractions in terms of what actions our types can execute. -[Jordan Orel](https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go)

- no `implements` keyword 
- if given `struct` has all methods from given interface defined it implements it 
- every type satisfy at least an empty interface `interface{}`

you can pass anything to the function below, but v is not `any` type but `interface{}` type
```Go
func myFunc(v interface{}){
    ...
}

```

everything in Go is passed by value