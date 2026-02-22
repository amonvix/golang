package main

// create a loop that runs 5 times
// nest a loop withing the first loop that also prints 5 times
// print something in each loop to illustrate qhat is occuring

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("External Loop %v\n", i)
		for j := 0; j < 5; j++ {
			fmt.Printf("Internal Loop number %v\n", j)
		}
	}
}
