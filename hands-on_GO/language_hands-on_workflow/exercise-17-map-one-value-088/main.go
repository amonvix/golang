package main

import (
	"fmt"
	"math/rand"
)

/*
use the "statement statement" idiom to
		initialize x with a random int between 0 inclusive and 5 exclusive
		if x is 3
			print "x is 3"
run that code 100 times
what's the benefit of using the "statement statement" idiom?
*/

func main() {
	for range 99 {
		if x := rand.Intn(5); x == 3 {
			fmt.Println(x)
		}
	}
}
