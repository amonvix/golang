package main // Declare the main package.

import ( // Start import block.
	"fmt"       // Import fmt for formatted I/O.
	"math/rand" // Import rand for random numbers.
	"time"      // Import time for time utilities.
) // End import block.

func main() { // Program entry point.
	x := 40 // Set x to 40.
	y := 5 // Set y to 5.
	fmt.Printf("\n x=%v \n y=%v \n", x, y) // Print x and y.

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
		ch1 <- 41 // Send value to ch1.
	}() // End goroutine.

	go func() { // Start goroutine for ch2.
		time.Sleep(d2 * time.Millisecond) // Sleep before sending.
		ch2 <- 42 // Send value to ch2.
	}() // End goroutine.

	// A select statement chooses which of a set of possible send or receive operations will proceed.
	// It looks similar to a switch statement but with the cases all referring to communication operations.
	// https://go.dev/ref/spec#select_statements
	select { // Wait for a channel receive.
	case v1 := <-ch1: // Receive from ch1.
		fmt.Printf("received %v from ch1\n", v1) // Print received value.
	case v2 := <-ch2: // Receive from ch2.
		fmt.Printf("received %v from ch2\n", v2) // Print received value.
	} // End select.
} // End main.
