package main

import (
	"fmt"
	"math/rand"
)

// Create a randon int btween 0 and 5
// assign the value to a variable with the identifier x
// use a switch statement to print a description of the variable and value
// Run the code 42 times and print the iteration number

func main() {
	for i := 0; i < 43; i++ {

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
