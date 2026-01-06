package main // Declare the main package.

import "fmt" // Import fmt for formatted I/O.

func main() { // Program entry point.
	x := 40 // Set x to 40.
	y := 5 // Set y to 5.
	fmt.Printf("\n x=%v \n y=%v \n", x, y) // Print x and y.

	// CONDITIONAL
	//  if statements
	//  switch statements

	switch { // Switch with boolean cases.
	case x < 42: // Match x < 42.
		fmt.Println("x is LESS THAN 42") // Print result.

	case x == 42: // Match x == 42.
		fmt.Println("x is EQUAL TO 42") // Print result.

	case x > 42: // Match x > 42.
		fmt.Println("x is GREATER THAN 42") // Print result.

	default: // Default branch.
		fmt.Println("this is the default for x") // Print default result.
	} // End switch.

	switch x { // Switch on x value.
	case 40: // Match x == 40.
		fmt.Println("x is 40") // Print result.

	case 41: // Match x == 41.
		fmt.Println("x is 41") // Print result.

	case 42: // Match x == 42.
		fmt.Println("x is 42") // Print result.

	case 43: // Match x == 43.
		fmt.Println("x is 43") // Print result.

	default: // Default branch.
		fmt.Println("this is the default case for x") // Print default result.
	} // End switch.

	switch x { // Switch on x with fallthrough.
	case 40: // Match x == 40.
		fmt.Println("x is 40") // Print case message.
		fallthrough // Continue to next case.
	case 41: // Match x == 41.
		fmt.Println("printed because of ALL OF THE fallthrough statements: x is 41") // Print case message.
		fallthrough // Continue to next case.
	case 42: // Match x == 42.
		fmt.Println("printed because of ALL OF THE fallthrough statements: x is 42") // Print case message.
		fallthrough // Continue to next case.
	case 43: // Match x == 43.
		fmt.Println("printed because of ALL OF THE fallthrough statements: x is 43") // Print case message.
		fallthrough // Continue to default.
	default: // Default case.
		fmt.Println("printed because of ALL OF THE fallthrough statements: this is the default case for x") // Print default message.
	} // End switch.

	// no default fallthrough
	switch x { // Switch on x without default fallthrough.
	case 40: // Match x == 40.
		fmt.Println("x is 40") // Print case message.
		fallthrough // Continue to next case.
	case 41: // Match x == 41.
		fmt.Println("printed because of ALL OF THE fallthrough statements: x is 41") // Print case message.

	case 42: // Match x == 42.
		fmt.Println("printed because of ALL OF THE fallthrough statements: x is 42") // Print case message.

	case 43: // Match x == 43.
		fmt.Println("printed because of ALL OF THE fallthrough statements: x is 43") // Print case message.

	default: // Default case.
		fmt.Println("printed because of ALL OF THE fallthrough statements: this is the default case for x") // Print default message.
	} // End switch.

} // End main.
