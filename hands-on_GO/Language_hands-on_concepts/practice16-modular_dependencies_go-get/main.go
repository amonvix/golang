package main

// learning the way to "get" functions from external sources

import (
	"fmt"

	"github.com/amonvix/dogao"
)

func main() {
	s1 := dogao.Latir()
	s2 := dogao.Latidos()
	fmt.Println(s1)
	fmt.Println(s2)

	s3 := dogao.LatidoAlto()  //This is an import that comes from the dog repository
	s4 := dogao.LatidoAltos() //This is an import that comes from the dog repository
	fmt.Println(s3)
	fmt.Println(s4)

}
