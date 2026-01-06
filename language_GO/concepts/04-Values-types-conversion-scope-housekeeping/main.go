package main

import "fmt"

var x int = 43

// Type inference can be determinated by the type of data storage in the variable
func main() {
	v := 42
	i := 88
	f := 3.142
	g := 0.867 + 5i
	y := 44

	// Cheking up the information provided by each variable we will see.
	fmt.Printf("%v is of Type %T\n",x, x) // numeral int is printed (here the variable is accesible inside the function).
	fmt.Printf("%v is of Type %T\n",y, y) // numeral int is printed (here the variable can be accesible only inside the function).
	fmt.Printf("%v is of Type %T\n",v, v) // numeral int is printed.
	fmt.Printf("%v is of Type %T\n",i, i) // numeral int is printed.
	fmt.Printf("%v is of Type %T\n",f, f) // decimal float64 is printed.
	fmt.Printf("%v is of Type %T\n",g, g) // because of the existance of the letter "i", representing an imaginary number.
	// fmt.Println(z)	// This usage is not permitted according to the scope of the code, and will throw an Build Error: undefined.
	out()	// Only through the function usage, the variable "z" can be accessed.
}

// Function created to explain the scope coverage
func out(){

	z := 45 
	fmt.Printf("%v is of Type %T\n",z, z)	// numeral int is printed
	}	
