package main

// Rule for the exercise: Modify the previous program to use switch statement

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(11)
	y := rand.Intn(11)

	fmt.Printf("The value of x is %v\n", x)
	fmt.Printf("The value of y is %v\n", y)

	if x < 4 && y < 4 {
		fmt.Printf("Both are less than 4. The value in x is %v, and in y is %v\n", x, y)
	} else if x > 6 && y > 6 {
		fmt.Printf("Both are greater than 6. The value in x is %v, and in y is %v\n", x, y)
	} else if x >= 4 && x <= 6 {
		fmt.Printf("x is grater than or equal to 4 and less than or equal to 6. The value in x is %v\n", x)
	} else if y != 5 {
		fmt.Printf("y is equal to 5. The value in y is %v\n", y)
	} else {
		fmt.Printf("None of the previous cases were met. x is equal to %v, and y is equal to %v\n", x, y)
	}

}
