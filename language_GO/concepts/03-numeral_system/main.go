package main

import (
	"fmt"
)

func main() {
	// Definition of two variables with the same mathematical value.
	y := 42
	z := 42.0

	// With Prinf, the value type can be showed using a Fotmat Specifier.
	fmt.Printf("%v of type %T \n", y, y)
	fmt.Printf("%v of type %T \n", z, z)

	// The next example is showing the type of value storaged in variable "m". The exit will print the exacly valeu of "m".
	var m float32 = 43.742
	fmt.Printf("%v of type %T \n", m, m)
	
	// This last example, show how bigger a float64 can be, if precision in calculus is needed. 15 number before the dot.
	z = float64(m)
	fmt.Printf("%v of type %T \n", z, z)

}
