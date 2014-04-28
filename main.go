package main

import (
    "fmt"
    "./calc"
    "./greet"
)

func main() {
    // Makes two calls to our GoAdd
    // wrapped function
    a := calc.GoAdd(5, 2)
    b := calc.GoAdd(10, 8)

    // Makes a call to our GoSubtract
    // wrapped function
    c := calc.GoSubtract(b, a)

    fmt.Println(c)

    // Makes a call to our SayHello
    // wrapped function
    greeting := greet.SayHello("Harry")

    fmt.Println(greeting)
}