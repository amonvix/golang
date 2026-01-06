package main

// Show tipes of inferences in the follow numbers

import "fmt"

//change the original value of v in the example bellow replacing the nunber by
// i := 88
// f := 3.142
// g := 0.867 + 05i

func main() {
	v := 42
	i := 88
	f := 3.142
	g := 0.867 + 5i
	fmt.Printf("v is of Type %T\n", v)
	fmt.Printf("v is of Type %T\n", i)
	fmt.Printf("v is of Type %T\n", f)
	fmt.Printf("v is of Type %T\n", g)
}
