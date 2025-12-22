package main

import (
	"fmt"
	"math/rand"
)

func main() {
	y := 251
	x := rand.Intn(y)
	fmt.Println(x)

	if x >= 0 && x <= 100 {
		fmt.Printf("%v is between 0 and 100", x)
	} else if x > 100 && x <= 200 {
		fmt.Printf("%v is between 101 and 200", x)
	} else {
		fmt.Printf("%v is between 201 and 250", x)
	}

	fmt.Println("\n######################")
	fmt.Println("Demonstrate the behavior included using the numbers **0–3**")
	z := rand.Intn(3)
	fmt.Println(z)
	fmt.Println("The last iten of the pseudo-rando is never accessed!")
	fmt.Println("######################")
}
