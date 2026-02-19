package main

import "fmt"

func main() {

	x := 40
	y := 5
	fmt.Printf("\n x=%v \n y=%v \n", x, y)

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

	for i := range 5 {
		fmt.Printf("Counting to five: %v \t first for loop\n", i)
	}

	for y < 20 {
		fmt.Printf("y is %v \t\t second for loop\n", y)
		y++
	}

	//  break
	//  takes you out of the loop
	for {
		fmt.Printf("y is %v \t\t third for loop\n", y)
		if y > 40 {
			break
		}
		y++
	}

	// continue
	// takes to next iteration
	for i := range 40 {
		if i%2 != 0 {
			continue
		}
		fmt.Println("Counting even numbers: ", i)
	}

}
