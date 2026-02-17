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

	// also like this
	fmt.Println(dogao.LatidoAlto())
	fmt.Println(dogao.LatidoAltos())
}
