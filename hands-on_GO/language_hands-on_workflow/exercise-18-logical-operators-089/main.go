package main

import "fmt"

/*
What do these print:

fmt.Println(true && true)
fmt.Println(true && false)g
fmt.Println(true || true)
fmt.Println(true || false)
fmt.Println(!true)
*/

func main() {
	fmt.Printf("Explaining the funcionality of all statement above")

	fmt.Println(true && true)  // not possible to use, because we have double & that means that is necessary both statements to continue
	fmt.Println(true && false) // possible, but not logical or easy to understand, since you will need a value true and false at same time
	fmt.Println(true || true)  // same patern than the first one. Pipes are used to test if one of the conditions happen. in this case both are equal
	fmt.Println(true || false) // This one is a real example of logical opertators usage, the program will keep rolling if one of the cases ocours
	fmt.Println(!true)         // The simple one. The negative form of true is not true, represented by ! simbol before the word
}
