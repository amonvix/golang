package main

import "fmt"

const (
	first  = 747
	second = 911
	third  = 90210
)

/*
the values in the variable above need to be converted to
#decimal
#binary
#hexadecimal
*/

func main() {
	fmt.Printf("\n%v in decimal is %d, in binary is %b and in hexadecimal is %#x\n", first, first, first, first)
	fmt.Printf("\n%v in decimal is %d, in binary is %b and in hexadecimal is %#x\n", second, second, second, second)
	fmt.Printf("\n%v in decimal is %d, in binary is %b and in hexadecimal is %#x\n", third, third, third, third)
}
