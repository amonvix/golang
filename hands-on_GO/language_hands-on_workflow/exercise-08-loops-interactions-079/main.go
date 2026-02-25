package main

// Create a program that has a loop that prints every number from 0 to 99
// Than modify the program in exercise 7 to run 100 times

import (
	"fmt"
	"math/rand"
)

func main() {
	for c := 0; c < 100; c++ {
		fmt.Println(c)
	}

	for i := 0; i < 101; i++ {
		x := rand.Intn(11)
		y := rand.Intn(11)

		fmt.Printf("Iteration %v -> x and y are %v and %v\t", i, x, y)

		switch {
		case x < 4 && y < 4:
			fmt.Printf("Both are less than 4. The value in x is %v, and in y is %v\n", x, y)
		case x > 6 && y > 6:
			fmt.Printf("Both are greater than 6. The value in x is %v, and in y is %v\n", x, y)
		case x >= 4 && x <= 6:
			fmt.Printf("Both are greater than 6. The value in x is %v, and in y is %v\n", x, y)
		case y != 5:
			fmt.Printf("y are not equal to 5. The value in y is %v\n", y)
		default:
			fmt.Printf("None of the previous cases were met. x is equal to %v, and y is equal to %v\n", x, y)
		}
	}
}
