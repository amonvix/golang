package main

import "fmt"

func main() {
	// SEQUENCE
	x := 40
	y := 5
	fmt.Printf("\n x=%v \n y=%v \n", x, y)

	// NESTED LOOPS
	for i := 0; i < 5; i++ {
		fmt.Println("--")
		for j := 0; j < 5; j++ {
			fmt.Printf("outer loop %v \t  inner loop %v\n", i, j)
		}
	}
}
