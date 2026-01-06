package main

// Show that you learn the types of zero and non informed values

import "fmt"

var zero int

func main() {
	fmt.Print(zero)

	z := 42
	fmt.Println(z)

	a, b := 43, "golang"
	fmt.Println(a, b)

	var x float32 = 42.42
	fmt.Printf("%v is of this Type of %T\n", x, x)

	e, f, _ := 45, 46, 47
	fmt.Println(e, f)
}

/*
var for zero value
short declaration operator
multiple initializations
var when specificity is required
blank identifier

*/
