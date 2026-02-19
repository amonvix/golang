package main

import "fmt"

// main is the entry point of the application.
// It demonstrates variable declaration, initialization, and multiple assignment with ignored values.
func main() {
	a := 42
	fmt.Println(a)

	// Multiple assignment with one ignored value (underscore).
	// The underscore is used to discard the value 3, which is not needed.
	b, c, d, _, f := 0, 1, 2, 3, "happiness"
	fmt.Println(b, c, d, f)

	// The following commented-out code illustrates that redeclaring variables with :=
	// in the same scope without using all variables is invalid.
	/*
		b, c, d, e := 0, 1, 2, 3
		fmt.Println(b, c, d)
	*/

	// The following commented-out code shows explicit variable declaration and initialization.
	// Declaring a variable without initialization sets it to the zero value of its type.
	/*
		var g int
		fmt.Println(g)
		g = 43
		fmt.Println(g)

		// Declaring and initializing a variable in one statement.
		var h int = 44
		fmt.Println(h)
	*/
}