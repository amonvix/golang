package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// this firts exercise is to collect the shasum of this file (SNOWY-EVENING.txt)
// and compare it with the expected shasum value after inserting a new character
// at the end of the file. The expected shasum value is:

func main() {
	x := exec.Command("shasum", "-a", "256", "./SNOWY-EVENING.txt")

	out, err := x.Output()
	if err != nil {
		panic(err)
	}

	resultado1 := strings.TrimSpace(string(out))
	fmt.Println("SHA256 of original file:", resultado1)

	// Now, let's add a new character at the end of the file
	x = exec.Command("bash", "-c", "echo '.' >> ./SNOWY-EVENING.txt")
	err = x.Run()
	if err != nil {
		panic(err)
	}
	// Now, let's get the shasum again
	x = exec.Command("shasum", "-a", "256", "./SNOWY-EVENING.txt")
	out, err = x.Output()
	if err != nil {
		panic(err)
	}
	resultado2 := strings.TrimSpace(string(out))
	fmt.Println("SHA256 of modified file:", resultado2)
	fmt.Printf("Original: %s\nModified: %s\n", resultado1, resultado2)
}
