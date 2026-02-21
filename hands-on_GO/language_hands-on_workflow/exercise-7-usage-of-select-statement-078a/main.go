package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(11)
	y := rand.Intn(11)

	fmt.Printf("The value of x is %v\n", x)
	fmt.Printf("The value of y is %v\n", y)

	// if x < 4 && y < 4 {
	// fmt.Printf("Both are less than 4. The value in x is %v, and in y is %v\n", x, y)
	// } else if x > 6 && y > 6 {
	// fmt.Printf("Both are greater than 6. The value in x is %v, and in y is %v\n", x, y)
	// } else if x >= 4 && x <= 6 {
	// fmt.Printf("x is grater than or equal to 4 and less than or equal to 6. The value in x is %v\n", x)
	// } else if y != 5 {
	// fmt.Printf("y is equal to 5. The value in y is %v\n", y)
	// } else {
	// fmt.Printf("None of the previous cases were met. x is equal to %v, and y is equal to %v\n", x, y)
	// }

	switch {
	case x < 4 && y < 4:
		fmt.Printf("Both are less than 4. The value in x is %v, and in y is %v\n", x, y)
	case x > 6 && y > 6:
		fmt.Printf("Both are greater than 6. The value in x is %v, and in y is %v\n", x, y)
	case x >= 4 && x <= 6:
		fmt.Printf("Both are greater than 6. The value in x is %v, and in y is %v\n", x, y)
	case y != 5:
		fmt.Printf("y is equal to 5. The value in y is %v\n", y)
	default:
		fmt.Printf("None of the previous cases were met. x is equal to %v, and y is equal to %v\n", x, y)
	}

	// Resolution using select
	z := rand.Intn(251)
	fmt.Println(x)

	low := make(chan int, 1)
	mid := make(chan int, 1)
	high := make(chan int, 1)

	if z <= 100 {
		low <- z
	} else if z <= 200 {
		mid <- z
	} else {
		high <- z
	}

	select {
	case z := <-low:
		fmt.Printf("The value of x is %v and it is equal or bellow 100\n", z)

	case z := <-mid:
		fmt.Printf("The value of x is %v and it is equal or bellow 200\n", z)

	case z := <-high:
		fmt.Printf("The value of x is %v and it is above 200\n", z)

	}
}
