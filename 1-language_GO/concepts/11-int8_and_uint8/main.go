package main

import "fmt"

var a int8 = 127
var b uint8 = 255

func main() {
	// Print the value and type of the signed 8-bit integer variable 'a'.
	fmt.Printf("%v is of Type of %T\n", a, a)
	// Print the value and type of the unsigned 8-bit integer variable 'b'.
	fmt.Printf("%v is of Type of %T\n", b, b)
}