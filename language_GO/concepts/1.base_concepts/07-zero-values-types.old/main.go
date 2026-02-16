package main // Define the main package.

import "fmt" // Import fmt for formatted I/O.

var zero int // Declare a zero-valued int.

func main() { // Program entry point.
	fmt.Println(zero) // Print the default zero value.

	z := 42        // Create a variable with short declaration.
	fmt.Println(z) // Print the value of z.

	a, b := 43, "golang" // Initialize multiple variables.
	fmt.Println(a, b)    // Print a and b.

	var x float32 = 42.42                          // Declare a float32 with explicit type.
	fmt.Printf("%v is of this Type of %T\n", x, x) // Print value and type of x.

	e, f, _ := 45, 46, 47 // Discard the third value with blank identifier.
	fmt.Println(e, f)     // Print e and f.
}
