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

	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and that is GREATER THAN OR EQUAL x wich is %v\n", z, x)
	} else {
		fmt.Printf("z is %v and that is LESS THAN x wich is %v\n", z, x)
	}

	if x := 42; x > 40 {
		fmt.Println("This is true that x is greater than 40!")
	}

	// does not run
	// fmt.Println(x)

	/*
		The comma ok idiom is also like this
	*/

	// 	https:// go.dev/play/p/OXGzjxVkag0

	m := make(map[string]int)
	m["James"] = 42
	m["Moneypenny"] = 32
	fmt.Println("James is ", m["James"])
	fmt.Println("Miss Moneypenny is ", m["Moneypenny"])
	fmt.Println("Q is ", m["Q"])

	// comma ok idiom
	if v, ok := m["James"]; ok {
		fmt.Println("James is ", v)
	}

	if v, ok := m["Moneypenny"]; ok {
		fmt.Println("Miss Moneypenny is ", v)
	}

	// will not print
	if v, ok := m["Q"]; ok {
		fmt.Println("Q is ", v)
	}
}
