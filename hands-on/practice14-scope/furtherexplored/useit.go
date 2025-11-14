package furtherexplored

import "fmt"

// this is also "package block" scope
// this is exported the identified "Fascinating" is capitalized

func Fascinating() {
	fmt.Println("Planet Speed", planetSpeed)

	planetRotationSpeed := 1000
	fmt.Println("Planet spinnning Speed", planetRotationSpeed)
}
