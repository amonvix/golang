package main

import "fmt"

const (
	first  = 747   // first is an integer constant used for formatted output in multiple bases.
	second = 911   // second is an integer constant used for formatted output in multiple bases.
	third  = 90210 // third is an integer constant used for formatted output in multiple bases.
)

// main prints the decimal, binary, and hexadecimal representations of predefined constants.
// Output formatting uses verbs to display values in multiple numeric bases for clarity.
func main() {
	fmt.Printf("\n%v in decimal is %d, in binary is %b and in hexadecimal is %#x\n", first, first, first, first)
	fmt.Printf("\n%v in decimal is %d, in binary is %b and in hexadecimal is %#x\n", second, second, second, second)
	fmt.Printf("\n%v in decimal is %d, in binary is %b and in hexadecimal is %#x\n", third, third, third, third)
}