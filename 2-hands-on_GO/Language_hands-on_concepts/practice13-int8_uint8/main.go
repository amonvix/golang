package main

// Let's see some examples of int8 and uint8, and their differences

import "fmt"

var a int8 = 127
var b uint8 = 255

func main() {
	fmt.Printf("%v is of Type of %T\n", a, a)
	fmt.Printf("%v is of Type of %T\n", b, b)
}
