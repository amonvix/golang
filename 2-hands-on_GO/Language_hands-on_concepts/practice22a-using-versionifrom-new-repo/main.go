package main

// Package main coordinates character combat actions by invoking external functions
// from the goimport package, managing the sequence of attack, defense, and evasion.

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