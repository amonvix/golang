package main

import (
	"fmt"
	"math/rand"
)

/*
Rules to be followed:
# Create one random int between 0 inclusive and 5 exclusive
	# assign the value to a variable with the identifier x

# Use a switch statement to print a description of the variable and the value
# Run the code 42 times and print the iteration number
*/

func main() {
	for i := 0; i < 42; i++ {
		x := rand.Intn(5)
		switch x {
		case 0:
			fmt.Printf("The iteration %v\t received the number %v\n", i, x)
		case 1:
			fmt.Printf("The iteration %v\t received the number %v\n", i, x)
		case 2:
			fmt.Printf("The iteration %v\t received the number %v\n", i, x)
		case 3:
			fmt.Printf("The iteration %v\t received the number %v\n", i, x)
		case 4:
			fmt.Printf("The iteration %v\t received the number %v\n", i, x)
		}
	}
}
