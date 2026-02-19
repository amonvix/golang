package main // Declare the main package.

// the imported package "fmt"
// is in the "file block" scope
// of this file

import ( // Start import block.
	"fmt" // Import fmt for formatted I/O.

	"github.com/amonvix/golang/language_GO/concepts/002.intermediate_concepts/05-scope/furtherexplored" // Import furtherexplored package.
) // End import block.

// x is in "package block" scope
var x = 42 // Define package-level x.

func main() { // Program entry point.
	// x can be accessed here
	fmt.Println(x) // Print package-level x.

	// x can be accessed here
	printMe() // Call printMe.

	// y does not exist here yet
	// so this won't work
	// fmt.Println(y)

	// y is in a "block scope"
	y := 43        // Define block-scoped y.
	fmt.Println(y) // Print y.

	p1 := person{ // Create a person.
		"James", // Set first name.
	} // End person literal.
	p1.sayHello() // Call sayHello on p1.

	// variable "shadowing"
	x := 32        // Shadow package-level x.
	fmt.Println(x) // Print shadowed x.

	// package clock x is still the sane
	printMe() // Call printMe again.

	furtherexplored.Fascinating() // Call Fascinating from furtherexplored.

	// also good to know

	if z := 82; z > 50 { // Define z and compare.
		fmt.Println(z) // Print z.
	} // End if.
	// you can't access z here
	// fmt.Println(z)
}

func printMe() { // Print the package-level x.
	// x can be accessed here
	fmt.Println(x) // Print package-level x.
} // End printMe.

// type person is in "package block" scope
type person struct { // Define person type.
	first string // Store first name.
} // End person type.

// the method sayHello
// witch is attached to values of TYPE person
// is in "package block" socope

func (p person) sayHello() { // Greet using the person's name.
	fmt.Println("Hi, my name is, ", p.first) // Print greeting.
} // End sayHello.
