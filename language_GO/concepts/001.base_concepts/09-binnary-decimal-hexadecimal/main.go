package main // Declare the main package.

import "fmt" // Import fmt for formatted I/O.

const ( // Start constant block.
	first  = 747   // Define first constant.
	second = 911   // Define second constant.
	third  = 90210 // Define third constant.
) // End constant block.

/*
the values in the variable above need to be converted to
#decimal
#binary
#hexadecimal
*/

func main() { // Program entry point.
	fmt.Printf("\n%v in decimal is %d, in binary is %b and in hexadecimal is %#x\n", first, first, first, first) // Print formats for first.
	fmt.Printf("\n%v in decimal is %d, in binary is %b and in hexadecimal is %#x\n", second, second, second, second) // Print formats for second.
	fmt.Printf("\n%v in decimal is %d, in binary is %b and in hexadecimal is %#x\n", third, third, third, third) // Print formats for third.
} // End main.
