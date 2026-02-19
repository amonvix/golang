package main

import "fmt"

var zero int

func main() {
	fmt.Println(zero)

	z := 42
	fmt.Println(z)

	a, b := 43, "golang"
	fmt.Println(a, b)

	var x float32 = 42.42
	// Print the value and its concrete type to verify type correctness.
	fmt.Printf("%v is of this Type of %T\n", x, x)

	e, f, _ := 45, 46, 47
	// The third value is intentionally discarded as it is not needed.
	fmt.Println(e, f)
}