package main // Declare the main package.

import ( // Start import block.
	"fmt"       // Import fmt for formatted I/O.
	"math/rand" // Import rand for random numbers.
) // End import block.

func main() { // Program entry point.

	x := 40 // Set x to 40.
	y := 5 // Set y to 5.
	fmt.Printf("\n x=%v \n y=%v \n", x, y) // Print x and y.

	/*
		The expression [evaluated in an if statemet ]may be preceded by a simple statement,
		wich executes before the experesssion is evaluated.
	*/

	// https://go.dev/ref/spec#If_statements

	if z := 2 * rand.Intn(x); z >= x { // Initialize z and compare with x.
		fmt.Printf("z is %v and that is GREATER THAN OR EQUAL x wich is %v\n", z, x) // Print result.
	} else { // Handle z < x.
		fmt.Printf("z is %v and that is LESS THAN x wich is %v\n", z, x) // Print result.
	} // End if/else.

	if x := 42; x > 40 { // Initialize x and compare to 40.
		fmt.Println("This is true that x is greater than 40!") // Print result.
	} // End if.

	// does not run
	// fmt.Println(x)

	/*
		The comma ok idiom is also like this
	*/

	// 	https:// go.dev/play/p/OXGzjxVkag0

	m := make(map[string]int) // Create a map of names to values.
	m["James"] = 42 // Store James.
	m["Moneypenny"] = 32 // Store Moneypenny.
	fmt.Println("James is ", m["James"]) // Print James value.
	fmt.Println("Miss Moneypenny is ", m["Moneypenny"]) // Print Moneypenny value.
	fmt.Println("Q is ", m["Q"]) // Print missing key value.

	// comma ok idiom
	if v, ok := m["James"]; ok { // Check if James exists.
		fmt.Println("James is ", v) // Print value.
	} // End if.

	if v, ok := m["Moneypenny"]; ok { // Check if Moneypenny exists.
		fmt.Println("Miss Moneypenny is ", v) // Print value.
	} // End if.

	// will not print
	if v, ok := m["Q"]; ok { // Check if Q exists.
		fmt.Println("Q is ", v) // Print value.
	} // End if.
} // End main.
