package main

import "fmt"

func main() {
	// NESTED LOOPS
	for i := 0; i < 5; i++ {
		fmt.Println("--")
		for j := 0; j < 5; j++ {
			fmt.Printf("outer loop %v \t  inner loop %v\n", i, j)
		}
	}

	//  New way to create nested loops using ranges
	for i := range 5 {
		fmt.Println("--")
		for j := range 5 {
			fmt.Printf("outer loop %v \t  inner loop %v\n", i, j)
		}
	}
}
