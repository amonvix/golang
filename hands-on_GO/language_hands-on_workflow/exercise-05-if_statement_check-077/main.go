package main

/*
Rules to be followed:
# Create 2 random ints between 0 and 10 exclusive
	# assign them to variables with the identifiers x and y

# Print their values and in same print statement
	# use an if statement to print the correct description:
		- x and y are both less than 4
		- x and y are both greater than 6
		- x is grater than or equal to 4 and less than or equal to 6
		- y is not 5
		- none of the previous cases were met
*/

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
		fmt.Printf("y is not equal to 5. The value in y is %v\n", y)
	} else {
		fmt.Printf("None of the previous cases were met. x is equal to %v, and y is equal to %v\n", x, y)
	}
}
