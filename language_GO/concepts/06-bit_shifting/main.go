package main // Declare the main package for an executable program.

import "fmt" // Import fmt for formatted I/O.

func main() { // Program entry point.
	fmt.Printf("%d \t %b \n", 1, 1) // Print 1 in decimal and binary.
	fmt.Printf("%d \t %b \n", 1<<1, 1<<1) // Print 1 shifted left by 1 in decimal and binary.
	fmt.Printf("%d \t %b \n", 1<<2, 1<<2) // Print 1 shifted left by 2 in decimal and binary.
	fmt.Printf("%d \t %b \n", 1<<3, 1<<3) // Print 1 shifted left by 3 in decimal and binary.
	fmt.Printf("%d \t %b \n", 1<<4, 1<<4) // Print 1 shifted left by 4 in decimal and binary.
	fmt.Printf("%d \t %b \n", 1<<5, 1<<5) // Print 1 shifted left by 5 in decimal and binary.
	fmt.Printf("%d \t %b \n", 1<<6, 1<<6) // Print 1 shifted left by 6 in decimal and binary.
	fmt.Printf("%d \t %b \n", 1<<7, 1<<7) // Print 1 shifted left by 7 in decimal and binary.
} // End of main.
