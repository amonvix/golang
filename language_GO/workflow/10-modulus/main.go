package main

import "fmt"

func main() {
	// MODULUS OPERATOR
	// The modulus operator (%) returns the remainder of a division operation.
	// It is commonly used to determine if a number is even or odd, or to cycle through a set of values.

	x := 83 / 40
	y := 83 % 40
	fmt.Println(x, y)

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Printf("%v is even\n", i)
		} else {
			fmt.Printf("%v is odd\n", i)
		}
	}
}
