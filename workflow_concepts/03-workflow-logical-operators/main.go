package main

import "fmt"

func main() {

	x := 40
	y := 5
	fmt.Printf("\n x=%v \n y=%v\n", x, y)

	// CONDITIONAL
	//  if statements
	//  switch statements

	if x < 42 && y < 42 {
		fmt.Println("both are less than the meaning of life")
	}

	if x > 30 || x < 42 {
		fmt.Println("x is getting closer to the meaning of life")
	}

	if x != 42 {
		fmt.Println("x is not 42")
	}
}
