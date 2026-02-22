package main // Declare the main package.

import ( // Start import block.
	"fmt"       // Import fmt for formatted I/O.
	"math/rand" // Import rand for random numbers.
	"time"      // Import time for time utilities.
) // End import block.

func init() { // Run initialization before main.
	fmt.Println("Begin initialization") // Print init message.
} // End init.

func main() { // Program entry point.
	// SEQUENCE
	fmt.Println("this is the first statement to run")  // Print the first message.
	fmt.Println("this is the second statement to run") // Print the second message.
	x := 40                                            // Set x to 40.
	y := 5                                             // Set y to 5.
	fmt.Println()                                      // Print a blank line.
	fmt.Printf("\n x=%v \n y=%v", x, y)                // Print x and y.

	// x := 40
	// y := 5
	// fmt.Printf("\n x=%v \n y=%v \n", x, y)

	// CONDITIONAL
	//  if statements
	//  switch statements

	if x < 42 { // Check if x is less than 42.
		fmt.Println("Less than the meaning of life") // Print branch message.
	} // End if.
	if x < 42 { // Check if x is less than 42.
		fmt.Println("Less than the meaning of life") // Print branch message.
	} else { // Handle x >= 42.
		fmt.Println("equal to, or greater than, the meaning of life") // Print branch message.
	} // End if/else.

	if x < 42 { // Check if x is less than 42.
		fmt.Println("Less than the meaning of life") // Print branch message.
	} else if x == 42 { // Check if x equals 42.
		fmt.Println("equal to the meaning of life") // Print branch message.
	} else { // Handle x > 42.
		fmt.Println("Greater than, the meaning of life") // Print branch message.
	} // End if/else if/else.
	/*
		"If" statements specify the conditional execution of two branches
		according to the value of boolean expression. If the expression evaluates
		to true, the if" branch is executed, otherwise, if present, the else "branch" is executed
	*/
	//  https://go.dev/ref/spec#If_statements

	/*

		Comparison operators
		Comparison operators comre to operands and yield an untyped boolean value.

		==		equal
		!=		not equal
		<		less
		<=		less or equal
		>		grater
		>=		greater or equal
	*/
	// https://go.dev/ref/spec#Comparison_operators

	if x < 42 && y < 42 { // Check if both x and y are less than 42.
		fmt.Println("both are less than the meaning of life") // Print combined result.
	} // End if.

	if x > 30 || x < 42 { // Check if x is above 30 or below 42.
		fmt.Println("x is getting closer to the meaning of life") // Print result.
	} // End if.

	if x != 42 { // Check if x is not 42.
		fmt.Println("x is not 42") // Print result.
	} // End if.

	/*
		Logical operators
		Logical operators apply to boolean values
		and yeld a result of the same type as the operands.
		The rigth operand is evaluated conditionally

		&&	conditional AND		p && q	is 	"if p then q else false"
		||	conditional OR		p || q	is 	"if p then true else q"
		!	NOT					!p		is	"not p"
	*/

	// https://go.dev/ref/spec#Logical_operators

	/*
		The expression [evaluated in an if statemet ]may be preceded by a simple statement,
		wich executes before the experesssion is evaluated.
	*/

	// https://go.dev/ref/spec#If_statements

	/*
		The comma ok idiom is also like this
	*/

	// 	https:// go.dev/play/p/OXGzjxVkag0

	if z := 2 * rand.Intn(x); z >= x { // Generate z and compare with x.
		fmt.Printf("z is %v and that is GREATER THAN OR EQUAL x wich is %v\n", z, x) // Print result.
	} else { // Handle z < x.
		fmt.Printf("z is %v and that is LESS THAN x wich is %v\n", z, x) // Print result.
	} // End if/else.

	switch {
	case x < 42:
		fmt.Println("x is LESS THAN 42")
	case x == 42:
		fmt.Println("x is EQUAL TO 42")
	case x > 42:
		fmt.Println("x is GREATER THAN 42")
	default:
		fmt.Println("this is the default case for x")
	}

	switch x { // Switch on x.
	case 40: // Match x == 40.
		fmt.Println("x is 40") // Print case message.
		fallthrough            // Continue to next case.
	case 41: // Match x == 41.
		fmt.Println("printed because of ALL OF THE fallthrough statements: x is 41") // Print case message.
		fallthrough                                                                  // Continue to next case.
	case 42: // Match x == 42.
		fmt.Println("printed because of ALL OF THE fallthrough statements: x is 42") // Print case message.
		fallthrough                                                                  // Continue to next case.
	case 43: // Match x == 43.
		fmt.Println("printed because of ALL OF THE fallthrough statements: x is 43") // Print case message.
		fallthrough                                                                  // Continue to default.
	default: // Default case.
		fmt.Println("printed because of ALL OF THE fallthrough statements: this is the default case for x") // Print default message.
	} // End switch.

	// no default fallthrough
	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("printed because of fallthrough: x is 41")
	case 42:
		fmt.Println("printed because of fallthrough: x is 42")
	case 43:
		fmt.Println("printed because of fallthrough: x is 43")
	default:
		fmt.Println("printed because of fallthrough: this is the default case for x")
	}

	// 	concurrency
	// 	switch on channel

	ch1 := make(chan int) // Create the first int channel.
	ch2 := make(chan int) // Create the second int channel.

	r := rand.New(rand.NewSource(time.Now().UnixNano())) // Create a seeded RNG.
	// get an int64, convert it to type time.Duration, then use it in time.Sleep
	// Int63n returns an int64
	// type duration's underlying type is int64
	// time.Sleep takes type duration
	// rime.Millisecond is a constant in the time package
	// https://pkg.go.dev/time#pkg-constants

	d1 := time.Duration(r.Int63n(250)) // Generate duration for first goroutine.
	d2 := time.Duration(r.Int63n(250)) // Generate duration for second goroutine.
	// fmt.Printf("%v \t %T\n", d1, d1)
	// time.Sleep(d1 * time.Millisecond)
	// fmt.Println("Done")

	go func() { // Start goroutine for ch1.
		time.Sleep(d1 * time.Millisecond) // Sleep before sending.
		ch1 <- 41                         // Send value to ch1.
	}() // End goroutine.

	go func() { // Start goroutine for ch2.
		time.Sleep(d2 * time.Millisecond) // Sleep before sending.
		ch2 <- 42                         // Send value to ch2.
	}() // End goroutine.

	// A "select" statement chooses which of a set of possible send or receive operations will proceed.
	// It looks similar to a "switch" statement but with the cases all referring to communication operations.
	// https://go.dev/ref/spec#Select_statements
	select {
	case v1 := <-ch1:
		fmt.Println("value from channel 1", v1)
	case v2 := <-ch2:
		fmt.Println("value from channel 2", v2)
	}

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

	for i := 0; i < 5; i++ { // Count from 0 to 4.
		fmt.Printf("Counting to five: %v \t first for loop\n", i) // Print counter.
	} // End for loop.

	for y < 10 { // Loop while y is less than 10.
		fmt.Printf("y is %v \t\t\t second for loop\n", y) // Print y.
		y++                                               // Increment y.
	} // End while-style loop.

	//  break
	//  takes you out of the loop

	for {
		fmt.Printf("y is %v \t\t third  for loop\n", y)
		if y > 20 {
			break
		}
		y++
	}

	// continue
	// takes to next iteration
	for i := 0; i < 20; i++ { // Loop from 0 to 19.
		if i%2 != 0 { // Skip odd numbers.
			continue // Continue to next iteration.
		} // End if.
		fmt.Println("Counting even numbers: ", i) // Print even number.
	} // End for loop.

	//  nested loops
	for i := 0; i < 5; i++ { // Outer loop.
		fmt.Println("--")        // Print separator.
		for j := 0; j < 5; j++ { // Inner loop.
			fmt.Printf("outer loop %v \t  inner loop %v\n", i, j) // Print indices.
		} // End inner loop.
	} // End outer loop.

	//  for range loop
	//  data structures - slices
	xi := []int{42, 43, 44, 45, 46, 47}
	for i, v := range xi {
		fmt.Println("ranging over a slice", i, v)
	}

	// for range loop
	// data structures - map
	m := map[string]int{ // Create a map of names to values.
		"James":      42, // Add James.
		"Moneypenny": 32, // Add Moneypenny.
	} // End map literal.

	for k, v := range m { // Range over the map.
		fmt.Println("ranging over a map", k, v) // Print key and value.
	} // End range loop.

} // End main.
