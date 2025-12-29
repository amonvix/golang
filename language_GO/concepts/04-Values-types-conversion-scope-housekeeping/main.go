package main

import "fmt"

// Type inference can be determinated by the type of data storage in the variable
func main() {
	v := 42
	i := 88
	f := 3.142
	g := 0.867 + 5i

	// Cheking up the information provided by each variable we will see.
	fmt.Printf("v is of Type %T\n", v) // numeral int is printed
	fmt.Printf("v is of Type %T\n", i) // numeral int is printed
	fmt.Printf("v is of Type %T\n", f) // decimal float64 is printed
	fmt.Printf("v is of Type %T\n", g) // because of the existance of the letter "i", representing an imaginary number.
}
