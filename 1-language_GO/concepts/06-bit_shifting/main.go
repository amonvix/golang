package main

import "fmt"

// main outputs decimal and binary representations of powers of two from 2^0 to 2^7.
// It uses bit shifting to generate the values and formats the output in aligned columns.
func main() {
	fmt.Printf("%d \t %b \n", 1, 1)
	fmt.Printf("%d \t %b \n", 1<<1, 1<<1)
	fmt.Printf("%d \t %b \n", 1<<2, 1<<2)
	fmt.Printf("%d \t %b \n", 1<<3, 1<<3)
	fmt.Printf("%d \t %b \n", 1<<4, 1<<4)
	fmt.Printf("%d \t %b \n", 1<<5, 1<<5)
	fmt.Printf("%d \t %b \n", 1<<6, 1<<6)
	fmt.Printf("%d \t %b \n", 1<<7, 1<<7)
}