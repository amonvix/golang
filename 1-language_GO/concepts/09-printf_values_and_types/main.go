package main

import "fmt"

// main is the entry point of the application.
// It initializes variables of different types and outputs their values and types using formatted printing.
func main() {
	s, i, f := "Sadness", 42, 42.42
	fmt.Printf("%v is of type %T\n", s, s)
	fmt.Printf("%v is of type %T\n", i, i)
	fmt.Printf("%v is of type %T\n", f, f)
}