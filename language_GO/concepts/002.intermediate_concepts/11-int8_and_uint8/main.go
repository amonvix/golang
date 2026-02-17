package main // Declare the main package.

import "fmt" // Import fmt for formatted I/O.

var a int8 = 127 // Define max int8 value.
var b uint8 = 255 // Define max uint8 value.

func main() { // Program entry point.
	fmt.Printf("%v is of Type of %T\n", a, a) // Print value and type of a.
	fmt.Printf("%v is of Type of %T\n", b, b) // Print value and type of b.
} // End main.
