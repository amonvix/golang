package furtherexplored

import "fmt"

// Fascinating prints information related to planetary speeds.
// It demonstrates usage of package-level and local variables.
func Fascinating() {
	fmt.Println("Planet Speed", planetSpeed)

	// Local variable representing the planet's rotation speed in arbitrary units.
	planetRotationSpeed := 1000
	fmt.Println("Planet spinnning Speed", planetRotationSpeed)
}