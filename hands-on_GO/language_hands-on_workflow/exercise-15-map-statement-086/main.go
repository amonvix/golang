package main

import "fmt"

/*
The below is the code to create a data structure called map.
Put this code into a program
m := map[string]int{
		"James": 42,
  		"Moneypenny": 32
}
Use a for range loop to print each value and the key associated with each value
*/

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for k, v := range m {
		fmt.Printf("The %v have %v years old\n", k, v)
	}

}
