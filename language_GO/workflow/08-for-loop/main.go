package main // Declare the main package.

import "fmt" // Import fmt for formatted I/O.

func main() { // Program entry point.

	x := 40 // Set x to 40.
	y := 5 // Set y to 5.
	fmt.Printf("\n x=%v \n y=%v \n", x, y) // Print x and y.

	// LOOPS  /  INTERATIVE
	// for statements

	/*
		The Go for loops is similar to-but not the same as-C's.
		It unifies for and while and there is no do-while.
		There are three forms, only one of wich has semicolons

		# Like a C for
		for init; condition; post { }

		# Like a C while
		for condition { }

		# Like a C for (;;)
		for { }

	*/
	//  https://go.dev/doc/effective_go#for

	for i := range 5 { // Range-based loop over 0..4.
		fmt.Printf("Counting to five: %v \t first for loop\n", i) // Print counter.
	} // End for loop.

	for y < 20 { // Loop while y is less than 20.
		fmt.Printf("y is %v \t\t second for loop\n", y) // Print y.
		y++ // Increment y.
	} // End while-style loop.

	//  break
	//  takes you out of the loop
	for { // Infinite loop with break.
		fmt.Printf("y is %v \t\t third for loop\n", y) // Print y.
		if y > 40 { // Check break condition.
			break // Exit the loop.
		} // End if.
		y++ // Increment y.
	} // End infinite loop.

	// continue
	// takes to next iteration
	for i := range 40 { // Range-based loop over 0..39.
		if i%2 != 0 { // Skip odd numbers.
			continue // Continue to next iteration.
		} // End if.
		fmt.Println("Counting even numbers: ", i) // Print even number.
	} // End for loop.

} // End main.
