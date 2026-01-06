package main

// Now prove that you compreend the system binary and hexadecimal givem exemples of each one

import "fmt"

func main() {
	number := 42
	fmt.Printf("42 as binary, %b \n", number)
	fmt.Printf("42 as hexadecimal, %#x \n", number)

	// print this values as both binary & hexadeimal
	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5

	//Solution to exercise proposed
	fmt.Printf("%v \t %b \t %#x \n", a, a, a)
	fmt.Printf("%v \t %b \t %#x \n", b, b, b)
	fmt.Printf("%v \t %b \t %#x \n", c, c, c)
	fmt.Printf("%v \t %b \t %#x \n", d, d, d)
	fmt.Printf("%v \t %b \t %#x \n", e, e, e)
	fmt.Printf("%v \t %b \t %#x \n", f, f, f)

}
