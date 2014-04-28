Very Simple Go bindings for C
=============================

This code was just written as a reference for myself
while I am learning how to interface C functions
in Go.

There are currently two "C libraries" in these examples,
a simple `calc` library and a `greet` library.

C source files are in both the calc and greet directories
along with their Go wrappers, which are then pulled into
`main.go`. This file then makes calls to the wrapped Go
functions that are defined in C. It demonstrates how the
types are converted from C to Go.

I will probably add more examples with different types
at some point.