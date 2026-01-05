package main // Declare the main package.

import "fmt" // Import fmt for formatted I/O.

func main() { // Program entry point.
	// MODULUS OPERATOR
	// The modulus operator (%) returns the remainder of a division operation.
	// It is commonly used to determine if a number is even or odd, or to cycle through a set of values.

	x := 83 / 40 // Compute integer division result.
	y := 83 % 40 // Compute remainder with modulus.
	fmt.Println(x, y) // Print quotient and remainder.

	for i := 0; i < 100; i++ { // Loop from 0 to 99.
		if i%2 == 0 { // Check if i is even.
			fmt.Printf("%v is even\n", i) // Print even message.
		} else { // Handle odd numbers.
			fmt.Printf("%v is odd\n", i) // Print odd message.
		} // End if/else.
	} // End for loop.
} // End main.
