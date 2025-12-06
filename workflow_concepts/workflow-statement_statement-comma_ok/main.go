package main

import (
	"fmt"
	"math/rand"
)

func main() {

	x := 40
	y := 5
	fmt.Printf("\n x=%v \n y=%v \n", x, y)

	/*
		The expression [evaluated in an if statemet ]may be preceded by a simple statement,
		wich executes before the experesssion is evaluated.
	*/

	// https://go.dev/ref/spec#If_statements

	/*
		The comma ok idiom is also like this
	*/

	// 	https:// go.dev/play/p/OXGzjxVkag0

	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and that is GREATER THAN OR EQUAL x wich is %v\n", z, x)
	} else {
		fmt.Printf("z is %v and that is LESS THAN x wich is %v\n", z, x)
	}

}
