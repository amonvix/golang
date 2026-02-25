package main

import "fmt"

/*
Use the code from the previous exercise
add this code to print a single value stored in the map

age := m["James"]
fmt.prinLn(age)

now use similar code to use the lookup of "Q"and print that value
	do that
now use the "comma ok" idiom to test whether Q is actually stored un the map, then print out a statement if is not stored in the map
	hint: check efective go for yhr "comma ok" idiom
*/

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for k, v := range m {
		fmt.Printf("The value %v storage the information %v\n", k, v)
	}

	fmt.Printf("Verification of the content in 'm\n'")
	age := m["James"]
	fmt.Println(age)

	fmt.Printf("verification of in content of 'Q'\n")
	if m1, ok := m["Q"]; ok {
		fmt.Println(m1)
	} else {
		fmt.Println("Q is not stored in the map")
	}

}
