package main

// show how to declare all the types of variables saw until now.

import "fmt"

func main() {
	a := 42
	fmt.Println(a)

	b, c, d, _, f := 0, 1, 2, 3, "happiness"
	fmt.Println(b, c, d, f)

	var h int
	fmt.Println(h)

	h = 666
	fmt.Println(h)

	// examples of wrong usage of the code
	/*
		b, c, d, e, f:= 0, 1 ,2 ,3
		fmt.Println(b, c, d)		<== An error is throw an error to show an unused variable

	*/

	// this type of usage does work
	/*

		var g int
		fmt.Println(g)
		g = 43
		fmt.Println(g)

		/// DECLARE a variable is used when you want hold a specific TYPE of VALUE
		/// then assign a VALUE of that TYPE to the variable
		/// initializing a variable
		var h int = 44
		fmt.Println(h)

	*/

}
