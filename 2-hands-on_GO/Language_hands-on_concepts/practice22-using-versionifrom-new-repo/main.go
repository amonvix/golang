package main

import (
	"fmt"

	"github.com/amonvix/dogao"
)

// main is the entry point of the application.
// It demonstrates usage of various functions from the dogao package
// and outputs their results to standard output.
func main() {
	fmt.Println("Hello Gophers")

	// Retrieve and print different dog-related sounds from the dogao package.
	s1 := dogao.Latir()
	s2 := dogao.Latidos()
	s3 := dogao.LatidoAlto()
	s4 := dogao.LatidoAltos()

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

	// Invoke From11 function from dogao package, side effects or output depend on implementation.
	dogao.From11()
}