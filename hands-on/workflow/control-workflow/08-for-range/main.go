package main

import "fmt"

func main() {
	//  for range loop
	// data structures - slice
	list_iteration := []int{42, 43, 44, 45, 46, 47}
	for i, v := range list_iteration {
		fmt.Println("ranging over a slice", i, v)
	}
	//  for range loop
	// data structures - slice
	m := map[string]int{
		"James":      32,
		"Moneypenny": 27,
	}
	for k, v := range m {
		fmt.Println("ranging over a map", k, v)
	}
}
