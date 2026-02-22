package main

import (
	"fmt"
)

// create a lop that uses break statement
// increment or decrement the variable being checked as a condition in the loop

func main() {

	x := 42

	for {
		fmt.Printf("x is %v \t\t keep it rolling until reach perfection²\n", x)
		if x > 83 {
			break
		}
		x++
	}

	fmt.Printf("\n\n\t\t\t***********\t\t\t\n\n")

	for {
		fmt.Printf("x is %v \t\t let's restore the perfection\n", x)
		if x < 43 {
			break
		}
		x--
	}
}
