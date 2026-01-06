package main // Declare the main package.

import ( // Start import block.
	"fmt" // Import fmt for formatted I/O.

	"github.com/amonvix/dogao" // Import dogao module.
) // End import block.

func main() { // Program entry point.
	s1 := dogao.Latir() // Call Latir from dogao.
	s2 := dogao.Latidos() // Call Latidos from dogao.
	fmt.Println(s1) // Print s1.
	fmt.Println(s2) // Print s2.

	s3 := dogao.LatidoAlto() // Call LatidoAlto from dogao.
	s4 := dogao.LatidoAltos() // Call LatidoAltos from dogao.
	fmt.Println(s3) // Print s3.
	fmt.Println(s4) // Print s4.

	dogao.From11() // Call From11 from dogao.

	// also like this
	fmt.Println(dogao.Latir()) // Print Latir result.
	fmt.Println(dogao.LatidoAltos()) // Print LatidoAltos result.
} // End main.
