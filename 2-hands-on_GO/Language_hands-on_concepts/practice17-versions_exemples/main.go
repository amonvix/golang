package main

// Working with versioning in the same original external file

import (
	"fmt"

	"github.com/amonvix/dogao"
)

func main() {

	dogao.From11()
	s1 := dogao.Latir()
	s2 := dogao.Latidos()
	fmt.Println(s1)
	fmt.Println(s2)

	s3 := dogao.LatidoAlto()  //This is an import that comes from the dog repository
	s4 := dogao.LatidoAltos() //This is an import that comes from the dog repository
	fmt.Println(s3)
	fmt.Println(s4)

	// also like this
	fmt.Println(dogao.Latir())
	fmt.Println(dogao.LatidoAltos())

}
