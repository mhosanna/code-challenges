package main

import (
    "fmt"

    "github.com/mhosanna/code-challenges/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}