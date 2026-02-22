package main

import (
	"fmt"
)

// create a lop using only a condition
// increment or decrement the variable being checked in the condition

func main() {
	fmt.Printf("\n\nAcress\n\n")
	var i int
	for i < 11 {
		fmt.Printf("%v é o valor da %v° iteration\n", i, i)
		i++
	}

	fmt.Printf("\n\nDecress\n\n")
	j := 20
	for j >= 0 {
		fmt.Printf("%v é o valor da %v° iteration\n", j, j)
		j--
	}
}
