package main

// the imported package "fmt"
// is in the "file block" scope
// of this file

import (
	"fmt"

	"github.com/amonvix/golang/hands-on/practice14-scope/furtherexplored"
)

// x is in "package block" scope
var x = 42

func main() {
	// x can be accessed here
	fmt.Println(x)

	// x can be accessed here
	printMe()

	// y does not exist here yet
	// so this won't work
	// fmt.Println(y)

	// y is in a "block scope"
	y := 43
	fmt.Println(y)

	p1 := person{
		"James",
	}
	p1.sayHello()

	// variable "shadowing"
	x := 32
	fmt.Println(x)

	// package clock x is still the sane
	printMe()

	furtherexplored.Fascinating()

	// also good to know

	if z := 82; z > 50 {
		fmt.Println(z)
	}
	// you can't access z here
	// fmt.Println(z)
}

func printMe() {
	// x can be accessed here
	fmt.Println(x)
}

// type person is in "package block" scope
type person struct {
	first string
}

// the method sayHello
// witch is attached to values of TYPE person
// is in "package block" socope

func (p person) sayHello() {
	fmt.Println("Hi, my name is, ", p.first)
}
