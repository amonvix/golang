package main // Declare the main package.

import "fmt" // Import fmt for formatted I/O.

func main() { // Program entry point.
	// NESTED LOOPS
	for i := 0; i < 5; i++ { // Outer loop counter.
		fmt.Println("--") // Print separator.
		for j := 0; j < 5; j++ { // Inner loop counter.
			fmt.Printf("outer loop %v \t  inner loop %v\n", i, j) // Print indices.
		} // End inner loop.
	} // End outer loop.

	//  New way to create nested loops using ranges
	for i := range 5 { // Range-based outer loop.
		fmt.Println("--") // Print separator.
		for j := range 5 { // Range-based inner loop.
			fmt.Printf("outer loop %v \t  inner loop %v\n", i, j) // Print indices.
		} // End inner loop.
	} // End outer loop.
} // End main.
