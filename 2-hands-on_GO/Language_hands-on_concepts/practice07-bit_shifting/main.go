package main

// Now that we compreend the bit wise, use bit shifting in the numbers above

import "fmt"

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
