package main

import "fmt"

func main() {
	number := 42                                       // a value is declared in the variable number
	fmt.Printf("42 as binary is, %b \n", number)       // here we print the value in binary using the Integer simbol %b
	fmt.Printf("42 as hexadecimal is, %#x \n", number) // here we print the value in hexadecimal using the Integer simbol %#x
}
