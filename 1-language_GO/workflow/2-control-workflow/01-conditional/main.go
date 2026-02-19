package main

import "fmt"

func main() {

	x := 45
	y := 5
	fmt.Printf(" x=%v \n y=%v", x, y)

	// CONDITIONAL
	//  if statements
	//  switch statements

	if x < 42 {
		fmt.Println("\nLess than the meaning of life")
	}
	if x < 42 {
		fmt.Println("\nLess than the meaning of life")
	} else {
		fmt.Println("\nequal to, or greater than, the meaning of life")
	}

	if x < 42 {
		fmt.Println("\nLess than the meaning of life")
	} else if x == 42 {
		fmt.Println("\nequal to the meaning of life")
	} else {
		fmt.Println("\nGreater than, the meaning of life")
	}

	if x%2 == 0 {
		fmt.Printf("\n%v é par\n", x)
	} else {
		fmt.Printf("\n%v é ímpar\n", x)
	}
	/*
		"If" statements specify the conditional execution of two branches
		according to the value of boolean expression. If the expression evaluates
		to true, the if" branch is executed, otherwise, if present, the else "branch" is executed
	*/
	//  https://go.dev/ref/spec#If_statements

}
