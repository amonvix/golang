package main

// The same as the last exercice, but with the external function created by the student

import (
	"fmt"

	goimports "github.com/amonvix/goimport"
)

var char string

func main() {
	char = "Hero"
	fmt.Printf("%s found an enemy and used Attack\n", char)
	goimports.Attack()
	fmt.Printf("%s found an enemy and used Defend\n", char)
	goimports.Defend()
	fmt.Printf("%s found an enemy and used Dodge\n", char)
	goimports.Dodge()

}
