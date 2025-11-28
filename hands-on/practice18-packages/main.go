package main

import "fmt"

var y = 43

const z = 44

func main() {
	x := 42
	fmt.Printf("The value of x is %v, in a type of %T\n", x, x)
	fmt.Printf("The value of y is %v, in a type of %T\n", y, y)
	fmt.Printf("The value of z is %v, in a type of %T\n", z, z)
}
