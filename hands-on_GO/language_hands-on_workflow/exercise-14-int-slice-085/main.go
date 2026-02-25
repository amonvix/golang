package main

import "fmt"

// Rules:
// below is the code to create a data structure called a slice of ints.
// put this code into a program
//  "xi := []int{42, 43, 44, 45, 46, 47}"
// use a range loop to print each value and the index position of each value

func main() {
	xi := []int{42, 43, 44, 45, 46, 47}

	for i, v := range xi {
		fmt.Printf("The position of the value %v is\the %v in the iteration\n", v, i)
	}
}
