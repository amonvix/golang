package main // Declare the main package.

import "fmt" // Import fmt for formatted I/O.

func main() { // Program entry point.
	//  for range loop
	// data structures - slice
	list_iteration := []int{42, 43, 44, 45, 46, 47} // Create a slice of ints.
	for i, v := range list_iteration { // Range over the slice.
		fmt.Println("ranging over a slice", i, v) // Print index and value.
	} // End range loop.
	//  for range loop
	// data structures - slice
	m := map[string]int{ // Create a map of names to values.
		"James":      32, // Add James.
		"Moneypenny": 27, // Add Moneypenny.
	} // End map literal.
	for k, v := range m { // Range over the map.
		fmt.Println("ranging over a map", k, v) // Print key and value.
	} // End range loop.
} // End main.
