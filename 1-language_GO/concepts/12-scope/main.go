package main

import (
	"fmt"
	"golang/1-language_GO/concepts/12-scope/furtherexplored"
)

// x holds a package-level integer accessible throughout this package.
var x = 42

func main() {
	fmt.Println(x) // Access package-level x.

	printMe() // Prints package-level x.

	// y is block-scoped to main function.
	y := 43
	fmt.Println(y)

	p1 := person{
		"James",
	}
	p1.sayHello() // Invoke method on person value.

	// Shadow package-level x with a new variable scoped to main.
	x := 32
	fmt.Println(x)

	// printMe still accesses package-level x, unaffected by shadowing.
	printMe()

	furtherexplored.Fascinating() // Invoke external package function.

	// z is scoped to the if statement block.
	if z := 82; z > 50 {
		fmt.Println(z)
	}
	// z is not accessible here.
}

// printMe outputs the package-level x variable.
func printMe() {
	fmt.Println(x)
}

// person represents an individual with a first name.
type person struct {
	first string
}

// sayHello prints a greeting including the person's first name.
func (p person) sayHello() {
	fmt.Println("Hi, my name is, ", p.first)
}
