package main // Define the main package for an executable program.

import "fmt" // Import the fmt package for formatted I/O.

const ( 
	Big   = 1 << 100 // Define a huge constant by shifting 1 left 100 bits.
	Small = Big >> 99 // Define a smaller constant by shifting Big right 99 bits.
) 

func needInt(x int) int { return x*10 + 1 } // Compute an int-based expression and return it.
func needFloat(x float64) float64 { // Declare a function that works with float64 input.
	return x * 0.1 // Scale the input by 0.1 and return it.
} 

func main() { 
	fmt.Println(needInt(Small)) // Print the int result using Small.
	fmt.Println(needFloat(Small)) // Print the float result using Small.
	fmt.Println(needFloat(Big)) // Print the float result using Big.
} 
