package main

// Show ways to use a formated print

import "fmt"

func main() {
	s, i, f := "Sadness", 42, 42.42
	fmt.Printf("%v is of type %T\n", s, s)
	fmt.Printf("%v is of type %T\n", i, i)
	fmt.Printf("%v is of type %T\n", f, f)
}
