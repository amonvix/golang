package main

import "fmt"

// use modulus and continue statement in a loop to print all ODD numbers

func main() {

	for i := 100; i > 1; i-- {
		if i%2 != 0 {
			fmt.Println(i)
		}

	}

}
