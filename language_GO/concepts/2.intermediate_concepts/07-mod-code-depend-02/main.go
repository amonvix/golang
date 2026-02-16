package main

import (
	"fmt"

	"github.com/amonvix/dogao"
)

func main() {
	s1 := dogao.Latir()
	s2 := dogao.Latidos()
	fmt.Println(s1)
	fmt.Println(s2)

	s3 := dogao.LatidoAlto()
	s4 := dogao.LatidoAltos()
	fmt.Println(s3)
	fmt.Println(s4)

	// also like this
	fmt.Println(dogao.Latir())
	fmt.Println(dogao.Latidos())
}
