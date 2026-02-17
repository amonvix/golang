package main // Declare the main package.

import "fmt" // Import fmt for formatted I/O.

func main() { // Program entry point.
	//SEQUENCE
	fmt.Println("this is the first statement to run")
	fmt.Println("this is the second statement to run")
	x := 40                             // Set x to 40.
	y := 5                              // Set y to 5.
	fmt.Printf(" x=%v \n y=%v\n", x, y) // Print x and y.

	// CONDITIONAL
	// if statements
	// switch statements

	if x < 42 && y < 42 {
		// && requries both to be true to run
		fmt.Println("both are less than the meaning of life") // Print result.
	} // End if.

	if x > 30 || x < 42 {
		// || requires one of them to be true to run
		fmt.Println("x is getting closer to the meaning of life") // Print result.
	} // End if.

	if x != 42 && y != 10 {
		// Check if x is NOT 42 AND if is NOT 10.
		fmt.Println("x is not 42") // Print result.
	} // End if.

	/*
		Logical operators
		Logical operators apply to boolean values
		and yield a result of the same type as the operands.
		The right operand is evaluated conditionally.

		&&    conditional AND    p && q  is  "if p then q else false"
		||    conditional OR     p || q  is  "if p then true else q"
		!     NOT                !p      is  "not p"
	*/
	// https://go.dev/ref/spec#Logical_operators

} // End main.
