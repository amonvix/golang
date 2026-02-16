package main // Declare the main package.

import "fmt" // Import fmt for formatted I/O.

func main() { // Program entry point.
	s, i, f := "Sadness", 42, 42.42 // Initialize string, int, and float.
	fmt.Printf("%v is of type %T\n", s, s) // Print value and type of s.
	fmt.Printf("%v is of type %T\n", i, i) // Print value and type of i.
	fmt.Printf("%v is of type %T\n", f, f) // Print value and type of f.
} // End main.
