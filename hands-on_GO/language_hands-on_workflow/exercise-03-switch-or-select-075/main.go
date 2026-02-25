package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Resolution using switch
	x := rand.Intn(251)
	fmt.Println(x)

	switch {
	case x <= 100:
		fmt.Printf("The value of x is %v and it is equal or bellow 100\n", x)

	case x <= 200:
		fmt.Printf("the value of x %v is between 101 and 200\n", x)

	case x <= 250:
		fmt.Printf("the value of x %v is between 201 and 250\n", x)
	}

	// Resolution using select

	low := make(chan int, 1)
	mid := make(chan int, 1)
	high := make(chan int, 1)

	if x <= 100 {
		low <- x
	} else if x <= 200 {
		mid <- x
	} else {
		high <- x
	}

	select {
	case v := <-low:
		fmt.Printf("the value of x is %v and it is equal or bellow 100\n", v)
	case v := <-mid:
		fmt.Printf("the value of x is %v and it is bellow 200\n", v)
	case v := <-high:
		fmt.Printf("the value of x is %v and it is bellow 250\n", v)
	}

}
