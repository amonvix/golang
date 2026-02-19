package main

import (
	"fmt"

	"github.com/amonvix/dogao"
)

// main is the entry point of the application.
// It demonstrates usage of various functions from the dogao package,
// printing their returned values to standard output.
func main() {
	s1 := dogao.Latir()
	s2 := dogao.Latidos()
	fmt.Println(s1)
	fmt.Println(s2)

	s3 := dogao.LatidoAlto()
	s4 := dogao.LatidoAltos()
	fmt.Println(s3)
	fmt.Println(s4)

	dogao.From11()

	// Directly print results of dogao functions without intermediate variables.
	fmt.Println(dogao.Latir())
	fmt.Println(dogao.LatidoAltos())
}
