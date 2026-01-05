package main // Declare the main package.

import "fmt" // Import fmt for formatted I/O.

func main() { // Program entry point.

	x := 40 // Set x to 40.
	y := 5 // Set y to 5.
	fmt.Printf("\n x=%v \n y=%v\n", x, y) // Print x and y.

	if x < 42 && y < 42 { // Check if both x and y are less than 42.
		fmt.Println("both are less than the meaning of life") // Print result.
	} // End if.

	if x > 30 || x < 42 { // Check if x is above 30 or below 42.
		fmt.Println("x is getting closer to the meaning of life") // Print result.
	} // End if.

	if x != 42 { // Check if x is not 42.
		fmt.Println("x is not 42") // Print result.
	} // End if.
} // End main.
