package main

// how to import a exernal function. (having in consideration that a function only can be exported if it beggins with a captalize letter)

import (
	"fmt"

	"github.com/amonvix/dogao"
)

func main() {
	fmt.Println("Hello Gophers")
	s1 := dogao.Latir()
	s2 := dogao.Latidos()
	s3 := dogao.LatidoAlto()
	s4 := dogao.LatidoAltos()

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	dogao.From13()
}
